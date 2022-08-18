package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/stateset/core/x/invoice/types"
)

var _ = strconv.Itoa(0)

func CmdCreateInvoice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-invoice [id] [did] [amount] [state] [purchaser]",
		Short: "Broadcast message create-invoice",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argDid := args[1]
			argAmount := args[2]
			argState := args[3]
			argPurchaser := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateInvoice(
				clientCtx.GetFromAddress().String(),
				argId,
				argDid,
				argAmount,
				argState,
				argPurchaser,
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
