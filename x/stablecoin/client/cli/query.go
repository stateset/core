package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		NewGetParamsCmd(),
		NewGetVaultCmd(),
		NewGetVaultsCmd(),
		NewGetReserveParamsCmd(),
		NewGetReserveCmd(),
		NewGetLockedReservesCmd(),
		NewGetTotalReservesCmd(),
		NewGetReserveDepositCmd(),
		NewGetReserveDepositsCmd(),
		NewGetRedemptionRequestCmd(),
		NewGetRedemptionRequestsCmd(),
		NewGetAttestationCmd(),
		NewGetLatestAttestationCmd(),
		NewGetDailyStatsCmd(),
	)

	return cmd
}

func NewGetParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query vault (CDP) parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Vault(context.Background(), &types.QueryVaultRequest{VaultId: id})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetVaultsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vaults [owner]",
		Short: "Query vaults (optionally filtered by owner)",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			owner := ""
			if len(args) == 1 {
				owner = args[0]
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Vaults(context.Background(), &types.QueryVaultsRequest{Owner: owner})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ReserveParams(context.Background(), &types.QueryReserveParamsRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Reserve(context.Background(), &types.QueryReserveRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetLockedReservesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "locked-reserves",
		Short: "Query reserves locked for pending redemptions",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			res, _, err := clientCtx.QueryStore(types.LockedReservesKey, types.StoreKey)
			if err != nil {
				return err
			}

			locked := sdk.NewCoins()
			if len(res) > 0 {
				if err := json.Unmarshal(res, &locked); err != nil {
					return err
				}
			}

			return clientCtx.PrintObjectLegacy(locked)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetTotalReservesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-reserves",
		Short: "Query combined on-chain and off-chain reserves",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.TotalReserves(context.Background(), &types.QueryTotalReservesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ReserveDeposit(context.Background(), &types.QueryReserveDepositRequest{DepositId: id})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetReserveDepositsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserve-deposits [depositor]",
		Short: "Query reserve deposit records (optionally filtered by depositor)",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			depositor := ""
			if len(args) == 1 {
				depositor = args[0]
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.ReserveDeposits(context.Background(), &types.QueryReserveDepositsRequest{Depositor: depositor})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.RedemptionRequest(context.Background(), &types.QueryRedemptionRequestRequest{RedemptionId: id})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func NewGetRedemptionRequestsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "redemptions [status]",
		Short: "Query redemption requests (optionally filtered by status)",
		Args:  cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			status := ""
			if len(args) == 1 {
				status = args[0]
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.RedemptionRequests(context.Background(), &types.QueryRedemptionRequestsRequest{Status: status})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Attestation(context.Background(), &types.QueryAttestationRequest{AttestationId: id})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
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

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.LatestAttestation(context.Background(), &types.QueryLatestAttestationRequest{})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
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
