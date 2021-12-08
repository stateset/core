package invoice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/keeper"
	"github.com/stateset/core/x/invoice/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the invoice
	for _, elem := range genState.InvoiceList {
		k.SetInvoice(ctx, elem)
	}

	// Set invoice count
	k.SetInvoiceCount(ctx, genState.InvoiceCount)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.InvoiceList = k.GetAllInvoice(ctx)
	genesis.InvoiceCount = k.GetInvoiceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
