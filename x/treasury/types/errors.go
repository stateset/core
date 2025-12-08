package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrUnauthorized   = errorsmod.Register(ModuleName, 1, "unauthorized")
	ErrInvalidReserve = errorsmod.Register(ModuleName, 2, "invalid reserve snapshot")
)
