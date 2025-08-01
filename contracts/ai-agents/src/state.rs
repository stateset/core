use cosmwasm_schema::cw_serde;
use cosmwasm_std::{Addr, Coin, Uint128};
use cw_storage_plus::{Item, Map, IndexedMap, Index, MultiIndex};

use crate::msg::{ServiceStatus, TransactionType};

// Configuration
#[cw_serde]
pub struct Config {
    pub admin: Addr,
    pub stablecoin_denom: String,
    pub min_agent_balance: Uint128,
    pub service_fee_percentage: u64,
    pub next_agent_id: u64,
    pub next_service_id: u64,
    pub next_tx_id: u64,
}

// Agent Information
#[cw_serde]
pub struct Agent {
    pub agent_id: String,
    pub owner: Addr,
    pub name: String,
    pub description: String,
    pub capabilities: Vec<String>,
    pub service_endpoints: Vec<String>,
    pub wallet_address: Addr,
    pub is_active: bool,
    pub created_at: u64,
    pub last_active: u64,
    pub total_services_provided: u64,
    pub total_services_requested: u64,
    pub reputation_score: u64,
}

// Agent Wallet
#[cw_serde]
pub struct AgentWallet {
    pub agent_id: String,
    pub balance: Coin,
    pub locked_balance: Coin,
}

// Service Record
#[cw_serde]
pub struct Service {
    pub service_id: String,
    pub requester_agent_id: String,
    pub provider_agent_id: String,
    pub service_type: String,
    pub payment: Coin,
    pub status: ServiceStatus,
    pub parameters: String,
    pub result: Option<String>,
    pub created_at: u64,
    pub completed_at: Option<u64>,
    pub escrow_released: bool,
}

// Transaction Record
#[cw_serde]
pub struct TransactionRecord {
    pub tx_id: u64,
    pub tx_type: TransactionType,
    pub from: Option<String>,
    pub to: Option<String>,
    pub amount: Coin,
    pub memo: Option<String>,
    pub timestamp: u64,
}

// Capability Index
pub struct CapabilityIndexes<'a> {
    pub capability: MultiIndex<'a, String, Agent, String>,
}

impl<'a> CapabilityIndexes<'a> {
    fn get_indexes(capabilities: &[String]) -> Vec<String> {
        capabilities.to_vec()
    }
}

impl<'a> IndexList<Agent> for CapabilityIndexes<'a> {
    fn get_indexes(&'_ self) -> Box<dyn Iterator<Item = &'_ dyn Index<Agent>> + '_> {
        let v: Vec<&dyn Index<Agent>> = vec![&self.capability];
        Box::new(v.into_iter())
    }
}

// Storage Items
pub const CONFIG: Item<Config> = Item::new("config");

// Agent storage with capability indexing
pub fn agents<'a>() -> IndexedMap<'a, String, Agent, CapabilityIndexes<'a>> {
    let indexes = CapabilityIndexes {
        capability: MultiIndex::new(
            |_pk, d: &Agent| CapabilityIndexes::get_indexes(&d.capabilities),
            "agents",
            "agents__capability",
        ),
    };
    IndexedMap::new("agents", indexes)
}

// Direct maps for quick lookups
pub const AGENT_WALLETS: Map<String, AgentWallet> = Map::new("agent_wallets");
pub const SERVICES: Map<String, Service> = Map::new("services");
pub const AGENT_SERVICES: Map<(&str, &str), Vec<String>> = Map::new("agent_services");
pub const TRANSACTIONS: Map<u64, TransactionRecord> = Map::new("transactions");
pub const AGENT_TRANSACTIONS: Map<(&str, u64), u64> = Map::new("agent_transactions");

// Helper to track agent addresses to IDs
pub const AGENT_ADDRESS_TO_ID: Map<Addr, String> = Map::new("agent_address_to_id");

// Service type registry
pub const SERVICE_TYPES: Map<String, ServiceTypeInfo> = Map::new("service_types");

#[cw_serde]
pub struct ServiceTypeInfo {
    pub name: String,
    pub description: String,
    pub min_payment: Uint128,
    pub max_duration: u64, // in seconds
    pub required_capabilities: Vec<String>,
}

// Reputation tracking
pub const REPUTATION_HISTORY: Map<(&str, u64), ReputationEvent> = Map::new("reputation_history");

#[cw_serde]
pub struct ReputationEvent {
    pub agent_id: String,
    pub event_type: ReputationEventType,
    pub score_change: i32,
    pub reason: String,
    pub timestamp: u64,
}

