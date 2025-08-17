package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/loan module sentinel errors
var (
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 1, "wrong loan state")
	ErrDeadline       = sdkerrors.Register(ModuleName, 2, "deadline")
	ErrInvalidAddress = sdkerrors.Register(ModuleName, 3, "invalid address")
)
