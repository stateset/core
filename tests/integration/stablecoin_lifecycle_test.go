//go:build integration
// +build integration

package integration

import (
	"testing"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	oraclekeeper "github.com/stateset/core/x/oracle/keeper"
	oracletypes "github.com/stateset/core/x/oracle/types"
	stablecoinkeeper "github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

// StablecoinLifecycleTestSuite tests the complete stablecoin lifecycle:
// Deposit collateral -> Mint ssUSD -> Price change -> Liquidation check -> Redeem
type StablecoinLifecycleTestSuite struct {
	suite.Suite

	ctx      sdk.Context
	cdc      codec.Codec
	storeKey storetypes.StoreKey

	// Keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	complianceKeeper compliancekeeper.Keeper
	oracleKeeper     oraclekeeper.Keeper
	stablecoinKeeper stablecoinkeeper.Keeper

	// Test accounts
	authority      sdk.AccAddress
	vaultOwner     sdk.AccAddress
	liquidator     sdk.AccAddress
	oracleProvider sdk.AccAddress
}

func TestStablecoinLifecycleTestSuite(t *testing.T) {
	suite.Run(t, new(StablecoinLifecycleTestSuite))
}

func (s *StablecoinLifecycleTestSuite) SetupTest() {
	// Set SDK config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("stateset", "statesetpub")
	config.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
	config.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")

	// Initialize codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	s.cdc = codec.NewProtoCodec(interfaceRegistry)

	// Create store keys
	storeKeys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		compliancetypes.StoreKey,
		oracletypes.StoreKey,
		stablecointypes.StoreKey,
	)
	s.storeKey = storeKeys[stablecointypes.StoreKey]

	// Create multistore
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	transientKey := storetypes.NewTransientStoreKey("transient_test")
	cms.MountStoreWithDB(transientKey, storetypes.StoreTypeTransient, db)
	
	err := cms.LoadLatestVersion()
	require.NoError(s.T(), err)

	// Create context
	s.ctx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).
		WithBlockHeight(1).
		WithBlockTime(time.Now())

	// Initialize test addresses
	s.authority = sdk.AccAddress([]byte("authority___________"))
	s.vaultOwner = sdk.AccAddress([]byte("vaultowner__________"))
	s.liquidator = sdk.AccAddress([]byte("liquidator__________"))
	s.oracleProvider = sdk.AccAddress([]byte("oracleprovider______"))

	// Initialize account keeper
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:        nil,
		stablecointypes.ModuleAccountName: {authtypes.Minter, authtypes.Burner},
	}
	s.accountKeeper = authkeeper.NewAccountKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
		address.NewBech32Codec("stateset"),
		"stateset",
		s.authority.String(),
	)

	// Initialize bank keeper
	s.bankKeeper = bankkeeper.NewBaseKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[banktypes.StoreKey]),
		s.accountKeeper,
		nil,
		s.authority.String(),
		log.NewNopLogger(),
	)

	// Initialize compliance keeper
	s.complianceKeeper = compliancekeeper.NewKeeper(
		s.cdc,
		storeKeys[compliancetypes.StoreKey],
		s.authority.String(),
	)

	// Initialize oracle keeper
	s.oracleKeeper = oraclekeeper.NewKeeper(
		s.cdc,
		storeKeys[oracletypes.StoreKey],
		s.authority.String(),
	)

	// Initialize stablecoin keeper
	s.stablecoinKeeper = stablecoinkeeper.NewKeeper(
		s.cdc,
		storeKeys[stablecointypes.StoreKey],
		s.authority.String(),
		s.bankKeeper,
		s.accountKeeper,
		s.oracleKeeper,
		s.complianceKeeper,
	)

	// Setup test data
	s.setupTestAccounts()
	s.setupComplianceProfiles()
	s.setupOraclePrices()
	s.setupStablecoinParams()
}

