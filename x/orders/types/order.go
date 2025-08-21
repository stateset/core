package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Order struct {
	Id              string        `json:"id"`
	Customer        string        `json:"customer"`
	Merchant        string        `json:"merchant"`
	Status          string        `json:"status"`
	TotalAmount     sdk.Int       `json:"total_amount"`
	Currency        string        `json:"currency"`
	Items           []OrderItem   `json:"items"`
	ShippingInfo    *ShippingInfo `json:"shipping_info,omitempty"`
	PaymentInfo     *PaymentInfo  `json:"payment_info,omitempty"`
	Metadata        string        `json:"metadata"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
	DueDate         *time.Time    `json:"due_date,omitempty"`
	FulfillmentType string        `json:"fulfillment_type"`
	Source          string        `json:"source"`
	Discounts       []Discount    `json:"discounts"`
	TaxInfo         *TaxInfo      `json:"tax_info,omitempty"`
}

type OrderItem struct {
	ProductId   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    uint64  `json:"quantity"`
	UnitPrice   sdk.Int `json:"unit_price"`
	TotalPrice  sdk.Int `json:"total_price"`
	Metadata    string  `json:"metadata"`
}

type ShippingInfo struct {
	Address        string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	PostalCode     string `json:"postal_code"`
	Method         string `json:"method"`
	Carrier        string `json:"carrier"`
	TrackingNumber string `json:"tracking_number"`
	Cost           sdk.Int `json:"cost"`
}

type PaymentInfo struct {
	PaymentMethod           string      `json:"payment_method"`
	PaymentStatus           string      `json:"payment_status"`
	TransactionId           string      `json:"transaction_id"`
	PaymentProcessor        string      `json:"payment_processor"`
	AmountPaid              sdk.Coins   `json:"amount_paid"`
	PaymentDate             *time.Time  `json:"payment_date,omitempty"`
	StablecoinDenom         *string     `json:"stablecoin_denom,omitempty"`
	ExchangeRate            *sdk.Dec    `json:"exchange_rate,omitempty"`
	UseEscrow               *bool       `json:"use_escrow,omitempty"`
	ConfirmationsRequired   *uint64     `json:"confirmations_required,omitempty"`
	ConfirmationCount       *uint64     `json:"confirmation_count,omitempty"`
	ConfirmationBlockHeight *uint64     `json:"confirmation_block_height,omitempty"`
	EscrowTimeout           *time.Time  `json:"escrow_timeout,omitempty"`
}

type Discount struct {
	Type   string  `json:"type"`
	Amount sdk.Int `json:"amount"`
	Code   string  `json:"code"`
}

type TaxInfo struct {
	TaxRate      sdk.Dec `json:"tax_rate"`
	TaxAmount    sdk.Int `json:"tax_amount"`
	TaxExempt    bool    `json:"tax_exempt"`
	TaxId        string  `json:"tax_id"`
}