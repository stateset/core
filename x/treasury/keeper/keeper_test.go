package keeper_test

import (
	"testing"
	"time"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	testutil "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/treasury/keeper"
	treasurytypes "github.com/stateset/core/x/treasury/types"
)

func setupTreasuryKeeperTest(t *testing.T) (keeper.Keeper, sdk.Context, string) {
	t.Helper()

	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("treasury-transient"))
	ctx = ctx.WithBlockTime(time.Now().UTC())

	authority := newTreasuryTestAddress().String()

	return keeper.NewKeeper(cdc, storeKey, authority), ctx, authority
}

func newTreasuryTestAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

func TestRecordSnapshot(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	snapshot := treasurytypes.ReserveSnapshot{
		Reporter:     newTreasuryTestAddress().String(),
		TotalSupply:  sdk.NewInt64Coin("ssusd", 1_000_000),
		FiatReserves: sdk.NewInt64Coin("usd", 1_200_000),
		OtherReserves: sdk.NewCoins(
			sdk.NewInt64Coin("eth", 10),
			sdk.NewInt64Coin("btc", 5),
		),
		Metadata: "monthly report",
	}

	id := k.RecordSnapshot(ctx, snapshot)
	require.Equal(t, uint64(1), id)

	// Verify stored snapshot
	stored, found := k.GetSnapshot(ctx, id)
	require.True(t, found)
	require.Equal(t, uint64(1), stored.Id)
	require.Equal(t, snapshot.Reporter, stored.Reporter)
	require.Equal(t, snapshot.TotalSupply, stored.TotalSupply)
	require.Equal(t, snapshot.FiatReserves, stored.FiatReserves)
	require.Equal(t, snapshot.OtherReserves, stored.OtherReserves)
	require.Equal(t, snapshot.Metadata, stored.Metadata)
	require.False(t, stored.Timestamp.IsZero())
}

func TestRecordMultipleSnapshots(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	for i := 1; i <= 5; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:     newTreasuryTestAddress().String(),
			TotalSupply:  sdk.NewInt64Coin("ssusd", int64(i*1_000_000)),
			FiatReserves: sdk.NewInt64Coin("usd", int64(i*1_100_000)),
			OtherReserves: sdk.NewCoins(
				sdk.NewInt64Coin("eth", int64(i*10)),
			),
			Metadata: "report-" + string(rune('0'+i)),
		}

		id := k.RecordSnapshot(ctx, snapshot)
		require.Equal(t, uint64(i), id)
	}

	// Verify all snapshots exist
	for i := uint64(1); i <= 5; i++ {
		stored, found := k.GetSnapshot(ctx, i)
		require.True(t, found)
		require.Equal(t, i, stored.Id)
	}
}

func TestGetSnapshotNotFound(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	_, found := k.GetSnapshot(ctx, 999)
	require.False(t, found)
}

func TestGetLatestSnapshot(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	// No snapshots yet
	_, found := k.GetLatestSnapshot(ctx)
	require.False(t, found)

	// Add multiple snapshots
	for i := 1; i <= 3; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      newTreasuryTestAddress().String(),
			TotalSupply:   sdk.NewInt64Coin("ssusd", int64(i*1_000_000)),
			FiatReserves:  sdk.NewInt64Coin("usd", int64(i*1_000_000)),
			OtherReserves: sdk.NewCoins(),
			Metadata:      "report-" + string(rune('0'+i)),
		}
		k.RecordSnapshot(ctx, snapshot)
	}

	// Get latest - should be the last one
	latest, found := k.GetLatestSnapshot(ctx)
	require.True(t, found)
	require.Equal(t, uint64(3), latest.Id)
	require.Equal(t, "report-3", latest.Metadata)
}

func TestIterateSnapshots(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	// Add snapshots
	for i := 1; i <= 5; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      newTreasuryTestAddress().String(),
			TotalSupply:   sdk.NewInt64Coin("ssusd", int64(i*1_000_000)),
			FiatReserves:  sdk.NewInt64Coin("usd", int64(i*1_000_000)),
			OtherReserves: sdk.NewCoins(),
		}
		k.RecordSnapshot(ctx, snapshot)
	}

	// Iterate and count
	count := 0
	k.IterateSnapshots(ctx, func(snapshot treasurytypes.ReserveSnapshot) bool {
		count++
		return false
	})
	require.Equal(t, 5, count)

	// Test early termination
	count = 0
	k.IterateSnapshots(ctx, func(snapshot treasurytypes.ReserveSnapshot) bool {
		count++
		return count >= 3 // Stop after 3
	})
	require.Equal(t, 3, count)
}

func TestGetAuthority(t *testing.T) {
	k, _, authority := setupTreasuryKeeperTest(t)
	require.Equal(t, authority, k.GetAuthority())
}

