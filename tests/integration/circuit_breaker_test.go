package integration

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	oraclekeeper "github.com/stateset/core/x/oracle/keeper"
	oracletypes "github.com/stateset/core/x/oracle/types"
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
	stablecoinkeeper "github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

// CircuitBreakerTestSuite tests circuit breaker functionality:
// Detect anomaly -> Trigger circuit breaker -> Block transactions -> Recovery
type CircuitBreakerTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	cdc            codec.Codec

	// Keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	circuitKeeper    circuitkeeper.Keeper
	complianceKeeper compliancekeeper.Keeper
	oracleKeeper     oraclekeeper.Keeper
	settlementKeeper settlementkeeper.Keeper
	stablecoinKeeper stablecoinkeeper.Keeper

	// Test accounts
	authority       sdk.AccAddress
	user1           sdk.AccAddress
	user2           sdk.AccAddress
	oracleProvider  sdk.AccAddress
}

func TestCircuitBreakerTestSuite(t *testing.T) {
	suite.Run(t, new(CircuitBreakerTestSuite))
}

func (s *CircuitBreakerTestSuite) SetupTest() {
	// Initialize codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	s.cdc = codec.NewProtoCodec(interfaceRegistry)

	// Create store keys
	storeKeys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		circuittypes.StoreKey,
		compliancetypes.StoreKey,
		oracletypes.StoreKey,
		settlementtypes.StoreKey,
		stablecointypes.StoreKey,
	)

	// Create transient store keys
	tKeys := storetypes.NewTransientStoreKeys(banktypes.TransientKey)

	// Create multistore
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		stateStore.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)
	}
	require.NoError(s.T(), stateStore.LoadLatestVersion())

	// Create context
	s.ctx = testutil.DefaultContextWithDB(s.T(), storeKeys[authtypes.StoreKey], tKeys[banktypes.TransientKey]).Ctx.
		WithBlockHeight(1).
		WithBlockTime(time.Now())

	// Initialize test addresses
	s.authority = sdk.AccAddress([]byte("authority___________"))
	s.user1 = sdk.AccAddress([]byte("user1_______________"))
	s.user2 = sdk.AccAddress([]byte("user2_______________"))
	s.oracleProvider = sdk.AccAddress([]byte("oracleprovider______"))

	// Initialize account keeper
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:         nil,
		settlementtypes.ModuleAccountName:  {authtypes.Minter, authtypes.Burner},
		stablecointypes.ModuleAccountName:  {authtypes.Minter, authtypes.Burner},
	}
	s.accountKeeper = authkeeper.NewAccountKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		"stateset",
		s.authority.String(),
	)

	// Initialize bank keeper
	s.bankKeeper = bankkeeper.NewBaseKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[banktypes.StoreKey]),
		s.accountKeeper,
		nil,
		s.authority.String(),
		log.NewNopLogger(),
	)

	// Initialize circuit keeper
	s.circuitKeeper = circuitkeeper.NewKeeper(
		s.cdc,
		storeKeys[circuittypes.StoreKey],
		s.authority.String(),
	)

	// Initialize compliance keeper
	s.complianceKeeper = compliancekeeper.NewKeeper(
		s.cdc,
		storeKeys[compliancetypes.StoreKey],
		s.authority.String(),
	)

	// Initialize oracle keeper
	s.oracleKeeper = oraclekeeper.NewKeeper(
		s.cdc,
		storeKeys[oracletypes.StoreKey],
		s.authority.String(),
	)

	// Initialize settlement keeper
	s.settlementKeeper = settlementkeeper.NewKeeper(
		s.cdc,
		storeKeys[settlementtypes.StoreKey],
		s.bankKeeper,
		s.complianceKeeper,
		s.accountKeeper,
		s.authority.String(),
	)

	// Initialize stablecoin keeper
	s.stablecoinKeeper = stablecoinkeeper.NewKeeper(
		s.cdc,
		storeKeys[stablecointypes.StoreKey],
		s.authority.String(),
		s.bankKeeper,
		s.accountKeeper,
		s.oracleKeeper,
		s.complianceKeeper,
	)

	// Setup test data
	s.setupTestAccounts()
	s.setupCircuitParams()
}

