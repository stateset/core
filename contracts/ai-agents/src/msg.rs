use cosmwasm_schema::{cw_serde, QueryResponses};
use cosmwasm_std::{Addr, Coin, Uint128};

#[cw_serde]
pub struct InstantiateMsg {
    pub stablecoin_denom: String,
    pub min_agent_balance: Uint128,
    pub service_fee_percentage: u64, // basis points (10000 = 100%)
}

#[cw_serde]
pub enum ExecuteMsg {
    // Agent Management
    RegisterAgent {
        name: String,
        description: String,
        capabilities: Vec<String>,
        service_endpoints: Vec<String>,
        initial_balance: Option<Coin>,
    },
    UpdateAgent {
        agent_id: String,
        name: Option<String>,
        description: Option<String>,
        capabilities: Option<Vec<String>>,
        service_endpoints: Option<Vec<String>>,
    },
    DeactivateAgent {
        agent_id: String,
    },
    
    // Wallet Management
    FundAgent {
        agent_id: String,
    },
    WithdrawFromAgent {
        agent_id: String,
        amount: Coin,
        recipient: String,
    },
    
    // Agent-to-Agent Transactions
    AgentTransfer {
        from_agent_id: String,
        to_agent_id: String,
        amount: Coin,
        memo: Option<String>,
    },
    
    // Service Interactions
    RequestService {
        requester_agent_id: String,
        provider_agent_id: String,
        service_type: String,
        payment: Coin,
        parameters: String, // JSON encoded parameters
    },
    CompleteService {
        service_id: String,
        result: String, // JSON encoded result
    },
    RefundService {
        service_id: String,
        reason: String,
    },
    
    // Batch Operations
    BatchAgentTransfer {
        from_agent_id: String,
        transfers: Vec<Transfer>,
    },
    
    // Agent Communication
    SendMessage {
        from_agent_id: String,
        to_agent_id: String,
        message_type: MessageType,
        content: String, // JSON encoded content
        requires_response: bool,
    },
    RespondToMessage {
        message_id: String,
        from_agent_id: String,
        response_content: String, // JSON encoded response
    },
    
    // Business Operations
    CreatePurchaseOrder {
        buyer_agent_id: String,
        seller_agent_id: String,
        items: Vec<PurchaseOrderItem>,
        delivery_terms: String,
        payment_terms: PaymentTerms,
        metadata: Option<String>,
    },
    UpdatePurchaseOrder {
        po_id: String,
        status: PurchaseOrderStatus,
        updater_agent_id: String,
        notes: Option<String>,
    },
    CreateInvoice {
        po_id: String,
        seller_agent_id: String,
        line_items: Vec<InvoiceLineItem>,
        tax_rate: Option<u64>, // basis points
        discount_rate: Option<u64>, // basis points
        due_date: u64,
        metadata: Option<String>,
    },
    PayInvoice {
        invoice_id: String,
        buyer_agent_id: String,
        payment_reference: Option<String>,
    },
    ConfirmReceipt {
        po_id: String,
        buyer_agent_id: String,
        items_received: Vec<ItemReceipt>,
        notes: Option<String>,
    },
    InitiateRefund {
        invoice_id: String,
        requester_agent_id: String,
        amount: Coin,
        reason: String,
    },
    ReconcileAccounts {
        agent_id: String,
        period_start: u64,
        period_end: u64,
    },
    
    // Governance
    UpdateConfig {
        stablecoin_denom: Option<String>,
        min_agent_balance: Option<Uint128>,
        service_fee_percentage: Option<u64>,
    },
}

#[cw_serde]
pub struct Transfer {
    pub to_agent_id: String,
    pub amount: Coin,
    pub memo: Option<String>,
}

#[cw_serde]
pub enum MessageType {
    ServiceRequest,
    ServiceResponse,
    Negotiation,
    Information,
    Alert,
    PurchaseOrder,
    Invoice,
    PaymentNotification,
    ReceiptConfirmation,
    Custom(String),
}

