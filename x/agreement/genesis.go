package agreement

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/agreement/keeper"
	"github.com/stateset/core/x/agreement/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the sentAgreement
	for _, elem := range genState.SentAgreementList {
		k.SetSentAgreement(ctx, elem)
	}

	// Set sentAgreement count
	k.SetSentAgreementCount(ctx, genState.SentAgreementCount)
	// Set all the timedoutAgreement
	for _, elem := range genState.TimedoutAgreementList {
		k.SetTimedoutAgreement(ctx, elem)
	}

	// Set timedoutAgreement count
	k.SetTimedoutAgreementCount(ctx, genState.TimedoutAgreementCount)
	// Set all the agreement
	for _, elem := range genState.AgreementList {
		k.SetAgreement(ctx, elem)
	}

	// Set agreement count
	k.SetAgreementCount(ctx, genState.AgreementCount)
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.SentAgreementList = k.GetAllSentAgreement(ctx)
	genesis.SentAgreementCount = k.GetSentAgreementCount(ctx)
	genesis.TimedoutAgreementList = k.GetAllTimedoutAgreement(ctx)
	genesis.TimedoutAgreementCount = k.GetTimedoutAgreementCount(ctx)
	genesis.AgreementList = k.GetAllAgreement(ctx)
	genesis.AgreementCount = k.GetAgreementCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
