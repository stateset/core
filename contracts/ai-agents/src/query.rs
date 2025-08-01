use cosmwasm_std::{Deps, StdResult, Coin, Uint128, Order};
use cw_storage_plus::Bound;

use crate::msg::{
    ConfigResponse, AgentResponse, AgentsResponse, AgentInfo,
    AgentBalanceResponse, ServiceResponse, ServicesResponse, ServiceInfo,
    AgentCapabilitiesResponse, TransactionHistoryResponse, Transaction,
    ServiceStatus, MessagesResponse, MessageInfo, MessageResponseInfo,
    MessageResponse, MessageType, PurchaseOrderResponse, PurchaseOrderInfo,
    PurchaseOrdersResponse, InvoiceResponse, InvoiceInfo, InvoicesResponse,
    AccountSummaryResponse, AgentRole, PurchaseOrderStatus,
};
use crate::state::{
    CONFIG, agents, AGENT_WALLETS, SERVICES, TRANSACTIONS, AGENT_TRANSACTIONS,
    MESSAGES, AGENT_MESSAGES, PURCHASE_ORDERS, AGENT_PURCHASE_ORDERS,
    INVOICES, AGENT_INVOICES, AGENT_FINANCIALS,
};

const DEFAULT_LIMIT: u32 = 10;
const MAX_LIMIT: u32 = 100;

pub fn config(deps: Deps) -> StdResult<ConfigResponse> {
    let config = CONFIG.load(deps.storage)?;
    
    // Count total agents
    let total_agents = agents()
        .keys(deps.storage, None, None, Order::Ascending)
        .count() as u64;
    
    // Count total services
    let total_services = SERVICES
        .keys(deps.storage, None, None, Order::Ascending)
        .count() as u64;
    
    Ok(ConfigResponse {
        stablecoin_denom: config.stablecoin_denom,
        min_agent_balance: config.min_agent_balance,
        service_fee_percentage: config.service_fee_percentage,
        total_agents,
        total_services,
    })
}

pub fn agent(deps: Deps, agent_id: String) -> StdResult<AgentResponse> {
    let agent = agents().load(deps.storage, agent_id.clone())?;
    let wallet = AGENT_WALLETS.load(deps.storage, agent_id)?;
    
    Ok(AgentResponse {
        agent_id: agent.agent_id,
        owner: agent.owner,
        name: agent.name,
        description: agent.description,
        capabilities: agent.capabilities,
        service_endpoints: agent.service_endpoints,
        wallet_address: agent.wallet_address,
        balance: wallet.balance,
        is_active: agent.is_active,
        created_at: agent.created_at,
        last_active: agent.last_active,
        total_services_provided: agent.total_services_provided,
        total_services_requested: agent.total_services_requested,
        reputation_score: agent.reputation_score,
    })
}

pub fn list_agents(
    deps: Deps,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<AgentsResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    let start = start_after.map(Bound::exclusive);
    
    let agents_list: StdResult<Vec<AgentInfo>> = agents()
        .range(deps.storage, start, None, Order::Ascending)
        .take(limit)
        .map(|item| {
            let (agent_id, agent) = item?;
            let wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
            
            Ok(AgentInfo {
                agent_id,
                name: agent.name,
                is_active: agent.is_active,
                balance: wallet.balance,
                reputation_score: agent.reputation_score,
            })
        })
        .collect();
    
    Ok(AgentsResponse {
        agents: agents_list?,
    })
}

pub fn agent_balance(deps: Deps, agent_id: String) -> StdResult<AgentBalanceResponse> {
    let wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
    let config = CONFIG.load(deps.storage)?;
    
    let available_balance = Coin {
        denom: config.stablecoin_denom.clone(),
        amount: wallet.balance.amount.saturating_sub(wallet.locked_balance.amount),
    };
    
    Ok(AgentBalanceResponse {
        agent_id,
        balance: wallet.balance,
        locked_balance: wallet.locked_balance,
        available_balance,
    })
}

pub fn service(deps: Deps, service_id: String) -> StdResult<ServiceResponse> {
    let service = SERVICES.load(deps.storage, service_id)?;
    
    Ok(ServiceResponse {
        service_id: service.service_id,
        requester_agent_id: service.requester_agent_id,
        provider_agent_id: service.provider_agent_id,
        service_type: service.service_type,
        payment: service.payment,
        status: service.status,
        parameters: service.parameters,
        result: service.result,
        created_at: service.created_at,
        completed_at: service.completed_at,
    })
}