#[cw_serde]
pub struct PurchaseOrderItem {
    pub item_id: String,
    pub description: String,
    pub quantity: u64,
    pub unit_price: Uint128,
    pub unit: String, // e.g., "piece", "kg", "hour"
}

#[cw_serde]
pub struct PaymentTerms {
    pub payment_type: PaymentType,
    pub deposit_percentage: Option<u64>, // basis points
    pub net_days: u64, // payment due in X days
}

#[cw_serde]
pub enum PaymentType {
    Immediate,
    Net, // Pay within net_days
    Deposit, // Deposit + remainder on delivery
    Milestone, // Pay per milestone
}

#[cw_serde]
pub enum PurchaseOrderStatus {
    Draft,
    Submitted,
    Accepted,
    Rejected,
    InProgress,
    Delivered,
    Completed,
    Cancelled,
}

#[cw_serde]
pub struct InvoiceLineItem {
    pub description: String,
    pub quantity: u64,
    pub unit_price: Uint128,
    pub po_item_id: Option<String>, // Reference to PO item
}

#[cw_serde]
pub struct ItemReceipt {
    pub po_item_id: String,
    pub quantity_received: u64,
    pub condition: ItemCondition,
    pub notes: Option<String>,
}

#[cw_serde]
pub enum ItemCondition {
    Good,
    Damaged,
    Missing,
    Wrong,
}

#[cw_serde]
pub enum AgentRole {
    Buyer,
    Seller,
    Both,
}

#[cw_serde]
#[derive(QueryResponses)]
pub enum QueryMsg {
    #[returns(ConfigResponse)]
    Config {},
    
    #[returns(AgentResponse)]
    Agent { agent_id: String },
    
    #[returns(AgentsResponse)]
    ListAgents {
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(AgentBalanceResponse)]
    AgentBalance { agent_id: String },
    
    #[returns(ServiceResponse)]
    Service { service_id: String },
    
    #[returns(ServicesResponse)]
    ListServices {
        agent_id: Option<String>,
        status: Option<ServiceStatus>,
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(AgentCapabilitiesResponse)]
    AgentsByCapability {
        capability: String,
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(TransactionHistoryResponse)]
    TransactionHistory {
        agent_id: String,
        start_after: Option<u64>,
        limit: Option<u32>,
    },
    
    #[returns(MessagesResponse)]
    AgentMessages {
        agent_id: String,
        message_type: Option<MessageType>,
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(MessageResponse)]
    Message { message_id: String },
    
    #[returns(PurchaseOrderResponse)]
    PurchaseOrder { po_id: String },
    
    #[returns(PurchaseOrdersResponse)]
    AgentPurchaseOrders {
        agent_id: String,
        role: AgentRole, // Buyer or Seller
        status: Option<PurchaseOrderStatus>,
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(InvoiceResponse)]
    Invoice { invoice_id: String },
    
    #[returns(InvoicesResponse)]
    AgentInvoices {
        agent_id: String,
        role: AgentRole,
        paid: Option<bool>,
        start_after: Option<String>,
        limit: Option<u32>,
    },
    
    #[returns(AccountSummaryResponse)]
    AccountSummary {
        agent_id: String,
        period_start: Option<u64>,
        period_end: Option<u64>,
    },
}

// Response Types
#[cw_serde]
pub struct ConfigResponse {
    pub stablecoin_denom: String,
    pub min_agent_balance: Uint128,
    pub service_fee_percentage: u64,
    pub total_agents: u64,
    pub total_services: u64,
}

#[cw_serde]
pub struct AgentResponse {
    pub agent_id: String,
    pub owner: Addr,
    pub name: String,
    pub description: String,
    pub capabilities: Vec<String>,
    pub service_endpoints: Vec<String>,
    pub wallet_address: Addr,
    pub balance: Coin,
    pub is_active: bool,
    pub created_at: u64,
    pub last_active: u64,
    pub total_services_provided: u64,
    pub total_services_requested: u64,
    pub reputation_score: u64,
}

#[cw_serde]
pub struct AgentsResponse {
    pub agents: Vec<AgentInfo>,
}

