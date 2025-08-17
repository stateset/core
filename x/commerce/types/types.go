package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "cosmossdk.io/errors"
)

const (
	ModuleName = "commerce"
	StoreKey   = ModuleName
	RouterKey  = ModuleName
)

// Payment Route Types
const (
	PaymentRouteDirect     = "direct"
	PaymentRouteMultiHop   = "multi_hop"
	PaymentRouteBridge     = "bridge"
	PaymentRouteLightning  = "lightning"
	PaymentRouteOptimized  = "optimized"
)

// Trade Finance Types
const (
	TradeFinanceTypeLC        = "letter_of_credit"
	TradeFinanceTypeGuarantee = "trade_guarantee"
	TradeFinanceTypeFactoring = "invoice_factoring"
	TradeFinanceTypeForwarding = "trade_forwarding"
	TradeFinanceTypeInsurance = "trade_insurance"
)

// Payment Status
const (
	PaymentStatusPending   = "pending"
	PaymentStatusProcessing = "processing"
	PaymentStatusCompleted = "completed"
	PaymentStatusFailed    = "failed"
	PaymentStatusCancelled = "cancelled"
	PaymentStatusRefunded  = "refunded"
)

// CommerceTransaction represents a comprehensive commerce transaction
type CommerceTransaction struct {
	ID                    string                 `json:"id"`
	Type                  string                 `json:"type"`
	Status               string                 `json:"status"`
	Parties              []TransactionParty     `json:"parties"`
	PaymentInfo          PaymentInfo            `json:"payment_info"`
	TradeFinanceInfo     *TradeFinanceInfo      `json:"trade_finance_info,omitempty"`
	ComplianceInfo       ComplianceInfo         `json:"compliance_info"`
	Analytics           TransactionAnalytics   `json:"analytics"`
	Metadata            map[string]interface{} `json:"metadata"`
	CreatedAt           time.Time              `json:"created_at"`
	UpdatedAt           time.Time              `json:"updated_at"`
	CompletedAt         *time.Time             `json:"completed_at,omitempty"`
	ExpiresAt           *time.Time             `json:"expires_at,omitempty"`
}

// TransactionParty represents a party in a commerce transaction
type TransactionParty struct {
	Address     string            `json:"address"`
	Role        string            `json:"role"` // buyer, seller, financier, guarantor, etc.
	Entity      EntityInfo        `json:"entity"`
	Permissions []string          `json:"permissions"`
	Signature   *PartySignature   `json:"signature,omitempty"`
}

// EntityInfo contains information about a business entity
type EntityInfo struct {
	Name           string            `json:"name"`
	Type           string            `json:"type"` // individual, corporation, partnership, etc.
	Country        string            `json:"country"`
	TaxID          string            `json:"tax_id,omitempty"`
	RegistrationID string            `json:"registration_id,omitempty"`
	KYCStatus      string            `json:"kyc_status"`
	RiskRating     int32             `json:"risk_rating"`
	Metadata       map[string]string `json:"metadata"`
}

// PaymentInfo contains comprehensive payment information
type PaymentInfo struct {
	Amount              sdk.Coins         `json:"amount"`
	Currency            string            `json:"currency"`
	ExchangeRate        sdk.Dec           `json:"exchange_rate,omitempty"`
	Route               PaymentRoute      `json:"route"`
	Fees                PaymentFees       `json:"fees"`
	Settlement          SettlementInfo    `json:"settlement"`
	CrossBorderInfo     *CrossBorderInfo  `json:"cross_border_info,omitempty"`
	InstallmentPlan     *InstallmentPlan  `json:"installment_plan,omitempty"`
}

// PaymentRoute defines how a payment will be routed
type PaymentRoute struct {
	Type        string             `json:"type"`
	Hops        []PaymentHop       `json:"hops"`
	EstimatedTime time.Duration    `json:"estimated_time"`
	Confidence    float64          `json:"confidence"`
	Backup        *PaymentRoute    `json:"backup,omitempty"`
}

// PaymentHop represents a single hop in a payment route
type PaymentHop struct {
	From        string    `json:"from"`
	To          string    `json:"to"`
	Amount      sdk.Coins `json:"amount"`
	Fee         sdk.Coins `json:"fee"`
	Network     string    `json:"network"`
	Protocol    string    `json:"protocol"`
	TimeEstimate time.Duration `json:"time_estimate"`
}

