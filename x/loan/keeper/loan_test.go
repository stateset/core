package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/loan/keeper"
	"github.com/stateset/core/x/loan/types"
	"github.com/stretchr/testify/require"
)

func createNLoan(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Loan {
	items := make([]types.Loan, n)
	for i := range items {
		items[i].Id = keeper.AppendLoan(ctx, items[i])
	}
	return items
}

func TestLoanGet(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLoan(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestLoanRemove(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLoan(ctx, item.Id)
		_, found := keeper.GetLoan(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLoanGetAll(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllLoan(ctx))
}

func TestLoanCount(t *testing.T) {
	keeper, ctx := keepertest.LoanKeeper(t)
	items := createNLoan(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLoanCount(ctx))
}
