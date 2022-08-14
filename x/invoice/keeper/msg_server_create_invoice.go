package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) CreateInvoice(goCtx context.Context, msg *types.MsgCreateInvoice) (*types.MsgCreateInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateInvoiceResponse{}, nil
}
