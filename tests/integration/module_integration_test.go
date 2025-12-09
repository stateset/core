package integration_test

import (
	"context"
	"sync"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	oraclekeeper "github.com/stateset/core/x/oracle/keeper"
	oracletypes "github.com/stateset/core/x/oracle/types"
)

var configOnce sync.Once

func setupConfig() {
	configOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount("stateset", "statesetpub")
		cfg.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		cfg.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		cfg.Seal()
	})
}

func newAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

// IntegrationTestSuite holds all keepers for integration testing
type IntegrationTestSuite struct {
	ctx               sdk.Context
	cdc               codec.Codec
	circuitKeeper     circuitkeeper.Keeper
	complianceKeeper  compliancekeeper.Keeper
	oracleKeeper      oraclekeeper.Keeper
	authority         string
}

func setupIntegrationSuite(t *testing.T) *IntegrationTestSuite {
	t.Helper()
	setupConfig()

	// Create store keys
	circuitStoreKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)
	complianceStoreKey := storetypes.NewKVStoreKey(compliancetypes.StoreKey)
	oracleStoreKey := storetypes.NewKVStoreKey(oracletypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(circuitStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(complianceStoreKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(oracleStoreKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-integration-test", Time: time.Now()}, false, log.NewNopLogger())

	authority := newAddress().String()

	// Create keepers
	circuitKeeper := circuitkeeper.NewKeeper(cdc, circuitStoreKey, authority)
	complianceKeeper := compliancekeeper.NewKeeper(cdc, complianceStoreKey, authority)
	oracleKeeper := oraclekeeper.NewKeeper(cdc, oracleStoreKey, authority)

	return &IntegrationTestSuite{
		ctx:              ctx,
		cdc:              cdc,
		circuitKeeper:    circuitKeeper,
		complianceKeeper: complianceKeeper,
		oracleKeeper:     oracleKeeper,
		authority:        authority,
	}
}

// TestCircuitBreakerBlocksCompliance tests that circuit breaker can halt compliance module
func TestCircuitBreakerBlocksCompliance(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Initially circuit should be closed
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "compliance"))

	// Compliance operations should work
	profile := compliancetypes.ComplianceProfile{
		Address:   newAddress().String(),
		KycLevel:  compliancetypes.KycLevelBasic,
		RiskLevel: compliancetypes.RiskLevelLow,
	}
	err := suite.complianceKeeper.SetProfile(suite.ctx, profile)
	require.NoError(t, err)

	// Trip the circuit for compliance module
	err = suite.circuitKeeper.TripCircuit(suite.ctx, "compliance", "emergency maintenance", suite.authority, nil)
	require.NoError(t, err)

	// Verify circuit is open
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "compliance"))

	// Reset the circuit
	err = suite.circuitKeeper.ResetCircuit(suite.ctx, "compliance", suite.authority)
	require.NoError(t, err)

	// Verify circuit is closed again
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "compliance"))
}

// TestGlobalPauseAffectsAllModules tests that global pause affects all modules
func TestGlobalPauseAffectsAllModules(t *testing.T) {
	suite := setupIntegrationSuite(t)

	modules := []string{"oracle", "compliance", "settlement", "stablecoin"}

	// Initially not paused
	require.False(t, suite.circuitKeeper.IsGloballyPaused(suite.ctx))

	// Pause the system
	err := suite.circuitKeeper.PauseSystem(suite.ctx, suite.authority, "security incident", 0)
	require.NoError(t, err)

	// Verify global pause is active
	require.True(t, suite.circuitKeeper.IsGloballyPaused(suite.ctx))

	// Check circuit breakers for all modules
	for _, module := range modules {
		err := suite.circuitKeeper.CheckCircuitBreakers(sdk.WrapSDKContext(suite.ctx), module, "/test", "sender")
		require.Error(t, err)
		require.ErrorIs(t, err, circuittypes.ErrGlobalPause)
	}

	// Resume the system
	err = suite.circuitKeeper.ResumeSystem(suite.ctx, suite.authority)
	require.NoError(t, err)

	// Verify checks pass again
	for _, module := range modules {
		err := suite.circuitKeeper.CheckCircuitBreakers(sdk.WrapSDKContext(suite.ctx), module, "/test", "sender")
		require.NoError(t, err)
	}
}

// TestComplianceBlocksNonKycAddress tests that compliance module blocks non-KYC addresses
func TestComplianceBlocksNonKycAddress(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Create a profile with no KYC
	noKycAddr := newAddress()
	profile := compliancetypes.ComplianceProfile{
		Address:   noKycAddr.String(),
		KycLevel:  compliancetypes.KycLevelNone,
		RiskLevel: compliancetypes.RiskLevelUnknown,
	}
	err := suite.complianceKeeper.SetProfile(suite.ctx, profile)
	require.NoError(t, err)

	// Address without KYC should fail compliance check for high-value operations
	err = suite.complianceKeeper.AssertCompliant(suite.ctx, noKycAddr)
	require.Error(t, err)

	// Upgrade to basic KYC
	profile.KycLevel = compliancetypes.KycLevelBasic
	profile.RiskLevel = compliancetypes.RiskLevelLow
	err = suite.complianceKeeper.SetProfile(suite.ctx, profile)
	require.NoError(t, err)

	// Now should pass
	err = suite.complianceKeeper.AssertCompliant(suite.ctx, noKycAddr)
	require.NoError(t, err)
}