func TestSetAuthority(t *testing.T) {
	k, _, _ := setupTreasuryKeeperTest(t)

	newAuthority := newTreasuryTestAddress().String()
	k.SetAuthority(newAuthority)
	require.Equal(t, newAuthority, k.GetAuthority())
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx, authority := setupTreasuryKeeperTest(t)

	// Add some snapshots
	for i := 1; i <= 3; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      newTreasuryTestAddress().String(),
			TotalSupply:   sdk.NewInt64Coin("ssusd", int64(i*1_000_000)),
			FiatReserves:  sdk.NewInt64Coin("usd", int64(i*1_000_000)),
			OtherReserves: sdk.NewCoins(),
			Metadata:      "export-test-" + string(rune('0'+i)),
		}
		k.RecordSnapshot(ctx, snapshot)
	}

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Equal(t, authority, genesis.Authority)
	require.Equal(t, uint64(4), genesis.NextID) // After 3 snapshots, next ID is 4
	require.Len(t, genesis.Snapshots, 3)

	// Create new keeper and import (use same authority from genesis)
	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey + "-import")
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	ctx2 := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("treasury-transient-import"))
	ctx2 = ctx2.WithBlockTime(time.Now().UTC())
	k2 := keeper.NewKeeper(cdc, storeKey, authority)
	k2.InitGenesis(ctx2, genesis)

	// Verify state was imported
	require.Equal(t, authority, k2.GetAuthority())

	for i := uint64(1); i <= 3; i++ {
		stored, found := k2.GetSnapshot(ctx2, i)
		require.True(t, found)
		require.Equal(t, i, stored.Id)
	}

	// Verify next ID
	newSnapshot := treasurytypes.ReserveSnapshot{
		Reporter:      newTreasuryTestAddress().String(),
		TotalSupply:   sdk.NewInt64Coin("ssusd", 4_000_000),
		FiatReserves:  sdk.NewInt64Coin("usd", 4_000_000),
		OtherReserves: sdk.NewCoins(),
	}
	id := k2.RecordSnapshot(ctx2, newSnapshot)
	require.Equal(t, uint64(4), id)
}

func TestInitGenesisWithNilState(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	// Should not panic with nil state
	k.InitGenesis(ctx, nil)

	// Should use defaults
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
}

func TestSnapshotTimestamp(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	// Snapshot without explicit timestamp should use block time
	snapshot := treasurytypes.ReserveSnapshot{
		Reporter:      newTreasuryTestAddress().String(),
		TotalSupply:   sdk.NewInt64Coin("ssusd", 1_000_000),
		FiatReserves:  sdk.NewInt64Coin("usd", 1_000_000),
		OtherReserves: sdk.NewCoins(),
	}

	id := k.RecordSnapshot(ctx, snapshot)
	stored, _ := k.GetSnapshot(ctx, id)
	require.False(t, stored.Timestamp.IsZero())

	// Snapshot with explicit timestamp should preserve it
	explicitTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	snapshotWithTime := treasurytypes.ReserveSnapshot{
		Reporter:      newTreasuryTestAddress().String(),
		TotalSupply:   sdk.NewInt64Coin("ssusd", 2_000_000),
		FiatReserves:  sdk.NewInt64Coin("usd", 2_000_000),
		OtherReserves: sdk.NewCoins(),
		Timestamp:     explicitTime,
	}

	id2 := k.RecordSnapshot(ctx, snapshotWithTime)
	stored2, _ := k.GetSnapshot(ctx, id2)
	require.Equal(t, explicitTime.Unix(), stored2.Timestamp.Unix())
}

func TestOtherReservesMultipleDenoms(t *testing.T) {
	k, ctx, _ := setupTreasuryKeeperTest(t)

	snapshot := treasurytypes.ReserveSnapshot{
		Reporter:     newTreasuryTestAddress().String(),
		TotalSupply:  sdk.NewInt64Coin("ssusd", 10_000_000),
		FiatReserves: sdk.NewInt64Coin("usd", 8_000_000),
		OtherReserves: sdk.NewCoins(
			sdk.NewInt64Coin("eth", 100),
			sdk.NewInt64Coin("btc", 50),
			sdk.NewInt64Coin("usdc", 1_000_000),
			sdk.NewInt64Coin("usdt", 500_000),
		),
		Metadata: "diversified reserves",
	}

	id := k.RecordSnapshot(ctx, snapshot)
	stored, found := k.GetSnapshot(ctx, id)
	require.True(t, found)
	require.Len(t, stored.OtherReserves, 4)
	require.Equal(t, int64(100), stored.OtherReserves.AmountOf("eth").Int64())
	require.Equal(t, int64(50), stored.OtherReserves.AmountOf("btc").Int64())
}

// Benchmark tests
func BenchmarkRecordSnapshot(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("test"))
	ctx = ctx.WithBlockTime(time.Now())

	k := keeper.NewKeeper(cdc, storeKey, "authority")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      "stateset1reporter",
			TotalSupply:   sdk.NewInt64Coin("ssusd", int64(i*1_000_000)),
			FiatReserves:  sdk.NewInt64Coin("usd", int64(i*1_000_000)),
			OtherReserves: sdk.NewCoins(sdk.NewInt64Coin("eth", 100)),
		}
		k.RecordSnapshot(ctx, snapshot)
	}
}

func BenchmarkGetSnapshot(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("test"))
	ctx = ctx.WithBlockTime(time.Now())

	k := keeper.NewKeeper(cdc, storeKey, "authority")

	// Pre-populate with snapshots
	for i := 0; i < 1000; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      "stateset1reporter",
			TotalSupply:   sdk.NewInt64Coin("ssusd", 1_000_000),
			FiatReserves:  sdk.NewInt64Coin("usd", 1_000_000),
			OtherReserves: sdk.NewCoins(),
		}
		k.RecordSnapshot(ctx, snapshot)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetSnapshot(ctx, uint64((i%1000)+1))
	}
}

func BenchmarkGetLatestSnapshot(b *testing.B) {
	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("test"))
	ctx = ctx.WithBlockTime(time.Now())

	k := keeper.NewKeeper(cdc, storeKey, "authority")

	// Pre-populate with snapshots
	for i := 0; i < 100; i++ {
		snapshot := treasurytypes.ReserveSnapshot{
			Reporter:      "stateset1reporter",
			TotalSupply:   sdk.NewInt64Coin("ssusd", 1_000_000),
			FiatReserves:  sdk.NewInt64Coin("usd", 1_000_000),
			OtherReserves: sdk.NewCoins(),
		}
		k.RecordSnapshot(ctx, snapshot)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetLatestSnapshot(ctx)
	}
}
