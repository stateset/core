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
	orderskeeper "github.com/stateset/core/x/orders/keeper"
	orderstypes "github.com/stateset/core/x/orders/types"
	paymentskeeper "github.com/stateset/core/x/payments/keeper"
	paymentstypes "github.com/stateset/core/x/payments/types"
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
)

// ECommerceFlowTestSuite tests the complete e-commerce payment workflow:
// User payment -> Compliance check -> Settlement -> Merchant receives funds
type ECommerceFlowTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	cdc            codec.Codec
	storeKey       storetypes.StoreKey

	// Keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	complianceKeeper compliancekeeper.Keeper
	ordersKeeper     orderskeeper.Keeper
	paymentsKeeper   paymentskeeper.Keeper
	settlementKeeper settlementkeeper.Keeper

	// Test accounts
	authority      sdk.AccAddress
	customer       sdk.AccAddress
	merchant       sdk.AccAddress
	blockedUser    sdk.AccAddress
}

func TestECommerceFlowTestSuite(t *testing.T) {
	suite.Run(t, new(ECommerceFlowTestSuite))
}

func (s *ECommerceFlowTestSuite) SetupTest() {
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
		orderstypes.StoreKey,
		paymentstypes.StoreKey,
		settlementtypes.StoreKey,
	)
	s.storeKey = storeKeys[settlementtypes.StoreKey]

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
	s.customer = sdk.AccAddress([]byte("customer____________"))
	s.merchant = sdk.AccAddress([]byte("merchant____________"))
	s.blockedUser = sdk.AccAddress([]byte("blocked_____________"))

	// Initialize account keeper
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:       nil,
		paymentstypes.ModuleAccountName:  {authtypes.Minter, authtypes.Burner},
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

	// Initialize orders keeper
	s.ordersKeeper = orderskeeper.NewKeeper(
		s.cdc,
		storeKeys[orderstypes.StoreKey],
		s.authority.String(),
		s.bankKeeper,
		s.complianceKeeper,
		s.settlementKeeper,
		s.accountKeeper,
	)

	// Setup test data
	s.setupTestAccounts()
	s.setupComplianceProfiles()

	// Configure orders module params
	ordersParams := orderstypes.DefaultParams()
	ordersParams.StablecoinDenom = "ssusd"
	ordersParams.MinOrderAmount = sdk.NewCoin("ssusd", sdkmath.NewInt(100))
	ordersParams.MaxOrderAmount = sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000000))
	err = s.ordersKeeper.SetParams(s.ctx, ordersParams)
	require.NoError(s.T(), err)
}

func (s *ECommerceFlowTestSuite) setupTestAccounts() {
	// Create module accounts
	paymentsModuleAcc := authtypes.NewEmptyModuleAccount(paymentstypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	paymentsModuleAcc = s.accountKeeper.NewAccount(s.ctx, paymentsModuleAcc).(*authtypes.ModuleAccount)
	settlementModuleAcc := authtypes.NewEmptyModuleAccount(settlementtypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	settlementModuleAcc = s.accountKeeper.NewAccount(s.ctx, settlementModuleAcc).(*authtypes.ModuleAccount)
	ordersModuleAcc := authtypes.NewEmptyModuleAccount(orderstypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	ordersModuleAcc = s.accountKeeper.NewAccount(s.ctx, ordersModuleAcc).(*authtypes.ModuleAccount)

	s.accountKeeper.SetModuleAccount(s.ctx, paymentsModuleAcc)
	s.accountKeeper.SetModuleAccount(s.ctx, settlementModuleAcc)
	s.accountKeeper.SetModuleAccount(s.ctx, ordersModuleAcc)

	// Create user accounts
	customerAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.customer)
	merchantAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.merchant)
	blockedAcc := s.accountKeeper.NewAccountWithAddress(s.ctx, s.blockedUser)
	s.accountKeeper.SetAccount(s.ctx, customerAcc)
	s.accountKeeper.SetAccount(s.ctx, merchantAcc)
	s.accountKeeper.SetAccount(s.ctx, blockedAcc)

	// Mint coins for customer
	initialCoins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000))) // 1000 ssUSD
	require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, initialCoins))
	require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, s.customer, initialCoins))
}

