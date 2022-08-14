package keeper

import (
	"context"
	"fmt"

	
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stateset/core/x/invoice/types"

)

func (k msgServer) CreateInvoice(goCtx context.Context, msg *types.MsgCreateInvoice) (*types.MsgCreateInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var invoice = types.Invoice {
		Did:    msg.Did,
		Amount: msg.Amount,
		State:  "requested",
	}

	k.AppendInvoice(
		ctx,
		invoice,
	)

	return &types.MsgCreateInvoiceResponse{}, nil
}
