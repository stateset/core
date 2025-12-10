package keeper_test

import (
	"context"
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

// Mock BankKeeper for msg server tests
type msgServerMockBankKeeper struct {
	balances       map[string]sdk.Coins
	moduleBalances map[string]sdk.Coins
}

func newMsgServerMockBankKeeper() *msgServerMockBankKeeper {
	return &msgServerMockBankKeeper{
		balances:       make(map[string]sdk.Coins),
		moduleBalances: make(map[string]sdk.Coins),
	}
}

func (m *msgServerMockBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.balances[addr.String()]
	return sdk.NewCoin(denom, coins.AmountOf(denom))
}

func (m *msgServerMockBankKeeper) GetAllBalances(ctx context.Context, addr sdk.AccAddress) sdk.Coins {
	return m.balances[addr.String()]
}

func (m *msgServerMockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	m.moduleBalances[senderModule] = m.moduleBalances[senderModule].Sub(amt...)
	m.balances[recipientAddr.String()] = m.balances[recipientAddr.String()].Add(amt...)
	return nil
}

func (m *msgServerMockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	m.balances[senderAddr.String()] = m.balances[senderAddr.String()].Sub(amt...)
	m.moduleBalances[recipientModule] = m.moduleBalances[recipientModule].Add(amt...)
	return nil
}

func (m *msgServerMockBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	m.moduleBalances[moduleName] = m.moduleBalances[moduleName].Sub(amt...)
	return nil
}

func (m *msgServerMockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	m.moduleBalances[moduleName] = m.moduleBalances[moduleName].Add(amt...)
	return nil
}

// Mock AccountKeeper for msg server tests
type msgServerMockAccountKeeper struct {
	moduleAddresses map[string]sdk.AccAddress
}

func newMsgServerMockAccountKeeper() *msgServerMockAccountKeeper {
	return &msgServerMockAccountKeeper{
		moduleAddresses: make(map[string]sdk.AccAddress),
	}
}

func (m *msgServerMockAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	if addr, ok := m.moduleAddresses[moduleName]; ok {
		return addr
	}
	// Generate deterministic address for module
	return sdk.AccAddress([]byte(moduleName))
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
	bankKeeper := newMsgServerMockBankKeeper()
	accountKeeper := newMsgServerMockAccountKeeper()

	return keeper.NewKeeper(cdc, storeKey, authority, bankKeeper, accountKeeper), ctx, authority
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
