package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/invoice/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group invoice queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdListInvoice())
	cmd.AddCommand(CmdShowInvoice())
	cmd.AddCommand(CmdListSentInvoice())
	cmd.AddCommand(CmdShowSentInvoice())
	cmd.AddCommand(CmdListTimedoutInvoice())
	cmd.AddCommand(CmdShowTimedoutInvoice())
	// this line is used by starport scaffolding # 1

	return cmd
}
