package types

import (
	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
)

const (
	DefaultBaseFeeChangeDenominator uint64 = 8
	DefaultElasticityMultiplier     uint64 = 2
	DefaultTargetGas                uint64 = 10_000_000
	DefaultMaxFeeHistory            uint32 = 100
)

var DefaultPriorityFloor = sdkmath.LegacyMustNewDecFromStr("0.0001") // suggested minimum priority tip

// DefaultParams returns the default parameters for the fee market.
func DefaultParams() Params {
	return Params{
		Enabled:                  true,
		BaseFeeChangeDenominator: DefaultBaseFeeChangeDenominator,
		ElasticityMultiplier:     DefaultElasticityMultiplier,
		TargetGas:                DefaultTargetGas,
		InitialBaseFee:           DefaultInitialBaseFee,
		MinBaseFee:               DefaultMinBaseFee,
		MaxBaseFee:               sdkmath.LegacyZeroDec(),
		PriorityFeeFloor:         DefaultPriorityFloor,
		MaxFeeHistory:            DefaultMaxFeeHistory,
	}
}

// Validate performs validation of parameters.
func (p Params) Validate() error {
	return p.ValidateBasic()
}

// ValidateBasic performs basic validation of parameters.
func (p Params) ValidateBasic() error {
	if p.BaseFeeChangeDenominator == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "base_fee_change_denominator must be > 0")
	}
	if p.ElasticityMultiplier == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "elasticity_multiplier must be > 0")
	}
	if p.TargetGas == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "target_gas must be > 0")
	}
	if p.InitialBaseFee.IsNegative() || p.MinBaseFee.IsNegative() || p.PriorityFeeFloor.IsNegative() {
		return errorsmod.Wrap(ErrInvalidParams, "fee values cannot be negative")
	}
	if !p.MaxBaseFee.IsZero() && p.MaxBaseFee.LT(p.MinBaseFee) {
		return errorsmod.Wrap(ErrInvalidParams, "max_base_fee must be >= min_base_fee or zero")
	}
	if p.MaxFeeHistory == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "max_fee_history must be > 0")
	}
	return nil
}

// TargetGasOrDefault returns the explicit target or a computed default.
func (p Params) TargetGasOrDefault(maxBlockGas uint64) uint64 {
	if p.TargetGas > 0 {
		return p.TargetGas
	}
	if p.ElasticityMultiplier > 0 && maxBlockGas > 0 {
		return maxBlockGas / p.ElasticityMultiplier
	}
	return DefaultTargetGas
}
