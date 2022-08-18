package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	factor, _ := sdk.AccAddressFromBech32(msg.Creator)
	seller, _ := sdk.AccAddressFromBech32(invoice.Seller)
	amount, _ := sdk.ParseCoinsNormalized(invoice.Amount)

	k.bankKeeper.SendCoins(ctx, factor, seller, amount)


	invoice.Factor = msg.Creator
	invoice.State = "factored"

	k.SetInvoice(ctx, invoice)

	return &types.MsgFactorInvoiceResponse{}, nil
}
