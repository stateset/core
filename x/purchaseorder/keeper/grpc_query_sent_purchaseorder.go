package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/purchaseorder/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SentPurchaseorderAll(c context.Context, req *types.QueryAllSentPurchaseorderRequest) (*types.QueryAllSentPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentPurchaseorders []types.SentPurchaseorder
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sentPurchaseorderStore := prefix.NewStore(store, types.KeyPrefix(types.SentPurchaseorderKey))

	pageRes, err := query.Paginate(sentPurchaseorderStore, req.Pagination, func(key []byte, value []byte) error {
		var sentPurchaseorder types.SentPurchaseorder
		if err := k.cdc.Unmarshal(value, &sentPurchaseorder); err != nil {
			return err
		}

		sentPurchaseorders = append(sentPurchaseorders, sentPurchaseorder)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentPurchaseorderResponse{SentPurchaseorder: sentPurchaseorders, Pagination: pageRes}, nil
}

func (k Keeper) SentPurchaseorder(c context.Context, req *types.QueryGetSentPurchaseorderRequest) (*types.QueryGetSentPurchaseorderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	sentPurchaseorder, found := k.GetSentPurchaseorder(ctx, req.Id)
	if !found {
		return nil, errorsmod.ErrKeyNotFound
	}

	return &types.QueryGetSentPurchaseorderResponse{SentPurchaseorder: sentPurchaseorder}, nil
}