func (s *ECommerceFlowTestSuite) setupComplianceProfiles() {
	// Setup compliant customer profile
	customerProfile := compliancetypes.Profile{
		Address:     s.customer.String(),
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
	s.complianceKeeper.SetProfile(s.ctx, customerProfile)

	// Setup compliant merchant profile
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

	// Setup blocked user profile
	blockedProfile := compliancetypes.Profile{
		Address:  s.blockedUser.String(),
		Status:   compliancetypes.StatusSuspended,
		KYCLevel: compliancetypes.KYCNone,
		Sanction: true,
	}
	s.complianceKeeper.SetProfile(s.ctx, blockedProfile)
}

// TestSuccessfulECommerceFlow tests the happy path: customer pays merchant successfully
func (s *ECommerceFlowTestSuite) TestSuccessfulECommerceFlow() {
	// Scenario: Customer purchases item for 100 ssUSD, payment flows through system to merchant

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000000)) // 100 ssUSD
	reference := "ORDER-12345"
	metadata := `{"item": "Widget Pro", "quantity": 1, "price": 100.00}`

	// Step 1: Verify customer has sufficient balance
	initialCustomerBalance := s.bankKeeper.GetBalance(s.ctx, s.customer, "ssusd")
	s.Require().True(initialCustomerBalance.Amount.GTE(amount.Amount), "Customer should have sufficient balance")

	// Step 2: Verify compliance for both parties
	err := s.complianceKeeper.AssertCompliant(s.ctx, s.customer)
	s.Require().NoError(err, "Customer should be compliant")

	err = s.complianceKeeper.AssertCompliant(s.ctx, s.merchant)
	s.Require().NoError(err, "Merchant should be compliant")

	// Step 3: Customer creates payment intent
	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.customer.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: metadata,
	}

	paymentID, err := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().NoError(err, "Payment creation should succeed")
	s.Require().Greater(paymentID, uint64(0), "Payment ID should be assigned")

	// Verify payment was created and escrowed
	payment, found := s.paymentsKeeper.GetPayment(s.ctx, paymentID)
	s.Require().True(found, "Payment should exist")
	s.Require().Equal(paymentstypes.PaymentStatusPending, payment.Status)

	// Verify funds moved to escrow
	customerBalanceAfterPayment := s.bankKeeper.GetBalance(s.ctx, s.customer, "ssusd")
	expectedBalance := initialCustomerBalance.Amount.Sub(amount.Amount)
	s.Require().Equal(expectedBalance, customerBalanceAfterPayment.Amount, "Customer balance should decrease")

	// Step 4: Execute instant settlement
	settlementID, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.customer.String(),
		s.merchant.String(),
		amount,
		reference,
		metadata,
	)
	s.Require().NoError(err, "Settlement should succeed")
	s.Require().Greater(settlementID, uint64(0), "Settlement ID should be assigned")

	// Step 5: Verify settlement was created and completed
	settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found, "Settlement should exist")
	s.Require().Equal(settlementtypes.SettlementStatusCompleted, settlement.Status)
	s.Require().Equal(settlementtypes.SettlementTypeInstant, settlement.Type)
	s.Require().Equal(s.customer.String(), settlement.Sender)
	s.Require().Equal(s.merchant.String(), settlement.Recipient)
	s.Require().Equal(amount, settlement.Amount)
	s.Require().Equal(reference, settlement.Reference)

	// Step 6: Verify merchant received funds (minus fee)
	merchantBalance := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	expectedMerchantReceived := settlement.NetAmount.Amount
	s.Require().Equal(expectedMerchantReceived, merchantBalance.Amount, "Merchant should receive net amount")

	// Step 7: Verify fee was collected
	s.Require().True(settlement.Fee.Amount.GT(sdkmath.ZeroInt()), "Fee should be positive")
	totalWithFee := settlement.NetAmount.Amount.Add(settlement.Fee.Amount)
	s.Require().Equal(amount.Amount, totalWithFee, "Net amount + fee should equal gross amount")

	// Step 8: Verify events were emitted
	events := s.ctx.EventManager().Events()
	s.Require().Greater(len(events), 0, "Events should be emitted")

	// Look for settlement completed event
	var foundSettlementEvent bool
	for _, event := range events {
		if event.Type == settlementtypes.EventTypeInstantTransfer {
			foundSettlementEvent = true
			break
		}
	}
	s.Require().True(foundSettlementEvent, "Settlement completed event should be emitted")
}

