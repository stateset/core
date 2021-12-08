package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/agreement/keeper"
	"github.com/stateset/core/x/agreement/types"
	"github.com/stretchr/testify/require"
)

func createNTimedoutAgreement(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutAgreement {
	items := make([]types.TimedoutAgreement, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutAgreement(ctx, items[i])
	}
	return items
}

func TestTimedoutAgreementGet(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNTimedoutAgreement(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutAgreement(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestTimedoutAgreementRemove(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNTimedoutAgreement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutAgreement(ctx, item.Id)
		_, found := keeper.GetTimedoutAgreement(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutAgreementGetAll(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNTimedoutAgreement(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllTimedoutAgreement(ctx))
}

func TestTimedoutAgreementCount(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNTimedoutAgreement(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutAgreementCount(ctx))
}
