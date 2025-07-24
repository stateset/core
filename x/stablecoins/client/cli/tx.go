package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/stablecoins/types"
)

var DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
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

	cmd.AddCommand(CmdCreateStablecoin())
	cmd.AddCommand(CmdUpdateStablecoin())
	cmd.AddCommand(CmdMintStablecoin())
	cmd.AddCommand(CmdBurnStablecoin())
	cmd.AddCommand(CmdPauseStablecoin())
	cmd.AddCommand(CmdUnpauseStablecoin())
	cmd.AddCommand(CmdUpdatePriceData())
	cmd.AddCommand(CmdWhitelistAddress())
	cmd.AddCommand(CmdBlacklistAddress())
	cmd.AddCommand(CmdRemoveFromWhitelist())
	cmd.AddCommand(CmdRemoveFromBlacklist())

	return cmd
}

func CmdCreateStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-stablecoin [denom] [name] [symbol] [decimals] [max-supply]",
		Short: "Create a new stablecoin",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argName := args[1]
			argSymbol := args[2]
			
			argDecimals, err := strconv.ParseUint(args[3], 10, 32)
			if err != nil {
				return err
			}

			argMaxSupply, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return fmt.Errorf("invalid max supply: %s", args[4])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get optional flags
			description, _ := cmd.Flags().GetString("description")
			targetAsset, _ := cmd.Flags().GetString("target-asset")
			targetPrice, _ := cmd.Flags().GetString("target-price")
			stabilityMechanism, _ := cmd.Flags().GetString("stability-mechanism")
			metadata, _ := cmd.Flags().GetString("metadata")

			// Create PegInfo if target asset is provided
			var pegInfo *types.PegInfo
			if targetAsset != "" {
				targetPriceDec := sdk.OneDec()
				if targetPrice != "" {
					targetPriceDec, err = sdk.NewDecFromStr(targetPrice)
					if err != nil {
						return fmt.Errorf("invalid target price: %s", targetPrice)
					}
				}

				pegInfo = &types.PegInfo{
					TargetAsset:  targetAsset,
					TargetPrice:  targetPriceDec,
					PriceTolerance: sdk.NewDecWithPrec(1, 2), // 1% default tolerance
				}
			}

			msg := types.NewMsgCreateStablecoin(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argName,
				argSymbol,
				uint32(argDecimals),
				description,
				argMaxSupply,
				pegInfo,
				nil, // reserve info
				stabilityMechanism,
				nil, // fee info
				nil, // access control
				metadata,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("description", "", "Description of the stablecoin")
	cmd.Flags().String("target-asset", "", "Target asset for pegging (e.g., USD, EUR)")
	cmd.Flags().String("target-price", "1.0", "Target price for pegging")
	cmd.Flags().String("stability-mechanism", "collateralized", "Stability mechanism")
	cmd.Flags().String("metadata", "", "Additional metadata")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-stablecoin [denom]",
		Short: "Update a stablecoin configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Get optional flags
			name, _ := cmd.Flags().GetString("name")
			description, _ := cmd.Flags().GetString("description")
			metadata, _ := cmd.Flags().GetString("metadata")

			msg := types.NewMsgUpdateStablecoin(
				clientCtx.GetFromAddress().String(),
				argDenom,
				name,
				description,
				nil, // peg info
				nil, // fee info
				nil, // access control
				metadata,
			)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String("name", "", "New name for the stablecoin")
	cmd.Flags().String("description", "", "New description for the stablecoin")
	cmd.Flags().String("metadata", "", "New metadata for the stablecoin")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMintStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-stablecoin [denom] [amount] [recipient]",
		Short: "Mint stablecoin tokens",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			
			argAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount: %s", args[1])
			}

			argRecipient := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintStablecoin(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argAmount,
				argRecipient,
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

func CmdBurnStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-stablecoin [denom] [amount]",
		Short: "Burn stablecoin tokens",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			
			argAmount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return fmt.Errorf("invalid amount: %s", args[1])
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnStablecoin(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argAmount,
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

func CmdPauseStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pause-stablecoin [denom] [operation] [reason]",
		Short: "Pause stablecoin operations",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argOperation := args[1]
			argReason := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPauseStablecoin(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argOperation,
				argReason,
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

func CmdUnpauseStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unpause-stablecoin [denom] [operation]",
		Short: "Unpause stablecoin operations",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argOperation := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgUnpauseStablecoin{
				Creator:   clientCtx.GetFromAddress().String(),
				Denom:    argDenom,
				Operation: argOperation,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePriceData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-price-data [denom] [price] [source]",
		Short: "Update price data for a stablecoin",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			
			argPrice, err := sdk.NewDecFromStr(args[1])
			if err != nil {
				return fmt.Errorf("invalid price: %s", args[1])
			}

			argSource := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgUpdatePriceData{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:  argDenom,
				Price:  argPrice,
				Source: argSource,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdWhitelistAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-address [denom] [address]",
		Short: "Add an address to the whitelist",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgWhitelistAddress{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:  argDenom,
				Address: argAddress,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdBlacklistAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blacklist-address [denom] [address] [reason]",
		Short: "Add an address to the blacklist",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddress := args[1]
			argReason := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgBlacklistAddress{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:  argDenom,
				Address: argAddress,
				Reason:  argReason,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveFromWhitelist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-from-whitelist [denom] [address]",
		Short: "Remove an address from the whitelist",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveFromWhitelist{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:  argDenom,
				Address: argAddress,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveFromBlacklist() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-from-blacklist [denom] [address]",
		Short: "Remove an address from the blacklist",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAddress := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveFromBlacklist{
				Creator: clientCtx.GetFromAddress().String(),
				Denom:  argDenom,
				Address: argAddress,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}