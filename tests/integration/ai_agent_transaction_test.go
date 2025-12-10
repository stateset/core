package integration

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
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

// AIAgentTransactionTestSuite tests AI agent autonomous transaction workflow:
// Agent message -> Negotiate -> Payment intent -> Escrow -> Service delivery -> Release
type AIAgentTransactionTestSuite struct {
	suite.Suite

	ctx            sdk.Context
	cdc            codec.Codec
	storeKey       storetypes.StoreKey

	// Keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	complianceKeeper compliancekeeper.Keeper
	paymentsKeeper   paymentskeeper.Keeper
	settlementKeeper settlementkeeper.Keeper

	// Test accounts - AI Agents and users
	authority       sdk.AccAddress
	aiAgent         sdk.AccAddress // AI service provider agent
	user            sdk.AccAddress // User requesting service
	serviceProvider sdk.AccAddress // Human service provider (for fallback)
	arbiter         sdk.AccAddress // Dispute resolution arbiter
}

func TestAIAgentTransactionTestSuite(t *testing.T) {
	suite.Run(t, new(AIAgentTransactionTestSuite))
}

func (s *AIAgentTransactionTestSuite) SetupTest() {
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
	s.storeKey = storeKeys[settlementtypes.StoreKey]

	// Create transient store keys
	tKeys := storetypes.NewTransientStoreKeys(banktypes.TransientKey)

	// Create multistore
	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range storeKeys {
		stateStore.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		stateStore.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)
	}
	require.NoError(s.T(), stateStore.LoadLatestVersion())

	// Create context
	s.ctx = testutil.DefaultContextWithDB(s.T(), storeKeys[authtypes.StoreKey], tKeys[banktypes.TransientKey]).Ctx.
		WithBlockHeight(1).
		WithBlockTime(time.Now())

	// Initialize test addresses
	s.authority = sdk.AccAddress([]byte("authority___________"))
	s.aiAgent = sdk.AccAddress([]byte("aiagent_____________"))
	s.user = sdk.AccAddress([]byte("user________________"))
	s.serviceProvider = sdk.AccAddress([]byte("serviceprovider_____"))
	s.arbiter = sdk.AccAddress([]byte("arbiter_____________"))

	// Initialize account keeper
	maccPerms := map[string][]string{
		authtypes.FeeCollectorName:         nil,
		paymentstypes.ModuleAccountName:    {authtypes.Minter, authtypes.Burner},
		settlementtypes.ModuleAccountName:  {authtypes.Minter, authtypes.Burner},
	}
	s.accountKeeper = authkeeper.NewAccountKeeper(
		s.cdc,
		runtime.NewKVStoreService(storeKeys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		maccPerms,
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

func (s *AIAgentTransactionTestSuite) setupTestAccounts() {
	// Create module accounts
	paymentsModuleAcc := authtypes.NewEmptyModuleAccount(paymentstypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	settlementModuleAcc := authtypes.NewEmptyModuleAccount(settlementtypes.ModuleAccountName, authtypes.Minter, authtypes.Burner)
	s.accountKeeper.SetModuleAccount(s.ctx, paymentsModuleAcc)
	s.accountKeeper.SetModuleAccount(s.ctx, settlementModuleAcc)

	// Create user accounts
	aiAgentAcc := authtypes.NewBaseAccountWithAddress(s.aiAgent)
	userAcc := authtypes.NewBaseAccountWithAddress(s.user)
	providerAcc := authtypes.NewBaseAccountWithAddress(s.serviceProvider)
	arbiterAcc := authtypes.NewBaseAccountWithAddress(s.arbiter)
	s.accountKeeper.SetAccount(s.ctx, aiAgentAcc)
	s.accountKeeper.SetAccount(s.ctx, userAcc)
	s.accountKeeper.SetAccount(s.ctx, providerAcc)
	s.accountKeeper.SetAccount(s.ctx, arbiterAcc)

	// Mint coins for user
	userCoins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000000))) // 10,000 ssUSD
	require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, userCoins))
	require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, s.user, userCoins))
}

