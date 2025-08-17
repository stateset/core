package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/agreement module sentinel errors
var (
	ErrWrongAgreementState = sdkerrors.Register(ModuleName, 1, "wrong agreement state")
	ErrDeadline            = sdkerrors.Register(ModuleName, 2, "deadline")
	ErrInvalidAddress      = sdkerrors.Register(ModuleName, 3, "invalid address")
)
