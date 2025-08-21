use cosmwasm_std::StdError;
use thiserror::Error;

#[derive(Error, Debug, PartialEq)]
pub enum ContractError {
    #[error("{0}")]
    Std(#[from] StdError),

    // unauthorized 
    #[error("Unauthorized")]
    Unauthorized {},

    // tokens whitelisted can only be used in the escrow contract
    #[error("Only accepts tokens in the cw20_whitelist")]
    NotInWhitelist {},

    // escrow contract expired
    #[error("Escrow is expired")]
    Expired {},

    // balance of escrow contract is empty
    #[error("Send some coins to create an escrow")]
    EmptyBalance {},

    // escrow id in use
    #[error("Escrow id already in use")]
    AlreadyInUse {},

    // need to set escrow recipient
    #[error("Recipient is not set")]
    RecipientNotSet {},
    
    // Invalid input validation
    #[error("Invalid input for field '{field}': {msg}")]
    InvalidInput { field: String, msg: String },
    
    // Insufficient balance
    #[error("Insufficient balance: required {required}, available {available}")]
    InsufficientBalance { required: String, available: String },
    
    // Contract is paused
    #[error("Contract is currently paused")]
    ContractPaused {},
    
    // Rate limit exceeded
    #[error("Rate limit exceeded: max {max_operations} operations per {time_window} seconds")]
    RateLimitExceeded {
        max_operations: u32,
        time_window: u64,
    },
    
    // Reentrancy detected
    #[error("Reentrancy detected - operation already in progress")]
    ReentrancyDetected {},
    
    // Amount too low
    #[error("Amount {got} is below minimum {min}")]
    AmountTooLow { min: String, got: String },
    
    // Amount too high
    #[error("Amount {got} exceeds maximum {max}")]
    AmountTooHigh { max: String, got: String },
    
    // Address blacklisted
    #[error("Address {address} is blacklisted")]
    AddressBlacklisted { address: String },
    
    // Already paused
    #[error("Contract is already paused")]
    AlreadyPaused {},
    
    // Not paused
    #[error("Contract is not paused")]
    NotPaused {},
}