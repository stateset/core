package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

// InitGenesis initializes the oracle module state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data *types.GenesisState) {
	if data == nil {
		data = types.DefaultGenesis()
	}
	k.InitGenesis(ctx, data)
}

// ExportGenesis exports the oracle module state.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return k.ExportGenesis(ctx)
}
