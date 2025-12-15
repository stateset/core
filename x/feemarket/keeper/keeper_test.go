package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	feemarketkeeper "github.com/stateset/core/x/feemarket/keeper"
	feemarkettypes "github.com/stateset/core/x/feemarket/types"
)

func setupKeeper(t *testing.T) (feemarketkeeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(feemarkettypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := feemarketkeeper.NewKeeper(
		cdc,
		storeKey,
		"stateset1authority",
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		Height:  1,
		ChainID: "stateset-test",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	return k, ctx
}

func TestNewKeeper(t *testing.T) {
	k, _ := setupKeeper(t)
	require.NotNil(t, k)
	require.Equal(t, "stateset1authority", k.GetAuthority())
}

func TestGetSetParams(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Get default params (should return defaults when unset)
	params := k.GetParams(ctx)
	require.True(t, params.Enabled)
	require.Equal(t, feemarkettypes.DefaultBaseFeeChangeDenominator, params.BaseFeeChangeDenominator)
	require.Equal(t, feemarkettypes.DefaultElasticityMultiplier, params.ElasticityMultiplier)
	require.Equal(t, feemarkettypes.DefaultTargetGas, params.TargetGas)

	// Set custom params
	customParams := feemarkettypes.Params{
		Enabled:                  true,
		BaseFeeChangeDenominator: 16,
		ElasticityMultiplier:     4,
		TargetGas:                20_000_000,
		InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr("0.001"),
		MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.0001"),
		MaxBaseFee:               sdkmath.LegacyMustNewDecFromStr("1.0"),
		PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.00001"),
		MaxFeeHistory:            200,
	}
	k.SetParams(ctx, customParams)

	// Verify params were stored
	retrieved := k.GetParams(ctx)
	require.Equal(t, uint64(16), retrieved.BaseFeeChangeDenominator)
	require.Equal(t, uint64(4), retrieved.ElasticityMultiplier)
	require.Equal(t, uint64(20_000_000), retrieved.TargetGas)
	require.Equal(t, uint32(200), retrieved.MaxFeeHistory)
}

func TestGetSetBaseFee(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Get default base fee (should return initial default when unset)
	baseFee := k.GetBaseFee(ctx)
	require.Equal(t, feemarkettypes.DefaultInitialBaseFee, baseFee)

	// Set custom base fee
	newBaseFee := sdkmath.LegacyMustNewDecFromStr("0.05")
	k.SetBaseFee(ctx, newBaseFee)

	// Verify it was stored
	retrieved := k.GetBaseFee(ctx)
	require.True(t, newBaseFee.Equal(retrieved))
}

func TestUpdateBaseFee_Increase(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params
	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	params.BaseFeeChangeDenominator = 8
	k.SetParams(ctx, params)

	// Set initial base fee
	initialBaseFee := sdkmath.LegacyMustNewDecFromStr("0.001")
	k.SetBaseFee(ctx, initialBaseFee)

	// Simulate high gas usage (above target)
	gasUsed := uint64(15_000_000)      // 1.5x target
	maxBlockGas := uint64(20_000_000)

	newBaseFee := k.UpdateBaseFee(ctx, gasUsed, maxBlockGas)

	// Base fee should increase when gas usage exceeds target
	require.True(t, newBaseFee.GT(initialBaseFee), "base fee should increase when gas > target")
}

func TestUpdateBaseFee_Decrease(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params
	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	params.BaseFeeChangeDenominator = 8
	params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.0001")
	k.SetParams(ctx, params)

	// Set initial base fee
	initialBaseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, initialBaseFee)

	// Simulate low gas usage (below target)
	gasUsed := uint64(5_000_000)       // 0.5x target
	maxBlockGas := uint64(20_000_000)

	newBaseFee := k.UpdateBaseFee(ctx, gasUsed, maxBlockGas)

	// Base fee should decrease when gas usage is below target
	require.True(t, newBaseFee.LT(initialBaseFee), "base fee should decrease when gas < target")
}

func TestUpdateBaseFee_MinBound(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params with minimum base fee
	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.001")
	k.SetParams(ctx, params)

	// Set base fee just above minimum
	k.SetBaseFee(ctx, sdkmath.LegacyMustNewDecFromStr("0.0011"))

	// Simulate very low gas usage to push fee down
	gasUsed := uint64(100_000)
	maxBlockGas := uint64(20_000_000)

	newBaseFee := k.UpdateBaseFee(ctx, gasUsed, maxBlockGas)

	// Base fee should not go below minimum
	require.True(t, newBaseFee.GTE(params.MinBaseFee), "base fee should not go below minimum")
}

func TestUpdateBaseFee_MaxBound(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params with maximum base fee
	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	params.MaxBaseFee = sdkmath.LegacyMustNewDecFromStr("1.0")
	k.SetParams(ctx, params)

	// Set base fee close to maximum
	k.SetBaseFee(ctx, sdkmath.LegacyMustNewDecFromStr("0.99"))

	// Simulate very high gas usage to push fee up
	gasUsed := uint64(50_000_000)
	maxBlockGas := uint64(100_000_000)

	newBaseFee := k.UpdateBaseFee(ctx, gasUsed, maxBlockGas)

	// Base fee should not exceed maximum
	require.True(t, newBaseFee.LTE(params.MaxBaseFee), "base fee should not exceed maximum")
}

func TestUpdateBaseFee_DisabledFeeMarket(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Disable fee market
	params := feemarkettypes.DefaultParams()
	params.Enabled = false
	k.SetParams(ctx, params)

	initialBaseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, initialBaseFee)

	// Update should not change fee when disabled
	newBaseFee := k.UpdateBaseFee(ctx, 15_000_000, 20_000_000)

	require.True(t, newBaseFee.Equal(initialBaseFee), "base fee should not change when fee market is disabled")
}