func (s *AIAgentTransactionTestSuite) setupComplianceProfiles() {
	// Setup profiles for all parties
	addresses := []sdk.AccAddress{s.aiAgent, s.user, s.serviceProvider, s.arbiter}
	for _, addr := range addresses {
		profile := compliancetypes.Profile{
			Address:    addr.String(),
			Status:     compliancetypes.StatusActive,
			KYCLevel:   compliancetypes.KYCStandard,
			Sanction:   false,
			DailyLimit: sdk.NewCoin("ssusd", sdkmath.NewInt(100000000000)),
			DailyUsed:  sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
			VerifiedAt: s.ctx.BlockTime(),
			ExpiresAt:  s.ctx.BlockTime().AddDate(1, 0, 0),
			LastLimitReset: s.ctx.BlockTime(),
		}
		s.complianceKeeper.SetProfile(s.ctx, profile)
	}
}

// TestSuccessfulAIAgentTransaction tests the complete AI agent workflow
func (s *AIAgentTransactionTestSuite) TestSuccessfulAIAgentTransaction() {
	// Scenario: User requests AI service, agent accepts, escrow created,
	// service delivered, payment released

	servicePrice := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)) // 1000 ssUSD
	serviceReference := "AI-SERVICE-GPT4-ANALYSIS-001"
	serviceMetadata := `{
		"service_type": "ai_analysis",
		"model": "gpt-4",
		"task": "financial_report_analysis",
		"estimated_completion": "10 minutes",
		"user_id": "user123",
		"agent_id": "ai-agent-001"
	}`

	initialUserBalance := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")
	initialAgentBalance := s.bankKeeper.GetBalance(s.ctx, s.aiAgent, "ssusd")

	// Step 1: User initiates service request with payment intent
	// This simulates the user sending a message to AI agent requesting service
	paymentIntent := paymentstypes.PaymentIntent{
		Payer:    s.user.String(),
		Payee:    s.aiAgent.String(),
		Amount:   servicePrice,
		Metadata: serviceMetadata,
	}

	paymentID, err := s.paymentsKeeper.CreatePayment(s.ctx, paymentIntent)
	s.Require().NoError(err, "Payment intent creation should succeed")
	s.Require().Greater(paymentID, uint64(0))

	// Verify payment intent is pending
	payment, found := s.paymentsKeeper.GetPayment(s.ctx, paymentID)
	s.Require().True(found)
	s.Require().Equal(paymentstypes.PaymentStatusPending, payment.Status)

	// Step 2: AI Agent negotiates terms (simulated by metadata update)
	// In real scenario, agent would analyze request and confirm capability
	// For this test, we assume agreement is reached

	// Step 3: Create escrow settlement to hold funds during service delivery
	expirationSeconds := int64(7200) // 2 hours for service completion

	settlementID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.user.String(),
		s.aiAgent.String(),
		servicePrice,
		serviceReference,
		serviceMetadata,
		expirationSeconds,
	)
	s.Require().NoError(err, "Escrow creation should succeed")

	// Verify escrow is created and funds are held
	settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusPending, settlement.Status)
	s.Require().Equal(settlementtypes.SettlementTypeEscrow, settlement.Type)
	s.Require().False(settlement.ExpiresAt.IsZero(), "Expiration should be set")

	// Verify user's funds are in escrow
	userBalanceAfterEscrow := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")
	s.Require().Equal(initialUserBalance.Amount.Sub(servicePrice.Amount), userBalanceAfterEscrow.Amount)

	// AI Agent has not received funds yet
	agentBalanceBeforeService := s.bankKeeper.GetBalance(s.ctx, s.aiAgent, "ssusd")
	s.Require().Equal(initialAgentBalance.Amount, agentBalanceBeforeService.Amount)

	// Step 4: AI Agent performs service (simulated by time passage)
	// In real scenario, agent would:
	// - Process the request
	// - Generate results
	// - Submit results to user
	// - Wait for user confirmation or auto-confirm after timeout

	serviceCompletionTime := s.ctx.BlockTime().Add(10 * time.Minute)
	s.ctx = s.ctx.WithBlockTime(serviceCompletionTime)

	// Simulate service completion with results metadata
	completionMetadata := `{
		"status": "completed",
		"completion_time": "2024-01-15T10:15:00Z",
		"results_hash": "0xabc123...",
		"quality_score": 0.98
	}`
	_ = completionMetadata

	// Step 5: User reviews service and releases escrow
	// (In automated scenario, this could be auto-released after confirmation period)
	err = s.settlementKeeper.ReleaseEscrow(s.ctx, settlementID, s.user)
	s.Require().NoError(err, "Escrow release should succeed")

	// Verify settlement is completed
	settlement, found = s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusCompleted, settlement.Status)

	// Step 6: Verify AI agent received payment (minus fee)
	agentBalanceAfterService := s.bankKeeper.GetBalance(s.ctx, s.aiAgent, "ssusd")
	expectedAgentReceived := settlement.NetAmount.Amount
	s.Require().Equal(initialAgentBalance.Amount.Add(expectedAgentReceived), agentBalanceAfterService.Amount)

	// Verify fee was collected
	s.Require().True(settlement.Fee.Amount.GT(sdkmath.ZeroInt()), "Fee should be collected")

	// Step 7: Update payment intent to settled
	err = s.paymentsKeeper.SettlePayment(s.ctx, paymentID, s.aiAgent)
	s.Require().NoError(err, "Payment settlement should succeed")

	payment, found = s.paymentsKeeper.GetPayment(s.ctx, paymentID)
	s.Require().True(found)
	s.Require().Equal(paymentstypes.PaymentStatusSettled, payment.Status)

	// Step 8: Verify events were emitted
	events := s.ctx.EventManager().Events()
	s.Require().Greater(len(events), 0)

	// Verify specific events
	var foundEscrowCreated, foundEscrowReleased bool
	for _, event := range events {
		if event.Type == settlementtypes.EventTypeSettlementCreated {
			foundEscrowCreated = true
		}
		if event.Type == settlementtypes.EventTypeSettlementCompleted {
			foundEscrowReleased = true
		}
	}
	s.Require().True(foundEscrowCreated, "Escrow created event should be emitted")
	s.Require().True(foundEscrowReleased, "Escrow released event should be emitted")
}

