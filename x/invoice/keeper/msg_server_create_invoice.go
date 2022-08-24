package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) CreateInvoice(goCtx context.Context, msg *types.MsgCreateInvoice) (*types.MsgCreateInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var invoice = types.Invoice{
		Did:    msg.Did,
		Amount: msg.Amount,
		State:  "requested",
		Seller: msg.Creator,
		Purchaser: msg.Purchaser,
	}

	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	
	k.AppendInvoice(
		ctx,
		invoice,
	)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.TypeEvtInvoiceCreated,
			sdk.NewAttribute(types.AttributeKeyInvoiceId, strconv.FormatUint(invoiceId, 10)),
		)
	})

	return &types.MsgCreateInvoiceResponse{}, nil
}
