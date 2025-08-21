use cosmwasm_std::{Addr, Deps, Env, MessageInfo, StdResult, Timestamp, Uint128};
use crate::error::ContractError;

/// Rate limiting structure for preventing DoS attacks
pub struct RateLimit {
    pub max_operations: u32,
    pub time_window: u64, // in seconds
}

/// Track operation counts for rate limiting
pub struct OperationTracker {
    pub address: Addr,
    pub count: u32,
    pub window_start: Timestamp,
}

impl RateLimit {
    pub fn check_and_update(
        &self,
        tracker: &mut OperationTracker,
        env: &Env,
    ) -> Result<(), ContractError> {
        let current_time = env.block.time;
        let window_duration = self.time_window * 1_000_000_000; // Convert to nanos
        
        // Reset counter if outside window
        if current_time.nanos() > tracker.window_start.nanos() + window_duration {
            tracker.count = 0;
            tracker.window_start = current_time;
        }
        
        // Check rate limit
        if tracker.count >= self.max_operations {
            return Err(ContractError::RateLimitExceeded {
                max_operations: self.max_operations,
                time_window: self.time_window,
            });
        }
        
        tracker.count += 1;
        Ok(())
    }
}

/// Security checks for escrow operations
pub struct SecurityChecks;

impl SecurityChecks {
    /// Validate that sender has required permissions
    pub fn validate_permission(
        sender: &Addr,
        required: &Addr,
        action: &str,
    ) -> Result<(), ContractError> {
        if sender != required {
            return Err(ContractError::Unauthorized {});
        }
        Ok(())
    }
    
    /// Check for reentrancy protection
    pub fn check_no_reentrancy(
        is_locked: bool,
    ) -> Result<(), ContractError> {
        if is_locked {
            return Err(ContractError::ReentrancyDetected {});
        }
        Ok(())
    }
    
    /// Validate amount is within reasonable bounds
    pub fn validate_amount_bounds(
        amount: Uint128,
        min: Uint128,
        max: Uint128,
    ) -> Result<(), ContractError> {
        if amount < min {
            return Err(ContractError::AmountTooLow {
                min: min.to_string(),
                got: amount.to_string(),
            });
        }
        if amount > max {
            return Err(ContractError::AmountTooHigh {
                max: max.to_string(),
                got: amount.to_string(),
            });
        }
        Ok(())
    }
    
    /// Validate address is not blacklisted
    pub fn check_not_blacklisted(
        address: &Addr,
        blacklist: &[Addr],
    ) -> Result<(), ContractError> {
        if blacklist.contains(address) {
            return Err(ContractError::AddressBlacklisted {
                address: address.to_string(),
            });
        }
        Ok(())
    }
    
    /// Validate contract is not paused
    pub fn check_not_paused(is_paused: bool) -> Result<(), ContractError> {
        if is_paused {
            return Err(ContractError::ContractPaused {});
        }
        Ok(())
    }
}

/// Emergency pause functionality
pub struct EmergencyPause {
    pub is_paused: bool,
    pub paused_at: Option<Timestamp>,
    pub paused_by: Option<Addr>,
}

impl EmergencyPause {
    pub fn new() -> Self {
        Self {
            is_paused: false,
            paused_at: None,
            paused_by: None,
        }
    }
    
    pub fn pause(&mut self, env: &Env, pauser: Addr) -> Result<(), ContractError> {
        if self.is_paused {
            return Err(ContractError::AlreadyPaused {});
        }
        
        self.is_paused = true;
        self.paused_at = Some(env.block.time);
        self.paused_by = Some(pauser);
        
        Ok(())
    }
    
    pub fn unpause(&mut self) -> Result<(), ContractError> {
        if !self.is_paused {
            return Err(ContractError::NotPaused {});
        }
        
        self.is_paused = false;
        self.paused_at = None;
        self.paused_by = None;
        
        Ok(())
    }
}