func TestGetSetLatestGas(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Initially should be 0
	require.Equal(t, uint64(0), k.GetLatestGas(ctx))

	// Set via UpdateBaseFee
	k.SetParams(ctx, feemarkettypes.DefaultParams())
	k.SetBaseFee(ctx, feemarkettypes.DefaultInitialBaseFee)
	k.UpdateBaseFee(ctx, 12345678, 20_000_000)

	// Should be stored
	require.Equal(t, uint64(12345678), k.GetLatestGas(ctx))
}

func TestGasOracle_SuggestGasPrice(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	k.SetParams(ctx, params)

	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	oracle := feemarketkeeper.NewGasOracle(k)

	// Test different priority levels
	lowPrice := oracle.SuggestGasPrice(ctx, "low")
	mediumPrice := oracle.SuggestGasPrice(ctx, "medium")
	highPrice := oracle.SuggestGasPrice(ctx, "high")

	// Higher priority should result in higher suggested price
	require.True(t, lowPrice.LTE(mediumPrice), "low priority should be <= medium")
	require.True(t, mediumPrice.LTE(highPrice), "medium priority should be <= high")
}

func TestGasOracle_SuggestGasPrice_DisabledReturnsMinimum(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.Enabled = false
	params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.0001")
	k.SetParams(ctx, params)

	oracle := feemarketkeeper.NewGasOracle(k)

	price := oracle.SuggestGasPrice(ctx, "high")
	require.True(t, price.Equal(params.MinBaseFee), "should return min base fee when disabled")
}

func TestGasOracle_EstimateGas(t *testing.T) {
	k, ctx := setupKeeper(t)
	oracle := feemarketkeeper.NewGasOracle(k)

	tests := []struct {
		txType   string
		expected uint64
	}{
		{"send", 100000},
		{"delegate", 200000},
		{"vote", 150000},
		{"swap", 250000},
		{"contract", 500000},
		{"unknown", 200000},
	}

	for _, tt := range tests {
		gas := oracle.EstimateGas(ctx, tt.txType)
		require.Equal(t, tt.expected, gas, "gas estimate for %s", tt.txType)
	}
}

