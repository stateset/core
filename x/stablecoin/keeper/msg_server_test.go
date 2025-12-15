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

	"github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
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

func setupKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *mockBankKeeper, *mockOracleKeeper, *mockComplianceKeeper) {
	t.Helper()

	setupConfig()

	storeKey := storetypes.NewKVStoreKey(stablecointypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	bankKeeper := newMockBankKeeper()
	oracleKeeper := newMockOracleKeeper()
	complianceKeeper := newMockComplianceKeeper()
	accountKeeper := newMockAccountKeeper()
	accountKeeper.SetAddress(stablecointypes.ModuleAccountName, newAddress())

	authority := newAddress()
	k := keeper.NewKeeper(cdc, storeKey, authority.String(), bankKeeper, accountKeeper, oracleKeeper, complianceKeeper)
	// Enable vault minting for keeper-level vault tests.
	params := stablecointypes.DefaultParams()
	params.VaultMintingEnabled = true
	// Add test collateral "stst" to supported collaterals
	params.CollateralParams = append(params.CollateralParams, stablecointypes.CollateralParam{
		Denom:            "stst",
		LiquidationRatio: sdkmath.LegacyMustNewDecFromStr("1.5"),
		StabilityFee:     sdkmath.LegacyMustNewDecFromStr("0.01"),
		DebtLimit:        sdkmath.NewInt(100_000_000_000_000),
		Active:           true,
	})
	k.SetParams(ctx, params)

	return k, ctx, bankKeeper, oracleKeeper, complianceKeeper
}

func newAddress() sdk.AccAddress {
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
	return sdk.NewCoins()
}

func (m *mockBankKeeper) ensureModule(module string) sdk.Coins {
	if coins, ok := m.moduleBalances[module]; ok {
		return coins
	}
	return sdk.NewCoins()
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

func (m *mockBankKeeper) SendCoinsFromAccountToModule(_ context.Context, sender sdk.AccAddress, module string, amt sdk.Coins) error {
	account := m.ensureAccount(sender.String())
	if !account.IsAllGTE(amt) {
		return errors.New("insufficient funds")
	}
	account = account.Sub(amt...).Sort()
	m.balances[sender.String()] = account

	moduleCoins := m.ensureModule(module)
	moduleCoins = moduleCoins.Add(amt...).Sort()
	m.moduleBalances[module] = moduleCoins
	return nil
}

func (m *mockBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, module string, recipient sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.ensureModule(module)
	if !moduleCoins.IsAllGTE(amt) {
		return errors.New("module insufficient funds")
	}
	moduleCoins = moduleCoins.Sub(amt...).Sort()
	m.moduleBalances[module] = moduleCoins

	account := m.ensureAccount(recipient.String())
	account = account.Add(amt...).Sort()
	m.balances[recipient.String()] = account
	return nil
}

func (m *mockBankKeeper) MintCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.ensureModule(module)
	moduleCoins = moduleCoins.Add(amt...).Sort()
	m.moduleBalances[module] = moduleCoins
	return nil
}

func (m *mockBankKeeper) BurnCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.ensureModule(module)
	if !moduleCoins.IsAllGTE(amt) {
		return errors.New("module insufficient funds to burn")
	}
	moduleCoins = moduleCoins.Sub(amt...).Sort()
	m.moduleBalances[module] = moduleCoins
	return nil
}

func (m *mockBankKeeper) GetSupply(_ context.Context, denom string) sdk.Coin {
	total := sdk.NewCoin(denom, sdkmath.ZeroInt())
	for _, coins := range m.balances {
		total = total.Add(sdk.NewCoin(denom, coins.AmountOf(denom)))
	}
	for _, coins := range m.moduleBalances {
		total = total.Add(sdk.NewCoin(denom, coins.AmountOf(denom)))
	}
	return total
}

type mockAccountKeeper struct {
	addresses map[string]sdk.AccAddress
}

func newMockAccountKeeper() *mockAccountKeeper {
	return &mockAccountKeeper{
		addresses: make(map[string]sdk.AccAddress),
	}
}

func (m *mockAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return m.addresses[moduleName]
}

func (m *mockAccountKeeper) SetModuleAccount(_ context.Context, macc sdk.ModuleAccountI) {
	m.addresses[macc.GetName()] = macc.GetAddress()
}

func (m *mockAccountKeeper) SetAddress(moduleName string, addr sdk.AccAddress) {
	m.addresses[moduleName] = addr
}

type mockOracleKeeper struct {
	prices map[string]sdkmath.LegacyDec
}

func newMockOracleKeeper() *mockOracleKeeper {
	return &mockOracleKeeper{
		prices: make(map[string]sdkmath.LegacyDec),
	}
}

func (m *mockOracleKeeper) SetPrice(denom string, price sdkmath.LegacyDec) {
	m.prices[denom] = price
}

func (m *mockOracleKeeper) GetPriceDec(_ context.Context, denom string) (sdkmath.LegacyDec, error) {
	price, ok := m.prices[denom]
	if !ok {
		return sdkmath.LegacyDec{}, stablecointypes.ErrPriceNotFound
	}
	return price, nil
}

