package keeper_test

import (
	"fmt"
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
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/oracle/keeper"
	"github.com/stateset/core/x/oracle/types"
)

func setupBenchmarkOracleKeeper() (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		ChainID: "stateset-bench",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	k := keeper.NewKeeper(cdc, storeKey, "authority")

	return k, ctx
}

func BenchmarkSetPrice(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	price := types.Price{
		Denom:       "stst",
		Amount:      sdkmath.LegacyMustNewDecFromStr("2.0"),
		LastUpdater: "stateset1updater",
		LastHeight:  100,
		UpdatedAt:   ctx.BlockTime(),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		price.LastHeight = int64(i)
		k.SetPrice(ctx, price)
	}
}

func BenchmarkGetPrice(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	// Set initial price
	price := types.Price{
		Denom:       "stst",
		Amount:      sdkmath.LegacyMustNewDecFromStr("2.0"),
		LastUpdater: "stateset1updater",
		LastHeight:  100,
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetPrice(ctx, "stst")
	}
}

func BenchmarkSetPriceWithValidation(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	// Set initial price
	initialPrice := types.Price{
		Denom:       "stst",
		Amount:      sdkmath.LegacyMustNewDecFromStr("2.0"),
		LastUpdater: "stateset1updater",
		LastHeight:  100,
		UpdatedAt:   ctx.BlockTime().Add(-time.Hour),
	}
	k.SetPrice(ctx, initialPrice)

	// Set default oracle config
	config := types.DefaultOracleConfig("stst")
	k.SetOracleConfig(ctx, config)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Small price changes that should pass validation
		newPrice := sdkmath.LegacyMustNewDecFromStr(fmt.Sprintf("2.%04d", i%100))
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Minute))
		k.SetPriceWithValidation(ctx, "authority", "stst", newPrice)
	}
}

func BenchmarkIteratePrices(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	// Set 100 different prices
	for i := 0; i < 100; i++ {
		price := types.Price{
			Denom:       fmt.Sprintf("denom%d", i),
			Amount:      sdkmath.LegacyMustNewDecFromStr(fmt.Sprintf("%d.0", i+1)),
			LastUpdater: "stateset1updater",
			LastHeight:  int64(i),
			UpdatedAt:   ctx.BlockTime(),
		}
		k.SetPrice(ctx, price)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		k.IteratePrices(ctx, func(price types.Price) bool {
			count++
			return false
		})
	}
}

func BenchmarkProviderOperations(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	provider := types.OracleProvider{
		Address:               "stateset1provider",
		IsActive:              true,
		Slashed:               false,
		TotalSubmissions:      0,
		SuccessfulSubmissions: 0,
	}

	b.Run("SetProvider", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			provider.TotalSubmissions = uint64(i)
			k.SetProvider(ctx, provider)
		}
	})

	k.SetProvider(ctx, provider)

	b.Run("GetProvider", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetProvider(ctx, "stateset1provider")
		}
	})

	b.Run("IsAuthorizedProvider", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.IsAuthorizedProvider(ctx, "stateset1provider")
		}
	})
}

func BenchmarkPriceHistory(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	// Set params with history
	params := types.DefaultOracleParams()
	params.PriceHistorySize = 100
	k.SetParams(ctx, params)

	// Set initial price
	price := types.Price{
		Denom:       "stst",
		Amount:      sdkmath.LegacyMustNewDecFromStr("2.0"),
		LastUpdater: "stateset1updater",
		LastHeight:  100,
		UpdatedAt:   ctx.BlockTime(),
	}
	k.SetPrice(ctx, price)

	config := types.DefaultOracleConfig("stst")
	config.MinUpdateIntervalSeconds = 0 // Disable rate limiting for benchmark
	k.SetOracleConfig(ctx, config)

	b.Run("RecordPriceHistory", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newPrice := sdkmath.LegacyMustNewDecFromStr(fmt.Sprintf("2.%04d", i%10000))
			k.SetPriceWithValidation(ctx, "authority", "stst", newPrice)
		}
	})

	b.Run("GetPriceHistory", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetPriceHistory(ctx, "stst")
		}
	})
}

func BenchmarkOracleParamsOperations(b *testing.B) {
	k, ctx := setupBenchmarkOracleKeeper()

	params := types.DefaultOracleParams()

	b.Run("SetParams", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			params.PriceHistorySize = uint32(i % 1000)
			k.SetParams(ctx, params)
		}
	})

	b.Run("GetParams", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetParams(ctx)
		}
	})
}
