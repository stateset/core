package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) PayInvoice(goCtx context.Context, msg *types.MsgPayInvoice) (*types.MsgPayInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	invoice, found := k.GetInvoice(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	if (invoice.State != "open" || "factored") {
		return nil, sdkerrors.Wrapf(types.ErrWrongInvoiceState, "%v", invoice.State)
	}

	factor, _ := sdk.AccAddressFromBech32(invoice.Factor)
	purchaser, _ := sdk.AccAddressFromBech32(invoice.Purchaser)

	if msg.Creator != invoice.Purchaser {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the purchaser")
	}

	amount, _ := sdk.ParseCoinsNormalized(invoice.Amount)

	err := k.bankKeeper.SendCoins(ctx, purchaser, factor, amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongInvoiceState, "Cannot send payment")
	}

	invoice.State = "paid"

	k.SetLoan(ctx, loan)

	return &types.MsgPayInvoiceResponse{}, nil
}
