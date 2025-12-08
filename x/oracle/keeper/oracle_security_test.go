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

	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

func setupOracleKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		"stateset1authority",
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	return k, ctx
}

func TestSetPriceWithValidation_Success(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config
	config := types.OracleConfig{
		Denom:                     "uatom",
		MaxDeviationBps:           500, // 5%
		StalenessThresholdSeconds: 3600,
		MinUpdateIntervalSeconds:  60,
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set initial price
	err := k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1000))
	require.NoError(t, err)

	// Verify price was set
	price, found := k.GetPrice(ctx, "uatom")
	require.True(t, found)
	require.Equal(t, sdkmath.LegacyNewDec(1000), price.Amount)
	require.Equal(t, "stateset1authority", price.LastUpdater)
}

func TestSetPriceWithValidation_DeviationTooLarge(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config with strict deviation limit
	config := types.OracleConfig{
		Denom:                     "uatom",
		MaxDeviationBps:           500, // 5%
		StalenessThresholdSeconds: 3600,
		MinUpdateIntervalSeconds:  1, // Short interval for testing
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set initial price
	initialPrice := types.Price{
		Denom:       "uatom",
		Amount:      sdkmath.LegacyNewDec(1000),
		LastUpdater: "stateset1authority",
		LastHeight:  1,
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetPrice(ctx, initialPrice)

	// Move time forward
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(2 * time.Second))

	// Try to set price with 10% deviation (should fail)
	err := k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1100))
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrDeviationTooLarge)

	// Try with acceptable deviation (3%)
	err = k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1030))
	require.NoError(t, err)
}

func TestSetPriceWithValidation_UpdateTooFrequent(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config
	config := types.OracleConfig{
		Denom:                     "uatom",
		MaxDeviationBps:           1000,
		StalenessThresholdSeconds: 3600,
		MinUpdateIntervalSeconds:  60, // 1 minute minimum
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set initial price
	err := k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1000))
	require.NoError(t, err)

	// Try to update immediately (should fail)
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(30 * time.Second))
	err = k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1001))
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUpdateTooFrequent)

	// Wait and try again (should succeed)
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(31 * time.Second))
	err = k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1001))
	require.NoError(t, err)
}

func TestSetPriceWithValidation_ConfigDisabled(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up disabled oracle config
	config := types.OracleConfig{
		Denom:   "uatom",
		Enabled: false,
	}
	k.SetOracleConfig(ctx, config)

	// Try to set price (should fail)
	err := k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(1000))
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrConfigDisabled)
}

func TestGetPriceWithStalenessCheck(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config
	config := types.OracleConfig{
		Denom:                     "uatom",
		MaxDeviationBps:           500,
		StalenessThresholdSeconds: 3600, // 1 hour
		MinUpdateIntervalSeconds:  60,
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set initial price
	price := types.Price{
		Denom:       "uatom",
		Amount:      sdkmath.LegacyNewDec(1000),
		LastUpdater: "stateset1authority",
		LastHeight:  1,
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	// Price should be fresh
	result, err := k.GetPriceWithStalenessCheck(ctx, "uatom")
	require.NoError(t, err)
	require.Equal(t, sdkmath.LegacyNewDec(1000), result.Amount)

	// Move time forward by 30 minutes (still fresh)
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(30 * time.Minute))
	result, err = k.GetPriceWithStalenessCheck(ctx, "uatom")
	require.NoError(t, err)

	// Move time forward past staleness threshold
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(45 * time.Minute))
	_, err = k.GetPriceWithStalenessCheck(ctx, "uatom")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPriceStale)
}

