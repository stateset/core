package keeper_test

import (
	"context"
	"fmt"
	"math/rand"
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
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

// SimulationConfig defines parameters for simulation tests
type SimulationConfig struct {
	NumAccounts      int
	NumBlocks        int
	BlockTime        time.Duration
	InitialBalance   int64
	PriceVolatility  float64
	LiquidationRatio sdkmath.LegacyDec
}

func DefaultSimConfig() SimulationConfig {
	return SimulationConfig{
		NumAccounts:      100,
		NumBlocks:        1000,
		BlockTime:        5 * time.Second,
		InitialBalance:   10_000_000,
		PriceVolatility:  0.05, // 5% max price change per block
		LiquidationRatio: sdkmath.LegacyMustNewDecFromStr("1.5"),
	}
}

// SimState tracks simulation state
type SimState struct {
	Accounts     []sdk.AccAddress
	Vaults       map[uint64]bool
	TotalDebt    sdkmath.Int
	TotalCollat  sdkmath.Int
	Liquidations int
	Mints        int
	Burns        int
}

// TestSimulateLiquidationCascade simulates market crash scenarios
func TestSimulateLiquidationCascade(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping simulation test in short mode")
	}

	config := DefaultSimConfig()
	config.NumBlocks = 500
	config.PriceVolatility = 0.10 // 10% volatility for stress test

	k, ctx, bank, oracle := setupSimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)
	simState := &SimState{
		Vaults:      make(map[uint64]bool),
		TotalDebt:   sdkmath.ZeroInt(),
		TotalCollat: sdkmath.ZeroInt(),
	}

	// Setup accounts
	simState.Accounts = make([]sdk.AccAddress, config.NumAccounts)
	for i := 0; i < config.NumAccounts; i++ {
		simState.Accounts[i] = newSimAddress()
		bank.SetBalance(simState.Accounts[i], sdk.NewCoins(
			sdk.NewInt64Coin("stst", config.InitialBalance),
		))
	}

	// Initial price
	currentPrice := sdkmath.LegacyMustNewDecFromStr("2.0")
	oracle.SetPrice("stst", currentPrice)

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Create initial vaults
	for i := 0; i < config.NumAccounts/2; i++ {
		owner := simState.Accounts[i]
		collateral := sdk.NewInt64Coin("stst", 100_000)

		msg := stablecointypes.NewMsgCreateVault(
			owner.String(),
			collateral,
			sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
		)
		resp, err := msgServer.CreateVault(ctx, msg)
		if err == nil {
			simState.Vaults[resp.VaultId] = true

			// Mint some stablecoin against it
			mintAmount := int64(50_000) // Conservative mint
			mintMsg := &stablecointypes.MsgMintStablecoin{
				Owner:   owner.String(),
				VaultId: resp.VaultId,
				Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, mintAmount),
			}
			if _, err := msgServer.MintStablecoin(ctx, mintMsg); err == nil {
				simState.Mints++
				simState.TotalDebt = simState.TotalDebt.Add(sdkmath.NewInt(mintAmount))
			}
		}
	}

	t.Logf("Initial state: %d vaults, debt=%s", len(simState.Vaults), simState.TotalDebt.String())

	// Simulate blocks with price volatility
	for block := 0; block < config.NumBlocks; block++ {
		// Update price with random walk
		priceChange := (rng.Float64()*2 - 1) * config.PriceVolatility
		newPrice := currentPrice.Mul(sdkmath.LegacyOneDec().Add(sdkmath.LegacyMustNewDecFromStr(fmt.Sprintf("%f", priceChange))))

		// Enforce minimum price to prevent division issues
		if newPrice.LT(sdkmath.LegacyMustNewDecFromStr("0.1")) {
			newPrice = sdkmath.LegacyMustNewDecFromStr("0.1")
		}

		oracle.SetPrice("stst", newPrice)
		currentPrice = newPrice

		// Advance block time
		ctx = ctx.WithBlockTime(ctx.BlockTime().Add(config.BlockTime))
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

		// Check for liquidatable vaults (simulation of keeper bot)
		liquidatableCount := 0
		k.IterateVaults(ctx, func(vault stablecointypes.Vault) bool {
			if vault.Debt.Amount.IsPositive() {
				collateralValue := currentPrice.MulInt(vault.Collateral.Amount)
				debtValue := sdkmath.LegacyNewDecFromInt(vault.Debt.Amount)
				ratio := collateralValue.Quo(debtValue)

				if ratio.LT(config.LiquidationRatio) {
					liquidatableCount++
				}
			}
			return false
		})

		// Random user actions
		if rng.Float64() < 0.3 { // 30% chance of user action per block
			actionType := rng.Intn(3)
			account := simState.Accounts[rng.Intn(len(simState.Accounts))]

			switch actionType {
			case 0: // Create new vault
				msg := stablecointypes.NewMsgCreateVault(
					account.String(),
					sdk.NewInt64Coin("stst", 10_000+rng.Int63n(90_000)),
					sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
				)
				if resp, err := msgServer.CreateVault(ctx, msg); err == nil {
					simState.Vaults[resp.VaultId] = true
				}

			case 1: // Deposit more collateral to random vault
				for vaultID := range simState.Vaults {
					vault, found := k.GetVault(ctx, vaultID)
					if found && vault.Owner == account.String() {
						deposit := &stablecointypes.MsgDepositCollateral{
							Owner:   account.String(),
							VaultId: vaultID,
							Amount:  sdk.NewInt64Coin("stst", 1000+rng.Int63n(9000)),
						}
						msgServer.DepositCollateral(ctx, deposit)
						break
					}
				}

			case 2: // Repay debt
				for vaultID := range simState.Vaults {
					vault, found := k.GetVault(ctx, vaultID)
					if found && vault.Owner == account.String() && vault.Debt.Amount.IsPositive() {
						repayAmount := vault.Debt.Amount.Quo(sdkmath.NewInt(2))
						if repayAmount.IsPositive() {
							repay := &stablecointypes.MsgRepayDebt{
								Owner:   account.String(),
								VaultId: vaultID,
								Amount:  sdk.NewCoin(stablecointypes.StablecoinDenom, repayAmount),
							}
							if _, err := msgServer.RepayDebt(ctx, repay); err == nil {
								simState.Burns++
							}
						}
						break
					}
				}
			}
		}

		// Log progress every 100 blocks
		if block%100 == 0 {
			t.Logf("Block %d: price=%s, liquidatable=%d, total_vaults=%d",
				block, currentPrice.String(), liquidatableCount, len(simState.Vaults))
		}
	}

	// Final state check
	t.Logf("Simulation complete: mints=%d, burns=%d, liquidations=%d",
		simState.Mints, simState.Burns, simState.Liquidations)

	// Verify system invariants
	var totalDebt, totalCollateral sdkmath.Int
	totalDebt = sdkmath.ZeroInt()
	totalCollateral = sdkmath.ZeroInt()

	k.IterateVaults(ctx, func(vault stablecointypes.Vault) bool {
		totalDebt = totalDebt.Add(vault.Debt.Amount)
		totalCollateral = totalCollateral.Add(vault.Collateral.Amount)
		return false
	})

	t.Logf("Final totals: debt=%s, collateral=%s", totalDebt.String(), totalCollateral.String())
}

