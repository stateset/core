package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrProfileNotFound             = errorsmod.Register(ModuleName, 1, "compliance profile not found")
	ErrUnauthorized                = errorsmod.Register(ModuleName, 2, "unauthorized")
	ErrInvalidAddress              = errorsmod.Register(ModuleName, 3, "invalid address")
	ErrSanctionedAddress           = errorsmod.Register(ModuleName, 4, "address sanctioned")
	ErrComplianceBlocked           = errorsmod.Register(ModuleName, 5, "address blocked from transacting")
	ErrProfileExpired              = errorsmod.Register(ModuleName, 6, "compliance profile expired")
	ErrLimitExceeded               = errorsmod.Register(ModuleName, 7, "transaction limit exceeded")
	ErrEnhancedDueDiligenceRequired = errorsmod.Register(ModuleName, 8, "enhanced due diligence required")
	ErrInvalidProfileStatus        = errorsmod.Register(ModuleName, 9, "invalid profile status")
	ErrBlockedJurisdiction         = errorsmod.Register(ModuleName, 10, "jurisdiction is blocked")
	ErrInvalidKYCLevel             = errorsmod.Register(ModuleName, 11, "invalid KYC level")
	ErrProfileAlreadyExists        = errorsmod.Register(ModuleName, 12, "profile already exists")
)
