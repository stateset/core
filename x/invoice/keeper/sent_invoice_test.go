package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/invoice/keeper"
	"github.com/stateset/core/x/invoice/types"
	"github.com/stretchr/testify/require"
)

func createNSentInvoice(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentInvoice {
	items := make([]types.SentInvoice, n)
	for i := range items {
		items[i].Id = keeper.AppendSentInvoice(ctx, items[i])
	}
	return items
}

func TestSentInvoiceGet(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNSentInvoice(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentInvoice(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestSentInvoiceRemove(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNSentInvoice(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentInvoice(ctx, item.Id)
		_, found := keeper.GetSentInvoice(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentInvoiceGetAll(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNSentInvoice(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllSentInvoice(ctx))
}

func TestSentInvoiceCount(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	items := createNSentInvoice(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentInvoiceCount(ctx))
}
