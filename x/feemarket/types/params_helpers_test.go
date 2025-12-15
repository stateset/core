package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/feemarket/types"
)

func TestDefaultParams(t *testing.T) {
	params := types.DefaultParams()

	require.True(t, params.Enabled)
	require.Equal(t, types.DefaultBaseFeeChangeDenominator, params.BaseFeeChangeDenominator)
	require.Equal(t, types.DefaultElasticityMultiplier, params.ElasticityMultiplier)
	require.Equal(t, types.DefaultTargetGas, params.TargetGas)
	require.Equal(t, types.DefaultMinBaseFee, params.MinBaseFee)
	require.Equal(t, types.DefaultMaxFeeHistory, params.MaxFeeHistory)
	require.False(t, params.InitialBaseFee.IsNegative())
	require.False(t, params.PriorityFeeFloor.IsNegative())
}

func TestParams_ValidateBasic(t *testing.T) {
	validParams := types.DefaultParams()

	tests := []struct {
		name      string
		params    types.Params
		expectErr bool
	}{
		{
			name:      "valid default params",
			params:    validParams,
			expectErr: false,
		},
		{
			name: "zero base fee change denominator",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 0,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "zero elasticity multiplier",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     0,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "zero target gas",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                0,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "negative initial base fee",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("-0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "negative min base fee",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("-0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "negative priority fee floor",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("-0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "max base fee less than min base fee",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.1"),
				MaxBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.05"), // less than min
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: true,
		},
		{
			name: "max base fee equals zero is valid (no max)",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: false,
		},
		{
			name: "zero max fee history",
			params: types.Params{
				Enabled:                  true,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            0,
			},
			expectErr: true,
		},
		{
			name: "disabled feemarket is valid",
			params: types.Params{
				Enabled:                  false,
				BaseFeeChangeDenominator: 8,
				ElasticityMultiplier:     2,
				TargetGas:                10_000_000,
				InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.025"),
				MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.025"),
				MaxBaseFee:               sdkmath.LegacyZeroDec(),
				PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
				MaxFeeHistory:            100,
			},
			expectErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestParams_Validate(t *testing.T) {
	// Validate is an alias for ValidateBasic
	params := types.DefaultParams()
	require.NoError(t, params.Validate())

	invalidParams := types.Params{
		BaseFeeChangeDenominator: 0,
	}
	require.Error(t, invalidParams.Validate())
}

func TestParams_TargetGasOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		params      types.Params
		maxBlockGas uint64
		expected    uint64
	}{
		{
			name: "explicit target gas",
			params: types.Params{
				TargetGas:            15_000_000,
				ElasticityMultiplier: 2,
			},
			maxBlockGas: 100_000_000,
			expected:    15_000_000,
		},
		{
			name: "computed from max block gas",
			params: types.Params{
				TargetGas:            0,
				ElasticityMultiplier: 2,
			},
			maxBlockGas: 100_000_000,
			expected:    50_000_000, // maxBlockGas / elasticityMultiplier
		},
		{
			name: "computed with different multiplier",
			params: types.Params{
				TargetGas:            0,
				ElasticityMultiplier: 4,
			},
			maxBlockGas: 100_000_000,
			expected:    25_000_000,
		},
		{
			name: "fallback to default when no max block gas",
			params: types.Params{
				TargetGas:            0,
				ElasticityMultiplier: 2,
			},
			maxBlockGas: 0,
			expected:    types.DefaultTargetGas,
		},
		{
			name: "fallback to default when no multiplier",
			params: types.Params{
				TargetGas:            0,
				ElasticityMultiplier: 0,
			},
			maxBlockGas: 100_000_000,
			expected:    types.DefaultTargetGas,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.params.TargetGasOrDefault(tc.maxBlockGas)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	require.Equal(t, uint64(8), types.DefaultBaseFeeChangeDenominator)
	require.Equal(t, uint64(2), types.DefaultElasticityMultiplier)
	require.Equal(t, uint64(10_000_000), types.DefaultTargetGas)
	require.Equal(t, uint32(100), types.DefaultMaxFeeHistory)
	require.Equal(t, sdkmath.LegacyMustNewDecFromStr("0.0001"), types.DefaultPriorityFloor)
}
