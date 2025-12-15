package keeper_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	feemarketkeeper "github.com/stateset/core/x/feemarket/keeper"
	feemarkettypes "github.com/stateset/core/x/feemarket/types"
)

func TestQueryBaseFee(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set base fee
	expectedBaseFee := sdkmath.LegacyMustNewDecFromStr("0.025")
	k.SetBaseFee(ctx, expectedBaseFee)

	// Query base fee
	resp, err := k.BaseFee(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryBaseFeeRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.True(t, expectedBaseFee.Equal(resp.BaseFee))
}

func TestQueryBaseFee_NilRequest(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.BaseFee(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidRequest)
}

func TestQueryParams(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set custom params
	customParams := feemarkettypes.Params{
		Enabled:                  true,
		BaseFeeChangeDenominator: 16,
		ElasticityMultiplier:     4,
		TargetGas:                20_000_000,
		InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.002"),
		MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.0002"),
		MaxBaseFee:               sdkmath.LegacyMustNewDecFromStr("2.0"),
		PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.00002"),
		MaxFeeHistory:            150,
	}
	k.SetParams(ctx, customParams)

	// Query params
	resp, err := k.Params(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryParamsRequest{})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(16), resp.Params.BaseFeeChangeDenominator)
	require.Equal(t, uint64(4), resp.Params.ElasticityMultiplier)
	require.Equal(t, uint64(20_000_000), resp.Params.TargetGas)
	require.Equal(t, uint32(150), resp.Params.MaxFeeHistory)
}

func TestQueryParams_NilRequest(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.Params(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidRequest)
}

func TestQueryGasPrice_Priorities(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set base fee
	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	tests := []struct {
		name           string
		priority       string
		minMultiplier  sdkmath.LegacyDec // Expected minimum multiplier
		maxMultiplier  sdkmath.LegacyDec // Expected maximum multiplier
	}{
		{
			name:          "low priority",
			priority:      "low",
			minMultiplier: sdkmath.LegacyOneDec(), // 1.0x
			maxMultiplier: sdkmath.LegacyMustNewDecFromStr("1.0"),
		},
		{
			name:          "medium priority",
			priority:      "medium",
			minMultiplier: sdkmath.LegacyMustNewDecFromStr("1.20"), // ~1.25x
			maxMultiplier: sdkmath.LegacyMustNewDecFromStr("1.30"),
		},
		{
			name:          "high priority",
			priority:      "high",
			minMultiplier: sdkmath.LegacyMustNewDecFromStr("1.45"), // ~1.5x
			maxMultiplier: sdkmath.LegacyMustNewDecFromStr("1.55"),
		},
		{
			name:          "default (standard)",
			priority:      "",
			minMultiplier: sdkmath.LegacyMustNewDecFromStr("1.05"), // ~1.1x
			maxMultiplier: sdkmath.LegacyMustNewDecFromStr("1.15"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := k.GasPrice(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryGasPriceRequest{
				Priority: tt.priority,
			})
			require.NoError(t, err)
			require.NotNil(t, resp)

			// Gas price should be base fee * multiplier
			expectedMin := baseFee.Mul(tt.minMultiplier)
			expectedMax := baseFee.Mul(tt.maxMultiplier)

			require.True(t, resp.GasPrice.GTE(expectedMin),
				"gas price %s should be >= %s for %s priority",
				resp.GasPrice, expectedMin, tt.priority)
			require.True(t, resp.GasPrice.LTE(expectedMax),
				"gas price %s should be <= %s for %s priority",
				resp.GasPrice, expectedMax, tt.priority)
		})
	}
}

func TestQueryGasPrice_NilRequest(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.GasPrice(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidRequest)
}

