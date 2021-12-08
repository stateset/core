package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/stateset/core/x/purchaseorder/types"
)

var _ = strconv.Itoa(0)

func CmdRequestPurchaseorder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-purchaseorder [did] [uri] [amount] [state]",
		Short: "Broadcast message request-purchaseorder",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argUri := args[1]
			argAmount := args[2]
			argState := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestPurchaseorder(
				clientCtx.GetFromAddress().String(),
				argDid,
				argUri,
				argAmount,
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
