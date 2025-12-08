package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/payments/types"
)

// NewQueryCmd returns the root query command for payments.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Payments query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetPaymentCmd(),
	)

	return cmd
}

// NewGetPaymentCmd retrieves a payment intent by id.
func NewGetPaymentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "payment [id]",
		Short: "Query a payment intent by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			key := types.PaymentStoreKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("payment %d not found", id)
			}

			var payment types.PaymentIntent
			types.ModuleCdc.MustUnmarshalJSON(res, &payment)

			return clientCtx.PrintObjectLegacy(payment)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