// TestComplianceSanctionedAddress tests that sanctioned addresses are blocked
func TestComplianceSanctionedAddress(t *testing.T) {
	suite := setupIntegrationSuite(t)

	sanctionedAddr := newAddress()

	// Create a profile but mark as sanctioned
	profile := compliancetypes.ComplianceProfile{
		Address:    sanctionedAddr.String(),
		KycLevel:   compliancetypes.KycLevelEnhanced,
		RiskLevel:  compliancetypes.RiskLevelHigh,
		Sanctioned: true,
	}
	err := suite.complianceKeeper.SetProfile(suite.ctx, profile)
	require.NoError(t, err)

	// Sanctioned address should fail compliance even with full KYC
	err = suite.complianceKeeper.AssertCompliant(suite.ctx, sanctionedAddr)
	require.Error(t, err)
}

// TestOraclePriceUpdates tests oracle price update flow
func TestOraclePriceUpdates(t *testing.T) {
	suite := setupIntegrationSuite(t)

	denom := "stst"
	provider := newAddress().String()

	// Set up oracle params
	params := oracletypes.DefaultParams()
	err := suite.oracleKeeper.SetParams(suite.ctx, params)
	require.NoError(t, err)

	// Initial price update
	price := sdkmath.LegacyMustNewDecFromStr("2.50")
	err = suite.oracleKeeper.UpdatePrice(suite.ctx, denom, price, provider)
	require.NoError(t, err)

	// Get price
	storedPrice, err := suite.oracleKeeper.GetPriceDec(suite.ctx, denom)
	require.NoError(t, err)
	require.True(t, price.Equal(storedPrice))

	// Update with small deviation (should succeed)
	newPrice := sdkmath.LegacyMustNewDecFromStr("2.55") // 2% increase
	err = suite.oracleKeeper.UpdatePrice(suite.ctx, denom, newPrice, provider)
	require.NoError(t, err)

	// Verify price was updated
	storedPrice, err = suite.oracleKeeper.GetPriceDec(suite.ctx, denom)
	require.NoError(t, err)
	require.True(t, newPrice.Equal(storedPrice))
}

// TestRateLimitingAcrossModules tests that rate limiting applies consistently
func TestRateLimitingAcrossModules(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Set up rate limits
	params := circuittypes.DefaultParams()
	params.RateLimits = []circuittypes.RateLimitConfig{
		{
			Name:          "global_tx_limit",
			MaxRequests:   5,
			WindowSeconds: 60,
			PerAddress:    true,
			Enabled:       true,
		},
	}
	err := suite.circuitKeeper.SetParams(suite.ctx, params)
	require.NoError(t, err)

	sender := "stateset1sender"

	// First 5 requests should succeed
	for i := 0; i < 5; i++ {
		err := suite.circuitKeeper.CheckRateLimit(suite.ctx, "global_tx_limit", sender, "/test")
		require.NoError(t, err, "request %d should succeed", i+1)
	}

	// 6th request should fail
	err = suite.circuitKeeper.CheckRateLimit(suite.ctx, "global_tx_limit", sender, "/test")
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrRateLimitExceeded)

	// Different sender should not be affected
	err = suite.circuitKeeper.CheckRateLimit(suite.ctx, "global_tx_limit", "stateset1different", "/test")
	require.NoError(t, err)
}

// TestLiquidationSurgeProtection tests liquidation surge protection
func TestLiquidationSurgeProtection(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Set up liquidation protection
	protection := circuittypes.LiquidationSurgeProtection{
		MaxLiquidationsPerBlock: 3,
		MaxLiquidationValue:     sdkmath.NewInt(1_000_000),
		CooldownBlocks:          5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          suite.ctx.BlockHeight(),
	}
	suite.circuitKeeper.SetLiquidationProtection(suite.ctx, protection)

	// First 3 liquidations should be allowed
	for i := 0; i < 3; i++ {
		err := suite.circuitKeeper.CheckLiquidationAllowed(suite.ctx, sdkmath.NewInt(100_000))
		require.NoError(t, err, "liquidation %d should be allowed", i+1)
		suite.circuitKeeper.RecordLiquidation(suite.ctx, sdkmath.NewInt(100_000))
	}

	// 4th liquidation should be blocked
	err := suite.circuitKeeper.CheckLiquidationAllowed(suite.ctx, sdkmath.NewInt(100_000))
	require.Error(t, err)
	require.ErrorIs(t, err, circuittypes.ErrLiquidationSurge)

	// Move to next block
	suite.ctx = suite.ctx.WithBlockHeight(suite.ctx.BlockHeight() + 1)

	// Should be allowed again
	err = suite.circuitKeeper.CheckLiquidationAllowed(suite.ctx, sdkmath.NewInt(100_000))
	require.NoError(t, err)
}

