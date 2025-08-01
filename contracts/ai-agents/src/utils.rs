use cosmwasm_std::Addr;
use crate::error::ContractError;

pub fn generate_agent_id(owner: &Addr, counter: u64) -> String {
    format!("agent_{:x}_{}", 
        &owner.as_bytes()[..8.min(owner.as_bytes().len())].iter()
            .fold(0u64, |acc, &b| acc.wrapping_mul(256).wrapping_add(b as u64)),
        counter
    )
}

pub fn generate_service_id(requester_id: &str, provider_id: &str, counter: u64) -> String {
    format!("service_{}_{}_{}",
        &requester_id[6..10.min(requester_id.len())],
        &provider_id[6..10.min(provider_id.len())],
        counter
    )
}

pub fn validate_agent_name(name: &str) -> Result<(), ContractError> {
    if name.is_empty() {
        return Err(ContractError::InvalidAgentName {
            reason: "Name cannot be empty".to_string(),
        });
    }
    
    if name.len() > 64 {
        return Err(ContractError::InvalidAgentName {
            reason: "Name cannot exceed 64 characters".to_string(),
        });
    }
    
    if !name.chars().all(|c| c.is_alphanumeric() || c == '_' || c == '-' || c == ' ') {
        return Err(ContractError::InvalidAgentName {
            reason: "Name can only contain alphanumeric characters, underscores, hyphens, and spaces".to_string(),
        });
    }
    
    Ok(())
}