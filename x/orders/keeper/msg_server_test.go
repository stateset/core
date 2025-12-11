package keeper_test

import (
	"context"
	"errors"
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

	"github.com/stateset/core/x/orders/keeper"
	ordertypes "github.com/stateset/core/x/orders/types"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

var ordersConfigOnce sync.Once

func setupOrdersConfig() {
	ordersConfigOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount("stateset", "statesetpub")
		cfg.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		cfg.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		cfg.Seal()
	})
}

func setupOrdersKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *mockSettlementKeeper) {
	t.Helper()
	setupOrdersConfig()

	storeKey := storetypes.NewKVStoreKey(ordertypes.StoreKey)
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	ctx := sdk.NewContext(stateStore, cmtproto.Header{ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	bankKeeper := newMockBankKeeper()
	complianceKeeper := newMockComplianceKeeper()
	settlementKeeper := newMockSettlementKeeper()
	accountKeeper := newMockAccountKeeper()

	k := keeper.NewKeeper(cdc, storeKey, "stateset1authority", bankKeeper, complianceKeeper, settlementKeeper, accountKeeper)
	return k, ctx, settlementKeeper
}

func newOrdersAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

type mockBankKeeper struct{}

func newMockBankKeeper() *mockBankKeeper { return &mockBankKeeper{} }

func (m *mockBankKeeper) GetBalance(_ context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	return sdk.NewCoin(denom, sdkmath.ZeroInt())
}

func (m *mockBankKeeper) SendCoins(_ context.Context, _, _ sdk.AccAddress, _ sdk.Coins) error {
	return nil
}

func (m *mockBankKeeper) SendCoinsFromAccountToModule(_ context.Context, _ sdk.AccAddress, _ string, _ sdk.Coins) error {
	return nil
}

func (m *mockBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, _ string, _ sdk.AccAddress, _ sdk.Coins) error {
	return nil
}

type mockComplianceKeeper struct {
	blocked map[string]bool
}

func newMockComplianceKeeper() *mockComplianceKeeper {
	return &mockComplianceKeeper{blocked: make(map[string]bool)}
}

func (m *mockComplianceKeeper) Block(addr sdk.AccAddress) { m.blocked[addr.String()] = true }

func (m *mockComplianceKeeper) AssertCompliant(_ context.Context, addr sdk.AccAddress) error {
	if m.blocked[addr.String()] {
		return errors.New("blocked")
	}
	return nil
}

type mockSettlementKeeper struct {
	lastMethod    string
	lastSender    string
	lastRecipient string
	lastAmount    sdk.Coin
	nextID        uint64
}

func newMockSettlementKeeper() *mockSettlementKeeper { return &mockSettlementKeeper{nextID: 7} }

func (m *mockSettlementKeeper) InstantTransfer(_ sdk.Context, sender, recipient string, amount sdk.Coin, _, _ string) (uint64, error) {
	m.lastMethod = "instant"
	m.lastSender = sender
	m.lastRecipient = recipient
	m.lastAmount = amount
	return m.nextID, nil
}

func (m *mockSettlementKeeper) CreateEscrow(_ sdk.Context, sender, recipient string, amount sdk.Coin, _, _ string, _ int64) (uint64, error) {
	m.lastMethod = "escrow"
	m.lastSender = sender
	m.lastRecipient = recipient
	m.lastAmount = amount
	return m.nextID, nil
}

func (m *mockSettlementKeeper) ReleaseEscrow(_ sdk.Context, _ uint64, _ sdk.AccAddress) error {
	return nil
}

func (m *mockSettlementKeeper) RefundEscrow(_ sdk.Context, _ uint64, _ sdk.AccAddress, _ string) error {
	return nil
}

func (m *mockSettlementKeeper) GetMerchant(_ sdk.Context, _ string) (interface{}, bool) {
	return nil, false
}

type mockAccountKeeper struct{}

func newMockAccountKeeper() *mockAccountKeeper { return &mockAccountKeeper{} }

func (m *mockAccountKeeper) GetModuleAddress(_ string) sdk.AccAddress { return nil }

func TestMsgCreateOrder(t *testing.T) {
	k, ctx, _ := setupOrdersKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	customer := newOrdersAddress()
	merchant := newOrdersAddress()

	items := []ordertypes.OrderItem{
		{
			Id:          "1",
			ProductId:   "sku-1",
			ProductName: "Widget",
			Quantity:    2,
			UnitPrice:   sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 100),
		},
	}

	msg := ordertypes.NewMsgCreateOrder(customer.String(), merchant.String(), items, ordertypes.ShippingInfo{}, "meta")
	resp, err := msgServer.CreateOrder(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.OrderId)

	order, found := k.GetOrder(ctx, resp.OrderId)
	require.True(t, found)
	require.Equal(t, ordertypes.OrderStatusPending, order.Status)
	require.Equal(t, customer.String(), order.Customer)
	require.Equal(t, merchant.String(), order.Merchant)
}

func TestMsgPayOrder_UsesInstantTransfer(t *testing.T) {
	k, ctx, settlement := setupOrdersKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	customer := newOrdersAddress()
	merchant := newOrdersAddress()

	items := []ordertypes.OrderItem{
		{
			Id:          "1",
			ProductId:   "sku-1",
			ProductName: "Widget",
			Quantity:    1,
			UnitPrice:   sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 500),
		},
	}

	createMsg := ordertypes.NewMsgCreateOrder(customer.String(), merchant.String(), items, ordertypes.ShippingInfo{}, "")
	resp, err := msgServer.CreateOrder(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(t, err)

	confirmMsg := ordertypes.NewMsgConfirmOrder(merchant.String(), resp.OrderId)
	_, err = msgServer.ConfirmOrder(sdk.WrapSDKContext(ctx), confirmMsg)
	require.NoError(t, err)

	payMsg := ordertypes.NewMsgPayOrder(customer.String(), resp.OrderId, sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 500), false)
	_, err = msgServer.PayOrder(sdk.WrapSDKContext(ctx), payMsg)
	require.NoError(t, err)

	order, found := k.GetOrder(ctx, resp.OrderId)
	require.True(t, found)
	require.Equal(t, ordertypes.OrderStatusPaid, order.Status)
	require.Equal(t, "instant", order.PaymentInfo.Method)
	require.Equal(t, settlement.nextID, order.SettlementId)
	require.Equal(t, "instant", settlement.lastMethod)
}
