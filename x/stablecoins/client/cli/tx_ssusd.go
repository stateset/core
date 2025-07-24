package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/stablecoins/types"
)

// GetSSUSDTxCmd returns the transaction commands for ssUSD management
func GetSSUSDTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "ssusd",
		Short:                      "ssUSD stablecoin transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdInitializeSSUSD(),
		CmdIssueSSUSD(),
		CmdRedeemSSUSD(),
		CmdUpdateSSUSDPrice(),
		CmdOptimizeSSUSDYield(),
		CmdAddSSUSDLiquidity(),
		CmdRemoveSSUSDLiquidity(),
		CmdStakeSSUSD(),
		CmdUnstakeSSUSD(),
		CmdClaimSSUSDRewards(),
		CmdRebalanceSSUSD(),
		CmdUpdateSSUSDCollateral(),
		CmdCreateSSUSDPool(),
		CmdBridgeSSUSD(),
	)

	return cmd
}

// CmdInitializeSSUSD initializes the ssUSD stablecoin
func CmdInitializeSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "initialize",
		Short: "Initialize the ssUSD stablecoin with enhanced features",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgInitializeSSUSD{
				Creator: clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdIssueSSUSD issues new ssUSD tokens backed by reserves
func CmdIssueSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue [amount] [reserve-payment]",
		Short: "Issue new ssUSD tokens backed 1:1 by conservative reserves",
		Long: `Issue new ssUSD tokens by providing backing reserves.
Reserve payment should be in the format: 100us_cash_token,500treasury_bill_token,75mmf_token,25repo_token
The total USD value of reserves must equal or exceed the ssUSD amount being issued.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid amount: %s", args[0])
			}

			reservePayment, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid reserve payment: %w", err)
			}

			msg := &types.MsgIssueSSUSD{
				Creator:        clientCtx.GetFromAddress().String(),
				Amount:         amount,
				ReservePayment: reservePayment,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdRedeemSSUSD redeems ssUSD tokens for underlying reserves
func CmdRedeemSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redeem [ssusd-amount] [preferred-asset]",
		Short: "Redeem ssUSD tokens for underlying reserve assets",
		Long: `Redeem ssUSD tokens and receive underlying reserve assets in return.
If preferred-asset is specified (us_cash_token, treasury_bill_token, mmf_token, or repo_token),
the system will try to fulfill the redemption with that asset type first.
If not specified or insufficient, redemption will be proportional across all reserve types.`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, ok := sdk.NewIntFromString(args[0])
			if !ok {
				return fmt.Errorf("invalid ssUSD amount: %s", args[0])
			}

			preferredAsset := ""
			if len(args) > 1 {
				preferredAsset = args[1]
				// Validate preferred asset
				validAssets := []string{"us_cash_token", "treasury_bill_token", "mmf_token", "repo_token"}
				valid := false
				for _, asset := range validAssets {
					if preferredAsset == asset {
						valid = true
						break
					}
				}
				if !valid {
					return fmt.Errorf("invalid preferred asset: %s. Valid options: %v", preferredAsset, validAssets)
				}
			}

			msg := &types.MsgRedeemSSUSD{
				Creator:        clientCtx.GetFromAddress().String(),
				SSUSDAmount:    amount,
				PreferredAsset: preferredAsset,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdUpdateSSUSDPrice updates the price feed for ssUSD
func CmdUpdateSSUSDPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-price [provider] [price]",
		Short: "Update ssUSD price from a specific provider",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			provider := args[0]
			price, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return fmt.Errorf("invalid price: %w", err)
			}

			msg := &types.MsgUpdateSSUSDPrice{
				Creator:  clientCtx.GetFromAddress().String(),
				Provider: provider,
				Price:    price,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdOptimizeSSUSDYield optimizes yield generation for ssUSD
func CmdOptimizeSSUSDYield() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "optimize-yield",
		Short: "Optimize yield generation for ssUSD",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgOptimizeSSUSDYield{
				Creator: clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdAddSSUSDLiquidity adds liquidity to an ssUSD pool
func CmdAddSSUSDLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-liquidity [pool-id] [amount-ssusd] [amount-other]",
		Short: "Add liquidity to an ssUSD liquidity pool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID := args[0]
			
			amountSSUSD, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid ssUSD amount: %w", err)
			}

			amountOther, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid other token amount: %w", err)
			}

			msg := &types.MsgAddSSUSDLiquidity{
				Creator:     clientCtx.GetFromAddress().String(),
				PoolID:      poolID,
				AmountSSUSD: amountSSUSD,
				AmountOther: amountOther,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdRemoveSSUSDLiquidity removes liquidity from an ssUSD pool
func CmdRemoveSSUSDLiquidity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-liquidity [pool-id] [lp-tokens]",
		Short: "Remove liquidity from an ssUSD liquidity pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			poolID := args[0]
			
			lpTokens, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid LP tokens: %w", err)
			}

			msg := &types.MsgRemoveSSUSDLiquidity{
				Creator:  clientCtx.GetFromAddress().String(),
				PoolID:   poolID,
				LPTokens: lpTokens,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdStakeSSUSD stakes ssUSD for yield generation
func CmdStakeSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake [amount] [strategy-id]",
		Short: "Stake ssUSD for yield generation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return fmt.Errorf("invalid amount: %w", err)
			}

			strategyID := args[1]

			msg := &types.MsgStakeSSUSD{
				Creator:    clientCtx.GetFromAddress().String(),
				Amount:     amount,
				StrategyID: strategyID,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdUnstakeSSUSD unstakes ssUSD from yield generation
func CmdUnstakeSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unstake [amount] [strategy-id]",
		Short: "Unstake ssUSD from yield generation",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return fmt.Errorf("invalid amount: %w", err)
			}

			strategyID := args[1]

			msg := &types.MsgUnstakeSSUSD{
				Creator:    clientCtx.GetFromAddress().String(),
				Amount:     amount,
				StrategyID: strategyID,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdClaimSSUSDRewards claims earned rewards from ssUSD staking
func CmdClaimSSUSDRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-rewards [strategy-id]",
		Short: "Claim earned rewards from ssUSD staking",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			strategyID := args[0]

			msg := &types.MsgClaimSSUSDRewards{
				Creator:    clientCtx.GetFromAddress().String(),
				StrategyID: strategyID,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdRebalanceSSUSD manually triggers ssUSD rebalancing
func CmdRebalanceSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rebalance",
		Short: "Manually trigger ssUSD rebalancing",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRebalanceSSUSD{
				Creator: clientCtx.GetFromAddress().String(),
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdUpdateSSUSDCollateral updates collateral backing ssUSD
func CmdUpdateSSUSDCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-collateral [denom] [amount] [action]",
		Short: "Update collateral backing ssUSD (add/remove)",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]
			
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid amount: %w", err)
			}

			action := args[2]
			if action != "add" && action != "remove" {
				return fmt.Errorf("action must be 'add' or 'remove'")
			}

			msg := &types.MsgUpdateSSUSDCollateral{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:   denom,
				Amount:  amount,
				Action:  action,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdCreateSSUSDPool creates a new ssUSD liquidity pool
func CmdCreateSSUSDPool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [token-pair] [initial-liquidity-ssusd] [initial-liquidity-other] [trading-fee] [apy]",
		Short: "Create a new ssUSD liquidity pool",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			tokenPair := args[0]
			
			initialLiquiditySSUSD, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return fmt.Errorf("invalid ssUSD liquidity: %w", err)
			}

			initialLiquidityOther, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid other token liquidity: %w", err)
			}

			tradingFee, err := sdk.NewDecFromStr(args[3])
			if err != nil {
				return fmt.Errorf("invalid trading fee: %w", err)
			}

			apy, err := sdk.NewDecFromStr(args[4])
			if err != nil {
				return fmt.Errorf("invalid APY: %w", err)
			}

			msg := &types.MsgCreateSSUSDPool{
				Creator:               clientCtx.GetFromAddress().String(),
				TokenPair:            tokenPair,
				InitialLiquiditySSUSD: initialLiquiditySSUSD,
				InitialLiquidityOther: initialLiquidityOther,
				TradingFee:           tradingFee,
				APY:                  apy,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdBridgeSSUSD bridges ssUSD to another chain
func CmdBridgeSSUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bridge [to-chain] [to-address] [amount]",
		Short: "Bridge ssUSD to another blockchain",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			toChain := args[0]
			toAddress := args[1]
			
			amount, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return fmt.Errorf("invalid amount: %w", err)
			}

			msg := &types.MsgBridgeSSUSD{
				Creator:   clientCtx.GetFromAddress().String(),
				ToChain:   toChain,
				ToAddress: toAddress,
				Amount:    amount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetSSUSDQueryCmd returns the query commands for ssUSD
func GetSSUSDQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "ssusd",
		Short:                      "Querying commands for ssUSD stablecoin",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQuerySSUSDPrice(),
		CmdQuerySSUSDMetrics(),
		CmdQuerySSUSDPools(),
		CmdQuerySSUSDPosition(),
		CmdQuerySSUSDRewards(),
		CmdQuerySSUSDCollateral(),
		CmdQuerySSUSDStrategies(),
		CmdQuerySSUSDRisk(),
	)

	return cmd
}

// CmdQuerySSUSDPrice queries the current ssUSD price
func CmdQuerySSUSDPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price",
		Short: "Query current ssUSD price from all feeds",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDPrice(cmd.Context(), &types.QuerySSUSDPriceRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDMetrics queries comprehensive ssUSD metrics
func CmdQuerySSUSDMetrics() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Query comprehensive ssUSD metrics",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDMetrics(cmd.Context(), &types.QuerySSUSDMetricsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDPools queries ssUSD liquidity pools
func CmdQuerySSUSDPools() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pools [pool-id]",
		Short: "Query ssUSD liquidity pools (specify pool-id for specific pool)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			
			var poolID string
			if len(args) > 0 {
				poolID = args[0]
			}

			res, err := queryClient.SSUSDPools(cmd.Context(), &types.QuerySSUSDPoolsRequest{
				PoolId: poolID,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDPosition queries user's ssUSD position
func CmdQuerySSUSDPosition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "position [user-address]",
		Short: "Query user's ssUSD staking and liquidity positions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDPosition(cmd.Context(), &types.QuerySSUSDPositionRequest{
				UserAddress: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDRewards queries pending ssUSD rewards
func CmdQuerySSUSDRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rewards [user-address] [strategy-id]",
		Short: "Query pending ssUSD rewards for a user",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDRewards(cmd.Context(), &types.QuerySSUSDRewardsRequest{
				UserAddress: args[0],
				StrategyId:  args[1],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDCollateral queries ssUSD collateral information
func CmdQuerySSUSDCollateral() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collateral [denom]",
		Short: "Query ssUSD collateral information (specify denom for specific collateral)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			
			var denom string
			if len(args) > 0 {
				denom = args[0]
			}

			res, err := queryClient.SSUSDCollateral(cmd.Context(), &types.QuerySSUSDCollateralRequest{
				Denom: denom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDStrategies queries available yield strategies
func CmdQuerySSUSDStrategies() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "strategies",
		Short: "Query available ssUSD yield strategies",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDStrategies(cmd.Context(), &types.QuerySSUSDStrategiesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySSUSDRisk queries ssUSD risk metrics
func CmdQuerySSUSDRisk() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "risk",
		Short: "Query ssUSD risk metrics and assessments",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.SSUSDRisk(cmd.Context(), &types.QuerySSUSDRiskRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}