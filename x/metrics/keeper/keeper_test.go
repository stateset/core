package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/metrics/keeper"
	"github.com/stateset/core/x/metrics/types"
)

func setupKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	t.Helper()

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(cdc, storeKey)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	return k, ctx
}

func TestSystemMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Get default metrics
	metricsData := k.GetSystemMetrics(ctx)
	require.NotNil(t, metricsData)

	// Update metrics
	metricsData.TotalSettlements = 100
	metricsData.TotalSettlementVolume = sdkmath.NewInt(1000000)
	k.SetSystemMetrics(ctx, metricsData)

	// Verify update
	retrieved := k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(100), retrieved.TotalSettlements)
	require.Equal(t, sdkmath.NewInt(1000000), retrieved.TotalSettlementVolume)
}

func TestCounters(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Test basic counter
	counter := k.GetCounter(ctx, "test_counter", nil)
	require.Equal(t, uint64(0), counter.Value)

	// Increment counter
	k.IncrementCounter(ctx, "test_counter", nil)
	counter = k.GetCounter(ctx, "test_counter", nil)
	require.Equal(t, uint64(1), counter.Value)

	// Increment multiple times
	k.IncrementCounter(ctx, "test_counter", nil)
	k.IncrementCounter(ctx, "test_counter", nil)
	counter = k.GetCounter(ctx, "test_counter", nil)
	require.Equal(t, uint64(3), counter.Value)

	// Add to counter
	k.AddToCounter(ctx, "test_counter", nil, 10)
	counter = k.GetCounter(ctx, "test_counter", nil)
	require.Equal(t, uint64(13), counter.Value)
}

func TestCountersWithLabels(t *testing.T) {
	k, ctx := setupKeeper(t)

	labels1 := map[string]string{"module": "settlement"}
	labels2 := map[string]string{"module": "stablecoin"}

	// Increment different labeled counters
	k.IncrementCounter(ctx, "transactions", labels1)
	k.IncrementCounter(ctx, "transactions", labels1)
	k.IncrementCounter(ctx, "transactions", labels2)

	// Verify they're tracked separately
	counter1 := k.GetCounter(ctx, "transactions", labels1)
	require.Equal(t, uint64(2), counter1.Value)

	counter2 := k.GetCounter(ctx, "transactions", labels2)
	require.Equal(t, uint64(1), counter2.Value)
}

func TestGauges(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Test basic gauge
	gauge := k.GetGauge(ctx, "test_gauge", nil)
	require.True(t, gauge.Value.IsZero())

	// Update gauge
	k.UpdateGauge(ctx, "test_gauge", nil, sdkmath.LegacyNewDec(150))
	gauge = k.GetGauge(ctx, "test_gauge", nil)
	require.Equal(t, sdkmath.LegacyNewDec(150), gauge.Value)

	// Update to different value
	k.UpdateGauge(ctx, "test_gauge", nil, sdkmath.LegacyNewDec(75))
	gauge = k.GetGauge(ctx, "test_gauge", nil)
	require.Equal(t, sdkmath.LegacyNewDec(75), gauge.Value)
}

func TestGaugesWithLabels(t *testing.T) {
	k, ctx := setupKeeper(t)

	labels := map[string]string{"denom": "ssusd"}

	k.UpdateGauge(ctx, "collateral_ratio", labels, sdkmath.LegacyNewDecWithPrec(150, 2))
	gauge := k.GetGauge(ctx, "collateral_ratio", labels)
	require.Equal(t, sdkmath.LegacyNewDecWithPrec(150, 2), gauge.Value)
}

func TestModuleHealth(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record module health
	k.UpdateModuleHealth(ctx, "settlement", "healthy", 0.01, 50.0)

	metricsData := k.GetSystemMetrics(ctx)
	require.NotNil(t, metricsData.ModuleHealth)

	health, ok := metricsData.ModuleHealth["settlement"]
	require.True(t, ok)
	require.NotNil(t, health)
	require.Equal(t, "healthy", health.Status)
	require.Equal(t, 0.01, health.ErrorRate)
	require.Equal(t, 50.0, health.Latency)
}