func (s *CircuitBreakerTestSuite) setupTestAccounts() {
	// Create module accounts
	settlementModuleAcc := authtypes.NewEmptyModuleAccount(settlementtypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	stablecoinModuleAcc := authtypes.NewEmptyModuleAccount(stablecointypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	s.accountKeeper.SetModuleAccount(s.ctx, settlementModuleAcc)
	s.accountKeeper.SetModuleAccount(s.ctx, stablecoinModuleAcc)

	// Create user accounts
	for _, addr := range []sdk.AccAddress{s.user1, s.user2, s.oracleProvider} {
		acc := authtypes.NewBaseAccountWithAddress(addr)
		s.accountKeeper.SetAccount(s.ctx, acc)

		// Mint coins
		coins := sdk.NewCoins(
			sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000)),
			sdk.NewCoin("uatom", sdkmath.NewInt(10000000000)),
		)
		require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, coins))
		require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, addr, coins))

		// Setup compliance
		profile := compliancetypes.Profile{
			Address:    addr.String(),
			Status:     compliancetypes.StatusActive,
			KYCLevel:   compliancetypes.KYCStandard,
			Sanction:   false,
			VerifiedAt: s.ctx.BlockTime(),
			ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
			LastLimitReset: s.ctx.BlockTime(),
		}
		s.complianceKeeper.SetProfile(s.ctx, profile)
	}
}

func (s *CircuitBreakerTestSuite) setupCircuitParams() {
	params := circuittypes.Params{
		Enabled:                  true,
		MaxPauseDuration:         86400, // 24 hours
		DefaultFailureThreshold:  5,
		DefaultRecoveryPeriod:    3600, // 1 hour
		Authorities:              []string{s.authority.String()},
		RateLimits:               []circuittypes.RateLimitConfig{},
	}
	err := s.circuitKeeper.SetParams(s.ctx, params)
	s.Require().NoError(err)
}

// TestGlobalSystemPause tests pausing the entire system
func (s *CircuitBreakerTestSuite) TestGlobalSystemPause() {
	// Scenario: Critical issue detected, pause entire system

	// Verify system is operational initially
	s.Require().False(s.circuitKeeper.IsGloballyPaused(s.ctx))

	// Pause the system
	reason := "Critical security vulnerability detected"
	durationSeconds := int64(3600) // 1 hour

	err := s.circuitKeeper.PauseSystem(s.ctx, s.authority.String(), reason, durationSeconds)
	s.Require().NoError(err, "System pause should succeed")

	// Verify system is paused
	s.Require().True(s.circuitKeeper.IsGloballyPaused(s.ctx))

	// Verify circuit state
	state := s.circuitKeeper.GetCircuitState(s.ctx)
	s.Require().True(state.GlobalPaused)
	s.Require().Equal(reason, state.Reason)
	s.Require().Equal(s.authority.String(), state.PausedBy)
	s.Require().False(state.AutoResumeAt.IsZero())

	// Try to perform operations - should be blocked
	// This would normally be checked in ante handler
	err = s.circuitKeeper.CheckCircuitBreakers(s.ctx, "settlement", "InstantTransfer", s.user1.String())
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "global pause")

	// Resume system manually
	err = s.circuitKeeper.ResumeSystem(s.ctx, s.authority.String())
	s.Require().NoError(err, "System resume should succeed")

	// Verify system is operational
	s.Require().False(s.circuitKeeper.IsGloballyPaused(s.ctx))
}

// TestModuleCircuitTrip tests tripping circuit for specific module
func (s *CircuitBreakerTestSuite) TestModuleCircuitTrip() {
	// Scenario: Stablecoin module has issues, trip its circuit

	moduleName := "stablecoin"
	reason := "Unusual liquidation activity detected"

	// Verify module is operational
	s.Require().False(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName))

	// Trip the circuit
	err := s.circuitKeeper.TripCircuit(s.ctx, moduleName, reason, s.authority.String(), nil)
	s.Require().NoError(err, "Circuit trip should succeed")

	// Verify module circuit is open (blocking)
	s.Require().True(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName))

	// Verify circuit state
	state, found := s.circuitKeeper.GetModuleCircuitState(s.ctx, moduleName)
	s.Require().True(found)
	s.Require().Equal(circuittypes.CircuitOpen, state.Status)
	s.Require().Equal(reason, state.Reason)

	// Check that operations are blocked
	err = s.circuitKeeper.CheckCircuitBreakers(s.ctx, moduleName, "MintStablecoin", s.user1.String())
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "circuit")

	// Reset the circuit
	err = s.circuitKeeper.ResetCircuit(s.ctx, moduleName, s.authority.String())
	s.Require().NoError(err, "Circuit reset should succeed")

	// Verify module is operational
	s.Require().False(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName))
}