func (s *StablecoinLifecycleTestSuite) setupTestAccounts() {
	// Create module account
	stablecoinModuleAcc := authtypes.NewEmptyModuleAccount(stablecointypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	stablecoinModuleAcc = s.accountKeeper.NewAccount(s.ctx, stablecoinModuleAcc).(*authtypes.ModuleAccount)
	s.accountKeeper.SetModuleAccount(s.ctx, stablecoinModuleAcc)

	// Create user accounts
	vaultOwnerAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.vaultOwner)
	liquidatorAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.liquidator)
	oracleAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.oracleProvider)
	s.accountKeeper.SetAccount(s.ctx, vaultOwnerAcc)
	s.accountKeeper.SetAccount(s.ctx, liquidatorAcc)
	s.accountKeeper.SetAccount(s.ctx, oracleAcc)

	// Mint collateral for vault owner (ATOM)
	collateralCoins := sdk.NewCoins(sdk.NewCoin("uatom", sdkmath.NewInt(10000000000))) // 10,000 ATOM
	require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, stablecointypes.ModuleAccountName, collateralCoins))
	require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, stablecointypes.ModuleAccountName, s.vaultOwner, collateralCoins))

	// Mint ssUSD for liquidator
	liquidatorCoins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000))) // 10,000 ssUSD
	require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, stablecointypes.ModuleAccountName, liquidatorCoins))
	require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, stablecointypes.ModuleAccountName, s.liquidator, liquidatorCoins))
}

func (s *StablecoinLifecycleTestSuite) setupComplianceProfiles() {
	// Setup compliant vault owner profile
	vaultOwnerProfile := compliancetypes.Profile{
		Address:        s.vaultOwner.String(),
		Status:         compliancetypes.StatusActive,
		KYCLevel:       compliancetypes.KYCStandard,
		Sanction:       false,
		VerifiedAt:     s.ctx.BlockTime(),
		ExpiresAt:      s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, vaultOwnerProfile)

	// Setup liquidator profile
	liquidatorProfile := compliancetypes.Profile{
		Address:        s.liquidator.String(),
		Status:         compliancetypes.StatusActive,
		KYCLevel:       compliancetypes.KYCStandard,
		Sanction:       false,
		VerifiedAt:     s.ctx.BlockTime(),
		ExpiresAt:      s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, liquidatorProfile)
}

func (s *StablecoinLifecycleTestSuite) setupOraclePrices() {
	// Set initial ATOM price: $10.00
	atomPrice := oracletypes.Price{
		Denom:       "uatom",
		Amount:      sdkmath.LegacyMustNewDecFromStr("10.00"),
		LastUpdater: s.oracleProvider.String(),
		LastHeight:  s.ctx.BlockHeight(),
		UpdatedAt:   s.ctx.BlockTime(),
	}
	s.oracleKeeper.SetPrice(s.ctx, atomPrice)

	// Set price for USDTBILL: $1.00
	usdtbillPrice := oracletypes.Price{
		Denom:       "USDTBILL",
		Amount:      sdkmath.LegacyOneDec(),
		LastUpdater: s.oracleProvider.String(),
		LastHeight:  s.ctx.BlockHeight(),
		UpdatedAt:   s.ctx.BlockTime(),
	}
	s.oracleKeeper.SetPrice(s.ctx, usdtbillPrice)

	// Enable config for USDTBILL
	usdtbillConfig := oracletypes.DefaultOracleConfig("USDTBILL")
	usdtbillConfig.Enabled = true
	require.NoError(s.T(), s.oracleKeeper.SetOracleConfig(s.ctx, usdtbillConfig))

	// Register oracle provider
	provider := oracletypes.OracleProvider{
		Address:               s.oracleProvider.String(),
		IsActive:              true,
		Slashed:               false,
		TotalSubmissions:      1,
		SuccessfulSubmissions: 1,
	}
	require.NoError(s.T(), s.oracleKeeper.SetProvider(s.ctx, provider))
}

func (s *StablecoinLifecycleTestSuite) setupStablecoinParams() {
	// Setup stablecoin parameters
	params := stablecointypes.Params{
		VaultMintingEnabled: true,
		CollateralParams: []stablecointypes.CollateralParam{
			{
				Denom:            "uatom",
				Active:           true,
				LiquidationRatio: sdkmath.LegacyMustNewDecFromStr("1.5"), // 150% collateralization
				DebtLimit:        sdkmath.NewInt(100000000000),           // 100k ssUSD
			},
		},
	}
	s.stablecoinKeeper.SetParams(s.ctx, params)
}

