package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "security"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_security"
)

// Security Alert Types
const (
	AlertTypeFraud              = "fraud"
	AlertTypeVelocity           = "velocity"
	AlertTypePattern            = "pattern"
	AlertTypeCompliance         = "compliance"
	AlertTypeUnusualActivity    = "unusual_activity"
	AlertTypeHighRisk           = "high_risk"
)

// Security Rule represents a security rule for fraud detection
type SecurityRule struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Enabled     bool      `json:"enabled"`
	Conditions  string    `json:"conditions"`  // JSON encoded conditions
	Actions     string    `json:"actions"`     // JSON encoded actions
	Severity    int32     `json:"severity"`    // 1=Low, 2=Medium, 3=High, 4=Critical
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SecurityAlert represents a security alert
type SecurityAlert struct {
	ID            string    `json:"id"`
	RuleID        string    `json:"rule_id"`
	Type          string    `json:"type"`
	Severity      int32     `json:"severity"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	TransactionID string    `json:"transaction_id"`
	Address       string    `json:"address"`
	Amount        sdk.Coins `json:"amount"`
	Data          string    `json:"data"`        // JSON encoded additional data
	Status        string    `json:"status"`      // pending, investigated, resolved, false_positive
	CreatedAt     time.Time `json:"created_at"`
	ResolvedAt    *time.Time `json:"resolved_at,omitempty"`
}

// RiskProfile represents a risk profile for an address
type RiskProfile struct {
	Address           string    `json:"address"`
	RiskScore         int32     `json:"risk_score"`      // 0-100
	VelocityScore     int32     `json:"velocity_score"`  // 0-100
	PatternScore      int32     `json:"pattern_score"`   // 0-100
	ComplianceScore   int32     `json:"compliance_score"` // 0-100
	TotalTransactions int64     `json:"total_transactions"`
	TotalVolume       sdk.Coins `json:"total_volume"`
	LastActivity      time.Time `json:"last_activity"`
	Flags             []string  `json:"flags"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// ComplianceRule represents compliance requirements
type ComplianceRule struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`              // kyc, aml, sanctions, tax
	Jurisdiction      string    `json:"jurisdiction"`
	Requirements      string    `json:"requirements"`      // JSON encoded requirements
	Enabled           bool      `json:"enabled"`
	EnforcementLevel  string    `json:"enforcement_level"` // warning, block, report
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// TransactionMonitor represents transaction monitoring data
type TransactionMonitor struct {
	TransactionHash string    `json:"transaction_hash"`
	BlockHeight     int64     `json:"block_height"`
	FromAddress     string    `json:"from_address"`
	ToAddress       string    `json:"to_address"`
	Amount          sdk.Coins `json:"amount"`
	Type            string    `json:"type"`
	RiskScore       int32     `json:"risk_score"`
	Flags           []string  `json:"flags"`
	ProcessedAt     time.Time `json:"processed_at"`
}

// Validate validates the SecurityRule
func (sr SecurityRule) Validate() error {
	if sr.ID == "" {
		return fmt.Errorf("security rule ID cannot be empty")
	}
	if sr.Name == "" {
		return fmt.Errorf("security rule name cannot be empty")
	}
	if sr.Type == "" {
		return fmt.Errorf("security rule type cannot be empty")
	}
	if sr.Severity < 1 || sr.Severity > 4 {
		return fmt.Errorf("severity must be between 1 and 4")
	}
	return nil
}

// Validate validates the SecurityAlert
func (sa SecurityAlert) Validate() error {
	if sa.ID == "" {
		return fmt.Errorf("security alert ID cannot be empty")
	}
	if sa.RuleID == "" {
		return fmt.Errorf("security alert rule ID cannot be empty")
	}
	if sa.Type == "" {
		return fmt.Errorf("security alert type cannot be empty")
	}
	if sa.Severity < 1 || sa.Severity > 4 {
		return fmt.Errorf("severity must be between 1 and 4")
	}
	return nil
}

// Validate validates the RiskProfile
func (rp RiskProfile) Validate() error {
	if rp.Address == "" {
		return fmt.Errorf("risk profile address cannot be empty")
	}
	if rp.RiskScore < 0 || rp.RiskScore > 100 {
		return fmt.Errorf("risk score must be between 0 and 100")
	}
	return nil
}

// IsHighRisk returns true if the risk profile indicates high risk
func (rp RiskProfile) IsHighRisk() bool {
	return rp.RiskScore >= 70
}

// IsCriticalRisk returns true if the risk profile indicates critical risk
func (rp RiskProfile) IsCriticalRisk() bool {
	return rp.RiskScore >= 90
}

// GetOverallRiskScore calculates an overall risk score
func (rp RiskProfile) GetOverallRiskScore() int32 {
	// Weighted average of different risk factors
	return (rp.RiskScore*40 + rp.VelocityScore*25 + rp.PatternScore*20 + rp.ComplianceScore*15) / 100
}