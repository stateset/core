package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/purchaseorder/keeper"
	"github.com/stateset/core/x/purchaseorder/types"
	"github.com/stretchr/testify/require"
)

func createNSentPurchaseorder(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentPurchaseorder {
	items := make([]types.SentPurchaseorder, n)
	for i := range items {
		items[i].Id = keeper.AppendSentPurchaseorder(ctx, items[i])
	}
	return items
}

func TestSentPurchaseorderGet(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNSentPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentPurchaseorder(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestSentPurchaseorderRemove(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNSentPurchaseorder(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentPurchaseorder(ctx, item.Id)
		_, found := keeper.GetSentPurchaseorder(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentPurchaseorderGetAll(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNSentPurchaseorder(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllSentPurchaseorder(ctx))
}

func TestSentPurchaseorderCount(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	items := createNSentPurchaseorder(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentPurchaseorderCount(ctx))
}
