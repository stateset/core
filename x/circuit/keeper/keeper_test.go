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

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
)

func setupKeeper(t *testing.T) (circuitkeeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := circuitkeeper.NewKeeper(
		cdc,
		storeKey,
		"stateset1authority",
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	return k, ctx
}

func TestGlobalPause(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Initially not paused
	require.False(t, k.IsGloballyPaused(ctx))

	// Pause the system
	err := k.PauseSystem(ctx, "stateset1authority", "testing pause", 0)
	require.NoError(t, err)

	// Should be paused now
	require.True(t, k.IsGloballyPaused(ctx))

	// Get circuit state
	state := k.GetCircuitState(ctx)
	require.True(t, state.GlobalPaused)
	require.Equal(t, "testing pause", state.Reason)
	require.Equal(t, "stateset1authority", state.PausedBy)

	// Double pause should fail
	err = k.PauseSystem(ctx, "stateset1authority", "another pause", 0)
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrAlreadyPaused)

	// Resume the system
	err = k.ResumeSystem(ctx, "stateset1authority")
	require.NoError(t, err)

	// Should not be paused anymore
	require.False(t, k.IsGloballyPaused(ctx))

	// Double resume should fail
	err = k.ResumeSystem(ctx, "stateset1authority")
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrNotPaused)
}

func TestGlobalPauseWithDuration(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Pause for 1 hour
	err := k.PauseSystem(ctx, "stateset1authority", "timed pause", 3600)
	require.NoError(t, err)
	require.True(t, k.IsGloballyPaused(ctx))

	// Move time forward by 30 minutes - should still be paused
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(30 * time.Minute))
	require.True(t, k.IsGloballyPaused(ctx))

	// Move time forward by 2 hours - should auto-resume
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(2 * time.Hour))
	require.False(t, k.IsGloballyPaused(ctx))
}

func TestModuleCircuitBreaker(t *testing.T) {
	k, ctx := setupKeeper(t)

	moduleName := "stablecoin"

	// Initially circuit should be closed
	require.False(t, k.IsModuleCircuitOpen(ctx, moduleName))

	// Trip the circuit
	err := k.TripCircuit(ctx, moduleName, "testing trip", "stateset1authority", nil)
	require.NoError(t, err)

	// Circuit should be open
	require.True(t, k.IsModuleCircuitOpen(ctx, moduleName))

	// Get module circuit state
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	require.True(t, found)
	require.Equal(t, circuittypes.CircuitOpen, state.Status)
	require.Equal(t, "testing trip", state.Reason)

	// Reset the circuit
	err = k.ResetCircuit(ctx, moduleName, "stateset1authority")
	require.NoError(t, err)

	// Circuit should be closed
	require.False(t, k.IsModuleCircuitOpen(ctx, moduleName))
}

func TestModuleCircuitWithDisabledMessages(t *testing.T) {
	k, ctx := setupKeeper(t)

	moduleName := "settlement"
	disabledMsgs := []string{"/stateset.settlement.v1.MsgInstantTransfer"}

	// Trip circuit with specific messages disabled
	err := k.TripCircuit(ctx, moduleName, "disable instant transfers", "stateset1authority", disabledMsgs)
	require.NoError(t, err)

	// Check specific message is disabled
	require.True(t, k.IsMessageDisabled(ctx, moduleName, "/stateset.settlement.v1.MsgInstantTransfer"))

	// Check other message is not disabled
	require.False(t, k.IsMessageDisabled(ctx, moduleName, "/stateset.settlement.v1.MsgCreateEscrow"))
}

func TestRateLimiting(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params with rate limits
	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "test_limit",
			MaxRequests:   5,
			WindowSeconds: 60,
			PerAddress:    true,
			Enabled:       true,
		},
	}
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	sender := "stateset1sender123"
	msgType := "/stateset.stablecoin.v1.MsgMint"

	// First 5 requests should succeed
	for i := 0; i < 5; i++ {
		err := k.CheckRateLimit(ctx, "test_limit", sender, msgType)
		require.NoError(t, err, "request %d should succeed", i+1)
	}

	// 6th request should fail
	err = k.CheckRateLimit(ctx, "test_limit", sender, msgType)
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrRateLimitExceeded)

	// Different sender should still work
	err = k.CheckRateLimit(ctx, "test_limit", "stateset1different", msgType)
	require.NoError(t, err)

	// Move time forward to reset window
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(61 * time.Second))

	// Original sender should work again
	err = k.CheckRateLimit(ctx, "test_limit", sender, msgType)
	require.NoError(t, err)
}

