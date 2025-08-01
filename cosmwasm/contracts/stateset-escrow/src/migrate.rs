use cosmwasm_std::{entry_point, DepsMut, Env, MessageInfo, Response, StdResult};
use cw2::{get_contract_version, set_contract_version};

use crate::error::ContractError;

const CONTRACT_NAME: &str = "crates.io:stateset-escrow";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[derive(serde::Deserialize)]
pub struct MigrateMsg {}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn migrate(
    deps: DepsMut,
    _env: Env,
    _info: MessageInfo,
    _msg: MigrateMsg,
) -> Result<Response, ContractError> {
    let version = get_contract_version(deps.storage)?;
    if version.contract != CONTRACT_NAME {
        return Err(ContractError::InvalidInput {
            field: "contract".to_string(),
            msg: format!("Cannot migrate from different contract: {}", version.contract),
        });
    }

    // Set the new version
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;

    Ok(Response::new()
        .add_attribute("action", "migrate")
        .add_attribute("from_version", version.version)
        .add_attribute("to_version", CONTRACT_VERSION))
}