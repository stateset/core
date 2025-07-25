package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/stateset/core/x/epochs/types"
)

var _ types.QueryServer = Keeper{}

// EpochInfos provide running epochInfos
func (k Keeper) EpochInfos(c context.Context, req *types.QueryEpochsInfoRequest) (*types.QueryEpochsInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	var epochs []types.EpochInfo
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEpoch)

	pageRes, err := query.Paginate(store, req.Pagination, func(_, value []byte) error {
		var epoch types.EpochInfo
		if err := k.cdc.Unmarshal(value, &epoch); err != nil {
			return err
		}
		epochs = append(epochs, epoch)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryEpochsInfoResponse{
		Epochs:     epochs,
		Pagination: pageRes,
	}, nil
}

// CurrentEpoch provides current epoch of specified identifier
func (k Keeper) CurrentEpoch(c context.Context, req *types.QueryCurrentEpochRequest) (*types.QueryCurrentEpochResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	info, _ := k.GetEpochInfo(ctx, req.Identifier)

	return &types.QueryCurrentEpochResponse{
		CurrentEpoch: info.CurrentEpoch,
	}, nil
}

// CurrentEpoch provides current epoch of specified identifier
func (k Keeper) EpochInfo(c context.Context, req *types.QueryEpochInfoRequest) (*types.QueryEpochInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	info, found := k.GetEpochInfo(ctx, req.Identifier)
	if !found {
		return nil, status.Error(codes.NotFound, "epoch info not found")
	}

	return &types.QueryEpochInfoResponse{
		Epoch: info,
	}, nil
}