// TestAutoCircuitTripOnFailures tests automatic circuit trip after threshold
func (s *CircuitBreakerTestSuite) TestAutoCircuitTripOnFailures() {
	// Scenario: Module experiences repeated failures, auto-trip after threshold

	moduleName := "settlement"
	params := s.circuitKeeper.GetParams(s.ctx)
	threshold := params.DefaultFailureThreshold

	// Record failures
	for i := uint32(0); i < threshold; i++ {
		s.circuitKeeper.RecordFailure(s.ctx, moduleName)

		// Circuit should not be open until threshold reached
		if i < threshold-1 {
			s.Require().False(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName),
				"Circuit should not trip before threshold")
		}
	}

	// After threshold failures, circuit should be open
	s.Require().True(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName),
		"Circuit should auto-trip after threshold")

	// Verify it was automatically tripped
	state, found := s.circuitKeeper.GetModuleCircuitState(s.ctx, moduleName)
	s.Require().True(found)
	s.Require().Equal(circuittypes.CircuitOpen, state.Status)
	s.Require().Equal("automatic", state.TrippedBy)
	s.Require().Contains(state.Reason, "threshold exceeded")
}

// TestCircuitAutoRecovery tests automatic circuit recovery after period
func (s *CircuitBreakerTestSuite) TestCircuitAutoRecovery() {
	// Scenario: Circuit trips, then auto-recovers after recovery period

	moduleName := "oracle"

	// Trip circuit
	err := s.circuitKeeper.TripCircuit(s.ctx, moduleName, "test auto-recovery", s.authority.String(), nil)
	s.Require().NoError(err)
	s.Require().True(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName))

	// Get recovery time
	state, _ := s.circuitKeeper.GetModuleCircuitState(s.ctx, moduleName)
	recoveryTime := state.RecoveryTime

	// Advance time past recovery period
	s.ctx = s.ctx.WithBlockTime(recoveryTime.Add(1 * time.Minute))

	// Check if circuit auto-recovered
	isOpen := s.circuitKeeper.IsModuleCircuitOpen(s.ctx, moduleName)
	s.Require().False(isOpen, "Circuit should auto-recover after recovery period")
}

// TestOraclePriceDeviationTriggersCircuit tests circuit trip on large price deviation
func (s *CircuitBreakerTestSuite) TestOraclePriceDeviationTriggersCircuit() {
	// Scenario: Oracle price changes drastically, triggering circuit breaker

	denom := "uatom"

	// Set initial price
	initialPrice := oracletypes.Price{
		Denom:       denom,
		Amount:      sdkmath.LegacyMustNewDecFromStr("10.00"),
		LastUpdater: s.oracleProvider.String(),
		LastHeight:  s.ctx.BlockHeight(),
		UpdatedAt:   s.ctx.BlockTime(),
	}
	s.oracleKeeper.SetPrice(s.ctx, initialPrice)

	// Configure oracle config with deviation limit
	config := oracletypes.OracleConfig{
		Denom:                      denom,
		Enabled:                    true,
		MinUpdateIntervalSeconds:   60,
		StalenessThresholdSeconds:  3600,
		MaxDeviationBps:            500, // 5% max deviation
	}
	err := s.oracleKeeper.SetOracleConfig(s.ctx, config)
	s.Require().NoError(err)

	// Register oracle provider
	provider := oracletypes.OracleProvider{
		Address:               s.oracleProvider.String(),
		IsActive:              true,
		Slashed:               false,
		TotalSubmissions:      1,
		SuccessfulSubmissions: 1,
	}
	require.NoError(s.T(), s.oracleKeeper.SetProvider(s.ctx, provider))

	// Advance time to allow update
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(2 * time.Minute))

	// Try to set price with large deviation (should fail)
	drasticPrice := sdkmath.LegacyMustNewDecFromStr("5.00") // 50% drop
	err = s.oracleKeeper.SetPriceWithValidation(s.ctx, s.oracleProvider.String(), denom, drasticPrice)
	s.Require().Error(err, "Large price deviation should be rejected")
	s.Require().Contains(err.Error(), "deviation")

	// In real system, this would trigger circuit breaker
	// Manually trip circuit to simulate
	err = s.circuitKeeper.TripCircuit(s.ctx, "stablecoin", "Large oracle price deviation detected", s.authority.String(), nil)
	s.Require().NoError(err)

	// Verify stablecoin operations are blocked
	s.Require().True(s.circuitKeeper.IsModuleCircuitOpen(s.ctx, "stablecoin"))
}

