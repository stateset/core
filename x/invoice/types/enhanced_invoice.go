package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InvoiceStatus represents the enhanced invoice statuses
const (
	InvoiceStatusDraft     = "draft"
	InvoiceStatusPending   = "pending"
	InvoiceStatusApproved  = "approved"
	InvoiceStatusPaid      = "paid"
	InvoiceStatusOverdue   = "overdue"
	InvoiceStatusDisputed  = "disputed"
	InvoiceStatusCancelled = "cancelled"
	InvoiceStatusPartiallyPaid = "partially_paid"
)

// PaymentTerms represents payment terms for an invoice
type PaymentTerms struct {
	DueDays          int32   `json:"due_days"`           // Payment due in X days
	EarlyPayDiscount float64 `json:"early_pay_discount"` // Discount for early payment
	LateFeeRate      float64 `json:"late_fee_rate"`      // Late fee rate per day
	Currency         string  `json:"currency"`           // Primary currency
	AcceptedCurrencies []string `json:"accepted_currencies"` // Other accepted currencies
}

// EnhancedInvoice extends the basic invoice with advanced features
type EnhancedInvoice struct {
	BaseInvoice      Invoice      `json:"base_invoice"`
	PaymentTerms     PaymentTerms `json:"payment_terms"`
	LineItems        []LineItem   `json:"line_items"`
	TaxDetails       []TaxDetail  `json:"tax_details"`
	PaymentSchedule  []PaymentScheduleItem `json:"payment_schedule"`
	PaymentHistory   []PaymentRecord `json:"payment_history"`
	Metadata         InvoiceMetadata `json:"metadata"`
	ComplianceFlags  []string     `json:"compliance_flags"`
	RiskScore        int32        `json:"risk_score"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	DueDate          time.Time    `json:"due_date"`
	PaidAt           *time.Time   `json:"paid_at,omitempty"`
}

// LineItem represents an individual line item on an invoice
type LineItem struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Quantity    sdk.Dec   `json:"quantity"`
	UnitPrice   sdk.Coins `json:"unit_price"`
	Total       sdk.Coins `json:"total"`
	TaxRate     float64   `json:"tax_rate"`
	Category    string    `json:"category"`
}

// TaxDetail represents tax calculations
type TaxDetail struct {
	Type        string    `json:"type"`        // federal, state, local, vat, etc.
	Rate        float64   `json:"rate"`
	Basis       sdk.Coins `json:"basis"`       // Amount tax is calculated on
	Amount      sdk.Coins `json:"amount"`      // Tax amount
	Jurisdiction string   `json:"jurisdiction"`
}

// PaymentScheduleItem represents a scheduled payment
type PaymentScheduleItem struct {
	ID          string    `json:"id"`
	DueDate     time.Time `json:"due_date"`
	Amount      sdk.Coins `json:"amount"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // pending, paid, overdue
	PaidAt      *time.Time `json:"paid_at,omitempty"`
}

// PaymentRecord represents a payment made against an invoice
type PaymentRecord struct {
	ID              string    `json:"id"`
	TransactionHash string    `json:"transaction_hash"`
	Amount          sdk.Coins `json:"amount"`
	Currency        string    `json:"currency"`
	ExchangeRate    sdk.Dec   `json:"exchange_rate,omitempty"`
	PaymentMethod   string    `json:"payment_method"`
	PayerAddress    string    `json:"payer_address"`
	ProcessedAt     time.Time `json:"processed_at"`
	Fees            sdk.Coins `json:"fees,omitempty"`
	Notes           string    `json:"notes,omitempty"`
}

// InvoiceMetadata contains additional invoice information
type InvoiceMetadata struct {
	PurchaseOrderNumber string            `json:"purchase_order_number,omitempty"`
	ContractReference   string            `json:"contract_reference,omitempty"`
	ProjectCode         string            `json:"project_code,omitempty"`
	Department          string            `json:"department,omitempty"`
	CustomFields        map[string]string `json:"custom_fields,omitempty"`
	AttachmentHashes    []string          `json:"attachment_hashes,omitempty"`
}

// InvoiceAnalytics provides analytics for invoice performance
type InvoiceAnalytics struct {
	AveragePaymentTime  time.Duration `json:"average_payment_time"`
	PaymentSuccessRate  float64       `json:"payment_success_rate"`
	EarlyPaymentRate    float64       `json:"early_payment_rate"`
	DisputeRate         float64       `json:"dispute_rate"`
	TotalVolume         sdk.Coins     `json:"total_volume"`
	OutstandingBalance  sdk.Coins     `json:"outstanding_balance"`
}

