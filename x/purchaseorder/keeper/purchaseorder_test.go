package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/purchaseorder/keeper"
	"github.com/stateset/core/x/purchaseorder/types"
	"github.com/stretchr/testify/require"
)

func createNPurchaseorder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Purchaseorder {
	items := make([]types.Purchaseorder, n)
	for i := range items {
		items[i].Id = keeper.AppendPurchaseorder(ctx, items[i])
	}
	return items
}

func TestPurchaseorderGet(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPurchaseorder(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestPurchaseorderRemove(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePurchaseorder(ctx, item.Id)
		_, found := keeper.GetPurchaseorder(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPurchaseorderGetAll(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNPurchaseorder(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllPurchaseorder(ctx))
}

func TestPurchaseorderCount(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNPurchaseorder(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPurchaseorderCount(ctx))
}
