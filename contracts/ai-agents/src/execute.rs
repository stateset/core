use cosmwasm_std::{
    DepsMut, Env, MessageInfo, Response, Coin, Addr, Uint128, 
    BankMsg, CosmosMsg, StdError, Order, Storage,
};
use cw_utils::must_pay;
use serde_json;

use crate::error::ContractError;
use crate::msg::{
    Transfer, ServiceStatus, TransactionType, MessageType, 
    PurchaseOrderItem, PaymentTerms, PurchaseOrderStatus, InvoiceLineItem, ItemReceipt,
};
use crate::state::{
    Config, Agent, AgentWallet, Service, TransactionRecord, AgentMessage, MessageResponse,
    PurchaseOrder, Invoice, Receipt, AgentFinancials,
    CONFIG, agents, AGENT_WALLETS, AGENT_ADDRESS_TO_ID, 
    SERVICES, AGENT_SERVICES, TRANSACTIONS, AGENT_TRANSACTIONS,
    MESSAGES, AGENT_MESSAGES, CONFIG_MESSAGE_COUNTER,
    PURCHASE_ORDERS, AGENT_PURCHASE_ORDERS, INVOICES, AGENT_INVOICES,
    RECEIPTS, CONFIG_PO_COUNTER, CONFIG_INVOICE_COUNTER, AGENT_FINANCIALS,
};
use crate::utils::{generate_agent_id, generate_service_id, validate_agent_name};

// Agent Management Functions

pub fn register_agent(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    name: String,
    description: String,
    capabilities: Vec<String>,
    service_endpoints: Vec<String>,
    initial_balance: Option<Coin>,
) -> Result<Response, ContractError> {
    let mut config = CONFIG.load(deps.storage)?;
    
    // Validate agent name
    validate_agent_name(&name)?;
    
    // Check if sender already has an agent
    if AGENT_ADDRESS_TO_ID.has(deps.storage, info.sender.clone()) {
        return Err(ContractError::AgentAlreadyExists {
            agent_id: "Agent already registered for this address".to_string(),
        });
    }
    
    // Generate unique agent ID
    let agent_id = generate_agent_id(&info.sender, config.next_agent_id);
    config.next_agent_id += 1;
    
    // Create agent wallet address (deterministic based on agent_id)
    let wallet_address = deps.api.addr_canonicalize(&format!("agent_{}", agent_id))
        .map_err(|_| StdError::generic_err("Failed to create wallet address"))?;
    let wallet_address = deps.api.addr_humanize(&wallet_address)?;
    
    // Create agent
    let agent = Agent {
        agent_id: agent_id.clone(),
        owner: info.sender.clone(),
        name,
        description,
        capabilities,
        service_endpoints,
        wallet_address: wallet_address.clone(),
        is_active: true,
        created_at: env.block.time.seconds(),
        last_active: env.block.time.seconds(),
        total_services_provided: 0,
        total_services_requested: 0,
        reputation_score: 100, // Start with base reputation
    };
    
    // Initialize wallet
    let mut wallet = AgentWallet {
        agent_id: agent_id.clone(),
        balance: Coin {
            denom: config.stablecoin_denom.clone(),
            amount: Uint128::zero(),
        },
        locked_balance: Coin {
            denom: config.stablecoin_denom.clone(),
            amount: Uint128::zero(),
        },
    };
    
    // Handle initial funding
    let mut messages = vec![];
    if let Some(initial_coin) = initial_balance {
        // Verify the coin matches stablecoin denom
        if initial_coin.denom != config.stablecoin_denom {
            return Err(ContractError::InvalidPayment {});
        }
        
        // Verify payment was sent
        must_pay(&info, &config.stablecoin_denom)?;
        
        wallet.balance = initial_coin.clone();
        
        // Record initial deposit transaction
        record_transaction(
            deps.storage,
            &mut config,
            TransactionType::Deposit,
            None,
            Some(agent_id.clone()),
            initial_coin,
            Some("Initial agent funding".to_string()),
            env.block.time.seconds(),
        )?;
    }
    
    // Save all state
    CONFIG.save(deps.storage, &config)?;
    agents().save(deps.storage, agent_id.clone(), &agent)?;
    AGENT_WALLETS.save(deps.storage, agent_id.clone(), &wallet)?;
    AGENT_ADDRESS_TO_ID.save(deps.storage, info.sender, &agent_id)?;
    
    Ok(Response::new()
        .add_messages(messages)
        .add_attribute("method", "register_agent")
        .add_attribute("agent_id", agent_id)
        .add_attribute("wallet_address", wallet_address)
        .add_attribute("initial_balance", wallet.balance.amount))
}

pub fn update_agent(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    agent_id: String,
    name: Option<String>,
    description: Option<String>,
    capabilities: Option<Vec<String>>,
    service_endpoints: Option<Vec<String>>,
) -> Result<Response, ContractError> {
    let mut agent = agents().load(deps.storage, agent_id.clone())?;
    
    // Verify ownership
    if agent.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Update fields if provided
    if let Some(new_name) = name {
        validate_agent_name(&new_name)?;
        agent.name = new_name;
    }
    if let Some(new_desc) = description {
        agent.description = new_desc;
    }
    if let Some(new_caps) = capabilities {
        agent.capabilities = new_caps;
    }
    if let Some(new_endpoints) = service_endpoints {
        agent.service_endpoints = new_endpoints;
    }
    
    agent.last_active = env.block.time.seconds();
    
    agents().save(deps.storage, agent_id.clone(), &agent)?;
    
    Ok(Response::new()
        .add_attribute("method", "update_agent")
        .add_attribute("agent_id", agent_id))
}

pub fn deactivate_agent(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    agent_id: String,
) -> Result<Response, ContractError> {
    let mut agent = agents().load(deps.storage, agent_id.clone())?;
    
    // Verify ownership
    if agent.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Check if agent has locked balance
    let wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
    if !wallet.locked_balance.amount.is_zero() {
        return Err(ContractError::CustomError {
            msg: "Cannot deactivate agent with locked balance".to_string(),
        });
    }
    
    agent.is_active = false;
    agent.last_active = env.block.time.seconds();
    
    agents().save(deps.storage, agent_id.clone(), &agent)?;
    
    Ok(Response::new()
        .add_attribute("method", "deactivate_agent")
        .add_attribute("agent_id", agent_id))
}

// Wallet Management Functions

pub fn fund_agent(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    agent_id: String,
) -> Result<Response, ContractError> {
    let config = CONFIG.load(deps.storage)?;
    let agent = agents().load(deps.storage, agent_id.clone())?;
    
    if !agent.is_active {
        return Err(ContractError::AgentNotActive { agent_id: agent_id.clone() });
    }
    
    // Verify payment
    let payment = must_pay(&info, &config.stablecoin_denom)?;
    
    // Update wallet balance
    let mut wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
    wallet.balance.amount = wallet.balance.amount.checked_add(payment)?;
    AGENT_WALLETS.save(deps.storage, agent_id.clone(), &wallet)?;
    
    // Record transaction
    let mut config_mut = config;
    record_transaction(
        deps.storage,
        &mut config_mut,
        TransactionType::Deposit,
        None,
        Some(agent_id.clone()),
        Coin {
            denom: config_mut.stablecoin_denom.clone(),
            amount: payment,
        },
        Some(format!("Funding from {}", info.sender)),
        env.block.time.seconds(),
    )?;
    CONFIG.save(deps.storage, &config_mut)?;
    
    Ok(Response::new()
        .add_attribute("method", "fund_agent")
        .add_attribute("agent_id", agent_id)
        .add_attribute("amount", payment))
}

