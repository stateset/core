use cosmwasm_std::StdError;
use thiserror::Error;

#[derive(Error, Debug)]
pub enum ContractError {
    #[error("{0}")]
    Std(#[from] StdError),

    #[error("Unauthorized: {reason}")]
    Unauthorized { reason: String },

    #[error("Agent {agent_id} not found")]
    AgentNotFound { agent_id: String },

    #[error("Agent {agent_id} is not active")]
    AgentNotActive { agent_id: String },

    #[error("Agent {agent_id} already exists")]
    AgentAlreadyExists { agent_id: String },

    #[error("Insufficient balance. Required: {required}, Available: {available}")]
    InsufficientBalance { required: String, available: String },

    #[error("Invalid agent name: {reason}")]
    InvalidAgentName { reason: String },

    #[error("Invalid capability: {capability}")]
    InvalidCapability { capability: String },

    #[error("Service {service_id} not found")]
    ServiceNotFound { service_id: String },

    #[error("Service {service_id} already completed")]
    ServiceAlreadyCompleted { service_id: String },

    #[error("Service {service_id} cannot be modified in current status")]
    InvalidServiceStatus { service_id: String },

    #[error("Agent {agent_id} does not have required capability: {capability}")]
    MissingCapability { agent_id: String, capability: String },

    #[error("Invalid payment amount")]
    InvalidPayment {},

    #[error("Invalid configuration parameter: {param}")]
    InvalidConfig { param: String },

    #[error("Withdrawal pending, please wait {seconds} seconds")]
    WithdrawalPending { seconds: u64 },

    #[error("Service type {service_type} not registered")]
    ServiceTypeNotFound { service_type: String },

    #[error("Agent wallet not initialized")]
    WalletNotInitialized {},

    #[error("Invalid service parameters: {reason}")]
    InvalidServiceParameters { reason: String },

    #[error("Maximum agents limit reached")]
    MaxAgentsReached {},

    #[error("Circular transfer detected")]
    CircularTransfer {},

    #[error("Invalid batch operation: {reason}")]
    InvalidBatchOperation { reason: String },

    #[error("Custom error: {msg}")]
    CustomError { msg: String },

    #[error("Invalid operation: {operation}")]
    InvalidOperation { operation: String },
}