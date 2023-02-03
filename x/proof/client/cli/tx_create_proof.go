package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/proof/types"
)

var _ = strconv.Itoa(0)

func CmdCreateProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-proof [id] [did] [uri] [hash] [state]",
		Short: "Broadcast message create-proof",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}
			argDid := args[1]
			argUri := args[2]
			argHash := args[3]
			argState := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProof(
				clientCtx.GetFromAddress().String(),
				argId,
				argDid,
				argUri,
				argHash,
				argState,
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