// TestAIAgentServiceRefund tests refund scenario when service is unsatisfactory
func (s *AIAgentTransactionTestSuite) TestAIAgentServiceRefund() {
	// Scenario: User requests service, agent fails to deliver, user requests refund

	servicePrice := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000)) // 500 ssUSD
	initialUserBalance := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")

	// Create escrow
	settlementID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.user.String(),
		s.aiAgent.String(),
		servicePrice,
		"AI-SERVICE-REFUND-TEST",
		`{"service": "refund_test"}`,
		3600,
	)
	s.Require().NoError(err)

	// Simulate service failure or user dissatisfaction
	// AI agent (payee) can initiate refund
	refundReason := "Service quality did not meet expectations"
	err = s.settlementKeeper.RefundEscrow(s.ctx, settlementID, s.aiAgent, refundReason)
	s.Require().NoError(err, "Refund should succeed")

	// Verify settlement is refunded
	settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusRefunded, settlement.Status)

	// Verify user received full refund (including fee)
	userBalanceAfterRefund := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")
	s.Require().Equal(initialUserBalance.Amount, userBalanceAfterRefund.Amount,
		"User should receive full refund")
}

// TestEscrowExpiration tests automatic refund when escrow expires
func (s *AIAgentTransactionTestSuite) TestEscrowExpiration() {
	// Scenario: Agent fails to deliver service before escrow expiration

	servicePrice := sdk.NewCoin("ssusd", sdkmath.NewInt(300000000)) // 300 ssUSD
	initialUserBalance := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")

	// Create escrow with short expiration
	expirationSeconds := int64(3600) // 1 hour
	settlementID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.user.String(),
		s.aiAgent.String(),
		servicePrice,
		"AI-SERVICE-TIMEOUT-TEST",
		`{"service": "timeout_test"}`,
		expirationSeconds,
	)
	s.Require().NoError(err)

	settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)

	// Advance time beyond expiration
	expirationTime := settlement.ExpiresAt.Add(1 * time.Minute)
	s.ctx = s.ctx.WithBlockTime(expirationTime)

	// Process expired escrows (normally done in EndBlock)
	s.settlementKeeper.ProcessExpiredEscrows(s.ctx)

	// Verify settlement was auto-cancelled and refunded
	settlement, found = s.settlementKeeper.GetSettlement(s.ctx, settlementID)
	s.Require().True(found)
	s.Require().Equal(settlementtypes.SettlementStatusCancelled, settlement.Status)

	// Verify user received refund
	userBalanceAfterExpiry := s.bankKeeper.GetBalance(s.ctx, s.user, "ssusd")
	s.Require().Equal(initialUserBalance.Amount, userBalanceAfterExpiry.Amount,
		"User should be refunded after expiration")
}