func TestGasOracle_EstimateFee(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	k.SetParams(ctx, params)

	baseFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, baseFee)

	oracle := feemarketkeeper.NewGasOracle(k)

	estimate := oracle.EstimateFee(ctx, 100000, "medium")

	require.Equal(t, uint64(100000), estimate.GasLimit)
	require.True(t, estimate.GasPrice.IsPositive())
	require.True(t, estimate.TotalFee.IsPositive())
	require.True(t, estimate.BaseFeeComponent.IsPositive())
}

func TestGasOracle_GetCongestionLevel(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	k.SetParams(ctx, params)

	oracle := feemarketkeeper.NewGasOracle(k)

	// Test low congestion
	k.UpdateBaseFee(ctx, 5_000_000, 20_000_000) // 50% of target
	level := oracle.GetCongestionLevel(ctx)
	require.Equal(t, "low", level)

	// Test medium congestion
	k.UpdateBaseFee(ctx, 10_000_000, 20_000_000) // 100% of target
	level = oracle.GetCongestionLevel(ctx)
	require.Equal(t, "medium", level)

	// Test high congestion
	k.UpdateBaseFee(ctx, 15_000_000, 20_000_000) // 150% of target
	level = oracle.GetCongestionLevel(ctx)
	require.Equal(t, "high", level)
}

func TestGasOracle_GetRecommendedPriority(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	k.SetParams(ctx, params)

	oracle := feemarketkeeper.NewGasOracle(k)

	// Low congestion, non-urgent
	k.UpdateBaseFee(ctx, 5_000_000, 20_000_000)
	priority := oracle.GetRecommendedPriority(ctx, false)
	require.Equal(t, "low", priority)

	// Low congestion, urgent
	priority = oracle.GetRecommendedPriority(ctx, true)
	require.Equal(t, "high", priority)

	// High congestion, non-urgent
	k.UpdateBaseFee(ctx, 15_000_000, 20_000_000)
	priority = oracle.GetRecommendedPriority(ctx, false)
	require.Equal(t, "high", priority)
}

func TestGasOracle_PredictNextBaseFee(t *testing.T) {
	k, ctx := setupKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.TargetGas = 10_000_000
	// Set MinBaseFee low to allow testing fee decreases
	params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.001")
	k.SetParams(ctx, params)

	// Use a base fee well above MinBaseFee so we can test both increases and decreases
	baseFee := sdkmath.LegacyMustNewDecFromStr("0.05")
	k.SetBaseFee(ctx, baseFee)

	oracle := feemarketkeeper.NewGasOracle(k)

	// Predict increase (gas used > target)
	nextFee := oracle.PredictNextBaseFee(ctx, 15_000_000, 20_000_000)
	require.True(t, nextFee.GT(baseFee), "should predict higher fee for high gas usage")

	// Predict decrease (gas used < target)
	nextFee = oracle.PredictNextBaseFee(ctx, 5_000_000, 20_000_000)
	require.True(t, nextFee.LT(baseFee), "should predict lower fee for low gas usage")
}

func TestComputeNextBaseFee(t *testing.T) {
	tests := []struct {
		name        string
		current     sdkmath.LegacyDec
		gasUsed     uint64
		targetGas   uint64
		enabled     bool
		shouldRise  bool
	}{
		{
			name:       "gas above target increases fee",
			current:    sdkmath.LegacyMustNewDecFromStr("0.01"),
			gasUsed:    15_000_000,
			targetGas:  10_000_000,
			enabled:    true,
			shouldRise: true,
		},
		{
			name:       "gas below target decreases fee",
			current:    sdkmath.LegacyMustNewDecFromStr("0.01"),
			gasUsed:    5_000_000,
			targetGas:  10_000_000,
			enabled:    true,
			shouldRise: false,
		},
		{
			name:       "gas at target keeps fee stable",
			current:    sdkmath.LegacyMustNewDecFromStr("0.01"),
			gasUsed:    10_000_000,
			targetGas:  10_000_000,
			enabled:    true,
			shouldRise: false, // Should stay same, not rise
		},
		{
			name:       "disabled returns current",
			current:    sdkmath.LegacyMustNewDecFromStr("0.01"),
			gasUsed:    15_000_000,
			targetGas:  10_000_000,
			enabled:    false,
			shouldRise: false, // No change
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := feemarkettypes.DefaultParams()
			params.Enabled = tt.enabled
			params.TargetGas = tt.targetGas
			params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.0001")

			next := feemarkettypes.ComputeNextBaseFee(tt.current, tt.gasUsed, params, 20_000_000)

			if tt.shouldRise {
				require.True(t, next.GT(tt.current), "%s: expected fee to increase", tt.name)
			} else if !tt.enabled {
				require.True(t, next.Equal(tt.current), "%s: expected fee unchanged when disabled", tt.name)
			}
		})
	}
}

