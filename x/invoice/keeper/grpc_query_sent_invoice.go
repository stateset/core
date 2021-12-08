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

func (k Keeper) SentInvoiceAll(c context.Context, req *types.QueryAllSentInvoiceRequest) (*types.QueryAllSentInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentInvoices []types.SentInvoice
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sentInvoiceStore := prefix.NewStore(store, types.KeyPrefix(types.SentInvoiceKey))

	pageRes, err := query.Paginate(sentInvoiceStore, req.Pagination, func(key []byte, value []byte) error {
		var sentInvoice types.SentInvoice
		if err := k.cdc.Unmarshal(value, &sentInvoice); err != nil {
			return err
		}

		sentInvoices = append(sentInvoices, sentInvoice)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentInvoiceResponse{SentInvoice: sentInvoices, Pagination: pageRes}, nil
}

func (k Keeper) SentInvoice(c context.Context, req *types.QueryGetSentInvoiceRequest) (*types.QueryGetSentInvoiceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	sentInvoice, found := k.GetSentInvoice(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentInvoiceResponse{SentInvoice: sentInvoice}, nil
}
