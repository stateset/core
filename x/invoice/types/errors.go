package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/invoice module sentinel errors
var (
	ErrWrongInvoiceState = sdkerrors.Register(ModuleName, 1, "wrong invoice state")
	ErrDeadline                = sdkerrors.Register(ModuleName, 2, "deadline")
)