func TestQueryEstimateFee(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set base fee
	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	gasLimit := uint64(100000)

	resp, err := k.EstimateFee(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryEstimateFeeRequest{
		GasLimit: gasLimit,
		Priority: "medium",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Base fee component should be baseFee * gasLimit
	expectedBaseFeeComponent := baseFee.MulInt64(int64(gasLimit))
	require.True(t, resp.BaseFeeComponent.Equal(expectedBaseFeeComponent),
		"base fee component mismatch: got %s, want %s",
		resp.BaseFeeComponent, expectedBaseFeeComponent)

	// Total fee should include priority component
	require.True(t, resp.EstimatedFee.GT(resp.BaseFeeComponent),
		"total fee should be greater than base fee component")

	// Priority component should be positive for medium priority
	require.True(t, resp.PriorityFeeComponent.IsPositive(),
		"priority fee component should be positive for medium priority")

	// Verify: estimated = base + priority
	expectedTotal := resp.BaseFeeComponent.Add(resp.PriorityFeeComponent)
	require.True(t, resp.EstimatedFee.Equal(expectedTotal),
		"estimated fee should equal base + priority: %s != %s + %s",
		resp.EstimatedFee, resp.BaseFeeComponent, resp.PriorityFeeComponent)
}

func TestQueryEstimateFee_ZeroGasLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.EstimateFee(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryEstimateFeeRequest{
		GasLimit: 0,
		Priority: "medium",
	})
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidGasLimit)
}

func TestQueryEstimateFee_NilRequest(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.EstimateFee(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidRequest)
}

func TestQueryEstimateFee_LowPriority(t *testing.T) {
	k, ctx := setupKeeper(t)

	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	gasLimit := uint64(100000)

	resp, err := k.EstimateFee(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryEstimateFeeRequest{
		GasLimit: gasLimit,
		Priority: "low",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	// For low priority, the multiplier is 0%, so priority component should be 0
	require.True(t, resp.PriorityFeeComponent.IsZero(),
		"priority fee should be zero for low priority, got %s", resp.PriorityFeeComponent)
}

func TestQueryFeeHistory(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Query with no history
	resp, err := k.FeeHistory(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryFeeHistoryRequest{
		Limit: 10,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Empty(t, resp.FeeHistory)
}

func TestQueryFeeHistory_DefaultLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Query with zero limit (should use default)
	resp, err := k.FeeHistory(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryFeeHistoryRequest{
		Limit: 0,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestQueryFeeHistory_MaxLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Query with very high limit (should be capped at 100)
	resp, err := k.FeeHistory(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryFeeHistoryRequest{
		Limit: 500,
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestQueryFeeHistory_NilRequest(t *testing.T) {
	k, ctx := setupKeeper(t)

	resp, err := k.FeeHistory(sdk.WrapSDKContext(ctx), nil)
	require.Error(t, err)
	require.Nil(t, resp)
	require.ErrorIs(t, err, feemarkettypes.ErrInvalidRequest)
}

func TestGasOracleIntegration(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up realistic scenario
	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	k.SetParams(ctx, params)

	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	// Create gas oracle
	oracle := feemarketkeeper.NewGasOracle(k)

	// Test gas estimation
	gas := oracle.EstimateGas(ctx, "send")
	require.Equal(t, uint64(100000), gas)

	// Test fee estimation
	estimate := oracle.EstimateFee(ctx, gas, "high")
	require.Equal(t, gas, estimate.GasLimit)
	require.True(t, estimate.TotalFee.IsPositive())

	// Test congestion level
	level := oracle.GetCongestionLevel(ctx)
	require.Contains(t, []string{"low", "medium", "high", "unknown"}, level)
}

func TestPriorityMultiplier(t *testing.T) {
	k, ctx := setupKeeper(t)

	baseFee := sdkmath.LegacyMustNewDecFromStr("1.0") // Use 1.0 for easy math
	k.SetBaseFee(ctx, baseFee)

	tests := []struct {
		priority     string
		expectedMult sdkmath.LegacyDec
	}{
		{"low", sdkmath.LegacyZeroDec()},
		{"medium", sdkmath.LegacyMustNewDecFromStr("0.25")},
		{"high", sdkmath.LegacyMustNewDecFromStr("0.50")},
		{"", sdkmath.LegacyMustNewDecFromStr("0.10")},       // default
		{"invalid", sdkmath.LegacyMustNewDecFromStr("0.10")}, // falls to default
	}

	for _, tt := range tests {
		t.Run("priority_"+tt.priority, func(t *testing.T) {
			resp, err := k.GasPrice(sdk.WrapSDKContext(ctx), &feemarkettypes.QueryGasPriceRequest{
				Priority: tt.priority,
			})
			require.NoError(t, err)

			// Expected: baseFee * (1 + multiplier)
			expectedGasPrice := baseFee.Mul(sdkmath.LegacyOneDec().Add(tt.expectedMult))
			require.True(t, resp.GasPrice.Equal(expectedGasPrice),
				"priority %q: expected %s, got %s", tt.priority, expectedGasPrice, resp.GasPrice)
		})
	}
}