pub fn withdraw_from_agent(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    agent_id: String,
    amount: Coin,
    recipient: String,
) -> Result<Response, ContractError> {
    let config = CONFIG.load(deps.storage)?;
    let agent = agents().load(deps.storage, agent_id.clone())?;
    
    // Verify ownership
    if agent.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Verify denom
    if amount.denom != config.stablecoin_denom {
        return Err(ContractError::InvalidPayment {});
    }
    
    // Check available balance
    let mut wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
    let available = wallet.balance.amount.saturating_sub(wallet.locked_balance.amount);
    if amount.amount > available {
        return Err(ContractError::InsufficientBalance {
            required: amount.amount.to_string(),
            available: available.to_string(),
        });
    }
    
    // Update wallet balance
    wallet.balance.amount = wallet.balance.amount.checked_sub(amount.amount)?;
    AGENT_WALLETS.save(deps.storage, agent_id.clone(), &wallet)?;
    
    // Validate recipient address
    let recipient_addr = deps.api.addr_validate(&recipient)?;
    
    // Create bank send message
    let msg = BankMsg::Send {
        to_address: recipient_addr.to_string(),
        amount: vec![amount.clone()],
    };
    
    // Record transaction
    let mut config_mut = config;
    record_transaction(
        deps.storage,
        &mut config_mut,
        TransactionType::Withdrawal,
        Some(agent_id.clone()),
        None,
        amount.clone(),
        Some(format!("Withdrawal to {}", recipient_addr)),
        env.block.time.seconds(),
    )?;
    CONFIG.save(deps.storage, &config_mut)?;
    
    Ok(Response::new()
        .add_message(msg)
        .add_attribute("method", "withdraw_from_agent")
        .add_attribute("agent_id", agent_id)
        .add_attribute("amount", amount.amount)
        .add_attribute("recipient", recipient_addr))
}

// Agent-to-Agent Transfer Functions

pub fn agent_transfer(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    from_agent_id: String,
    to_agent_id: String,
    amount: Coin,
    memo: Option<String>,
) -> Result<Response, ContractError> {
    let config = CONFIG.load(deps.storage)?;
    
    // Verify from agent ownership
    let from_agent = agents().load(deps.storage, from_agent_id.clone())?;
    if from_agent.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Verify agents are different
    if from_agent_id == to_agent_id {
        return Err(ContractError::CircularTransfer {});
    }
    
    // Verify to agent exists and is active
    let mut to_agent = agents().load(deps.storage, to_agent_id.clone())?;
    if !to_agent.is_active {
        return Err(ContractError::AgentNotActive { agent_id: to_agent_id.clone() });
    }
    
    // Verify denom
    if amount.denom != config.stablecoin_denom {
        return Err(ContractError::InvalidPayment {});
    }
    
    // Update wallets
    let mut from_wallet = AGENT_WALLETS.load(deps.storage, from_agent_id.clone())?;
    let available = from_wallet.balance.amount.saturating_sub(from_wallet.locked_balance.amount);
    if amount.amount > available {
        return Err(ContractError::InsufficientBalance {
            required: amount.amount.to_string(),
            available: available.to_string(),
        });
    }
    
    from_wallet.balance.amount = from_wallet.balance.amount.checked_sub(amount.amount)?;
    AGENT_WALLETS.save(deps.storage, from_agent_id.clone(), &from_wallet)?;
    
    let mut to_wallet = AGENT_WALLETS.load(deps.storage, to_agent_id.clone())?;
    to_wallet.balance.amount = to_wallet.balance.amount.checked_add(amount.amount)?;
    AGENT_WALLETS.save(deps.storage, to_agent_id.clone(), &to_wallet)?;
    
    // Update last active timestamps
    let mut from_agent_mut = from_agent;
    from_agent_mut.last_active = env.block.time.seconds();
    agents().save(deps.storage, from_agent_id.clone(), &from_agent_mut)?;
    
    to_agent.last_active = env.block.time.seconds();
    agents().save(deps.storage, to_agent_id.clone(), &to_agent)?;
    
    // Record transaction
    let mut config_mut = config;
    record_transaction(
        deps.storage,
        &mut config_mut,
        TransactionType::Transfer,
        Some(from_agent_id.clone()),
        Some(to_agent_id.clone()),
        amount.clone(),
        memo,
        env.block.time.seconds(),
    )?;
    CONFIG.save(deps.storage, &config_mut)?;
    
    Ok(Response::new()
        .add_attribute("method", "agent_transfer")
        .add_attribute("from", from_agent_id)
        .add_attribute("to", to_agent_id)
        .add_attribute("amount", amount.amount))
}

pub fn batch_agent_transfer(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    from_agent_id: String,
    transfers: Vec<Transfer>,
) -> Result<Response, ContractError> {
    let config = CONFIG.load(deps.storage)?;
    
    // Verify from agent ownership
    let from_agent = agents().load(deps.storage, from_agent_id.clone())?;
    if from_agent.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Validate all transfers and calculate total
    let mut total_amount = Uint128::zero();
    for transfer in &transfers {
        // Verify denom
        if transfer.amount.denom != config.stablecoin_denom {
            return Err(ContractError::InvalidPayment {});
        }
        
        // Verify recipient exists and is active
        let to_agent = agents().load(deps.storage, transfer.to_agent_id.clone())?;
        if !to_agent.is_active {
            return Err(ContractError::AgentNotActive { 
                agent_id: transfer.to_agent_id.clone() 
            });
        }
        
        // Check for self-transfer
        if from_agent_id == transfer.to_agent_id {
            return Err(ContractError::CircularTransfer {});
        }
        
        total_amount = total_amount.checked_add(transfer.amount.amount)?;
    }
    
    // Check sufficient balance
    let mut from_wallet = AGENT_WALLETS.load(deps.storage, from_agent_id.clone())?;
    let available = from_wallet.balance.amount.saturating_sub(from_wallet.locked_balance.amount);
    if total_amount > available {
        return Err(ContractError::InsufficientBalance {
            required: total_amount.to_string(),
            available: available.to_string(),
        });
    }
    
    // Execute all transfers
    let mut config_mut = config;
    for transfer in transfers {
        // Update recipient wallet
        let mut to_wallet = AGENT_WALLETS.load(deps.storage, transfer.to_agent_id.clone())?;
        to_wallet.balance.amount = to_wallet.balance.amount.checked_add(transfer.amount.amount)?;
        AGENT_WALLETS.save(deps.storage, transfer.to_agent_id.clone(), &to_wallet)?;
        
        // Update recipient agent last active
        let mut to_agent = agents().load(deps.storage, transfer.to_agent_id.clone())?;
        to_agent.last_active = env.block.time.seconds();
        agents().save(deps.storage, transfer.to_agent_id.clone(), &to_agent)?;
        
        // Record transaction
        record_transaction(
            deps.storage,
            &mut config_mut,
            TransactionType::Transfer,
            Some(from_agent_id.clone()),
            Some(transfer.to_agent_id.clone()),
            transfer.amount.clone(),
            transfer.memo,
            env.block.time.seconds(),
        )?;
    }
    
    // Update sender wallet
    from_wallet.balance.amount = from_wallet.balance.amount.checked_sub(total_amount)?;
    AGENT_WALLETS.save(deps.storage, from_agent_id.clone(), &from_wallet)?;
    
    // Update sender agent last active
    let mut from_agent_mut = from_agent;
    from_agent_mut.last_active = env.block.time.seconds();
    agents().save(deps.storage, from_agent_id.clone(), &from_agent_mut)?;
    
    CONFIG.save(deps.storage, &config_mut)?;
    
    Ok(Response::new()
        .add_attribute("method", "batch_agent_transfer")
        .add_attribute("from", from_agent_id)
        .add_attribute("total_amount", total_amount)
        .add_attribute("transfer_count", transfers.len().to_string()))
}

