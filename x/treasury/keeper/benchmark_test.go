package keeper_test

import (
	"fmt"
	"testing"
	"time"

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

	"github.com/stateset/core/x/treasury/keeper"
	"github.com/stateset/core/x/treasury/types"
)

func setupBenchmarkTreasuryKeeper() (keeper.Keeper, sdk.Context, *benchTreasuryBankKeeper) {
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
		Height:  1,
	}, false, log.NewNopLogger())

	bankKeeper := newBenchTreasuryBankKeeper()
	accountKeeper := newBenchTreasuryAccountKeeper()

	treasuryAddr := newBenchTreasuryAddress()
	accountKeeper.SetModuleAddress(types.ModuleAccountName, treasuryAddr)
	bankKeeper.moduleAddresses[types.ModuleAccountName] = treasuryAddr

	// Fund treasury
	bankKeeper.SetModuleBalance(types.ModuleAccountName, sdk.NewCoins(
		sdk.NewInt64Coin("ustate", 1_000_000_000_000),
	))

	k := keeper.NewKeeper(cdc, storeKey, "authority", bankKeeper, accountKeeper)
	k.SetParams(ctx, types.DefaultTreasuryParams())

	return k, ctx, bankKeeper
}

func newBenchTreasuryAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

func BenchmarkCreateSpendProposal(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		recipient := newBenchTreasuryAddress()
		k.CreateSpendProposal(
			ctx,
			"authority",
			recipient.String(),
			sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000_000)),
			types.CategoryGrants,
			fmt.Sprintf("Grant proposal %d", i),
			24*time.Hour,
		)
	}
}

func BenchmarkGetSpendProposal(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	// Create a proposal first
	recipient := newBenchTreasuryAddress()
	proposalID, _ := k.CreateSpendProposal(
		ctx,
		"authority",
		recipient.String(),
		sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000_000)),
		types.CategoryGrants,
		"Test grant",
		24*time.Hour,
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetSpendProposal(ctx, proposalID)
	}
}

func BenchmarkExecuteSpendProposal(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	// Create proposals and advance time past timelock
	proposals := make([]uint64, b.N)
	for i := 0; i < b.N; i++ {
		recipient := newBenchTreasuryAddress()
		id, _ := k.CreateSpendProposal(
			ctx,
			"authority",
			recipient.String(),
			sdk.NewCoins(sdk.NewInt64Coin("ustate", 1000)),
			types.CategoryGrants,
			fmt.Sprintf("Grant %d", i),
			time.Hour, // Short timelock for benchmark
		)
		proposals[i] = id
	}

	// Advance time past timelock
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(2 * time.Hour))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.ExecuteSpendProposal(ctx, proposals[i])
	}
}

func BenchmarkIterateSpendProposals(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	// Create 100 proposals
	for i := 0; i < 100; i++ {
		recipient := newBenchTreasuryAddress()
		k.CreateSpendProposal(
			ctx,
			"authority",
			recipient.String(),
			sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000_000)),
			types.CategoryGrants,
			fmt.Sprintf("Grant %d", i),
			24*time.Hour,
		)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		k.IterateSpendProposals(ctx, func(p types.SpendProposal) bool {
			count++
			return false
		})
	}
}

func BenchmarkBudgetOperations(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	budget := types.Budget{
		Category:       types.CategoryGrants,
		TotalLimit:     sdk.NewCoins(sdk.NewInt64Coin("ustate", 100_000_000_000)),
		PeriodLimit:    sdk.NewCoins(sdk.NewInt64Coin("ustate", 10_000_000_000)),
		PeriodDuration: 30 * 24 * time.Hour,
		PeriodStart:    ctx.BlockTime(),
		Enabled:        true,
	}

	b.Run("SetBudget", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.SetBudget(ctx, budget)
		}
	})

	k.SetBudget(ctx, budget)

	b.Run("GetBudget", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetBudget(ctx, types.CategoryGrants)
		}
	})
}

func BenchmarkSnapshotOperations(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	snapshot := types.ReserveSnapshot{
		Reporter:     "stateset1reporter",
		TotalSupply:  sdk.NewInt64Coin("ssusd", 1_000_000_000),
		FiatReserves: sdk.NewInt64Coin("usd", 1_000_000_000),
		Timestamp:    ctx.BlockTime(),
		Metadata:     "Daily attestation",
	}

	b.Run("RecordSnapshot", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.RecordSnapshot(ctx, snapshot)
		}
	})

	b.Run("GetLatestSnapshot", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetLatestSnapshot(ctx)
		}
	})
}

func BenchmarkTreasuryParams(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	params := types.DefaultTreasuryParams()

	b.Run("SetParams", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.SetParams(ctx, params)
		}
	})

	b.Run("GetParams", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			k.GetParams(ctx)
		}
	})
}

func BenchmarkRecordRevenue(b *testing.B) {
	k, ctx, _ := setupBenchmarkTreasuryKeeper()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.RecordRevenue(ctx, "transaction_fees", sdk.NewCoins(sdk.NewInt64Coin("ustate", 1000)), "block_fees")
	}
}

// Mock keepers for benchmarks
type benchTreasuryBankKeeper struct {
	balances        map[string]sdk.Coins
	moduleBalances  map[string]sdk.Coins
	moduleAddresses map[string]sdk.AccAddress
}

func newBenchTreasuryBankKeeper() *benchTreasuryBankKeeper {
	return &benchTreasuryBankKeeper{
		balances:        make(map[string]sdk.Coins),
		moduleBalances:  make(map[string]sdk.Coins),
		moduleAddresses: make(map[string]sdk.AccAddress),
	}
}

func (m *benchTreasuryBankKeeper) GetBalance(_ context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.balances[addr.String()]
	return sdk.NewCoin(denom, coins.AmountOf(denom))
}

func (m *benchTreasuryBankKeeper) GetAllBalances(_ context, addr sdk.AccAddress) sdk.Coins {
	for module, moduleAddr := range m.moduleAddresses {
		if moduleAddr.Equals(addr) {
			return m.moduleBalances[module]
		}
	}
	return m.balances[addr.String()]
}

func (m *benchTreasuryBankKeeper) SetModuleBalance(module string, coins sdk.Coins) {
	m.moduleBalances[module] = coins.Sort()
}

func (m *benchTreasuryBankKeeper) SendCoinsFromModuleToAccount(_ context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[senderModule]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	newBalance, _ := moduleCoins.SafeSub(amt...)
	m.moduleBalances[senderModule] = newBalance.Sort()

	account := m.balances[recipientAddr.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	m.balances[recipientAddr.String()] = account.Add(amt...).Sort()
	return nil
}

func (m *benchTreasuryBankKeeper) SendCoinsFromAccountToModule(_ context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	return nil
}

func (m *benchTreasuryBankKeeper) BurnCoins(_ context, moduleName string, amt sdk.Coins) error {
	return nil
}

func (m *benchTreasuryBankKeeper) MintCoins(_ context, moduleName string, amt sdk.Coins) error {
	return nil
}

type benchTreasuryAccountKeeper struct {
	addresses map[string]sdk.AccAddress
}

func newBenchTreasuryAccountKeeper() *benchTreasuryAccountKeeper {
	return &benchTreasuryAccountKeeper{
		addresses: make(map[string]sdk.AccAddress),
	}
}

func (m *benchTreasuryAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return m.addresses[moduleName]
}

func (m *benchTreasuryAccountKeeper) SetModuleAddress(moduleName string, addr sdk.AccAddress) {
	m.addresses[moduleName] = addr
}

type context = sdk.Context
