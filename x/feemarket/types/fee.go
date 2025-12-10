package types

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FeeEstimate represents a fee estimate for a transaction
type FeeEstimate struct {
	GasPrice             sdkmath.LegacyDec `json:"gas_price"`
	GasLimit             uint64            `json:"gas_limit"`
	TotalFee             sdkmath.LegacyDec `json:"total_fee"`
	BaseFeeComponent     sdkmath.LegacyDec `json:"base_fee_component"`
	PriorityFeeComponent sdkmath.LegacyDec `json:"priority_fee_component"`
}

// GetFeeHistoryKey returns the store key for a fee history entry
func GetFeeHistoryKey(height int64) []byte {
	return append(HistoryKey, sdk.Uint64ToBigEndian(uint64(height))...)
}

// MsgServer defines the interface for fee market message server
type MsgServer interface {
	UpdateParams(ctx context.Context, msg *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
}

// MsgUpdateParams updates the fee market parameters
type MsgUpdateParams struct {
	Authority string `json:"authority"`
	Params    Params `json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return "MsgUpdateParams" }
func (*MsgUpdateParams) ProtoMessage()    {}
func (m MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Authority)
	return []sdk.AccAddress{addr}
}
func (m MsgUpdateParams) ValidateBasic() error {
	return m.Params.ValidateBasic()
}

type MsgUpdateParamsResponse struct{}

func (*MsgUpdateParamsResponse) Reset()         {}
func (*MsgUpdateParamsResponse) String() string { return "MsgUpdateParamsResponse" }
func (*MsgUpdateParamsResponse) ProtoMessage()  {}

// QueryServer defines the interface for fee market queries
type QueryServer interface {
	BaseFee(ctx context.Context, req *QueryBaseFeeRequest) (*QueryBaseFeeResponse, error)
	Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error)
	GasPrice(ctx context.Context, req *QueryGasPriceRequest) (*QueryGasPriceResponse, error)
	EstimateFee(ctx context.Context, req *QueryEstimateFeeRequest) (*QueryEstimateFeeResponse, error)
	FeeHistory(ctx context.Context, req *QueryFeeHistoryRequest) (*QueryFeeHistoryResponse, error)
}

// QueryBaseFeeRequest is the request type for querying the base fee
type QueryBaseFeeRequest struct{}

func (*QueryBaseFeeRequest) Reset()         {}
func (*QueryBaseFeeRequest) String() string { return "QueryBaseFeeRequest" }
func (*QueryBaseFeeRequest) ProtoMessage()  {}

// QueryBaseFeeResponse is the response type for querying the base fee
type QueryBaseFeeResponse struct {
	BaseFee sdkmath.LegacyDec `json:"base_fee"`
}

func (*QueryBaseFeeResponse) Reset()         {}
func (*QueryBaseFeeResponse) String() string { return "QueryBaseFeeResponse" }
func (*QueryBaseFeeResponse) ProtoMessage()  {}

// QueryParamsRequest is the request type for querying parameters
type QueryParamsRequest struct{}

func (*QueryParamsRequest) Reset()         {}
func (*QueryParamsRequest) String() string { return "QueryParamsRequest" }
func (*QueryParamsRequest) ProtoMessage()  {}

// QueryParamsResponse is the response type for querying parameters
type QueryParamsResponse struct {
	Params Params `json:"params"`
}

func (*QueryParamsResponse) Reset()         {}
func (*QueryParamsResponse) String() string { return "QueryParamsResponse" }
func (*QueryParamsResponse) ProtoMessage()  {}

// QueryGasPriceRequest is the request type for gas price query
type QueryGasPriceRequest struct {
	Priority string `json:"priority"`
}
func (*QueryGasPriceRequest) Reset()         {}
func (*QueryGasPriceRequest) String() string { return "QueryGasPriceRequest" }
func (*QueryGasPriceRequest) ProtoMessage()  {}

// QueryGasPriceResponse is the response type for gas price query
type QueryGasPriceResponse struct {
	GasPrice sdkmath.LegacyDec `json:"gas_price"`
}
func (*QueryGasPriceResponse) Reset()         {}
func (*QueryGasPriceResponse) String() string { return "QueryGasPriceResponse" }
func (*QueryGasPriceResponse) ProtoMessage()  {}

// QueryEstimateFeeRequest is the request type for fee estimation
type QueryEstimateFeeRequest struct {
	GasLimit uint64 `json:"gas_limit"`
	Priority string `json:"priority"`
}
func (*QueryEstimateFeeRequest) Reset()         {}
func (*QueryEstimateFeeRequest) String() string { return "QueryEstimateFeeRequest" }
func (*QueryEstimateFeeRequest) ProtoMessage()  {}

// QueryEstimateFeeResponse is the response type for fee estimation
type QueryEstimateFeeResponse struct {
	EstimatedFee         sdkmath.LegacyDec `json:"estimated_fee"`
	BaseFeeComponent     sdkmath.LegacyDec `json:"base_fee_component"`
	PriorityFeeComponent sdkmath.LegacyDec `json:"priority_fee_component"`
}
func (*QueryEstimateFeeResponse) Reset()         {}
func (*QueryEstimateFeeResponse) String() string { return "QueryEstimateFeeResponse" }
func (*QueryEstimateFeeResponse) ProtoMessage()  {}

// QueryFeeHistoryRequest is the request type for fee history
type QueryFeeHistoryRequest struct {
	Limit uint64 `json:"limit"`
}
func (*QueryFeeHistoryRequest) Reset()         {}
func (*QueryFeeHistoryRequest) String() string { return "QueryFeeHistoryRequest" }
func (*QueryFeeHistoryRequest) ProtoMessage()  {}

// QueryFeeHistoryResponse is the response type for fee history
type QueryFeeHistoryResponse struct {
	FeeHistory []FeeHistoryEntry `json:"fee_history"`
}
func (*QueryFeeHistoryResponse) Reset()         {}
func (*QueryFeeHistoryResponse) String() string { return "QueryFeeHistoryResponse" }
func (*QueryFeeHistoryResponse) ProtoMessage()  {}

// ComputeNextBaseFee applies an EIP-1559-style adjustment based on gas used vs. target gas.
func ComputeNextBaseFee(current sdkmath.LegacyDec, gasUsed uint64, params Params, maxBlockGas uint64) sdkmath.LegacyDec {
	if !params.Enabled {
		return current
	}

	target := params.TargetGasOrDefault(maxBlockGas)
	if target == 0 {
		return current
	}

	// guard against zero current base fee: bootstrap from params.InitialBaseFee
	if current.IsZero() {
		current = params.InitialBaseFee
	}

	gasUsedDec := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(gasUsed))
	targetDec := sdkmath.LegacyNewDecFromInt(sdkmath.NewIntFromUint64(target))

	delta := gasUsedDec.Sub(targetDec)

	// change = current * delta / target / denominator
	change := current.Mul(delta).Quo(targetDec.MulInt64(int64(params.BaseFeeChangeDenominator)))
	next := current.Add(change)

	if next.IsNegative() {
		next = sdkmath.LegacyZeroDec()
	}
	if next.LT(params.MinBaseFee) {
		next = params.MinBaseFee
	}
	if !params.MaxBaseFee.IsZero() && next.GT(params.MaxBaseFee) {
		next = params.MaxBaseFee
	}

	return next
}
