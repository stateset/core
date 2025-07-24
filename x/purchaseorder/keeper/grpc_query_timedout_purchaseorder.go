package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/purchaseorder/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TimedoutPurchaseorderAll(c context.Context, req *types.QueryAllTimedoutPurchaseorderRequest) (*types.QueryAllTimedoutPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timedoutPurchaseorders []types.TimedoutPurchaseorder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	timedoutPurchaseorderStore := prefix.NewStore(store, types.KeyPrefix(types.TimedoutPurchaseorderKey))

	pageRes, err := query.Paginate(timedoutPurchaseorderStore, req.Pagination, func(key []byte, value []byte) error {
		var timedoutPurchaseorder types.TimedoutPurchaseorder
		if err := k.cdc.Unmarshal(value, &timedoutPurchaseorder); err != nil {
			return err
		}

		timedoutPurchaseorders = append(timedoutPurchaseorders, timedoutPurchaseorder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimedoutPurchaseorderResponse{TimedoutPurchaseorder: timedoutPurchaseorders, Pagination: pageRes}, nil
}

func (k Keeper) TimedoutPurchaseorder(c context.Context, req *types.QueryGetTimedoutPurchaseorderRequest) (*types.QueryGetTimedoutPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	timedoutPurchaseorder, found := k.GetTimedoutPurchaseorder(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimedoutPurchaseorderResponse{TimedoutPurchaseorder: timedoutPurchaseorder}, nil
}
