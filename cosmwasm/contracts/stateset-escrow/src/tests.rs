#[cfg(test)]
mod tests {
    use super::*;
    use cosmwasm_std::testing::{mock_dependencies, mock_env, mock_info};
    use cosmwasm_std::{attr, coin, coins, CosmosMsg, StdError, Uint128, from_binary};
    use cw20::{Cw20Coin, Cw20ReceiveMsg};

    use crate::contract::{execute, instantiate, query};
    use crate::msg::{CreateMsg, ExecuteMsg, InstantiateMsg, QueryMsg, ReceiveMsg};
    use crate::error::ContractError;

    // Helper function to create a basic escrow
    fn create_basic_escrow() -> CreateMsg {
        CreateMsg {
            id: "test_escrow".to_string(),
            arbiter: "arbiter".to_string(),
            recipient: Some("recipient".to_string()),
            title: "Test Escrow".to_string(),
            description: "A test escrow for integration testing".to_string(),
            end_height: Some(123456),
            end_time: None,
            cw20_whitelist: None,
        }
    }

    #[test]
    fn test_proper_instantiation() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        let msg = InstantiateMsg {};

        let res = instantiate(deps.as_mut(), mock_env(), info, msg).unwrap();
        
        // Check that proper event was emitted
        assert_eq!(res.events.len(), 1);
        assert_eq!(res.events[0].ty, "instantiate");
    }

    #[test]
    fn test_create_escrow_with_validation() {
        let mut deps = mock_dependencies();
        
        // Initialize contract
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Test valid escrow creation
        let create_msg = create_basic_escrow();
        let info = mock_info("creator", &coins(1000, "usdc"));
        let res = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();
        
        // Verify event emission
        assert_eq!(res.events.len(), 1);
        assert_eq!(res.events[0].ty, "escrow_created");
        
        // Test duplicate ID rejection
        let info = mock_info("creator2", &coins(500, "usdc"));
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg)).unwrap_err();
        assert_eq!(err, ContractError::AlreadyInUse {});
    }

    #[test]
    fn test_input_validation() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Test empty ID
        let mut create_msg = create_basic_escrow();
        create_msg.id = "".to_string();
        let info = mock_info("creator", &coins(1000, "usdc"));
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg)).unwrap_err();
        assert!(matches!(err, ContractError::InvalidInput { .. }));

        // Test overly long title
        let mut create_msg = create_basic_escrow();
        create_msg.title = "a".repeat(200);
        let info = mock_info("creator", &coins(1000, "usdc"));
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg)).unwrap_err();
        assert!(matches!(err, ContractError::InvalidInput { .. }));

        // Test missing end conditions
        let mut create_msg = create_basic_escrow();
        create_msg.end_height = None;
        create_msg.end_time = None;
        let info = mock_info("creator", &coins(1000, "usdc"));
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg)).unwrap_err();
        assert!(matches!(err, ContractError::InvalidInput { .. }));
    }

    #[test]
    fn test_authorization_controls() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Create escrow
        let create_msg = create_basic_escrow();
        let info = mock_info("creator", &coins(1000, "usdc"));
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();

        // Test unauthorized approval attempt
        let info = mock_info("unauthorized", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Approve { id: create_msg.id.clone() }).unwrap_err();
        assert_eq!(err, ContractError::Unauthorized {});

        // Test unauthorized set recipient attempt
        let info = mock_info("unauthorized", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::SetRecipient {
            id: create_msg.id.clone(),
            recipient: "new_recipient".to_string(),
        }).unwrap_err();
        assert_eq!(err, ContractError::Unauthorized {});
    }

    #[test]
    fn test_reentrancy_protection() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Create and approve escrow
        let create_msg = create_basic_escrow();
        let info = mock_info("creator", &coins(1000, "usdc"));
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();

        // Approve (should delete the escrow)
        let info = mock_info("arbiter", &[]);
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Approve { id: create_msg.id.clone() }).unwrap();

        // Try to approve again (should fail because escrow is deleted)
        let info = mock_info("arbiter", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Approve { id: create_msg.id }).unwrap_err();
        assert!(matches!(err, ContractError::Std(StdError::NotFound { .. })));
    }

    #[test]
    fn test_whitelist_validation() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Create escrow with whitelist
        let mut create_msg = create_basic_escrow();
        create_msg.cw20_whitelist = Some(vec!["token1".to_string(), "token2".to_string()]);
        let info = mock_info("creator", &coins(1000, "usdc"));
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();

        // Try to top up with non-whitelisted token
        let top_up_msg = ReceiveMsg::TopUp { id: create_msg.id.clone() };
        let receive_msg = Cw20ReceiveMsg {
            sender: "user".to_string(),
            amount: Uint128::new(500),
            msg: to_binary(&top_up_msg).unwrap(),
        };
        let info = mock_info("unauthorized_token", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Receive(receive_msg)).unwrap_err();
        assert_eq!(err, ContractError::NotInWhitelist {});
    }

    #[test]
    fn test_expiration_logic() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Create escrow with end height
        let mut create_msg = create_basic_escrow();
        create_msg.end_height = Some(100);
        let info = mock_info("creator", &coins(1000, "usdc"));
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();

        // Try to approve after expiration
        let mut env = mock_env();
        env.block.height = 200; // Past expiration
        let info = mock_info("arbiter", &[]);
        let err = execute(deps.as_mut(), env, info, ExecuteMsg::Approve { id: create_msg.id.clone() }).unwrap_err();
        assert_eq!(err, ContractError::Expired {});

        // Refund should work after expiration
        let mut env = mock_env();
        env.block.height = 200;
        let info = mock_info("anyone", &[]);
        let res = execute(deps.as_mut(), env, info, ExecuteMsg::Refund { id: create_msg.id }).unwrap();
        assert_eq!(res.events.len(), 1);
        assert_eq!(res.events[0].ty, "escrow_refunded");
    }

    #[test] 
    fn test_gas_optimization_constants() {
        // Test that our constants are reasonable for gas optimization
        assert!(MAX_ESCROW_ID_LENGTH <= 64);
        assert!(MAX_TITLE_LENGTH <= 128);
        assert!(MAX_DESCRIPTION_LENGTH <= 512);
        assert!(MAX_WHITELIST_SIZE <= 50);
    }

    #[test]
    fn test_empty_balance_rejection() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        let create_msg = create_basic_escrow();
        let info = mock_info("creator", &[]); // No funds
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg)).unwrap_err();
        assert_eq!(err, ContractError::EmptyBalance {});
    }

    #[test]
    fn test_recipient_validation() {
        let mut deps = mock_dependencies();
        let info = mock_info("creator", &[]);
        instantiate(deps.as_mut(), mock_env(), info, InstantiateMsg {}).unwrap();

        // Create escrow without recipient
        let mut create_msg = create_basic_escrow();
        create_msg.recipient = None;
        let info = mock_info("creator", &coins(1000, "usdc"));
        execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Create(create_msg.clone())).unwrap();

        // Try to approve without setting recipient
        let info = mock_info("arbiter", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::Approve { id: create_msg.id.clone() }).unwrap_err();
        assert_eq!(err, ContractError::RecipientNotSet {});

        // Set recipient with invalid address format
        let info = mock_info("arbiter", &[]);
        let err = execute(deps.as_mut(), mock_env(), info, ExecuteMsg::SetRecipient {
            id: create_msg.id.clone(),
            recipient: "invalid_address_format_!@#$%".to_string(),
        }).unwrap_err();
        assert!(matches!(err, ContractError::Std(StdError::GenericErr { .. })));
    }
}