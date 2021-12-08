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
	// Set all the sentInvoice
	for _, elem := range genState.SentInvoiceList {
		k.SetSentInvoice(ctx, elem)
	}

	// Set sentInvoice count
	k.SetSentInvoiceCount(ctx, genState.SentInvoiceCount)
	// Set all the timedoutInvoice
	for _, elem := range genState.TimedoutInvoiceList {
		k.SetTimedoutInvoice(ctx, elem)
	}

	// Set timedoutInvoice count
	k.SetTimedoutInvoiceCount(ctx, genState.TimedoutInvoiceCount)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.InvoiceList = k.GetAllInvoice(ctx)
	genesis.InvoiceCount = k.GetInvoiceCount(ctx)
	genesis.SentInvoiceList = k.GetAllSentInvoice(ctx)
	genesis.SentInvoiceCount = k.GetSentInvoiceCount(ctx)
	genesis.TimedoutInvoiceList = k.GetAllTimedoutInvoice(ctx)
	genesis.TimedoutInvoiceCount = k.GetTimedoutInvoiceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
