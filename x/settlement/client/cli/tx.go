package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/settlement/types"
)

const (
	flagMetadata       = "metadata"
	flagReference      = "reference"
	flagReason         = "reason"
	flagUseEscrow      = "use-escrow"
	flagFeeRateBps     = "fee-rate-bps"
	flagMinSettlement  = "min-settlement"
	flagMaxSettlement  = "max-settlement"
	flagBatchEnabled   = "batch-enabled"
	flagBatchThreshold = "batch-threshold"
	flagWebhookURL     = "webhook-url"
	flagName           = "name"
	flagIsActive       = "is-active"
)

// NewTxCmd returns the root tx command for settlement operations.
func NewTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Settlement transaction subcommands",
		Aliases:                    []string{"settlement", "settlements"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewInstantTransferCmd(),
		NewCreateEscrowCmd(),
		NewReleaseEscrowCmd(),
		NewRefundEscrowCmd(),
		NewOpenChannelCmd(),
		NewCloseChannelCmd(),
		NewClaimChannelCmd(),
		NewInstantCheckoutCmd(),
		NewPartialRefundCmd(),
		NewRegisterMerchantCmd(),
		NewUpdateMerchantCmd(),
		NewCreateBatchCmd(),
		NewSettleBatchCmd(),
	)

	return cmd
}

func NewInstantTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instant-transfer [recipient] [amount] [reference]",
		Short: "Send an instant ssUSD transfer",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient := args[0]
			if _, err := sdk.AccAddressFromBech32(recipient); err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			reference := ""
			if len(args) == 3 {
				reference = args[2]
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			msg := types.NewMsgInstantTransfer(clientCtx.GetFromAddress().String(), recipient, amount, reference, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagMetadata, "", "Optional metadata for the transfer")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCreateEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-escrow [recipient] [amount] [expires-in]",
		Short: "Create an escrow settlement",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient := args[0]
			if _, err := sdk.AccAddressFromBech32(recipient); err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			expiresIn, err := time.ParseDuration(args[2])
			if err != nil {
				seconds, serr := strconv.ParseInt(args[2], 10, 64)
				if serr != nil {
					return err
				}
				expiresIn = time.Duration(seconds) * time.Second
			}

			reference, err := cmd.Flags().GetString(flagReference)
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateEscrow(clientCtx.GetFromAddress().String(), recipient, amount, reference, metadata, expiresIn)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReference, "", "Optional settlement reference")
	cmd.Flags().String(flagMetadata, "", "Optional metadata")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewReleaseEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release-escrow [settlement-id]",
		Short: "Release escrowed funds to the recipient",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgReleaseEscrow(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRefundEscrowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refund-escrow [settlement-id]",
		Short: "Refund escrowed funds back to the sender",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			reason, err := cmd.Flags().GetString(flagReason)
			if err != nil {
				return err
			}

			msg := types.NewMsgRefundEscrow(clientCtx.GetFromAddress().String(), id, reason)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReason, "", "Optional refund reason")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewOpenChannelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "open-channel [recipient] [deposit] [expires-in-blocks]",
		Short: "Open a payment channel with an ssUSD deposit",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient := args[0]
			if _, err := sdk.AccAddressFromBech32(recipient); err != nil {
				return err
			}

			deposit, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			expires, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgOpenChannel(clientCtx.GetFromAddress().String(), recipient, deposit, expires)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCloseChannelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close-channel [channel-id]",
		Short: "Close an open payment channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCloseChannel(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewClaimChannelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-channel [channel-id] [amount] [nonce] [signature]",
		Short: "Claim funds from a payment channel (recipient)",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			channelID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			nonce, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			signature := args[3]

			msg := types.NewMsgClaimChannel(clientCtx.GetFromAddress().String(), channelID, amount, nonce, signature)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewInstantCheckoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instant-checkout [merchant] [amount] [order-reference]",
		Short: "Perform a streamlined ecommerce checkout",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merchant := args[0]
			if _, err := sdk.AccAddressFromBech32(merchant); err != nil {
				return err
			}

			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			orderRef := args[2]

			useEscrow, err := cmd.Flags().GetBool(flagUseEscrow)
			if err != nil {
				return err
			}

			metadata, err := cmd.Flags().GetString(flagMetadata)
			if err != nil {
				return err
			}

			msg := types.NewMsgInstantCheckout(clientCtx.GetFromAddress().String(), merchant, amount, orderRef, useEscrow, nil, metadata)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Bool(flagUseEscrow, false, "Place funds in escrow instead of instant settlement")
	cmd.Flags().String(flagMetadata, "", "Optional metadata")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewPartialRefundCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "partial-refund [settlement-id] [refund-amount]",
		Short: "Issue a partial refund for a settlement (merchant/authority)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			refundAmount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			reason, err := cmd.Flags().GetString(flagReason)
			if err != nil {
				return err
			}

			msg := types.NewMsgPartialRefund(clientCtx.GetFromAddress().String(), id, refundAmount, reason)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagReason, "", "Refund reason")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewRegisterMerchantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-merchant [name]",
		Short: "Register a merchant configuration for settlements",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			feeRate, err := cmd.Flags().GetUint32(flagFeeRateBps)
			if err != nil {
				return err
			}

			minStr, err := cmd.Flags().GetString(flagMinSettlement)
			if err != nil {
				return err
			}
			maxStr, err := cmd.Flags().GetString(flagMaxSettlement)
			if err != nil {
				return err
			}

			minSettlement := sdk.NewCoin(types.StablecoinDenom, sdkmath.ZeroInt())
			maxSettlement := sdk.NewCoin(types.StablecoinDenom, sdkmath.ZeroInt())
			if minStr != "" {
				minSettlement, err = sdk.ParseCoinNormalized(minStr)
				if err != nil {
					return err
				}
			}
			if maxStr != "" {
				maxSettlement, err = sdk.ParseCoinNormalized(maxStr)
				if err != nil {
					return err
				}
			}

			batchEnabled, err := cmd.Flags().GetBool(flagBatchEnabled)
			if err != nil {
				return err
			}

			thresholdStr, err := cmd.Flags().GetString(flagBatchThreshold)
			if err != nil {
				return err
			}
			batchThreshold := sdk.NewCoin(types.StablecoinDenom, sdkmath.ZeroInt())
			if thresholdStr != "" {
				batchThreshold, err = sdk.ParseCoinNormalized(thresholdStr)
				if err != nil {
					return err
				}
			}

			webhookURL, err := cmd.Flags().GetString(flagWebhookURL)
			if err != nil {
				return err
			}

			from := clientCtx.GetFromAddress().String()
			msg := types.NewMsgRegisterMerchant(from, from, name, feeRate, minSettlement, maxSettlement, batchEnabled, batchThreshold, webhookURL)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint32(flagFeeRateBps, 0, "Custom fee rate in basis points")
	cmd.Flags().String(flagMinSettlement, "", "Optional per-merchant min settlement coin (e.g. 10ssusd)")
	cmd.Flags().String(flagMaxSettlement, "", "Optional per-merchant max settlement coin (e.g. 10000ssusd)")
	cmd.Flags().Bool(flagBatchEnabled, false, "Enable batch settlements for this merchant")
	cmd.Flags().String(flagBatchThreshold, "", "Optional batch threshold coin (e.g. 1000ssusd)")
	cmd.Flags().String(flagWebhookURL, "", "Optional HTTPS webhook URL for off-chain notifications")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewUpdateMerchantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-merchant [merchant-address]",
		Short: "Update an existing merchant configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merchant := args[0]
			if _, err := sdk.AccAddressFromBech32(merchant); err != nil {
				return err
			}

			msg := types.NewMsgUpdateMerchant(clientCtx.GetFromAddress().String(), merchant)

			if cmd.Flags().Changed(flagName) {
				name, _ := cmd.Flags().GetString(flagName)
				msg.Name = name
			}
			if cmd.Flags().Changed(flagFeeRateBps) {
				rate, _ := cmd.Flags().GetUint32(flagFeeRateBps)
				msg.FeeRateBps = rate
			}
			if cmd.Flags().Changed(flagBatchEnabled) {
				enabled, _ := cmd.Flags().GetBool(flagBatchEnabled)
				msg.BatchEnabled = &enabled
			}
			if cmd.Flags().Changed(flagIsActive) {
				active, _ := cmd.Flags().GetBool(flagIsActive)
				msg.IsActive = &active
			}
			if cmd.Flags().Changed(flagWebhookURL) {
				url, _ := cmd.Flags().GetString(flagWebhookURL)
				msg.WebhookUrl = url
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagName, "", "Updated merchant name")
	cmd.Flags().Uint32(flagFeeRateBps, 0, "Updated fee rate in basis points")
	cmd.Flags().Bool(flagBatchEnabled, false, "Enable/disable batch settlements")
	cmd.Flags().Bool(flagIsActive, true, "Set merchant active status")
	cmd.Flags().String(flagWebhookURL, "", "Updated HTTPS webhook URL")
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCreateBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-batch [merchant] [senders] [amounts] [references]",
		Short: "Create a batch settlement (authority only)",
		Long:  "Create a batch settlement for multiple senders. Lists are comma-separated.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			merchant := args[0]
			senders := strings.Split(args[1], ",")
			amountStrs := strings.Split(args[2], ",")
			references := strings.Split(args[3], ",")

			if len(senders) != len(amountStrs) || len(senders) != len(references) {
				return fmt.Errorf("senders, amounts, and references must have same length")
			}

			amounts := make([]sdk.Coin, len(amountStrs))
			for i, amtStr := range amountStrs {
				amt, err := sdk.ParseCoinNormalized(strings.TrimSpace(amtStr))
				if err != nil {
					return err
				}
				amounts[i] = amt
			}

			msg := types.NewMsgCreateBatch(clientCtx.GetFromAddress().String(), merchant, senders, amounts, references)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewSettleBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settle-batch [batch-id]",
		Short: "Settle a pending batch settlement (authority only)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgSettleBatch(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
