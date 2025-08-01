use cosmwasm_std::{
    entry_point, to_json_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult,
    Coin, Addr, Uint128, BankMsg, CosmosMsg, WasmMsg, StdError,
};
use cw2::set_contract_version;

use crate::error::ContractError;
use crate::msg::{ExecuteMsg, InstantiateMsg, QueryMsg};
use crate::state::{Config, Agent, AgentWallet, CONFIG, agents, AGENT_WALLETS, AGENT_ADDRESS_TO_ID};
use crate::{execute, query};

// Contract version info
const CONTRACT_NAME: &str = "crates.io:ai-agent-registry";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[entry_point]
pub fn instantiate(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;

    // Validate configuration
    if msg.service_fee_percentage > 10000 {
        return Err(ContractError::InvalidConfig {
            param: "service_fee_percentage cannot exceed 10000 (100%)".to_string(),
        });
    }

    let config = Config {
        admin: info.sender,
        stablecoin_denom: msg.stablecoin_denom,
        min_agent_balance: msg.min_agent_balance,
        service_fee_percentage: msg.service_fee_percentage,
        next_agent_id: 1,
        next_service_id: 1,
        next_tx_id: 1,
    };

    CONFIG.save(deps.storage, &config)?;

    Ok(Response::new()
        .add_attribute("method", "instantiate")
        .add_attribute("admin", config.admin)
        .add_attribute("stablecoin_denom", config.stablecoin_denom))
}

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response, ContractError> {
    match msg {
        // Agent Management
        ExecuteMsg::RegisterAgent {
            name,
            description,
            capabilities,
            service_endpoints,
            initial_balance,
        } => execute::register_agent(
            deps,
            env,
            info,
            name,
            description,
            capabilities,
            service_endpoints,
            initial_balance,
        ),
        ExecuteMsg::UpdateAgent {
            agent_id,
            name,
            description,
            capabilities,
            service_endpoints,
        } => execute::update_agent(
            deps,
            env,
            info,
            agent_id,
            name,
            description,
            capabilities,
            service_endpoints,
        ),
        ExecuteMsg::DeactivateAgent { agent_id } => {
            execute::deactivate_agent(deps, env, info, agent_id)
        }
        
        // Wallet Management
        ExecuteMsg::FundAgent { agent_id } => execute::fund_agent(deps, env, info, agent_id),
        ExecuteMsg::WithdrawFromAgent {
            agent_id,
            amount,
            recipient,
        } => execute::withdraw_from_agent(deps, env, info, agent_id, amount, recipient),
        
        // Agent-to-Agent Transactions
        ExecuteMsg::AgentTransfer {
            from_agent_id,
            to_agent_id,
            amount,
            memo,
        } => execute::agent_transfer(deps, env, info, from_agent_id, to_agent_id, amount, memo),
        
        // Service Interactions
        ExecuteMsg::RequestService {
            requester_agent_id,
            provider_agent_id,
            service_type,
            payment,
            parameters,
        } => execute::request_service(
            deps,
            env,
            info,
            requester_agent_id,
            provider_agent_id,
            service_type,
            payment,
            parameters,
        ),
        ExecuteMsg::CompleteService { service_id, result } => {
            execute::complete_service(deps, env, info, service_id, result)
        }
        ExecuteMsg::RefundService { service_id, reason } => {
            execute::refund_service(deps, env, info, service_id, reason)
        }
        
        // Batch Operations
        ExecuteMsg::BatchAgentTransfer {
            from_agent_id,
            transfers,
        } => execute::batch_agent_transfer(deps, env, info, from_agent_id, transfers),
        
        // Agent Communication
        ExecuteMsg::SendMessage {
            from_agent_id,
            to_agent_id,
            message_type,
            content,
            requires_response,
        } => execute::send_message(
            deps,
            env,
            info,
            from_agent_id,
            to_agent_id,
            message_type,
            content,
            requires_response,
        ),
        ExecuteMsg::RespondToMessage {
            message_id,
            from_agent_id,
            response_content,
        } => execute::respond_to_message(deps, env, info, message_id, from_agent_id, response_content),
        
        // Business Operations
        ExecuteMsg::CreatePurchaseOrder {
            buyer_agent_id,
            seller_agent_id,
            items,
            delivery_terms,
            payment_terms,
            metadata,
        } => execute::create_purchase_order(
            deps,
            env,
            info,
            buyer_agent_id,
            seller_agent_id,
            items,
            delivery_terms,
            payment_terms,
            metadata,
        ),
        ExecuteMsg::UpdatePurchaseOrder {
            po_id,
            status,
            updater_agent_id,
            notes,
        } => execute::update_purchase_order(deps, env, info, po_id, status, updater_agent_id, notes),
        ExecuteMsg::CreateInvoice {
            po_id,
            seller_agent_id,
            line_items,
            tax_rate,
            discount_rate,
            due_date,
            metadata,
        } => execute::create_invoice(
            deps,
            env,
            info,
            po_id,
            seller_agent_id,
            line_items,
            tax_rate,
            discount_rate,
            due_date,
            metadata,
        ),
        ExecuteMsg::PayInvoice {
            invoice_id,
            buyer_agent_id,
            payment_reference,
        } => execute::pay_invoice(deps, env, info, invoice_id, buyer_agent_id, payment_reference),
        ExecuteMsg::ConfirmReceipt {
            po_id,
            buyer_agent_id,
            items_received,
            notes,
        } => execute::confirm_receipt(deps, env, info, po_id, buyer_agent_id, items_received, notes),
        ExecuteMsg::InitiateRefund {
            invoice_id,
            requester_agent_id,
            amount,
            reason,
        } => execute::initiate_refund(deps, env, info, invoice_id, requester_agent_id, amount, reason),
        ExecuteMsg::ReconcileAccounts {
            agent_id,
            period_start,
            period_end,
        } => execute::reconcile_accounts(deps, env, info, agent_id, period_start, period_end),
        
        // Governance
        ExecuteMsg::UpdateConfig {
            stablecoin_denom,
            min_agent_balance,
            service_fee_percentage,
        } => execute::update_config(
            deps,
            env,
            info,
            stablecoin_denom,
            min_agent_balance,
            service_fee_percentage,
        ),
    }
}

