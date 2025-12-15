//go:build integration
// +build integration

package integration

import (
	"testing"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
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
	"cosmossdk.io/log"
	dbm "github.com/cosmos/cosmos-db"

	compliancekeeper "github.com/stateset/core/x/compliance/keeper"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	paymentskeeper "github.com/stateset/core/x/payments/keeper"
	paymentstypes "github.com/stateset/core/x/payments/types"
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
)

// CrossModuleComplianceTestSuite tests compliance integration across all modules:
// Payment + Compliance integration, Settlement + Compliance integration
type CrossModuleComplianceTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	cdc            codec.Codec

	// Keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	complianceKeeper compliancekeeper.Keeper
	paymentsKeeper   paymentskeeper.Keeper
	settlementKeeper settlementkeeper.Keeper

	// Test accounts with different compliance statuses
	authority        sdk.AccAddress
	compliantUser    sdk.AccAddress
	expiredUser      sdk.AccAddress
	suspendedUser    sdk.AccAddress
	sanctionedUser   sdk.AccAddress
	lowKYCUser       sdk.AccAddress
	limitExceededUser sdk.AccAddress
	merchant         sdk.AccAddress
}

func TestCrossModuleComplianceTestSuite(t *testing.T) {
	suite.Run(t, new(CrossModuleComplianceTestSuite))
}

func (s *CrossModuleComplianceTestSuite) SetupTest() {
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
		paymentstypes.StoreKey,
		settlementtypes.StoreKey,
	)

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
	s.compliantUser = sdk.AccAddress([]byte("compliantuser_______"))
	s.expiredUser = sdk.AccAddress([]byte("expireduser_________"))
	s.suspendedUser = sdk.AccAddress([]byte("suspendeduser_______"))
	s.sanctionedUser = sdk.AccAddress([]byte("sanctioneduser______"))
	s.lowKYCUser = sdk.AccAddress([]byte("lowkycuser__________"))
	s.limitExceededUser = sdk.AccAddress([]byte("limitexceededuser___"))
	s.merchant = sdk.AccAddress([]byte("merchant____________"))

	// Initialize account keeper
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:        nil,
		paymentstypes.ModuleAccountName:   {authtypes.Minter, authtypes.Burner},
		settlementtypes.ModuleAccountName: {authtypes.Minter, authtypes.Burner},
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

	// Initialize payments keeper
	s.paymentsKeeper = paymentskeeper.NewKeeper(
		s.cdc,
		storeKeys[paymentstypes.StoreKey],
		s.bankKeeper,
		s.complianceKeeper,
		paymentstypes.ModuleAccountName,
	)

	// Initialize settlement keeper
	s.settlementKeeper = settlementkeeper.NewKeeper(
		s.cdc,
		storeKeys[settlementtypes.StoreKey],
		s.bankKeeper,
		s.complianceKeeper,
		s.accountKeeper,
		s.authority.String(),
	)

	// Setup test data
	s.setupTestAccounts()
	s.setupComplianceProfiles()
}

func (s *CrossModuleComplianceTestSuite) setupTestAccounts() {
	// Create module accounts
	paymentsModuleAcc := authtypes.NewEmptyModuleAccount(paymentstypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	paymentsModuleAcc = s.accountKeeper.NewAccount(s.ctx, paymentsModuleAcc).(*authtypes.ModuleAccount)
	settlementModuleAcc := authtypes.NewEmptyModuleAccount(settlementtypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	settlementModuleAcc = s.accountKeeper.NewAccount(s.ctx, settlementModuleAcc).(*authtypes.ModuleAccount)
	s.accountKeeper.SetModuleAccount(s.ctx, paymentsModuleAcc)
	s.accountKeeper.SetModuleAccount(s.ctx, settlementModuleAcc)

	// Create user accounts and fund them
	allUsers := []sdk.AccAddress{
		s.compliantUser, s.expiredUser, s.suspendedUser, s.sanctionedUser,
		s.lowKYCUser, s.limitExceededUser, s.merchant,
	}

	for _, addr := range allUsers {
		acc := s.accountKeeper.NewAccountWithAddress(s.ctx, addr)
		s.accountKeeper.SetAccount(s.ctx, acc)

		// Mint coins
		coins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000)))
		require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, coins))
		require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, addr, coins))
	}
}