// TestComplianceJurisdictionBlocking tests jurisdiction-based blocking
func TestComplianceJurisdictionBlocking(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Address from blocked jurisdiction
	blockedAddr := newAddress()
	profile := compliancetypes.ComplianceProfile{
		Address:      blockedAddr.String(),
		KycLevel:     compliancetypes.KycLevelEnhanced,
		RiskLevel:    compliancetypes.RiskLevelLow,
		Jurisdiction: "KP", // North Korea - blocked
	}
	err := suite.complianceKeeper.SetProfile(suite.ctx, profile)
	require.NoError(t, err)

	// Should fail due to jurisdiction
	err = suite.complianceKeeper.AssertCompliant(suite.ctx, blockedAddr)
	require.Error(t, err)

	// Address from allowed jurisdiction
	allowedAddr := newAddress()
	allowedProfile := compliancetypes.ComplianceProfile{
		Address:      allowedAddr.String(),
		KycLevel:     compliancetypes.KycLevelBasic,
		RiskLevel:    compliancetypes.RiskLevelLow,
		Jurisdiction: "US",
	}
	err = suite.complianceKeeper.SetProfile(suite.ctx, allowedProfile)
	require.NoError(t, err)

	// Should pass
	err = suite.complianceKeeper.AssertCompliant(suite.ctx, allowedAddr)
	require.NoError(t, err)
}

// TestAutoFailureTripping tests automatic circuit tripping on failures
func TestAutoFailureTripping(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Set params with low failure threshold
	params := circuittypes.DefaultParams()
	params.DefaultFailureThreshold = 3
	err := suite.circuitKeeper.SetParams(suite.ctx, params)
	require.NoError(t, err)

	moduleName := "oracle"

	// Record failures
	for i := 0; i < 2; i++ {
		suite.circuitKeeper.RecordFailure(suite.ctx, moduleName)
		require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, moduleName))
	}

	// Third failure should trip the circuit
	suite.circuitKeeper.RecordFailure(suite.ctx, moduleName)
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, moduleName))

	// Verify state
	state, found := suite.circuitKeeper.GetModuleCircuitState(suite.ctx, moduleName)
	require.True(t, found)
	require.Equal(t, "automatic", state.TrippedBy)
}

// TestTimedPauseAutoResume tests that timed pause auto-resumes
func TestTimedPauseAutoResume(t *testing.T) {
	suite := setupIntegrationSuite(t)

	// Pause for 1 hour
	err := suite.circuitKeeper.PauseSystem(suite.ctx, suite.authority, "scheduled maintenance", 3600)
	require.NoError(t, err)
	require.True(t, suite.circuitKeeper.IsGloballyPaused(suite.ctx))

	// After 30 minutes, should still be paused
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(30 * time.Minute))
	require.True(t, suite.circuitKeeper.IsGloballyPaused(suite.ctx))

	// After 2 hours, should auto-resume
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(90 * time.Minute))
	require.False(t, suite.circuitKeeper.IsGloballyPaused(suite.ctx))
}

// TestMultipleModuleCircuits tests independent module circuits
func TestMultipleModuleCircuits(t *testing.T) {
	suite := setupIntegrationSuite(t)

	modules := []string{"oracle", "stablecoin", "settlement"}

	// Trip only oracle circuit
	err := suite.circuitKeeper.TripCircuit(suite.ctx, "oracle", "price feed issue", suite.authority, nil)
	require.NoError(t, err)

	// Only oracle should be tripped
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "oracle"))
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "stablecoin"))
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "settlement"))

	// Trip stablecoin circuit
	err = suite.circuitKeeper.TripCircuit(suite.ctx, "stablecoin", "vault issue", suite.authority, nil)
	require.NoError(t, err)

	// Oracle and stablecoin should be tripped
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "oracle"))
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "stablecoin"))
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "settlement"))

	// Reset oracle
	err = suite.circuitKeeper.ResetCircuit(suite.ctx, "oracle", suite.authority)
	require.NoError(t, err)

	// Only stablecoin should remain tripped
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "oracle"))
	require.True(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "stablecoin"))
	require.False(t, suite.circuitKeeper.IsModuleCircuitOpen(suite.ctx, "settlement"))
}

// Mock bank keeper for compliance integration
type mockBankKeeper struct {
	balances map[string]sdk.Coins
}

func newMockBankKeeper() *mockBankKeeper {
	return &mockBankKeeper{
		balances: make(map[string]sdk.Coins),
	}
}

func (m *mockBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.balances[addr.String()]
	return sdk.NewCoin(denom, coins.AmountOf(denom))
}

func (m *mockBankKeeper) SendCoins(ctx context.Context, from, to sdk.AccAddress, amt sdk.Coins) error {
	return nil
}