func (m *mockOracleKeeper) GetPriceDecSafe(ctx context.Context, denom string) (sdkmath.LegacyDec, error) {
	return m.GetPriceDec(ctx, denom)
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

func TestMsgCreateVault(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000)))
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	msg := stablecointypes.NewMsgCreateVault(owner.String(), sdk.NewInt64Coin("stst", 1_000), sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 500))

	resp, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.Equal(t, uint64(1), resp.VaultId)

	vault, found := k.GetVault(ctx, resp.VaultId)
	require.True(t, found)
	require.Equal(t, owner.String(), vault.Owner)
	require.Equal(t, msg.Collateral, vault.Collateral)
	require.True(t, vault.Debt.Equal(sdkmath.NewInt(500)))

	require.True(t, bank.Balance(owner).AmountOf("stst").IsZero())
	require.True(t, bank.Balance(owner).AmountOf(stablecointypes.StablecoinDenom).Equal(sdkmath.NewInt(500)))
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf("stst").Equal(sdkmath.NewInt(1_000)))
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf(stablecointypes.StablecoinDenom).IsZero())
}

func TestMsgDepositAndWithdrawCollateral(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 800)))
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	create := stablecointypes.NewMsgCreateVault(owner.String(), sdk.NewInt64Coin("stst", 500), sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0))
	resp, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), create)
	require.NoError(t, err)

	deposit := stablecointypes.NewMsgDepositCollateral(owner.String(), resp.VaultId, sdk.NewInt64Coin("stst", 200))
	_, err = msgServer.DepositCollateral(sdk.WrapSDKContext(ctx), deposit)
	require.NoError(t, err)

	withdraw := &stablecointypes.MsgWithdrawCollateral{
		Owner:      owner.String(),
		VaultId:    resp.VaultId,
		Collateral: sdk.NewInt64Coin("stst", 100),
	}
	_, err = msgServer.WithdrawCollateral(sdk.WrapSDKContext(ctx), withdraw)
	require.NoError(t, err)

	vault, found := k.GetVault(ctx, resp.VaultId)
	require.True(t, found)
	require.Equal(t, sdkmath.NewInt(600), vault.Collateral.Amount)
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf("stst").Equal(sdkmath.NewInt(600)))
	require.True(t, bank.Balance(owner).AmountOf("stst").Equal(sdkmath.NewInt(200)))
}

func TestMsgMintAndRepayStablecoin(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000)))
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	create := stablecointypes.NewMsgCreateVault(owner.String(), sdk.NewInt64Coin("stst", 1_000), sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0))
	resp, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), create)
	require.NoError(t, err)

	mint := &stablecointypes.MsgMintStablecoin{
		Owner:   owner.String(),
		VaultId: resp.VaultId,
		Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 300),
	}
	_, err = msgServer.MintStablecoin(sdk.WrapSDKContext(ctx), mint)
	require.NoError(t, err)

	vault, found := k.GetVault(ctx, resp.VaultId)
	require.True(t, found)
	require.True(t, vault.Debt.Equal(sdkmath.NewInt(300)))
	require.True(t, bank.Balance(owner).AmountOf(stablecointypes.StablecoinDenom).Equal(sdkmath.NewInt(300)))

	repay := &stablecointypes.MsgRepayStablecoin{
		Owner:   owner.String(),
		VaultId: resp.VaultId,
		Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 200),
	}
	_, err = msgServer.RepayStablecoin(sdk.WrapSDKContext(ctx), repay)
	require.NoError(t, err)

	vault, found = k.GetVault(ctx, resp.VaultId)
	require.True(t, found)
	require.True(t, vault.Debt.Equal(sdkmath.NewInt(100)))
	require.True(t, bank.Balance(owner).AmountOf(stablecointypes.StablecoinDenom).Equal(sdkmath.NewInt(100)))
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf(stablecointypes.StablecoinDenom).IsZero())
}

func TestMsgLiquidateHealthyVaultFails(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 500)))
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	create := stablecointypes.NewMsgCreateVault(owner.String(), sdk.NewInt64Coin("stst", 500), sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0))
	resp, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), create)
	require.NoError(t, err)

	liquidator := newAddress()
	liquidate := &stablecointypes.MsgLiquidateVault{
		Liquidator: liquidator.String(),
		VaultId:    resp.VaultId,
	}
	_, err = msgServer.LiquidateVault(sdk.WrapSDKContext(ctx), liquidate)
	require.Error(t, err)
	require.ErrorIs(t, err, stablecointypes.ErrVaultHealthy)
}

func TestMsgLiquidateVaultRepaysDebtFromLiquidator(t *testing.T) {
	k, ctx, bank, oracle, _ := setupKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newAddress()
	liquidator := newAddress()

	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000)))
	bank.SetBalance(liquidator, sdk.NewCoins(sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 600)))

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	create := stablecointypes.NewMsgCreateVault(
		owner.String(),
		sdk.NewInt64Coin("stst", 1_000),
		sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 400),
	)
	resp, err := msgServer.CreateVault(sdk.WrapSDKContext(ctx), create)
	require.NoError(t, err)

	// Price crash pushes the vault below the liquidation ratio.
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("0.5"))

	_, err = msgServer.LiquidateVault(sdk.WrapSDKContext(ctx), &stablecointypes.MsgLiquidateVault{
		Liquidator: liquidator.String(),
		VaultId:    resp.VaultId,
	})
	require.NoError(t, err)

	_, found := k.GetVault(ctx, resp.VaultId)
	require.False(t, found, "vault should be removed after liquidation")

	// Liquidator paid the ssusd debt and received the seized collateral.
	require.True(t, bank.Balance(liquidator).AmountOf(stablecointypes.StablecoinDenom).Equal(sdkmath.NewInt(200)))
	require.True(t, bank.Balance(liquidator).AmountOf("stst").Equal(sdkmath.NewInt(1_000)))

	// Module account should end up empty after the transfer/burn.
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf(stablecointypes.StablecoinDenom).IsZero())
	require.True(t, bank.ModuleBalance(stablecointypes.ModuleAccountName).AmountOf("stst").IsZero())
}
