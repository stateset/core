use cosmwasm_std::{
    entry_point, to_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult, Order,
};
use cw_storage_plus::{Item, Map};

use crate::error::ContractError;
use crate::msg::{ConfigResponse, ExecuteMsg, InstantiateMsg, QueryMsg};
use crate::state::{config, config_read, State};

// Storage for proofs
pub const PROOF_INFO: Map<String, State> = Map::new("proof_info");

// Instantiate the Smart Contract for a Stateset Proof
#[entry_point]
pub fn instantiate(
    deps: DepsMut, // Mutable Deps
    _env: Env,
    info: MessageInfo,
    msg: InstantiateMsg, // InstantiateMsg from msg.rs crate
) -> Result<Response, ContractError> {

    // State from state.rs crate
    let state = State {
        provider: info.sender.clone(), // creator of the proof is the provider
        did: msg.did, // DID for the proof
        payload: msg.payload, // payload for the proof
        status: msg.status // status for the proof
    };

    // Save the State defined in state.rs crate to the state
    config(deps.storage).save(&state)?;

    Ok(Response::new()
        .add_attribute("action", "instantiate")
        .add_attribute("provider", info.sender)
        .add_attribute("did", &state.did))
}


// Entry Point for Execution of the Proof
#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg, // ExecuteMsg as defined in the msg.rs crate
) -> Result<Response, ContractError> {
    match msg {
        // ExecuteMsg::Verify from msg.rs crate
        ExecuteMsg::Verify { proof } => execute_verify(deps, env, info, proof),
    }
}


// Verify the Proof
pub fn execute_verify(
    deps: DepsMut,
    _env: Env,
    info: MessageInfo,
    proof: String, // proof payload to verify
) -> Result<Response, ContractError> {

    // Check Contract Error Conditions
    // load the state from deps.storage imported from cosmwasm_std crate
    let mut state = config(deps.storage).load()?;
    
    // Only the provider can verify the proof
    if info.sender != state.provider {
        // ContractError::Unauthorized from error.rs crate
        return Err(ContractError::Unauthorized {});
    }

    // Simple proof verification (in a real implementation, this would be more sophisticated)
    let is_valid = proof == state.payload;
    
    // Update the status based on verification result
    state.status = if is_valid {
        "verified".to_string()
    } else {
        "invalid".to_string()
    };
    
    // save the state to storage
    config(deps.storage).save(&state)?;

    let mut res = Response::new();
    res = res.add_attribute("action", "verify");
    res = res.add_attribute("status", &state.status);
    res = res.add_attribute("proof_valid", is_valid.to_string());
    
    if !is_valid {
        return Err(ContractError::VerificationFailed {});
    }
    
    Ok(res)
}


#[cfg_attr(not(feature = "library"), entry_point)]
pub fn query(deps: Deps, _env: Env, msg: QueryMsg) -> StdResult<Binary> {
    match msg {
        QueryMsg::Config {} => to_binary(&query_config(deps)?),
    }
}

fn query_config(deps: Deps) -> StdResult<ConfigResponse> {
    let state = config_read(deps.storage).load()?;
    Ok(state)
}