package ibc

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrInvalidAmount       = errorsmod.Register("settlement-ibc", 1, "invalid amount")
	ErrEscrowNotFound      = errorsmod.Register("settlement-ibc", 2, "escrow not found")
	ErrInvalidEscrowStatus = errorsmod.Register("settlement-ibc", 3, "invalid escrow status")
	ErrUnauthorized        = errorsmod.Register("settlement-ibc", 4, "unauthorized")
	ErrInvalidMemo         = errorsmod.Register("settlement-ibc", 5, "invalid memo format")
)
