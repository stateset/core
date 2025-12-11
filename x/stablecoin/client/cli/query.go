package cli

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/stablecoin/types"
)

// NewQueryCmd returns the root query command for stablecoin.
func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Stablecoin query subcommands",
		Aliases:                    []string{"stablecoin", "stablecoins"},
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewGetVaultCmd(),
		NewGetReserveParamsCmd(),
		NewGetReserveCmd(),
		NewGetReserveDepositCmd(),
		NewGetRedemptionRequestCmd(),
		NewGetAttestationCmd(),
		NewGetLatestAttestationCmd(),
		NewGetDailyStatsCmd(),
	)

	return cmd
}

// NewGetVaultCmd retrieves vault information.
func NewGetVaultCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vault [id]",
		Short: "Query a vault by id",
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

			key := types.VaultStoreKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("vault %d not found", id)
			}

			var vault types.Vault
			types.ModuleCdc.MustUnmarshalJSON(res, &vault)

			return clientCtx.PrintObjectLegacy(vault)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetReserveParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserve-params",
		Short: "Query reserve parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(types.ReserveParamsKey, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("reserve params not found")
			}

			var params types.ReserveParams
			types.ModuleCdc.MustUnmarshalJSON(res, &params)
			return clientCtx.PrintObjectLegacy(params)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetReserveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserve",
		Short: "Query on-chain reserve state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(types.ReserveKey, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("reserve not found")
			}

			var reserve types.Reserve
			types.ModuleCdc.MustUnmarshalJSON(res, &reserve)
			return clientCtx.PrintObjectLegacy(reserve)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetReserveDepositCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserve-deposit [id]",
		Short: "Query a reserve deposit by id",
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

			key := types.ReserveDepositKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("reserve deposit %d not found", id)
			}

			var deposit types.ReserveDeposit
			types.ModuleCdc.MustUnmarshalJSON(res, &deposit)
			return clientCtx.PrintObjectLegacy(deposit)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetRedemptionRequestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redemption [id]",
		Short: "Query a redemption request by id",
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

			key := types.RedemptionRequestKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("redemption request %d not found", id)
			}

			var redemption types.RedemptionRequest
			types.ModuleCdc.MustUnmarshalJSON(res, &redemption)
			return clientCtx.PrintObjectLegacy(redemption)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attestation [id]",
		Short: "Query an off-chain reserve attestation by id",
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

			key := types.OffChainAttestationKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("attestation %d not found", id)
			}

			var att types.OffChainReserveAttestation
			types.ModuleCdc.MustUnmarshalJSON(res, &att)
			return clientCtx.PrintObjectLegacy(att)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetLatestAttestationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "latest-attestation",
		Short: "Query the latest off-chain reserve attestation",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			nextBz, _, err := clientCtx.QueryStore(types.NextAttestationIDKey, types.StoreKey)
			if err != nil {
				return err
			}
			if len(nextBz) == 0 {
				return fmt.Errorf("no attestations found")
			}

			next := binary.BigEndian.Uint64(nextBz)
			if next <= 1 {
				return fmt.Errorf("no attestations found")
			}
			id := next - 1

			key := types.OffChainAttestationKey(id)
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("attestation %d not found", id)
			}

			var att types.OffChainReserveAttestation
			types.ModuleCdc.MustUnmarshalJSON(res, &att)
			return clientCtx.PrintObjectLegacy(att)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetDailyStatsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "daily-stats [yyyy-mm-dd]",
		Short: "Query daily mint/redeem stats for a date",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			key := types.DailyMintStatsKey(args[0])
			res, _, err := clientCtx.QueryStore(key, types.StoreKey)
			if err != nil {
				return err
			}
			if len(res) == 0 {
				return fmt.Errorf("daily stats for %s not found", args[0])
			}

			var stats types.DailyMintStats
			types.ModuleCdc.MustUnmarshalJSON(res, &stats)
			return clientCtx.PrintObjectLegacy(stats)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}
