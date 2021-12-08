package purchaseorder_test

import (
	"testing"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/purchaseorder"
	"github.com/stateset/core/x/purchaseorder/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		PurchaseorderList: []types.Purchaseorder{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PurchaseorderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PurchaseorderKeeper(t)
	purchaseorder.InitGenesis(ctx, *k, genesisState)
	got := purchaseorder.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.PurchaseorderList, len(genesisState.PurchaseorderList))
	require.Subset(t, genesisState.PurchaseorderList, got.PurchaseorderList)
	require.Equal(t, genesisState.PurchaseorderCount, got.PurchaseorderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
