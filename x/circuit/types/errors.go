package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrCircuitOpen         = errorsmod.Register(ModuleName, 1, "circuit breaker is open")
	ErrGlobalPause         = errorsmod.Register(ModuleName, 2, "system is globally paused")
	ErrRateLimitExceeded   = errorsmod.Register(ModuleName, 3, "rate limit exceeded")
	ErrUnauthorized        = errorsmod.Register(ModuleName, 4, "unauthorized")
	ErrInvalidParams       = errorsmod.Register(ModuleName, 5, "invalid parameters")
	ErrModuleNotFound      = errorsmod.Register(ModuleName, 6, "module circuit not found")
	ErrInvalidDuration     = errorsmod.Register(ModuleName, 7, "invalid duration")
	ErrAlreadyPaused       = errorsmod.Register(ModuleName, 8, "system already paused")
	ErrNotPaused           = errorsmod.Register(ModuleName, 9, "system not paused")
	ErrMessageDisabled     = errorsmod.Register(ModuleName, 10, "message type is disabled")
	ErrOracleDeviation     = errorsmod.Register(ModuleName, 11, "oracle price deviation exceeds threshold")
	ErrOracleStaleness     = errorsmod.Register(ModuleName, 12, "oracle price is stale")
	ErrLiquidationSurge    = errorsmod.Register(ModuleName, 13, "liquidation surge protection triggered")
	ErrOracleUpdateTooFast = errorsmod.Register(ModuleName, 14, "oracle update too frequent")
)
