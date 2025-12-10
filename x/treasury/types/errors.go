package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrUnauthorized        = errorsmod.Register(ModuleName, 1, "unauthorized")
	ErrInvalidReserve      = errorsmod.Register(ModuleName, 2, "invalid reserve snapshot")
	ErrInsufficientFunds   = errorsmod.Register(ModuleName, 3, "insufficient treasury funds")
	ErrProposalNotFound    = errorsmod.Register(ModuleName, 4, "spend proposal not found")
	ErrProposalNotPending  = errorsmod.Register(ModuleName, 5, "proposal is not in pending status")
	ErrTimelockNotExpired  = errorsmod.Register(ModuleName, 6, "timelock period has not expired")
	ErrProposalExpired     = errorsmod.Register(ModuleName, 7, "proposal has expired")
	ErrInvalidCategory     = errorsmod.Register(ModuleName, 8, "invalid budget category")
	ErrBudgetExceeded      = errorsmod.Register(ModuleName, 9, "budget limit exceeded for category")
	ErrInvalidAmount       = errorsmod.Register(ModuleName, 10, "invalid amount")
	ErrInvalidRecipient    = errorsmod.Register(ModuleName, 11, "invalid recipient address")
	ErrBudgetNotFound      = errorsmod.Register(ModuleName, 12, "budget not found for category")
	ErrAllocationNotFound  = errorsmod.Register(ModuleName, 13, "allocation not found")
	ErrTimelockTooShort    = errorsmod.Register(ModuleName, 14, "timelock duration too short")
	ErrMaxProposalsReached = errorsmod.Register(ModuleName, 15, "maximum pending proposals reached")
)
