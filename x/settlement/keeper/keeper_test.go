package keeper_test

import (
	"context"
	"sync"
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/stateset/core/x/settlement/keeper"
	"github.com/stateset/core/x/settlement/types"
)

var settlementConfigOnce sync.Once

func setupSettlementConfig() {
	settlementConfigOnce.Do(func() {
		cfg := sdk.GetConfig()
		cfg.SetBech32PrefixForAccount("stateset", "statesetpub")
		cfg.SetBech32PrefixForValidator("statesetvaloper", "statesetvaloperpub")
		cfg.SetBech32PrefixForConsensusNode("statesetvalcons", "statesetvalconspub")
		cfg.Seal()
	})
}

func newSettlementAddress() sdk.AccAddress {
	key := secp256k1.GenPrivKey()
	return sdk.AccAddress(key.PubKey().Address())
}

// Mock bank keeper
type mockBankKeeper struct {
	balances map[string]sdk.Coins
	moduleBalances map[string]sdk.Coins
}

func newMockBankKeeper() *mockBankKeeper {
	return &mockBankKeeper{
		balances:       make(map[string]sdk.Coins),
		moduleBalances: make(map[string]sdk.Coins),
	}
}

func (m *mockBankKeeper) SetBalance(addr string, coins sdk.Coins) {
	m.balances[addr] = coins
}

func (m *mockBankKeeper) GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	coins := m.balances[addr.String()]
	return sdk.NewCoin(denom, coins.AmountOf(denom))
}

func (m *mockBankKeeper) SendCoins(ctx context.Context, from, to sdk.AccAddress, amt sdk.Coins) error {
	fromCoins := m.balances[from.String()]
	toCoins := m.balances[to.String()]

	newFrom, _ := fromCoins.SafeSub(amt...)
	m.balances[from.String()] = newFrom
	m.balances[to.String()] = toCoins.Add(amt...)
	return nil
}

func (m *mockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	fromCoins := m.balances[senderAddr.String()]
	moduleCoins := m.moduleBalances[recipientModule]

	newFrom, _ := fromCoins.SafeSub(amt...)
	m.balances[senderAddr.String()] = newFrom
	m.moduleBalances[recipientModule] = moduleCoins.Add(amt...)
	return nil
}

func (m *mockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[senderModule]
	toCoins := m.balances[recipientAddr.String()]

	newModule, _ := moduleCoins.SafeSub(amt...)
	m.moduleBalances[senderModule] = newModule
	m.balances[recipientAddr.String()] = toCoins.Add(amt...)
	return nil
}

func (m *mockBankKeeper) MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[moduleName]
	m.moduleBalances[moduleName] = moduleCoins.Add(amt...)
	return nil
}

func (m *mockBankKeeper) BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error {
	moduleCoins := m.moduleBalances[moduleName]
	newModule, _ := moduleCoins.SafeSub(amt...)
	m.moduleBalances[moduleName] = newModule
	return nil
}

func (m *mockBankKeeper) SendCoinsFromModuleToModule(ctx context.Context, senderModule, recipientModule string, amt sdk.Coins) error {
	senderCoins := m.moduleBalances[senderModule]
	recipientCoins := m.moduleBalances[recipientModule]

	newSender, _ := senderCoins.SafeSub(amt...)
	m.moduleBalances[senderModule] = newSender
	m.moduleBalances[recipientModule] = recipientCoins.Add(amt...)
	return nil
}

// Mock compliance keeper
type mockComplianceKeeper struct {
	sanctionedAddresses map[string]bool
}

func newMockComplianceKeeper() *mockComplianceKeeper {
	return &mockComplianceKeeper{
		sanctionedAddresses: make(map[string]bool),
	}
}

func (m *mockComplianceKeeper) AssertCompliant(ctx context.Context, addr sdk.AccAddress) error {
	if m.sanctionedAddresses[addr.String()] {
		return types.ErrComplianceCheckFailed
	}
	return nil
}

func (m *mockComplianceKeeper) SetSanctioned(addr string, sanctioned bool) {
	m.sanctionedAddresses[addr] = sanctioned
}

func setupSettlementKeeper(t *testing.T) (keeper.Keeper, sdk.Context, *mockBankKeeper, *mockComplianceKeeper) {
	t.Helper()
	setupSettlementConfig()

	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	bankKeeper := newMockBankKeeper()
	complianceKeeper := newMockComplianceKeeper()

	authority := newSettlementAddress()

	k := keeper.NewKeeper(
		cdc,
		storeKey,
		bankKeeper,
		complianceKeeper,
		authority.String(),
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{Height: 1, ChainID: "stateset-test", Time: time.Now()}, false, log.NewNopLogger())

	// Initialize genesis
	k.InitGenesis(ctx, types.DefaultGenesis())

	return k, ctx, bankKeeper, complianceKeeper
}

