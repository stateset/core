package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) VoidInvoice(goCtx context.Context, msg *types.MsgVoidInvoice) (*types.MsgVoidInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if invoice.State != "requested" {
		return nil, errorsmod.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
	}

	invoice.State = "void"

	k.SetInvoice(ctx, invoice)

	return &types.MsgVoidInvoiceResponse{}, nil
}
