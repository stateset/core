package types

import (
	"fmt"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RiskLevel enumerates rough classifications used for compliance checks.
type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

// KYCLevel represents the level of KYC verification completed.
type KYCLevel string

const (
	KYCNone     KYCLevel = "none"     // No KYC completed
	KYCBasic    KYCLevel = "basic"    // Basic identity verification
	KYCStandard KYCLevel = "standard" // Standard KYC with document verification
	KYCEnhanced KYCLevel = "enhanced" // Enhanced due diligence completed
)

// ProfileStatus represents the current status of a compliance profile.
type ProfileStatus string

const (
	StatusPending  ProfileStatus = "pending"  // Awaiting verification
	StatusActive   ProfileStatus = "active"   // Fully verified and active
	StatusSuspended ProfileStatus = "suspended" // Temporarily suspended
	StatusRejected ProfileStatus = "rejected" // Verification rejected
	StatusExpired  ProfileStatus = "expired"  // Verification expired
)

// Profile captures compliance metadata tied to an address.
type Profile struct {
	Address            string        `json:"address" yaml:"address"`
	KYCLevel           KYCLevel      `json:"kyc_level" yaml:"kyc_level"`
	Risk               RiskLevel     `json:"risk" yaml:"risk"`
	Status             ProfileStatus `json:"status" yaml:"status"`
	Sanction           bool          `json:"sanction" yaml:"sanction"`
	Jurisdiction       string        `json:"jurisdiction" yaml:"jurisdiction"`         // ISO 3166-1 alpha-2 country code
	BusinessType       string        `json:"business_type" yaml:"business_type"`       // individual, business, institution
	DailyLimit         sdk.Coin      `json:"daily_limit" yaml:"daily_limit"`           // Maximum daily transaction volume
	MonthlyLimit       sdk.Coin      `json:"monthly_limit" yaml:"monthly_limit"`       // Maximum monthly transaction volume
	DailyUsed          sdk.Coin      `json:"daily_used" yaml:"daily_used"`             // Current daily usage
	MonthlyUsed        sdk.Coin      `json:"monthly_used" yaml:"monthly_used"`         // Current monthly usage
	LastLimitReset     time.Time     `json:"last_limit_reset" yaml:"last_limit_reset"` // Last time limits were reset
	VerifiedAt         time.Time     `json:"verified_at" yaml:"verified_at"`           // When verification completed
	ExpiresAt          time.Time     `json:"expires_at" yaml:"expires_at"`             // When verification expires
	Metadata           string        `json:"metadata" yaml:"metadata"`
	UpdatedBy          string        `json:"updated_by" yaml:"updated_by"`
	UpdatedAt          time.Time     `json:"updated_at" yaml:"updated_at"`
	AuditLog           []AuditEntry  `json:"audit_log" yaml:"audit_log"` // Compliance action history
}

// AuditEntry records a compliance action for audit purposes.
type AuditEntry struct {
	Timestamp time.Time `json:"timestamp" yaml:"timestamp"`
	Action    string    `json:"action" yaml:"action"`       // created, updated, suspended, reactivated, etc.
	Actor     string    `json:"actor" yaml:"actor"`         // Address that performed the action
	Reason    string    `json:"reason" yaml:"reason"`       // Reason for the action
	OldStatus string    `json:"old_status" yaml:"old_status"`
	NewStatus string    `json:"new_status" yaml:"new_status"`
}

// BlockedJurisdictions lists jurisdictions that are blocked for compliance reasons.
var BlockedJurisdictions = map[string]bool{
	"KP": true, // North Korea
	"IR": true, // Iran
	"SY": true, // Syria
	"CU": true, // Cuba
	"RU": true, // Russia (sanctions)
}

// HighRiskJurisdictions lists jurisdictions requiring enhanced due diligence.
var HighRiskJurisdictions = map[string]bool{
	"AF": true, // Afghanistan
	"BY": true, // Belarus
	"MM": true, // Myanmar
	"VE": true, // Venezuela
	"YE": true, // Yemen
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
	// Keep last 100 audit entries
	if len(p.AuditLog) >= 100 {
		p.AuditLog = p.AuditLog[1:]
	}
	p.AuditLog = append(p.AuditLog, entry)
}