func TestOracleProvider(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Create provider
	provider := types.OracleProvider{
		Address:               "stateset1provider1",
		Name:                  "Test Provider",
		Weight:                100,
		IsActive:              true,
		TotalSubmissions:      0,
		SuccessfulSubmissions: 0,
		Slashed:               false,
		SlashCount:            0,
	}
	k.SetProvider(ctx, provider)

	// Get provider
	result, found := k.GetProvider(ctx, "stateset1provider1")
	require.True(t, found)
	require.Equal(t, "Test Provider", result.Name)
	require.True(t, result.IsActive)

	// Check authorization
	require.True(t, k.IsAuthorizedProvider(ctx, "stateset1provider1"))

	// Non-existent provider should not be authorized
	require.False(t, k.IsAuthorizedProvider(ctx, "stateset1unknown"))

	// Authority should always be authorized
	require.True(t, k.IsAuthorizedProvider(ctx, "stateset1authority"))
}

func TestOracleProviderSlashing(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Create provider
	provider := types.OracleProvider{
		Address:               "stateset1badprovider",
		Name:                  "Bad Provider",
		Weight:                100,
		IsActive:              true,
		TotalSubmissions:      0,
		SuccessfulSubmissions: 0,
		Slashed:               false,
		SlashCount:            0,
	}
	k.SetProvider(ctx, provider)

	// Manual slash
	err := k.SlashProvider(ctx, "stateset1badprovider", "malicious behavior")
	require.NoError(t, err)

	// Verify slashed
	result, found := k.GetProvider(ctx, "stateset1badprovider")
	require.True(t, found)
	require.True(t, result.Slashed)
	require.False(t, result.IsActive)
	require.Equal(t, uint32(1), result.SlashCount)

	// Slashed provider should not be authorized
	require.False(t, k.IsAuthorizedProvider(ctx, "stateset1badprovider"))
}

func TestPriceHistorySecurity(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config
	config := types.OracleConfig{
		Denom:                     "uatom",
		MaxDeviationBps:           1000, // 10%
		StalenessThresholdSeconds: 3600,
		MinUpdateIntervalSeconds:  1, // Short for testing
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set params with history size
	params := types.DefaultOracleParams()
	params.PriceHistorySize = 5
	k.SetParams(ctx, params)

	// Set multiple prices
	prices := []int64{1000, 1010, 1020, 1015, 1025, 1030, 1035}
	for i, p := range prices {
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(2 * time.Second))
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
		err := k.SetPriceWithValidation(ctx, "stateset1authority", "uatom", sdkmath.LegacyNewDec(p))
		require.NoError(t, err, "failed at price %d (index %d)", p, i)
	}

	// Get history
	history, found := k.GetPriceHistory(ctx, "uatom")
	require.True(t, found)

	// Should only have last 5 entries
	require.Len(t, history.Prices, 5)

	// Verify last price is most recent
	require.Equal(t, sdkmath.LegacyNewDec(1035), history.Prices[len(history.Prices)-1].Amount)
}

