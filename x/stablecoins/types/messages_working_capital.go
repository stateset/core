package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/math"
	"cosmossdk.io/errors"
)

const (
	TypeMsgRequestWorkingCapital = "request_working_capital"
	TypeMsgApproveWorkingCapital = "approve_working_capital"
	TypeMsgDisburseWorkingCapital = "disburse_working_capital"
	TypeMsgRepayWorkingCapital = "repay_working_capital"
	TypeMsgCreateCapitalPool = "create_capital_pool"
	TypeMsgFundCapitalPool = "fund_capital_pool"
	TypeMsgCreateCreditLine = "create_credit_line"
	TypeMsgDrawFromCreditLine = "draw_from_credit_line"
)

type MsgRequestWorkingCapital struct {
	Creator         string    `json:"creator"`
	Amount          sdk.Coins `json:"amount"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	Purpose         string    `json:"purpose"`
	OrderId         string    `json:"order_id,omitempty"`
	Term            int64     `json:"term"`
	Collateral      sdk.Coins `json:"collateral,omitempty"`
	BusinessProfile *BusinessProfile `json:"business_profile"`
}

func NewMsgRequestWorkingCapital(
	creator string,
	amount sdk.Coins,
	stablecoinDenom string,
	purpose string,
	term int64,
	collateral sdk.Coins,
	businessProfile *BusinessProfile,
) *MsgRequestWorkingCapital {
	return &MsgRequestWorkingCapital{
		Creator:         creator,
		Amount:          amount,
		StablecoinDenom: stablecoinDenom,
		Purpose:         purpose,
		Term:            term,
		Collateral:      collateral,
		BusinessProfile: businessProfile,
	}
}

func (msg *MsgRequestWorkingCapital) Route() string {
	return RouterKey
}

func (msg *MsgRequestWorkingCapital) Type() string {
	return TypeMsgRequestWorkingCapital
}

func (msg *MsgRequestWorkingCapital) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestWorkingCapital) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestWorkingCapital) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "amount must be positive")
	}
	
	if msg.StablecoinDenom == "" {
		return errors.Wrap(ErrInvalidDenom, "stablecoin denom cannot be empty")
	}
	
	if msg.Term <= 0 {
		return errors.Wrap(ErrInvalidTerm, "term must be positive")
	}
	
	return nil
}

type MsgApproveWorkingCapital struct {
	Authority       string    `json:"authority"`
	RequestId       string    `json:"request_id"`
	ApprovedAmount  sdk.Coins `json:"approved_amount"`
	InterestRate    math.LegacyDec `json:"interest_rate"`
	CollateralRatio math.LegacyDec `json:"collateral_ratio"`
}

func NewMsgApproveWorkingCapital(
	authority string,
	requestId string,
	approvedAmount sdk.Coins,
	interestRate math.LegacyDec,
	collateralRatio math.LegacyDec,
) *MsgApproveWorkingCapital {
	return &MsgApproveWorkingCapital{
		Authority:       authority,
		RequestId:       requestId,
		ApprovedAmount:  approvedAmount,
		InterestRate:    interestRate,
		CollateralRatio: collateralRatio,
	}
}

func (msg *MsgApproveWorkingCapital) Route() string {
	return RouterKey
}

func (msg *MsgApproveWorkingCapital) Type() string {
	return TypeMsgApproveWorkingCapital
}

func (msg *MsgApproveWorkingCapital) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgApproveWorkingCapital) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveWorkingCapital) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	
	if msg.RequestId == "" {
		return errors.Wrap(ErrInvalidRequest, "request ID cannot be empty")
	}
	
	if !msg.ApprovedAmount.IsValid() || msg.ApprovedAmount.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "approved amount must be positive")
	}
	
	if msg.InterestRate.IsNegative() {
		return errors.Wrap(ErrInvalidRate, "interest rate cannot be negative")
	}
	
	if msg.CollateralRatio.IsNegative() {
		return errors.Wrap(ErrInvalidRatio, "collateral ratio cannot be negative")
	}
	
	return nil
}

type MsgDisburseWorkingCapital struct {
	Authority string `json:"authority"`
	LoanId    string `json:"loan_id"`
	PoolId    string `json:"pool_id"`
}

func NewMsgDisburseWorkingCapital(authority, loanId, poolId string) *MsgDisburseWorkingCapital {
	return &MsgDisburseWorkingCapital{
		Authority: authority,
		LoanId:    loanId,
		PoolId:    poolId,
	}
}

func (msg *MsgDisburseWorkingCapital) Route() string {
	return RouterKey
}

func (msg *MsgDisburseWorkingCapital) Type() string {
	return TypeMsgDisburseWorkingCapital
}

func (msg *MsgDisburseWorkingCapital) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgDisburseWorkingCapital) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisburseWorkingCapital) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	
	if msg.LoanId == "" {
		return errors.Wrap(ErrInvalidLoan, "loan ID cannot be empty")
	}
	
	if msg.PoolId == "" {
		return errors.Wrap(ErrInvalidPool, "pool ID cannot be empty")
	}
	
	return nil
}

type MsgRepayWorkingCapital struct {
	Creator string    `json:"creator"`
	LoanId  string    `json:"loan_id"`
	Amount  sdk.Coins `json:"amount"`
}

func NewMsgRepayWorkingCapital(creator, loanId string, amount sdk.Coins) *MsgRepayWorkingCapital {
	return &MsgRepayWorkingCapital{
		Creator: creator,
		LoanId:  loanId,
		Amount:  amount,
	}
}

func (msg *MsgRepayWorkingCapital) Route() string {
	return RouterKey
}

func (msg *MsgRepayWorkingCapital) Type() string {
	return TypeMsgRepayWorkingCapital
}

func (msg *MsgRepayWorkingCapital) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayWorkingCapital) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayWorkingCapital) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if msg.LoanId == "" {
		return errors.Wrap(ErrInvalidLoan, "loan ID cannot be empty")
	}
	
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "repayment amount must be positive")
	}
	
	return nil
}

type MsgCreateCapitalPool struct {
	Creator            string    `json:"creator"`
	Name               string    `json:"name"`
	StablecoinDenom    string    `json:"stablecoin_denom"`
	InitialFunds       sdk.Coins `json:"initial_funds"`
	MinLoanAmount      sdk.Int   `json:"min_loan_amount"`
	MaxLoanAmount      sdk.Int   `json:"max_loan_amount"`
	MinCollateralRatio math.LegacyDec `json:"min_collateral_ratio"`
	MaxTerm            int64     `json:"max_term"`
}

func NewMsgCreateCapitalPool(
	creator string,
	name string,
	stablecoinDenom string,
	initialFunds sdk.Coins,
	minLoanAmount sdk.Int,
	maxLoanAmount sdk.Int,
	minCollateralRatio math.LegacyDec,
	maxTerm int64,
) *MsgCreateCapitalPool {
	return &MsgCreateCapitalPool{
		Creator:            creator,
		Name:               name,
		StablecoinDenom:    stablecoinDenom,
		InitialFunds:       initialFunds,
		MinLoanAmount:      minLoanAmount,
		MaxLoanAmount:      maxLoanAmount,
		MinCollateralRatio: minCollateralRatio,
		MaxTerm:            maxTerm,
	}
}

func (msg *MsgCreateCapitalPool) Route() string {
	return RouterKey
}

func (msg *MsgCreateCapitalPool) Type() string {
	return TypeMsgCreateCapitalPool
}

func (msg *MsgCreateCapitalPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCapitalPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCapitalPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if msg.Name == "" {
		return errors.Wrap(ErrInvalidPool, "pool name cannot be empty")
	}
	
	if msg.StablecoinDenom == "" {
		return errors.Wrap(ErrInvalidDenom, "stablecoin denom cannot be empty")
	}
	
	if !msg.InitialFunds.IsValid() {
		return errors.Wrap(ErrInvalidAmount, "invalid initial funds")
	}
	
	if msg.MinLoanAmount.IsNegative() || msg.MaxLoanAmount.IsNegative() {
		return errors.Wrap(ErrInvalidAmount, "loan amounts cannot be negative")
	}
	
	if msg.MinLoanAmount.GT(msg.MaxLoanAmount) {
		return errors.Wrap(ErrInvalidAmount, "min loan amount cannot exceed max loan amount")
	}
	
	if msg.MinCollateralRatio.IsNegative() {
		return errors.Wrap(ErrInvalidRatio, "collateral ratio cannot be negative")
	}
	
	if msg.MaxTerm <= 0 {
		return errors.Wrap(ErrInvalidTerm, "max term must be positive")
	}
	
	return nil
}

type MsgFundCapitalPool struct {
	Creator string    `json:"creator"`
	PoolId  string    `json:"pool_id"`
	Amount  sdk.Coins `json:"amount"`
}

func NewMsgFundCapitalPool(creator, poolId string, amount sdk.Coins) *MsgFundCapitalPool {
	return &MsgFundCapitalPool{
		Creator: creator,
		PoolId:  poolId,
		Amount:  amount,
	}
}

func (msg *MsgFundCapitalPool) Route() string {
	return RouterKey
}

func (msg *MsgFundCapitalPool) Type() string {
	return TypeMsgFundCapitalPool
}

func (msg *MsgFundCapitalPool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFundCapitalPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFundCapitalPool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if msg.PoolId == "" {
		return errors.Wrap(ErrInvalidPool, "pool ID cannot be empty")
	}
	
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "funding amount must be positive")
	}
	
	return nil
}

type MsgCreateCreditLine struct {
	Creator         string    `json:"creator"`
	Borrower        string    `json:"borrower"`
	MaxLimit        sdk.Coins `json:"max_limit"`
	StablecoinDenom string    `json:"stablecoin_denom"`
	InterestRate    math.LegacyDec `json:"interest_rate"`
	ExpiryMonths    int64     `json:"expiry_months"`
}

func NewMsgCreateCreditLine(
	creator string,
	borrower string,
	maxLimit sdk.Coins,
	stablecoinDenom string,
	interestRate math.LegacyDec,
	expiryMonths int64,
) *MsgCreateCreditLine {
	return &MsgCreateCreditLine{
		Creator:         creator,
		Borrower:        borrower,
		MaxLimit:        maxLimit,
		StablecoinDenom: stablecoinDenom,
		InterestRate:    interestRate,
		ExpiryMonths:    expiryMonths,
	}
}

func (msg *MsgCreateCreditLine) Route() string {
	return RouterKey
}

func (msg *MsgCreateCreditLine) Type() string {
	return TypeMsgCreateCreditLine
}

func (msg *MsgCreateCreditLine) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCreditLine) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCreditLine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	_, err = sdk.AccAddressFromBech32(msg.Borrower)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid borrower address (%s)", err)
	}
	
	if !msg.MaxLimit.IsValid() || msg.MaxLimit.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "max limit must be positive")
	}
	
	if msg.StablecoinDenom == "" {
		return errors.Wrap(ErrInvalidDenom, "stablecoin denom cannot be empty")
	}
	
	if msg.InterestRate.IsNegative() {
		return errors.Wrap(ErrInvalidRate, "interest rate cannot be negative")
	}
	
	if msg.ExpiryMonths <= 0 {
		return errors.Wrap(ErrInvalidTerm, "expiry months must be positive")
	}
	
	return nil
}

type MsgDrawFromCreditLine struct {
	Creator      string    `json:"creator"`
	CreditLineId string    `json:"credit_line_id"`
	Amount       sdk.Coins `json:"amount"`
	Purpose      string    `json:"purpose"`
}

func NewMsgDrawFromCreditLine(creator, creditLineId string, amount sdk.Coins, purpose string) *MsgDrawFromCreditLine {
	return &MsgDrawFromCreditLine{
		Creator:      creator,
		CreditLineId: creditLineId,
		Amount:       amount,
		Purpose:      purpose,
	}
}

func (msg *MsgDrawFromCreditLine) Route() string {
	return RouterKey
}

func (msg *MsgDrawFromCreditLine) Type() string {
	return TypeMsgDrawFromCreditLine
}

func (msg *MsgDrawFromCreditLine) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDrawFromCreditLine) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDrawFromCreditLine) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	
	if msg.CreditLineId == "" {
		return errors.Wrap(ErrInvalidRequest, "credit line ID cannot be empty")
	}
	
	if !msg.Amount.IsValid() || msg.Amount.IsZero() {
		return errors.Wrap(ErrInvalidAmount, "draw amount must be positive")
	}
	
	return nil
}