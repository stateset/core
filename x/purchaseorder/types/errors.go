package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/purchaseorder module sentinel errors
var (
	ErrWrongPurchaseOrderState = errorsmod.Register(ModuleName, 1, "wrong purchaseorder state")
	ErrDeadline                = errorsmod.Register(ModuleName, 2, "deadline")
	ErrInvalidAddress          = errorsmod.Register(ModuleName, 3, "invalid address")
)
