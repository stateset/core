package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrOrderNotFound         = errorsmod.Register(ModuleName, 2, "order not found")
	ErrInvalidOrder          = errorsmod.Register(ModuleName, 3, "invalid order")
	ErrUnauthorized          = errorsmod.Register(ModuleName, 4, "unauthorized")
	ErrInvalidStatus         = errorsmod.Register(ModuleName, 5, "invalid order status")
	ErrInvalidTransition     = errorsmod.Register(ModuleName, 6, "invalid status transition")
	ErrOrderExpired          = errorsmod.Register(ModuleName, 7, "order expired")
	ErrInsufficientFunds     = errorsmod.Register(ModuleName, 8, "insufficient funds")
	ErrPaymentFailed         = errorsmod.Register(ModuleName, 9, "payment failed")
	ErrAlreadyPaid           = errorsmod.Register(ModuleName, 10, "order already paid")
	ErrNotPaid               = errorsmod.Register(ModuleName, 11, "order not paid")
	ErrCannotRefund          = errorsmod.Register(ModuleName, 12, "order cannot be refunded")
	ErrCannotCancel          = errorsmod.Register(ModuleName, 13, "order cannot be cancelled")
	ErrDisputeNotFound       = errorsmod.Register(ModuleName, 14, "dispute not found")
	ErrDisputeAlreadyExists  = errorsmod.Register(ModuleName, 15, "dispute already exists")
	ErrCannotDispute         = errorsmod.Register(ModuleName, 16, "order cannot be disputed")
	ErrInvalidAmount         = errorsmod.Register(ModuleName, 17, "invalid amount")
	ErrComplianceFailed      = errorsmod.Register(ModuleName, 18, "compliance check failed")
	ErrSettlementFailed      = errorsmod.Register(ModuleName, 19, "settlement failed")
	ErrOrderAlreadyCompleted = errorsmod.Register(ModuleName, 20, "order already completed")
	ErrEmptyItems            = errorsmod.Register(ModuleName, 21, "order must have at least one item")
	ErrInvalidMerchant       = errorsmod.Register(ModuleName, 22, "invalid merchant address")
	ErrInvalidCustomer       = errorsmod.Register(ModuleName, 23, "invalid customer address")
)
