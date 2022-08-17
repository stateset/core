package keeper_test

import (
	"testing"

	testkeeper "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/refund/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RefundKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