// Validate validates the enhanced invoice
func (ei EnhancedInvoice) Validate() error {
	// Validate base invoice
	if err := ei.BaseInvoice.Validate(); err != nil {
		return fmt.Errorf("invalid base invoice: %w", err)
	}

	// Validate payment terms
	if ei.PaymentTerms.DueDays < 0 {
		return fmt.Errorf("due days cannot be negative")
	}
	if ei.PaymentTerms.EarlyPayDiscount < 0 || ei.PaymentTerms.EarlyPayDiscount > 1 {
		return fmt.Errorf("early payment discount must be between 0 and 1")
	}
	if ei.PaymentTerms.LateFeeRate < 0 {
		return fmt.Errorf("late fee rate cannot be negative")
	}

	// Validate line items
	if len(ei.LineItems) == 0 {
		return fmt.Errorf("invoice must have at least one line item")
	}

	totalAmount := sdk.NewCoins()
	for _, item := range ei.LineItems {
		if item.Quantity.IsNegative() {
			return fmt.Errorf("line item quantity cannot be negative")
		}
		if !item.UnitPrice.IsValid() || item.UnitPrice.IsAnyNegative() {
			return fmt.Errorf("line item unit price is invalid")
		}
		
		// Calculate expected total
		expectedTotal := item.UnitPrice.MulDec(item.Quantity)
		if !item.Total.IsEqual(expectedTotal) {
			return fmt.Errorf("line item total does not match quantity * unit price")
		}
		
		totalAmount = totalAmount.Add(item.Total...)
	}

	// Validate risk score
	if ei.RiskScore < 0 || ei.RiskScore > 100 {
		return fmt.Errorf("risk score must be between 0 and 100")
	}

	return nil
}

// CalculateTotal calculates the total invoice amount including taxes
func (ei EnhancedInvoice) CalculateTotal() sdk.Coins {
	subtotal := sdk.NewCoins()
	for _, item := range ei.LineItems {
		subtotal = subtotal.Add(item.Total...)
	}

	totalTax := sdk.NewCoins()
	for _, tax := range ei.TaxDetails {
		totalTax = totalTax.Add(tax.Amount...)
	}

	return subtotal.Add(totalTax...)
}

// CalculateAmountDue calculates the remaining amount due
func (ei EnhancedInvoice) CalculateAmountDue() sdk.Coins {
	total := ei.CalculateTotal()
	paid := sdk.NewCoins()

	for _, payment := range ei.PaymentHistory {
		paid = paid.Add(payment.Amount...)
	}

	return total.Sub(paid...)
}

// IsOverdue checks if the invoice is overdue
func (ei EnhancedInvoice) IsOverdue() bool {
	return time.Now().After(ei.DueDate) && ei.BaseInvoice.State != InvoiceStatusPaid
}

// GetPaymentProgress returns the payment progress as a percentage
func (ei EnhancedInvoice) GetPaymentProgress() float64 {
	total := ei.CalculateTotal()
	paid := sdk.NewCoins()

	for _, payment := range ei.PaymentHistory {
		paid = paid.Add(payment.Amount...)
	}

	if total.Empty() {
		return 0
	}

	// Calculate percentage based on the primary denomination
	if len(total) > 0 && len(paid) > 0 {
		totalAmount := total[0].Amount.ToDec()
		paidAmount := sdk.ZeroDec()
		
		for _, coin := range paid {
			if coin.Denom == total[0].Denom {
				paidAmount = coin.Amount.ToDec()
				break
			}
		}
		
		if totalAmount.IsZero() {
			return 0
		}
		
		return paidAmount.Quo(totalAmount).MustFloat64() * 100
	}

	return 0
}

// GetNextPaymentDue returns the next payment due from the schedule
func (ei EnhancedInvoice) GetNextPaymentDue() *PaymentScheduleItem {
	now := time.Now()
	for _, payment := range ei.PaymentSchedule {
		if payment.Status == "pending" && payment.DueDate.After(now) {
			return &payment
		}
	}
	return nil
}

// RequiresApproval checks if the invoice requires approval based on amount or risk
func (ei EnhancedInvoice) RequiresApproval() bool {
	// High-value invoices require approval
	total := ei.CalculateTotal()
	if len(total) > 0 {
		// Define threshold (this could be configurable)
		threshold := sdk.NewInt(1000000) // 1 million units
		if total[0].Amount.GT(threshold) {
			return true
		}
	}

	// High-risk invoices require approval
	if ei.RiskScore >= 70 {
		return true
	}

	// Invoices with compliance flags require approval
	if len(ei.ComplianceFlags) > 0 {
		return true
	}

	return false
}

// GeneratePaymentSchedule generates a payment schedule based on terms
func (ei *EnhancedInvoice) GeneratePaymentSchedule() {
	total := ei.CalculateTotal()
	if total.Empty() {
		return
	}

	// For now, create a simple single payment schedule
	// In production, this could support installments, milestone payments, etc.
	dueDate := ei.CreatedAt.AddDate(0, 0, int(ei.PaymentTerms.DueDays))
	
	scheduleItem := PaymentScheduleItem{
		ID:          fmt.Sprintf("%s-payment-1", ei.BaseInvoice.Id),
		DueDate:     dueDate,
		Amount:      total,
		Description: "Full payment",
		Status:      "pending",
	}

	ei.PaymentSchedule = []PaymentScheduleItem{scheduleItem}
	ei.DueDate = dueDate
}