// TestCompleteStablecoinLifecycle tests the full lifecycle
func (s *StablecoinLifecycleTestSuite) TestCompleteStablecoinLifecycle() {
	// Scenario: Create vault, mint, check health, price drops, liquidate

	initialAtomBalance := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "uatom")
	s.Require().True(initialAtomBalance.Amount.GT(sdkmath.ZeroInt()))

	// Step 1: Create vault with collateral
	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000)) // 1000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(
		s.ctx,
		s.vaultOwner,
		collateral,
		sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
	)
	s.Require().NoError(err, "Vault creation should succeed")
	s.Require().Greater(vaultID, uint64(0))

	// Verify vault exists
	vault, found := s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	s.Require().True(found)
	s.Require().Equal(s.vaultOwner.String(), vault.Owner)
	s.Require().Equal(collateral, vault.Collateral)
	s.Require().True(vault.Debt.IsZero())

	// Verify collateral was transferred
	atomBalanceAfterVault := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "uatom")
	s.Require().Equal(initialAtomBalance.Amount.Sub(collateral.Amount), atomBalanceAfterVault.Amount)

	// Step 2: Mint stablecoin against collateral
	// At $10/ATOM, 1000 ATOM = $10,000 value
	// With 150% ratio, can mint max $10,000 / 1.5 = $6,666 ssUSD
	mintAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)) // 5000 ssUSD (safe)

	initialSsUsdBalance := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "ssusd")

	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID, mintAmount)
	s.Require().NoError(err, "Minting should succeed")

	// Verify ssUSD was minted
	ssUsdBalanceAfterMint := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "ssusd")
	s.Require().Equal(initialSsUsdBalance.Amount.Add(mintAmount.Amount), ssUsdBalanceAfterMint.Amount)

	// Verify vault debt increased
	vault, _ = s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	s.Require().Equal(mintAmount.Amount, vault.Debt)

	// Step 3: Verify vault is healthy at current price
	// Collateral value: 1000 ATOM * $10 = $10,000
	// Debt: 5000 ssUSD
	// Ratio: $10,000 / $5,000 = 200% (healthy)
	collateralValue := collateral.Amount.ToLegacyDec().Mul(sdkmath.LegacyMustNewDecFromStr("10.00"))
	debtValue := vault.Debt.ToLegacyDec()
	currentRatio := collateralValue.Quo(debtValue)
	s.Require().True(currentRatio.GTE(sdkmath.LegacyMustNewDecFromStr("1.5")), "Vault should be healthy")

	// Step 4: Simulate price drop (Oracle update)
	// Price drops from $10 to $6 per ATOM
	newPrice := oracletypes.Price{
		Denom:       "uatom",
		Amount:      sdkmath.LegacyMustNewDecFromStr("6.00"),
		LastUpdater: s.oracleProvider.String(),
		LastHeight:  s.ctx.BlockHeight() + 1,
		UpdatedAt:   s.ctx.BlockTime().Add(time.Hour),
	}
	s.oracleKeeper.SetPrice(s.ctx, newPrice)
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + 1).WithBlockTime(s.ctx.BlockTime().Add(time.Hour))

	// Step 5: Verify vault is now undercollateralized
	// New collateral value: 1000 ATOM * $6 = $6,000
	// Debt: 5000 ssUSD
	// Ratio: $6,000 / $5,000 = 120% (under 150% threshold)
	newCollateralValue := collateral.Amount.ToLegacyDec().Mul(sdkmath.LegacyMustNewDecFromStr("6.00"))
	newRatio := newCollateralValue.Quo(debtValue)
	s.Require().True(newRatio.LT(sdkmath.LegacyMustNewDecFromStr("1.5")), "Vault should be undercollateralized")

	// Step 6: Liquidate the vault
	initialLiquidatorSsUsd := s.bankKeeper.GetBalance(s.ctx, s.liquidator, "ssusd")
	initialLiquidatorAtom := s.bankKeeper.GetBalance(s.ctx, s.liquidator, "uatom")

	receivedCollateral, err := s.stablecoinKeeper.LiquidateVault(s.ctx, s.liquidator, vaultID)
	s.Require().NoError(err, "Liquidation should succeed")
	s.Require().NotNil(receivedCollateral)
	s.Require().Equal(collateral.Amount, receivedCollateral.AmountOf("uatom"))

	// Verify liquidator paid debt and received collateral
	liquidatorSsUsdAfter := s.bankKeeper.GetBalance(s.ctx, s.liquidator, "ssusd")
	liquidatorAtomAfter := s.bankKeeper.GetBalance(s.ctx, s.liquidator, "uatom")

	s.Require().Equal(initialLiquidatorSsUsd.Amount.Sub(vault.Debt), liquidatorSsUsdAfter.Amount,
		"Liquidator should pay debt")
	s.Require().Equal(initialLiquidatorAtom.Amount.Add(collateral.Amount), liquidatorAtomAfter.Amount,
		"Liquidator should receive collateral")

	// Verify vault is deleted
	_, found = s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	s.Require().False(found, "Vault should be deleted after liquidation")

	// Step 7: Verify liquidator profit
	// Liquidator paid 5000 ssUSD and received 1000 ATOM worth $6,000 at new price
	// Profit = $6,000 - $5,000 = $1,000 (in ATOM terms ~166 ATOM)
	profit := newCollateralValue.Sub(debtValue)
	s.Require().True(profit.GT(sdkmath.LegacyZeroDec()), "Liquidator should profit")
}

