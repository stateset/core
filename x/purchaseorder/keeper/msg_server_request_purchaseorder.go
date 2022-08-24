package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/core/x/purchaseorder/types"
)

func (k msgServer) RequestPurchaseorder(goCtx context.Context, msg *types.MsgRequestPurchaseorder) (*types.MsgRequestPurchaseorderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	var purchaseorder = types.Purchaseorder{
		Did:    msg.Did,
		Uri:    msg.Uri,
		Amount: msg.Amount,
		State:  "requested",
		Purchaser: msg.Creator,
		Seller: msg.Seller,
	}

	// Create NFT for the Purchase Order

	k.AppendPurchaseorder(
		ctx,
		purchaseorder,
	)
	return &types.MsgRequestPurchaseorderResponse{}, nil
}