// PaymentFees breakdown of all fees
type PaymentFees struct {
	NetworkFee    sdk.Coins `json:"network_fee"`
	BridgeFee     sdk.Coins `json:"bridge_fee"`
	ExchangeFee   sdk.Coins `json:"exchange_fee"`
	ProcessingFee sdk.Coins `json:"processing_fee"`
	TotalFee      sdk.Coins `json:"total_fee"`
}

// SettlementInfo contains settlement details
type SettlementInfo struct {
	Method           string    `json:"method"`
	ExpectedTime     time.Duration `json:"expected_time"`
	FinalityBlocks   int64     `json:"finality_blocks"`
	ConfirmationHash string    `json:"confirmation_hash,omitempty"`
	SettledAt        *time.Time `json:"settled_at,omitempty"`
}

// CrossBorderInfo for international payments
type CrossBorderInfo struct {
	SourceCountry      string            `json:"source_country"`
	DestinationCountry string            `json:"destination_country"`
	CorrespondentBanks []string          `json:"correspondent_banks"`
	SwiftCode          string            `json:"swift_code,omitempty"`
	ComplianceChecks   []ComplianceCheck `json:"compliance_checks"`
	RegulatoryInfo     map[string]string `json:"regulatory_info"`
}

// InstallmentPlan for payment scheduling
type InstallmentPlan struct {
	TotalInstallments int32               `json:"total_installments"`
	Installments      []InstallmentItem   `json:"installments"`
	InterestRate      sdk.Dec             `json:"interest_rate"`
	LateFeeRate       sdk.Dec             `json:"late_fee_rate"`
}

// InstallmentItem represents a single installment
type InstallmentItem struct {
	Number      int32     `json:"number"`
	Amount      sdk.Coins `json:"amount"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
	PaidAt      *time.Time `json:"paid_at,omitempty"`
	LateFee     sdk.Coins `json:"late_fee,omitempty"`
}

// TradeFinanceInfo contains trade finance specific information
type TradeFinanceInfo struct {
	Type             string               `json:"type"`
	Instruments      []FinancialInstrument `json:"instruments"`
	TermsConditions  string               `json:"terms_conditions"`
	CollateralInfo   *CollateralInfo      `json:"collateral_info,omitempty"`
	GuaranteeInfo    *GuaranteeInfo       `json:"guarantee_info,omitempty"`
	InsuranceInfo    *InsuranceInfo       `json:"insurance_info,omitempty"`
	DocumentsRequired []string            `json:"documents_required"`
	ExpiryDate       time.Time            `json:"expiry_date"`
}

// FinancialInstrument represents various trade finance instruments
type FinancialInstrument struct {
	ID              string            `json:"id"`
	Type            string            `json:"type"`
	Amount          sdk.Coins         `json:"amount"`
	Currency        string            `json:"currency"`
	Issuer          string            `json:"issuer"`
	Beneficiary     string            `json:"beneficiary"`
	Terms           map[string]string `json:"terms"`
	Status          string            `json:"status"`
	IssuedAt        time.Time         `json:"issued_at"`
	ExpiresAt       time.Time         `json:"expires_at"`
	UtilizedAmount  sdk.Coins         `json:"utilized_amount"`
}

// CollateralInfo for secured trade finance
type CollateralInfo struct {
	Type           string            `json:"type"`
	Value          sdk.Coins         `json:"value"`
	Description    string            `json:"description"`
	Location       string            `json:"location"`
	Valuation      ValuationInfo     `json:"valuation"`
	Insurance      *InsuranceInfo    `json:"insurance,omitempty"`
	LienHolders    []string          `json:"lien_holders"`
	Documentation  []string          `json:"documentation"`
}

// GuaranteeInfo for trade guarantees
type GuaranteeInfo struct {
	Guarantor       string            `json:"guarantor"`
	Amount          sdk.Coins         `json:"amount"`
	Type            string            `json:"type"`
	Terms           string            `json:"terms"`
	ValidUntil      time.Time         `json:"valid_until"`
	Beneficiary     string            `json:"beneficiary"`
	ClaimConditions []string          `json:"claim_conditions"`
	Status          string            `json:"status"`
}

// InsuranceInfo for trade insurance
type InsuranceInfo struct {
	Provider        string            `json:"provider"`
	PolicyNumber    string            `json:"policy_number"`
	Coverage        sdk.Coins         `json:"coverage"`
	Premium         sdk.Coins         `json:"premium"`
	CoverageType    string            `json:"coverage_type"`
	ValidFrom       time.Time         `json:"valid_from"`
	ValidUntil      time.Time         `json:"valid_until"`
	Beneficiaries   []string          `json:"beneficiaries"`
	ClaimHistory    []InsuranceClaim  `json:"claim_history"`
}

// InsuranceClaim represents an insurance claim
type InsuranceClaim struct {
	ID          string    `json:"id"`
	Amount      sdk.Coins `json:"amount"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"`
	ClaimedAt   time.Time `json:"claimed_at"`
	ProcessedAt *time.Time `json:"processed_at,omitempty"`
	PaidAt      *time.Time `json:"paid_at,omitempty"`
}