// Service Functions

pub fn request_service(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    requester_agent_id: String,
    provider_agent_id: String,
    service_type: String,
    payment: Coin,
    parameters: String,
) -> Result<Response, ContractError> {
    let mut config = CONFIG.load(deps.storage)?;
    
    // Verify requester agent ownership
    let mut requester = agents().load(deps.storage, requester_agent_id.clone())?;
    if requester.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Verify both agents are active
    if !requester.is_active {
        return Err(ContractError::AgentNotActive { 
            agent_id: requester_agent_id.clone() 
        });
    }
    
    let mut provider = agents().load(deps.storage, provider_agent_id.clone())?;
    if !provider.is_active {
        return Err(ContractError::AgentNotActive { 
            agent_id: provider_agent_id.clone() 
        });
    }
    
    // Verify payment denom
    if payment.denom != config.stablecoin_denom {
        return Err(ContractError::InvalidPayment {});
    }
    
    // Check requester has sufficient balance
    let mut requester_wallet = AGENT_WALLETS.load(deps.storage, requester_agent_id.clone())?;
    let available = requester_wallet.balance.amount
        .saturating_sub(requester_wallet.locked_balance.amount);
    if payment.amount > available {
        return Err(ContractError::InsufficientBalance {
            required: payment.amount.to_string(),
            available: available.to_string(),
        });
    }
    
    // Lock payment in escrow
    requester_wallet.locked_balance.amount = requester_wallet.locked_balance.amount
        .checked_add(payment.amount)?;
    AGENT_WALLETS.save(deps.storage, requester_agent_id.clone(), &requester_wallet)?;
    
    // Generate service ID
    let service_id = generate_service_id(
        &requester_agent_id,
        &provider_agent_id,
        config.next_service_id
    );
    config.next_service_id += 1;
    
    // Create service record
    let service = Service {
        service_id: service_id.clone(),
        requester_agent_id: requester_agent_id.clone(),
        provider_agent_id: provider_agent_id.clone(),
        service_type,
        payment: payment.clone(),
        status: ServiceStatus::Pending,
        parameters,
        result: None,
        created_at: env.block.time.seconds(),
        completed_at: None,
        escrow_released: false,
    };
    
    SERVICES.save(deps.storage, service_id.clone(), &service)?;
    
    // Update agent statistics
    requester.total_services_requested += 1;
    requester.last_active = env.block.time.seconds();
    agents().save(deps.storage, requester_agent_id.clone(), &requester)?;
    
    provider.last_active = env.block.time.seconds();
    agents().save(deps.storage, provider_agent_id.clone(), &provider)?;
    
    // Save updated config
    CONFIG.save(deps.storage, &config)?;
    
    Ok(Response::new()
        .add_attribute("method", "request_service")
        .add_attribute("service_id", service_id)
        .add_attribute("requester", requester_agent_id)
        .add_attribute("provider", provider_agent_id)
        .add_attribute("payment", payment.amount))
}

pub fn complete_service(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    service_id: String,
    result: String,
) -> Result<Response, ContractError> {
    let mut config = CONFIG.load(deps.storage)?;
    let mut service = SERVICES.load(deps.storage, service_id.clone())?;
    
    // Verify service is pending
    if service.status != ServiceStatus::Pending {
        return Err(ContractError::InvalidServiceStatus { 
            service_id: service_id.clone() 
        });
    }
    
    // Verify caller is the provider
    let provider = agents().load(deps.storage, service.provider_agent_id.clone())?;
    if provider.owner != info.sender {
        return Err(ContractError::Unauthorized {});
    }
    
    // Update service status
    service.status = ServiceStatus::Completed;
    service.result = Some(result);
    service.completed_at = Some(env.block.time.seconds());
    service.escrow_released = true;
    SERVICES.save(deps.storage, service_id.clone(), &service)?;
    
    // Calculate fee
    let fee_amount = service.payment.amount
        .multiply_ratio(config.service_fee_percentage, 10000u128);
    let provider_payment = service.payment.amount.checked_sub(fee_amount)?;
    
    // Release escrow from requester
    let mut requester_wallet = AGENT_WALLETS.load(deps.storage, service.requester_agent_id.clone())?;
    requester_wallet.balance.amount = requester_wallet.balance.amount
        .checked_sub(service.payment.amount)?;
    requester_wallet.locked_balance.amount = requester_wallet.locked_balance.amount
        .checked_sub(service.payment.amount)?;
    AGENT_WALLETS.save(deps.storage, service.requester_agent_id.clone(), &requester_wallet)?;
    
    // Pay provider
    let mut provider_wallet = AGENT_WALLETS.load(deps.storage, service.provider_agent_id.clone())?;
    provider_wallet.balance.amount = provider_wallet.balance.amount
        .checked_add(provider_payment)?;
    AGENT_WALLETS.save(deps.storage, service.provider_agent_id.clone(), &provider_wallet)?;
    
    // Update agent statistics
    let mut provider_mut = provider;
    provider_mut.total_services_provided += 1;
    provider_mut.reputation_score = provider_mut.reputation_score.saturating_add(5); // Reward for completion
    provider_mut.last_active = env.block.time.seconds();
    agents().save(deps.storage, service.provider_agent_id.clone(), &provider_mut)?;
    
    // Record transactions
    record_transaction(
        deps.storage,
        &mut config,
        TransactionType::ServicePayment,
        Some(service.requester_agent_id.clone()),
        Some(service.provider_agent_id.clone()),
        Coin {
            denom: config.stablecoin_denom.clone(),
            amount: provider_payment,
        },
        Some(format!("Service {} completed", service_id)),
        env.block.time.seconds(),
    )?;
    
    if !fee_amount.is_zero() {
        record_transaction(
            deps.storage,
            &mut config,
            TransactionType::Fee,
            Some(service.requester_agent_id.clone()),
            None,
            Coin {
                denom: config.stablecoin_denom.clone(),
                amount: fee_amount,
            },
            Some(format!("Service fee for {}", service_id)),
            env.block.time.seconds(),
        )?;
    }
    
    CONFIG.save(deps.storage, &config)?;
    
    Ok(Response::new()
        .add_attribute("method", "complete_service")
        .add_attribute("service_id", service_id)
        .add_attribute("provider_payment", provider_payment)
        .add_attribute("fee", fee_amount))
}

