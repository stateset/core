package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/testutil/nullify"
	"github.com/stateset/core/x/refund/keeper"
	"github.com/stateset/core/x/refund/types"
	"github.com/stretchr/testify/require"
)

func createNRefund(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Refund {
	items := make([]types.Refund, n)
	for i := range items {
		items[i].Id = keeper.AppendRefund(ctx, items[i])
	}
	return items
}

func TestRefundGet(t *testing.T) {
	keeper, ctx := keepertest.RefundKeeper(t)
	items := createNRefund(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRefund(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRefundRemove(t *testing.T) {
	keeper, ctx := keepertest.RefundKeeper(t)
	items := createNRefund(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRefund(ctx, item.Id)
		_, found := keeper.GetRefund(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRefundGetAll(t *testing.T) {
	keeper, ctx := keepertest.RefundKeeper(t)
	items := createNRefund(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRefund(ctx)),
	)
}

func TestRefundCount(t *testing.T) {
	keeper, ctx := keepertest.RefundKeeper(t)
	items := createNRefund(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRefundCount(ctx))
}
