package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/oracle/types"
)

// TestStalenessLogic verifies that prices are rejected or marked stale correctly
func TestStalenessLogic(t *testing.T) {
	k, ctx := setupKeeper(t)

	// 1. Setup params with 60 second staleness threshold
	params := k.GetParams(ctx)
	params.DefaultStalenessThreshold = 60
	k.SetParams(ctx, params)

	// 1b. Setup Config with 60 second staleness threshold
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.StalenessThresholdSeconds = 60
	require.NoError(t, k.SetOracleConfig(ctx, config))

	// 2. Set a price at T=0
	initialTime := ctx.BlockTime()
	price := types.Price{
		Denom:       "uusdc",
		Amount:      sdkmath.LegacyOneDec(),
		LastUpdater: "provider1",
		UpdatedAt:   initialTime,
	}
	k.SetPrice(ctx, price)

	// 3. Query at T+30s (Fresh)
	ctx = ctx.WithBlockTime(initialTime.Add(30 * time.Second))
	retrieved, err := k.GetPriceWithStalenessCheck(ctx, "uusdc")
	require.NoError(t, err)
	require.Equal(t, price.Amount, retrieved.Amount)

	// 4. Query at T+61s (Stale)
	ctx = ctx.WithBlockTime(initialTime.Add(61 * time.Second))
	_, err = k.GetPriceWithStalenessCheck(ctx, "uusdc")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPriceStale)

	// 5. Verify ProcessStalePrices emits event
	// Reset context to stale time
	ctx = ctx.WithBlockTime(initialTime.Add(61 * time.Second))
	
	// Create a fresh event manager to capture events
	ctx = ctx.WithEventManager(sdk.NewEventManager())
	k.ProcessStalePrices(ctx)

	events := ctx.EventManager().Events()
	require.Len(t, events, 1)
	require.Equal(t, "oracle_price_stale", events[0].Type)
	require.Equal(t, "uusdc", string(events[0].Attributes[0].Value))
}

// TestAutomaticSlashing verifies that providers are slashed after poor performance
func TestAutomaticSlashing(t *testing.T) {
	k, ctx := setupKeeper(t)

	// 1. Register a provider
	providerAddr := "provider_bad_performance"
	provider := types.OracleProvider{
		Address:               providerAddr,
		IsActive:              true,
		TotalSubmissions:      0,
		SuccessfulSubmissions: 0,
	}
	require.NoError(t, k.SetProvider(ctx, provider))

	// 2. Simulate 10 successful submissions (100% success)
	for i := 0; i < 10; i++ {
		// Manually simulate success recording (as we can't easily trigger SetPriceWithValidation success in this loop without complex setup)
		// We use the internal logic simulation here or helper
		// But keeper_test doesn't expose `recordProviderSuccess`. 
		// We will simulate by updating the provider state directly which logic relies on.
		
		p, _ := k.GetProvider(ctx, providerAddr)
		p.TotalSubmissions++
		p.SuccessfulSubmissions++
		k.SetProvider(ctx, p)
	}

	p, _ := k.GetProvider(ctx, providerAddr)
	require.False(t, p.Slashed)

	// 3. Simulate 15 failed submissions (bringing rate below 50%)
	// Total: 25, Success: 10. Rate: 40%
	
	// We need to trigger the *logic* that checks for slashing.
	// This logic resides in `recordProviderFailure` inside keeper.go.
	// Since that method is private, we must trigger it via `SetPriceWithValidation`.
	
	// Setup config for failure
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.MaxDeviationBps = 100 // 1% deviation
	config.MinUpdateIntervalSeconds = 0 // Disable rate limiting for test
	require.NoError(t, k.SetOracleConfig(ctx, config))

	// Set baseline price
	k.SetPrice(ctx, types.Price{
		Denom:     "uusdc",
		Amount:    sdkmath.LegacyNewDec(100),
		UpdatedAt: ctx.BlockTime(),
	})

	for i := 0; i < 15; i++ {
		// Advance time to avoid update frequency check
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second))

		// Submit price with huge deviation to trigger failure
		// Price 200 vs 100 = 100% deviation
		err := k.SetPriceWithValidation(ctx, providerAddr, "uusdc", sdkmath.LegacyNewDec(200))
		require.Error(t, err) // Should fail deviation check
	}

	// 4. Verify Provider is Slashed
	p, found := k.GetProvider(ctx, providerAddr)
	require.True(t, found)
	require.True(t, p.Slashed, "Provider should be slashed due to low success rate")
	require.False(t, p.IsActive, "Slashed provider should be inactive")
}

// TestPendingPricesCleanup verifies that pending prices don't leak state
func TestPendingPricesCleanup(t *testing.T) {
	k, ctx := setupKeeper(t)

	// 1. Configure for multiple confirmations
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.RequiredConfirmations = 3
	require.NoError(t, k.SetOracleConfig(ctx, config))

	// 2. Submit 2/3 confirmations
	require.NoError(t, k.SetPriceWithValidation(ctx, "p1", "uusdc", sdkmath.LegacyNewDec(100)))
	require.NoError(t, k.SetPriceWithValidation(ctx, "p2", "uusdc", sdkmath.LegacyNewDec(100)))

	// Check price not set yet
	_, found := k.GetPrice(ctx, "uusdc")
	require.False(t, found)

	// 3. Submit 3rd confirmation -> Should trigger set and cleanup
	require.NoError(t, k.SetPriceWithValidation(ctx, "p3", "uusdc", sdkmath.LegacyNewDec(100)))

	// Check price set
	_, found = k.GetPrice(ctx, "uusdc")
	require.True(t, found)

	// 4. Verify internal storage cleared (using exported genesis to check)
	// If cleanup worked, PendingPrices in genesis should be empty for this denom
	genesis := k.ExportGenesis(ctx)
	for _, pp := range genesis.PendingPrices {
		require.NotEqual(t, "uusdc", pp.Denom, "Pending prices for uusdc should be cleared")
	}
}