// ValuationInfo for asset valuation
type ValuationInfo struct {
	Method        string    `json:"method"`
	Valuator      string    `json:"valuator"`
	Value         sdk.Coins `json:"value"`
	Date          time.Time `json:"date"`
	Currency      string    `json:"currency"`
	Confidence    float64   `json:"confidence"`
	Documentation []string  `json:"documentation"`
}

// ComplianceInfo contains compliance and regulatory information
type ComplianceInfo struct {
	Checks          []ComplianceCheck    `json:"checks"`
	RiskAssessment  RiskAssessment       `json:"risk_assessment"`
	AMLStatus       string               `json:"aml_status"`
	KYCStatus       string               `json:"kyc_status"`
	SanctionStatus  string               `json:"sanction_status"`
	TaxCompliance   TaxComplianceInfo    `json:"tax_compliance"`
	Regulations     []ApplicableRegulation `json:"regulations"`
	Authorizations  []string             `json:"authorizations"`
}

// ComplianceCheck represents a single compliance check
type ComplianceCheck struct {
	Type        string            `json:"type"`
	Status      string            `json:"status"`
	Result      string            `json:"result"`
	Timestamp   time.Time         `json:"timestamp"`
	Details     map[string]string `json:"details"`
	Score       int32             `json:"score"`
	Automated   bool              `json:"automated"`
}

// RiskAssessment contains risk analysis
type RiskAssessment struct {
	OverallScore    int32             `json:"overall_score"`
	CreditRisk      int32             `json:"credit_risk"`
	CountryRisk     int32             `json:"country_risk"`
	CurrencyRisk    int32             `json:"currency_risk"`
	OperationalRisk int32             `json:"operational_risk"`
	MarketRisk      int32             `json:"market_risk"`
	Factors         []RiskFactor      `json:"factors"`
	Mitigation      []string          `json:"mitigation"`
	LastUpdated     time.Time         `json:"last_updated"`
}

// RiskFactor represents an individual risk factor
type RiskFactor struct {
	Type        string  `json:"type"`
	Score       int32   `json:"score"`
	Weight      float64 `json:"weight"`
	Description string  `json:"description"`
	Impact      string  `json:"impact"`
}

// TaxComplianceInfo for tax-related compliance
type TaxComplianceInfo struct {
	JurisdictionRules []TaxJurisdiction `json:"jurisdiction_rules"`
	WithholdingTax    sdk.Coins         `json:"withholding_tax"`
	TaxReporting      []TaxReport       `json:"tax_reporting"`
	TaxIdentifiers    map[string]string `json:"tax_identifiers"`
}

// TaxJurisdiction represents tax rules for a jurisdiction
type TaxJurisdiction struct {
	Country     string    `json:"country"`
	TaxRate     sdk.Dec   `json:"tax_rate"`
	TaxType     string    `json:"tax_type"`
	Threshold   sdk.Coins `json:"threshold"`
	Exemptions  []string  `json:"exemptions"`
	ReportingRequired bool `json:"reporting_required"`
}