pub fn list_services(
    deps: Deps,
    agent_id: Option<String>,
    status: Option<ServiceStatus>,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<ServicesResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    let start = start_after.map(Bound::exclusive);
    
    let services: StdResult<Vec<ServiceInfo>> = SERVICES
        .range(deps.storage, start, None, Order::Ascending)
        .filter(|item| {
            if let Ok((_, service)) = item {
                // Filter by agent if specified
                if let Some(ref aid) = agent_id {
                    if service.requester_agent_id != *aid && service.provider_agent_id != *aid {
                        return false;
                    }
                }
                
                // Filter by status if specified
                if let Some(ref s) = status {
                    if service.status != *s {
                        return false;
                    }
                }
                
                true
            } else {
                false
            }
        })
        .take(limit)
        .map(|item| {
            let (service_id, service) = item?;
            
            Ok(ServiceInfo {
                service_id,
                service_type: service.service_type,
                status: service.status,
                payment: service.payment,
                created_at: service.created_at,
            })
        })
        .collect();
    
    Ok(ServicesResponse {
        services: services?,
    })
}

pub fn agents_by_capability(
    deps: Deps,
    capability: String,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<AgentCapabilitiesResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    
    let agents_list: StdResult<Vec<AgentInfo>> = agents()
        .idx.capability
        .prefix(capability.clone())
        .range(deps.storage, None, None, Order::Ascending)
        .skip(
            start_after
                .as_ref()
                .and_then(|s| agents().idx.capability.prefix(capability.clone())
                    .range(deps.storage, None, None, Order::Ascending)
                    .position(|item| {
                        item.map(|(k, _)| k == s.clone()).unwrap_or(false)
                    }))
                .map(|p| p + 1)
                .unwrap_or(0)
        )
        .take(limit)
        .map(|item| {
            let (agent_id, agent) = item?;
            let wallet = AGENT_WALLETS.load(deps.storage, agent_id.clone())?;
            
            Ok(AgentInfo {
                agent_id,
                name: agent.name,
                is_active: agent.is_active,
                balance: wallet.balance,
                reputation_score: agent.reputation_score,
            })
        })
        .collect();
    
    Ok(AgentCapabilitiesResponse {
        capability,
        agents: agents_list?,
    })
}

pub fn transaction_history(
    deps: Deps,
    agent_id: String,
    start_after: Option<u64>,
    limit: Option<u32>,
) -> StdResult<TransactionHistoryResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    
    // Verify agent exists
    agents().load(deps.storage, agent_id.clone())?;
    
    let start_bound = start_after.map(|id| Bound::exclusive((agent_id.as_str(), id)));
    
    let transactions: StdResult<Vec<Transaction>> = AGENT_TRANSACTIONS
        .prefix(&agent_id)
        .range(deps.storage, start_bound, None, Order::Descending)
        .take(limit)
        .map(|item| {
            let (_, tx_id) = item?;
            let tx_record = TRANSACTIONS.load(deps.storage, tx_id)?;
            
            Ok(Transaction {
                tx_id: tx_record.tx_id,
                tx_type: tx_record.tx_type,
                from: tx_record.from,
                to: tx_record.to,
                amount: tx_record.amount,
                memo: tx_record.memo,
                timestamp: tx_record.timestamp,
            })
        })
        .collect();
    
    Ok(TransactionHistoryResponse {
        transactions: transactions?,
    })
}

pub fn agent_messages(
    deps: Deps,
    agent_id: String,
    message_type: Option<MessageType>,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<MessagesResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    
    // Verify agent exists
    agents().load(deps.storage, agent_id.clone())?;
    
    // Get inbox messages
    let inbox_ids = AGENT_MESSAGES
        .may_load(deps.storage, (&agent_id, "inbox"))?
        .unwrap_or_default();
    
    // Get outbox messages
    let outbox_ids = AGENT_MESSAGES
        .may_load(deps.storage, (&agent_id, "outbox"))?
        .unwrap_or_default();
    
    // Combine and deduplicate message IDs
    let mut all_message_ids: Vec<String> = inbox_ids;
    all_message_ids.extend(outbox_ids);
    all_message_ids.sort();
    all_message_ids.dedup();
    
    // Filter by start_after if provided
    let start_index = if let Some(start) = start_after {
        all_message_ids.iter().position(|id| id > &start).unwrap_or(all_message_ids.len())
    } else {
        0
    };
    
    // Load messages and filter by type if specified
    let messages: StdResult<Vec<MessageInfo>> = all_message_ids
        .into_iter()
        .skip(start_index)
        .take(limit)
        .filter_map(|msg_id| {
            MESSAGES.load(deps.storage, msg_id.clone()).ok().and_then(|msg| {
                // Filter by message type if specified
                if let Some(ref msg_type) = message_type {
                    match (&msg.message_type, msg_type) {
                        (MessageType::ServiceRequest, MessageType::ServiceRequest) |
                        (MessageType::ServiceResponse, MessageType::ServiceResponse) |
                        (MessageType::Negotiation, MessageType::Negotiation) |
                        (MessageType::Information, MessageType::Information) |
                        (MessageType::Alert, MessageType::Alert) => Some(msg),
                        (MessageType::Custom(a), MessageType::Custom(b)) if a == b => Some(msg),
                        _ => None,
                    }
                } else {
                    Some(msg)
                }
            })
        })
        .map(|msg| {
            Ok(MessageInfo {
                message_id: msg.message_id,
                from_agent_id: msg.from_agent_id,
                to_agent_id: msg.to_agent_id,
                message_type: msg.message_type,
                content: msg.content,
                requires_response: msg.requires_response,
                timestamp: msg.timestamp,
                response: msg.response.map(|r| MessageResponseInfo {
                    response_content: r.response_content,
                    responded_at: r.responded_at,
                }),
            })
        })
        .collect();
    
    Ok(MessagesResponse {
        messages: messages?,
    })
}

