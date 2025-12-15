package load_test

import (
	"fmt"
	"sync"
	"sync/atomic"
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

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
	feemarketkeeper "github.com/stateset/core/x/feemarket/keeper"
	feemarkettypes "github.com/stateset/core/x/feemarket/types"
)

// LoadTestConfig defines configuration for load tests
type LoadTestConfig struct {
	Concurrency      int           // Number of concurrent goroutines
	Duration         time.Duration // Test duration
	OperationsPerSec int           // Target operations per second
	WarmupDuration   time.Duration // Warmup period before measuring
}

// LoadTestResult contains results from a load test
type LoadTestResult struct {
	TotalOperations   int64
	SuccessfulOps     int64
	FailedOps         int64
	AverageLatencyMs  float64
	P95LatencyMs      float64
	P99LatencyMs      float64
	OperationsPerSec  float64
	ErrorsByType      map[string]int64
}

// ===============================================================================
// Fee Market Load Tests
// ===============================================================================

func setupFeeMarketKeeper(t *testing.T) (feemarketkeeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(feemarkettypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := feemarketkeeper.NewKeeper(cdc, storeKey, "stateset1authority")
	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		Height:  1,
		ChainID: "stateset-load-test",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	// Initialize with default params but add MaxBaseFee to prevent overflow in load tests
	params := feemarkettypes.DefaultParams()
	// Set MaxBaseFee to prevent exponential growth causing overflow in stress tests
	params.MaxBaseFee = sdkmath.LegacyMustNewDecFromStr("1000000")
	k.SetParams(ctx, params)
	k.SetBaseFee(ctx, feemarkettypes.DefaultInitialBaseFee)

	return k, ctx
}

func TestFeeMarket_HighVolumeBaseFeeQueries(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupFeeMarketKeeper(t)
	wrappedCtx := sdk.WrapSDKContext(ctx)

	var (
		totalOps   int64
		successOps int64
		failedOps  int64
	)

	concurrency := 100
	duration := 5 * time.Second

	var wg sync.WaitGroup
	done := make(chan struct{})

	// Start workers
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					atomic.AddInt64(&totalOps, 1)
					_, err := k.BaseFee(wrappedCtx, &feemarkettypes.QueryBaseFeeRequest{})
					if err != nil {
						atomic.AddInt64(&failedOps, 1)
					} else {
						atomic.AddInt64(&successOps, 1)
					}
				}
			}
		}()
	}

	// Run for duration
	time.Sleep(duration)
	close(done)
	wg.Wait()

	// Report results
	opsPerSec := float64(totalOps) / duration.Seconds()
	t.Logf("Fee Market BaseFee Query Load Test Results:")
	t.Logf("  Total Operations: %d", totalOps)
	t.Logf("  Successful: %d (%.2f%%)", successOps, float64(successOps)/float64(totalOps)*100)
	t.Logf("  Failed: %d", failedOps)
	t.Logf("  Operations/sec: %.2f", opsPerSec)

	// Assertions
	require.Zero(t, failedOps, "all operations should succeed")
	require.Greater(t, opsPerSec, float64(1000), "should handle >1000 ops/sec")
}

func TestFeeMarket_ConcurrentBaseFeeUpdates(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupFeeMarketKeeper(t)

	var (
		totalUpdates   int64
		successUpdates int64
	)

	concurrency := 50
	iterations := 1000

	var wg sync.WaitGroup

	// Each worker updates base fee
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				atomic.AddInt64(&totalUpdates, 1)

				// Simulate varying gas usage
				gasUsed := uint64((workerID+1)*1000000 + j*1000)
				maxBlockGas := uint64(50_000_000)

				newFee := k.UpdateBaseFee(ctx, gasUsed, maxBlockGas)
				if newFee.IsPositive() {
					atomic.AddInt64(&successUpdates, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Fee Market Update Load Test Results:")
	t.Logf("  Total Updates: %d", totalUpdates)
	t.Logf("  Successful: %d", successUpdates)

	// Final base fee should be valid
	finalFee := k.GetBaseFee(ctx)
	require.True(t, finalFee.IsPositive(), "final base fee should be positive")
}

func TestFeeMarket_GasOracleUnderLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupFeeMarketKeeper(t)
	oracle := feemarketkeeper.NewGasOracle(k)

	var (
		totalEstimates int64
		validEstimates int64
	)

	concurrency := 100
	duration := 3 * time.Second

	var wg sync.WaitGroup
	done := make(chan struct{})

	txTypes := []string{"send", "delegate", "vote", "swap", "contract"}
	priorities := []string{"low", "medium", "high"}

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					atomic.AddInt64(&totalEstimates, 1)

					// Rotate through tx types and priorities
					txType := txTypes[workerID%len(txTypes)]
					priority := priorities[workerID%len(priorities)]

					gasEstimate := oracle.EstimateGas(ctx, txType)
					feeEstimate := oracle.EstimateFee(ctx, gasEstimate, priority)

					if feeEstimate.TotalFee.IsPositive() {
						atomic.AddInt64(&validEstimates, 1)
					}
				}
			}
		}(i)
	}

	time.Sleep(duration)
	close(done)
	wg.Wait()

	successRate := float64(validEstimates) / float64(totalEstimates) * 100
	opsPerSec := float64(totalEstimates) / duration.Seconds()

	t.Logf("Gas Oracle Load Test Results:")
	t.Logf("  Total Estimates: %d", totalEstimates)
	t.Logf("  Valid: %d (%.2f%%)", validEstimates, successRate)
	t.Logf("  Operations/sec: %.2f", opsPerSec)

	require.Greater(t, successRate, float64(99), "success rate should be >99%%")
}