func TestRecordModuleError(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record multiple transactions to establish baseline
	for i := 0; i < 10; i++ {
		k.UpdateModuleHealth(ctx, "payments", "healthy", 0.0, 10.0)
	}

	// Record an error
	k.RecordModuleError(ctx, "payments", "connection timeout")

	metricsData := k.GetSystemMetrics(ctx)
	health := metricsData.ModuleHealth["payments"]
	require.NotNil(t, health)

	require.Equal(t, "connection timeout", health.LastError)
	require.False(t, health.LastErrorAt.IsZero())
}

func TestBlockMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Update block metrics
	k.UpdateBlockMetrics(ctx)

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, int64(1), metricsData.LastBlockHeight)
	require.False(t, metricsData.LastBlockTime.IsZero())

	// Simulate next block
	ctx = ctx.WithBlockHeight(2).WithBlockTime(ctx.BlockTime().Add(6 * time.Second))
	k.UpdateBlockMetrics(ctx)

	metricsData = k.GetSystemMetrics(ctx)
	require.Equal(t, int64(2), metricsData.LastBlockHeight)
	require.True(t, metricsData.AverageBlockTime > 0)
}

func TestCollateralMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	totalCollateral := sdkmath.NewInt(1000000)
	totalDebt := sdkmath.NewInt(500000)

	k.UpdateCollateralMetrics(ctx, totalCollateral, totalDebt)

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, totalCollateral, metricsData.TotalCollateralValue)
	require.Equal(t, totalDebt, metricsData.TotalDebtValue)

	// Ratio should be 2.0 (1000000 / 500000)
	expectedRatio := sdkmath.LegacyNewDec(2)
	require.Equal(t, expectedRatio, metricsData.SystemCollateralRatio)
}

func TestCollateralMetricsZeroDebt(t *testing.T) {
	k, ctx := setupKeeper(t)

	totalCollateral := sdkmath.NewInt(1000000)
	totalDebt := sdkmath.ZeroInt()

	k.UpdateCollateralMetrics(ctx, totalCollateral, totalDebt)

	metricsData := k.GetSystemMetrics(ctx)
	require.True(t, metricsData.SystemCollateralRatio.IsZero())
}

func TestSettlementMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record settlements
	k.RecordSettlement(ctx, sdkmath.NewInt(100000))
	k.RecordSettlement(ctx, sdkmath.NewInt(200000))
	k.RecordSettlement(ctx, sdkmath.NewInt(150000))

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(3), metricsData.TotalSettlements)
	require.Equal(t, sdkmath.NewInt(450000), metricsData.TotalSettlementVolume)
}

func TestActiveEscrowsAndChannels(t *testing.T) {
	k, ctx := setupKeeper(t)

	k.UpdateActiveEscrows(ctx, 25)
	k.UpdateActiveChannels(ctx, 10)

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(25), metricsData.ActiveEscrows)
	require.Equal(t, uint64(10), metricsData.ActiveChannels)
}

func TestSecurityMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record circuit trips
	k.RecordCircuitTrip(ctx)
	k.RecordCircuitTrip(ctx)

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(2), metricsData.CircuitTrips)

	// Record rate limit hits
	k.RecordRateLimitHit(ctx, "mint_limit")
	k.RecordRateLimitHit(ctx, "mint_limit")
	k.RecordRateLimitHit(ctx, "transfer_limit")

	metricsData = k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(3), metricsData.RateLimitHits)

	// Record compliance blocks
	k.RecordComplianceBlock(ctx, "sanctioned")

	metricsData = k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(1), metricsData.ComplianceBlocks)
}