// TestVaultWithdrawCollateral tests withdrawing collateral while maintaining health
func (s *StablecoinLifecycleTestSuite) TestVaultWithdrawCollateral() {
	// Scenario: Create vault, mint, then withdraw some collateral

	// Create vault
	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(2000000000)) // 2000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Mint ssUSD
	mintAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)) // 5000 ssUSD
	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID, mintAmount)
	s.Require().NoError(err)

	// Try to withdraw collateral that would break health
	// Current: 2000 ATOM * $10 = $20,000 collateral, 5000 debt
	// Ratio: 400% (very healthy)
	// To break health: need ratio < 150%, i.e., collateral < 7500
	// So can withdraw at most: 2000 ATOM - 750 ATOM = 1250 ATOM

	// Try withdrawing too much (should fail)
	excessiveWithdraw := sdk.NewCoin("uatom", sdkmath.NewInt(1300000000)) // 1300 ATOM
	err = s.stablecoinKeeper.WithdrawCollateral(s.ctx, s.vaultOwner, vaultID, excessiveWithdraw)
	s.Require().Error(err, "Excessive withdrawal should fail")

	// Withdraw safe amount
	safeWithdraw := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000)) // 1000 ATOM
	initialAtomBalance := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "uatom")

	err = s.stablecoinKeeper.WithdrawCollateral(s.ctx, s.vaultOwner, vaultID, safeWithdraw)
	s.Require().NoError(err, "Safe withdrawal should succeed")

	// Verify withdrawal
	atomBalanceAfter := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "uatom")
	s.Require().Equal(initialAtomBalance.Amount.Add(safeWithdraw.Amount), atomBalanceAfter.Amount)

	// Verify vault collateral decreased
	vault, _ := s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	expectedCollateral := collateral.Amount.Sub(safeWithdraw.Amount)
	s.Require().Equal(expectedCollateral, vault.Collateral.Amount)
}

// TestRepayDebt tests repaying stablecoin debt
func (s *StablecoinLifecycleTestSuite) TestRepayDebt() {
	// Scenario: Create vault, mint, then repay debt

	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000)) // 1000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Mint ssUSD
	mintAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(3000000000)) // 3000 ssUSD
	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID, mintAmount)
	s.Require().NoError(err)

	vault, _ := s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	initialDebt := vault.Debt

	// Repay partial debt
	repayAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 ssUSD
	initialSsUsdBalance := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "ssusd")

	err = s.stablecoinKeeper.RepayStablecoin(s.ctx, s.vaultOwner, vaultID, repayAmount)
	s.Require().NoError(err, "Repayment should succeed")

	// Verify ssUSD was burned
	ssUsdBalanceAfter := s.bankKeeper.GetBalance(s.ctx, s.vaultOwner, "ssusd")
	s.Require().Equal(initialSsUsdBalance.Amount.Sub(repayAmount.Amount), ssUsdBalanceAfter.Amount)

	// Verify debt decreased
	vault, _ = s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	expectedDebt := initialDebt.Sub(repayAmount.Amount)
	s.Require().Equal(expectedDebt, vault.Debt)

	// Repay remaining debt
	remainingRepay := sdk.NewCoin("ssusd", vault.Debt)
	err = s.stablecoinKeeper.RepayStablecoin(s.ctx, s.vaultOwner, vaultID, remainingRepay)
	s.Require().NoError(err)

	// Verify debt is zero
	vault, _ = s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	s.Require().True(vault.Debt.IsZero(), "Debt should be zero after full repayment")
}