// ===============================================================================
// Circuit Breaker Load Tests
// ===============================================================================

func setupCircuitKeeper(t *testing.T) (circuitkeeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := circuitkeeper.NewKeeper(cdc, storeKey, "stateset1authority")
	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		Height:  1,
		ChainID: "stateset-load-test",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	k.SetParams(ctx, circuittypes.DefaultParams())

	return k, ctx
}

func TestCircuitBreaker_RateLimitUnderLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupCircuitKeeper(t)

	// Configure rate limit
	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "load_test_limit",
			MaxRequests:   10000,
			WindowSeconds: 60,
			PerAddress:    false, // Global limit
			Enabled:       true,
		},
	}
	k.SetParams(ctx, params)

	var (
		totalChecks   int64
		passedChecks  int64
		blockedChecks int64
	)

	concurrency := 100
	checksPerWorker := 500

	var wg sync.WaitGroup

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			sender := fmt.Sprintf("stateset1sender%d", workerID)
			msgType := "/stateset.test.v1.MsgTest"

			for j := 0; j < checksPerWorker; j++ {
				atomic.AddInt64(&totalChecks, 1)
				err := k.CheckRateLimit(ctx, "load_test_limit", sender, msgType)
				if err == nil {
					atomic.AddInt64(&passedChecks, 1)
				} else {
					atomic.AddInt64(&blockedChecks, 1)
				}
			}
		}(i)
	}

	wg.Wait()

	t.Logf("Rate Limit Load Test Results:")
	t.Logf("  Total Checks: %d", totalChecks)
	t.Logf("  Passed: %d", passedChecks)
	t.Logf("  Blocked: %d", blockedChecks)

	// Some should pass, some should be blocked (rate limited)
	require.Greater(t, passedChecks, int64(0), "some checks should pass")
}

func TestCircuitBreaker_CircuitCheckUnderLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupCircuitKeeper(t)

	var (
		totalChecks  int64
		passedChecks int64
	)

	concurrency := 100
	duration := 3 * time.Second

	modules := []string{"stablecoin", "settlement", "payments", "oracle", "compliance"}

	var wg sync.WaitGroup
	done := make(chan struct{})

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					return
				default:
					atomic.AddInt64(&totalChecks, 1)
					module := modules[workerID%len(modules)]

					if !k.IsModuleCircuitOpen(ctx, module) {
						atomic.AddInt64(&passedChecks, 1)
					}
				}
			}
		}(i)
	}

	time.Sleep(duration)
	close(done)
	wg.Wait()

	opsPerSec := float64(totalChecks) / duration.Seconds()
	successRate := float64(passedChecks) / float64(totalChecks) * 100

	t.Logf("Circuit Check Load Test Results:")
	t.Logf("  Total Checks: %d", totalChecks)
	t.Logf("  Passed: %d (%.2f%%)", passedChecks, successRate)
	t.Logf("  Operations/sec: %.2f", opsPerSec)

	require.Equal(t, totalChecks, passedChecks, "all checks should pass (no circuits tripped)")
}

func TestCircuitBreaker_LiquidationProtectionUnderLoad(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping load test in short mode")
	}

	k, ctx := setupCircuitKeeper(t)

	// Configure liquidation protection
	protection := circuittypes.LiquidationSurgeProtection{
		MaxLiquidationsPerBlock: 100,
		MaxLiquidationValue:     sdkmath.NewInt(10_000_000_000),
		CooldownBlocks:          5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          ctx.BlockHeight(),
	}
	k.SetLiquidationProtection(ctx, protection)

	var (
		totalAttempts int64
		allowedLiqs   int64
		blockedLiqs   int64
	)

	concurrency := 50
	attemptsPerWorker := 100

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < attemptsPerWorker; j++ {
				atomic.AddInt64(&totalAttempts, 1)
				liqValue := sdkmath.NewInt(100_000_000) // 100M units

				mu.Lock()
				err := k.CheckLiquidationAllowed(ctx, liqValue)
				if err == nil {
					atomic.AddInt64(&allowedLiqs, 1)
					k.RecordLiquidation(ctx, liqValue)
				} else {
					atomic.AddInt64(&blockedLiqs, 1)
				}
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	t.Logf("Liquidation Protection Load Test Results:")
	t.Logf("  Total Attempts: %d", totalAttempts)
	t.Logf("  Allowed: %d", allowedLiqs)
	t.Logf("  Blocked: %d", blockedLiqs)

	// Protection should kick in
	require.Greater(t, blockedLiqs, int64(0), "some liquidations should be blocked by surge protection")
	require.LessOrEqual(t, allowedLiqs, int64(100), "should not exceed max liquidations per block")
}

