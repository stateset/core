package cli

import (
	"github.com/spf13/cobra"
	
	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for the oracle module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "oracle",
		Short:                      "Oracle transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	
	cmd.AddCommand(
		NewSubmitPriceFeedCmd(),
		NewRegisterOracleCmd(),
		NewUpdateOracleCmd(),
		NewRemoveOracleCmd(),
		NewRequestPriceCmd(),
	)
	
	return cmd
}

// GetQueryCmd returns the cli query commands for the oracle module
func GetQueryCmd(queryRoute string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "oracle",
		Short:                      "Querying commands for the oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	
	cmd.AddCommand(
		CmdQueryPrice(),
		CmdQueryPriceFeed(),
		CmdQueryOracle(),
		CmdListOracles(),
		CmdQueryPriceHistory(),
		CmdQueryParams(),
	)
	
	return cmd
}