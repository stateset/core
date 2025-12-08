package types

import (
	errorsmod "cosmossdk.io/errors"
)

const (
	moduleName = ModuleName
)

var (
	ErrInvalidAuthority    = errorsmod.Register(moduleName, 1, "invalid authority")
	ErrUnauthorized        = errorsmod.Register(moduleName, 2, "unauthorized")
	ErrInvalidDenom        = errorsmod.Register(moduleName, 3, "invalid denom")
	ErrInvalidPrice        = errorsmod.Register(moduleName, 4, "invalid price")
	ErrPriceNotFound       = errorsmod.Register(moduleName, 5, "price not found")
	ErrPriceStale          = errorsmod.Register(moduleName, 6, "price is stale")
	ErrDeviationTooLarge   = errorsmod.Register(moduleName, 7, "price deviation exceeds maximum")
	ErrUpdateTooFrequent   = errorsmod.Register(moduleName, 8, "price update too frequent")
	ErrProviderNotFound    = errorsmod.Register(moduleName, 9, "oracle provider not found")
	ErrProviderInactive    = errorsmod.Register(moduleName, 10, "oracle provider is inactive")
	ErrProviderSlashed     = errorsmod.Register(moduleName, 11, "oracle provider has been slashed")
	ErrConfigNotFound      = errorsmod.Register(moduleName, 12, "oracle config not found")
	ErrConfigDisabled      = errorsmod.Register(moduleName, 13, "oracle config is disabled")
	ErrMaxProvidersReached = errorsmod.Register(moduleName, 14, "maximum oracle providers reached")
	ErrInvalidProvider     = errorsmod.Register(moduleName, 15, "invalid oracle provider")
)
