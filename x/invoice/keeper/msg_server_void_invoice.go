package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) VoidInvoice(goCtx context.Context, msg *types.MsgVoidInvoice) (*types.MsgVoidInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if invoice.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
	}

	invoice.State = "void"

	k.SetInvoice(ctx, invoice)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtInvoiceVoided,
			sdk.NewAttribute(types.AttributeKeyInvoiceId, strconv.FormatUint(msg.InvoiceId, 10)),
		)
	})

	return &types.MsgVoidInvoiceResponse{}, nil
}
