package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUnauthorized         = sdkerrors.Register(ModuleName, 1100, "unauthorized")
	ErrOracleAlreadyExists  = sdkerrors.Register(ModuleName, 1101, "oracle provider already exists")
	ErrOracleNotFound       = sdkerrors.Register(ModuleName, 1102, "oracle provider not found")
	ErrInvalidPrice         = sdkerrors.Register(ModuleName, 1103, "invalid price")
	ErrPriceNotAvailable    = sdkerrors.Register(ModuleName, 1104, "price not available")
	ErrPriceTooOld          = sdkerrors.Register(ModuleName, 1105, "price too old")
	ErrInsufficientProviders = sdkerrors.Register(ModuleName, 1106, "insufficient oracle providers")
	ErrExcessiveDeviation   = sdkerrors.Register(ModuleName, 1107, "excessive price deviation")
	ErrInvalidAsset         = sdkerrors.Register(ModuleName, 1108, "invalid asset")
	ErrInvalidProvider      = sdkerrors.Register(ModuleName, 1109, "invalid provider")
	ErrEmergencyMode        = sdkerrors.Register(ModuleName, 1110, "oracle in emergency mode")
)