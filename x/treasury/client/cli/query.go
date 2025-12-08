package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/treasury/types"
)

// NewQueryCmd returns the root query command for treasury.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Treasury query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetSnapshotCmd(),
	)

	return cmd
}

// NewGetSnapshotCmd fetches a reserve snapshot by id.
func NewGetSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot [id]",
		Short: "Query a reserve snapshot by id",
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

			key := types.SnapshotStoreKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("snapshot %d not found", id)
			}

			var snapshot types.ReserveSnapshot
			types.ModuleCdc.MustUnmarshalJSON(res, &snapshot)

			return clientCtx.PrintObjectLegacy(snapshot)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
