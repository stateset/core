package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/purchaseorder/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PurchaseorderAll(c context.Context, req *types.QueryAllPurchaseorderRequest) (*types.QueryAllPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var purchaseorders []types.Purchaseorder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	purchaseorderStore := prefix.NewStore(store, types.KeyPrefix(types.PurchaseorderKey))

	pageRes, err := query.Paginate(purchaseorderStore, req.Pagination, func(key []byte, value []byte) error {
		var purchaseorder types.Purchaseorder
		if err := k.cdc.Unmarshal(value, &purchaseorder); err != nil {
			return err
		}

		purchaseorders = append(purchaseorders, purchaseorder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPurchaseorderResponse{Purchaseorder: purchaseorders, Pagination: pageRes}, nil
}

func (k Keeper) Purchaseorder(c context.Context, req *types.QueryGetPurchaseorderRequest) (*types.QueryGetPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	purchaseorder, found := k.GetPurchaseorder(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetPurchaseorderResponse{Purchaseorder: purchaseorder}, nil
}
