package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/oracle/types"
)

// NewQueryCmd returns the root query command for oracle data.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Oracle query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewGetPriceCmd())
	return cmd
}

// NewGetPriceCmd retrieves the latest price for a denom.
func NewGetPriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price [denom]",
		Short: "Query the oracle price for a denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			key := types.PriceStoreKey(args[0])
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("price for denom %s not found", args[0])
			}

			var price types.Price
			types.ModuleCdc.MustUnmarshalJSON(res, &price)

			return clientCtx.PrintObjectLegacy(price)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
