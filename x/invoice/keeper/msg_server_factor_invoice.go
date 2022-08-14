package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) FactorInvoice(goCtx context.Context, msg *types.MsgFactorInvoice) (*types.MsgFactorInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if invoice.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
	}

	invoice.State = "factored"

	k.SetInvoice(ctx, invoice)

	return &types.MsgFactorInvoiceResponse{}, nil
}
