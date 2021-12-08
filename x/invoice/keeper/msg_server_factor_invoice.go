package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/invoice/types"
)

func (k msgServer) FactorInvoice(goCtx context.Context, msg *types.MsgFactorInvoice) (*types.MsgFactorInvoiceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFactorInvoiceResponse{}, nil
}