func TestCalculateDeviation(t *testing.T) {
	tests := []struct {
		name     string
		oldPrice sdkmath.LegacyDec
		newPrice sdkmath.LegacyDec
		expected sdkmath.LegacyDec
	}{
		{
			name:     "10% increase",
			oldPrice: sdkmath.LegacyNewDec(1000),
			newPrice: sdkmath.LegacyNewDec(1100),
			expected: sdkmath.LegacyNewDec(1000), // 10% = 1000 bps
		},
		{
			name:     "10% decrease",
			oldPrice: sdkmath.LegacyNewDec(1000),
			newPrice: sdkmath.LegacyNewDec(900),
			expected: sdkmath.LegacyNewDec(1000), // 10% = 1000 bps
		},
		{
			name:     "5% increase",
			oldPrice: sdkmath.LegacyNewDec(1000),
			newPrice: sdkmath.LegacyNewDec(1050),
			expected: sdkmath.LegacyNewDec(500), // 5% = 500 bps
		},
		{
			name:     "zero old price",
			oldPrice: sdkmath.LegacyZeroDec(),
			newPrice: sdkmath.LegacyNewDec(1000),
			expected: sdkmath.LegacyZeroDec(), // No deviation from zero
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := types.CalculateDeviation(tt.oldPrice, tt.newPrice)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestPriceIsStale(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name       string
		updatedAt  time.Time
		threshold  time.Duration
		isStale    bool
	}{
		{
			name:      "fresh price",
			updatedAt: now.Add(-30 * time.Minute),
			threshold: 1 * time.Hour,
			isStale:   false,
		},
		{
			name:      "stale price",
			updatedAt: now.Add(-2 * time.Hour),
			threshold: 1 * time.Hour,
			isStale:   true,
		},
		{
			name:      "exactly at threshold",
			updatedAt: now.Add(-1 * time.Hour),
			threshold: 1 * time.Hour,
			isStale:   false,
		},
		{
			name:      "zero timestamp",
			updatedAt: time.Time{},
			threshold: 1 * time.Hour,
			isStale:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price := types.Price{
				UpdatedAt: tt.updatedAt,
			}
			result := price.IsStale(now, tt.threshold)
			require.Equal(t, tt.isStale, result)
		})
	}
}

func TestIterateProviders(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Create multiple providers
	providers := []types.OracleProvider{
		{Address: "stateset1provider1", Name: "Provider 1", IsActive: true},
		{Address: "stateset1provider2", Name: "Provider 2", IsActive: true},
		{Address: "stateset1provider3", Name: "Provider 3", IsActive: false},
	}

	for _, p := range providers {
		k.SetProvider(ctx, p)
	}

	// Iterate and count
	count := 0
	activeCount := 0
	k.IterateProviders(ctx, func(p types.OracleProvider) bool {
		count++
		if p.IsActive {
			activeCount++
		}
		return false
	})

	require.Equal(t, 3, count)
	require.Equal(t, 2, activeCount)
}

func TestGetPriceDecSafe(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Set up oracle config
	config := types.OracleConfig{
		Denom:                     "uatom",
		StalenessThresholdSeconds: 60, // 1 minute
		Enabled:                   true,
	}
	k.SetOracleConfig(ctx, config)

	// Set fresh price
	price := types.Price{
		Denom:     "uatom",
		Amount:    sdkmath.LegacyNewDec(1000),
		UpdatedAt: ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	// Should work when fresh
	result, err := k.GetPriceDecSafe(ctx, "uatom")
	require.NoError(t, err)
	require.Equal(t, sdkmath.LegacyNewDec(1000), result)

	// Move time forward past threshold
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(2 * time.Minute))

	// Should fail when stale
	_, err = k.GetPriceDecSafe(ctx, "uatom")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPriceStale)
}

func TestRemoveProvider(t *testing.T) {
	k, ctx := setupOracleKeeper(t)

	// Create provider
	provider := types.OracleProvider{
		Address:  "stateset1provider1",
		Name:     "Provider 1",
		IsActive: true,
	}
	k.SetProvider(ctx, provider)

	// Verify exists
	_, found := k.GetProvider(ctx, "stateset1provider1")
	require.True(t, found)

	// Remove
	k.RemoveProvider(ctx, "stateset1provider1")

	// Verify removed
	_, found = k.GetProvider(ctx, "stateset1provider1")
	require.False(t, found)
}

func TestDefaultOracleConfig(t *testing.T) {
	config := types.DefaultOracleConfig("uatom")

	require.Equal(t, "uatom", config.Denom)
	require.Equal(t, uint64(500), config.MaxDeviationBps)
	require.Equal(t, int64(3600), config.StalenessThresholdSeconds)
	require.Equal(t, int64(60), config.MinUpdateIntervalSeconds)
	require.Equal(t, uint32(1), config.RequiredConfirmations)
	require.True(t, config.Enabled)
}

func TestDefaultOracleParams(t *testing.T) {
	params := types.DefaultOracleParams()

	require.Equal(t, uint64(500), params.DefaultMaxDeviationBps)
	require.Equal(t, int64(3600), params.DefaultStalenessThreshold)
	require.Equal(t, uint64(1000), params.SlashFractionBps)
	require.Equal(t, uint32(10), params.MaxProviders)
	require.Equal(t, uint32(100), params.PriceHistorySize)
}
