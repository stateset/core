package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/proof module sentinel errors
var (
	ErrSample = errorsmod.Register(ModuleName, 1100, "sample error")
)
