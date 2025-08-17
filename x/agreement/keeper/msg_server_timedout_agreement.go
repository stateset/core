package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
	"github.com/stateset/core/x/agreement/types"
)

func (k msgServer) CreateTimedoutAgreement(goCtx context.Context, msg *types.MsgCreateTimedoutAgreement) (*types.MsgCreateTimedoutAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutAgreement = types.TimedoutAgreement{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendTimedoutAgreement(
		ctx,
		timedoutAgreement,
	)

	return &types.MsgCreateTimedoutAgreementResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateTimedoutAgreement(goCtx context.Context, msg *types.MsgUpdateTimedoutAgreement) (*types.MsgUpdateTimedoutAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var timedoutAgreement = types.TimedoutAgreement{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetTimedoutAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetTimedoutAgreement(ctx, timedoutAgreement)

	return &types.MsgUpdateTimedoutAgreementResponse{}, nil
}

func (k msgServer) DeleteTimedoutAgreement(goCtx context.Context, msg *types.MsgDeleteTimedoutAgreement) (*types.MsgDeleteTimedoutAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetTimedoutAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveTimedoutAgreement(ctx, msg.Id)

	return &types.MsgDeleteTimedoutAgreementResponse{}, nil
}