// TestBlockedUserCannotPay tests that sanctioned users cannot make payments
func (s *ECommerceFlowTestSuite) TestBlockedUserCannotPay() {
	// Scenario: Blocked user attempts to make payment, should fail at compliance check

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(50000000)) // 50 ssUSD

	// Verify blocked user fails compliance check
	err := s.complianceKeeper.AssertCompliant(s.ctx, s.blockedUser)
	s.Require().Error(err, "Blocked user should fail compliance check")
	s.Require().Contains(err.Error(), "sanctioned", "Error should mention sanction")

	// Attempt to create payment should fail
	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.blockedUser.String(),
		Payee:    s.merchant.String(),
		Amount:   amount,
		Metadata: "Blocked user payment",
	}

	_, err = s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().Error(err, "Payment from blocked user should fail")
}

// TestInsufficientFunds tests payment failure when customer has insufficient balance
func (s *ECommerceFlowTestSuite) TestInsufficientFunds() {
	// Scenario: Customer attempts to pay more than available balance

	customerBalance := s.bankKeeper.GetBalance(s.ctx, s.customer, "ssusd")
	excessiveAmount := sdk.NewCoin("ssusd", customerBalance.Amount.Add(sdkmath.NewInt(1000000)))

	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.customer.String(),
		Payee:    s.merchant.String(),
		Amount:   excessiveAmount,
		Metadata: "Insufficient funds test",
	}

	_, err := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().Error(err, "Payment should fail with insufficient funds")
	s.Require().Contains(err.Error(), "insufficient", "Error should mention insufficient balance")
}

// TestPaymentToBlockedMerchant tests that payments to sanctioned merchants are blocked
func (s *ECommerceFlowTestSuite) TestPaymentToBlockedMerchant() {
	// Scenario: Customer attempts to pay sanctioned merchant

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000000))

	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.customer.String(),
		Payee:    s.blockedUser.String(),
		Amount:   amount,
		Metadata: "Payment to blocked merchant",
	}

	_, err := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().Error(err, "Payment to blocked merchant should fail")
}

// TestEscrowSettlement tests escrow-based settlement flow
func (s *ECommerceFlowTestSuite) TestEscrowSettlement() {
	// Scenario: Create escrow settlement, then release funds

	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(200000000)) // 200 ssUSD
	reference := "ESCROW-001"
	expirationSeconds := int64(3600) // 1 hour

	initialCustomerBalance := s.bankKeeper.GetBalance(s.ctx, s.customer, "ssusd")

	// Step 1: Create escrow settlement
	settlementID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.customer.String(),
		s.merchant.String(),
		amount,
		reference,
		"Escrow payment",
		expirationSeconds,
	)
	s.Require().NoError(err, "Escrow creation should succeed")

	// Verify settlement is pending
	settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusPending, settlement.Status)
	s.Require().Equal(settlementtypes.SettlementTypeEscrow, settlement.Type)

	// Verify funds are held in escrow (customer balance decreased)
	customerBalanceAfterEscrow := s.bankKeeper.GetBalance(s.ctx, s.customer, "ssusd")
	s.Require().Equal(initialCustomerBalance.Amount.Sub(amount.Amount), customerBalanceAfterEscrow.Amount)

	// Merchant should not have received funds yet
	merchantBalance := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	s.Require().True(merchantBalance.IsZero())

	// Step 2: Release escrow
	err = s.settlementKeeper.ReleaseEscrow(s.ctx, settlementID, s.customer)
	s.Require().NoError(err, "Escrow release should succeed")

	// Verify settlement is completed
	settlement, found = s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusCompleted, settlement.Status)

	// Merchant should now have funds (minus fee)
	merchantBalance = s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	s.Require().Equal(settlement.NetAmount.Amount, merchantBalance.Amount)
}

