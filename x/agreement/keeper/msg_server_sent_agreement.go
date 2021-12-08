package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/agreement/types"
)

func (k msgServer) CreateSentAgreement(goCtx context.Context, msg *types.MsgCreateSentAgreement) (*types.MsgCreateSentAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentAgreement = types.SentAgreement{
		Creator: msg.Creator,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	id := k.AppendSentAgreement(
		ctx,
		sentAgreement,
	)

	return &types.MsgCreateSentAgreementResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateSentAgreement(goCtx context.Context, msg *types.MsgUpdateSentAgreement) (*types.MsgUpdateSentAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var sentAgreement = types.SentAgreement{
		Creator: msg.Creator,
		Id:      msg.Id,
		Did:     msg.Did,
		Chain:   msg.Chain,
	}

	// Checks that the element exists
	val, found := k.GetSentAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetSentAgreement(ctx, sentAgreement)

	return &types.MsgUpdateSentAgreementResponse{}, nil
}

func (k msgServer) DeleteSentAgreement(goCtx context.Context, msg *types.MsgDeleteSentAgreement) (*types.MsgDeleteSentAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSentAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSentAgreement(ctx, msg.Id)

	return &types.MsgDeleteSentAgreementResponse{}, nil
}