// TestDepositMoreCollateral tests adding collateral to existing vault
func (s *StablecoinLifecycleTestSuite) TestDepositMoreCollateral() {
	// Scenario: Create vault, then add more collateral

	initialCollateral := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000)) // 1000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, initialCollateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Add more collateral
	additionalCollateral := sdk.NewCoin("uatom", sdkmath.NewInt(500000000)) // 500 ATOM
	err = s.stablecoinKeeper.DepositCollateral(s.ctx, s.vaultOwner, vaultID, additionalCollateral)
	s.Require().NoError(err, "Deposit should succeed")

	// Verify vault collateral increased
	vault, _ := s.stablecoinKeeper.GetVault(s.ctx, vaultID)
	expectedCollateral := initialCollateral.Amount.Add(additionalCollateral.Amount)
	s.Require().Equal(expectedCollateral, vault.Collateral.Amount)
}

// TestCannotMintOverDebtLimit tests debt limit enforcement
func (s *StablecoinLifecycleTestSuite) TestCannotMintOverDebtLimit() {
	// Scenario: Try to mint more than allowed by debt limit

	// Create vault with large collateral
	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(100000000000)) // 100,000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Try to mint over debt limit (100k ssUSD limit in params)
	excessiveMint := sdk.NewCoin("ssusd", sdkmath.NewInt(200000000000)) // 200k ssUSD
	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID, excessiveMint)
	s.Require().Error(err, "Should fail due to debt limit")
	s.Require().Contains(err.Error(), "debt limit")
}

// TestLiquidationFailsForHealthyVault tests that healthy vaults cannot be liquidated
func (s *StablecoinLifecycleTestSuite) TestLiquidationFailsForHealthyVault() {
	// Scenario: Try to liquidate a healthy vault

	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(2000000000)) // 2000 ATOM
	vaultID, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Mint conservative amount
	mintAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)) // 5000 ssUSD
	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID, mintAmount)
	s.Require().NoError(err)

	// Vault is healthy: 2000 ATOM * $10 = $20,000 / 5000 = 400%
	// Try to liquidate
	_, err = s.stablecoinKeeper.LiquidateVault(s.ctx, s.liquidator, vaultID)
	s.Require().Error(err, "Cannot liquidate healthy vault")
	s.Require().Contains(err.Error(), "still healthy")
}

// TestMultipleVaults tests managing multiple vaults
func (s *StablecoinLifecycleTestSuite) TestMultipleVaults() {
	// Scenario: Create multiple vaults and manage them independently

	// Create first vault
	collateral1 := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000))
	vaultID1, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral1, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Create second vault
	collateral2 := sdk.NewCoin("uatom", sdkmath.NewInt(2000000000))
	vaultID2, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral2, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	s.Require().NotEqual(vaultID1, vaultID2, "Vault IDs should be unique")

	// Mint from both vaults
	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID1, sdk.NewCoin("ssusd", sdkmath.NewInt(2000000000)))
	s.Require().NoError(err)

	err = s.stablecoinKeeper.MintStablecoin(s.ctx, s.vaultOwner, vaultID2, sdk.NewCoin("ssusd", sdkmath.NewInt(4000000000)))
	s.Require().NoError(err)

	// Verify both vaults exist with correct debt
	vault1, found1 := s.stablecoinKeeper.GetVault(s.ctx, vaultID1)
	vault2, found2 := s.stablecoinKeeper.GetVault(s.ctx, vaultID2)
	s.Require().True(found1)
	s.Require().True(found2)
	s.Require().Equal(sdkmath.NewInt(2000000000), vault1.Debt)
	s.Require().Equal(sdkmath.NewInt(4000000000), vault2.Debt)
}

// TestEventEmissions tests that proper events are emitted
func (s *StablecoinLifecycleTestSuite) TestEventEmissions() {
	// Scenario: Perform operations and verify events

	collateral := sdk.NewCoin("uatom", sdkmath.NewInt(1000000000))
	_, err := s.stablecoinKeeper.CreateVault(s.ctx, s.vaultOwner, collateral, sdk.NewCoin("ssusd", sdkmath.ZeroInt()))
	s.Require().NoError(err)

	// Check events were emitted
	events := s.ctx.EventManager().Events()
	s.Require().Greater(len(events), 0, "Events should be emitted")
}
