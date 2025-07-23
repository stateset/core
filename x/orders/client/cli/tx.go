package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/orders/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateOrder())
	cmd.AddCommand(CmdUpdateOrder())
	cmd.AddCommand(CmdCancelOrder())
	cmd.AddCommand(CmdFulfillOrder())
	cmd.AddCommand(CmdRefundOrder())
	cmd.AddCommand(CmdUpdateOrderStatus())

	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdCreateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-order [customer] [merchant] [items] [total-amount] [currency]",
		Short: "Create a new order",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCustomer := args[0]
			argMerchant := args[1]
			argItems := args[2]
			argTotalAmount := args[3]
			argCurrency := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse total amount
			totalAmount, err := sdk.ParseCoinsNormalized(argTotalAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateOrder(
				clientCtx.GetFromAddress().String(),
				argCustomer,
				argMerchant,
				nil, // items would need to be parsed from JSON
				nil, // shipping info
				nil, // payment info
				argCurrency,
				totalAmount,
				"", // fulfillment type
				"", // source
				nil, // discounts
				nil, // tax info
				"", // metadata
				nil, // due date
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdCancelOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-order [order-id] [reason]",
		Short: "Cancel an existing order",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderID := args[0]
			argReason := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelOrder(
				clientCtx.GetFromAddress().String(),
				argOrderID,
				argReason,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-order [order-id]",
		Short: "Update an existing order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateOrder(
				clientCtx.GetFromAddress().String(),
				argOrderID,
				nil, // items
				nil, // shipping info
				nil, // payment info
				"", // metadata
				nil, // due date
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdFulfillOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fulfill-order [order-id] [tracking-number] [carrier]",
		Short: "Fulfill an existing order",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderID := args[0]
			argTrackingNumber := args[1]
			argCarrier := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgFulfillOrder(
				clientCtx.GetFromAddress().String(),
				argOrderID,
				argTrackingNumber,
				argCarrier,
				nil, // shipped at
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRefundOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refund-order [order-id] [refund-amount] [reason]",
		Short: "Refund an existing order",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderID := args[0]
			argRefundAmount := args[1]
			argReason := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse refund amount
			refundAmount, err := sdk.ParseCoinsNormalized(argRefundAmount)
			if err != nil {
				return err
			}

			msg := types.NewMsgRefundOrder(
				clientCtx.GetFromAddress().String(),
				argOrderID,
				refundAmount,
				argReason,
				false, // partial refund
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateOrderStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-order-status [order-id] [status] [notes]",
		Short: "Update order status",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argOrderID := args[0]
			argStatus := args[1]
			argNotes := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateOrderStatus(
				clientCtx.GetFromAddress().String(),
				argOrderID,
				argStatus,
				argNotes,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}