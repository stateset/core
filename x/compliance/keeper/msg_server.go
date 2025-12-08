package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/compliance/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl constructs a MsgServer backed by the compliance keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{keeper: k}
}

func (m msgServer) UpsertProfile(goCtx context.Context, msg *types.MsgUpsertProfile) (*types.MsgUpsertProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority for profile upsert")
	}

	profile := msg.Profile
	profile.UpdatedBy = msg.Authority

	if err := profile.ValidateBasic(); err != nil {
		return nil, err
	}

	m.keeper.SetProfile(ctx, profile)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProfileUpserted,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
			sdk.NewAttribute(types.AttributeKeyAddress, profile.Address),
		),
	)

	return &types.MsgUpsertProfileResponse{}, nil
}

func (m msgServer) SetSanction(goCtx context.Context, msg *types.MsgSetSanction) (*types.MsgSetSanctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority for sanction update")
	}

	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidAddress, err.Error())
	}

	profile, found := m.keeper.GetProfile(ctx, addr)
	if !found {
		return nil, errorsmod.Wrap(types.ErrProfileNotFound, "profile not found for address")
	}

	profile.Sanction = msg.Sanction
	if msg.Sanction {
		profile.Metadata = msg.Reason
	}
	profile.UpdatedBy = msg.Authority

	if err := profile.ValidateBasic(); err != nil {
		return nil, err
	}

	m.keeper.SetProfile(ctx, profile)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProfileSanctioned,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
			sdk.NewAttribute(types.AttributeKeyAddress, profile.Address),
			sdk.NewAttribute(types.AttributeKeySanction, strconv.FormatBool(msg.Sanction)),
		),
	)

	return &types.MsgSetSanctionResponse{}, nil
}
