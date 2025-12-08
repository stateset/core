package keeper_test

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	"cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/payments/keeper"
	paymentstypes "github.com/stateset/core/x/payments/types"
)

var paymentsConfigOnce sync.Once

func setupPaymentsConfig() {
	paymentsConfigOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount("stateset", "statesetpub")
		cfg.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		cfg.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		cfg.Seal()
	})
}

func setupPaymentsKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *mockBankKeeper, *mockComplianceKeeper) {
	t.Helper()

	setupPaymentsConfig()

	storeKey := types.NewKVStoreKey(paymentstypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, types.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	bankKeeper := newMockBankKeeper()
	complianceKeeper := newMockComplianceKeeper()

	k := keeper.NewKeeper(cdc, storeKey, bankKeeper, complianceKeeper, paymentstypes.ModuleAccountName)

	return k, ctx, bankKeeper, complianceKeeper
}

func newPaymentsAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

type mockBankKeeper struct {
	balances       map[string]sdk.Coins
	moduleBalances map[string]sdk.Coins
}

func newMockBankKeeper() *mockBankKeeper {
	return &mockBankKeeper{
		balances:       make(map[string]sdk.Coins),
		moduleBalances: make(map[string]sdk.Coins),
	}
}

func (m *mockBankKeeper) ensureAccount(addr string) sdk.Coins {
	if coins, ok := m.balances[addr]; ok {
		return coins
	}
	return sdk.Coins{}
}

func (m *mockBankKeeper) ensureModule(module string) sdk.Coins {
	if coins, ok := m.moduleBalances[module]; ok {
		return coins
	}
	return sdk.Coins{}
}

func (m *mockBankKeeper) SetBalance(addr sdk.AccAddress, coins sdk.Coins) {
	m.balances[addr.String()] = coins.Sort()
}

func (m *mockBankKeeper) Balance(addr sdk.AccAddress) sdk.Coins {
	return m.ensureAccount(addr.String())
}

func (m *mockBankKeeper) ModuleBalance(module string) sdk.Coins {
	return m.ensureModule(module)
}

func (m *mockBankKeeper) GetBalance(_ context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.ensureAccount(addr.String())
	amount := coins.AmountOf(denom)
	return sdk.NewCoin(denom, amount)
}

func (m *mockBankKeeper) SendCoinsFromAccountToModule(_ context.Context, sender sdk.AccAddress, module string, amt sdk.Coins) error {
	if !m.ensureAccount(sender.String()).IsAllGTE(amt) {
		return errors.New("insufficient funds")
	}
	m.balances[sender.String()] = m.ensureAccount(sender.String()).Sub(amt...)
	m.moduleBalances[module] = m.ensureModule(module).Add(amt...)
	return nil
}

func (m *mockBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, module string, recipient sdk.AccAddress, amt sdk.Coins) error {
	if !m.ensureModule(module).IsAllGTE(amt) {
		return errors.New("module insufficient funds")
	}
	m.moduleBalances[module] = m.ensureModule(module).Sub(amt...)
	m.balances[recipient.String()] = m.ensureAccount(recipient.String()).Add(amt...)
	return nil
}

type mockComplianceKeeper struct {
	blocked map[string]bool
}

func newMockComplianceKeeper() *mockComplianceKeeper {
	return &mockComplianceKeeper{
		blocked: make(map[string]bool),
	}
}

func (m *mockComplianceKeeper) Block(addr sdk.AccAddress) {
	m.blocked[addr.String()] = true
}

func (m *mockComplianceKeeper) AssertCompliant(_ context.Context, addr sdk.AccAddress) error {
	if m.blocked[addr.String()] {
		return errors.New("address not compliant")
	}
	return nil
}

func TestMsgCreatePayment(t *testing.T) {
	k, ctx, bank, _ := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	// Set generous balance and ensure compliance passes.
	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000)))
	// leave compliance unblocked by default

	msg := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 250), "invoice-42")
	resp, err := msgServer.CreatePayment(ctx, msg)
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.PaymentId)

	stored, found := k.GetPayment(ctx, resp.PaymentId)
	require.True(t, found)
	require.Equal(t, paymentstypes.PaymentStatusPending, stored.Status)
	require.Equal(t, msg.Metadata, stored.Metadata)
	require.Equal(t, msg.Amount, stored.Amount)
	require.Equal(t, payer.String(), stored.Payer)
	require.Equal(t, payee.String(), stored.Payee)

	require.Equal(t, sdk.NewCoins(msg.Amount), bank.ModuleBalance(paymentstypes.ModuleAccountName))
	require.Equal(t, sdk.NewCoins(sdk.NewInt64Coin("ustate", 750)), bank.Balance(payer))
}

