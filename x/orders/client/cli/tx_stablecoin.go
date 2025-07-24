package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/orders/types"
)

// GetCmdPayWithStablecoin returns the command for paying with stablecoins
func GetCmdPayWithStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pay-with-stablecoin [order-id] [stablecoin-denom] [amount] [customer-address] [merchant-address] [exchange-rate]",
		Short: "Pay for an order using stablecoins",
		Long: `Pay for an order using stablecoins. This command supports both direct payment and escrow.

Examples:
- Direct payment: pay-with-stablecoin ORDER-123 uusdc 1000000 cosmos1abc... cosmos1def... 1.0
- Escrow payment: pay-with-stablecoin ORDER-123 uusdc 1000000 cosmos1abc... cosmos1def... 1.0 --use-escrow --confirmations-required 6
`,
		Args: cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orderId := args[0]
			stablecoinDenom := args[1]
			
			amount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid amount: %s", args[2])
			}

			customerAddress := args[3]
			merchantAddress := args[4]
			
			exchangeRate, err := sdk.NewDecFromStr(args[5])
			if err != nil {
				return fmt.Errorf("invalid exchange rate: %s", args[5])
			}

			useEscrow, _ := cmd.Flags().GetBool("use-escrow")
			confirmationsRequired, _ := cmd.Flags().GetUint64("confirmations-required")
			escrowTimeoutHours, _ := cmd.Flags().GetUint64("escrow-timeout-hours")

			var escrowTimeout *time.Time
			if useEscrow && escrowTimeoutHours > 0 {
				timeout := time.Now().Add(time.Hour * time.Duration(escrowTimeoutHours))
				escrowTimeout = &timeout
			}

			msg := &types.MsgPayWithStablecoin{
				Creator:               clientCtx.GetFromAddress().String(),
				OrderId:               orderId,
				StablecoinDenom:       stablecoinDenom,
				StablecoinAmount:      amount,
				CustomerAddress:       customerAddress,
				MerchantAddress:       merchantAddress,
				ExchangeRate:          exchangeRate,
				UseEscrow:             useEscrow,
				ConfirmationsRequired: confirmationsRequired,
				EscrowTimeout:         escrowTimeout,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool("use-escrow", false, "Use escrow for payment security")
	cmd.Flags().Uint64("confirmations-required", 1, "Number of confirmations required for payment")
	cmd.Flags().Uint64("escrow-timeout-hours", 720, "Escrow timeout in hours (default 30 days)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdConfirmStablecoinPayment returns the command for confirming stablecoin payments
func GetCmdConfirmStablecoinPayment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "confirm-stablecoin-payment [order-id] [confirmation-count] [block-height]",
		Short: "Confirm a stablecoin payment after required confirmations",
		Long: `Confirm a stablecoin payment after the required number of confirmations have been reached.
This is typically called by the merchant once they have verified the payment on-chain.

Example:
confirm-stablecoin-payment ORDER-123 6 1234567
`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orderId := args[0]
			
			confirmationCount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid confirmation count: %s", args[1])
			}

			blockHeight, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid block height: %s", args[2])
			}

			msg := &types.MsgConfirmStablecoinPayment{
				Creator:           clientCtx.GetFromAddress().String(),
				OrderId:           orderId,
				ConfirmationCount: confirmationCount,
				BlockHeight:       blockHeight,
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

// GetCmdRefundStablecoinPayment returns the command for refunding stablecoin payments
func GetCmdRefundStablecoinPayment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refund-stablecoin-payment [order-id] [customer-address] [refund-amount] [reason]",
		Short: "Process a refund for a stablecoin payment",
		Long: `Process a refund for a stablecoin payment. This can be a full or partial refund.
The refund will be sent back to the customer's address.

Examples:
- Full refund: refund-stablecoin-payment ORDER-123 cosmos1abc... 1000000 "Order cancelled"
- Partial refund: refund-stablecoin-payment ORDER-123 cosmos1abc... 500000 "Partial return" --partial
`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orderId := args[0]
			customerAddress := args[1]
			
			refundAmount, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid refund amount: %s", args[2])
			}

			reason := args[3]
			partialRefund, _ := cmd.Flags().GetBool("partial")

			msg := &types.MsgRefundStablecoinPayment{
				Creator:         clientCtx.GetFromAddress().String(),
				OrderId:         orderId,
				CustomerAddress: customerAddress,
				RefundAmount:    refundAmount,
				Reason:          reason,
				PartialRefund:   partialRefund,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool("partial", false, "Mark this as a partial refund")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdReleaseEscrow returns the command for releasing escrow
func GetCmdReleaseEscrow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release-escrow [order-id]",
		Short: "Release escrowed stablecoins to the merchant",
		Long: `Release escrowed stablecoins to the merchant. This can be called by either the customer 
or the merchant once the order has been satisfactorily completed.

Example:
release-escrow ORDER-123
`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			orderId := args[0]

			msg := &types.MsgReleaseEscrow{
				Creator: clientCtx.GetFromAddress().String(),
				OrderId: orderId,
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

// GetStablecoinTxCmd returns the root command for stablecoin payment transactions
func GetStablecoinTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "stablecoin",
		Short:                      "Stablecoin payment transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdPayWithStablecoin(),
		GetCmdConfirmStablecoinPayment(),
		GetCmdRefundStablecoinPayment(),
		GetCmdReleaseEscrow(),
	)

	return cmd
}