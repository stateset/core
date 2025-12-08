package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

func setupKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	testCtx := testutil.DefaultContextWithDB(t, key, storetypes.NewTransientStoreKey("transient_test"))
	ctx := testCtx.Ctx.WithBlockTime(time.Now().UTC())

	k := keeper.NewKeeper(nil, key, "authority")

	return k, ctx
}

func TestSetGetPrice(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Test setting and getting a price
	price := types.Price{
		Denom:       "uusdc",
		Amount:      sdkmath.LegacyNewDec(100),
		LastUpdater: "provider1",
		LastHeight:  100,
		UpdatedAt:   ctx.BlockTime(),
	}

	k.SetPrice(ctx, price)

	// Get the price back
	retrieved, found := k.GetPrice(ctx, "uusdc")
	require.True(t, found)
	require.Equal(t, price.Denom, retrieved.Denom)
	require.Equal(t, price.Amount.String(), retrieved.Amount.String())
	require.Equal(t, price.LastUpdater, retrieved.LastUpdater)

	// Test price not found
	_, found = k.GetPrice(ctx, "nonexistent")
	require.False(t, found)
}

func TestGetPriceDec(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set a price
	price := types.Price{
		Denom:     "uusdc",
		Amount:    sdkmath.LegacyNewDec(150),
		UpdatedAt: ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	// Get as Dec
	dec, err := k.GetPriceDec(ctx, "uusdc")
	require.NoError(t, err)
	require.Equal(t, sdkmath.LegacyNewDec(150), dec)

	// Test error for non-existent price
	_, err = k.GetPriceDec(ctx, "nonexistent")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrPriceNotFound)
}

func TestSetPriceWithValidation(t *testing.T) {
	k, ctx := setupKeeper(t)

	// First, set up an oracle config
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.MaxDeviationBps = 1000 // 10%
	err := k.SetOracleConfig(ctx, config)
	require.NoError(t, err)

	// Set initial price
	err = k.SetPriceWithValidation(ctx, "provider1", "uusdc", sdkmath.LegacyNewDec(100))
	require.NoError(t, err)

	// Verify price was set
	price, found := k.GetPrice(ctx, "uusdc")
	require.True(t, found)
	require.Equal(t, sdkmath.LegacyNewDec(100).String(), price.Amount.String())
}

func TestPriceDeviationRejection(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up oracle config with strict deviation limit
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.MaxDeviationBps = 500    // 5%
	config.MinUpdateIntervalSeconds = 0 // No minimum interval for testing
	err := k.SetOracleConfig(ctx, config)
	require.NoError(t, err)

	// Set initial price
	err = k.SetPriceWithValidation(ctx, "provider1", "uusdc", sdkmath.LegacyNewDec(100))
	require.NoError(t, err)

	// Try to update with price that exceeds deviation (20% change)
	err = k.SetPriceWithValidation(ctx, "provider1", "uusdc", sdkmath.LegacyNewDec(120))
	require.Error(t, err)
	require.ErrorContains(t, err, "deviation")
}

func TestProviderManagement(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Add a provider
	provider := types.OracleProvider{
		Address:  "provider1",
		IsActive: true,
		Slashed:  false,
	}
	err := k.SetProvider(ctx, provider)
	require.NoError(t, err)

	// Get the provider
	retrieved, found := k.GetProvider(ctx, "provider1")
	require.True(t, found)
	require.Equal(t, provider.Address, retrieved.Address)
	require.True(t, retrieved.IsActive)
	require.False(t, retrieved.Slashed)

	// Check authorization
	require.True(t, k.IsAuthorizedProvider(ctx, "provider1"))
	require.False(t, k.IsAuthorizedProvider(ctx, "unknown"))
}

func TestProviderSlashing(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Add a provider
	provider := types.OracleProvider{
		Address:  "provider1",
		IsActive: true,
	}
	err := k.SetProvider(ctx, provider)
	require.NoError(t, err)

	// Slash the provider
	err = k.SlashProvider(ctx, "provider1", "malicious behavior")
	require.NoError(t, err)

	// Verify slashing
	slashed, found := k.GetProvider(ctx, "provider1")
	require.True(t, found)
	require.True(t, slashed.Slashed)
	require.False(t, slashed.IsActive)
	require.Equal(t, uint32(1), slashed.SlashCount)

	// Slashed provider should not be authorized
	require.False(t, k.IsAuthorizedProvider(ctx, "provider1"))
}

func TestIteratePrices(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Add multiple prices
	denoms := []string{"uusdc", "uusdt", "uatom"}
	for i, denom := range denoms {
		price := types.Price{
			Denom:     denom,
			Amount:    sdkmath.LegacyNewDec(int64(100 + i*10)),
			UpdatedAt: ctx.BlockTime(),
		}
		k.SetPrice(ctx, price)
	}

	// Iterate and collect
	var collected []string
	k.IteratePrices(ctx, func(price types.Price) bool {
		collected = append(collected, price.Denom)
		return false
	})

	require.Len(t, collected, 3)
}

func TestDeletePrice(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set a price
	price := types.Price{
		Denom:     "uusdc",
		Amount:    sdkmath.LegacyNewDec(100),
		UpdatedAt: ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	// Verify it exists
	_, found := k.GetPrice(ctx, "uusdc")
	require.True(t, found)

	// Delete it
	k.DeletePrice(ctx, "uusdc")

	// Verify it's gone
	_, found = k.GetPrice(ctx, "uusdc")
	require.False(t, found)
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up some state
	price := types.Price{
		Denom:     "uusdc",
		Amount:    sdkmath.LegacyNewDec(100),
		UpdatedAt: ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Prices, 1)

	// Create new keeper and import
	k2, ctx2 := setupKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify state was imported
	imported, found := k2.GetPrice(ctx2, "uusdc")
	require.True(t, found)
	require.Equal(t, price.Denom, imported.Denom)
}

func TestOracleParams(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Get default params
	params := k.GetParams(ctx)
	require.NotZero(t, params.PriceHistorySize)

	// Update params
	params.PriceHistorySize = 200
	err := k.SetParams(ctx, params)
	require.NoError(t, err)

	// Verify update
	retrieved := k.GetParams(ctx)
	require.Equal(t, uint32(200), retrieved.PriceHistorySize)
}

func TestPriceHistory(t *testing.T) {
	k, ctx := setupKeeper(t)

	// Set up config
	config := types.DefaultOracleConfig("uusdc")
	config.Enabled = true
	config.MinUpdateIntervalSeconds = 0
	config.MaxDeviationBps = 10000 // Allow large changes for testing
	err := k.SetOracleConfig(ctx, config)
	require.NoError(t, err)

	// Add multiple price updates
	for i := 0; i < 5; i++ {
		ctx = ctx.WithBlockHeight(int64(i + 1))
		err := k.SetPriceWithValidation(ctx, "provider1", "uusdc", sdkmath.LegacyNewDec(int64(100+i)))
		require.NoError(t, err)
	}

	// Get price history
	history, found := k.GetPriceHistory(ctx, "uusdc")
	require.True(t, found)
	require.GreaterOrEqual(t, len(history.Prices), 1)
}
