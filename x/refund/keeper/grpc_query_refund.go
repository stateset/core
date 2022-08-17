package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/refund/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RefundAll(c context.Context, req *types.QueryAllRefundRequest) (*types.QueryAllRefundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var refunds []types.Refund
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	refundStore := prefix.NewStore(store, types.KeyPrefix(types.RefundKey))

	pageRes, err := query.Paginate(refundStore, req.Pagination, func(key []byte, value []byte) error {
		var refund types.Refund
		if err := k.cdc.Unmarshal(value, &refund); err != nil {
			return err
		}

		refunds = append(refunds, refund)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRefundResponse{Refund: refunds, Pagination: pageRes}, nil
}

func (k Keeper) Refund(c context.Context, req *types.QueryGetRefundRequest) (*types.QueryGetRefundResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	refund, found := k.GetRefund(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRefundResponse{Refund: refund}, nil
}
