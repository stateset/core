package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	"github.com/stateset/core/x/agreement/types"
)

func (k msgServer) TerminateAgreement(goCtx context.Context, msg *types.MsgTerminateAgreement) (*types.MsgTerminateAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	agreement, found := k.GetAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if agreement.State != "activated" {
		return nil, errorsmod.Wrapf(types.ErrWrongAgreementState, "%v", agreement.State)
	}

	agreement.State = "terminated"

	return &types.MsgTerminateAgreementResponse{}, nil
}