pub fn refund_service(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    service_id: String,
    reason: String,
) -> Result<Response, ContractError> {
    let mut config = CONFIG.load(deps.storage)?;
    let mut service = SERVICES.load(deps.storage, service_id.clone())?;
    
    // Verify service is pending
    if service.status != ServiceStatus::Pending {
        return Err(ContractError::InvalidServiceStatus { 
            service_id: service_id.clone() 
        });
    }
    
    // Verify caller is either requester or provider
    let requester = agents().load(deps.storage, service.requester_agent_id.clone())?;
    let provider = agents().load(deps.storage, service.provider_agent_id.clone())?;
    
    if info.sender != requester.owner && info.sender != provider.owner {
        return Err(ContractError::Unauthorized {});
    }
    
    // Update service status
    service.status = ServiceStatus::Refunded;
    service.result = Some(reason.clone());
    service.completed_at = Some(env.block.time.seconds());
    service.escrow_released = true;
    SERVICES.save(deps.storage, service_id.clone(), &service)?;
    
    // Release escrow back to requester
    let mut requester_wallet = AGENT_WALLETS.load(deps.storage, service.requester_agent_id.clone())?;
    requester_wallet.locked_balance.amount = requester_wallet.locked_balance.amount
        .checked_sub(service.payment.amount)?;
    AGENT_WALLETS.save(deps.storage, service.requester_agent_id.clone(), &requester_wallet)?;
    
    // Update reputation if provider initiated refund
    if info.sender == provider.owner {
        let mut provider_mut = provider;
        provider_mut.reputation_score = provider_mut.reputation_score.saturating_sub(2);
        agents().save(deps.storage, service.provider_agent_id.clone(), &provider_mut)?;
    }
    
    // Record transaction
    record_transaction(
        deps.storage,
        &mut config,
        TransactionType::ServiceRefund,
        Some(service.provider_agent_id.clone()),
        Some(service.requester_agent_id.clone()),
        service.payment,
        Some(format!("Refund: {}", reason)),
        env.block.time.seconds(),
    )?;
    
    CONFIG.save(deps.storage, &config)?;
    
    Ok(Response::new()
        .add_attribute("method", "refund_service")
        .add_attribute("service_id", service_id)
        .add_attribute("reason", reason))
}

// Governance Functions

pub fn update_config(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    stablecoin_denom: Option<String>,
    min_agent_balance: Option<Uint128>,
    service_fee_percentage: Option<u64>,
) -> Result<Response, ContractError> {
    let mut config = CONFIG.load(deps.storage)?;
    
    // Only admin can update config
    if info.sender != config.admin {
        return Err(ContractError::Unauthorized {});
    }
    
    if let Some(denom) = stablecoin_denom {
        config.stablecoin_denom = denom;
    }
    
    if let Some(min_balance) = min_agent_balance {
        config.min_agent_balance = min_balance;
    }
    
    if let Some(fee_percentage) = service_fee_percentage {
        if fee_percentage > 10000 {
            return Err(ContractError::InvalidConfig {
                param: "service_fee_percentage cannot exceed 10000 (100%)".to_string(),
            });
        }
        config.service_fee_percentage = fee_percentage;
    }
    
    CONFIG.save(deps.storage, &config)?;
    
    Ok(Response::new()
        .add_attribute("method", "update_config")
        .add_attribute("updated_by", info.sender))
}

// Helper function to record transactions
fn record_transaction(
    storage: &mut dyn Storage,
    config: &mut Config,
    tx_type: TransactionType,
    from: Option<String>,
    to: Option<String>,
    amount: Coin,
    memo: Option<String>,
    timestamp: u64,
) -> Result<(), ContractError> {
    let tx_id = config.next_tx_id;
    config.next_tx_id += 1;
    
    let transaction = TransactionRecord {
        tx_id,
        tx_type,
        from: from.clone(),
        to: to.clone(),
        amount,
        memo,
        timestamp,
    };
    
    TRANSACTIONS.save(storage, tx_id, &transaction)?;
    
    // Index by agent
    if let Some(agent_id) = from {
        AGENT_TRANSACTIONS.save(storage, (&agent_id, tx_id), &tx_id)?;
    }
    if let Some(agent_id) = to {
        AGENT_TRANSACTIONS.save(storage, (&agent_id, tx_id), &tx_id)?;
    }
    
    Ok(())
}

// Agent Communication Functions

pub fn send_message(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    from_agent_id: String,
    to_agent_id: String,
    message_type: MessageType,
    content: String,
    requires_response: bool,
) -> Result<Response, ContractError> {
    // Verify sender owns the from_agent
    let from_agent = agents().load(deps.storage, from_agent_id.clone())?;
    if from_agent.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can send messages".to_string(),
        });
    }
    
    // Verify from_agent is active
    if !from_agent.is_active {
        return Err(ContractError::AgentNotActive {
            agent_id: from_agent_id.clone(),
        });
    }
    
    // Verify to_agent exists and is active
    let to_agent = agents().load(deps.storage, to_agent_id.clone())?;
    if !to_agent.is_active {
        return Err(ContractError::AgentNotActive {
            agent_id: to_agent_id.clone(),
        });
    }
    
    // Generate message ID
    let mut counter = CONFIG_MESSAGE_COUNTER.may_load(deps.storage)?.unwrap_or(0);
    counter += 1;
    CONFIG_MESSAGE_COUNTER.save(deps.storage, &counter)?;
    let message_id = format!("msg_{}", counter);
    
    // Create message
    let message = AgentMessage {
        message_id: message_id.clone(),
        from_agent_id: from_agent_id.clone(),
        to_agent_id: to_agent_id.clone(),
        message_type,
        content,
        requires_response,
        timestamp: env.block.time.seconds(),
        response: None,
    };
    
    // Save message
    MESSAGES.save(deps.storage, message_id.clone(), &message)?;
    
    // Update agent message boxes
    // Add to sender's outbox
    let mut outbox = AGENT_MESSAGES
        .may_load(deps.storage, (&from_agent_id, "outbox"))?
        .unwrap_or_default();
    outbox.push(message_id.clone());
    AGENT_MESSAGES.save(deps.storage, (&from_agent_id, "outbox"), &outbox)?;
    
    // Add to recipient's inbox
    let mut inbox = AGENT_MESSAGES
        .may_load(deps.storage, (&to_agent_id, "inbox"))?
        .unwrap_or_default();
    inbox.push(message_id.clone());
    AGENT_MESSAGES.save(deps.storage, (&to_agent_id, "inbox"), &inbox)?;
    
    Ok(Response::new()
        .add_attribute("method", "send_message")
        .add_attribute("message_id", message_id)
        .add_attribute("from_agent", from_agent_id)
        .add_attribute("to_agent", to_agent_id)
        .add_attribute("requires_response", requires_response.to_string()))
}

