package cli

import (
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/oracle/types"
)

// NewTxCmd returns the root tx command for oracle updates.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Oracle transaction subcommands",
		Aliases:                    []string{"oracle"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewUpdatePriceCmd())
	return cmd
}

// NewUpdatePriceCmd submits a price update.
func NewUpdatePriceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-price [denom] [price]",
		Short: "Update the oracle price for a denom",
		Long:  "Submit a new price for the provided denom (e.g. 1.23). Only the configured authority may broadcast this message.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			price, err := sdkmath.LegacyNewDecFromStr(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePrice(clientCtx.GetFromAddress().String(), args[0], price)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