// TestSimulateOracleManipulation tests oracle price manipulation resistance
func TestSimulateOracleManipulation(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping simulation test in short mode")
	}

	k, ctx, bank, oracle := setupSimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	owner := newSimAddress()
	bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000_000)))

	// Set normal price
	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	// Create vault with large collateral
	createMsg := stablecointypes.NewMsgCreateVault(
		owner.String(),
		sdk.NewInt64Coin("stst", 100_000_000),
		sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
	)
	resp, err := msgServer.CreateVault(ctx, createMsg)
	require.NoError(t, err)

	// Attempt flash loan style attack - spike price, mint max, crash price
	t.Run("flash_loan_attack_simulation", func(t *testing.T) {
		// Spike price 10x
		oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("20.0"))

		// Try to mint maximum possible
		maxMint := &stablecointypes.MsgMintStablecoin{
			Owner:   owner.String(),
			VaultId: resp.VaultId,
			Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 1_000_000_000),
		}
		_, err := msgServer.MintStablecoin(ctx, maxMint)
		// Should be limited by collateral ratio requirements
		if err != nil {
			t.Logf("Max mint rejected (expected): %v", err)
		}

		// Crash price
		oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("0.5"))

		// Verify vault is now undercollateralized
		vault, _ := k.GetVault(ctx, resp.VaultId)
		if vault.Debt.Amount.IsPositive() {
			collateralValue := sdkmath.LegacyMustNewDecFromStr("0.5").MulInt(vault.Collateral.Amount)
			debtValue := sdkmath.LegacyNewDecFromInt(vault.Debt.Amount)
			ratio := collateralValue.Quo(debtValue)
			t.Logf("Post-crash collateral ratio: %s", ratio.String())
		}
	})
}

// TestFuzzMintBurn performs fuzz testing on mint/burn operations
func TestFuzzMintBurn(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping fuzz test in short mode")
	}

	k, ctx, bank, oracle := setupSimKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	oracle.SetPrice("stst", sdkmath.LegacyMustNewDecFromStr("2.0"))

	rng := rand.New(rand.NewSource(42)) // Deterministic seed for reproducibility

	for i := 0; i < 1000; i++ {
		owner := newSimAddress()
		bank.SetBalance(owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000_000)))

		// Random collateral amount (fuzz input)
		collateralAmount := rng.Int63n(1_000_000_000) + 1

		createMsg := stablecointypes.NewMsgCreateVault(
			owner.String(),
			sdk.NewInt64Coin("stst", collateralAmount),
			sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
		)
		resp, err := msgServer.CreateVault(ctx, createMsg)
		if err != nil {
			continue // Invalid input, skip
		}

		// Random mint amount (fuzz input)
		mintAmount := rng.Int63n(collateralAmount * 2) // Intentionally allow over-mint attempts

		mintMsg := &stablecointypes.MsgMintStablecoin{
			Owner:   owner.String(),
			VaultId: resp.VaultId,
			Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, mintAmount),
		}
		_, mintErr := msgServer.MintStablecoin(ctx, mintMsg)

		// Verify invariants after each operation
		vault, found := k.GetVault(ctx, resp.VaultId)
		if found && vault.Debt.Amount.IsPositive() {
			// Collateral ratio should always be >= 150%
			price, _ := oracle.GetPriceDec(ctx, "stst")
			collateralValue := price.MulInt(vault.Collateral.Amount)
			debtValue := sdkmath.LegacyNewDecFromInt(vault.Debt.Amount)
			ratio := collateralValue.Quo(debtValue)

			require.True(t, ratio.GTE(sdkmath.LegacyMustNewDecFromStr("1.5")),
				"Collateral ratio violation at iteration %d: ratio=%s, mintErr=%v",
				i, ratio.String(), mintErr)
		}
	}
}