pub fn respond_to_message(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    message_id: String,
    from_agent_id: String,
    response_content: String,
) -> Result<Response, ContractError> {
    // Load the original message
    let mut message = MESSAGES.load(deps.storage, message_id.clone())?;
    
    // Verify the responder is the recipient of the original message
    if message.to_agent_id != from_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only the message recipient can respond".to_string(),
        });
    }
    
    // Verify sender owns the from_agent
    let from_agent = agents().load(deps.storage, from_agent_id.clone())?;
    if from_agent.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can respond to messages".to_string(),
        });
    }
    
    // Verify message requires response
    if !message.requires_response {
        return Err(ContractError::InvalidOperation {
            operation: "Message does not require a response".to_string(),
        });
    }
    
    // Verify message hasn't already been responded to
    if message.response.is_some() {
        return Err(ContractError::InvalidOperation {
            operation: "Message has already been responded to".to_string(),
        });
    }
    
    // Add response to message
    message.response = Some(MessageResponse {
        response_content,
        responded_at: env.block.time.seconds(),
    });
    
    // Save updated message
    MESSAGES.save(deps.storage, message_id.clone(), &message)?;
    
    Ok(Response::new()
        .add_attribute("method", "respond_to_message")
        .add_attribute("message_id", message_id)
        .add_attribute("responder_agent", from_agent_id)
        .add_attribute("responded_at", env.block.time.seconds().to_string()))
}

// Business Operations Functions

pub fn create_purchase_order(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    buyer_agent_id: String,
    seller_agent_id: String,
    items: Vec<PurchaseOrderItem>,
    delivery_terms: String,
    payment_terms: PaymentTerms,
    metadata: Option<String>,
) -> Result<Response, ContractError> {
    // Verify buyer agent exists and sender owns it
    let buyer = agents().load(deps.storage, buyer_agent_id.clone())?;
    if buyer.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can create purchase orders".to_string(),
        });
    }
    
    // Verify buyer is active
    if !buyer.is_active {
        return Err(ContractError::AgentNotActive {
            agent_id: buyer_agent_id.clone(),
        });
    }
    
    // Verify seller agent exists and is active
    let seller = agents().load(deps.storage, seller_agent_id.clone())?;
    if !seller.is_active {
        return Err(ContractError::AgentNotActive {
            agent_id: seller_agent_id.clone(),
        });
    }
    
    // Calculate total amount
    let total_amount = items.iter()
        .map(|item| item.unit_price.checked_mul(Uint128::from(item.quantity)).unwrap())
        .fold(Uint128::zero(), |acc, amount| acc + amount);
    
    // Generate PO ID
    let mut counter = CONFIG_PO_COUNTER.may_load(deps.storage)?.unwrap_or(0);
    counter += 1;
    CONFIG_PO_COUNTER.save(deps.storage, &counter)?;
    let po_id = format!("PO-{}-{}", counter, env.block.time.seconds());
    
    // Create purchase order
    let purchase_order = PurchaseOrder {
        po_id: po_id.clone(),
        buyer_agent_id: buyer_agent_id.clone(),
        seller_agent_id: seller_agent_id.clone(),
        items,
        total_amount,
        status: PurchaseOrderStatus::Draft,
        created_at: env.block.time.seconds(),
        updated_at: env.block.time.seconds(),
        delivery_terms,
        payment_terms,
        invoice_id: None,
        metadata,
    };
    
    // Save purchase order
    PURCHASE_ORDERS.save(deps.storage, po_id.clone(), &purchase_order)?;
    
    // Update agent PO lists
    let mut buyer_pos = AGENT_PURCHASE_ORDERS
        .may_load(deps.storage, (&buyer_agent_id, "buyer"))?
        .unwrap_or_default();
    buyer_pos.push(po_id.clone());
    AGENT_PURCHASE_ORDERS.save(deps.storage, (&buyer_agent_id, "buyer"), &buyer_pos)?;
    
    let mut seller_pos = AGENT_PURCHASE_ORDERS
        .may_load(deps.storage, (&seller_agent_id, "seller"))?
        .unwrap_or_default();
    seller_pos.push(po_id.clone());
    AGENT_PURCHASE_ORDERS.save(deps.storage, (&seller_agent_id, "seller"), &seller_pos)?;
    
    // Send notification message to seller
    let message_content = serde_json::json!({
        "po_id": po_id,
        "buyer": buyer_agent_id,
        "total_amount": total_amount.to_string(),
        "items_count": purchase_order.items.len(),
        "delivery_terms": purchase_order.delivery_terms,
    });
    
    let msg_counter = CONFIG_MESSAGE_COUNTER.may_load(deps.storage)?.unwrap_or(0) + 1;
    CONFIG_MESSAGE_COUNTER.save(deps.storage, &msg_counter)?;
    let message_id = format!("msg_{}", msg_counter);
    
    let notification = AgentMessage {
        message_id: message_id.clone(),
        from_agent_id: buyer_agent_id.clone(),
        to_agent_id: seller_agent_id.clone(),
        message_type: MessageType::PurchaseOrder,
        content: message_content.to_string(),
        requires_response: true,
        timestamp: env.block.time.seconds(),
        response: None,
    };
    
    MESSAGES.save(deps.storage, message_id.clone(), &notification)?;
    
    // Update message boxes
    let mut seller_inbox = AGENT_MESSAGES
        .may_load(deps.storage, (&seller_agent_id, "inbox"))?
        .unwrap_or_default();
    seller_inbox.push(message_id);
    AGENT_MESSAGES.save(deps.storage, (&seller_agent_id, "inbox"), &seller_inbox)?;
    
    Ok(Response::new()
        .add_attribute("method", "create_purchase_order")
        .add_attribute("po_id", po_id)
        .add_attribute("buyer", buyer_agent_id)
        .add_attribute("seller", seller_agent_id)
        .add_attribute("total_amount", total_amount))
}

