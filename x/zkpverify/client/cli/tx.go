package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/zkpverify/types"
)

// GetTxCmd returns the transaction commands for the zkpverify module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transaction subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdRegisterCircuit(),
		CmdDeactivateCircuit(),
		CmdRegisterSymbolicRule(),
		CmdSubmitProof(),
		CmdChallengeProof(),
	)

	return cmd
}

// CmdRegisterCircuit returns a command to register a new circuit
func CmdRegisterCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-circuit [name] [vk-file] [schema-file]",
		Short: "Register a new STARK verification circuit",
		Long: `Register a new STARK verification circuit with a verification key and public input schema.

Arguments:
  name        - Unique name for the circuit
  vk-file     - Path to file containing the verification key (hex encoded)
  schema-file - Path to JSON file containing the public input schema

Example schema file:
[
  {"name": "input_hash", "type": "hash", "required": true},
  {"name": "output_value", "type": "uint64", "required": true}
]`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]

			// Read verification key
			vkHex, err := os.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("failed to read verification key file: %w", err)
			}
			vk, err := hex.DecodeString(strings.TrimSpace(string(vkHex)))
			if err != nil {
				return fmt.Errorf("failed to decode verification key: %w", err)
			}

			// Read schema
			schemaBytes, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read schema file: %w", err)
			}
			var schema []types.PublicInputField
			if err := json.Unmarshal(schemaBytes, &schema); err != nil {
				return fmt.Errorf("failed to parse schema: %w", err)
			}

			description, _ := cmd.Flags().GetString("description")
			maxRecursion, _ := cmd.Flags().GetUint32("max-recursion")

			msg := &types.MsgRegisterCircuit{
				Authority:         clientCtx.GetFromAddress().String(),
				Name:              name,
				VerificationKey:   vk,
				PublicInputSchema: schema,
				Description:       description,
				MaxRecursionDepth: maxRecursion,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("description", "", "Description of the circuit")
	cmd.Flags().Uint32("max-recursion", 16, "Maximum recursion depth for proofs")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdDeactivateCircuit returns a command to deactivate a circuit
func CmdDeactivateCircuit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deactivate-circuit [name]",
		Short: "Deactivate an existing circuit",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgDeactivateCircuit{
				Authority:   clientCtx.GetFromAddress().String(),
				CircuitName: args[0],
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdRegisterSymbolicRule returns a command to register a symbolic rule
func CmdRegisterSymbolicRule() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-rule [circuit-name] [rule-name] [rule-file]",
		Short: "Register a symbolic logic rule for a circuit",
		Long: `Register a symbolic logic rule that proofs must satisfy.

Example rule file:
{
  "rule_type": "implication",
  "conditions": [
    {"field": "amount", "operator": "gt", "value": "0"},
    {"field": "approved", "operator": "eq", "value": "true"}
  ],
  "description": "If amount > 0, then approved must be true"
}

Rule types: implication, conjunction, disjunction, equality, inequality, comparison, set_membership
Operators: eq, neq, gt, lt, gte, lte, in, not_in`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			circuitName := args[0]
			ruleName := args[1]

			// Read rule definition
			ruleBytes, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read rule file: %w", err)
			}

			var ruleDef struct {
				RuleType    types.RuleType    `json:"rule_type"`
				Conditions  []types.Condition `json:"conditions"`
				Description string            `json:"description"`
			}
			if err := json.Unmarshal(ruleBytes, &ruleDef); err != nil {
				return fmt.Errorf("failed to parse rule: %w", err)
			}

			msg := &types.MsgRegisterSymbolicRule{
				Authority:   clientCtx.GetFromAddress().String(),
				CircuitName: circuitName,
				RuleName:    ruleName,
				RuleType:    ruleDef.RuleType,
				Conditions:  ruleDef.Conditions,
				Description: ruleDef.Description,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// CmdSubmitProof returns a command to submit a proof for verification
func CmdSubmitProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proof [circuit-name] [proof-file] [inputs-file] [commitment-hex]",
		Short: "Submit a STARK proof for verification",
		Long: `Submit a STARK proof for on-chain verification.

Arguments:
  circuit-name   - Name of the registered circuit
  proof-file     - Path to file containing the proof (hex encoded)
  inputs-file    - Path to JSON file containing public inputs
  commitment-hex - Hex-encoded data commitment (Merkle root of off-chain data)

Optional flags:
  --recursive    - Comma-separated list of proof IDs this proof aggregates

Example inputs file:
{
  "fields": {
    "input_hash": "abc123...",
    "output_value": 42,
    "data_commitment": "def456..."
  }
}`,
		Args: cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			circuitName := args[0]

			// Read proof
			proofHex, err := os.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("failed to read proof file: %w", err)
			}
			proofData, err := hex.DecodeString(strings.TrimSpace(string(proofHex)))
			if err != nil {
				return fmt.Errorf("failed to decode proof: %w", err)
			}

			// Read public inputs
			inputsBytes, err := os.ReadFile(args[2])
			if err != nil {
				return fmt.Errorf("failed to read inputs file: %w", err)
			}

			// Decode data commitment
			commitment, err := hex.DecodeString(args[3])
			if err != nil {
				return fmt.Errorf("failed to decode commitment: %w", err)
			}

			// Parse recursive proofs if provided
			var recursiveProofs []uint64
			recursiveStr, _ := cmd.Flags().GetString("recursive")
			if recursiveStr != "" {
				for _, idStr := range strings.Split(recursiveStr, ",") {
					id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 64)
					if err != nil {
						return fmt.Errorf("invalid recursive proof ID: %s", idStr)
					}
					recursiveProofs = append(recursiveProofs, id)
				}
			}

			msg := &types.MsgSubmitProof{
				Submitter:       clientCtx.GetFromAddress().String(),
				CircuitName:     circuitName,
				ProofData:       proofData,
				PublicInputs:    inputsBytes,
				DataCommitment:  commitment,
				RecursiveProofs: recursiveProofs,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("recursive", "", "Comma-separated proof IDs for recursive aggregation")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdChallengeProof returns a command to challenge a proof
func CmdChallengeProof() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "challenge-proof [proof-id] [fraud-proof-file]",
		Short: "Challenge a verified proof with fraud evidence",
		Long: `Submit a challenge against a previously verified proof.

The fraud proof should demonstrate that the original proof was invalid.
Challenges must be submitted within the challenge window (typically 24 hours).`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proofID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("invalid proof ID: %w", err)
			}

			// Read fraud proof
			fraudProofHex, err := os.ReadFile(args[1])
			if err != nil {
				return fmt.Errorf("failed to read fraud proof file: %w", err)
			}
			fraudProof, err := hex.DecodeString(strings.TrimSpace(string(fraudProofHex)))
			if err != nil {
				return fmt.Errorf("failed to decode fraud proof: %w", err)
			}

			reason, _ := cmd.Flags().GetString("reason")

			msg := &types.MsgChallengeProof{
				Challenger: clientCtx.GetFromAddress().String(),
				ProofID:    proofID,
				FraudProof: fraudProof,
				Reason:     reason,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("reason", "", "Reason for the challenge")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
