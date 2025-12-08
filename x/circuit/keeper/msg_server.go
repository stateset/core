package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/circuit/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// PauseSystem handles MsgPauseSystem
func (m msgServer) PauseSystem(ctx context.Context, msg *types.MsgPauseSystem) (*types.MsgPauseSystemResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if !m.keeper.IsAuthorized(sdkCtx, msg.Authority) {
		return nil, types.ErrUnauthorized
	}

	if err := m.keeper.PauseSystem(sdkCtx, msg.Authority, msg.Reason, msg.DurationSeconds); err != nil {
		return nil, err
	}

	return &types.MsgPauseSystemResponse{}, nil
}

// ResumeSystem handles MsgResumeSystem
func (m msgServer) ResumeSystem(ctx context.Context, msg *types.MsgResumeSystem) (*types.MsgResumeSystemResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if !m.keeper.IsAuthorized(sdkCtx, msg.Authority) {
		return nil, types.ErrUnauthorized
	}

	if err := m.keeper.ResumeSystem(sdkCtx, msg.Authority); err != nil {
		return nil, err
	}

	return &types.MsgResumeSystemResponse{}, nil
}

// TripCircuit handles MsgTripCircuit
func (m msgServer) TripCircuit(ctx context.Context, msg *types.MsgTripCircuit) (*types.MsgTripCircuitResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if !m.keeper.IsAuthorized(sdkCtx, msg.Authority) {
		return nil, types.ErrUnauthorized
	}

	if err := m.keeper.TripCircuit(sdkCtx, msg.ModuleName, msg.Reason, msg.Authority, msg.DisableMessages); err != nil {
		return nil, err
	}

	return &types.MsgTripCircuitResponse{}, nil
}

// ResetCircuit handles MsgResetCircuit
func (m msgServer) ResetCircuit(ctx context.Context, msg *types.MsgResetCircuit) (*types.MsgResetCircuitResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	if !m.keeper.IsAuthorized(sdkCtx, msg.Authority) {
		return nil, types.ErrUnauthorized
	}

	if err := m.keeper.ResetCircuit(sdkCtx, msg.ModuleName, msg.Authority); err != nil {
		return nil, err
	}

	return &types.MsgResetCircuitResponse{}, nil
}

// UpdateParams handles MsgUpdateParams
func (m msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Only governance can update params
	if msg.Authority != m.keeper.GetAuthority() {
		return nil, types.ErrUnauthorized
	}

	if err := m.keeper.SetParams(sdkCtx, msg.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