func (s *CrossModuleComplianceTestSuite) setupComplianceProfiles() {
	// 1. Fully compliant user
	compliantProfile := compliancetypes.Profile{
		Address:     s.compliantUser.String(),
		Status:      compliancetypes.StatusActive,
		KYCLevel:    compliancetypes.KYCStandard,
		Sanction:    false,
		DailyLimit:  sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000)),
		DailyUsed:   sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		MonthlyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(100000000000)),
		MonthlyUsed: sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		VerifiedAt:  s.ctx.BlockTime(),
		ExpiresAt:   s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, compliantProfile)

	// 2. Expired KYC user
	expiredProfile := compliancetypes.Profile{
		Address:    s.expiredUser.String(),
		Status:     compliancetypes.StatusActive,
		KYCLevel:   compliancetypes.KYCStandard,
		Sanction:   false,
		VerifiedAt: s.ctx.BlockTime().AddDate(-2, 0, 0),
		ExpiresAt:  s.ctx.BlockTime().AddDate(-1, 0, 0), // Expired 1 year ago
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, expiredProfile)

	// 3. Suspended user
	suspendedProfile := compliancetypes.Profile{
		Address:    s.suspendedUser.String(),
		Status:     compliancetypes.StatusSuspended,
		KYCLevel:   compliancetypes.KYCStandard,
		Sanction:   false,
		VerifiedAt: s.ctx.BlockTime(),
		ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, suspendedProfile)

	// 4. Sanctioned user
	sanctionedProfile := compliancetypes.Profile{
		Address:    s.sanctionedUser.String(),
		Status:     compliancetypes.StatusActive,
		KYCLevel:   compliancetypes.KYCStandard,
		Sanction:   true, // Sanctioned
		VerifiedAt: s.ctx.BlockTime(),
		ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, sanctionedProfile)

	// 5. Low KYC level user
	lowKYCProfile := compliancetypes.Profile{
		Address:    s.lowKYCUser.String(),
		Status:     compliancetypes.StatusActive,
		KYCLevel:   compliancetypes.KYCNone,
		Sanction:   false,
		VerifiedAt: s.ctx.BlockTime(),
		ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
		DailyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(100000000)), // Low limit
		DailyUsed:  sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, lowKYCProfile)

	// 6. User who exceeded limits
	limitExceededProfile := compliancetypes.Profile{
		Address:     s.limitExceededUser.String(),
		Status:      compliancetypes.StatusActive,
		KYCLevel:    compliancetypes.KYCStandard,
		Sanction:    false,
		DailyLimit:  sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)),
		DailyUsed:   sdk.NewCoin("ssusd", sdkmath.NewInt(999000000)), // Almost at limit
		MonthlyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000)),
		MonthlyUsed: sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)),
		VerifiedAt:  s.ctx.BlockTime(),
		ExpiresAt:   s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, limitExceededProfile)

	// 7. Compliant merchant
	merchantProfile := compliancetypes.Profile{
		Address:     s.merchant.String(),
		Status:      compliancetypes.StatusActive,
		KYCLevel:    compliancetypes.KYCEnhanced,
		Sanction:    false,
		DailyLimit:  sdk.NewCoin("ssusd", sdkmath.NewInt(100000000000)),
		DailyUsed:   sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		MonthlyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000000)),
		MonthlyUsed: sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		VerifiedAt:  s.ctx.BlockTime(),
		ExpiresAt:   s.ctx.BlockTime().AddDate(1, 0, 0),
		LastLimitReset: s.ctx.BlockTime(),
	}
	s.complianceKeeper.SetProfile(s.ctx, merchantProfile)
}

