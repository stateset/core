package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/invoice/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TimedoutInvoiceAll(c context.Context, req *types.QueryAllTimedoutInvoiceRequest) (*types.QueryAllTimedoutInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timedoutInvoices []types.TimedoutInvoice
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	timedoutInvoiceStore := prefix.NewStore(store, types.KeyPrefix(types.TimedoutInvoiceKey))

	pageRes, err := query.Paginate(timedoutInvoiceStore, req.Pagination, func(key []byte, value []byte) error {
		var timedoutInvoice types.TimedoutInvoice
		if err := k.cdc.Unmarshal(value, &timedoutInvoice); err != nil {
			return err
		}

		timedoutInvoices = append(timedoutInvoices, timedoutInvoice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimedoutInvoiceResponse{TimedoutInvoice: timedoutInvoices, Pagination: pageRes}, nil
}

func (k Keeper) TimedoutInvoice(c context.Context, req *types.QueryGetTimedoutInvoiceRequest) (*types.QueryGetTimedoutInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	timedoutInvoice, found := k.GetTimedoutInvoice(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimedoutInvoiceResponse{TimedoutInvoice: timedoutInvoice}, nil
}
