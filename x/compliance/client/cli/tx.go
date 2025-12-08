package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/compliance/types"
)

const (
	flagMetadata = "metadata"
	flagSanction = "sanction"
	flagReason   = "reason"
)

// NewTxCmd builds the root tx command for compliance operations.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Compliance transaction subcommands",
		Aliases:                    []string{"compliance"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewUpsertProfileCmd(),
		NewSetSanctionCmd(),
	)

	return cmd
}

// NewUpsertProfileCmd creates or updates a compliance profile.
func NewUpsertProfileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upsert-profile [address] [kyc-level] [risk-level]",
		Short: "Create or update a compliance profile",
		Long: `Create or update compliance data for an address.
Risk level must be one of: low, medium, high.`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr := args[0]
			if _, err := sdk.AccAddressFromBech32(addr); err != nil {
				return err
			}

			kycLevel := types.KYCLevel(args[1])
			switch kycLevel {
			case types.KYCNone, types.KYCBasic, types.KYCStandard, types.KYCEnhanced:
			default:
				return fmt.Errorf("invalid kyc level: %s (expected none|basic|standard|enhanced)", args[1])
			}

			risk := types.RiskLevel(args[2])
			switch risk {
			case types.RiskLow, types.RiskMedium, types.RiskHigh:
			default:
				return fmt.Errorf("invalid risk level: %s (expected low|medium|high)", args[2])
			}

			sanction, err := cmd.Flags().GetBool(flagSanction)
			if err != nil {
				return err
			}
			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			profile := types.Profile{
				Address:   addr,
				KYCLevel:  kycLevel,
				Risk:      risk,
				Status:    types.StatusActive,
				Sanction:  sanction,
				Metadata:  metadata,
				UpdatedBy: clientCtx.GetFromAddress().String(),
			}

			msg := types.NewMsgUpsertProfile(clientCtx.GetFromAddress().String(), profile)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(flagSanction, false, "Set sanction flag on the profile")
	cmd.Flags().String(flagMetadata, "", "Optional metadata note for the profile")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// NewSetSanctionCmd toggles the sanction flag for a profile.
func NewSetSanctionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-sanction [address] [true|false]",
		Short: "Update the sanction status of a profile",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			addr := args[0]
			if _, err := sdk.AccAddressFromBech32(addr); err != nil {
				return err
			}

			sanction, err := parseBool(args[1])
			if err != nil {
				return err
			}

			reason, err := cmd.Flags().GetString(flagReason)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetSanction(clientCtx.GetFromAddress().String(), addr, sanction, reason)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReason, "", "Optional reason for the sanction update")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func parseBool(input string) (bool, error) {
	switch input {
	case "true", "True", "TRUE", "1":
		return true, nil
	case "false", "False", "FALSE", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", input)
	}
}
