package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
)

type WorkingCapitalLoan struct {
	Id               string    `json:"id"`
	Borrower         string    `json:"borrower"`
	Amount           sdk.Coins `json:"amount"`
	StablecoinDenom  string    `json:"stablecoin_denom"`
	Purpose          string    `json:"purpose"`
	OrderId          string    `json:"order_id,omitempty"`
	Status           string    `json:"status"`
	InterestRate     math.LegacyDec `json:"interest_rate"`
	Term             int64     `json:"term"`
	Collateral       sdk.Coins `json:"collateral,omitempty"`
	CollateralRatio  math.LegacyDec `json:"collateral_ratio"`
	DisbursedAt      *time.Time `json:"disbursed_at,omitempty"`
	DueDate          time.Time `json:"due_date"`
	RepaidAmount     sdk.Coins `json:"repaid_amount"`
	OutstandingAmount sdk.Coins `json:"outstanding_amount"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Metadata         string    `json:"metadata"`
}

type WorkingCapitalRequest struct {
	Id              string    `json:"id"`
	Requester       string    `json:"requester"`
	BusinessProfile *BusinessProfile `json:"business_profile"`
	RequestedAmount sdk.Coins `json:"requested_amount"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	Purpose         string    `json:"purpose"`
	OrderId         string    `json:"order_id,omitempty"`
	ProposedCollateral sdk.Coins `json:"proposed_collateral,omitempty"`
	RequestedTerm   int64     `json:"requested_term"`
	Status          string    `json:"status"`
	RiskScore       math.LegacyDec `json:"risk_score"`
	ApprovedAmount  sdk.Coins `json:"approved_amount,omitempty"`
	ApprovedRate    math.LegacyDec `json:"approved_rate,omitempty"`
	ApprovedBy      string    `json:"approved_by,omitempty"`
	ApprovedAt      *time.Time `json:"approved_at,omitempty"`
	RejectionReason string    `json:"rejection_reason,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type BusinessProfile struct {
	BusinessId      string    `json:"business_id"`
	BusinessName    string    `json:"business_name"`
	BusinessType    string    `json:"business_type"`
	CreditScore     int       `json:"credit_score"`
	MonthlyRevenue  sdk.Coins `json:"monthly_revenue"`
	YearsInBusiness int       `json:"years_in_business"`
	PreviousLoans   []string  `json:"previous_loans"`
	RepaymentHistory string   `json:"repayment_history"`
}

type WorkingCapitalPool struct {
	PoolId          string    `json:"pool_id"`
	Name            string    `json:"name"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	TotalFunds      sdk.Coins `json:"total_funds"`
	AvailableFunds  sdk.Coins `json:"available_funds"`
	LentFunds       sdk.Coins `json:"lent_funds"`
	DefaultRate     math.LegacyDec `json:"default_rate"`
	AverageInterestRate math.LegacyDec `json:"average_interest_rate"`
	MinLoanAmount   sdk.Int   `json:"min_loan_amount"`
	MaxLoanAmount   sdk.Int   `json:"max_loan_amount"`
	MinCollateralRatio math.LegacyDec `json:"min_collateral_ratio"`
	MaxTerm         int64     `json:"max_term"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type RepaymentSchedule struct {
	LoanId           string    `json:"loan_id"`
	InstallmentNumber int      `json:"installment_number"`
	DueDate          time.Time `json:"due_date"`
	PrincipalAmount  sdk.Coins `json:"principal_amount"`
	InterestAmount   sdk.Coins `json:"interest_amount"`
	TotalAmount      sdk.Coins `json:"total_amount"`
	Status           string    `json:"status"`
	PaidAt           *time.Time `json:"paid_at,omitempty"`
	PaidAmount       sdk.Coins `json:"paid_amount,omitempty"`
}

type CreditLine struct {
	Id              string    `json:"id"`
	Borrower        string    `json:"borrower"`
	MaxLimit        sdk.Coins `json:"max_limit"`
	AvailableLimit  sdk.Coins `json:"available_limit"`
	UsedAmount      sdk.Coins `json:"used_amount"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	InterestRate    math.LegacyDec `json:"interest_rate"`
	Status          string    `json:"status"`
	ExpiryDate      time.Time `json:"expiry_date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

const (
	LoanStatusPending    = "pending"
	LoanStatusApproved   = "approved"
	LoanStatusDisbursed  = "disbursed"
	LoanStatusActive     = "active"
	LoanStatusOverdue    = "overdue"
	LoanStatusRepaid     = "repaid"
	LoanStatusDefaulted  = "defaulted"
	LoanStatusCancelled  = "cancelled"
)

const (
	PoolStatusActive   = "active"
	PoolStatusPaused   = "paused"
	PoolStatusDepleted = "depleted"
	PoolStatusClosed   = "closed"
)

const (
	RepaymentStatusPending  = "pending"
	RepaymentStatusPaid     = "paid"
	RepaymentStatusOverdue  = "overdue"
	RepaymentStatusPartial  = "partial"
	RepaymentStatusWaived   = "waived"
)