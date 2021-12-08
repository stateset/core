package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/purchaseorder module sentinel errors
var (
	ErrWrongPurchaseOrderState = sdkerrors.Register(ModuleName, 1, "wrong purchaseorder state")
	ErrDeadline                = sdkerrors.Register(ModuleName, 2, "deadline")
)
