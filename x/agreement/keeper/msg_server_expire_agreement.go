package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/agreement/types"
)

func (k msgServer) ExpireAgreement(goCtx context.Context, msg *types.MsgExpireAgreement) (*types.MsgExpireAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	agreement, found := k.GetAgreement(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if agreement.State != "activated" {
		return nil, sdkerrors.Wrapf(types.ErrWrongAgreementState, "%v", agreement.State)
	}

	agreement.State = "expired"

	k.SetAgreement(ctx, agreement)
	return &types.MsgExpireAgreementResponse{}, nil
}