func TestInstantTransfer(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Execute instant transfer
	settlementId, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "test transfer")
	require.NoError(t, err)
	require.Equal(t, uint64(1), settlementId)

	// Verify settlement record
	settlement, found := k.GetSettlement(ctx, settlementId)
	require.True(t, found)
	require.Equal(t, types.SettlementStatusCompleted, settlement.Status)
	require.Equal(t, types.SettlementTypeInstant, settlement.Type)
	require.Equal(t, sender.String(), settlement.Sender)
	require.Equal(t, recipient.String(), settlement.Recipient)
	require.Equal(t, amount, settlement.Amount)
	require.Equal(t, "REF001", settlement.Reference)
}

func TestInstantTransfer_InsufficientBalance(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender with less than needed
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(500000))))

	// Execute instant transfer should fail
	_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "test transfer")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInsufficientFunds)
}

func TestInstantTransfer_ComplianceBlocked(t *testing.T) {
	k, ctx, bankKeeper, complianceKeeper := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Sanction sender
	complianceKeeper.SetSanctioned(sender.String(), true)

	// Execute instant transfer should fail
	_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "test transfer")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrComplianceCheckFailed)
}

func TestInstantTransfer_AmountTooSmall(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(100)) // Too small

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Execute instant transfer should fail
	_, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "test transfer")
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrSettlementTooSmall)
}

func TestCreateEscrow(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	settlementId, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESCROW001", "test escrow", 86400)
	require.NoError(t, err)
	require.Equal(t, uint64(1), settlementId)

	// Verify settlement record
	settlement, found := k.GetSettlement(ctx, settlementId)
	require.True(t, found)
	require.Equal(t, types.SettlementStatusPending, settlement.Status)
	require.Equal(t, types.SettlementTypeEscrow, settlement.Type)
	require.False(t, settlement.ExpiresAt.IsZero())
}

func TestReleaseEscrow(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	settlementId, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESCROW001", "test escrow", 86400)
	require.NoError(t, err)

	// Release escrow
	senderAddr := sender
	err = k.ReleaseEscrow(ctx, settlementId, senderAddr)
	require.NoError(t, err)

	// Verify settlement status
	settlement, found := k.GetSettlement(ctx, settlementId)
	require.True(t, found)
	require.Equal(t, types.SettlementStatusCompleted, settlement.Status)
}

func TestReleaseEscrow_WrongSender(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	wrongSender := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	settlementId, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESCROW001", "test escrow", 86400)
	require.NoError(t, err)

	// Try to release with wrong sender
	wrongSenderAddr := wrongSender
	err = k.ReleaseEscrow(ctx, settlementId, wrongSenderAddr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

func TestRefundEscrow(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	settlementId, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESCROW001", "test escrow", 86400)
	require.NoError(t, err)

	// Refund escrow (by recipient)
	recipientAddr := recipient
	err = k.RefundEscrow(ctx, settlementId, recipientAddr, "item not received")
	require.NoError(t, err)

	// Verify settlement status
	settlement, found := k.GetSettlement(ctx, settlementId)
	require.True(t, found)
	require.Equal(t, types.SettlementStatusRefunded, settlement.Status)
}

func TestOpenChannel(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel
	channelId, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	require.NoError(t, err)
	require.Equal(t, uint64(1), channelId)

	// Verify channel
	channel, found := k.GetChannel(ctx, channelId)
	require.True(t, found)
	require.True(t, channel.IsOpen)
	require.Equal(t, deposit, channel.Deposit)
	require.Equal(t, deposit, channel.Balance)
	require.Equal(t, uint64(0), channel.Nonce)
}

func TestClaimChannel(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel
	channelId, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	require.NoError(t, err)

	// Claim from channel
	recipientAddr := recipient
	claimAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(300000))
	err = k.ClaimChannel(ctx, channelId, recipientAddr, claimAmount, 1)
	require.NoError(t, err)

	// Verify channel state
	channel, _ := k.GetChannel(ctx, channelId)
	require.Equal(t, uint64(1), channel.Nonce)
	require.Equal(t, claimAmount, channel.Spent)
	expectedBalance := sdk.NewCoin("ssusd", sdkmath.NewInt(700000))
	require.Equal(t, expectedBalance, channel.Balance)
}

func TestClaimChannel_InvalidNonce(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel
	channelId, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 1000)
	require.NoError(t, err)

	// First claim
	recipientAddr := recipient
	claimAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(100000))
	err = k.ClaimChannel(ctx, channelId, recipientAddr, claimAmount, 1)
	require.NoError(t, err)

	// Try to claim with same or lower nonce (replay attack)
	err = k.ClaimChannel(ctx, channelId, recipientAddr, claimAmount, 1)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidNonce)

	err = k.ClaimChannel(ctx, channelId, recipientAddr, claimAmount, 0)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidNonce)
}

