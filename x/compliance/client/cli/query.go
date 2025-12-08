package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/compliance/types"
)

// NewQueryCmd returns the root query command for the compliance module.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliance query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetProfileCmd(),
	)

	return cmd
}

// NewGetProfileCmd fetches a compliance profile by address.
func NewGetProfileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "profile [address]",
		Short: "Query a compliance profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			key := append(append([]byte{}, types.ProfileKeyPrefix...), []byte(addr.String())...)

			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("profile for %s not found", addr.String())
			}

			var profile types.Profile
			types.ModuleCdc.MustUnmarshalJSON(res, &profile)

			return clientCtx.PrintObjectLegacy(profile)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
