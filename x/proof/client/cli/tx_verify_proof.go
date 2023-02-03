package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/proof/types"
)

var _ = strconv.Itoa(0)

func CmdVerifyProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify-proof [proof] [public-signals]",
		Short: "Broadcast message verify-proof",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProof := args[0]
			argPublicSignals := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerifyProof(
				clientCtx.GetFromAddress().String(),
				argProof,
				argPublicSignals,
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