#[cw_serde]
pub struct AgentInfo {
    pub agent_id: String,
    pub name: String,
    pub is_active: bool,
    pub balance: Coin,
    pub reputation_score: u64,
}

#[cw_serde]
pub struct AgentBalanceResponse {
    pub agent_id: String,
    pub balance: Coin,
    pub locked_balance: Coin,
    pub available_balance: Coin,
}

#[cw_serde]
pub struct ServiceResponse {
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
}

#[cw_serde]
pub struct ServicesResponse {
    pub services: Vec<ServiceInfo>,
}

#[cw_serde]
pub struct ServiceInfo {
    pub service_id: String,
    pub service_type: String,
    pub status: ServiceStatus,
    pub payment: Coin,
    pub created_at: u64,
}

#[cw_serde]
pub struct AgentCapabilitiesResponse {
    pub capability: String,
    pub agents: Vec<AgentInfo>,
}

#[cw_serde]
pub struct TransactionHistoryResponse {
    pub transactions: Vec<Transaction>,
}

#[cw_serde]
pub struct Transaction {
    pub tx_id: u64,
    pub tx_type: TransactionType,
    pub from: Option<String>,
    pub to: Option<String>,
    pub amount: Coin,
    pub memo: Option<String>,
    pub timestamp: u64,
}

#[cw_serde]
pub enum ServiceStatus {
    Pending,
    InProgress,
    Completed,
    Failed,
    Refunded,
}

#[cw_serde]
pub enum TransactionType {
    Deposit,
    Withdrawal,
    Transfer,
    ServicePayment,
    ServiceRefund,
    Fee,
}

#[cw_serde]
pub struct MessagesResponse {
    pub messages: Vec<MessageInfo>,
}

#[cw_serde]
pub struct MessageInfo {
    pub message_id: String,
    pub from_agent_id: String,
    pub to_agent_id: String,
    pub message_type: MessageType,
    pub content: String,
    pub requires_response: bool,
    pub timestamp: u64,
    pub response: Option<MessageResponseInfo>,
}

#[cw_serde]
pub struct MessageResponseInfo {
    pub response_content: String,
    pub responded_at: u64,
}

#[cw_serde]
pub struct MessageResponse {
    pub message: MessageInfo,
}

#[cw_serde]
pub struct PurchaseOrderResponse {
    pub purchase_order: PurchaseOrderInfo,
}

#[cw_serde]
pub struct PurchaseOrderInfo {
    pub po_id: String,
    pub buyer_agent_id: String,
    pub seller_agent_id: String,
    pub items: Vec<PurchaseOrderItem>,
    pub total_amount: Uint128,
    pub status: PurchaseOrderStatus,
    pub created_at: u64,
    pub updated_at: u64,
    pub delivery_terms: String,
    pub payment_terms: PaymentTerms,
    pub invoice_id: Option<String>,
}

#[cw_serde]
pub struct PurchaseOrdersResponse {
    pub purchase_orders: Vec<PurchaseOrderInfo>,
}

#[cw_serde]
pub struct InvoiceResponse {
    pub invoice: InvoiceInfo,
}

#[cw_serde]
pub struct InvoiceInfo {
    pub invoice_id: String,
    pub po_id: String,
    pub seller_agent_id: String,
    pub buyer_agent_id: String,
    pub line_items: Vec<InvoiceLineItem>,
    pub subtotal: Uint128,
    pub tax_amount: Uint128,
    pub discount_amount: Uint128,
    pub total_amount: Uint128,
    pub paid: bool,
    pub paid_at: Option<u64>,
    pub created_at: u64,
    pub due_date: u64,
    pub payment_reference: Option<String>,
}

#[cw_serde]
pub struct InvoicesResponse {
    pub invoices: Vec<InvoiceInfo>,
}

#[cw_serde]
pub struct AccountSummaryResponse {
    pub agent_id: String,
    pub period_start: u64,
    pub period_end: u64,
    pub total_sales: Uint128,
    pub total_purchases: Uint128,
    pub outstanding_receivables: Uint128,
    pub outstanding_payables: Uint128,
    pub completed_orders: u64,
    pub pending_orders: u64,
}