pub fn message(deps: Deps, message_id: String) -> StdResult<MessageResponse> {
    let msg = MESSAGES.load(deps.storage, message_id)?;
    
    Ok(MessageResponse {
        message: MessageInfo {
            message_id: msg.message_id,
            from_agent_id: msg.from_agent_id,
            to_agent_id: msg.to_agent_id,
            message_type: msg.message_type,
            content: msg.content,
            requires_response: msg.requires_response,
            timestamp: msg.timestamp,
            response: msg.response.map(|r| MessageResponseInfo {
                response_content: r.response_content,
                responded_at: r.responded_at,
            }),
        },
    })
}

pub fn purchase_order(deps: Deps, po_id: String) -> StdResult<PurchaseOrderResponse> {
    let po = PURCHASE_ORDERS.load(deps.storage, po_id)?;
    
    Ok(PurchaseOrderResponse {
        purchase_order: PurchaseOrderInfo {
            po_id: po.po_id,
            buyer_agent_id: po.buyer_agent_id,
            seller_agent_id: po.seller_agent_id,
            items: po.items,
            total_amount: po.total_amount,
            status: po.status,
            created_at: po.created_at,
            updated_at: po.updated_at,
            delivery_terms: po.delivery_terms,
            payment_terms: po.payment_terms,
            invoice_id: po.invoice_id,
        },
    })
}

pub fn agent_purchase_orders(
    deps: Deps,
    agent_id: String,
    role: AgentRole,
    status: Option<PurchaseOrderStatus>,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<PurchaseOrdersResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    
    // Verify agent exists
    agents().load(deps.storage, agent_id.clone())?;
    
    // Get PO IDs based on role
    let po_ids = match role {
        AgentRole::Buyer => AGENT_PURCHASE_ORDERS
            .may_load(deps.storage, (&agent_id, "buyer"))?
            .unwrap_or_default(),
        AgentRole::Seller => AGENT_PURCHASE_ORDERS
            .may_load(deps.storage, (&agent_id, "seller"))?
            .unwrap_or_default(),
        AgentRole::Both => {
            let mut buyer_pos = AGENT_PURCHASE_ORDERS
                .may_load(deps.storage, (&agent_id, "buyer"))?
                .unwrap_or_default();
            let seller_pos = AGENT_PURCHASE_ORDERS
                .may_load(deps.storage, (&agent_id, "seller"))?
                .unwrap_or_default();
            buyer_pos.extend(seller_pos);
            buyer_pos.sort();
            buyer_pos.dedup();
            buyer_pos
        }
    };
    
    // Filter by start_after if provided
    let start_index = if let Some(start) = start_after {
        po_ids.iter().position(|id| id > &start).unwrap_or(po_ids.len())
    } else {
        0
    };
    
    // Load POs and filter by status if specified
    let purchase_orders: StdResult<Vec<PurchaseOrderInfo>> = po_ids
        .into_iter()
        .skip(start_index)
        .take(limit)
        .filter_map(|po_id| {
            PURCHASE_ORDERS.load(deps.storage, po_id).ok().and_then(|po| {
                if let Some(ref req_status) = status {
                    if po.status == *req_status {
                        Some(po)
                    } else {
                        None
                    }
                } else {
                    Some(po)
                }
            })
        })
        .map(|po| {
            Ok(PurchaseOrderInfo {
                po_id: po.po_id,
                buyer_agent_id: po.buyer_agent_id,
                seller_agent_id: po.seller_agent_id,
                items: po.items,
                total_amount: po.total_amount,
                status: po.status,
                created_at: po.created_at,
                updated_at: po.updated_at,
                delivery_terms: po.delivery_terms,
                payment_terms: po.payment_terms,
                invoice_id: po.invoice_id,
            })
        })
        .collect();
    
    Ok(PurchaseOrdersResponse {
        purchase_orders: purchase_orders?,
    })
}