// TaxReport represents a tax report
type TaxReport struct {
	Type         string    `json:"type"`
	Jurisdiction string    `json:"jurisdiction"`
	Amount       sdk.Coins `json:"amount"`
	DueDate      time.Time `json:"due_date"`
	Status       string    `json:"status"`
	Reference    string    `json:"reference"`
}

// ApplicableRegulation represents applicable regulations
type ApplicableRegulation struct {
	Name         string            `json:"name"`
	Jurisdiction string            `json:"jurisdiction"`
	Type         string            `json:"type"`
	Requirements []string          `json:"requirements"`
	Compliance   string            `json:"compliance"`
	Evidence     []string          `json:"evidence"`
	LastChecked  time.Time         `json:"last_checked"`
}

// TransactionAnalytics contains analytics and metrics
type TransactionAnalytics struct {
	ProcessingTime    time.Duration     `json:"processing_time"`
	Cost              AnalyticsCost     `json:"cost"`
	Performance       PerformanceMetrics `json:"performance"`
	UserExperience    UXMetrics         `json:"user_experience"`
	Carbon            CarbonFootprint   `json:"carbon"`
	Efficiency        EfficiencyMetrics `json:"efficiency"`
}

// AnalyticsCost breakdown of transaction costs
type AnalyticsCost struct {
	TotalCost       sdk.Coins `json:"total_cost"`
	NetworkCost     sdk.Coins `json:"network_cost"`
	ProcessingCost  sdk.Coins `json:"processing_cost"`
	ComplianceCost  sdk.Coins `json:"compliance_cost"`
	OpportunityCost sdk.Coins `json:"opportunity_cost"`
}

// PerformanceMetrics for transaction performance
type PerformanceMetrics struct {
	Throughput      float64 `json:"throughput"`
	Latency         time.Duration `json:"latency"`
	SuccessRate     float64 `json:"success_rate"`
	ErrorRate       float64 `json:"error_rate"`
	RetryCount      int32   `json:"retry_count"`
	OptimizationScore int32 `json:"optimization_score"`
}

// UXMetrics for user experience
type UXMetrics struct {
	StepsCompleted  int32         `json:"steps_completed"`
	TimeToComplete  time.Duration `json:"time_to_complete"`
	UserSatisfaction float64      `json:"user_satisfaction"`
	FrictionPoints  []string      `json:"friction_points"`
	Abandonment     bool          `json:"abandonment"`
}

// CarbonFootprint for environmental impact
type CarbonFootprint struct {
	TotalEmissions    float64 `json:"total_emissions"` // in kg CO2
	NetworkEmissions  float64 `json:"network_emissions"`
	ComputeEmissions  float64 `json:"compute_emissions"`
	Offsets           float64 `json:"offsets"`
	NetEmissions      float64 `json:"net_emissions"`
}

// EfficiencyMetrics for overall efficiency
type EfficiencyMetrics struct {
	CostEfficiency    float64 `json:"cost_efficiency"`
	TimeEfficiency    float64 `json:"time_efficiency"`
	ResourceUtilization float64 `json:"resource_utilization"`
	WasteReduction    float64 `json:"waste_reduction"`
	AutomationLevel   float64 `json:"automation_level"`
}

// PartySignature represents a digital signature
type PartySignature struct {
	Algorithm   string    `json:"algorithm"`
	Signature   []byte    `json:"signature"`
	PublicKey   []byte    `json:"public_key"`
	Timestamp   time.Time `json:"timestamp"`
	ValidUntil  time.Time `json:"valid_until"`
	SignedData  []byte    `json:"signed_data"`
}

// Validate validates a CommerceTransaction
func (ct CommerceTransaction) Validate() error {
	if ct.ID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "transaction ID cannot be empty")
	}

	if len(ct.Parties) < 2 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "transaction must have at least 2 parties")
	}

	if !ct.PaymentInfo.Amount.IsValid() || ct.PaymentInfo.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "payment amount must be valid and positive")
	}

	// Validate parties
	for _, party := range ct.Parties {
		if err := party.Validate(); err != nil {
			return fmt.Errorf("invalid party: %w", err)
		}
	}

	return nil
}

// Validate validates a TransactionParty
func (tp TransactionParty) Validate() error {
	if tp.Address == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "party address cannot be empty")
	}

	if tp.Role == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "party role cannot be empty")
	}

	return nil
}