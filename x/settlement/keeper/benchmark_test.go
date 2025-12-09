package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/stateset/core/x/settlement/keeper"
	"github.com/stateset/core/x/settlement/types"
)

func setupBenchmarkSettlementKeeper() (keeper.Keeper, sdk.Context, *mockBankKeeper) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	bankKeeper := newMockBankKeeper()
	complianceKeeper := newMockComplianceKeeper()
	accountKeeper := newMockAccountKeeper()

	authority := newSettlementAddress()

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		bankKeeper,
		complianceKeeper,
		accountKeeper,
		authority.String(),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-bench", Time: time.Now()}, false, log.NewNopLogger())

	// Initialize genesis
	k.InitGenesis(ctx, types.DefaultGenesis())

	return k, ctx, bankKeeper
}

func BenchmarkInstantTransfer(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender with large balance
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF"+string(rune(i)), "")
	}
}

func BenchmarkCreateEscrow(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender with large balance
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESC"+string(rune(i)), "", 86400)
	}
}

func BenchmarkOpenChannel(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender with large balance
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	}
}

func BenchmarkGetSettlement(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Pre-populate with settlements
	for i := 0; i < 1000; i++ {
		k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "", "")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetSettlement(ctx, uint64((i%1000)+1))
	}
}

func BenchmarkGetChannel(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Pre-populate with channels
	for i := 0; i < 1000; i++ {
		k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetChannel(ctx, uint64((i%1000)+1))
	}
}

func BenchmarkIterateSettlements(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000_000_000))))

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Pre-populate with 100 settlements
	for i := 0; i < 100; i++ {
		k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "", "")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		k.IterateSettlements(ctx, func(s types.Settlement) bool {
			count++
			return false
		})
	}
}

func BenchmarkCreateBatch(b *testing.B) {
	k, ctx, bankKeeper := setupBenchmarkSettlementKeeper()

	merchant := newSettlementAddress()

	// Create senders
	senders := make([]string, 10)
	amounts := make([]sdk.Coin, 10)
	references := make([]string, 10)

	for i := 0; i < 10; i++ {
		sender := newSettlementAddress()
		senders[i] = sender.String()
		amounts[i] = sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
		references[i] = "REF" + string(rune('0'+i))
		bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1_000_000))))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
	}
}

func BenchmarkRegisterMerchant(b *testing.B) {
	k, ctx, _ := setupBenchmarkSettlementKeeper()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merchantAddr := newSettlementAddress()
		config := types.MerchantConfig{
			Address:      merchantAddr.String(),
			Name:         "Merchant " + string(rune(i)),
			FeeRateBps:   25,
			BatchEnabled: true,
		}
		k.RegisterMerchant(ctx, config)
	}
}

func BenchmarkGetMerchant(b *testing.B) {
	k, ctx, _ := setupBenchmarkSettlementKeeper()

	// Pre-populate with merchants
	merchants := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		merchantAddr := newSettlementAddress()
		merchants[i] = merchantAddr.String()
		config := types.MerchantConfig{
			Address: merchantAddr.String(),
			Name:    "Merchant " + string(rune(i)),
		}
		k.RegisterMerchant(ctx, config)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetMerchant(ctx, merchants[i%1000])
	}
}
