package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stateset/core/x/agreement/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AgreementAll(c context.Context, req *types.QueryAllAgreementRequest) (*types.QueryAllAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var agreements []types.Agreement
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	agreementStore := prefix.NewStore(store, types.KeyPrefix(types.AgreementKey))

	pageRes, err := query.Paginate(agreementStore, req.Pagination, func(key []byte, value []byte) error {
		var agreement types.Agreement
		if err := k.cdc.Unmarshal(value, &agreement); err != nil {
			return err
		}

		agreements = append(agreements, agreement)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAgreementResponse{Agreement: agreements, Pagination: pageRes}, nil
}

func (k Keeper) Agreement(c context.Context, req *types.QueryGetAgreementRequest) (*types.QueryGetAgreementResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	agreement, found := k.GetAgreement(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAgreementResponse{Agreement: agreement}, nil
}
