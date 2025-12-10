package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidParams   = errorsmod.Register(ModuleName, 1, "invalid fee market params")
	ErrInvalidBaseFee  = errorsmod.Register(ModuleName, 2, "invalid base fee")
	ErrInvalidRequest  = errorsmod.Register(ModuleName, 3, "invalid request")
	ErrInvalidGasLimit = errorsmod.Register(ModuleName, 4, "invalid gas limit")
	ErrInsufficientFee = errorsmod.Register(ModuleName, 5, "insufficient fee")
)
