package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/stablecoin/types"
)

const (
	flagDebt = "debt"
)

// NewTxCmd returns the root tx command for stablecoin operations.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Stablecoin transaction subcommands",
		Aliases:                    []string{"stablecoin", "stablecoins"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreateVaultCmd(),
		NewDepositCollateralCmd(),
		NewWithdrawCollateralCmd(),
		NewMintStablecoinCmd(),
		NewRepayStablecoinCmd(),
		NewLiquidateVaultCmd(),
		// Reserve-backed ssUSD operations
		NewDepositReserveCmd(),
		NewRequestRedemptionCmd(),
		NewExecuteRedemptionCmd(),
		NewCancelRedemptionCmd(),
		NewUpdateReserveParamsCmd(),
		NewRecordAttestationCmd(),
		NewSetApprovedAttesterCmd(),
	)

	return cmd
}

func NewCreateVaultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vault [collateral]",
		Short: "Create a vault by locking collateral",
		Long:  "Create a new vault and optionally mint stablecoin debt in a single transaction.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			debtStr, err := cmd.Flags().GetString(flagDebt)
			if err != nil {
				return err
			}

			var debt sdk.Coin
			if debtStr != "" {
				debt, err = sdk.ParseCoinNormalized(debtStr)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgCreateVault(clientCtx.GetFromAddress().String(), collateral, debt)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagDebt, "", "Optional stablecoin debt to mint (e.g. 1000ssusd)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func NewDepositCollateralCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit [vault-id] [collateral]",
		Short: "Deposit additional collateral into a vault",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositCollateral(clientCtx.GetFromAddress().String(), id, collateral)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewWithdrawCollateralCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw [vault-id] [collateral]",
		Short: "Withdraw collateral from a vault",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			collateral, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgWithdrawCollateral{
				Owner:      clientCtx.GetFromAddress().String(),
				VaultId:    id,
				Collateral: collateral,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewMintStablecoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [vault-id] [amount]",
		Short: "Mint stablecoin from a vault",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgMintStablecoin{
				Owner:   clientCtx.GetFromAddress().String(),
				VaultId: id,
				Amount:  amount,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRepayStablecoinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [vault-id] [amount]",
		Short: "Repay stablecoin debt for a vault",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgRepayStablecoin{
				Owner:   clientCtx.GetFromAddress().String(),
				VaultId: id,
				Amount:  amount,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewLiquidateVaultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidate [vault-id]",
		Short: "Liquidate an undercollateralized vault",
		Long: "Liquidate an undercollateralized vault by repaying its outstanding ssusd debt in " +
			"exchange for the locked collateral. You must have sufficient ssusd to cover the debt.",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.MsgLiquidateVault{
				Liquidator: clientCtx.GetFromAddress().String(),
				VaultId:    id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// ============================================================================
// Reserve-backed ssUSD CLI
// ============================================================================

func NewDepositReserveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-reserve [amount]",
		Short: "Deposit approved tokenized treasuries to mint ssUSD",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			msg := types.MsgDepositReserve{
				Depositor: clientCtx.GetFromAddress().String(),
				Amount:    amount,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRequestRedemptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-redemption [ssusd-amount] [output-denom]",
		Short: "Request redemption of ssUSD into an approved reserve asset",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			if coin.Denom != types.StablecoinDenom {
				return fmt.Errorf("ssusd amount denom must be %s", types.StablecoinDenom)
			}

			msg := types.MsgRequestRedemption{
				Requester:   clientCtx.GetFromAddress().String(),
				SsusdAmount: coin.Amount.String(),
				OutputDenom: args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewExecuteRedemptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute-redemption [redemption-id]",
		Short: "Execute a pending redemption request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.MsgExecuteRedemption{
				Executor:     clientCtx.GetFromAddress().String(),
				RedemptionId: id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCancelRedemptionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-redemption [redemption-id]",
		Short: "Cancel a pending redemption request (authority only)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.MsgCancelRedemption{
				Authority:    clientCtx.GetFromAddress().String(),
				RedemptionId: id,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewUpdateReserveParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-reserve-params [params-json-file]",
		Short: "Update reserve parameters from a JSON file (authority only)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bz, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var params types.ReserveParams
			if err := json.Unmarshal(bz, &params); err != nil {
				return err
			}

			msg := types.MsgUpdateReserveParams{
				Authority: clientCtx.GetFromAddress().String(),
				Params:    params,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRecordAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record-attestation [attestation-json-file]",
		Short: "Record an off-chain reserve attestation from a JSON file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			bz, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			var msg types.MsgRecordAttestation
			if err := json.Unmarshal(bz, &msg); err != nil {
				return err
			}
			if msg.Attester == "" {
				msg.Attester = clientCtx.GetFromAddress().String()
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewSetApprovedAttesterCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-approved-attester [attester] [true|false]",
		Short: "Add or remove an approved attester (authority only)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			approved, err := strconv.ParseBool(args[1])
			if err != nil {
				return err
			}

			msg := types.MsgSetApprovedAttester{
				Authority: clientCtx.GetFromAddress().String(),
				Attester:  args[0],
				Approved:  approved,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