// TestBatchSettlement tests batch processing of multiple payments
func (s *ECommerceFlowTestSuite) TestBatchSettlement() {
	// Scenario: Multiple customers pay merchant, settled in batch

	// Create additional customers
	customer2 := sdk.AccAddress([]byte("customer2___________"))
	customer3 := sdk.AccAddress([]byte("customer3___________"))

	// Setup accounts and compliance for additional customers
	for _, addr := range []sdk.AccAddress{customer2, customer3} {
		acc := s.accountKeeper.NewAccountWithAddress(s.ctx, addr)
		s.accountKeeper.SetAccount(s.ctx, acc)

		// Mint coins
		coins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(500000000)))
		require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, coins))
		require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, addr, coins))

		// Setup compliance
		profile := compliancetypes.Profile{
			Address:    addr.String(),
			Status:     compliancetypes.StatusActive,
			KYCLevel:   compliancetypes.KYCStandard,
			Sanction:   false,
			DailyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000)),
			DailyUsed:  sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
			VerifiedAt: s.ctx.BlockTime(),
			ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
			LastLimitReset: s.ctx.BlockTime(),
		}
		s.complianceKeeper.SetProfile(s.ctx, profile)
	}

	// Create batch settlement
	senders := []string{s.customer.String(), customer2.String(), customer3.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(150000000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000000)),
	}
	references := []string{"ORDER-001", "ORDER-002", "ORDER-003"}

	batchID, settlementIDs, err := s.settlementKeeper.CreateBatch(
		s.ctx,
		s.merchant.String(),
		senders,
		amounts,
		references,
	)
	s.Require().NoError(err, "Batch creation should succeed")
	s.Require().Greater(batchID, uint64(0))
	s.Require().Len(settlementIDs, 3)

	// Verify batch exists
	batch, found := s.settlementKeeper.GetBatch(s.ctx, batchID)
	s.Require().True(found)
	s.Require().Equal(uint64(3), batch.Count)
	s.Require().Equal(settlementtypes.SettlementStatusPending, batch.Status)

	// Settle the batch
	err = s.settlementKeeper.SettleBatch(s.ctx, batchID, s.authority.String())
	s.Require().NoError(err, "Batch settlement should succeed")

	// Verify batch is completed
	batch, found = s.settlementKeeper.GetBatch(s.ctx, batchID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusCompleted, batch.Status)

	// Verify merchant received net amount
	merchantBalance := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	s.Require().Equal(batch.NetAmount.Amount, merchantBalance.Amount)
}

// TestStateChangesAreConsistent verifies all state changes are consistent
func (s *ECommerceFlowTestSuite) TestStateChangesAreConsistent() {
	// Scenario: Verify balances sum to expected totals before and after operations

	// Calculate total supply before
	totalSupplyBefore := s.bankKeeper.GetSupply(s.ctx, "ssusd")

	// Execute payment flow
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000000))
	settlementID, err := s.settlementKeeper.InstantTransfer(
		s.ctx,
		s.customer.String(),
		s.merchant.String(),
		amount,
		"TEST-001",
		"consistency test",
	)
	s.Require().NoError(err)

	// Calculate total supply after
	totalSupplyAfter := s.bankKeeper.GetSupply(s.ctx, "ssusd")

	// Total supply should remain constant (no minting/burning in transfers)
	s.Require().Equal(totalSupplyBefore.Amount, totalSupplyAfter.Amount, "Total supply should remain constant")

	// Verify settlement accounting
	settlement, _ := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().Equal(amount.Amount, settlement.NetAmount.Amount.Add(settlement.Fee.Amount),
		"Amount should equal net + fee")
}

