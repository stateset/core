package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/compliance/types"
)

var _ types.QueryServer = queryServer{}

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the compliance QueryServer interface
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

// Profile returns a compliance profile by address
func (q queryServer) Profile(req *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	return nil, types.ErrProfileNotFound
}

// Profiles returns all profiles with pagination
func (q queryServer) Profiles(req *types.QueryProfilesRequest) (*types.QueryProfilesResponse, error) {
	return nil, types.ErrProfileNotFound
}

// ProfilesByStatus returns profiles filtered by status
func (q queryServer) ProfilesByStatus(req *types.QueryProfilesByStatusRequest) (*types.QueryProfilesByStatusResponse, error) {
	return nil, types.ErrProfileNotFound
}

// Context-aware query implementations

// ProfileWithContext returns a compliance profile by address
func (q queryServer) ProfileWithContext(goCtx context.Context, req *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, types.ErrInvalidAddress
	}

	profile, found := q.Keeper.GetProfile(ctx, addr)
	if !found {
		return nil, types.ErrProfileNotFound
	}

	return &types.QueryProfileResponse{
		Profile: profile,
	}, nil
}

// ProfilesWithContext returns all profiles with pagination
func (q queryServer) ProfilesWithContext(goCtx context.Context, req *types.QueryProfilesRequest) (*types.QueryProfilesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	maxLimit := uint64(100)
	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var profiles []types.Profile
	var total uint64

	q.Keeper.IterateProfiles(ctx, func(p types.Profile) bool {
		if total >= offset && uint64(len(profiles)) < limit {
			profiles = append(profiles, p)
		}
		total++
		return false
	})

	return &types.QueryProfilesResponse{
		Profiles: profiles,
		Total:    total,
	}, nil
}

// ProfilesByStatusWithContext returns profiles filtered by status
func (q queryServer) ProfilesByStatusWithContext(goCtx context.Context, req *types.QueryProfilesByStatusRequest) (*types.QueryProfilesByStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	maxLimit := uint64(100)
	limit := req.Limit
	if limit == 0 || limit > maxLimit {
		limit = maxLimit
	}
	offset := req.Offset

	var profiles []types.Profile
	var matched uint64

	q.Keeper.IterateProfiles(ctx, func(p types.Profile) bool {
		if p.Status == req.Status {
			if matched >= offset && uint64(len(profiles)) < limit {
				profiles = append(profiles, p)
			}
			matched++
		}
		return false
	})

	return &types.QueryProfilesByStatusResponse{
		Profiles: profiles,
		Total:    matched,
	}, nil
}