#[entry_point]
pub fn query(deps: Deps, env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Config {} => to_json_binary(&query::config(deps)?),
        QueryMsg::Agent { agent_id } => to_json_binary(&query::agent(deps, agent_id)?),
        QueryMsg::ListAgents { start_after, limit } => {
            to_json_binary(&query::list_agents(deps, start_after, limit)?)
        }
        QueryMsg::AgentBalance { agent_id } => {
            to_json_binary(&query::agent_balance(deps, agent_id)?)
        }
        QueryMsg::Service { service_id } => to_json_binary(&query::service(deps, service_id)?),
        QueryMsg::ListServices {
            agent_id,
            status,
            start_after,
            limit,
        } => to_json_binary(&query::list_services(
            deps,
            agent_id,
            status,
            start_after,
            limit,
        )?),
        QueryMsg::AgentsByCapability {
            capability,
            start_after,
            limit,
        } => to_json_binary(&query::agents_by_capability(
            deps,
            capability,
            start_after,
            limit,
        )?),
        QueryMsg::TransactionHistory {
            agent_id,
            start_after,
            limit,
        } => to_json_binary(&query::transaction_history(
            deps,
            agent_id,
            start_after,
            limit,
        )?),
        QueryMsg::AgentMessages {
            agent_id,
            message_type,
            start_after,
            limit,
        } => to_json_binary(&query::agent_messages(
            deps,
            agent_id,
            message_type,
            start_after,
            limit,
        )?),
        QueryMsg::Message { message_id } => to_json_binary(&query::message(deps, message_id)?),
        QueryMsg::PurchaseOrder { po_id } => to_json_binary(&query::purchase_order(deps, po_id)?),
        QueryMsg::AgentPurchaseOrders {
            agent_id,
            role,
            status,
            start_after,
            limit,
        } => to_json_binary(&query::agent_purchase_orders(
            deps,
            agent_id,
            role,
            status,
            start_after,
            limit,
        )?),
        QueryMsg::Invoice { invoice_id } => to_json_binary(&query::invoice(deps, invoice_id)?),
        QueryMsg::AgentInvoices {
            agent_id,
            role,
            paid,
            start_after,
            limit,
        } => to_json_binary(&query::agent_invoices(
            deps,
            agent_id,
            role,
            paid,
            start_after,
            limit,
        )?),
        QueryMsg::AccountSummary {
            agent_id,
            period_start,
            period_end,
        } => to_json_binary(&query::account_summary(
            deps,
            agent_id,
            period_start,
            period_end,
        )?),
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
    use cosmwasm_std::{coins, from_json};

    #[test]
    fn proper_initialization() {
        let mut deps = mock_dependencies();

        let msg = InstantiateMsg {
            stablecoin_denom: "usdc".to_string(),
            min_agent_balance: Uint128::new(1000000),
            service_fee_percentage: 250, // 2.5%
        };
        let info = mock_info("creator", &coins(1000, "usdc"));

        let res = instantiate(deps.as_mut(), mock_env(), info, msg).unwrap();
        assert_eq!(0, res.messages.len());
    }
}