pub fn invoice(deps: Deps, invoice_id: String) -> StdResult<InvoiceResponse> {
    let inv = INVOICES.load(deps.storage, invoice_id)?;
    
    Ok(InvoiceResponse {
        invoice: InvoiceInfo {
            invoice_id: inv.invoice_id,
            po_id: inv.po_id,
            seller_agent_id: inv.seller_agent_id,
            buyer_agent_id: inv.buyer_agent_id,
            line_items: inv.line_items,
            subtotal: inv.subtotal,
            tax_amount: inv.tax_amount,
            discount_amount: inv.discount_amount,
            total_amount: inv.total_amount,
            paid: inv.paid,
            paid_at: inv.paid_at,
            created_at: inv.created_at,
            due_date: inv.due_date,
            payment_reference: inv.payment_reference,
        },
    })
}

pub fn agent_invoices(
    deps: Deps,
    agent_id: String,
    role: AgentRole,
    paid: Option<bool>,
    start_after: Option<String>,
    limit: Option<u32>,
) -> StdResult<InvoicesResponse> {
    let limit = limit.unwrap_or(DEFAULT_LIMIT).min(MAX_LIMIT) as usize;
    
    // Verify agent exists
    agents().load(deps.storage, agent_id.clone())?;
    
    // Get invoice IDs based on role
    let invoice_ids = match role {
        AgentRole::Buyer => AGENT_INVOICES
            .may_load(deps.storage, (&agent_id, "buyer"))?
            .unwrap_or_default(),
        AgentRole::Seller => AGENT_INVOICES
            .may_load(deps.storage, (&agent_id, "seller"))?
            .unwrap_or_default(),
        AgentRole::Both => {
            let mut buyer_invs = AGENT_INVOICES
                .may_load(deps.storage, (&agent_id, "buyer"))?
                .unwrap_or_default();
            let seller_invs = AGENT_INVOICES
                .may_load(deps.storage, (&agent_id, "seller"))?
                .unwrap_or_default();
            buyer_invs.extend(seller_invs);
            buyer_invs.sort();
            buyer_invs.dedup();
            buyer_invs
        }
    };
    
    // Filter by start_after if provided
    let start_index = if let Some(start) = start_after {
        invoice_ids.iter().position(|id| id > &start).unwrap_or(invoice_ids.len())
    } else {
        0
    };
    
    // Load invoices and filter by paid status if specified
    let invoices: StdResult<Vec<InvoiceInfo>> = invoice_ids
        .into_iter()
        .skip(start_index)
        .take(limit)
        .filter_map(|invoice_id| {
            INVOICES.load(deps.storage, invoice_id).ok().and_then(|inv| {
                if let Some(req_paid) = paid {
                    if inv.paid == req_paid {
                        Some(inv)
                    } else {
                        None
                    }
                } else {
                    Some(inv)
                }
            })
        })
        .map(|inv| {
            Ok(InvoiceInfo {
                invoice_id: inv.invoice_id,
                po_id: inv.po_id,
                seller_agent_id: inv.seller_agent_id,
                buyer_agent_id: inv.buyer_agent_id,
                line_items: inv.line_items,
                subtotal: inv.subtotal,
                tax_amount: inv.tax_amount,
                discount_amount: inv.discount_amount,
                total_amount: inv.total_amount,
                paid: inv.paid,
                paid_at: inv.paid_at,
                created_at: inv.created_at,
                due_date: inv.due_date,
                payment_reference: inv.payment_reference,
            })
        })
        .collect();
    
    Ok(InvoicesResponse {
        invoices: invoices?,
    })
}

pub fn account_summary(
    deps: Deps,
    agent_id: String,
    period_start: Option<u64>,
    period_end: Option<u64>,
) -> StdResult<AccountSummaryResponse> {
    // Verify agent exists
    agents().load(deps.storage, agent_id.clone())?;
    
    // Load financials (may not exist yet)
    let financials = AGENT_FINANCIALS
        .may_load(deps.storage, agent_id.clone())?
        .unwrap_or_default();
    
    Ok(AccountSummaryResponse {
        agent_id,
        period_start: period_start.unwrap_or(0),
        period_end: period_end.unwrap_or(u64::MAX),
        total_sales: financials.total_sales,
        total_purchases: financials.total_purchases,
        outstanding_receivables: financials.outstanding_receivables,
        outstanding_payables: financials.outstanding_payables,
        completed_orders: financials.completed_orders,
        pending_orders: financials.pending_orders,
    })
}