pub fn update_purchase_order(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    po_id: String,
    status: PurchaseOrderStatus,
    updater_agent_id: String,
    notes: Option<String>,
) -> Result<Response, ContractError> {
    // Load PO
    let mut po = PURCHASE_ORDERS.load(deps.storage, po_id.clone())?;
    
    // Verify updater is either buyer or seller
    let updater = agents().load(deps.storage, updater_agent_id.clone())?;
    if updater.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can update purchase orders".to_string(),
        });
    }
    
    if updater_agent_id != po.buyer_agent_id && updater_agent_id != po.seller_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only buyer or seller can update this PO".to_string(),
        });
    }
    
    // Validate status transitions
    match (&po.status, &status) {
        (PurchaseOrderStatus::Draft, PurchaseOrderStatus::Submitted) => {
            if updater_agent_id != po.buyer_agent_id {
                return Err(ContractError::Unauthorized {
                    reason: "Only buyer can submit PO".to_string(),
                });
            }
        },
        (PurchaseOrderStatus::Submitted, PurchaseOrderStatus::Accepted) |
        (PurchaseOrderStatus::Submitted, PurchaseOrderStatus::Rejected) => {
            if updater_agent_id != po.seller_agent_id {
                return Err(ContractError::Unauthorized {
                    reason: "Only seller can accept/reject PO".to_string(),
                });
            }
        },
        (PurchaseOrderStatus::Accepted, PurchaseOrderStatus::InProgress) => {
            if updater_agent_id != po.seller_agent_id {
                return Err(ContractError::Unauthorized {
                    reason: "Only seller can mark PO as in progress".to_string(),
                });
            }
        },
        _ => {
            return Err(ContractError::InvalidOperation {
                operation: format!("Invalid status transition from {:?} to {:?}", po.status, status),
            });
        }
    }
    
    // Update PO
    po.status = status.clone();
    po.updated_at = env.block.time.seconds();
    PURCHASE_ORDERS.save(deps.storage, po_id.clone(), &po)?;
    
    // Send notification
    let recipient = if updater_agent_id == po.buyer_agent_id {
        po.seller_agent_id.clone()
    } else {
        po.buyer_agent_id.clone()
    };
    
    let message_content = serde_json::json!({
        "po_id": po_id,
        "new_status": format!("{:?}", status),
        "updated_by": updater_agent_id,
        "notes": notes,
    });
    
    let msg_counter = CONFIG_MESSAGE_COUNTER.may_load(deps.storage)?.unwrap_or(0) + 1;
    CONFIG_MESSAGE_COUNTER.save(deps.storage, &msg_counter)?;
    let message_id = format!("msg_{}", msg_counter);
    
    let notification = AgentMessage {
        message_id: message_id.clone(),
        from_agent_id: updater_agent_id.clone(),
        to_agent_id: recipient.clone(),
        message_type: MessageType::PurchaseOrder,
        content: message_content.to_string(),
        requires_response: false,
        timestamp: env.block.time.seconds(),
        response: None,
    };
    
    MESSAGES.save(deps.storage, message_id.clone(), &notification)?;
    
    // Update financials if accepted
    if matches!(status, PurchaseOrderStatus::Accepted) {
        update_agent_financials(deps.storage, &po.buyer_agent_id, None, Some(po.total_amount), env.block.time.seconds())?;
        update_agent_financials(deps.storage, &po.seller_agent_id, Some(po.total_amount), None, env.block.time.seconds())?;
    }
    
    Ok(Response::new()
        .add_attribute("method", "update_purchase_order")
        .add_attribute("po_id", po_id)
        .add_attribute("new_status", format!("{:?}", status))
        .add_attribute("updated_by", updater_agent_id))
}

pub fn create_invoice(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    po_id: String,
    seller_agent_id: String,
    line_items: Vec<InvoiceLineItem>,
    tax_rate: Option<u64>,
    discount_rate: Option<u64>,
    due_date: u64,
    metadata: Option<String>,
) -> Result<Response, ContractError> {
    // Load PO
    let mut po = PURCHASE_ORDERS.load(deps.storage, po_id.clone())?;
    
    // Verify seller
    let seller = agents().load(deps.storage, seller_agent_id.clone())?;
    if seller.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can create invoices".to_string(),
        });
    }
    
    if seller_agent_id != po.seller_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only the seller of this PO can create invoice".to_string(),
        });
    }
    
    // Verify PO status
    if !matches!(po.status, PurchaseOrderStatus::Delivered | PurchaseOrderStatus::InProgress) {
        return Err(ContractError::InvalidOperation {
            operation: "Can only invoice delivered or in-progress orders".to_string(),
        });
    }
    
    // Calculate amounts
    let subtotal = line_items.iter()
        .map(|item| item.unit_price.checked_mul(Uint128::from(item.quantity)).unwrap())
        .fold(Uint128::zero(), |acc, amount| acc + amount);
    
    let tax_amount = if let Some(rate) = tax_rate {
        subtotal.checked_mul(Uint128::from(rate)).unwrap() / Uint128::from(10000u64)
    } else {
        Uint128::zero()
    };
    
    let discount_amount = if let Some(rate) = discount_rate {
        subtotal.checked_mul(Uint128::from(rate)).unwrap() / Uint128::from(10000u64)
    } else {
        Uint128::zero()
    };
    
    let total_amount = subtotal + tax_amount - discount_amount;
    
    // Generate invoice ID
    let mut counter = CONFIG_INVOICE_COUNTER.may_load(deps.storage)?.unwrap_or(0);
    counter += 1;
    CONFIG_INVOICE_COUNTER.save(deps.storage, &counter)?;
    let invoice_id = format!("INV-{}-{}", counter, env.block.time.seconds());
    
    // Create invoice
    let invoice = Invoice {
        invoice_id: invoice_id.clone(),
        po_id: po_id.clone(),
        seller_agent_id: seller_agent_id.clone(),
        buyer_agent_id: po.buyer_agent_id.clone(),
        line_items,
        subtotal,
        tax_rate,
        tax_amount,
        discount_rate,
        discount_amount,
        total_amount,
        paid: false,
        paid_at: None,
        created_at: env.block.time.seconds(),
        due_date,
        payment_reference: None,
        metadata,
    };
    
    // Save invoice
    INVOICES.save(deps.storage, invoice_id.clone(), &invoice)?;
    
    // Update PO with invoice reference
    po.invoice_id = Some(invoice_id.clone());
    PURCHASE_ORDERS.save(deps.storage, po_id.clone(), &po)?;
    
    // Update agent invoice lists
    let mut seller_invoices = AGENT_INVOICES
        .may_load(deps.storage, (&seller_agent_id, "seller"))?
        .unwrap_or_default();
    seller_invoices.push(invoice_id.clone());
    AGENT_INVOICES.save(deps.storage, (&seller_agent_id, "seller"), &seller_invoices)?;
    
    let mut buyer_invoices = AGENT_INVOICES
        .may_load(deps.storage, (&po.buyer_agent_id, "buyer"))?
        .unwrap_or_default();
    buyer_invoices.push(invoice_id.clone());
    AGENT_INVOICES.save(deps.storage, (&po.buyer_agent_id, "buyer"), &buyer_invoices)?;
    
    // Send invoice notification
    let message_content = serde_json::json!({
        "invoice_id": invoice_id,
        "po_id": po_id,
        "seller": seller_agent_id,
        "total_amount": total_amount.to_string(),
        "due_date": due_date,
    });
    
    send_business_notification(
        deps.storage,
        &env,
        seller_agent_id.clone(),
        po.buyer_agent_id.clone(),
        MessageType::Invoice,
        message_content,
        true,
    )?;
    
    Ok(Response::new()
        .add_attribute("method", "create_invoice")
        .add_attribute("invoice_id", invoice_id)
        .add_attribute("po_id", po_id)
        .add_attribute("total_amount", total_amount))
}

