package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/invoice/keeper"
	"github.com/stateset/core/x/invoice/types"
	"github.com/stretchr/testify/require"
)

func createNTimedoutInvoice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutInvoice {
	items := make([]types.TimedoutInvoice, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutInvoice(ctx, items[i])
	}
	return items
}

func TestTimedoutInvoiceGet(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNTimedoutInvoice(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutInvoice(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestTimedoutInvoiceRemove(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNTimedoutInvoice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutInvoice(ctx, item.Id)
		_, found := keeper.GetTimedoutInvoice(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutInvoiceGetAll(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNTimedoutInvoice(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllTimedoutInvoice(ctx))
}

func TestTimedoutInvoiceCount(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNTimedoutInvoice(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutInvoiceCount(ctx))
}
