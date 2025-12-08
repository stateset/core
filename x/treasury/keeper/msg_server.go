package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/treasury/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns a MsgServer backed by the treasury keeper.
func NewMsgServerImpl(k Keeper) types.MsgServer {
	return msgServer{keeper: k}
}

func (m msgServer) RecordReserve(goCtx context.Context, msg *types.MsgRecordReserve) (*types.MsgRecordReserveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != m.keeper.GetAuthority() {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "invalid authority")
	}

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	id := m.keeper.RecordSnapshot(goCtx, msg.Snapshot)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSnapshotRecorded,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
			sdk.NewAttribute(types.AttributeKeySnapshotID, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.Snapshot.Reporter),
			sdk.NewAttribute(types.AttributeKeyTotal, msg.Snapshot.TotalSupply.String()),
		),
	)

	return &types.MsgRecordReserveResponse{SnapshotID: id}, nil
}