// TestPaymentComplianceIntegration tests payment module with compliance checks
func (s *CrossModuleComplianceTestSuite) TestPaymentComplianceIntegration() {
	// Test 1: Compliant user can create payment
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000000))
	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.compliantUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "compliant payment",
	}

	paymentID, err := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().NoError(err, "Compliant user should create payment")
	s.Require().Greater(paymentID, uint64(0))

	// Test 2: Expired user cannot create payment
	expiredIntent := paymentstypes.PaymentIntent{
		Payer:    s.expiredUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "expired payment",
	}

	_, err = s.paymentsKeeper.CreatePayment(s.ctx, expiredIntent)
	s.Require().Error(err, "Expired user should be blocked")
	s.Require().Contains(err.Error(), "expired")

	// Test 3: Suspended user cannot create payment
	suspendedIntent := paymentstypes.PaymentIntent{
		Payer:    s.suspendedUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "suspended payment",
	}

	_, err = s.paymentsKeeper.CreatePayment(s.ctx, suspendedIntent)
	s.Require().Error(err, "Suspended user should be blocked")
	s.Require().Contains(err.Error(), "blocked")

	// Test 4: Sanctioned user cannot create payment
	sanctionedIntent := paymentstypes.PaymentIntent{
		Payer:    s.sanctionedUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "sanctioned payment",
	}

	_, err = s.paymentsKeeper.CreatePayment(s.ctx, sanctionedIntent)
	s.Require().Error(err, "Sanctioned user should be blocked")
	s.Require().Contains(err.Error(), "sanction")

	// Test 5: Cannot send to sanctioned recipient
	toSanctionedIntent := paymentstypes.PaymentIntent{
		Payer:    s.compliantUser.String(),
		Payee:    s.sanctionedUser.String(),
		Amount:   amount,
		Metadata: "payment to sanctioned",
	}

	_, err = s.paymentsKeeper.CreatePayment(s.ctx, toSanctionedIntent)
	s.Require().Error(err, "Payment to sanctioned user should be blocked")
}

// TestSettlementComplianceIntegration tests settlement module with compliance
func (s *CrossModuleComplianceTestSuite) TestSettlementComplianceIntegration() {
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000))

	// Test 1: Compliant users can settle
	settlementID, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.compliantUser.String(),
		s.merchant.String(),
		amount,
		"COMPLIANT-001",
		"test settlement",
	)
	s.Require().NoError(err, "Compliant settlement should succeed")
	s.Require().Greater(settlementID, uint64(0))

	// Test 2: Suspended user cannot settle
	_, err = s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.suspendedUser.String(),
		s.merchant.String(),
		amount,
		"SUSPENDED-001",
		"suspended settlement",
	)
	s.Require().Error(err, "Suspended user settlement should fail")

	// Test 3: Cannot settle to sanctioned recipient
	_, err = s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.compliantUser.String(),
		s.sanctionedUser.String(),
		amount,
		"TO-SANCTIONED-001",
		"settlement to sanctioned",
	)
	s.Require().Error(err, "Settlement to sanctioned should fail")

	// Test 4: Escrow with compliance checks
	escrowID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.compliantUser.String(),
		s.merchant.String(),
		amount,
		"ESCROW-001",
		"escrow test",
		3600,
	)
	s.Require().NoError(err, "Compliant escrow should succeed")
	s.Require().Greater(escrowID, uint64(0))

	// Test 5: Expired user cannot create escrow
	_, err = s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.expiredUser.String(),
		s.merchant.String(),
		amount,
		"EXPIRED-ESCROW",
		"expired escrow",
		3600,
	)
	s.Require().Error(err, "Expired user escrow should fail")
}

// TestTransactionLimitsEnforcement tests daily/monthly limit enforcement
func (s *CrossModuleComplianceTestSuite) TestTransactionLimitsEnforcement() {
	// Test 1: User at daily limit cannot transact more
	overLimitAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(50000000)) // Pushes over limit

	_, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.limitExceededUser.String(),
		s.merchant.String(),
		overLimitAmount,
		"OVER-LIMIT",
		"over limit test",
	)
	s.Require().Error(err, "Over limit transaction should fail")
	s.Require().Contains(err.Error(), "limit")

	// Test 2: User under limit can transact
	underLimitAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(500000)) // Under remaining limit

	settlementID, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.limitExceededUser.String(),
		s.merchant.String(),
		underLimitAmount,
		"UNDER-LIMIT",
		"under limit test",
	)
	s.Require().NoError(err, "Under limit transaction should succeed")
	s.Require().Greater(settlementID, uint64(0))

	// Test 3: Limits reset after time period
	profile, _ := s.complianceKeeper.GetProfile(s.ctx, s.limitExceededUser)
	s.Require().True(profile.DailyUsed.Amount.GT(sdkmath.ZeroInt()))

	// Advance time by 25 hours
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(25 * time.Hour))

	// Check if compliance allows transaction (limits should have reset)
	err = s.complianceKeeper.AssertCompliantForAmount(s.ctx, s.limitExceededUser, overLimitAmount)
	s.Require().NoError(err, "Should succeed after limit reset")
}