// TestMultipleAIAgentTransactions tests handling multiple concurrent services
func (s *AIAgentTransactionTestSuite) TestMultipleAIAgentTransactions() {
	// Scenario: User requests multiple AI services simultaneously

	services := []struct {
		price     sdkmath.Int
		reference string
		metadata  string
	}{
		{
			price:     sdkmath.NewInt(500000000),
			reference: "AI-SERVICE-001",
			metadata:  `{"service": "text_analysis"}`,
		},
		{
			price:     sdkmath.NewInt(750000000),
			reference: "AI-SERVICE-002",
			metadata:  `{"service": "image_generation"}`,
		},
		{
			price:     sdkmath.NewInt(1000000000),
			reference: "AI-SERVICE-003",
			metadata:  `{"service": "code_review"}`,
		},
	}

	settlementIDs := make([]uint64, len(services))

	// Create escrows for all services
	for i, service := range services {
		settlementID, err := s.settlementKeeper.CreateEscrow(
			s.ctx,
			s.user.String(),
			s.aiAgent.String(),
			sdk.NewCoin("ssusd", service.price),
			service.reference,
			service.metadata,
			7200,
		)
		s.Require().NoError(err)
		settlementIDs[i] = settlementID
	}

	// Verify all escrows are pending
	for _, settlementID := range settlementIDs {
		settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
		s.Require().True(found)
		s.Require().Equal(settlementtypes.SettlementStatusPending, settlement.Status)
	}

	// Complete services in order
	for _, settlementID := range settlementIDs {
		err := s.settlementKeeper.ReleaseEscrow(s.ctx, settlementID, s.user)
		s.Require().NoError(err)
	}

	// Verify all settlements completed
	for _, settlementID := range settlementIDs {
		settlement, found := s.settlementKeeper.GetSettlement(s.ctx, settlementID)
		s.Require().True(found)
		s.Require().Equal(settlementtypes.SettlementStatusCompleted, settlement.Status)
	}
}

// TestAIAgentNonCompliantUser tests rejection of non-compliant users
func (s *AIAgentTransactionTestSuite) TestAIAgentNonCompliantUser() {
	// Scenario: Blocked user tries to request AI service

	// Create blocked user
	blockedUser := sdk.AccAddress([]byte("blockeduser_________"))
	blockedAcc := authtypes.NewBaseAccountWithAddress(blockedUser)
	s.accountKeeper.SetAccount(s.ctx, blockedAcc)

	// Set non-compliant profile
	blockedProfile := compliancetypes.Profile{
		Address:  blockedUser.String(),
		Status:   compliancetypes.StatusSuspended,
		Sanction: true,
	}
	s.complianceKeeper.SetProfile(s.ctx, blockedProfile)

	// Mint coins for blocked user (for testing)
	coins := sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000000)))
	require.NoError(s.T(), s.bankKeeper.MintCoins(s.ctx, settlementtypes.ModuleAccountName, coins))
	require.NoError(s.T(), s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, settlementtypes.ModuleAccountName, blockedUser, coins))

	// Try to create escrow - should fail at compliance check
	servicePrice := sdk.NewCoin("ssusd", sdkmath.NewInt(500000000))
	_, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		blockedUser.String(),
		s.aiAgent.String(),
		servicePrice,
		"BLOCKED-USER-TEST",
		`{"test": "blocked"}`,
		3600,
	)
	s.Require().Error(err, "Blocked user should not be able to create escrow")
}