#[cw_serde]
pub enum ReputationEventType {
    ServiceCompleted,
    ServiceFailed,
    ServiceRefunded,
    PositiveFeedback,
    NegativeFeedback,
}

// Pending withdrawals for security
pub const PENDING_WITHDRAWALS: Map<(&str, u64), PendingWithdrawal> = Map::new("pending_withdrawals");

#[cw_serde]
pub struct PendingWithdrawal {
    pub agent_id: String,
    pub amount: Coin,
    pub recipient: Addr,
    pub requested_at: u64,
    pub execute_at: u64,
}

use cw_storage_plus::IndexList;

// Agent Message
#[cw_serde]
pub struct AgentMessage {
    pub message_id: String,
    pub from_agent_id: String,
    pub to_agent_id: String,
    pub message_type: crate::msg::MessageType,
    pub content: String,
    pub requires_response: bool,
    pub timestamp: u64,
    pub response: Option<MessageResponse>,
}

#[cw_serde]
pub struct MessageResponse {
    pub response_content: String,
    pub responded_at: u64,
}

// Message storage
pub const MESSAGES: Map<String, AgentMessage> = Map::new("messages");
pub const AGENT_MESSAGES: Map<(&str, &str), Vec<String>> = Map::new("agent_messages"); // (agent_id, "inbox"|"outbox") -> message_ids
pub const CONFIG_MESSAGE_COUNTER: Item<u64> = Item::new("message_counter");

// Business Operations Storage
pub const PURCHASE_ORDERS: Map<String, PurchaseOrder> = Map::new("purchase_orders");
pub const AGENT_PURCHASE_ORDERS: Map<(&str, &str), Vec<String>> = Map::new("agent_purchase_orders"); // (agent_id, "buyer"|"seller") -> po_ids
pub const INVOICES: Map<String, Invoice> = Map::new("invoices");
pub const AGENT_INVOICES: Map<(&str, &str), Vec<String>> = Map::new("agent_invoices"); // (agent_id, "buyer"|"seller") -> invoice_ids
pub const RECEIPTS: Map<String, Vec<Receipt>> = Map::new("receipts"); // po_id -> receipts
pub const CONFIG_PO_COUNTER: Item<u64> = Item::new("po_counter");
pub const CONFIG_INVOICE_COUNTER: Item<u64> = Item::new("invoice_counter");

// Business Structures
#[cw_serde]
pub struct PurchaseOrder {
    pub po_id: String,
    pub buyer_agent_id: String,
    pub seller_agent_id: String,
    pub items: Vec<crate::msg::PurchaseOrderItem>,
    pub total_amount: Uint128,
    pub status: crate::msg::PurchaseOrderStatus,
    pub created_at: u64,
    pub updated_at: u64,
    pub delivery_terms: String,
    pub payment_terms: crate::msg::PaymentTerms,
    pub invoice_id: Option<String>,
    pub metadata: Option<String>,
}

#[cw_serde]
pub struct Invoice {
    pub invoice_id: String,
    pub po_id: String,
    pub seller_agent_id: String,
    pub buyer_agent_id: String,
    pub line_items: Vec<crate::msg::InvoiceLineItem>,
    pub subtotal: Uint128,
    pub tax_rate: Option<u64>,
    pub tax_amount: Uint128,
    pub discount_rate: Option<u64>,
    pub discount_amount: Uint128,
    pub total_amount: Uint128,
    pub paid: bool,
    pub paid_at: Option<u64>,
    pub created_at: u64,
    pub due_date: u64,
    pub payment_reference: Option<String>,
    pub metadata: Option<String>,
}

#[cw_serde]
pub struct Receipt {
    pub po_id: String,
    pub confirmed_by: String,
    pub items_received: Vec<crate::msg::ItemReceipt>,
    pub confirmed_at: u64,
    pub notes: Option<String>,
}

// Financial tracking
pub const AGENT_FINANCIALS: Map<String, AgentFinancials> = Map::new("agent_financials");

#[cw_serde]
#[derive(Default)]
pub struct AgentFinancials {
    pub agent_id: String,
    pub total_sales: Uint128,
    pub total_purchases: Uint128,
    pub outstanding_receivables: Uint128,
    pub outstanding_payables: Uint128,
    pub completed_orders: u64,
    pub pending_orders: u64,
    pub last_updated: u64,
}