package keeper

import (
	"context"

	"github.com/stateset/core/x/stst/types"
)

type queryServer struct {
	Keeper
}

// NewQueryServerImpl returns an implementation of the QueryServer interface
// for the provided Keeper.
func NewQueryServerImpl(keeper Keeper) types.QueryServer {
	return &queryServer{Keeper: keeper}
}

var _ types.QueryServer = queryServer{}

func (k queryServer) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	params, err := k.Keeper.GetParams(goCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k queryServer) StakingState(goCtx context.Context, req *types.QueryStakingStateRequest) (*types.QueryStakingStateResponse, error) {
	stakingState, err := k.Keeper.GetStakingState(goCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryStakingStateResponse{StakingState: stakingState}, nil
}

func (k queryServer) FeeBurnState(goCtx context.Context, req *types.QueryFeeBurnStateRequest) (*types.QueryFeeBurnStateResponse, error) {
	feeBurnState, err := k.Keeper.GetFeeBurnState(goCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryFeeBurnStateResponse{FeeBurnState: feeBurnState}, nil
}

func (k queryServer) StakerInfo(goCtx context.Context, req *types.QueryStakerInfoRequest) (*types.QueryStakerInfoResponse, error) {
	// TODO: Implement staker info query logic
	return &types.QueryStakerInfoResponse{}, nil
}

func (k queryServer) Proposals(goCtx context.Context, req *types.QueryProposalsRequest) (*types.QueryProposalsResponse, error) {
	// TODO: Implement proposals query logic
	return &types.QueryProposalsResponse{}, nil
}

func (k queryServer) Proposal(goCtx context.Context, req *types.QueryProposalRequest) (*types.QueryProposalResponse, error) {
	// TODO: Implement proposal query logic
	return &types.QueryProposalResponse{}, nil
}

func (k queryServer) VestingSchedules(goCtx context.Context, req *types.QueryVestingSchedulesRequest) (*types.QueryVestingSchedulesResponse, error) {
	vestingSchedules, err := k.Keeper.GetAllVestingSchedules(goCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryVestingSchedulesResponse{VestingSchedules: vestingSchedules}, nil
}

func (k queryServer) VestingSchedule(goCtx context.Context, req *types.QueryVestingScheduleRequest) (*types.QueryVestingScheduleResponse, error) {
	vestingSchedule, found, err := k.Keeper.GetVestingSchedule(goCtx, req.Category)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, types.ErrVestingScheduleNotFound
	}

	return &types.QueryVestingScheduleResponse{VestingSchedule: vestingSchedule}, nil
}

func (k queryServer) VestingInfo(goCtx context.Context, req *types.QueryVestingInfoRequest) (*types.QueryVestingInfoResponse, error) {
	// TODO: Implement vesting info query logic
	return &types.QueryVestingInfoResponse{}, nil
}

func (k queryServer) TokenSupply(goCtx context.Context, req *types.QueryTokenSupplyRequest) (*types.QueryTokenSupplyResponse, error) {
	// TODO: Implement token supply query logic
	return &types.QueryTokenSupplyResponse{}, nil
}