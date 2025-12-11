package keeper_test

import (
	"sync"
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/oracle/keeper"
	oracletypes "github.com/stateset/core/x/oracle/types"
)

var msgCfgOnce sync.Once

func setupMsgConfig() {
	msgCfgOnce.Do(func() {
		config := sdk.GetConfig()
		config.SetBech32PrefixForAccount("stateset", "statesetpub")
		config.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		config.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		config.Seal()
	})
}

func setupMsgKeeper(t *testing.T) (keeper.Keeper, sdk.Context, string) {
	t.Helper()

	setupMsgConfig()

	storeKey := storetypes.NewKVStoreKey(oracletypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{ChainID: "stateset-test", Height: 42, Time: time.Now()}, false, log.NewNopLogger())

	authority := newMsgAddress().String()

	return keeper.NewKeeper(cdc, storeKey, authority), ctx, authority
}

func newMsgAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

func TestUpdatePrice(t *testing.T) {
	k, ctx, authority := setupMsgKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := oracletypes.NewMsgUpdatePrice(authority, "ustate", sdkmath.LegacyNewDec(12345))

	_, err := msgServer.UpdatePrice(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)

	price, found := k.GetPrice(ctx, "ustate")
	require.True(t, found)
	require.Equal(t, msg.Price, price.Amount)
	require.Equal(t, msg.Authority, price.LastUpdater)
	require.Equal(t, ctx.BlockHeight(), price.LastHeight)
}

func TestUpdatePriceUnauthorized(t *testing.T) {
	k, ctx, authority := setupMsgKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	other := newMsgAddress().String()
	require.NotEqual(t, authority, other)

	msg := oracletypes.NewMsgUpdatePrice(other, "ustate", sdkmath.LegacyNewDec(10))

	_, err := msgServer.UpdatePrice(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, oracletypes.ErrUnauthorized)
}

func TestUpdatePriceInvalidValue(t *testing.T) {
	k, ctx, authority := setupMsgKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := oracletypes.NewMsgUpdatePrice(authority, "ustate", sdkmath.LegacyZeroDec())

	_, err := msgServer.UpdatePrice(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, oracletypes.ErrInvalidPrice)
}
