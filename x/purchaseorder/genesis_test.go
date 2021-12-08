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
		SentPurchaseorderList: []types.SentPurchaseorder{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentPurchaseorderCount: 2,
		TimedoutPurchaseorderList: []types.TimedoutPurchaseorder{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimedoutPurchaseorderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PurchaseorderKeeper(t)
	purchaseorder.InitGenesis(ctx, *k, genesisState)
	got := purchaseorder.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.PurchaseorderList, len(genesisState.PurchaseorderList))
	require.Subset(t, genesisState.PurchaseorderList, got.PurchaseorderList)
	require.Equal(t, genesisState.PurchaseorderCount, got.PurchaseorderCount)
	require.Len(t, got.SentPurchaseorderList, len(genesisState.SentPurchaseorderList))
	require.Subset(t, genesisState.SentPurchaseorderList, got.SentPurchaseorderList)
	require.Equal(t, genesisState.SentPurchaseorderCount, got.SentPurchaseorderCount)
	require.Len(t, got.TimedoutPurchaseorderList, len(genesisState.TimedoutPurchaseorderList))
	require.Subset(t, genesisState.TimedoutPurchaseorderList, got.TimedoutPurchaseorderList)
	require.Equal(t, genesisState.TimedoutPurchaseorderCount, got.TimedoutPurchaseorderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
