package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
)

type OrderFinancing struct {
	OrderId          string    `json:"order_id"`
	FinancingType    string    `json:"financing_type"`
	RequestedAmount  sdk.Coins `json:"requested_amount"`
	ApprovedAmount   sdk.Coins `json:"approved_amount,omitempty"`
	StablecoinDenom  string    `json:"stablecoin_denom"`
	InterestRate     math.LegacyDec `json:"interest_rate"`
	Term             int64     `json:"term"`
	Status           string    `json:"status"`
	LoanId           string    `json:"loan_id,omitempty"`
	Supplier         string    `json:"supplier"`
	Buyer            string    `json:"buyer"`
	InvoiceAmount    sdk.Coins `json:"invoice_amount"`
	DisbursedAt      *time.Time `json:"disbursed_at,omitempty"`
	DueDate          time.Time `json:"due_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type PurchaseOrderFinancing struct {
	OrderId          string    `json:"order_id"`
	Buyer            string    `json:"buyer"`
	Supplier         string    `json:"supplier"`
	OrderAmount      sdk.Coins `json:"order_amount"`
	FinancingAmount  sdk.Coins `json:"financing_amount"`
	StablecoinDenom  string    `json:"stablecoin_denom"`
	AdvanceRate      math.LegacyDec `json:"advance_rate"`
	InterestRate     math.LegacyDec `json:"interest_rate"`
	PaymentTerms     int64     `json:"payment_terms"`
	Status           string    `json:"status"`
	LoanId           string    `json:"loan_id,omitempty"`
	ApprovedBy       string    `json:"approved_by,omitempty"`
	ApprovedAt       *time.Time `json:"approved_at,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
}

type InvoiceFinancing struct {
	InvoiceId        string    `json:"invoice_id"`
	OrderId          string    `json:"order_id"`
	Seller           string    `json:"seller"`
	Buyer            string    `json:"buyer"`
	InvoiceAmount    sdk.Coins `json:"invoice_amount"`
	FinancingAmount  sdk.Coins `json:"financing_amount"`
	StablecoinDenom  string    `json:"stablecoin_denom"`
	DiscountRate     math.LegacyDec `json:"discount_rate"`
	MaturityDate     time.Time `json:"maturity_date"`
	Status           string    `json:"status"`
	LoanId           string    `json:"loan_id,omitempty"`
	FundedAt         *time.Time `json:"funded_at,omitempty"`
	RepaidAt         *time.Time `json:"repaid_at,omitempty"`
	CreatedAt        time.Time `json:"created_at"`
}

type SupplyChainFinancing struct {
	FinancingId      string    `json:"financing_id"`
	ChainPosition    string    `json:"chain_position"`
	AnchorBuyer      string    `json:"anchor_buyer"`
	Supplier         string    `json:"supplier"`
	OrderIds         []string  `json:"order_ids"`
	TotalAmount      sdk.Coins `json:"total_amount"`
	FinancingAmount  sdk.Coins `json:"financing_amount"`
	StablecoinDenom  string    `json:"stablecoin_denom"`
	InterestRate     math.LegacyDec `json:"interest_rate"`
	RiskScore        math.LegacyDec `json:"risk_score"`
	CollateralType   string    `json:"collateral_type"`
	CollateralValue  sdk.Coins `json:"collateral_value,omitempty"`
	Status           string    `json:"status"`
	PoolId           string    `json:"pool_id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

const (
	FinancingTypePurchaseOrder = "purchase_order"
	FinancingTypeInvoice       = "invoice"
	FinancingTypeSupplyChain   = "supply_chain"
	FinancingTypeInventory     = "inventory"
	
	FinancingStatusPending   = "pending"
	FinancingStatusApproved  = "approved"
	FinancingStatusFunded    = "funded"
	FinancingStatusActive    = "active"
	FinancingStatusRepaid    = "repaid"
	FinancingStatusDefaulted = "defaulted"
	FinancingStatusCancelled = "cancelled"
)

type MsgRequestOrderFinancing struct {
	Creator         string    `json:"creator"`
	OrderId         string    `json:"order_id"`
	FinancingType   string    `json:"financing_type"`
	Amount          sdk.Coins `json:"amount"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	Term            int64     `json:"term"`
}

func NewMsgRequestOrderFinancing(
	creator string,
	orderId string,
	financingType string,
	amount sdk.Coins,
	stablecoinDenom string,
	term int64,
) *MsgRequestOrderFinancing {
	return &MsgRequestOrderFinancing{
		Creator:         creator,
		OrderId:         orderId,
		FinancingType:   financingType,
		Amount:          amount,
		StablecoinDenom: stablecoinDenom,
		Term:            term,
	}
}

func (msg *MsgRequestOrderFinancing) Route() string {
	return RouterKey
}

func (msg *MsgRequestOrderFinancing) Type() string {
	return "request_order_financing"
}

func (msg *MsgRequestOrderFinancing) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestOrderFinancing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestOrderFinancing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return ErrInvalidAddress
	}
	
	if msg.OrderId == "" {
		return ErrInvalidOrderId
	}
	
	if msg.FinancingType == "" {
		return ErrInvalidFinancingType
	}
	
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return ErrInvalidAmount
	}
	
	if msg.StablecoinDenom == "" {
		return ErrInvalidDenom
	}
	
	if msg.Term <= 0 {
		return ErrInvalidTerm
	}
	
	return nil
}

type MsgApproveOrderFinancing struct {
	Authority       string    `json:"authority"`
	OrderId         string    `json:"order_id"`
	ApprovedAmount  sdk.Coins `json:"approved_amount"`
	InterestRate    math.LegacyDec `json:"interest_rate"`
	PoolId          string    `json:"pool_id"`
}

func NewMsgApproveOrderFinancing(
	authority string,
	orderId string,
	approvedAmount sdk.Coins,
	interestRate math.LegacyDec,
	poolId string,
) *MsgApproveOrderFinancing {
	return &MsgApproveOrderFinancing{
		Authority:      authority,
		OrderId:        orderId,
		ApprovedAmount: approvedAmount,
		InterestRate:   interestRate,
		PoolId:         poolId,
	}
}

func (msg *MsgApproveOrderFinancing) Route() string {
	return RouterKey
}

func (msg *MsgApproveOrderFinancing) Type() string {
	return "approve_order_financing"
}

func (msg *MsgApproveOrderFinancing) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgApproveOrderFinancing) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveOrderFinancing) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return ErrInvalidAddress
	}
	
	if msg.OrderId == "" {
		return ErrInvalidOrderId
	}
	
	if !msg.ApprovedAmount.IsValid() || msg.ApprovedAmount.IsZero() {
		return ErrInvalidAmount
	}
	
	if msg.InterestRate.IsNegative() {
		return ErrInvalidRate
	}
	
	if msg.PoolId == "" {
		return ErrInvalidPoolId
	}
	
	return nil
}