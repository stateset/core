package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsmod "cosmossdk.io/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) PayInvoice(goCtx context.Context, msg *types.MsgPayInvoice) (*types.MsgPayInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(errorsmod.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	factor, _ := sdk.AccAddressFromBech32(invoice.Factor)
	seller, _ := sdk.AccAddressFromBech32(invoice.Seller)
	purchaser, _ := sdk.AccAddressFromBech32(invoice.Purchaser)

	if msg.Creator != invoice.Purchaser {
		return nil, errorsmod.Wrap(errorsmod.ErrUnauthorized, "Cannot repay: not the purchaser")
	}

	amount, _ := sdk.ParseCoinsNormalized(invoice.Amount)

	if invoice.Factor != "" {

		if invoice.State != "factored" {
			return nil, errorsmod.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
		}

		k.bankKeeper.SendCoins(ctx, purchaser, factor, amount)

	} else {

		if invoice.State != "requested" {
			return nil, errorsmod.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
		}

		k.bankKeeper.SendCoins(ctx, purchaser, seller, amount)

	}

	invoice.State = "paid"

	k.SetInvoice(ctx, invoice)

	return &types.MsgPayInvoiceResponse{}, nil
}
