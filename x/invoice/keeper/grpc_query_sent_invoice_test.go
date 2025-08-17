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
	"github.com/stateset/core/x/invoice/types"
)

func TestSentInvoiceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSentInvoice(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSentInvoiceRequest
		response *types.QueryGetSentInvoiceResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetSentInvoiceRequest{Id: msgs[0].Id},
			response: &types.QueryGetSentInvoiceResponse{SentInvoice: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetSentInvoiceRequest{Id: msgs[1].Id},
			response: &types.QueryGetSentInvoiceResponse{SentInvoice: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetSentInvoiceRequest{Id: uint64(len(msgs))},
			err:     errorsmod.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SentInvoice(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestSentInvoiceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.InvoiceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSentInvoice(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSentInvoiceRequest {
		return &types.QueryAllSentInvoiceRequest{
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
			resp, err := keeper.SentInvoiceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentInvoice), step)
			require.Subset(t, msgs, resp.SentInvoice)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SentInvoiceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SentInvoice), step)
			require.Subset(t, msgs, resp.SentInvoice)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SentInvoiceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SentInvoiceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