pub fn pay_invoice(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    invoice_id: String,
    buyer_agent_id: String,
    payment_reference: Option<String>,
) -> Result<Response, ContractError> {
    // Load invoice
    let mut invoice = INVOICES.load(deps.storage, invoice_id.clone())?;
    
    // Verify buyer
    let buyer = agents().load(deps.storage, buyer_agent_id.clone())?;
    if buyer.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can pay invoices".to_string(),
        });
    }
    
    if buyer_agent_id != invoice.buyer_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only the buyer can pay this invoice".to_string(),
        });
    }
    
    // Verify invoice not already paid
    if invoice.paid {
        return Err(ContractError::InvalidOperation {
            operation: "Invoice already paid".to_string(),
        });
    }
    
    // Process payment (transfer from buyer to seller)
    let mut buyer_wallet = AGENT_WALLETS.load(deps.storage, buyer_agent_id.clone())?;
    let mut seller_wallet = AGENT_WALLETS.load(deps.storage, invoice.seller_agent_id.clone())?;
    
    // Check buyer balance
    if buyer_wallet.balance.amount < invoice.total_amount {
        return Err(ContractError::InsufficientBalance {
            required: invoice.total_amount.to_string(),
            available: buyer_wallet.balance.amount.to_string(),
        });
    }
    
    // Transfer funds
    buyer_wallet.balance.amount -= invoice.total_amount;
    seller_wallet.balance.amount += invoice.total_amount;
    
    AGENT_WALLETS.save(deps.storage, buyer_agent_id.clone(), &buyer_wallet)?;
    AGENT_WALLETS.save(deps.storage, invoice.seller_agent_id.clone(), &seller_wallet)?;
    
    // Update invoice
    invoice.paid = true;
    invoice.paid_at = Some(env.block.time.seconds());
    invoice.payment_reference = payment_reference;
    INVOICES.save(deps.storage, invoice_id.clone(), &invoice)?;
    
    // Update PO status
    let mut po = PURCHASE_ORDERS.load(deps.storage, invoice.po_id.clone())?;
    po.status = PurchaseOrderStatus::Completed;
    po.updated_at = env.block.time.seconds();
    PURCHASE_ORDERS.save(deps.storage, invoice.po_id.clone(), &po)?;
    
    // Record transaction
    let mut config = CONFIG.load(deps.storage)?;
    record_transaction(
        deps.storage,
        &mut config,
        TransactionType::ServicePayment,
        Some(buyer_agent_id.clone()),
        Some(invoice.seller_agent_id.clone()),
        Coin {
            denom: config.stablecoin_denom.clone(),
            amount: invoice.total_amount,
        },
        Some(format!("Payment for invoice {}", invoice_id)),
        env.block.time.seconds(),
    )?;
    CONFIG.save(deps.storage, &config)?;
    
    // Update financials
    update_agent_financials(deps.storage, &buyer_agent_id, None, None, env.block.time.seconds())?;
    update_agent_financials(deps.storage, &invoice.seller_agent_id, None, None, env.block.time.seconds())?;
    
    // Send payment notification
    let message_content = serde_json::json!({
        "invoice_id": invoice_id,
        "po_id": invoice.po_id,
        "payer": buyer_agent_id,
        "amount": invoice.total_amount.to_string(),
        "payment_reference": invoice.payment_reference,
    });
    
    send_business_notification(
        deps.storage,
        &env,
        buyer_agent_id.clone(),
        invoice.seller_agent_id.clone(),
        MessageType::PaymentNotification,
        message_content,
        false,
    )?;
    
    Ok(Response::new()
        .add_attribute("method", "pay_invoice")
        .add_attribute("invoice_id", invoice_id)
        .add_attribute("amount", invoice.total_amount))
}

pub fn confirm_receipt(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    po_id: String,
    buyer_agent_id: String,
    items_received: Vec<ItemReceipt>,
    notes: Option<String>,
) -> Result<Response, ContractError> {
    // Load PO
    let mut po = PURCHASE_ORDERS.load(deps.storage, po_id.clone())?;
    
    // Verify buyer
    let buyer = agents().load(deps.storage, buyer_agent_id.clone())?;
    if buyer.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can confirm receipt".to_string(),
        });
    }
    
    if buyer_agent_id != po.buyer_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only the buyer can confirm receipt".to_string(),
        });
    }
    
    // Create receipt record
    let receipt = Receipt {
        po_id: po_id.clone(),
        confirmed_by: buyer_agent_id.clone(),
        items_received,
        confirmed_at: env.block.time.seconds(),
        notes,
    };
    
    // Save receipt
    let mut receipts = RECEIPTS.may_load(deps.storage, po_id.clone())?.unwrap_or_default();
    receipts.push(receipt);
    RECEIPTS.save(deps.storage, po_id.clone(), &receipts)?;
    
    // Update PO status
    po.status = PurchaseOrderStatus::Delivered;
    po.updated_at = env.block.time.seconds();
    PURCHASE_ORDERS.save(deps.storage, po_id.clone(), &po)?;
    
    // Send confirmation notification
    let message_content = serde_json::json!({
        "po_id": po_id,
        "confirmed_by": buyer_agent_id,
        "confirmation_time": env.block.time.seconds(),
    });
    
    send_business_notification(
        deps.storage,
        &env,
        buyer_agent_id.clone(),
        po.seller_agent_id.clone(),
        MessageType::ReceiptConfirmation,
        message_content,
        false,
    )?;
    
    Ok(Response::new()
        .add_attribute("method", "confirm_receipt")
        .add_attribute("po_id", po_id)
        .add_attribute("confirmed_by", buyer_agent_id))
}

// Helper functions

fn update_agent_financials(
    storage: &mut dyn Storage,
    agent_id: &str,
    add_sales: Option<Uint128>,
    add_purchases: Option<Uint128>,
    timestamp: u64,
) -> Result<(), ContractError> {
    let mut financials = AGENT_FINANCIALS
        .may_load(storage, agent_id.to_string())?
        .unwrap_or(AgentFinancials {
            agent_id: agent_id.to_string(),
            total_sales: Uint128::zero(),
            total_purchases: Uint128::zero(),
            outstanding_receivables: Uint128::zero(),
            outstanding_payables: Uint128::zero(),
            completed_orders: 0,
            pending_orders: 0,
            last_updated: timestamp,
        });
    
    if let Some(sales) = add_sales {
        financials.total_sales += sales;
        financials.outstanding_receivables += sales;
        financials.pending_orders += 1;
    }
    
    if let Some(purchases) = add_purchases {
        financials.total_purchases += purchases;
        financials.outstanding_payables += purchases;
        financials.pending_orders += 1;
    }
    
    financials.last_updated = timestamp;
    AGENT_FINANCIALS.save(storage, agent_id.to_string(), &financials)?;
    
    Ok(())
}