func TestGlobalRateLimiting(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up params with global rate limit
	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "global_limit",
			MaxRequests:   3,
			WindowSeconds: 60,
			PerAddress:    false, // Global
			Enabled:       true,
		},
	}
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	msgType := "/stateset.test.v1.MsgTest"

	// First 3 requests from different addresses should succeed
	err = k.CheckRateLimit(ctx, "global_limit", "sender1", msgType)
	require.NoError(t, err)
	err = k.CheckRateLimit(ctx, "global_limit", "sender2", msgType)
	require.NoError(t, err)
	err = k.CheckRateLimit(ctx, "global_limit", "sender3", msgType)
	require.NoError(t, err)

	// 4th request should fail even from new sender
	err = k.CheckRateLimit(ctx, "global_limit", "sender4", msgType)
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrRateLimitExceeded)
}

func TestLiquidationSurgeProtection(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up liquidation protection
	protection := circuittypes.LiquidationSurgeProtection{
		MaxLiquidationsPerBlock: 3,
		MaxLiquidationValue:     sdkmath.NewInt(1000000),
		CooldownBlocks:          5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          0,
	}
	k.SetLiquidationProtection(ctx, protection)

	// First liquidations should be allowed
	for i := 0; i < 3; i++ {
		err := k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(100000))
		require.NoError(t, err)
		k.RecordLiquidation(ctx, sdkmath.NewInt(100000))
	}

	// 4th liquidation should fail
	err := k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(100000))
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrLiquidationSurge)

	// Move to next block
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	// Should work again
	err = k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(100000))
	require.NoError(t, err)
}

func TestLiquidationValueLimit(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up with value limit
	protection := circuittypes.LiquidationSurgeProtection{
		MaxLiquidationsPerBlock: 100, // High count limit
		MaxLiquidationValue:     sdkmath.NewInt(500000),
		CooldownBlocks:          5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          ctx.BlockHeight(),
	}
	k.SetLiquidationProtection(ctx, protection)

	// First liquidation with large value
	err := k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(300000))
	require.NoError(t, err)
	k.RecordLiquidation(ctx, sdkmath.NewInt(300000))

	// Second liquidation pushing over limit should fail
	err = k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(300000))
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrLiquidationSurge)

	// Smaller value should still fail (cumulative check)
	err = k.CheckLiquidationAllowed(ctx, sdkmath.NewInt(250000))
	require.Error(t, err)
}

func TestAuthorization(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Authority should be authorized
	require.True(t, k.IsAuthorized(ctx, "stateset1authority"))

	// Random address should not be authorized
	require.False(t, k.IsAuthorized(ctx, "stateset1random"))

	// Add address to authorities in params
	params := k.GetParams(ctx)
	params.Authorities = append(params.Authorities, "stateset1operator")
	k.SetParams(ctx, params)

	// Now operator should be authorized
	require.True(t, k.IsAuthorized(ctx, "stateset1operator"))
}

func TestRecordFailureAutoTrip(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set params with low failure threshold
	params := circuittypes.DefaultParams()
	params.DefaultFailureThreshold = 3
	k.SetParams(ctx, params)

	moduleName := "payments"

	// Record failures
	k.RecordFailure(ctx, moduleName)
	require.False(t, k.IsModuleCircuitOpen(ctx, moduleName))

	k.RecordFailure(ctx, moduleName)
	require.False(t, k.IsModuleCircuitOpen(ctx, moduleName))

	// Third failure should trip the circuit
	k.RecordFailure(ctx, moduleName)
	require.True(t, k.IsModuleCircuitOpen(ctx, moduleName))

	// Verify state
	state, found := k.GetModuleCircuitState(ctx, moduleName)
	require.True(t, found)
	require.Equal(t, "automatic", state.TrippedBy)
	require.Equal(t, "failure threshold exceeded", state.Reason)
}

