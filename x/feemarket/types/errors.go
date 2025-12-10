package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidParams  = errorsmod.Register(ModuleName, 1, "invalid fee market params")
	ErrInvalidBaseFee = errorsmod.Register(ModuleName, 2, "invalid base fee")
)
