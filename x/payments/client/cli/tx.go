package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/payments/types"
)

const (
	flagMetadata = "metadata"
	flagReason   = "reason"
)

// NewTxCmd builds the root tx command for payments.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Payments transaction subcommands",
		Aliases:                    []string{"payments"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCreatePaymentCmd(),
		NewSettlePaymentCmd(),
		NewCancelPaymentCmd(),
	)

	return cmd
}

// NewCreatePaymentCmd escrows funds for a payment intent.
func NewCreatePaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [payee] [amount]",
		Short: "Create a new payment intent",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			payee := args[0]
			if _, err := sdk.AccAddressFromBech32(payee); err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePayment(clientCtx.GetFromAddress().String(), payee, amount, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMetadata, "", "Optional metadata for the payment")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewSettlePaymentCmd releases escrowed funds to the payee.
func NewSettlePaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle [payment-id]",
		Short: "Settle a payment intent and receive funds",
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

			msg := types.NewMsgSettlePayment(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewCancelPaymentCmd cancels a pending payment and refunds the payer.
func NewCancelPaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [payment-id]",
		Short: "Cancel a pending payment intent",
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

			reason, err := cmd.Flags().GetString(flagReason)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelPayment(clientCtx.GetFromAddress().String(), id, reason)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReason, "", "Optional cancellation reason")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