// TestPaymentChannelForStreamingAI tests payment channels for streaming AI services
func (s *AIAgentTransactionTestSuite) TestPaymentChannelForStreamingAI() {
	// Scenario: User opens payment channel for continuous AI service (e.g., chatbot)

	channelDeposit := sdk.NewCoin("ssusd", sdkmath.NewInt(5000000000)) // 5000 ssUSD
	expiresInBlocks := int64(1000) // Channel expires in 1000 blocks

	// Open payment channel
	channelID, err := s.settlementKeeper.OpenChannel(
		s.ctx,
		s.user.String(),
		s.aiAgent.String(),
		channelDeposit,
		expiresInBlocks,
	)
	s.Require().NoError(err, "Channel opening should succeed")
	s.Require().Greater(channelID, uint64(0))

	// Verify channel is open
	channel, found := s.settlementKeeper.GetChannel(s.ctx, channelID)
	s.Require().True(found)
	s.Require().True(channel.IsOpen)
	s.Require().Equal(channelDeposit, channel.Deposit)
	s.Require().Equal(channelDeposit, channel.Balance)
	s.Require().True(channel.Spent.IsZero())

	// Simulate multiple service calls with incremental payments
	// In real scenario, AI agent would sign payment authorizations
	// For this test, we skip signature verification

	// Note: ClaimChannel requires signature verification which we can't properly test here
	// In production, the AI agent would sign off-chain and user would submit claims

	// Close channel after use period
	s.ctx = s.ctx.WithBlockHeight(s.ctx.BlockHeight() + expiresInBlocks + 1)

	returnedBalance, err := s.settlementKeeper.CloseChannel(s.ctx, channelID, s.user)
	s.Require().NoError(err, "Channel closing should succeed")

	// Verify remaining balance returned
	s.Require().Equal(channel.Balance.Amount, returnedBalance.Amount)

	// Verify channel is closed
	channel, found = s.settlementKeeper.GetChannel(s.ctx, channelID)
	s.Require().True(found)
	s.Require().False(channel.IsOpen)
}

// TestAIAgentComplianceRecording tests that transactions are recorded for compliance
func (s *AIAgentTransactionTestSuite) TestAIAgentComplianceRecording() {
	// Scenario: Verify compliance system tracks AI agent transactions

	servicePrice := sdk.NewCoin("ssusd", sdkmath.NewInt(800000000))

	// Get initial usage
	userProfile, _ := s.complianceKeeper.GetProfile(s.ctx, s.user)
	initialUsed := userProfile.DailyUsed.Amount

	// Create and complete transaction
	settlementID, err := s.settlementKeeper.CreateEscrow(
		s.ctx,
		s.user.String(),
		s.aiAgent.String(),
		servicePrice,
		"COMPLIANCE-TEST",
		`{"test": "compliance"}`,
		3600,
	)
	s.Require().NoError(err)

	err = s.settlementKeeper.ReleaseEscrow(s.ctx, settlementID, s.user)
	s.Require().NoError(err)

	// Record transaction in compliance
	err = s.complianceKeeper.RecordTransaction(s.ctx, s.user, servicePrice)
	s.Require().NoError(err)

	// Verify usage was recorded
	userProfile, _ = s.complianceKeeper.GetProfile(s.ctx, s.user)
	expectedUsed := initialUsed.Add(servicePrice.Amount)
	s.Require().Equal(expectedUsed, userProfile.DailyUsed.Amount,
		"Transaction should be recorded in compliance")
}
