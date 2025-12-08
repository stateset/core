package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/circuit/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdPauseSystem(),
		CmdResumeSystem(),
		CmdTripCircuit(),
		CmdResetCircuit(),
	)

	return cmd
}

// CmdPauseSystem creates a command to pause the entire system
func CmdPauseSystem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-system [reason] --duration [seconds]",
		Short: "Pause the entire system (requires authority)",
		Long: `Pause all system operations. This is an emergency measure.

Example:
  statesetd tx circuit pause-system "Security incident detected" --duration 3600 --from authority`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			reason := args[0]
			durationStr, _ := cmd.Flags().GetString("duration")
			duration, _ := strconv.ParseInt(durationStr, 10, 64)

			msg := types.NewMsgPauseSystem(
				clientCtx.GetFromAddress().String(),
				reason,
				duration,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("duration", "0", "Duration in seconds (0 = indefinite)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdResumeSystem creates a command to resume the system
func CmdResumeSystem() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume-system",
		Short: "Resume the system from pause (requires authority)",
		Long: `Resume system operations after a pause.

Example:
  statesetd tx circuit resume-system --from authority`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgResumeSystem(clientCtx.GetFromAddress().String())

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdTripCircuit creates a command to trip a module's circuit breaker
func CmdTripCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "trip-circuit [module-name] [reason]",
		Short: "Trip a module's circuit breaker (requires authority)",
		Long: `Trip the circuit breaker for a specific module, disabling its operations.

Example:
  statesetd tx circuit trip-circuit stablecoin "Investigating anomaly" --from authority
  statesetd tx circuit trip-circuit stablecoin "Emergency" --disable-messages "/stateset.stablecoin.v1.MsgMint,/stateset.stablecoin.v1.MsgLiquidate" --from authority`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moduleName := args[0]
			reason := args[1]

			disableMsgsStr, _ := cmd.Flags().GetString("disable-messages")
			var disableMsgs []string
			if disableMsgsStr != "" {
				disableMsgs = strings.Split(disableMsgsStr, ",")
			}

			msg := types.NewMsgTripCircuit(
				clientCtx.GetFromAddress().String(),
				moduleName,
				reason,
				disableMsgs,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("disable-messages", "", "Comma-separated list of message types to disable (empty = all)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdResetCircuit creates a command to reset a module's circuit breaker
func CmdResetCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset-circuit [module-name]",
		Short: "Reset a module's circuit breaker (requires authority)",
		Long: `Reset the circuit breaker for a specific module, re-enabling its operations.

Example:
  statesetd tx circuit reset-circuit stablecoin --from authority`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			moduleName := args[0]

			msg := types.NewMsgResetCircuit(
				clientCtx.GetFromAddress().String(),
				moduleName,
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
