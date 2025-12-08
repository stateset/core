package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrPaymentNotFound     = errorsmod.Register(ModuleName, 1, "payment not found")
	ErrPaymentCompleted    = errorsmod.Register(ModuleName, 2, "payment already completed")
	ErrPaymentCancelled    = errorsmod.Register(ModuleName, 3, "payment cancelled")
	ErrNotAuthorized       = errorsmod.Register(ModuleName, 4, "not authorized")
	ErrInvalidPayment      = errorsmod.Register(ModuleName, 5, "invalid payment")
	ErrInsufficientBalance = errorsmod.Register(ModuleName, 6, "insufficient balance for escrow")
	ErrInvalidAddress      = errorsmod.Register(ModuleName, 7, "invalid address")
	ErrInvalidAmount       = errorsmod.Register(ModuleName, 8, "invalid amount")
)
