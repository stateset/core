package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/stablecoin/types"
)

// NewQueryCmd returns the root query command for stablecoin.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Stablecoin query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetVaultCmd(),
	)

	return cmd
}

// NewGetVaultCmd retrieves vault information.
func NewGetVaultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vault [id]",
		Short: "Query a vault by id",
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

			key := types.VaultStoreKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("vault %d not found", id)
			}

			var vault types.Vault
			types.ModuleCdc.MustUnmarshalJSON(res, &vault)

			return clientCtx.PrintObjectLegacy(vault)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
