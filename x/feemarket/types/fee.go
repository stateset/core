package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ComputeNextBaseFee applies an EIP-1559-style adjustment based on gas used vs. target gas.
func ComputeNextBaseFee(current sdk.Dec, gasUsed uint64, params Params, maxBlockGas uint64) sdk.Dec {
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

	gasUsedDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(gasUsed))
	targetDec := sdk.NewDecFromInt(sdk.NewIntFromUint64(target))

	delta := gasUsedDec.Sub(targetDec)

	// change = current * delta / target / denominator
	change := current.Mul(delta).Quo(targetDec.MulInt64(int64(params.BaseFeeChangeDenominator)))
	next := current.Add(change)

	if next.IsNegative() {
		next = sdk.ZeroDec()
	}
	if next.LT(params.MinBaseFee) {
		next = params.MinBaseFee
	}
	if !params.MaxBaseFee.IsZero() && next.GT(params.MaxBaseFee) {
		next = params.MaxBaseFee
	}

	return next
}
