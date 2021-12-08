package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/invoice/keeper"
	"github.com/stateset/core/x/invoice/types"
	"github.com/stretchr/testify/require"
)

func createNInvoice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Invoice {
	items := make([]types.Invoice, n)
	for i := range items {
		items[i].Id = keeper.AppendInvoice(ctx, items[i])
	}
	return items
}

func TestInvoiceGet(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetInvoice(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestInvoiceRemove(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveInvoice(ctx, item.Id)
		_, found := keeper.GetInvoice(ctx, item.Id)
		require.False(t, found)
	}
}

func TestInvoiceGetAll(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllInvoice(ctx))
}

func TestInvoiceCount(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNInvoice(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetInvoiceCount(ctx))
}
