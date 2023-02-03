package proof_test

import (
	"testing"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/testutil/nullify"
	"github.com/stateset/core/x/proof"
	"github.com/stateset/core/x/proof/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProofKeeper(t)
	proof.InitGenesis(ctx, *k, genesisState)
	got := proof.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
