package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/stateset/core/testutil/keeper"
	"github.com/stateset/core/x/purchaseorder/types"
)

func TestSentPurchaseorderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSentPurchaseorder(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSentPurchaseorderRequest
		response *types.QueryGetSentPurchaseorderResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSentPurchaseorderRequest{Id: msgs[0].Id},
			response: &types.QueryGetSentPurchaseorderResponse{SentPurchaseorder: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetSentPurchaseorderRequest{Id: msgs[1].Id},
			response: &types.QueryGetSentPurchaseorderResponse{SentPurchaseorder: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetSentPurchaseorderRequest{Id: uint64(len(msgs))},
			err:     errorsmod.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SentPurchaseorder(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestSentPurchaseorderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.PurchaseorderKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSentPurchaseorder(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSentPurchaseorderRequest {
		return &types.QueryAllSentPurchaseorderRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SentPurchaseorderAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentPurchaseorder), step)
			require.Subset(t, msgs, resp.SentPurchaseorder)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SentPurchaseorderAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentPurchaseorder), step)
			require.Subset(t, msgs, resp.SentPurchaseorder)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SentPurchaseorderAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SentPurchaseorderAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
