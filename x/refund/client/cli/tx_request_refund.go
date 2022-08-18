package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/refund/types"
)

var _ = strconv.Itoa(0)

func CmdRequestRefund() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-refund [did] [amount] [fee] [deadline]",
		Short: "Broadcast message request-refund",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDid := args[0]
			argAmount := args[1]
			argFee := args[2]
			argDeadline := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestRefund(
				clientCtx.GetFromAddress().String(),
				argDid,
				argAmount,
				argFee,
				argDeadline,
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
