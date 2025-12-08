package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/zkpverify/types"
)

// GetQueryCmd returns the query commands for the zkpverify module
func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Query commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdQueryCircuit(),
		CmdQueryCircuits(),
		CmdQueryProof(),
		CmdQueryVerificationResult(),
		CmdQuerySymbolicRules(),
		CmdQueryParams(),
	)

	return cmd
}

// CmdQueryCircuit returns a command to query a circuit by name
func CmdQueryCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circuit [name]",
		Short: "Query a circuit by name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			fmt.Printf("Querying circuit: %s\n", args[0])
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryCircuits returns a command to query all circuits
func CmdQueryCircuits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circuits",
		Short: "Query all registered circuits",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			activeOnly, _ := cmd.Flags().GetBool("active-only")
			fmt.Printf("Querying circuits (active-only: %v)\n", activeOnly)
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	cmd.Flags().Bool("active-only", false, "Only return active circuits")
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryProof returns a command to query a proof by ID
func CmdQueryProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proof [proof-id]",
		Short: "Query a proof by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			proofID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid proof ID: %w", err)
			}

			fmt.Printf("Querying proof: %d\n", proofID)
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryVerificationResult returns a command to query verification result
func CmdQueryVerificationResult() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "result [proof-id]",
		Short: "Query verification result for a proof",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			proofID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid proof ID: %w", err)
			}

			fmt.Printf("Querying verification result for proof: %d\n", proofID)
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQuerySymbolicRules returns a command to query symbolic rules for a circuit
func CmdQuerySymbolicRules() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rules [circuit-name]",
		Short: "Query symbolic rules for a circuit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			fmt.Printf("Querying symbolic rules for circuit: %s\n", args[0])
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// CmdQueryParams returns a command to query module parameters
func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			fmt.Println("Querying zkpverify parameters")
			fmt.Println("Note: Full gRPC client integration required for actual query")
			_ = clientCtx
			return nil
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
