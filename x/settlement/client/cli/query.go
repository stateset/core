package cli

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/settlement/types"
)

// NewQueryCmd returns the root query command for settlement.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Settlement query subcommands",
		Aliases:                    []string{"settlement", "settlements"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetSettlementCmd(),
		NewGetBatchCmd(),
		NewGetChannelCmd(),
		NewGetMerchantCmd(),
		NewGetParamsCmd(),
	)

	return cmd
}

func settlementKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, types.SettlementKeyPrefix...), bz...)
}

func batchKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, types.BatchKeyPrefix...), bz...)
}

func channelKey(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return append(append([]byte{}, types.ChannelKeyPrefix...), bz...)
}

func merchantKey(addr string) []byte {
	return append(append([]byte{}, types.MerchantKeyPrefix...), []byte(addr)...)
}

func NewGetSettlementCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "settlement [id]",
		Short: "Query a settlement by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(settlementKey(id), types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("settlement %d not found", id)
			}

			var settlement types.Settlement
			types.ModuleCdc.MustUnmarshalJSON(res, &settlement)
			return clientCtx.PrintObjectLegacy(settlement)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetBatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batch [id]",
		Short: "Query a batch settlement by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(batchKey(id), types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("batch %d not found", id)
			}

			var batch types.BatchSettlement
			types.ModuleCdc.MustUnmarshalJSON(res, &batch)
			return clientCtx.PrintObjectLegacy(batch)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetChannelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "channel [id]",
		Short: "Query a payment channel by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(channelKey(id), types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("channel %d not found", id)
			}

			var channel types.PaymentChannel
			types.ModuleCdc.MustUnmarshalJSON(res, &channel)
			return clientCtx.PrintObjectLegacy(channel)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetMerchantCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "merchant [address]",
		Short: "Query a merchant configuration by address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			addr := args[0]
			res, _, err := clientCtx.QueryStore(merchantKey(addr), types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("merchant %s not found", addr)
			}

			var merchant types.MerchantConfig
			types.ModuleCdc.MustUnmarshalJSON(res, &merchant)
			return clientCtx.PrintObjectLegacy(merchant)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query settlement module parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(types.ParamsKey, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("params not found")
			}

			var params types.Params
			types.ModuleCdc.MustUnmarshalJSON(res, &params)
			return clientCtx.PrintObjectLegacy(params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