// TestComplianceProfileManagement tests profile updates and their effect
func (s *CrossModuleComplianceTestSuite) TestComplianceProfileManagement() {
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000000))

	// Test 1: Suspend active user
	err := s.complianceKeeper.SuspendProfile(
		s.ctx,
		s.compliantUser,
		s.authority.String(),
		"suspicious activity",
	)
	s.Require().NoError(err, "Profile suspension should succeed")

	// Verify user can no longer transact
	_, err = s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.compliantUser.String(),
		s.merchant.String(),
		amount,
		"AFTER-SUSPEND",
		"after suspension",
	)
	s.Require().Error(err, "Suspended user should not transact")

	// Test 2: Reactivate suspended user
	err = s.complianceKeeper.ReactivateProfile(
		s.ctx,
		s.compliantUser,
		s.authority.String(),
		"investigation cleared",
	)
	s.Require().NoError(err, "Profile reactivation should succeed")

	// Verify user can transact again
	settlementID, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.compliantUser.String(),
		s.merchant.String(),
		amount,
		"AFTER-REACTIVATE",
		"after reactivation",
	)
	s.Require().NoError(err, "Reactivated user should transact")
	s.Require().Greater(settlementID, uint64(0))

	// Test 3: Update KYC level
	err = s.complianceKeeper.UpdateKYCLevel(
		s.ctx,
		s.lowKYCUser,
		compliancetypes.KYCEnhanced,
		s.authority.String(),
		"verification completed",
	)
	s.Require().NoError(err, "KYC update should succeed")

	// Verify updated profile
	profile, found := s.complianceKeeper.GetProfile(s.ctx, s.lowKYCUser)
	s.Require().True(found)
	s.Require().Equal(compliancetypes.KYCEnhanced, profile.KYCLevel)
}

// TestAuditTrail tests compliance audit logging
func (s *CrossModuleComplianceTestSuite) TestAuditTrail() {
	// Suspend user
	err := s.complianceKeeper.SuspendProfile(
		s.ctx,
		s.compliantUser,
		s.authority.String(),
		"test audit trail",
	)
	s.Require().NoError(err)

	// Get profile and check audit entries
	profile, found := s.complianceKeeper.GetProfile(s.ctx, s.compliantUser)
	s.Require().True(found)
	s.Require().Greater(len(profile.AuditLog), 0, "Audit log should have entries")

	// Verify audit entry details
	lastEntry := profile.AuditLog[len(profile.AuditLog)-1]
	s.Require().Equal("suspended", lastEntry.Action)
	s.Require().Equal(s.authority.String(), lastEntry.Actor)
	s.Require().Equal("test audit trail", lastEntry.Reason)
}

// TestBatchComplianceChecks tests batch operations with compliance
func (s *CrossModuleComplianceTestSuite) TestBatchComplianceChecks() {
	// Setup additional compliant users
	user2 := sdk.AccAddress([]byte("user2_______________"))
	user3 := sdk.AccAddress([]byte("user3_______________"))

	for _, addr := range []sdk.AccAddress{user2, user3} {
		acc := s.accountKeeper.NewAccountWithAddress(s.ctx, addr)
		s.accountKeeper.SetAccount(s.ctx, acc)

		coins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)))
		require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, coins))
		require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, addr, coins))

		profile := compliancetypes.Profile{
			Address:    addr.String(),
			Status:     compliancetypes.StatusActive,
			KYCLevel:   compliancetypes.KYCStandard,
			Sanction:   false,
			VerifiedAt: s.ctx.BlockTime(),
			ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
			LastLimitReset: s.ctx.BlockTime(),
		}
		s.complianceKeeper.SetProfile(s.ctx, profile)
	}

	// Test 1: Batch with all compliant users
	senders := []string{s.compliantUser.String(), user2.String(), user3.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(150000000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000000)),
	}
	references := []string{"BATCH-001", "BATCH-002", "BATCH-003"}

	batchID, settlementIDs, err := s.settlementKeeper.CreateBatch(
		s.ctx,
		s.merchant.String(),
		senders,
		amounts,
		references,
	)
	s.Require().NoError(err, "Batch with compliant users should succeed")
	s.Require().Greater(batchID, uint64(0))
	s.Require().Len(settlementIDs, 3)

	// Test 2: Batch with one non-compliant user should fail
	sendersWithSanctioned := []string{s.compliantUser.String(), s.sanctionedUser.String(), user2.String()}

	_, _, err = s.settlementKeeper.CreateBatch(
		s.ctx,
		s.merchant.String(),
		sendersWithSanctioned,
		amounts,
		references,
	)
	s.Require().Error(err, "Batch with sanctioned user should fail")
}