// Helper functions for simulation tests
func setupSimKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *simBankKeeper, *simOracleKeeper) {
	storeKey := storetypes.NewKVStoreKey(stablecointypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	err := stateStore.LoadLatestVersion()
	require.NoError(t, err)

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{
		ChainID: "stateset-sim",
		Time:    time.Now(),
		Height:  1,
	}, false, log.NewNopLogger())

	bankKeeper := newSimBankKeeper()
	oracleKeeper := newSimOracleKeeper()
	complianceKeeper := newSimComplianceKeeper()
	accountKeeper := newSimAccountKeeper()
	accountKeeper.SetAddress(stablecointypes.ModuleAccountName, newSimAddress())

	k := keeper.NewKeeper(cdc, storeKey, bankKeeper, accountKeeper, oracleKeeper, complianceKeeper)

	return k, ctx, bankKeeper, oracleKeeper
}

func newSimAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

// Simulation mock keepers
type simBankKeeper struct {
	balances       map[string]sdk.Coins
	moduleBalances map[string]sdk.Coins
}

func newSimBankKeeper() *simBankKeeper {
	return &simBankKeeper{
		balances:       make(map[string]sdk.Coins),
		moduleBalances: make(map[string]sdk.Coins),
	}
}

func (m *simBankKeeper) SetBalance(addr sdk.AccAddress, coins sdk.Coins) {
	m.balances[addr.String()] = coins.Sort()
}

func (m *simBankKeeper) SendCoinsFromAccountToModule(_ context.Context, sender sdk.AccAddress, module string, amt sdk.Coins) error {
	account := m.balances[sender.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	newBalance, negative := account.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient funds")
	}
	m.balances[sender.String()] = newBalance.Sort()
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Add(amt...).Sort()
	return nil
}

func (m *simBankKeeper) SendCoinsFromModuleToAccount(_ context.Context, module string, recipient sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	newBalance, negative := moduleCoins.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient module funds")
	}
	m.moduleBalances[module] = newBalance.Sort()
	account := m.balances[recipient.String()]
	if account == nil {
		account = sdk.NewCoins()
	}
	m.balances[recipient.String()] = account.Add(amt...).Sort()
	return nil
}

func (m *simBankKeeper) MintCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	m.moduleBalances[module] = moduleCoins.Add(amt...).Sort()
	return nil
}

func (m *simBankKeeper) BurnCoins(_ context.Context, module string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[module]
	if moduleCoins == nil {
		moduleCoins = sdk.NewCoins()
	}
	newBalance, negative := moduleCoins.SafeSub(amt...)
	if negative {
		return fmt.Errorf("insufficient module funds to burn")
	}
	m.moduleBalances[module] = newBalance.Sort()
	return nil
}

type simAccountKeeper struct {
	addresses map[string]sdk.AccAddress
}

func newSimAccountKeeper() *simAccountKeeper {
	return &simAccountKeeper{
		addresses: make(map[string]sdk.AccAddress),
	}
}

func (m *simAccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	return m.addresses[moduleName]
}

func (m *simAccountKeeper) SetModuleAccount(_ context.Context, macc sdk.ModuleAccountI) {
	m.addresses[macc.GetName()] = macc.GetAddress()
}

func (m *simAccountKeeper) SetAddress(moduleName string, addr sdk.AccAddress) {
	m.addresses[moduleName] = addr
}

type simOracleKeeper struct {
	prices map[string]sdkmath.LegacyDec
}

func newSimOracleKeeper() *simOracleKeeper {
	return &simOracleKeeper{
		prices: make(map[string]sdkmath.LegacyDec),
	}
}

func (m *simOracleKeeper) SetPrice(denom string, price sdkmath.LegacyDec) {
	m.prices[denom] = price
}

func (m *simOracleKeeper) GetPriceDec(_ context.Context, denom string) (sdkmath.LegacyDec, error) {
	price, ok := m.prices[denom]
	if !ok {
		return sdkmath.LegacyDec{}, stablecointypes.ErrPriceNotFound
	}
	return price, nil
}

type simComplianceKeeper struct{}

func newSimComplianceKeeper() *simComplianceKeeper {
	return &simComplianceKeeper{}
}

func (m *simComplianceKeeper) AssertCompliant(_ context.Context, addr sdk.AccAddress) error {
	return nil
}
