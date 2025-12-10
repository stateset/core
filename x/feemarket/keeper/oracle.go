package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/feemarket/types"
)

// GasOracle provides intelligent gas price and fee estimation capabilities.
type GasOracle struct {
	keeper Keeper
}

// NewGasOracle creates a new gas oracle instance.
func NewGasOracle(keeper Keeper) GasOracle {
	return GasOracle{keeper: keeper}
}

// SuggestGasPrice returns the recommended gas price for a transaction.
// It considers recent base fee trends, network congestion, and the requested priority.
func (o GasOracle) SuggestGasPrice(ctx sdk.Context, priority string) sdkmath.LegacyDec {
	baseFee := o.keeper.GetBaseFee(ctx)
	params := o.keeper.GetParams(ctx)

	if !params.Enabled {
		// If fee market is disabled, return minimum base fee
		return params.MinBaseFee
	}

	// Get congestion multiplier based on recent gas usage
	congestionMultiplier := o.calculateCongestionMultiplier(ctx)

	// Get priority multiplier
	priorityMultiplier := o.keeper.getPriorityMultiplier(priority)

	// Suggested gas price = baseFee * (1 + congestionMultiplier) * (1 + priorityMultiplier)
	suggestedPrice := baseFee.
		Mul(sdkmath.LegacyOneDec().Add(congestionMultiplier)).
		Mul(sdkmath.LegacyOneDec().Add(priorityMultiplier))

	// Ensure within bounds
	if suggestedPrice.LT(params.MinBaseFee) {
		return params.MinBaseFee
	}
	if suggestedPrice.GT(params.MaxBaseFee) {
		return params.MaxBaseFee
	}

	return suggestedPrice
}

// EstimateGas estimates the gas required for a transaction based on historical data.
// This is a simple estimation - actual gas usage may vary.
func (o GasOracle) EstimateGas(ctx sdk.Context, txType string) uint64 {
	// Default gas estimates for common transaction types
	switch txType {
	case "send":
		return 100000 // Simple bank send
	case "delegate":
		return 200000 // Staking delegation
	case "vote":
		return 150000 // Governance vote
	case "swap":
		return 250000 // Token swap
	case "contract":
		return 500000 // Smart contract execution
	default:
		return 200000 // Default estimate
	}
}

// EstimateFee provides a complete fee estimation including base fee and priority fee.
func (o GasOracle) EstimateFee(ctx sdk.Context, gasLimit uint64, priority string) types.FeeEstimate {
	gasPrice := o.SuggestGasPrice(ctx, priority)
	baseFee := o.keeper.GetBaseFee(ctx)

	// Calculate components
	baseFeeComponent := baseFee.MulInt64(int64(gasLimit))
	totalFee := gasPrice.MulInt64(int64(gasLimit))
	priorityFeeComponent := totalFee.Sub(baseFeeComponent)

	return types.FeeEstimate{
		GasPrice:             gasPrice,
		GasLimit:             gasLimit,
		TotalFee:             totalFee,
		BaseFeeComponent:     baseFeeComponent,
		PriorityFeeComponent: priorityFeeComponent,
	}
}

// GetHistoricalFees returns historical fee data for trend analysis.
func (o GasOracle) GetHistoricalFees(ctx sdk.Context, blocks uint64) []types.FeeHistoryEntry {
	if blocks == 0 {
		blocks = 20 // Default to last 20 blocks
	}
	if blocks > 100 {
		blocks = 100 // Cap at 100 blocks
	}

	return o.keeper.GetFeeHistory(ctx, blocks)
}

// calculateCongestionMultiplier calculates a multiplier based on network congestion.
// Returns 0.0 for low congestion, up to 0.5 for high congestion.
func (o GasOracle) calculateCongestionMultiplier(ctx sdk.Context) sdkmath.LegacyDec {
	params := o.keeper.GetParams(ctx)
	latestGas := o.keeper.GetLatestGas(ctx)
	targetGas := params.TargetGas

	if targetGas == 0 || latestGas == 0 {
		return sdkmath.LegacyZeroDec()
	}

	// Calculate congestion ratio: actual gas / target gas
	gasRatio := sdkmath.LegacyNewDec(int64(latestGas)).
		Quo(sdkmath.LegacyNewDec(int64(targetGas)))

	// If below target, no congestion multiplier
	if gasRatio.LTE(sdkmath.LegacyOneDec()) {
		return sdkmath.LegacyZeroDec()
	}

	// Calculate congestion multiplier (capped at 0.5 = 50% increase)
	// Formula: min(0.5, (gasRatio - 1.0) * 0.5)
	congestionMultiplier := gasRatio.Sub(sdkmath.LegacyOneDec()).
		Mul(sdkmath.LegacyMustNewDecFromStr("0.5"))

	maxCongestionMultiplier := sdkmath.LegacyMustNewDecFromStr("0.5")
	if congestionMultiplier.GT(maxCongestionMultiplier) {
		return maxCongestionMultiplier
	}

	return congestionMultiplier
}

// PredictNextBaseFee predicts what the base fee will be in the next block
// based on current gas usage trends.
func (o GasOracle) PredictNextBaseFee(ctx sdk.Context, estimatedGasUsed uint64, maxBlockGas uint64) sdkmath.LegacyDec {
	params := o.keeper.GetParams(ctx)
	currentBaseFee := o.keeper.GetBaseFee(ctx)

	return types.ComputeNextBaseFee(currentBaseFee, estimatedGasUsed, params, maxBlockGas)
}

// GetCongestionLevel returns a human-readable congestion level (low, medium, high).
func (o GasOracle) GetCongestionLevel(ctx sdk.Context) string {
	params := o.keeper.GetParams(ctx)
	latestGas := o.keeper.GetLatestGas(ctx)
	targetGas := params.TargetGas

	if targetGas == 0 {
		return "unknown"
	}

	gasRatio := float64(latestGas) / float64(targetGas)

	switch {
	case gasRatio < 0.7:
		return "low"
	case gasRatio < 1.2:
		return "medium"
	default:
		return "high"
	}
}

// GetRecommendedPriority suggests a priority level based on current network conditions.
func (o GasOracle) GetRecommendedPriority(ctx sdk.Context, urgent bool) string {
	congestionLevel := o.GetCongestionLevel(ctx)

	if urgent {
		// For urgent transactions, always recommend high priority
		return "high"
	}

	// For non-urgent transactions, adjust based on congestion
	switch congestionLevel {
	case "low":
		return "low"
	case "medium":
		return "medium"
	default:
		return "high"
	}
}