func TestCloseChannel(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel with short expiration
	channelId, err := k.OpenChannel(ctx, sender.String(), recipient.String(), deposit, 10)
	require.NoError(t, err)

	// Try to close before expiration (should fail)
	senderAddr := sender
	_, err = k.CloseChannel(ctx, channelId, senderAddr)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrChannelNotExpired)

	// Move block height past expiration
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 11)

	// Now close should work
	balance, err := k.CloseChannel(ctx, channelId, senderAddr)
	require.NoError(t, err)
	require.Equal(t, deposit, balance)

	// Verify channel is closed
	channel, _ := k.GetChannel(ctx, channelId)
	require.False(t, channel.IsOpen)
}

func TestCreateBatch(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()
	sender2 := newSettlementAddress()
	sender3 := newSettlementAddress()
	senders := []string{sender1.String(), sender2.String(), sender3.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(150000)),
	}
	references := []string{"REF1", "REF2", "REF3"}

	// Fund senders
	for _, sender := range senders {
		bankKeeper.SetBalance(sender, sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))
	}

	// Create batch
	batchId, settlementIds, err := k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
	require.NoError(t, err)
	require.Equal(t, uint64(1), batchId)
	require.Len(t, settlementIds, 3)

	// Verify batch
	batch, found := k.GetBatch(ctx, batchId)
	require.True(t, found)
	require.Equal(t, types.SettlementStatusPending, batch.Status)
	require.Equal(t, uint64(3), batch.Count)
	require.Equal(t, merchant.String(), batch.Merchant)
}

func TestSettleBatch(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()
	sender2 := newSettlementAddress()
	senders := []string{sender1.String(), sender2.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000)),
	}
	references := []string{"REF1", "REF2"}

	// Fund senders
	for _, sender := range senders {
		bankKeeper.SetBalance(sender, sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))
	}

	// Create batch
	batchId, _, err := k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
	require.NoError(t, err)

	// Settle batch (as authority)
	authority := k.GetAuthority()
	err = k.SettleBatch(ctx, batchId, authority)
	require.NoError(t, err)

	// Verify batch is completed
	batch, _ := k.GetBatch(ctx, batchId)
	require.Equal(t, types.SettlementStatusCompleted, batch.Status)
}

func TestSettleBatch_Unauthorized(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()
	senders := []string{sender1.String()}
	amounts := []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000))}
	references := []string{"REF1"}

	// Fund sender
	bankKeeper.SetBalance(sender1.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))

	// Create batch
	batchId, _, err := k.CreateBatch(ctx, merchant.String(), senders, amounts, references)
	require.NoError(t, err)

	// Try to settle as non-authority
	randomAddr := newSettlementAddress()
	err = k.SettleBatch(ctx, batchId, randomAddr.String())
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

func TestRegisterMerchant(t *testing.T) {
	k, ctx, _, _ := setupSettlementKeeper(t)

	merchantAddr := newSettlementAddress()
	config := types.MerchantConfig{
		Address:      merchantAddr.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25, // 0.25%
		BatchEnabled: true,
	}

	err := k.RegisterMerchant(ctx, config)
	require.NoError(t, err)

	// Verify merchant
	merchant, found := k.GetMerchant(ctx, merchantAddr.String())
	require.True(t, found)
	require.Equal(t, "Test Merchant", merchant.Name)
	require.Equal(t, uint32(25), merchant.FeeRateBps)
	require.True(t, merchant.IsActive)
	require.True(t, merchant.BatchEnabled)
}

func TestUpdateMerchant(t *testing.T) {
	k, ctx, _, _ := setupSettlementKeeper(t)

	merchantAddr := newSettlementAddress()
	// Register merchant
	config := types.MerchantConfig{
		Address:      merchantAddr.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25,
		BatchEnabled: true,
	}
	k.RegisterMerchant(ctx, config)

	// Update merchant
	updates := map[string]interface{}{
		"name":          "Updated Merchant",
		"fee_rate_bps":  uint32(30),
		"batch_enabled": false,
	}
	err := k.UpdateMerchant(ctx, merchantAddr.String(), updates)
	require.NoError(t, err)

	// Verify updates
	merchant, _ := k.GetMerchant(ctx, merchantAddr.String())
	require.Equal(t, "Updated Merchant", merchant.Name)
	require.Equal(t, uint32(30), merchant.FeeRateBps)
	require.False(t, merchant.BatchEnabled)
}

