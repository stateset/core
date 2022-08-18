package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) VoidInvoice(goCtx context.Context, msg *types.MsgVoidInvoice) (*types.MsgVoidInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if (invoice.State != "open" || "requested") {
		return nil, sdkerrors.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
	}

	invoice.State = "void"

	k.SetInvoice(ctx, invoice)
	return &types.MsgVoidInvoiceResponse{}, nil
}
