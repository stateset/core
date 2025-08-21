#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
    use cosmwasm_std::{coins, Addr, BankMsg, CosmosMsg, Event, SubMsg, Uint128};
    use crate::contract::{execute, instantiate, query};
    use crate::error::ContractError;
    use crate::msg::{ConfigResponse, ExecuteMsg, InstantiateMsg, QueryMsg};

    const CREATOR: &str = "creator";
    const OWNER: &str = "owner";
    const RECIPIENT: &str = "recipient";

    fn setup_contract(
        expires: u64,
        collateral: Vec<cosmwasm_std::Coin>,
        counter_offer: Vec<cosmwasm_std::Coin>,
    ) -> (cosmwasm_std::OwnedDeps<cosmwasm_std::MemoryStorage, cosmwasm_std::testing::MockApi, cosmwasm_std::testing::MockQuerier>, cosmwasm_std::Env) {
        let mut deps = mock_dependencies();
        let mut env = mock_env();
        env.block.height = 100;

        let msg = InstantiateMsg {
            counter_offer,
            expires,
        };

        let info = mock_info(CREATOR, &collateral);
        let _res = instantiate(deps.as_mut(), env.clone(), info, msg).unwrap();

        (deps, env)
    }

    #[test]
    fn test_proper_initialization() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (deps, _env) = setup_contract(200, collateral.clone(), counter_offer.clone());

        // Query config
        let res = query(deps.as_ref(), mock_env(), QueryMsg::Config {}).unwrap();
        let config: ConfigResponse = cosmwasm_std::from_binary(&res).unwrap();

        assert_eq!(config.owner, Addr::unchecked(CREATOR));
        assert_eq!(config.creator, Addr::unchecked(CREATOR));
        assert_eq!(config.collateral, collateral);
        assert_eq!(config.counter_offer, counter_offer);
        assert_eq!(config.expires, 200);
    }

    #[test]
    fn test_instantiate_with_invalid_expiry() {
        let mut deps = mock_dependencies();
        let mut env = mock_env();
        env.block.height = 100;

        let msg = InstantiateMsg {
            counter_offer: coins(50, "atom"),
            expires: 100, // Same as current block height
        };

        let info = mock_info(CREATOR, &coins(1000, "usdc"));
        let err = instantiate(deps.as_mut(), env, info, msg).unwrap_err();
        
        match err {
            ContractError::OptionExpired { expired } => assert_eq!(expired, 100),
            _ => panic!("Expected OptionExpired error"),
        }
    }

    #[test]
    fn test_instantiate_with_empty_collateral() {
        let mut deps = mock_dependencies();
        let mut env = mock_env();
        env.block.height = 100;

        let msg = InstantiateMsg {
            counter_offer: coins(50, "atom"),
            expires: 200,
        };

        let info = mock_info(CREATOR, &[]); // Empty collateral
        let err = instantiate(deps.as_mut(), env, info, msg).unwrap_err();
        
        match err {
            ContractError::InvalidCollateral {} => {},
            _ => panic!("Expected InvalidCollateral error"),
        }
    }

    #[test]
    fn test_instantiate_with_empty_counter_offer() {
        let mut deps = mock_dependencies();
        let mut env = mock_env();
        env.block.height = 100;

        let msg = InstantiateMsg {
            counter_offer: vec![], // Empty counter offer
            expires: 200,
        };

        let info = mock_info(CREATOR, &coins(1000, "usdc"));
        let err = instantiate(deps.as_mut(), env, info, msg).unwrap_err();
        
        match err {
            ContractError::InvalidCounterOffer {} => {},
            _ => panic!("Expected InvalidCounterOffer error"),
        }
    }

    #[test]
    fn test_transfer_option() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral, counter_offer);

        // Transfer from creator to new owner
        let msg = ExecuteMsg::Transfer {
            recipient: OWNER.to_string(),
        };
        let info = mock_info(CREATOR, &[]);
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Check event was emitted
        assert_eq!(res.events.len(), 1);
        assert_eq!(res.events[0].ty, "option_transferred");

        // Query config to verify new owner
        let res = query(deps.as_ref(), env, QueryMsg::Config {}).unwrap();
        let config: ConfigResponse = cosmwasm_std::from_binary(&res).unwrap();
        assert_eq!(config.owner, Addr::unchecked(OWNER));
    }

    #[test]
    fn test_transfer_unauthorized() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral, counter_offer);

        // Try to transfer from non-owner
        let msg = ExecuteMsg::Transfer {
            recipient: RECIPIENT.to_string(),
        };
        let info = mock_info("random_user", &[]);
        let err = execute(deps.as_mut(), env, info, msg).unwrap_err();

        match err {
            ContractError::Unauthorized {} => {},
            _ => panic!("Expected Unauthorized error"),
        }
    }

    #[test]
    fn test_execute_option_success() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral.clone(), counter_offer.clone());

        // Execute option with correct counter offer
        let msg = ExecuteMsg::Execute {};
        let info = mock_info(CREATOR, &counter_offer);
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Verify messages sent
        assert_eq!(res.messages.len(), 2);
        
        // Check that collateral is sent to owner and counter_offer to creator
        let expected_msgs = vec![
            SubMsg::new(BankMsg::Send {
                to_address: CREATOR.to_string(),
                amount: counter_offer,
            }),
            SubMsg::new(BankMsg::Send {
                to_address: CREATOR.to_string(), // Since owner is still creator
                amount: collateral,
            }),
        ];
        
        assert_eq!(res.messages, expected_msgs);

        // Verify option was deleted
        let err = query(deps.as_ref(), env, QueryMsg::Config {}).unwrap_err();
        assert!(matches!(err, cosmwasm_std::StdError::NotFound { .. }));
    }

    #[test]
    fn test_execute_expired_option() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, mut env) = setup_contract(200, collateral, counter_offer.clone());

        // Move time forward past expiry
        env.block.height = 201;

        // Try to execute expired option
        let msg = ExecuteMsg::Execute {};
        let info = mock_info(CREATOR, &counter_offer);
        let err = execute(deps.as_mut(), env, info, msg).unwrap_err();

        match err {
            ContractError::OptionExpired { expired } => assert_eq!(expired, 200),
            _ => panic!("Expected OptionExpired error"),
        }
    }

    #[test]
    fn test_execute_wrong_counter_offer() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral, counter_offer.clone());

        // Try to execute with wrong counter offer
        let wrong_offer = coins(100, "atom"); // Wrong amount
        let msg = ExecuteMsg::Execute {};
        let info = mock_info(CREATOR, &wrong_offer);
        let err = execute(deps.as_mut(), env, info, msg).unwrap_err();

        match err {
            ContractError::CounterOfferMismatch { .. } => {},
            _ => panic!("Expected CounterOfferMismatch error"),
        }
    }

    #[test]
    fn test_burn_expired_option() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, mut env) = setup_contract(200, collateral.clone(), counter_offer);

        // Move time forward past expiry
        env.block.height = 201;

        // Burn expired option
        let msg = ExecuteMsg::Burn {};
        let info = mock_info("anyone", &[]); // Anyone can burn expired option
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Verify collateral is returned to creator
        assert_eq!(res.messages.len(), 1);
        assert_eq!(
            res.messages[0],
            SubMsg::new(BankMsg::Send {
                to_address: CREATOR.to_string(),
                amount: collateral,
            })
        );

        // Check event was emitted
        assert_eq!(res.events.len(), 1);
        assert_eq!(res.events[0].ty, "option_burned");

        // Verify option was deleted
        let err = query(deps.as_ref(), env, QueryMsg::Config {}).unwrap_err();
        assert!(matches!(err, cosmwasm_std::StdError::NotFound { .. }));
    }

    #[test]
    fn test_burn_not_expired() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral, counter_offer);

        // Try to burn before expiry
        let msg = ExecuteMsg::Burn {};
        let info = mock_info("anyone", &[]);
        let err = execute(deps.as_mut(), env, info, msg).unwrap_err();

        match err {
            ContractError::OptionNotExpired { expires } => assert_eq!(expires, 200),
            _ => panic!("Expected OptionNotExpired error"),
        }
    }

    #[test]
    fn test_burn_with_funds() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, mut env) = setup_contract(200, collateral, counter_offer);

        // Move time forward past expiry
        env.block.height = 201;

        // Try to burn with funds attached
        let msg = ExecuteMsg::Burn {};
        let info = mock_info("anyone", &coins(100, "usdc"));
        let err = execute(deps.as_mut(), env, info, msg).unwrap_err();

        match err {
            ContractError::FundsSentWithBurn {} => {},
            _ => panic!("Expected FundsSentWithBurn error"),
        }
    }

    #[test]
    fn test_option_lifecycle() {
        let collateral = coins(1000, "usdc");
        let counter_offer = coins(50, "atom");
        let (mut deps, env) = setup_contract(200, collateral.clone(), counter_offer.clone());

        // 1. Transfer option to new owner
        let msg = ExecuteMsg::Transfer {
            recipient: OWNER.to_string(),
        };
        let info = mock_info(CREATOR, &[]);
        execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // 2. New owner executes the option
        let msg = ExecuteMsg::Execute {};
        let info = mock_info(OWNER, &counter_offer);
        let res = execute(deps.as_mut(), env.clone(), info, msg).unwrap();

        // Verify both transfers happened
        assert_eq!(res.messages.len(), 2);
        
        // Counter offer goes to creator
        assert_eq!(
            res.messages[0].msg,
            CosmosMsg::Bank(BankMsg::Send {
                to_address: CREATOR.to_string(),
                amount: counter_offer,
            })
        );
        
        // Collateral goes to new owner
        assert_eq!(
            res.messages[1].msg,
            CosmosMsg::Bank(BankMsg::Send {
                to_address: OWNER.to_string(),
                amount: collateral,
            })
        );
    }
}