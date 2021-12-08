package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/agreement module sentinel errors
var (
	ErrWrongAgreementState = sdkerrors.Register(ModuleName, 1, "wrong agreement state")
	ErrDeadline            = sdkerrors.Register(ModuleName, 2, "deadline")
)