// TestComplianceEventEmission tests that compliance events are emitted
func (s *CrossModuleComplianceTestSuite) TestComplianceEventEmission() {
	// Suspend user
	err := s.complianceKeeper.SuspendProfile(
		s.ctx,
		s.compliantUser,
		s.authority.String(),
		"event test",
	)
	s.Require().NoError(err)

	// Check events
	events := s.ctx.EventManager().Events()
	s.Require().Greater(len(events), 0)

	var foundSuspendEvent bool
	for _, event := range events {
		if event.Type == "compliance_profile_suspended" {
			foundSuspendEvent = true
			// Verify attributes
			for _, attr := range event.Attributes {
				if attr.Key == "address" {
					s.Require().Equal(s.compliantUser.String(), attr.Value)
				}
			}
		}
	}
	s.Require().True(foundSuspendEvent, "Suspend event should be emitted")
}

// TestCrossModuleComplianceConsistency tests consistency across modules
func (s *CrossModuleComplianceTestSuite) TestCrossModuleComplianceConsistency() {
	// Scenario: Verify compliance checks are consistent across payments and settlements

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(300000000))

	// Test 1: Both modules reject sanctioned user
	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.sanctionedUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "consistency test",
	}

	_, paymentErr := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().Error(paymentErr)

	_, settlementErr := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.sanctionedUser.String(),
		s.merchant.String(),
		amount,
		"CONSISTENCY-TEST",
		"consistency test",
	)
	s.Require().Error(settlementErr)

	// Both should fail for same reason (sanctioned)
	s.Require().Contains(paymentErr.Error(), "sanction")
	s.Require().Contains(settlementErr.Error(), "sanctioned")
}

// TestComplianceRecordTransaction tests transaction recording
func (s *CrossModuleComplianceTestSuite) TestComplianceRecordTransaction() {
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000))

	// Get initial usage
	profile, _ := s.complianceKeeper.GetProfile(s.ctx, s.compliantUser)
	initialDailyUsed := profile.DailyUsed.Amount
	initialMonthlyUsed := profile.MonthlyUsed.Amount

	// Execute settlement
	_, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.compliantUser.String(),
		s.merchant.String(),
		amount,
		"RECORD-TEST",
		"record transaction test",
	)
	s.Require().NoError(err)

	// Record transaction
	err = s.complianceKeeper.RecordTransaction(s.ctx, s.compliantUser, amount)
	s.Require().NoError(err)

	// Verify usage was updated
	profile, _ = s.complianceKeeper.GetProfile(s.ctx, s.compliantUser)
	s.Require().Equal(initialDailyUsed.Add(amount.Amount), profile.DailyUsed.Amount)
	s.Require().Equal(initialMonthlyUsed.Add(amount.Amount), profile.MonthlyUsed.Amount)
}

// TestComplianceFailurePreventsSettlement tests that compliance failure stops settlement
func (s *CrossModuleComplianceTestSuite) TestComplianceFailurePreventsSettlement() {
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(200000000))

	// Record initial balance
	initialMerchantBalance := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")

	// Try settlement from suspended user
	_, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.suspendedUser.String(),
		s.merchant.String(),
		amount,
		"FAIL-TEST",
		"should fail",
	)
	s.Require().Error(err, "Settlement should fail")

	// Verify merchant did not receive funds
	finalMerchantBalance := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	s.Require().Equal(initialMerchantBalance.Amount, finalMerchantBalance.Amount,
		"Merchant balance should not change on failed settlement")
}
