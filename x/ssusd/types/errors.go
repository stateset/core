package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrInvalidCollateral       = sdkerrors.Register(ModuleName, 1100, "invalid collateral")
	ErrInsufficientCollateral  = sdkerrors.Register(ModuleName, 1101, "insufficient collateral")
	ErrDebtCeilingExceeded     = sdkerrors.Register(ModuleName, 1102, "debt ceiling exceeded")
	ErrPositionNotFound        = sdkerrors.Register(ModuleName, 1103, "position not found")
	ErrInvalidAmount           = sdkerrors.Register(ModuleName, 1104, "invalid amount")
	ErrNotLiquidatable         = sdkerrors.Register(ModuleName, 1105, "position not liquidatable")
	ErrAgentExists             = sdkerrors.Register(ModuleName, 1106, "agent already exists")
	ErrAgentNotFound           = sdkerrors.Register(ModuleName, 1107, "agent not found")
	ErrAgentInactive           = sdkerrors.Register(ModuleName, 1108, "agent inactive")
	ErrInsufficientReputation  = sdkerrors.Register(ModuleName, 1109, "insufficient reputation")
	ErrUnauthorized            = sdkerrors.Register(ModuleName, 1110, "unauthorized")
	ErrAuctionNotFound         = sdkerrors.Register(ModuleName, 1111, "auction not found")
	ErrAuctionInactive         = sdkerrors.Register(ModuleName, 1112, "auction inactive")
	ErrAuctionExpired          = sdkerrors.Register(ModuleName, 1113, "auction expired")
	ErrInsufficientBid         = sdkerrors.Register(ModuleName, 1114, "insufficient bid")
	ErrProviderNotFound        = sdkerrors.Register(ModuleName, 1115, "stability provider not found")
	ErrEmergencyShutdown       = sdkerrors.Register(ModuleName, 1116, "emergency shutdown active")
	ErrInvalidOraclePrice      = sdkerrors.Register(ModuleName, 1117, "invalid oracle price")
	ErrOracleNotWhitelisted    = sdkerrors.Register(ModuleName, 1118, "oracle not whitelisted")
	ErrInvalidProof            = sdkerrors.Register(ModuleName, 1119, "invalid proof")
	ErrSystemOverloaded        = sdkerrors.Register(ModuleName, 1120, "system overloaded")
)