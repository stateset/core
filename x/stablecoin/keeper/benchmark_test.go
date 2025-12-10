package keeper_test

import (
	"context"
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
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

func setupBenchmarkKeeper() (keeper.Keeper, sdk.Context, *benchBankKeeper, *benchOracleKeeper) {
	storeKey := storetypes.NewKVStoreKey(stablecointypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.LoadLatestVersion()

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{ChainID: "stateset-bench", Time: time.Now()}, false, log.NewNopLogger())

	bankKeeper := newBenchBankKeeper()
	oracleKeeper := newBenchOracleKeeper()
	complianceKeeper := newBenchComplianceKeeper()
	accountKeeper := newBenchAccountKeeper()
	accountKeeper.SetAddress(stablecointypes.ModuleAccountName, newBenchAddress())

	k := keeper.NewKeeper(cdc, storeKey, "stateset1authority", bankKeeper, accountKeeper, oracleKeeper, complianceKeeper)

	return k, ctx, bankKeeper, oracleKeeper
}

func newBenchAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

type benchBankKeeper struct {
	balances       map[string]sdk.Coins
	moduleBalances map[string]sdk.Coins
}

func newBenchBankKeeper() *benchBankKeeper {
	return &benchBankKeeper{
		balances:       make(map[string]sdk.Coins),
		moduleBalances: make(map[string]sdk.Coins),
	}
}

func (m *benchBankKeeper) SetBalance(addr sdk.AccAddress, coins sdk.Coins) {
	m.balances[addr.String()] = coins.Sort()
}

func (m *benchBankKeeper) SendCoinsFromAccountToModule(_ context.Context, sender sdk.AccAddress, module string, amt sdk.Coins) error {
	account := m.balances[sender.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	m.balances[sender.String()] = account.Sub(amt...).Sort()
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Add(amt...).Sort()
	return nil
}

func (m *benchBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, module string, recipient sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Sub(amt...).Sort()
	account := m.balances[recipient.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	m.balances[recipient.String()] = account.Add(amt...).Sort()
	return nil
}

func (m *benchBankKeeper) MintCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Add(amt...).Sort()
	return nil
}

func (m *benchBankKeeper) BurnCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Sub(amt...).Sort()
	return nil
}

type benchAccountKeeper struct {
	addresses map[string]sdk.AccAddress
}

func newBenchAccountKeeper() *benchAccountKeeper {
	return &benchAccountKeeper{
		addresses: make(map[string]sdk.AccAddress),
	}
}

func (m *benchAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return m.addresses[moduleName]
}

func (m *benchAccountKeeper) SetModuleAccount(_ context.Context, macc sdk.ModuleAccountI) {
	m.addresses[macc.GetName()] = macc.GetAddress()
}

func (m *benchAccountKeeper) SetAddress(moduleName string, addr sdk.AccAddress) {
	m.addresses[moduleName] = addr
}

type benchOracleKeeper struct {
	prices map[string]sdkmath.LegacyDec
}

func newBenchOracleKeeper() *benchOracleKeeper {
	return &benchOracleKeeper{
		prices: make(map[string]sdkmath.LegacyDec),
	}
}

func (m *benchOracleKeeper) SetPrice(denom string, price sdkmath.LegacyDec) {
	m.prices[denom] = price
}

func (m *benchOracleKeeper) GetPriceDec(_ context.Context, denom string) (sdkmath.LegacyDec, error) {
	price, ok := m.prices[denom]
	if !ok {
		return sdkmath.LegacyDec{}, stablecointypes.ErrPriceNotFound
	}
	return price, nil
}

type benchComplianceKeeper struct{}

func newBenchComplianceKeeper() *benchComplianceKeeper {
	return &benchComplianceKeeper{}
}

func (m *benchComplianceKeeper) AssertCompliant(_ context.Context, addr sdk.AccAddress) error {
	return nil
}

func BenchmarkCreateVault(b *testing.B) {
	k, ctx, bank, oracle := setupBenchmarkKeeper()
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		owner := newBenchAddress()
		bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000)))

		msg := stablecointypes.NewMsgCreateVault(
			owner.String(),
			sdk.NewInt64Coin("stst", 1_000),
			sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
		)
		msgServer.CreateVault(ctx, msg)
	}
}

func BenchmarkDepositCollateral(b *testing.B) {
	k, ctx, bank, oracle := setupBenchmarkKeeper()
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	owner := newBenchAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000_000)))

	createMsg := stablecointypes.NewMsgCreateVault(
		owner.String(),
		sdk.NewInt64Coin("stst", 1_000),
		sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
	)
	resp, _ := msgServer.CreateVault(ctx, createMsg)
	vaultId := resp.VaultId

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		deposit := stablecointypes.NewMsgDepositCollateral(
			owner.String(),
			vaultId,
			sdk.NewInt64Coin("stst", 100),
		)
		msgServer.DepositCollateral(ctx, deposit)
	}
}

func BenchmarkMintStablecoin(b *testing.B) {
	k, ctx, bank, oracle := setupBenchmarkKeeper()
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	owner := newBenchAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000_000)))

	createMsg := stablecointypes.NewMsgCreateVault(
		owner.String(),
		sdk.NewInt64Coin("stst", 1_000_000),
		sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
	)
	resp, _ := msgServer.CreateVault(ctx, createMsg)
	vaultId := resp.VaultId

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mint := &stablecointypes.MsgMintStablecoin{
			Owner:   owner.String(),
			VaultId: vaultId,
			Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 10),
		}
		msgServer.MintStablecoin(ctx, mint)
	}
}

func BenchmarkGetVault(b *testing.B) {
	k, ctx, bank, oracle := setupBenchmarkKeeper()
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	// Create 1000 vaults
	for i := 0; i < 1000; i++ {
		owner := newBenchAddress()
		bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000)))

		createMsg := stablecointypes.NewMsgCreateVault(
			owner.String(),
			sdk.NewInt64Coin("stst", 1_000),
			sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
		)
		msgServer.CreateVault(ctx, createMsg)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k.GetVault(ctx, uint64((i%1000)+1))
	}
}

func BenchmarkIterateVaults(b *testing.B) {
	k, ctx, bank, oracle := setupBenchmarkKeeper()
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	// Create 100 vaults
	for i := 0; i < 100; i++ {
		owner := newBenchAddress()
		bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000)))

		createMsg := stablecointypes.NewMsgCreateVault(
			owner.String(),
			sdk.NewInt64Coin("stst", 1_000),
			sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
		)
		msgServer.CreateVault(ctx, createMsg)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count := 0
		k.IterateVaults(ctx, func(vault stablecointypes.Vault) bool {
			count++
			return false
		})
	}
}