func TestRecordSuccess(t *testing.T) {
	k, ctx := setupKeeper(t)

	moduleName := "oracle"

	// Record some failures
	k.RecordFailure(ctx, moduleName)
	k.RecordFailure(ctx, moduleName)

	state, found := k.GetModuleCircuitState(ctx, moduleName)
	require.True(t, found)
	require.Equal(t, uint64(2), state.FailureCount)

	// Record success should reset failure count
	k.RecordSuccess(ctx, moduleName)

	state, _ = k.GetModuleCircuitState(ctx, moduleName)
	require.Equal(t, uint64(0), state.FailureCount)
}

func TestCircuitRecovery(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set params with short recovery period
	params := circuittypes.DefaultParams()
	params.DefaultRecoveryPeriod = 60 // 60 seconds
	k.SetParams(ctx, params)

	moduleName := "compliance"

	// Trip the circuit
	err := k.TripCircuit(ctx, moduleName, "test", "authority", nil)
	require.NoError(t, err)
	require.True(t, k.IsModuleCircuitOpen(ctx, moduleName))

	// Verify recovery time is set
	state, _ := k.GetModuleCircuitState(ctx, moduleName)
	require.False(t, state.RecoveryTime.IsZero())

	// Move time forward past recovery
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(61 * time.Second))

	// Circuit should auto-recover (when checked with half-open logic)
	// For this test, manually reset as the half-open check is in IsModuleCircuitOpen
	state.Status = circuittypes.CircuitHalfOpen
	k.SetModuleCircuitState(ctx, state)

	// Now the check should close the circuit
	require.False(t, k.IsModuleCircuitOpen(ctx, moduleName))
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up some state
	err := k.PauseSystem(ctx, "authority", "genesis test", 3600)
	require.NoError(t, err)

	err = k.TripCircuit(ctx, "testmodule", "test trip", "authority", nil)
	require.NoError(t, err)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.True(t, genesis.CircuitState.GlobalPaused)
	require.Len(t, genesis.ModuleCircuits, 1)

	// Create new keeper and import
	k2, ctx2 := setupKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify state was imported
	require.True(t, k2.IsGloballyPaused(ctx2))
	require.True(t, k2.IsModuleCircuitOpen(ctx2, "testmodule"))
}

func TestCheckCircuitBreakers(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Initially all checks should pass
	err := k.CheckCircuitBreakers(sdk.WrapSDKContext(ctx), "stablecoin", "/stateset.stablecoin.v1.MsgMint", "sender1")
	require.NoError(t, err)

	// Pause system
	err = k.PauseSystem(ctx, "authority", "test", 0)
	require.NoError(t, err)

	// Check should fail
	err = k.CheckCircuitBreakers(sdk.WrapSDKContext(ctx), "stablecoin", "/stateset.stablecoin.v1.MsgMint", "sender1")
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrGlobalPause)
}

func TestParamsValidation(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Valid params should work
	params := circuittypes.DefaultParams()
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	// Invalid params should fail
	params.DefaultFailureThreshold = 0
	err = k.SetParams(ctx, params)
	require.Error(t, err)

	params = circuittypes.DefaultParams()
	params.DefaultRecoveryPeriod = 0
	err = k.SetParams(ctx, params)
	require.Error(t, err)

	params = circuittypes.DefaultParams()
	params.MaxPauseDuration = -1
	err = k.SetParams(ctx, params)
	require.Error(t, err)
}

func TestMessageTypeFiltering(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up rate limit for specific message types
	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "mint_limit",
			MaxRequests:   2,
			WindowSeconds: 60,
			PerAddress:    true,
			Enabled:       true,
			MessageTypes:  []string{"/stateset.stablecoin.v1.MsgMint"},
		},
	}
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	sender := "stateset1sender"

	// Mint messages should be limited
	err = k.CheckRateLimit(ctx, "mint_limit", sender, "/stateset.stablecoin.v1.MsgMint")
	require.NoError(t, err)
	err = k.CheckRateLimit(ctx, "mint_limit", sender, "/stateset.stablecoin.v1.MsgMint")
	require.NoError(t, err)
	err = k.CheckRateLimit(ctx, "mint_limit", sender, "/stateset.stablecoin.v1.MsgMint")
	require.Error(t, err)

	// Other messages should not be affected
	err = k.CheckRateLimit(ctx, "mint_limit", sender, "/stateset.stablecoin.v1.MsgRepay")
	require.NoError(t, err)
}