// TestFullOrderLifecycle tests the complete order management flow linked with settlement
func (s *ECommerceFlowTestSuite) TestFullOrderLifecycle() {
	// Scenario: Full lifecycle from Order Creation -> Confirmation -> Payment -> Shipping -> Delivery -> Completion

	// Verify params are set correctly
	params := s.ordersKeeper.GetParams(s.ctx)
	s.Require().Equal("ssusd", params.StablecoinDenom)

	// 1. Create Order
	items := []orderstypes.OrderItem{
		{
			Id:          "ITEM-1",
			ProductName: "Stablecoin Handbook",
			Quantity:    1,
			UnitPrice:   sdk.NewCoin("ssusd", sdkmath.NewInt(50000000)), // 50 ssUSD
		},
	}
	shipping := orderstypes.ShippingInfo{
		Address: orderstypes.Address{
			Name:       "John Doe",
			Line1:      "123 Blockchain Blvd",
			City:       "Crypto City",
			Country:    "Internet",
			PostalCode: "10101",
		},
		Method: "Standard",
	}

	orderID, err := s.ordersKeeper.CreateOrder(
		s.ctx,
		s.customer.String(),
		s.merchant.String(),
		items,
		shipping,
		"Order metadata",
	)
	s.Require().NoError(err, "Order creation should succeed")
	s.Require().NotEmpty(orderID)

	// Verify status PENDING
	order, found := s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().True(found)
	s.Require().Equal(orderstypes.OrderStatusPending, order.Status)

	// 2. Merchant Confirms Order
	err = s.ordersKeeper.ConfirmOrder(s.ctx, s.merchant.String(), orderID)
	s.Require().NoError(err, "Order confirmation should succeed")

	// Verify status CONFIRMED
	order, _ = s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().Equal(orderstypes.OrderStatusConfirmed, order.Status)

	// 3. Customer Pays for Order (using Escrow)
	amountToPay := sdk.NewCoin("ssusd", sdkmath.NewInt(50000000))
	err = s.ordersKeeper.PayOrder(s.ctx, s.customer.String(), orderID, amountToPay, true)
	s.Require().NoError(err, "Order payment should succeed")

	// Verify status PAID and funds escrowed
	order, _ = s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().Equal(orderstypes.OrderStatusPaid, order.Status)
	s.Require().NotEmpty(order.SettlementId)

	// 4. Merchant Ships Order
	err = s.ordersKeeper.ShipOrder(s.ctx, s.merchant.String(), orderID, "FedEx", "TRACK-999")
	s.Require().NoError(err, "Order shipping should succeed")

	// Verify status SHIPPED
	order, _ = s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().Equal(orderstypes.OrderStatusShipped, order.Status)

	// 5. Customer marks Delivered
	err = s.ordersKeeper.DeliverOrder(s.ctx, s.customer.String(), orderID)
	s.Require().NoError(err, "Order delivery should succeed")

	// Verify status DELIVERED
	order, _ = s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().Equal(orderstypes.OrderStatusDelivered, order.Status)

	// 6. Complete Order (Release Escrow)
	// Capture merchant balance before release
	merchantBalBefore := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")

	err = s.ordersKeeper.CompleteOrder(s.ctx, s.customer.String(), orderID)
	s.Require().NoError(err, "Order completion should succeed")

	// Verify status COMPLETED
	order, _ = s.ordersKeeper.GetOrder(s.ctx, orderID)
	s.Require().Equal(orderstypes.OrderStatusCompleted, order.Status)

	// Verify merchant got paid
	merchantBalAfter := s.bankKeeper.GetBalance(s.ctx, s.merchant, "ssusd")
	s.Require().True(merchantBalAfter.Amount.GT(merchantBalBefore.Amount), "Merchant should receive funds")
}

