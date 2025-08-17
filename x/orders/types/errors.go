package types

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/orders module sentinel errors
var (
	ErrOrderNotFound     = sdkerrors.Register(ModuleName, 1100, "order not found")
	ErrInvalidOrderID    = sdkerrors.Register(ModuleName, 1101, "invalid order ID")
	ErrOrderExists       = sdkerrors.Register(ModuleName, 1102, "order already exists")
	ErrInvalidStatus     = sdkerrors.Register(ModuleName, 1103, "invalid order status")
	ErrUnauthorized      = sdkerrors.Register(ModuleName, 1104, "unauthorized to perform action")
	ErrOrderNotPending   = sdkerrors.Register(ModuleName, 1105, "order is not in pending status")
	ErrOrderNotCancellable = sdkerrors.Register(ModuleName, 1106, "order cannot be cancelled")
	ErrInvalidAmount     = sdkerrors.Register(ModuleName, 1107, "invalid amount")
	ErrEmptyOrderItems   = sdkerrors.Register(ModuleName, 1108, "order must have at least one item")
	ErrInvalidRefundAmount = sdkerrors.Register(ModuleName, 1109, "refund amount exceeds order total")
	ErrOrderAlreadyFulfilled = sdkerrors.Register(ModuleName, 1110, "order is already fulfilled")
	ErrOrderAlreadyCancelled = sdkerrors.Register(ModuleName, 1111, "order is already cancelled")
	ErrInvalidCustomer   = sdkerrors.Register(ModuleName, 1112, "invalid customer address")
	ErrInvalidMerchant   = sdkerrors.Register(ModuleName, 1113, "invalid merchant address")
	
	// Stablecoin payment errors
	ErrInvalidStablecoin      = sdkerrors.Register(ModuleName, 1200, "invalid stablecoin denomination")
	ErrOrderAlreadyPaid       = sdkerrors.Register(ModuleName, 1201, "order has already been paid")
	ErrPaymentFailed          = sdkerrors.Register(ModuleName, 1202, "payment transaction failed")
	ErrInvalidPaymentAmount   = sdkerrors.Register(ModuleName, 1203, "invalid payment amount")
	ErrNoStablecoinPayment    = sdkerrors.Register(ModuleName, 1204, "no stablecoin payment found for order")
	ErrRefundFailed           = sdkerrors.Register(ModuleName, 1205, "refund transaction failed")
	ErrNoEscrow               = sdkerrors.Register(ModuleName, 1206, "no escrow found for order")
	ErrEscrowReleaseFailed    = sdkerrors.Register(ModuleName, 1207, "escrow release failed")
	ErrInsufficientConfirmations = sdkerrors.Register(ModuleName, 1208, "insufficient payment confirmations")
	ErrEscrowTimeout          = sdkerrors.Register(ModuleName, 1209, "escrow timeout exceeded")
)