func TestOracleMetrics(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record price updates
	k.RecordPriceUpdate(ctx, "ssusd")
	k.RecordPriceUpdate(ctx, "stst")
	k.RecordPriceUpdate(ctx, "ssusd")

	metricsData := k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(3), metricsData.PricesUpdated)

	// Update stale price count
	k.UpdateStalePriceCount(ctx, 2)

	metricsData = k.GetSystemMetrics(ctx)
	require.Equal(t, uint64(2), metricsData.StalePriceCount)
}

func TestAlertConfigs(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set custom alert config
	config := types.AlertConfig{
		Name:       "low_collateral",
		MetricName: "system_collateral_ratio",
		Condition:  types.AlertConditionLessThan,
		Threshold:  sdkmath.LegacyNewDecWithPrec(15, 1), // 1.5
		Severity:   types.AlertSeverityCritical,
		Enabled:    true,
	}
	k.SetAlertConfig(ctx, config)

	// Get configs
	configs := k.GetAlertConfigs(ctx)
	require.NotEmpty(t, configs)

	// Find our config
	found := false
	for _, c := range configs {
		if c.Name == "low_collateral" {
			found = true
			require.Equal(t, types.AlertSeverityCritical, c.Severity)
		}
	}
	require.True(t, found)
}

func TestCheckAlerts(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up alert config
	config := types.AlertConfig{
		Name:       "test_alert",
		MetricName: "circuit_trips",
		Condition:  types.AlertConditionGreaterThan,
		Threshold:  sdkmath.LegacyNewDec(5),
		Severity:   types.AlertSeverityWarning,
		Enabled:    true,
	}
	k.SetAlertConfig(ctx, config)

	// Initially no alerts should trigger
	alerts := k.CheckAlerts(ctx)
	triggeredTestAlert := false
	for _, alert := range alerts {
		if alert.ConfigName == "test_alert" {
			triggeredTestAlert = true
		}
	}
	require.False(t, triggeredTestAlert)

	// Record enough circuit trips to trigger
	for i := 0; i < 10; i++ {
		k.RecordCircuitTrip(ctx)
	}

	// Now alert should trigger
	alerts = k.CheckAlerts(ctx)
	triggeredTestAlert = false
	for _, alert := range alerts {
		if alert.ConfigName == "test_alert" {
			triggeredTestAlert = true
			require.Equal(t, types.AlertSeverityWarning, alert.Severity)
		}
	}
	require.True(t, triggeredTestAlert)
}

func TestRecordTransaction(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record successful transaction
	k.RecordTransaction(sdk.WrapSDKContext(ctx), "settlement", true, 25.5)

	metricsData := k.GetSystemMetrics(ctx)
	health := metricsData.ModuleHealth["settlement"]
	require.NotNil(t, health)
	require.Equal(t, uint64(1), health.Transactions)
	require.Equal(t, 25.5, health.Latency)
	require.Equal(t, float64(0), health.ErrorRate)

	// Record failed transaction
	k.RecordTransaction(sdk.WrapSDKContext(ctx), "settlement", false, 100.0)

	metricsData = k.GetSystemMetrics(ctx)
	health = metricsData.ModuleHealth["settlement"]
	require.NotNil(t, health)
	require.Equal(t, uint64(2), health.Transactions)
	require.Greater(t, health.ErrorRate, float64(0))
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up state
	k.RecordSettlement(ctx, sdkmath.NewInt(500000))
	k.RecordCircuitTrip(ctx)
	k.UpdateActiveChannels(ctx, 5)

	config := types.AlertConfig{
		Name:       "custom_alert",
		MetricName: "stale_price_count",
		Condition:  types.AlertConditionGreaterThan,
		Threshold:  sdkmath.LegacyNewDec(0),
		Severity:   types.AlertSeverityInfo,
		Enabled:    true,
	}
	k.SetAlertConfig(ctx, config)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Equal(t, uint64(1), genesis.SystemMetrics.TotalSettlements)
	require.Equal(t, uint64(1), genesis.SystemMetrics.CircuitTrips)
	require.Equal(t, uint64(5), genesis.SystemMetrics.ActiveChannels)

	// Create new keeper and import
	k2, ctx2 := setupKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify state was imported
	imported := k2.GetSystemMetrics(ctx2)
	require.Equal(t, uint64(1), imported.TotalSettlements)
	require.Equal(t, uint64(1), imported.CircuitTrips)
	require.Equal(t, uint64(5), imported.ActiveChannels)

	// Verify alert configs imported
	importedConfigs := k2.GetAlertConfigs(ctx2)
	foundCustom := false
	for _, c := range importedConfigs {
		if c.Name == "custom_alert" {
			foundCustom = true
		}
	}
	require.True(t, foundCustom)
}