// TestLiquidationSurgeProtection tests liquidation surge protection
func (s *CircuitBreakerTestSuite) TestLiquidationSurgeProtection() {
	// Scenario: Multiple liquidations in same block trigger protection

	protection := s.circuitKeeper.GetLiquidationProtection(s.ctx)
	maxPerBlock := protection.MaxLiquidationsPerBlock

	// Simulate liquidations up to limit
	for i := uint64(0); i < maxPerBlock; i++ {
		value := sdkmath.NewInt(1000000000)
		err := s.circuitKeeper.CheckLiquidationAllowed(s.ctx, value)
		s.Require().NoError(err, "Liquidation %d should be allowed", i+1)

		s.circuitKeeper.RecordLiquidation(s.ctx, value)
	}

	// Next liquidation should be blocked
	err := s.circuitKeeper.CheckLiquidationAllowed(s.ctx, sdkmath.NewInt(1000000000))
	s.Require().Error(err, "Liquidation should be blocked after surge")
	s.Require().Contains(err.Error(), "surge")

	// New block resets counter
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1)
	err = s.circuitKeeper.CheckLiquidationAllowed(s.ctx, sdkmath.NewInt(1000000000))
	s.Require().NoError(err, "Liquidation should be allowed in new block")
}

// TestRateLimiting tests rate limiting functionality
func (s *CircuitBreakerTestSuite) TestRateLimiting() {
	// Scenario: High-frequency operations are rate limited

	// Setup rate limit
	params := s.circuitKeeper.GetParams(s.ctx)
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "settlement_rate",
			Enabled:       true,
			MaxRequests:   10,
			WindowSeconds: 60,
			PerAddress:    true,
			MessageTypes:  []string{"InstantTransfer"},
		},
	}
	err := s.circuitKeeper.SetParams(s.ctx, params)
	s.Require().NoError(err)

	msgType := "InstantTransfer"
	configName := "settlement_rate"

	// Make requests up to limit
	for i := uint64(0); i < 10; i++ {
		err := s.circuitKeeper.CheckRateLimit(s.ctx, configName, s.user1.String(), msgType)
		s.Require().NoError(err, "Request %d should be allowed", i+1)
	}

	// Next request should be rate limited
	err = s.circuitKeeper.CheckRateLimit(s.ctx, configName, s.user1.String(), msgType)
	s.Require().Error(err, "Request should be rate limited")
	s.Require().Contains(err.Error(), "rate limit")

	// Different user should have own limit
	err = s.circuitKeeper.CheckRateLimit(s.ctx, configName, s.user2.String(), msgType)
	s.Require().NoError(err, "Different user should have own rate limit")

	// After time window, limit resets
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(61 * time.Second))
	err = s.circuitKeeper.CheckRateLimit(s.ctx, configName, s.user1.String(), msgType)
	s.Require().NoError(err, "Rate limit should reset after window")
}