func TestProcessExpiredEscrows(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow with short expiration
	settlementId, err := k.CreateEscrow(ctx, sender.String(), recipient.String(), amount, "ESCROW001", "test escrow", 60) // 60 seconds
	require.NoError(t, err)

	// Verify it's pending
	settlement, _ := k.GetSettlement(ctx, settlementId)
	require.Equal(t, types.SettlementStatusPending, settlement.Status)

	// Move time forward past expiration
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(61 * time.Second))

	// Process expired escrows
	k.ProcessExpiredEscrows(ctx)

	// Verify it's cancelled and refunded
	settlement, _ = k.GetSettlement(ctx, settlementId)
	require.Equal(t, types.SettlementStatusCancelled, settlement.Status)
	require.Contains(t, settlement.Metadata, "expired")
}

func TestFeeCalculation(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000)) // 1 SSUSD

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Execute transfer
	settlementId, err := k.InstantTransfer(ctx, sender.String(), recipient.String(), amount, "REF001", "")
	require.NoError(t, err)

	// Verify fee was calculated (default 0.5%)
	settlement, _ := k.GetSettlement(ctx, settlementId)
	expectedFee := sdk.NewCoin("ssusd", sdkmath.NewInt(5000)) // 0.5% of 1M
	expectedNet := sdk.NewCoin("ssusd", sdkmath.NewInt(995000))

	require.Equal(t, expectedFee, settlement.Fee)
	require.Equal(t, expectedNet, settlement.NetAmount)
}

func TestMerchantCustomFeeRate(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	merchant := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	// Register merchant with custom fee rate
	config := types.MerchantConfig{
		Address:    merchant.String(),
		Name:       "Test Merchant",
		FeeRateBps: 25, // 0.25%
	}
	k.RegisterMerchant(ctx, config)

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Execute transfer to merchant
	settlementId, err := k.InstantTransfer(ctx, sender.String(), merchant.String(), amount, "REF001", "")
	require.NoError(t, err)

	// Verify custom fee was applied
	settlement, _ := k.GetSettlement(ctx, settlementId)
	expectedFee := sdk.NewCoin("ssusd", sdkmath.NewInt(2500)) // 0.25% of 1M

	require.Equal(t, expectedFee, settlement.Fee)
}

func TestGenesisExportImport(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	merchantAddr := newSettlementAddress()

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create some state
	k.InstantTransfer(ctx, sender.String(), recipient.String(), sdk.NewCoin("ssusd", sdkmath.NewInt(100000)), "REF1", "")
	k.CreateEscrow(ctx, sender.String(), recipient.String(), sdk.NewCoin("ssusd", sdkmath.NewInt(200000)), "ESC1", "", 86400)
	k.OpenChannel(ctx, sender.String(), recipient.String(), sdk.NewCoin("ssusd", sdkmath.NewInt(300000)), 1000)

	config := types.MerchantConfig{
		Address: merchantAddr.String(),
		Name:    "Test Merchant",
	}
	k.RegisterMerchant(ctx, config)

	// Export genesis
	genesis := k.ExportGenesis(ctx)
	require.NotNil(t, genesis)
	require.Len(t, genesis.Settlements, 2)
	require.Len(t, genesis.Channels, 1)
	require.Len(t, genesis.Merchants, 1)

	// Create new keeper and import
	k2, ctx2, _, _ := setupSettlementKeeper(t)
	k2.InitGenesis(ctx2, genesis)

	// Verify state was imported
	settlement1, found := k2.GetSettlement(ctx2, 1)
	require.True(t, found)
	require.Equal(t, "REF1", settlement1.Reference)

	channel, found := k2.GetChannel(ctx2, 1)
	require.True(t, found)
	require.True(t, channel.IsOpen)

	merchant, found := k2.GetMerchant(ctx2, merchantAddr.String())
	require.True(t, found)
	require.Equal(t, "Test Merchant", merchant.Name)
}

func TestIterators(t *testing.T) {
	k, ctx, bankKeeper, _ := setupSettlementKeeper(t)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()

	// Fund sender
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(10000000))))

	// Create multiple settlements
	for i := 0; i < 5; i++ {
		k.InstantTransfer(ctx, sender.String(), recipient.String(), sdk.NewCoin("ssusd", sdkmath.NewInt(100000)), "", "")
	}

	// Iterate and count
	count := 0
	k.IterateSettlements(ctx, func(s types.Settlement) bool {
		count++
		return false
	})
	require.Equal(t, 5, count)

	// Test early termination
	count = 0
	k.IterateSettlements(ctx, func(s types.Settlement) bool {
		count++
		return count >= 3 // Stop after 3
	})
	require.Equal(t, 3, count)
}
