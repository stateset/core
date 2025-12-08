package settlement

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/settlement/keeper"
	"github.com/stateset/core/x/settlement/types"
)

// InitGenesis initializes the settlement module's state from a provided genesis state
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.InitGenesis(ctx, &genState)
}

// ExportGenesis returns the settlement module's exported genesis state
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
