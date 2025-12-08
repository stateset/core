package cli

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/treasury/types"
)

const (
	flagReporter      = "reporter"
	flagOtherReserves = "other-reserves"
	flagMetadata      = "metadata"
	flagTimestamp     = "timestamp"
)

// NewTxCmd returns the root tx command for the treasury module.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Treasury transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewRecordReserveCmd())
	return cmd
}

// NewRecordReserveCmd builds the cobra command to record a reserve snapshot.
func NewRecordReserveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "record-reserve [total-supply] [fiat-reserves]",
		Short: "Record a new reserve snapshot",
		Long: `Record a reserve snapshot linking circulating supply with attested reserves.
Arguments expect coin strings, for example:

statesetd tx treasury record-reserve 1000000ssusd 1200000usd \
  --other-reserves "10eth,50000usdc" --metadata "monthly report" --from treasury-admin`,
		Example: "statesetd tx treasury record-reserve 1000000ssusd 1200000usd --other-reserves 1000usdc --metadata 'snapshot' --from validator",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reporter, err := cmd.Flags().GetString(flagReporter)
			if err != nil {
				return err
			}
			if reporter == "" {
				reporter = clientCtx.GetFromAddress().String()
			}

			totalSupply, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}

			fiatReserves, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			otherResStr, err := cmd.Flags().GetString(flagOtherReserves)
			if err != nil {
				return err
			}
			var otherReserves sdk.Coins
			if len(otherResStr) > 0 {
				otherReserves, err = sdk.ParseCoinsNormalized(otherResStr)
				if err != nil {
					return err
				}
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			var timestamp time.Time
			tsStr, err := cmd.Flags().GetString(flagTimestamp)
			if err != nil {
				return err
			}
			if tsStr != "" {
				timestamp, err = time.Parse(time.RFC3339, tsStr)
				if err != nil {
					return err
				}
			} else {
				timestamp = time.Now().UTC()
			}

			snapshot := types.ReserveSnapshot{
				Reporter:      reporter,
				TotalSupply:   totalSupply,
				FiatReserves:  fiatReserves,
				OtherReserves: otherReserves,
				Timestamp:     timestamp,
				Metadata:      metadata,
			}

			msg := types.NewMsgRecordReserve(clientCtx.GetFromAddress().String(), snapshot)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReporter, "", "Reporter (defaults to --from address)")
	cmd.Flags().String(flagOtherReserves, "", "Comma-separated list of other reserve coins")
	cmd.Flags().String(flagMetadata, "", "Optional metadata describing the snapshot")
	cmd.Flags().String(flagTimestamp, "", "Snapshot timestamp in RFC3339 format (defaults to current time)")

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