fn send_business_notification(
    storage: &mut dyn Storage,
    env: &Env,
    from_agent_id: String,
    to_agent_id: String,
    message_type: MessageType,
    content: serde_json::Value,
    requires_response: bool,
) -> Result<(), ContractError> {
    let msg_counter = CONFIG_MESSAGE_COUNTER.may_load(storage)?.unwrap_or(0) + 1;
    CONFIG_MESSAGE_COUNTER.save(storage, &msg_counter)?;
    let message_id = format!("msg_{}", msg_counter);
    
    let notification = AgentMessage {
        message_id: message_id.clone(),
        from_agent_id: from_agent_id.clone(),
        to_agent_id: to_agent_id.clone(),
        message_type,
        content: content.to_string(),
        requires_response,
        timestamp: env.block.time.seconds(),
        response: None,
    };
    
    MESSAGES.save(storage, message_id.clone(), &notification)?;
    
    // Update inboxes
    let mut inbox = AGENT_MESSAGES
        .may_load(storage, (&to_agent_id, "inbox"))?
        .unwrap_or_default();
    inbox.push(message_id.clone());
    AGENT_MESSAGES.save(storage, (&to_agent_id, "inbox"), &inbox)?;
    
    let mut outbox = AGENT_MESSAGES
        .may_load(storage, (&from_agent_id, "outbox"))?
        .unwrap_or_default();
    outbox.push(message_id);
    AGENT_MESSAGES.save(storage, (&from_agent_id, "outbox"), &outbox)?;
    
    Ok(())
}

pub fn initiate_refund(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    invoice_id: String,
    requester_agent_id: String,
    amount: Coin,
    reason: String,
) -> Result<Response, ContractError> {
    // Load invoice
    let invoice = INVOICES.load(deps.storage, invoice_id.clone())?;
    
    // Verify requester is either buyer or seller
    let requester = agents().load(deps.storage, requester_agent_id.clone())?;
    if requester.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can request refunds".to_string(),
        });
    }
    
    if requester_agent_id != invoice.buyer_agent_id && requester_agent_id != invoice.seller_agent_id {
        return Err(ContractError::Unauthorized {
            reason: "Only buyer or seller can request refund".to_string(),
        });
    }
    
    // Verify invoice is paid
    if !invoice.paid {
        return Err(ContractError::InvalidOperation {
            operation: "Cannot refund unpaid invoice".to_string(),
        });
    }
    
    // For now, we'll implement a simple refund that requires seller approval
    // In production, this would involve more complex dispute resolution
    
    // Send refund request notification
    let recipient = if requester_agent_id == invoice.buyer_agent_id {
        invoice.seller_agent_id.clone()
    } else {
        invoice.buyer_agent_id.clone()
    };
    
    let message_content = serde_json::json!({
        "refund_request": {
            "invoice_id": invoice_id,
            "requested_by": requester_agent_id,
            "amount": amount.amount.to_string(),
            "reason": reason,
            "original_payment": invoice.total_amount.to_string(),
        }
    });
    
    send_business_notification(
        deps.storage,
        &env,
        requester_agent_id.clone(),
        recipient,
        MessageType::Custom("RefundRequest".to_string()),
        message_content,
        true,
    )?;
    
    Ok(Response::new()
        .add_attribute("method", "initiate_refund")
        .add_attribute("invoice_id", invoice_id)
        .add_attribute("requested_by", requester_agent_id)
        .add_attribute("amount", amount.amount))
}

pub fn reconcile_accounts(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    agent_id: String,
    period_start: u64,
    period_end: u64,
) -> Result<Response, ContractError> {
    // Verify agent
    let agent = agents().load(deps.storage, agent_id.clone())?;
    if agent.owner != info.sender {
        return Err(ContractError::Unauthorized {
            reason: "Only agent owner can reconcile accounts".to_string(),
        });
    }
    
    // Load agent financials
    let mut financials = AGENT_FINANCIALS
        .may_load(deps.storage, agent_id.clone())?
        .unwrap_or(AgentFinancials {
            agent_id: agent_id.clone(),
            total_sales: Uint128::zero(),
            total_purchases: Uint128::zero(),
            outstanding_receivables: Uint128::zero(),
            outstanding_payables: Uint128::zero(),
            completed_orders: 0,
            pending_orders: 0,
            last_updated: env.block.time.seconds(),
        });
    
    // Calculate outstanding amounts
    let buyer_invoices = AGENT_INVOICES
        .may_load(deps.storage, (&agent_id, "buyer"))?
        .unwrap_or_default();
    
    let seller_invoices = AGENT_INVOICES
        .may_load(deps.storage, (&agent_id, "seller"))?
        .unwrap_or_default();
    
    // Reset counters
    financials.outstanding_receivables = Uint128::zero();
    financials.outstanding_payables = Uint128::zero();
    financials.completed_orders = 0;
    financials.pending_orders = 0;
    
    // Calculate receivables (unpaid invoices where agent is seller)
    for invoice_id in seller_invoices {
        if let Ok(invoice) = INVOICES.load(deps.storage, invoice_id) {
            if invoice.created_at >= period_start && invoice.created_at <= period_end {
                if !invoice.paid {
                    financials.outstanding_receivables += invoice.total_amount;
                }
            }
        }
    }
    
    // Calculate payables (unpaid invoices where agent is buyer)
    for invoice_id in buyer_invoices {
        if let Ok(invoice) = INVOICES.load(deps.storage, invoice_id) {
            if invoice.created_at >= period_start && invoice.created_at <= period_end {
                if !invoice.paid {
                    financials.outstanding_payables += invoice.total_amount;
                }
            }
        }
    }
    
    // Count orders
    let buyer_pos = AGENT_PURCHASE_ORDERS
        .may_load(deps.storage, (&agent_id, "buyer"))?
        .unwrap_or_default();
    
    let seller_pos = AGENT_PURCHASE_ORDERS
        .may_load(deps.storage, (&agent_id, "seller"))?
        .unwrap_or_default();
    
    for po_id in buyer_pos.iter().chain(seller_pos.iter()) {
        if let Ok(po) = PURCHASE_ORDERS.load(deps.storage, po_id.clone()) {
            if po.created_at >= period_start && po.created_at <= period_end {
                match po.status {
                    PurchaseOrderStatus::Completed => financials.completed_orders += 1,
                    PurchaseOrderStatus::Draft | PurchaseOrderStatus::Cancelled => {},
                    _ => financials.pending_orders += 1,
                }
            }
        }
    }
    
    financials.last_updated = env.block.time.seconds();
    AGENT_FINANCIALS.save(deps.storage, agent_id.clone(), &financials)?;
    
    Ok(Response::new()
        .add_attribute("method", "reconcile_accounts")
        .add_attribute("agent_id", agent_id)
        .add_attribute("outstanding_receivables", financials.outstanding_receivables)
        .add_attribute("outstanding_payables", financials.outstanding_payables)
        .add_attribute("completed_orders", financials.completed_orders.to_string())
        .add_attribute("pending_orders", financials.pending_orders.to_string()))
}