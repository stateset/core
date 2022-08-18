package refund_test

import (
	"testing"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/testutil/nullify"
	"github.com/stateset/core/x/refund"
	"github.com/stateset/core/x/refund/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RefundList: []types.Refund{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RefundCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RefundKeeper(t)
	refund.InitGenesis(ctx, *k, genesisState)
	got := refund.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RefundList, got.RefundList)
	require.Equal(t, genesisState.RefundCount, got.RefundCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
