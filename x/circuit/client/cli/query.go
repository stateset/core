package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/circuit/types"
)

// GetQueryCmd returns the query commands for this module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryParams(),
		CmdQueryCircuitState(),
		CmdQueryModuleCircuit(),
		CmdQueryRateLimits(),
		CmdQueryLiquidationProtection(),
	)

	return cmd
}

// CmdQueryParams creates a command to query module params
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current circuit breaker parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			// Note: This would need a query client implementation
			fmt.Fprintf(clientCtx.Output, "Query params for circuit module\n")
			fmt.Fprintf(clientCtx.Output, "Note: Full query implementation requires gRPC query server\n")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCircuitState creates a command to query global circuit state
func CmdQueryCircuitState() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "state",
		Short: "Query the global circuit breaker state",
		Long: `Query whether the system is globally paused and related information.

Example:
  statesetd query circuit state`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			fmt.Fprintf(clientCtx.Output, "Query global circuit state\n")
			fmt.Fprintf(clientCtx.Output, "Note: Full query implementation requires gRPC query server\n")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryModuleCircuit creates a command to query a specific module's circuit
func CmdQueryModuleCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "module [module-name]",
		Short: "Query a specific module's circuit breaker state",
		Long: `Query the circuit breaker state for a specific module.

Example:
  statesetd query circuit module stablecoin
  statesetd query circuit module settlement`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			moduleName := args[0]
			fmt.Fprintf(clientCtx.Output, "Query circuit state for module: %s\n", moduleName)
			fmt.Fprintf(clientCtx.Output, "Note: Full query implementation requires gRPC query server\n")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryRateLimits creates a command to query rate limits
func CmdQueryRateLimits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rate-limits [address]",
		Short: "Query rate limit status for an address",
		Long: `Query the current rate limit status for a specific address.

Example:
  statesetd query circuit rate-limits stateset1abc...`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			address := args[0]
			fmt.Fprintf(clientCtx.Output, "Query rate limits for address: %s\n", address)
			fmt.Fprintf(clientCtx.Output, "Note: Full query implementation requires gRPC query server\n")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryLiquidationProtection creates a command to query liquidation protection
func CmdQueryLiquidationProtection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "liquidation-protection",
		Short: "Query the current liquidation surge protection status",
		Long: `Query the liquidation surge protection state including current block limits.

Example:
  statesetd query circuit liquidation-protection`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			fmt.Fprintf(clientCtx.Output, "Query liquidation surge protection state\n")
			fmt.Fprintf(clientCtx.Output, "Note: Full query implementation requires gRPC query server\n")
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
