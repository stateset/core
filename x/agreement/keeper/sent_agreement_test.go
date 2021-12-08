package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/agreement/keeper"
	"github.com/stateset/core/x/agreement/types"
	"github.com/stretchr/testify/require"
)

func createNSentAgreement(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentAgreement {
	items := make([]types.SentAgreement, n)
	for i := range items {
		items[i].Id = keeper.AppendSentAgreement(ctx, items[i])
	}
	return items
}

func TestSentAgreementGet(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNSentAgreement(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentAgreement(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestSentAgreementRemove(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNSentAgreement(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentAgreement(ctx, item.Id)
		_, found := keeper.GetSentAgreement(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentAgreementGetAll(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNSentAgreement(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllSentAgreement(ctx))
}

func TestSentAgreementCount(t *testing.T) {
	keeper, ctx := keepertest.AgreementKeeper(t)
	items := createNSentAgreement(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentAgreementCount(ctx))
}