func TestMultipleModuleHealth(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Record health for multiple modules
	modules := []string{"settlement", "stablecoin", "payments", "oracle", "compliance"}
	for i, module := range modules {
		k.UpdateModuleHealth(ctx, module, "healthy", 0.01*float64(i), 10.0+float64(i*5))
	}

	metricsData := k.GetSystemMetrics(ctx)
	require.Len(t, metricsData.ModuleHealth, 5)

	for _, module := range modules {
		health, ok := metricsData.ModuleHealth[module]
		require.True(t, ok, "module %s should exist", module)
		require.NotNil(t, health)
		require.Equal(t, module, health.Module)
	}
}

func TestAlertConditions(t *testing.T) {
	k, ctx := setupKeeper(t)

	tests := []struct {
		name      string
		condition types.AlertCondition
		threshold sdkmath.LegacyDec
		value     uint64
		expected  bool
	}{
		{
			name:      "greater than - triggered",
			condition: types.AlertConditionGreaterThan,
			threshold: sdkmath.LegacyNewDec(5),
			value:     10,
			expected:  true,
		},
		{
			name:      "greater than - not triggered",
			condition: types.AlertConditionGreaterThan,
			threshold: sdkmath.LegacyNewDec(15),
			value:     10,
			expected:  false,
		},
		{
			name:      "less than - triggered",
			condition: types.AlertConditionLessThan,
			threshold: sdkmath.LegacyNewDec(15),
			value:     10,
			expected:  true,
		},
		{
			name:      "less than - not triggered",
			condition: types.AlertConditionLessThan,
			threshold: sdkmath.LegacyNewDec(5),
			value:     10,
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Reset metrics
			metricsData := types.DefaultSystemMetrics()
			metricsData.CircuitTrips = tt.value
			k.SetSystemMetrics(ctx, metricsData)

			config := types.AlertConfig{
				Name:       "test_condition",
				MetricName: "circuit_trips",
				Condition:  tt.condition,
				Threshold:  tt.threshold,
				Severity:   types.AlertSeverityWarning,
				Enabled:    true,
			}
			k.SetAlertConfig(ctx, config)

			alerts := k.CheckAlerts(ctx)
			triggered := false
			for _, alert := range alerts {
				if alert.ConfigName == "test_condition" {
					triggered = true
				}
			}
			require.Equal(t, tt.expected, triggered, "alert should be triggered=%v", tt.expected)
		})
	}
}

func TestDisabledAlerts(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set metrics that would trigger alert
	metricsData := types.DefaultSystemMetrics()
	metricsData.CircuitTrips = 100
	k.SetSystemMetrics(ctx, metricsData)

	// Set disabled alert
	config := types.AlertConfig{
		Name:       "disabled_alert",
		MetricName: "circuit_trips",
		Condition:  types.AlertConditionGreaterThan,
		Threshold:  sdkmath.LegacyNewDec(5),
		Severity:   types.AlertSeverityCritical,
		Enabled:    false, // Disabled
	}
	k.SetAlertConfig(ctx, config)

	// Alert should not trigger
	alerts := k.CheckAlerts(ctx)
	for _, alert := range alerts {
		require.NotEqual(t, "disabled_alert", alert.ConfigName)
	}
}
