package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/stateset/core/x/feemarket/types"
)

var _ types.MsgServer = msgServer{}

// msgServer is a wrapper of Keeper that implements MsgServer interface.
type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the feemarket MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// UpdateParams updates the feemarket module parameters.
// This can only be executed by the governance module authority.
func (ms msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if msg == nil {
		return nil, types.ErrInvalidRequest
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate authority
	if ms.GetAuthority() != msg.Authority {
		return nil, errorsmod.Wrapf(
			govtypes.ErrInvalidSigner,
			"invalid authority; expected %s, got %s",
			ms.GetAuthority(),
			msg.Authority,
		)
	}

	// Validate params
	if err := msg.Params.Validate(); err != nil {
		return nil, errorsmod.Wrap(types.ErrInvalidParams, err.Error())
	}

	// Store updated params
	ms.SetParams(ctx, msg.Params)

	// Emit event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateParams,
			sdk.NewAttribute(types.AttributeKeyAuthority, msg.Authority),
			sdk.NewAttribute(types.AttributeKeyEnabled, strconv.FormatBool(msg.Params.Enabled)),
			sdk.NewAttribute(types.AttributeKeyMinBaseFee, msg.Params.MinBaseFee.String()),
			sdk.NewAttribute(types.AttributeKeyMaxBaseFee, msg.Params.MaxBaseFee.String()),
		),
	)

	return &types.MsgUpdateParamsResponse{}, nil
}
