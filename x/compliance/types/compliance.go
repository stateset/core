package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ComplianceStatus string

const (
	ComplianceStatusPending  ComplianceStatus = "pending"
	ComplianceStatusApproved ComplianceStatus = "approved"
	ComplianceStatusRejected ComplianceStatus = "rejected"
	ComplianceStatusFlagged  ComplianceStatus = "flagged"
)

type RiskLevel string

const (
	RiskLevelLow    RiskLevel = "low"
	RiskLevelMedium RiskLevel = "medium"
	RiskLevelHigh   RiskLevel = "high"
	RiskLevelCritical RiskLevel = "critical"
)

type ComplianceCheck struct {
	ID              string           `json:"id"`
	TransactionID   string           `json:"transaction_id"`
	CheckType       string           `json:"check_type"`
	Status          ComplianceStatus `json:"status"`
	RiskScore       sdk.Dec          `json:"risk_score"`
	RiskLevel       RiskLevel        `json:"risk_level"`
	Details         string           `json:"details"`
	Timestamp       time.Time        `json:"timestamp"`
}

type KYCRecord struct {
	Address        string    `json:"address"`
	VerificationID string    `json:"verification_id"`
	Level          int       `json:"level"`
	Country        string    `json:"country"`
	Status         string    `json:"status"`
	VerifiedAt     time.Time `json:"verified_at"`
	ExpiresAt      time.Time `json:"expires_at"`
}

type AMLScreening struct {
	Address       string    `json:"address"`
	ScreeningID   string    `json:"screening_id"`
	IsSanctioned  bool      `json:"is_sanctioned"`
	IsPEP         bool      `json:"is_pep"`
	RiskScore     sdk.Dec   `json:"risk_score"`
	LastScreened  time.Time `json:"last_screened"`
	NextScreening time.Time `json:"next_screening"`
}

type TransactionMonitor struct {
	Address           string    `json:"address"`
	DailyVolume       sdk.Coins `json:"daily_volume"`
	MonthlyVolume     sdk.Coins `json:"monthly_volume"`
	TransactionCount  uint64    `json:"transaction_count"`
	SuspiciousCount   uint64    `json:"suspicious_count"`
	LastTransaction   time.Time `json:"last_transaction"`
}

type JurisdictionRule struct {
	FromCountry      string   `json:"from_country"`
	ToCountry        string   `json:"to_country"`
	AllowedCurrencies []string `json:"allowed_currencies"`
	MaxAmount        sdk.Int  `json:"max_amount"`
	RequiredKYCLevel int      `json:"required_kyc_level"`
	RequiresApproval bool     `json:"requires_approval"`
	TaxRate          sdk.Dec  `json:"tax_rate"`
	ReportingRequired bool    `json:"reporting_required"`
}

type ComplianceReport struct {
	ID               string    `json:"id"`
	ReportType       string    `json:"report_type"`
	Period           string    `json:"period"`
	Jurisdiction     string    `json:"jurisdiction"`
	TotalTransactions uint64   `json:"total_transactions"`
	TotalVolume      sdk.Coins `json:"total_volume"`
	FlaggedTransactions uint64 `json:"flagged_transactions"`
	GeneratedAt      time.Time `json:"generated_at"`
	SubmittedAt      time.Time `json:"submitted_at"`
	Status           string    `json:"status"`
}

type CrossBorderPayment struct {
	ID                string           `json:"id"`
	Sender            string           `json:"sender"`
	Receiver          string           `json:"receiver"`
	Amount            sdk.Coin         `json:"amount"`
	OriginCountry     string           `json:"origin_country"`
	DestinationCountry string          `json:"destination_country"`
	Purpose           string           `json:"purpose"`
	ComplianceStatus  ComplianceStatus `json:"compliance_status"`
	ComplianceChecks  []ComplianceCheck `json:"compliance_checks"`
	EstimatedFees     sdk.Coins        `json:"estimated_fees"`
	EstimatedTax      sdk.Coin         `json:"estimated_tax"`
	EstimatedArrival  time.Time        `json:"estimated_arrival"`
	CreatedAt         time.Time        `json:"created_at"`
	ProcessedAt       time.Time        `json:"processed_at"`
}

type ComplianceEngine struct {
	KYCProvider       string   `json:"kyc_provider"`
	AMLProvider       string   `json:"aml_provider"`
	SanctionLists     []string `json:"sanction_lists"`
	RiskThreshold     sdk.Dec  `json:"risk_threshold"`
	AutoApproveLimit  sdk.Int  `json:"auto_approve_limit"`
	ScreeningInterval uint64   `json:"screening_interval"`
	ReportingInterval uint64   `json:"reporting_interval"`
}

func NewComplianceCheck(txID, checkType string) ComplianceCheck {
	return ComplianceCheck{
		ID:            sdk.AccAddress([]byte(txID + checkType)).String(),
		TransactionID: txID,
		CheckType:     checkType,
		Status:        ComplianceStatusPending,
		RiskScore:     sdk.ZeroDec(),
		RiskLevel:     RiskLevelLow,
		Timestamp:     time.Now(),
	}
}

func (c *ComplianceCheck) UpdateRiskScore(score sdk.Dec) {
	c.RiskScore = score
	
	if score.LTE(sdk.NewDecWithPrec(25, 2)) {
		c.RiskLevel = RiskLevelLow
	} else if score.LTE(sdk.NewDecWithPrec(50, 2)) {
		c.RiskLevel = RiskLevelMedium
	} else if score.LTE(sdk.NewDecWithPrec(75, 2)) {
		c.RiskLevel = RiskLevelHigh
	} else {
		c.RiskLevel = RiskLevelCritical
	}
}

func (k *KYCRecord) IsValid() bool {
	return k.Status == "verified" && time.Now().Before(k.ExpiresAt)
}

func (a *AMLScreening) NeedsRescreening() bool {
	return time.Now().After(a.NextScreening)
}

func (j *JurisdictionRule) IsCompliant(amount sdk.Int, kycLevel int) bool {
	if amount.GT(j.MaxAmount) {
		return false
	}
	if kycLevel < j.RequiredKYCLevel {
		return false
	}
	return true
}

func (p *CrossBorderPayment) CalculateTotalCost() sdk.Coins {
	totalCost := sdk.NewCoins(p.Amount)
	totalCost = totalCost.Add(p.EstimatedFees...)
	if !p.EstimatedTax.IsZero() {
		totalCost = totalCost.Add(p.EstimatedTax)
	}
	return totalCost
}

func (p *CrossBorderPayment) IsCompliant() bool {
	for _, check := range p.ComplianceChecks {
		if check.Status == ComplianceStatusRejected {
			return false
		}
		if check.RiskLevel == RiskLevelCritical {
			return false
		}
	}
	return p.ComplianceStatus == ComplianceStatusApproved
}