//go:build integration
// +build integration

package integration

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
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

// RWAFlowTestSuite tests the Reserve-Backed Stablecoin flow (Real World Assets via IBC)
type RWAFlowTestSuite struct {
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
	investor       sdk.AccAddress
	oracleProvider sdk.AccAddress
}

func TestRWAFlowTestSuite(t *testing.T) {
	suite.Run(t, new(RWAFlowTestSuite))
}

func (s *RWAFlowTestSuite) SetupTest() {
	// 1. Basic Setup (Codec, Store, Context)
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(interfaceRegistry)
	banktypes.RegisterInterfaces(interfaceRegistry)
	s.cdc = codec.NewProtoCodec(interfaceRegistry)

	storeKeys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey,
		banktypes.StoreKey,
		compliancetypes.StoreKey,
		oracletypes.StoreKey,
		stablecointypes.StoreKey,
	)
	s.storeKey = storeKeys[stablecointypes.StoreKey]

	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	transientKey := storetypes.NewTransientStoreKey("transient_test")
	cms.MountStoreWithDB(transientKey, storetypes.StoreTypeTransient, db)
	
	err := cms.LoadLatestVersion()
	require.NoError(s.T(), err)

	s.ctx = sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).
		WithBlockHeight(1).
		WithBlockTime(time.Now())

	// 2. Accounts
	s.authority = sdk.AccAddress([]byte("authority___________"))
	s.investor = sdk.AccAddress([]byte("investor____________"))
	s.oracleProvider = sdk.AccAddress([]byte("oracleprovider______"))

	// 3. Keepers Initialization
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

	s.bankKeeper = bankkeeper.NewBaseKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[banktypes.StoreKey]),
		s.accountKeeper,
		nil,
		s.authority.String(),
		log.NewNopLogger(),
	)

	s.complianceKeeper = compliancekeeper.NewKeeper(
		s.cdc,
		storeKeys[compliancetypes.StoreKey],
		s.authority.String(),
	)

	s.oracleKeeper = oraclekeeper.NewKeeper(
		s.cdc,
		storeKeys[oracletypes.StoreKey],
		s.authority.String(),
	)

	s.stablecoinKeeper = stablecoinkeeper.NewKeeper(
		s.cdc,
		storeKeys[stablecointypes.StoreKey],
		s.authority.String(),
		s.bankKeeper,
		s.accountKeeper,
		s.oracleKeeper,
		s.complianceKeeper,
	)

	// 4. State Setup
	s.setupAccounts()
	s.setupCompliance()
	s.setupOracle()
	s.setupStablecoinParams()
}

func (s *RWAFlowTestSuite) setupAccounts() {
	// Create module account
	stablecoinModuleAcc := authtypes.NewEmptyModuleAccount(stablecointypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	stablecoinModuleAcc = s.accountKeeper.NewAccount(s.ctx, stablecoinModuleAcc).(*authtypes.ModuleAccount)
	s.accountKeeper.SetModuleAccount(s.ctx, stablecoinModuleAcc)

	// Fund investor with "ibc/USDC" (simulated IBC token)
	// 10,000 USDC (6 decimals usually, but simplified here to 1 unit = $1 for test clarity if needed, or stick to decimals)
	// Let's assume 1_000_000 = 1 USDC for consistency with standard tests, or just 10000_000000
	ibcDenom := "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	amount := sdk.NewCoins(sdk.NewCoin(ibcDenom, sdkmath.NewInt(10_000_000_000))) // $10,000 USDC
	
	s.bankKeeper.MintCoins(s.ctx, stablecointypes.ModuleAccountName, amount)
	s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, stablecointypes.ModuleAccountName, s.investor, amount)
}

func (s *RWAFlowTestSuite) setupCompliance() {
	// Investor must be KYC verified
	profile := compliancetypes.Profile{
		Address:    s.investor.String(),
		Status:     compliancetypes.StatusActive,
		KYCLevel:   compliancetypes.KYCStandard,
		VerifiedAt: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, profile)
}

func (s *RWAFlowTestSuite) setupOracle() {
	// Register provider
	provider := oracletypes.OracleProvider{
		Address:  s.oracleProvider.String(),
		IsActive: true,
	}
	s.oracleKeeper.SetProvider(s.ctx, provider)

	// Set price for USDC (The underlying asset of the IBC token)
	// Price $1.00
	price := oracletypes.Price{
		Denom:       "USDC",
		Amount:      sdkmath.LegacyOneDec(),
		LastUpdater: s.oracleProvider.String(),
		UpdatedAt:   s.ctx.BlockTime(),
	}
	s.oracleKeeper.SetPrice(s.ctx, price)
	
	// Enable config for USDC
	config := oracletypes.DefaultOracleConfig("USDC")
	config.Enabled = true
	s.oracleKeeper.SetOracleConfig(s.ctx, config)
}

func (s *RWAFlowTestSuite) setupStablecoinParams() {
	ibcDenom := "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	
	params := stablecointypes.DefaultReserveParams()
	params.RequireKyc = true
	params.MintFeeBps = 0 // Simplify calc
	params.RedeemFeeBps = 0
	params.TokenizedTreasuries = []stablecointypes.TokenizedTreasuryConfig{
		{
			Denom:            ibcDenom,
			OracleDenom:      "USDC",
			Active:           true,
			UnderlyingType:   stablecointypes.ReserveAssetCash, // 1:1 map
			HaircutBps:       0,
			MaxAllocationBps: 10000,
		},
	}
	s.stablecoinKeeper.SetReserveParams(s.ctx, params)
}

