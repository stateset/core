package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/agreement/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SentAgreementAll(c context.Context, req *types.QueryAllSentAgreementRequest) (*types.QueryAllSentAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentAgreements []types.SentAgreement
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sentAgreementStore := prefix.NewStore(store, types.KeyPrefix(types.SentAgreementKey))

	pageRes, err := query.Paginate(sentAgreementStore, req.Pagination, func(key []byte, value []byte) error {
		var sentAgreement types.SentAgreement
		if err := k.cdc.Unmarshal(value, &sentAgreement); err != nil {
			return err
		}

		sentAgreements = append(sentAgreements, sentAgreement)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentAgreementResponse{SentAgreement: sentAgreements, Pagination: pageRes}, nil
}

func (k Keeper) SentAgreement(c context.Context, req *types.QueryGetSentAgreementRequest) (*types.QueryGetSentAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	sentAgreement, found := k.GetSentAgreement(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentAgreementResponse{SentAgreement: sentAgreement}, nil
}