// TestSelectiveMessageDisabling tests disabling specific message types
func (s *CircuitBreakerTestSuite) TestSelectiveMessageDisabling() {
	// Scenario: Disable only certain operations while keeping others active

	moduleName := "settlement"
	disabledMessages := []string{"CreateEscrow", "OpenChannel"}

	// Trip circuit with specific disabled messages
	err := s.circuitKeeper.TripCircuit(s.ctx, moduleName, "Disable risky operations", s.authority.String(), disabledMessages)
	s.Require().NoError(err)

	// Verify disabled messages are blocked
	for _, msgType := range disabledMessages {
		isDisabled := s.circuitKeeper.IsMessageDisabled(s.ctx, moduleName, msgType)
		s.Require().True(isDisabled, "%s should be disabled", msgType)
	}

	// Other message types should still work
	isDisabled := s.circuitKeeper.IsMessageDisabled(s.ctx, moduleName, "InstantTransfer")
	s.Require().False(isDisabled, "InstantTransfer should not be disabled")
}

// TestCircuitBreakerEvents tests event emission
func (s *CircuitBreakerTestSuite) TestCircuitBreakerEvents() {
	// Scenario: Verify circuit breaker emits proper events

	// Trip circuit
	err := s.circuitKeeper.TripCircuit(s.ctx, "test", "event test", s.authority.String(), nil)
	s.Require().NoError(err)

	// Check events
	events := s.ctx.EventManager().Events()
	s.Require().Greater(len(events), 0)

	var foundCircuitEvent bool
	for _, event := range events {
		if event.Type == "circuit_tripped" {
			foundCircuitEvent = true
			// Verify attributes
			for _, attr := range event.Attributes {
				if attr.Key == "module" {
					s.Require().Equal("test", attr.Value)
				}
			}
		}
	}
	s.Require().True(foundCircuitEvent, "Circuit tripped event should be emitted")
}

// TestUnauthorizedCircuitControl tests that only authorized addresses can control circuits
func (s *CircuitBreakerTestSuite) TestUnauthorizedCircuitControl() {
	// Scenario: Unauthorized user tries to trip circuit

	// User1 is not in authorities list
	s.Require().False(s.circuitKeeper.IsAuthorized(s.ctx, s.user1.String()))

	// Should not be able to trip circuit (would normally be checked in msg handler)
	// For this test, we verify the IsAuthorized check works
	s.Require().True(s.circuitKeeper.IsAuthorized(s.ctx, s.authority.String()))
}

// TestCircuitStateConsistency tests state consistency across operations
func (s *CircuitBreakerTestSuite) TestCircuitStateConsistency() {
	// Scenario: Verify circuit state remains consistent through multiple operations

	moduleName := "test_module"

	// Trip circuit
	err := s.circuitKeeper.TripCircuit(s.ctx, moduleName, "test", s.authority.String(), nil)
	s.Require().NoError(err)

	// Get state
	state1, found1 := s.circuitKeeper.GetModuleCircuitState(s.ctx, moduleName)
	s.Require().True(found1)

	// Reset circuit
	err = s.circuitKeeper.ResetCircuit(s.ctx, moduleName, s.authority.String())
	s.Require().NoError(err)

	// Get state again
	state2, found2 := s.circuitKeeper.GetModuleCircuitState(s.ctx, moduleName)
	s.Require().True(found2)

	// Verify states are different but consistent
	s.Require().NotEqual(state1.Status, state2.Status)
	s.Require().Equal(circuittypes.CircuitClosed, state2.Status)
	s.Require().Equal(uint32(0), state2.FailureCount)
}

// TestMultipleModuleCircuits tests independent circuit states for different modules
func (s *CircuitBreakerTestSuite) TestMultipleModuleCircuits() {
	// Scenario: Different modules can have different circuit states

	modules := []string{"stablecoin", "settlement", "oracle"}

	// Trip circuits for some modules
	for i, module := range modules {
		if i%2 == 0 {
			err := s.circuitKeeper.TripCircuit(s.ctx, module, "test", s.authority.String(), nil)
			s.Require().NoError(err)
		}
	}

	// Verify states are independent
	for i, module := range modules {
		isOpen := s.circuitKeeper.IsModuleCircuitOpen(s.ctx, module)
		if i%2 == 0 {
			s.Require().True(isOpen, "%s circuit should be open", module)
		} else {
			s.Require().False(isOpen, "%s circuit should be closed", module)
		}
	}
}
