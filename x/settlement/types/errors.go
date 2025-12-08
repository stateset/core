package types

import (
	errorsmod "cosmossdk.io/errors"
)

// Module errors
var (
	ErrInvalidSettlement          = errorsmod.Register(ModuleName, 1, "invalid settlement")
	ErrSettlementNotFound         = errorsmod.Register(ModuleName, 2, "settlement not found")
	ErrInsufficientFunds          = errorsmod.Register(ModuleName, 3, "insufficient funds")
	ErrUnauthorized               = errorsmod.Register(ModuleName, 4, "unauthorized")
	ErrSettlementCompleted        = errorsmod.Register(ModuleName, 5, "settlement already completed")
	ErrSettlementCancelled        = errorsmod.Register(ModuleName, 6, "settlement cancelled")
	ErrSettlementExpired          = errorsmod.Register(ModuleName, 7, "settlement expired")
	ErrInvalidAmount              = errorsmod.Register(ModuleName, 8, "invalid amount")
	ErrInvalidRecipient           = errorsmod.Register(ModuleName, 9, "invalid recipient")
	ErrMerchantNotFound           = errorsmod.Register(ModuleName, 10, "merchant not found")
	ErrMerchantInactive           = errorsmod.Register(ModuleName, 11, "merchant is inactive")
	ErrBatchNotFound              = errorsmod.Register(ModuleName, 12, "batch not found")
	ErrBatchAlreadySettled        = errorsmod.Register(ModuleName, 13, "batch already settled")
	ErrChannelNotFound            = errorsmod.Register(ModuleName, 14, "payment channel not found")
	ErrChannelClosed              = errorsmod.Register(ModuleName, 15, "payment channel is closed")
	ErrChannelExpired             = errorsmod.Register(ModuleName, 16, "payment channel has expired")
	ErrChannelInsufficientBalance = errorsmod.Register(ModuleName, 17, "payment channel has insufficient balance")
	ErrInvalidNonce               = errorsmod.Register(ModuleName, 18, "invalid nonce")
	ErrSettlementBelowMin         = errorsmod.Register(ModuleName, 19, "settlement amount below minimum")
	ErrSettlementAboveMax         = errorsmod.Register(ModuleName, 20, "settlement amount above maximum")
	ErrComplianceCheckFailed      = errorsmod.Register(ModuleName, 21, "compliance check failed")
	ErrInvalidDenom               = errorsmod.Register(ModuleName, 22, "invalid denomination, expected stablecoin")
	ErrSettlementTooSmall         = errorsmod.Register(ModuleName, 23, "settlement amount too small")
	ErrSettlementTooLarge         = errorsmod.Register(ModuleName, 24, "settlement amount too large")
	ErrBatchTooLarge              = errorsmod.Register(ModuleName, 25, "batch size exceeds maximum")
	ErrChannelNotExpired          = errorsmod.Register(ModuleName, 26, "channel has not expired yet")
	ErrInvalidFeeCollector        = errorsmod.Register(ModuleName, 27, "invalid fee collector address")
	ErrInvalidEscrowExpiration    = errorsmod.Register(ModuleName, 28, "escrow expiration exceeds maximum")
	ErrInvalidChannelExpiration   = errorsmod.Register(ModuleName, 29, "channel expiration out of bounds")
)