// TestEndToEndMintRedeem verifies the full user flow
func (s *RWAFlowTestSuite) TestEndToEndMintRedeem() {
	ibcDenom := "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	
	// 1. Check Initial Balance
	initialBal := s.bankKeeper.GetBalance(s.ctx, s.investor, ibcDenom)
	s.Require().Equal(int64(10_000_000_000), initialBal.Amount.Int64())

	// 2. Deposit 1000 USDC -> Mint 1000 ssUSD
	depositAmount := sdk.NewCoin(ibcDenom, sdkmath.NewInt(1_000_000_000))
	
	// Reset event manager to catch our specific event
	s.ctx = s.ctx.WithEventManager(sdk.NewEventManager())
	
	depositID, minted, err := s.stablecoinKeeper.DepositReserve(s.ctx, s.investor, depositAmount)
	s.Require().NoError(err)
	s.Require().Equal(int64(1_000_000_000), minted.Int64())
	s.Require().Equal(uint64(1), depositID)

	// 3. Verify Balances
	// Investor should have -1000 USDC, +1000 ssUSD
	usdcBal := s.bankKeeper.GetBalance(s.ctx, s.investor, ibcDenom)
	ssusdBal := s.bankKeeper.GetBalance(s.ctx, s.investor, "ssusd")
	
	s.Require().Equal(int64(9_000_000_000), usdcBal.Amount.Int64())
	s.Require().Equal(int64(1_000_000_000), ssusdBal.Amount.Int64())

	// 4. Verify Event (Safety Check)
	events := s.ctx.EventManager().Events()
	var depositEvent sdk.Event
	found := false
	for _, e := range events {
		if e.Type == stablecointypes.EventTypeReserveDeposit {
			depositEvent = e
			found = true
			break
		}
	}
	s.Require().True(found, "Reserve deposit event not found")
	
	// Check for the "price" attribute I added in the hardening step
	hasPrice := false
	for _, attr := range depositEvent.Attributes {
		if attr.Key == stablecointypes.AttributeKeyPrice {
			hasPrice = true
			s.Require().Equal("1.000000000000000000", attr.Value) // LegacyDec string format
		}
	}
	s.Require().True(hasPrice, "Price attribute missing from event log")

	// 5. Redemption
	// Redeem 500 ssUSD -> Get 500 USDC
	redeemAmount := sdkmath.NewInt(500_000_000)
	redemptionID, err := s.stablecoinKeeper.RequestRedemption(s.ctx, s.investor, redeemAmount, ibcDenom)
	s.Require().NoError(err)
	
	// Since RedemptionDelay is 0 in DefaultParams (and setup), it should execute immediately (or be pending if logic dictates)
	// Checking the code: RequestRedemption calls executeRedemption immediately if delay == 0.
	
	// Verify Request
	req, found := s.stablecoinKeeper.GetRedemptionRequest(s.ctx, redemptionID)
	s.Require().True(found)
	s.Require().Equal(stablecointypes.RedeemStatusExecuted, req.Status)

	// Verify Final Balances
	// Investor: 9000 + 500 = 9500 USDC
	// Investor: 1000 - 500 = 500 ssUSD
	finalUsdc := s.bankKeeper.GetBalance(s.ctx, s.investor, ibcDenom)
	finalSsusd := s.bankKeeper.GetBalance(s.ctx, s.investor, "ssusd")
	
	s.Require().Equal(int64(9_500_000_000), finalUsdc.Amount.Int64())
	s.Require().Equal(int64(500_000_000), finalSsusd.Amount.Int64())
}

// TestBadOraclePrice verifies that 0 or negative prices block minting
func (s *RWAFlowTestSuite) TestBadOraclePrice() {
	ibcDenom := "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	depositAmount := sdk.NewCoin(ibcDenom, sdkmath.NewInt(1_000_000_000))

	// 1. Set Price to Zero
	s.oracleKeeper.SetPrice(s.ctx, oracletypes.Price{
		Denom:       "USDC",
		Amount:      sdkmath.LegacyZeroDec(),
		LastUpdater: s.oracleProvider.String(),
		UpdatedAt:   s.ctx.BlockTime(),
	})

	// 2. Attempt Deposit (Should Fail)
	_, _, err := s.stablecoinKeeper.DepositReserve(s.ctx, s.investor, depositAmount)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid price") // Matches the error we added "invalid price %s"

	// 3. Set Price to Negative (Should impossible via LegacyDec usually, but checking logic path)
	// LegacyDec can be negative.
	negPrice := sdkmath.LegacyNewDec(-1)
	s.oracleKeeper.SetPrice(s.ctx, oracletypes.Price{
		Denom:       "USDC",
		Amount:      negPrice,
		LastUpdater: s.oracleProvider.String(),
		UpdatedAt:   s.ctx.BlockTime(),
	})

	_, _, err = s.stablecoinKeeper.DepositReserve(s.ctx, s.investor, depositAmount)
	s.Require().Error(err)
	s.Require().Contains(err.Error(), "invalid price")
}
