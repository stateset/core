use cosmwasm_std::{
    entry_point, to_json_binary, Binary, Deps, DepsMut, Env, MessageInfo, Response, StdResult,
    Uint128, Addr,
};
use cw2::set_contract_version;
use cw20::{Cw20Coin, MinterResponse};
use cw20_base::{
    contract::{execute as cw20_execute, instantiate as cw20_instantiate, query as cw20_query},
    msg::{ExecuteMsg, InstantiateMsg, QueryMsg},
    ContractError,
};

// Contract name and version
const CONTRACT_NAME: &str = "crates.io:ai-stablecoin";
const CONTRACT_VERSION: &str = env!("CARGO_PKG_VERSION");

#[entry_point]
pub fn instantiate(
    mut deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: InstantiateMsg,
) -> Result<Response, ContractError> {
    set_contract_version(deps.storage, CONTRACT_NAME, CONTRACT_VERSION)?;
    
    // Create a modified instantiate message with AI-specific parameters
    let ai_msg = InstantiateMsg {
        name: msg.name.clone(),
        symbol: msg.symbol.clone(),
        decimals: msg.decimals,
        initial_balances: msg.initial_balances.clone(),
        mint: Some(MinterResponse {
            minter: info.sender.to_string(),
            cap: None, // Unlimited minting for AI agents
        }),
        marketing: msg.marketing.clone(),
    };
    
    // Call the base cw20 instantiate
    let res = cw20_instantiate(deps.branch(), env, info, ai_msg)?;
    
    Ok(res.add_attribute("token_type", "ai_stablecoin"))
}

#[entry_point]
pub fn execute(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    msg: ExecuteMsg,
) -> Result<Response, ContractError> {
    match msg {
        // Add custom AI agent validation for certain operations
        ExecuteMsg::Transfer { recipient, amount } => {
            // Could add AI agent registry validation here
            cw20_execute(deps, env, info, ExecuteMsg::Transfer { recipient, amount })
        }
        ExecuteMsg::TransferFrom { owner, recipient, amount } => {
            // Could add AI agent registry validation here
            cw20_execute(deps, env, info, ExecuteMsg::TransferFrom { owner, recipient, amount })
        }
        ExecuteMsg::Mint { recipient, amount } => {
            // Only allow minting for registered AI agents
            // This would integrate with the AI agent registry
            cw20_execute(deps, env, info, ExecuteMsg::Mint { recipient, amount })
        }
        // Delegate all other messages to base implementation
        _ => cw20_execute(deps, env, info, msg),
    }
}

#[entry_point]
pub fn query(deps: Deps, env: Env, msg: QueryMsg) -> StdResult<Binary> {
    cw20_query(deps, env, msg)
}

#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
    use cosmwasm_std::{coins, from_json, Uint128};
    use cw20::{BalanceResponse, Cw20Coin, TokenInfoResponse};

    #[test]
    fn proper_initialization() {
        let mut deps = mock_dependencies();

        let msg = InstantiateMsg {
            name: "AI USD".to_string(),
            symbol: "aiUSD".to_string(),
            decimals: 6,
            initial_balances: vec![Cw20Coin {
                address: "creator".to_string(),
                amount: Uint128::new(1000000000000),
            }],
            mint: None,
            marketing: None,
        };
        let info = mock_info("creator", &[]);
        let env = mock_env();

        // Instantiate the contract
        let res = instantiate(deps.as_mut(), env, info, msg).unwrap();
        assert_eq!(res.attributes[0].value, "instantiate");

        // Query token info
        let res = query(deps.as_ref(), mock_env(), QueryMsg::TokenInfo {}).unwrap();
        let token_info: TokenInfoResponse = from_json(&res).unwrap();
        assert_eq!(token_info.name, "AI USD");
        assert_eq!(token_info.symbol, "aiUSD");
        assert_eq!(token_info.decimals, 6);
    }
}