package types

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCommerceTransaction{}, "commerce/CreateCommerceTransaction", nil)
	cdc.RegisterConcrete(&MsgProcessPayment{}, "commerce/ProcessPayment", nil)
	cdc.RegisterConcrete(&MsgCreateTradeFinanceInstrument{}, "commerce/CreateTradeFinanceInstrument", nil)
	cdc.RegisterConcrete(&MsgOptimizePaymentRoute{}, "commerce/OptimizePaymentRoute", nil)
	cdc.RegisterConcrete(&MsgRunComplianceCheck{}, "commerce/RunComplianceCheck", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCommerceTransaction{},
		&MsgProcessPayment{},
		&MsgCreateTradeFinanceInstrument{},
		&MsgOptimizePaymentRoute{},
		&MsgRunComplianceCheck{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

// Message types for the commerce module
type MsgCreateCommerceTransaction struct {
	Creator     string             `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Transaction CommerceTransaction `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

type MsgProcessPayment struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TransactionID string `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

type MsgCreateTradeFinanceInstrument struct {
	Creator    string              `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Instrument FinancialInstrument `protobuf:"bytes,2,opt,name=instrument,proto3" json:"instrument,omitempty"`
}

type MsgOptimizePaymentRoute struct {
	Creator     string      `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PaymentInfo PaymentInfo `protobuf:"bytes,2,opt,name=paymentInfo,proto3" json:"paymentInfo,omitempty"`
}

type MsgRunComplianceCheck struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	TransactionID string `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

// Response types
type MsgCreateCommerceTransactionResponse struct {
	TransactionID string `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
}

type MsgProcessPaymentResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TxHash  string `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
}

type MsgCreateTradeFinanceInstrumentResponse struct {
	InstrumentID string `protobuf:"bytes,1,opt,name=instrumentId,proto3" json:"instrumentId,omitempty"`
}

type MsgOptimizePaymentRouteResponse struct {
	OptimalRoute PaymentRoute `protobuf:"bytes,1,opt,name=optimalRoute,proto3" json:"optimalRoute,omitempty"`
}

type MsgRunComplianceCheckResponse struct {
	ComplianceScore int32  `protobuf:"varint,1,opt,name=complianceScore,proto3" json:"complianceScore,omitempty"`
	Passed          bool   `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Issues          string `protobuf:"bytes,3,opt,name=issues,proto3" json:"issues,omitempty"`
}

// Implement sdk.Msg interface for all message types
var _ sdk.Msg = &MsgCreateCommerceTransaction{}
var _ sdk.Msg = &MsgProcessPayment{}
var _ sdk.Msg = &MsgCreateTradeFinanceInstrument{}
var _ sdk.Msg = &MsgOptimizePaymentRoute{}
var _ sdk.Msg = &MsgRunComplianceCheck{}

// Route implementations
func (msg *MsgCreateCommerceTransaction) Route() string {
	return RouterKey
}

func (msg *MsgProcessPayment) Route() string {
	return RouterKey
}

func (msg *MsgCreateTradeFinanceInstrument) Route() string {
	return RouterKey
}

func (msg *MsgOptimizePaymentRoute) Route() string {
	return RouterKey
}

func (msg *MsgRunComplianceCheck) Route() string {
	return RouterKey
}

// Type implementations
func (msg *MsgCreateCommerceTransaction) Type() string {
	return "CreateCommerceTransaction"
}

func (msg *MsgProcessPayment) Type() string {
	return "ProcessPayment"
}

func (msg *MsgCreateTradeFinanceInstrument) Type() string {
	return "CreateTradeFinanceInstrument"
}

func (msg *MsgOptimizePaymentRoute) Type() string {
	return "OptimizePaymentRoute"
}

func (msg *MsgRunComplianceCheck) Type() string {
	return "RunComplianceCheck"
}

// GetSigners implementations
func (msg *MsgCreateCommerceTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProcessPayment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTradeFinanceInstrument) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgOptimizePaymentRoute) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRunComplianceCheck) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implementations
func (msg *MsgCreateCommerceTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProcessPayment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTradeFinanceInstrument) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgOptimizePaymentRoute) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRunComplianceCheck) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implementations
func (msg *MsgCreateCommerceTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	return msg.Transaction.Validate()
}

func (msg *MsgProcessPayment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if msg.TransactionID == "" {
		return fmt.Errorf("transaction ID cannot be empty")
	}
	return nil
}

func (msg *MsgCreateTradeFinanceInstrument) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if msg.Instrument.ID == "" {
		return fmt.Errorf("instrument ID cannot be empty")
	}
	return nil
}

func (msg *MsgOptimizePaymentRoute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if msg.PaymentInfo.Amount.Empty() {
		return fmt.Errorf("payment amount cannot be empty")
	}
	return nil
}

func (msg *MsgRunComplianceCheck) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return err
	}
	if msg.TransactionID == "" {
		return fmt.Errorf("transaction ID cannot be empty")
	}
	return nil
}

// Service descriptor for gRPC
var _Msg_serviceDesc = struct {
	ServiceName string
}{
	ServiceName: "stateset.commerce.v1.Msg",
}