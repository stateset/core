package keeper_test

import (
	"sync"
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

func setupKeeper(t *testing.T) (keeper.Keeper, sdk.Context, string) {
	t.Helper()

	setupConfig()

	storeKey := storetypes.NewKVStoreKey(treasurytypes.StoreKey)
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := testutil.DefaultContext(storeKey, storetypes.NewTransientStoreKey("treasury-transient"))
	ctx = ctx.WithBlockTime(time.Now().UTC())

	authority := newAddress().String()

	return keeper.NewKeeper(cdc, storeKey, authority), ctx, authority
}

func newAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

func TestRecordReserve(t *testing.T) {
	k, ctx, authority := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	reporter := newAddress().String()
	snapshot := treasurytypes.ReserveSnapshot{
		Reporter:     reporter,
		TotalSupply:  sdk.NewInt64Coin("ssusd", 1_000_000),
		FiatReserves: sdk.NewInt64Coin("usd", 1_200_000),
		OtherReserves: sdk.NewCoins(
			sdk.NewInt64Coin("eth", 10),
		),
		Metadata: "monthly report",
	}

	msg := treasurytypes.NewMsgRecordReserve(authority, snapshot)
	resp, err := msgServer.RecordReserve(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.SnapshotID)

	stored, found := k.GetSnapshot(sdk.WrapSDKContext(ctx), resp.SnapshotID)
	require.True(t, found)
	require.Equal(t, snapshot.Reporter, stored.Reporter)
	require.Equal(t, snapshot.TotalSupply, stored.TotalSupply)
	require.Equal(t, snapshot.FiatReserves, stored.FiatReserves)
	require.Equal(t, snapshot.OtherReserves, stored.OtherReserves)
	require.False(t, stored.Timestamp.IsZero())
}

func TestRecordReserveUnauthorized(t *testing.T) {
	k, ctx, authority := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	require.NotEqual(t, authority, newAddress().String())

	msg := treasurytypes.NewMsgRecordReserve(newAddress().String(), treasurytypes.ReserveSnapshot{
		Reporter:     newAddress().String(),
		TotalSupply:  sdk.NewInt64Coin("ssusd", 1),
		FiatReserves: sdk.NewInt64Coin("usd", 1),
		OtherReserves: sdk.NewCoins(
			sdk.NewInt64Coin("eth", 1),
		),
	})

	_, err := msgServer.RecordReserve(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, treasurytypes.ErrUnauthorized)
}

func TestRecordReserveInvalidSnapshot(t *testing.T) {
	k, ctx, authority := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := treasurytypes.NewMsgRecordReserve(authority, treasurytypes.ReserveSnapshot{
		Reporter:     "invalid",
		TotalSupply:  sdk.NewInt64Coin("ssusd", 1),
		FiatReserves: sdk.NewInt64Coin("usd", 1),
	})

	_, err := msgServer.RecordReserve(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, treasurytypes.ErrInvalidReserve)
}
