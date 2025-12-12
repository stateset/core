package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RiskLevel enumerates rough classifications used for compliance checks.
type RiskLevel = string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

// KYCLevel represents the level of KYC verification completed.
type KYCLevel = string

const (
	KYCNone     KYCLevel = "none"
	KYCBasic    KYCLevel = "basic"
	KYCStandard KYCLevel = "standard"
	KYCEnhanced KYCLevel = "enhanced"
)

// ProfileStatus represents the current status of a compliance profile.
type ProfileStatus = string

const (
	StatusPending   ProfileStatus = "pending"
	StatusActive    ProfileStatus = "active"
	StatusSuspended ProfileStatus = "suspended"
	StatusRejected  ProfileStatus = "rejected"
	StatusExpired   ProfileStatus = "expired"
)

// BlockedJurisdictions lists jurisdictions that are blocked for compliance reasons.
var BlockedJurisdictions = map[string]bool{
	"KP": true,
	"IR": true,
	"SY": true,
	"CU": true,
	"RU": true,
}

// HighRiskJurisdictions lists jurisdictions requiring enhanced due diligence.
var HighRiskJurisdictions = map[string]bool{
	"AF": true,
	"BY": true,
	"MM": true,
	"VE": true,
	"YE": true,
}

// ValidateBasic ensures the profile is well-formed.
func (p Profile) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(p.Address); err != nil {
		return errorsmod.Wrap(ErrInvalidAddress, err.Error())
	}
	if p.KYCLevel == "" {
		return fmt.Errorf("kyc level cannot be empty")
	}
	if p.Risk == "" {
		return fmt.Errorf("risk level required")
	}
	if p.Jurisdiction != "" && len(p.Jurisdiction) != 2 {
		return fmt.Errorf("jurisdiction must be ISO 3166-1 alpha-2 code")
	}
	return nil
}

// IsBlocked returns true if the profile is blocked from transacting.
func (p Profile) IsBlocked() bool {
	if p.Sanction {
		return true
	}
	if p.Status != StatusActive {
		return true
	}
	if BlockedJurisdictions[p.Jurisdiction] {
		return true
	}
	return false
}

// RequiresEnhancedDueDiligence returns true if enhanced checks are required.
func (p Profile) RequiresEnhancedDueDiligence() bool {
	if p.Risk == RiskHigh {
		return true
	}
	if HighRiskJurisdictions[p.Jurisdiction] {
		return true
	}
	return false
}

// IsExpired returns true if the verification has expired.
func (p Profile) IsExpired(now time.Time) bool {
	if p.ExpiresAt.IsZero() {
		return false
	}
	return now.After(p.ExpiresAt)
}

// AddAuditEntry adds an audit log entry to the profile.
func (p *Profile) AddAuditEntry(action, actor, reason string, oldStatus, newStatus ProfileStatus) {
	entry := AuditEntry{
		Timestamp: time.Now().UTC(),
		Action:    action,
		Actor:     actor,
		Reason:    reason,
		OldStatus: string(oldStatus),
		NewStatus: string(newStatus),
	}
	// Keep last 100 audit entries.
	if len(p.AuditLog) >= 100 {
		p.AuditLog = p.AuditLog[1:]
	}
	p.AuditLog = append(p.AuditLog, entry)
}
