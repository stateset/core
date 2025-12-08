package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/oracle/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the oracle MsgServer backed by keeper state.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return msgServer{keeper: k}
}

func (m msgServer) UpdatePrice(goCtx context.Context, msg *types.MsgUpdatePrice) (*types.MsgUpdatePriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "authority mismatch")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	price := types.Price{
		Denom:       msg.Denom,
		Amount:      msg.Price,
		LastUpdater: msg.Authority,
		LastHeight:  ctx.BlockHeight(),
	}

	if err := price.ValidateBasic(); err != nil {
		return nil, err
	}

	m.keeper.SetPrice(goCtx, price)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypePriceUpdated,
			sdk.NewAttribute(types.AttributeKeyDenom, price.Denom),
			sdk.NewAttribute(types.AttributeKeyPrice, price.Amount.String()),
			sdk.NewAttribute(types.AttributeKeySource, msg.Authority),
		),
	)

	return &types.MsgUpdatePriceResponse{}, nil
}
