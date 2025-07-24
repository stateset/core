package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/stateset/core/x/stablecoins/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group stablecoins queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowStablecoin())
	cmd.AddCommand(CmdListStablecoin())
	cmd.AddCommand(CmdStablecoinsByIssuer())
	cmd.AddCommand(CmdStablecoinSupply())
	cmd.AddCommand(CmdPriceData())
	cmd.AddCommand(CmdReserveInfo())
	cmd.AddCommand(CmdIsWhitelisted())
	cmd.AddCommand(CmdIsBlacklisted())
	cmd.AddCommand(CmdStablecoinStats())
	cmd.AddCommand(CmdMintRequests())
	cmd.AddCommand(CmdBurnRequests())

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stablecoin [denom]",
		Short: "shows a stablecoin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]

			params := &types.QueryStablecoinRequest{
				Denom: argDenom,
			}

			res, err := queryClient.Stablecoin(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdListStablecoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stablecoin",
		Short: "list all stablecoin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStablecoinsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.Stablecoins(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdStablecoinsByIssuer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stablecoins-by-issuer [issuer]",
		Short: "Query stablecoins by issuer address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			argIssuer := args[0]

			params := &types.QueryStablecoinsByIssuerRequest{
				Issuer:     argIssuer,
				Pagination: pageReq,
			}

			res, err := queryClient.StablecoinsByIssuer(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdStablecoinSupply() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "supply [denom]",
		Short: "Query stablecoin supply information",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]

			params := &types.QueryStablecoinSupplyRequest{
				Denom: argDenom,
			}

			res, err := queryClient.StablecoinSupply(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdPriceData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price-data [denom]",
		Short: "Query price data for a stablecoin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]

			params := &types.QueryPriceDataRequest{
				Denom: argDenom,
			}

			res, err := queryClient.PriceData(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdReserveInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reserve-info [denom]",
		Short: "Query reserve information for a stablecoin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]

			params := &types.QueryReserveInfoRequest{
				Denom: argDenom,
			}

			res, err := queryClient.ReserveInfo(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdIsWhitelisted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-whitelisted [denom] [address]",
		Short: "Check if an address is whitelisted for a stablecoin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]
			argAddress := args[1]

			params := &types.QueryIsWhitelistedRequest{
				Denom:   argDenom,
				Address: argAddress,
			}

			res, err := queryClient.IsWhitelisted(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdIsBlacklisted() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-blacklisted [denom] [address]",
		Short: "Check if an address is blacklisted for a stablecoin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argDenom := args[0]
			argAddress := args[1]

			params := &types.QueryIsBlacklistedRequest{
				Denom:   argDenom,
				Address: argAddress,
			}

			res, err := queryClient.IsBlacklisted(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdStablecoinStats() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stats",
		Short: "Query stablecoin ecosystem statistics",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			timeRange, _ := cmd.Flags().GetString("time-range")

			params := &types.QueryStablecoinStatsRequest{
				TimeRange: timeRange,
			}

			res, err := queryClient.StablecoinStats(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("time-range", "day", "Time range for statistics (day, week, month, year)")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdMintRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-requests",
		Short: "Query pending mint requests",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			denom, _ := cmd.Flags().GetString("denom")
			status, _ := cmd.Flags().GetString("status")

			params := &types.QueryMintRequestsRequest{
				Denom:      denom,
				Status:     status,
				Pagination: pageReq,
			}

			res, err := queryClient.MintRequests(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("denom", "", "Filter by stablecoin denom")
	cmd.Flags().String("status", "", "Filter by request status")
	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdBurnRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-requests",
		Short: "Query pending burn requests",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			denom, _ := cmd.Flags().GetString("denom")
			status, _ := cmd.Flags().GetString("status")

			params := &types.QueryBurnRequestsRequest{
				Denom:      denom,
				Status:     status,
				Pagination: pageReq,
			}

			res, err := queryClient.BurnRequests(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	cmd.Flags().String("denom", "", "Filter by stablecoin denom")
	cmd.Flags().String("status", "", "Filter by request status")
	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}