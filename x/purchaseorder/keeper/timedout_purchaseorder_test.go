package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/purchaseorder/keeper"
	"github.com/stateset/core/x/purchaseorder/types"
	"github.com/stretchr/testify/require"
)

func createNTimedoutPurchaseorder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutPurchaseorder {
	items := make([]types.TimedoutPurchaseorder, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutPurchaseorder(ctx, items[i])
	}
	return items
}

func TestTimedoutPurchaseorderGet(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNTimedoutPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutPurchaseorder(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestTimedoutPurchaseorderRemove(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNTimedoutPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutPurchaseorder(ctx, item.Id)
		_, found := keeper.GetTimedoutPurchaseorder(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutPurchaseorderGetAll(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNTimedoutPurchaseorder(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllTimedoutPurchaseorder(ctx))
}

func TestTimedoutPurchaseorderCount(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNTimedoutPurchaseorder(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutPurchaseorderCount(ctx))
}