func TestMsgCreatePaymentInsufficientBalance(t *testing.T) {
	k, ctx, _, _ := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	msg := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 10), "insufficient")
	_, err := msgServer.CreatePayment(ctx, msg)
	require.Error(t, err)
	require.ErrorIs(t, err, paymentstypes.ErrInsufficientBalance)
}

func TestMsgSettlePayment(t *testing.T) {
	k, ctx, bank, _ := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 1_000)))

	create := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 300), "order-1")
	resp, err := msgServer.CreatePayment(ctx, create)
	require.NoError(t, err)

	settle := paymentstypes.NewMsgSettlePayment(payee.String(), resp.PaymentId)
	_, err = msgServer.SettlePayment(ctx, settle)
	require.NoError(t, err)

	stored, found := k.GetPayment(ctx, resp.PaymentId)
	require.True(t, found)
	require.Equal(t, paymentstypes.PaymentStatusSettled, stored.Status)

	require.Equal(t, sdk.NewCoins(), bank.ModuleBalance(paymentstypes.ModuleAccountName))
	require.Equal(t, sdk.NewCoins(sdk.NewInt64Coin("ustate", 700)), bank.Balance(payer))
	require.Equal(t, sdk.NewCoins(sdk.NewInt64Coin("ustate", 300)), bank.Balance(payee))
}

func TestMsgSettlePaymentWrongPayee(t *testing.T) {
	k, ctx, bank, _ := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()
	other := newPaymentsAddress()

	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 400)))

	create := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 200), "order-2")
	resp, err := msgServer.CreatePayment(ctx, create)
	require.NoError(t, err)

	settle := paymentstypes.NewMsgSettlePayment(other.String(), resp.PaymentId)
	_, err = msgServer.SettlePayment(ctx, settle)
	require.Error(t, err)
	require.ErrorIs(t, err, paymentstypes.ErrNotAuthorized)

	stored, found := k.GetPayment(ctx, resp.PaymentId)
	require.True(t, found)
	require.Equal(t, paymentstypes.PaymentStatusPending, stored.Status)
}

func TestMsgCancelPayment(t *testing.T) {
	k, ctx, bank, _ := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 500)))

	create := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 200), "order-3")
	resp, err := msgServer.CreatePayment(ctx, create)
	require.NoError(t, err)

	cancel := paymentstypes.NewMsgCancelPayment(payer.String(), resp.PaymentId, "buyer request")
	_, err = msgServer.CancelPayment(ctx, cancel)
	require.NoError(t, err)

	stored, found := k.GetPayment(ctx, resp.PaymentId)
	require.True(t, found)
	require.Equal(t, paymentstypes.PaymentStatusCancelled, stored.Status)

	require.Equal(t, sdk.NewCoins(), bank.ModuleBalance(paymentstypes.ModuleAccountName))
	require.Equal(t, sdk.NewCoins(sdk.NewInt64Coin("ustate", 500)), bank.Balance(payer))
	require.Equal(t, sdk.Coins{}, bank.Balance(payee))
}

func TestMsgCreatePaymentComplianceBlocked(t *testing.T) {
	k, ctx, bank, compliance := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 100)))
	compliance.Block(payee)

	msg := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 50), "blocked")
	_, err := msgServer.CreatePayment(ctx, msg)
	require.Error(t, err)
}

func TestMsgSettlePaymentComplianceBlocked(t *testing.T) {
	k, ctx, bank, compliance := setupPaymentsKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	payer := newPaymentsAddress()
	payee := newPaymentsAddress()

	bank.SetBalance(payer, sdk.NewCoins(sdk.NewInt64Coin("ustate", 250)))

	create := paymentstypes.NewMsgCreatePayment(payer.String(), payee.String(), sdk.NewInt64Coin("ustate", 100), "order-4")
	resp, err := msgServer.CreatePayment(ctx, create)
	require.NoError(t, err)

	compliance.Block(payee)

	settle := paymentstypes.NewMsgSettlePayment(payee.String(), resp.PaymentId)
	_, err = msgServer.SettlePayment(ctx, settle)
	require.Error(t, err)
}