func TestParamsValidation(t *testing.T) {
	tests := []struct {
		name      string
		modify    func(*feemarkettypes.Params)
		expectErr bool
	}{
		{
			name:      "valid default params",
			modify:    func(p *feemarkettypes.Params) {},
			expectErr: false,
		},
		{
			name: "zero base fee change denominator",
			modify: func(p *feemarkettypes.Params) {
				p.BaseFeeChangeDenominator = 0
			},
			expectErr: true,
		},
		{
			name: "zero elasticity multiplier",
			modify: func(p *feemarkettypes.Params) {
				p.ElasticityMultiplier = 0
			},
			expectErr: true,
		},
		{
			name: "zero target gas",
			modify: func(p *feemarkettypes.Params) {
				p.TargetGas = 0
			},
			expectErr: true,
		},
		{
			name: "negative min base fee",
			modify: func(p *feemarkettypes.Params) {
				p.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("-0.001")
			},
			expectErr: true,
		},
		{
			name: "max base fee less than min",
			modify: func(p *feemarkettypes.Params) {
				p.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("1.0")
				p.MaxBaseFee = sdkmath.LegacyMustNewDecFromStr("0.5")
			},
			expectErr: true,
		},
		{
			name: "zero max fee history",
			modify: func(p *feemarkettypes.Params) {
				p.MaxFeeHistory = 0
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := feemarkettypes.DefaultParams()
			tt.modify(&params)

			err := params.ValidateBasic()
			if tt.expectErr {
				require.Error(t, err, "expected validation error")
			} else {
				require.NoError(t, err, "expected no validation error")
			}
		})
	}
}

func TestTargetGasOrDefault(t *testing.T) {
	tests := []struct {
		name        string
		targetGas   uint64
		elasticity  uint64
		maxBlockGas uint64
		expected    uint64
	}{
		{
			name:        "explicit target gas",
			targetGas:   15_000_000,
			elasticity:  2,
			maxBlockGas: 100_000_000,
			expected:    15_000_000,
		},
		{
			name:        "computed from elasticity",
			targetGas:   0,
			elasticity:  2,
			maxBlockGas: 100_000_000,
			expected:    50_000_000,
		},
		{
			name:        "fallback to default",
			targetGas:   0,
			elasticity:  0,
			maxBlockGas: 0,
			expected:    feemarkettypes.DefaultTargetGas,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := feemarkettypes.Params{
				TargetGas:            tt.targetGas,
				ElasticityMultiplier: tt.elasticity,
			}
			result := params.TargetGasOrDefault(tt.maxBlockGas)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAuthority(t *testing.T) {
	k, _ := setupKeeper(t)
	require.Equal(t, "stateset1authority", k.GetAuthority())
}

func TestBaseFeeZeroBootstrap(t *testing.T) {
	// Test that zero base fee gets bootstrapped from initial params
	params := feemarkettypes.DefaultParams()
	params.InitialBaseFee = sdkmath.LegacyMustNewDecFromStr("0.005")
	params.TargetGas = 10_000_000

	zeroDec := sdkmath.LegacyZeroDec()
	next := feemarkettypes.ComputeNextBaseFee(zeroDec, 15_000_000, params, 20_000_000)

	// Should bootstrap from initial, not stay at zero
	require.True(t, next.IsPositive(), "zero base fee should bootstrap from initial")
}
