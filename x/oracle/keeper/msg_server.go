package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"

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
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	if !m.keeper.IsAuthorizedProvider(goCtx, msg.Authority) {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "unauthorized provider")
	}

	if err := m.keeper.SetPriceWithValidation(goCtx, msg.Authority, msg.Denom, msg.Price); err != nil {
		return nil, err
	}

	return &types.MsgUpdatePriceResponse{}, nil
}
