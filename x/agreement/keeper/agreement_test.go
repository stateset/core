package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/agreement/keeper"
	"github.com/stateset/core/x/agreement/types"
	"github.com/stretchr/testify/require"
)

func createNAgreement(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Agreement {
	items := make([]types.Agreement, n)
	for i := range items {
		items[i].Id = keeper.AppendAgreement(ctx, items[i])
	}
	return items
}

func TestAgreementGet(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNAgreement(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAgreement(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestAgreementRemove(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNAgreement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAgreement(ctx, item.Id)
		_, found := keeper.GetAgreement(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAgreementGetAll(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNAgreement(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllAgreement(ctx))
}

func TestAgreementCount(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNAgreement(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAgreementCount(ctx))
}
