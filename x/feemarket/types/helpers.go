package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// FeeEstimate represents a fee estimate for a transaction.
type FeeEstimate struct {
	GasPrice             sdkmath.LegacyDec `json:"gas_price"`
	GasLimit             uint64            `json:"gas_limit"`
	TotalFee             sdkmath.LegacyDec `json:"total_fee"`
	BaseFeeComponent     sdkmath.LegacyDec `json:"base_fee_component"`
	PriorityFeeComponent sdkmath.LegacyDec `json:"priority_fee_component"`
}

// GetFeeHistoryKey returns the store key for a fee history entry.
func GetFeeHistoryKey(height int64) []byte {
	return append(HistoryKey, sdk.Uint64ToBigEndian(uint64(height))...)
}

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
