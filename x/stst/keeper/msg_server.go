package keeper

import (
	"context"

	"github.com/stateset/core/x/stst/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) StakeTokens(goCtx context.Context, msg *types.MsgStakeTokens) (*types.MsgStakeTokensResponse, error) {
	err := k.Keeper.StakeTokens(goCtx, msg.Staker, msg.ValidatorAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgStakeTokensResponse{}, nil
}

func (k msgServer) UnstakeTokens(goCtx context.Context, msg *types.MsgUnstakeTokens) (*types.MsgUnstakeTokensResponse, error) {
	completionTime, err := k.Keeper.UnstakeTokens(goCtx, msg.Staker, msg.ValidatorAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgUnstakeTokensResponse{
		CompletionTime: completionTime,
	}, nil
}

func (k msgServer) ClaimStakingRewards(goCtx context.Context, msg *types.MsgClaimStakingRewards) (*types.MsgClaimStakingRewardsResponse, error) {
	// TODO: Implement staking rewards claiming logic
	return &types.MsgClaimStakingRewardsResponse{}, nil
}

func (k msgServer) SubmitProposal(goCtx context.Context, msg *types.MsgSubmitProposal) (*types.MsgSubmitProposalResponse, error) {
	// TODO: Implement governance proposal submission logic
	return &types.MsgSubmitProposalResponse{}, nil
}

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	// TODO: Implement governance voting logic
	return &types.MsgVoteResponse{}, nil
}

func (k msgServer) ClaimVestedTokens(goCtx context.Context, msg *types.MsgClaimVestedTokens) (*types.MsgClaimVestedTokensResponse, error) {
	// TODO: Implement vested token claiming logic
	return &types.MsgClaimVestedTokensResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != msg.Authority {
		return nil, types.ErrUnauthorized
	}

	if err := k.SetParams(goCtx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}