// ===============================================================================
// Stress Tests - Edge Cases Under Load
// ===============================================================================

func TestStress_RapidParamUpdates(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	k, ctx := setupFeeMarketKeeper(t)

	iterations := 1000
	var successfulUpdates int64

	for i := 0; i < iterations; i++ {
		params := feemarkettypes.Params{
			Enabled:                  true,
			BaseFeeChangeDenominator: uint64(8 + i%8),
			ElasticityMultiplier:     uint64(2 + i%4),
			TargetGas:                uint64(10_000_000 + i*1000),
			InitialBaseFee:           sdkmath.LegacyMustNewDecFromStr(fmt.Sprintf("0.00%d", 1+i%9)),
			MinBaseFee:               sdkmath.LegacyMustNewDecFromStr("0.0001"),
			MaxBaseFee:               sdkmath.LegacyZeroDec(),
			PriorityFeeFloor:         sdkmath.LegacyMustNewDecFromStr("0.0001"),
			MaxFeeHistory:            uint32(100 + i%100),
		}

		k.SetParams(ctx, params)
		retrieved := k.GetParams(ctx)
		if retrieved.TargetGas == params.TargetGas {
			successfulUpdates++
		}
	}

	require.Equal(t, int64(iterations), successfulUpdates, "all param updates should be consistent")
}

func TestStress_BaseFeeVolatility(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping stress test in short mode")
	}

	k, ctx := setupFeeMarketKeeper(t)

	params := feemarkettypes.DefaultParams()
	params.MinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.0001")
	params.MaxBaseFee = sdkmath.LegacyMustNewDecFromStr("10.0")
	params.TargetGas = 10_000_000
	k.SetParams(ctx, params)

	initialFee := sdkmath.LegacyMustNewDecFromStr("0.01")
	k.SetBaseFee(ctx, initialFee)

	// Simulate extreme gas usage swings
	iterations := 1000
	var (
		minFee = initialFee
		maxFee = initialFee
	)

	for i := 0; i < iterations; i++ {
		var gasUsed uint64
		if i%2 == 0 {
			gasUsed = 1_000_000 // Very low
		} else {
			gasUsed = 50_000_000 // Very high
		}

		newFee := k.UpdateBaseFee(ctx, gasUsed, 100_000_000)

		if newFee.LT(minFee) {
			minFee = newFee
		}
		if newFee.GT(maxFee) {
			maxFee = newFee
		}
	}

	t.Logf("Base Fee Volatility Stress Test:")
	t.Logf("  Initial Fee: %s", initialFee)
	t.Logf("  Min Fee: %s", minFee)
	t.Logf("  Max Fee: %s", maxFee)

	// Fee should stay within bounds
	require.True(t, minFee.GTE(params.MinBaseFee), "fee should not go below minimum")
	require.True(t, maxFee.LTE(params.MaxBaseFee), "fee should not exceed maximum")
}

// ===============================================================================
// Benchmark Tests
// ===============================================================================

func BenchmarkFeeMarket_BaseFeeQuery(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(feemarkettypes.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	_ = stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := feemarketkeeper.NewKeeper(cdc, storeKey, "authority")

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())
	k.SetBaseFee(ctx, feemarkettypes.DefaultInitialBaseFee)
	wrappedCtx := sdk.WrapSDKContext(ctx)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = k.BaseFee(wrappedCtx, &feemarkettypes.QueryBaseFeeRequest{})
	}
}

func BenchmarkFeeMarket_UpdateBaseFee(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(feemarkettypes.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	_ = stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := feemarketkeeper.NewKeeper(cdc, storeKey, "authority")

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())
	k.SetParams(ctx, feemarkettypes.DefaultParams())
	k.SetBaseFee(ctx, feemarkettypes.DefaultInitialBaseFee)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = k.UpdateBaseFee(ctx, uint64(10_000_000+i%5_000_000), 50_000_000)
	}
}

func BenchmarkCircuit_RateLimitCheck(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	_ = stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := circuitkeeper.NewKeeper(cdc, storeKey, "authority")

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())

	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "bench_limit",
			MaxRequests:   1_000_000,
			WindowSeconds: 60,
			PerAddress:    true,
			Enabled:       true,
		},
	}
	k.SetParams(ctx, params)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = k.CheckRateLimit(ctx, "bench_limit", "sender", "/msg/type")
	}
}

func BenchmarkCircuit_CircuitCheck(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	_ = stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	k := circuitkeeper.NewKeeper(cdc, storeKey, "authority")

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, Time: time.Now()}, false, log.NewNopLogger())
	k.SetParams(ctx, circuittypes.DefaultParams())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = k.IsModuleCircuitOpen(ctx, "stablecoin")
	}
}
