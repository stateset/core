package keeper_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/settlement/keeper"
	"github.com/stateset/core/x/settlement/types"
)

// TestMsgServer_InstantTransfer tests the InstantTransfer message handler
func TestMsgServer_InstantTransfer(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgInstantTransfer{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: "TEST001",
		Metadata:  "test metadata",
	}

	resp, err := msgServer.InstantTransfer(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(1), resp.SettlementId)
}

func TestMsgServer_InstantTransfer_SelfTransfer(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgInstantTransfer{
		Sender:    sender.String(),
		Recipient: sender.String(),
		Amount:    amount,
		Reference: "TEST001",
	}

	_, err := msgServer.InstantTransfer(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidRecipient)
}

func TestMsgServer_InstantTransfer_InvalidSender(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := &types.MsgInstantTransfer{
		Sender:    "invalid",
		Recipient: newSettlementAddress().String(),
		Amount:    sdk.NewCoin("ssusd", sdkmath.NewInt(1000000)),
		Reference: "TEST001",
	}

	_, err := msgServer.InstantTransfer(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgServer_InstantTransfer_InvalidRecipient(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgInstantTransfer{
		Sender:    sender.String(),
		Recipient: "invalid",
		Amount:    sdk.NewCoin("ssusd", sdkmath.NewInt(1000000)),
		Reference: "TEST001",
	}

	_, err := msgServer.InstantTransfer(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgServer_InstantTransfer_ZeroAmount(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgInstantTransfer{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    sdk.NewCoin("ssusd", sdkmath.ZeroInt()),
		Reference: "TEST001",
	}

	_, err := msgServer.InstantTransfer(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_CreateEscrow tests the CreateEscrow message handler
func TestMsgServer_CreateEscrow(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgCreateEscrow{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: "ESCROW001",
		Metadata:  "test escrow",
		ExpiresIn: time.Hour * 24,
	}

	resp, err := msgServer.CreateEscrow(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(1), resp.SettlementId)
	require.False(t, resp.ExpiresAt.IsZero())
}

func TestMsgServer_CreateEscrow_SelfTransfer(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgCreateEscrow{
		Sender:    sender.String(),
		Recipient: sender.String(),
		Amount:    amount,
		Reference: "ESCROW001",
		ExpiresIn: time.Hour * 24,
	}

	_, err := msgServer.CreateEscrow(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidRecipient)
}

func TestMsgServer_CreateEscrow_ShortExpiration(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgCreateEscrow{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: "ESCROW001",
		ExpiresIn: time.Second * 60, // 1 minute
	}

	resp, err := msgServer.CreateEscrow(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)
}

// TestMsgServer_ReleaseEscrow tests the ReleaseEscrow message handler
func TestMsgServer_ReleaseEscrow(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	createMsg := &types.MsgCreateEscrow{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: "ESCROW001",
		ExpiresIn: time.Hour * 24,
	}
	createResp, err := msgServer.CreateEscrow(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(t, err)

	// Release escrow
	releaseMsg := &types.MsgReleaseEscrow{
		SettlementId: createResp.SettlementId,
		Sender:       sender.String(),
	}

	releaseResp, err := msgServer.ReleaseEscrow(sdk.WrapSDKContext(ctx), releaseMsg)
	require.NoError(t, err)
	require.NotNil(t, releaseResp)
}

func TestMsgServer_ReleaseEscrow_InvalidSettlementId(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()

	msg := &types.MsgReleaseEscrow{
		SettlementId: 9999, // Non-existent
		Sender:       sender.String(),
	}

	_, err := msgServer.ReleaseEscrow(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgServer_ReleaseEscrow_InvalidAddress(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := &types.MsgReleaseEscrow{
		SettlementId: 1,
		Sender:       "invalid",
	}

	_, err := msgServer.ReleaseEscrow(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_RefundEscrow tests the RefundEscrow message handler
func TestMsgServer_RefundEscrow(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	amount := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Create escrow
	createMsg := &types.MsgCreateEscrow{
		Sender:    sender.String(),
		Recipient: recipient.String(),
		Amount:    amount,
		Reference: "ESCROW001",
		ExpiresIn: time.Hour * 24,
	}
	createResp, err := msgServer.CreateEscrow(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(t, err)

	// Refund escrow
	refundMsg := &types.MsgRefundEscrow{
		SettlementId: createResp.SettlementId,
		Recipient:    recipient.String(),
		Reason:       "item not received",
	}

	refundResp, err := msgServer.RefundEscrow(sdk.WrapSDKContext(ctx), refundMsg)
	require.NoError(t, err)
	require.NotNil(t, refundResp)
}

func TestMsgServer_RefundEscrow_InvalidRecipientAddress(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := &types.MsgRefundEscrow{
		SettlementId: 1,
		Recipient:    "invalid",
		Reason:       "test",
	}

	_, err := msgServer.RefundEscrow(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_CreateBatch tests the CreateBatch message handler
func TestMsgServer_CreateBatch(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	authority := k.GetAuthority()
	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()
	sender2 := newSettlementAddress()

	senders := []string{sender1.String(), sender2.String()}
	amounts := []sdk.Coin{
		sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		sdk.NewCoin("ssusd", sdkmath.NewInt(200000)),
	}
	references := []string{"REF1", "REF2"}

	for _, sender := range senders {
		bankKeeper.SetBalance(sender, sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))
	}

	msg := &types.MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant.String(),
		Senders:    senders,
		Amounts:    amounts,
		References: references,
	}

	resp, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(1), resp.BatchId)
	require.Len(t, resp.SettlementIds, 2)
}

func TestMsgServer_CreateBatch_EmptySenders(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	authority := k.GetAuthority()
	merchant := newSettlementAddress()

	msg := &types.MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant.String(),
		Senders:    []string{},
		Amounts:    []sdk.Coin{},
		References: []string{},
	}

	_, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgServer_CreateBatch_MismatchedArrays(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	authority := k.GetAuthority()
	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()

	msg := &types.MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant.String(),
		Senders:    []string{sender1.String()},
		Amounts:    []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000)), sdk.NewCoin("ssusd", sdkmath.NewInt(200000))},
		References: []string{"REF1"},
	}

	_, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

func TestMsgServer_CreateBatch_Unauthorized(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()
	sender := newSettlementAddress()

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))

	msg := &types.MsgCreateBatch{
		Authority:  newSettlementAddress().String(),
		Merchant:   merchant.String(),
		Senders:    []string{sender.String()},
		Amounts:    []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000))},
		References: []string{"REF1"},
	}

	_, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

// TestMsgServer_SettleBatch tests the SettleBatch message handler
func TestMsgServer_SettleBatch(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	authority := k.GetAuthority()
	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()

	senders := []string{sender1.String()}
	amounts := []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000))}
	references := []string{"REF1"}

	bankKeeper.SetBalance(sender1.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))

	// Create batch
	createMsg := &types.MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant.String(),
		Senders:    senders,
		Amounts:    amounts,
		References: references,
	}
	createResp, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(t, err)

	// Settle batch
	settleMsg := &types.MsgSettleBatch{
		BatchId:   createResp.BatchId,
		Authority: authority,
	}

	settleResp, err := msgServer.SettleBatch(sdk.WrapSDKContext(ctx), settleMsg)
	require.NoError(t, err)
	require.NotNil(t, settleResp)
	require.True(t, settleResp.TotalAmount.IsPositive())
}

func TestMsgServer_SettleBatch_Unauthorized(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	authority := k.GetAuthority()
	merchant := newSettlementAddress()
	sender1 := newSettlementAddress()

	senders := []string{sender1.String()}
	amounts := []sdk.Coin{sdk.NewCoin("ssusd", sdkmath.NewInt(100000))}
	references := []string{"REF1"}

	bankKeeper.SetBalance(sender1.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))))

	// Create batch
	createMsg := &types.MsgCreateBatch{
		Authority:  authority,
		Merchant:   merchant.String(),
		Senders:    senders,
		Amounts:    amounts,
		References: references,
	}
	createResp, err := msgServer.CreateBatch(sdk.WrapSDKContext(ctx), createMsg)
	require.NoError(t, err)

	// Try to settle with wrong authority
	settleMsg := &types.MsgSettleBatch{
		BatchId:   createResp.BatchId,
		Authority: newSettlementAddress().String(),
	}

	_, err = msgServer.SettleBatch(sdk.WrapSDKContext(ctx), settleMsg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

// TestMsgServer_OpenChannel tests the OpenChannel message handler
func TestMsgServer_OpenChannel(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       recipient.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 1000,
	}

	resp, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, uint64(1), resp.ChannelId)
	require.Equal(t, ctx.BlockHeight()+1000, resp.ExpiresAtHeight)
}

func TestMsgServer_OpenChannel_SelfTransfer(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	msg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       sender.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 1000,
	}

	_, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrInvalidRecipient)
}

func TestMsgServer_OpenChannel_InsufficientBalance(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(500000))))

	msg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       recipient.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 1000,
	}

	_, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_CloseChannel tests the CloseChannel message handler
func TestMsgServer_CloseChannel(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel
	openMsg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       recipient.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 10,
	}
	openResp, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), openMsg)
	require.NoError(t, err)

	// Move past expiration
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 11)

	// Close channel
	closeMsg := &types.MsgCloseChannel{
		ChannelId: openResp.ChannelId,
		Closer:    sender.String(),
	}

	closeResp, err := msgServer.CloseChannel(sdk.WrapSDKContext(ctx), closeMsg)
	require.NoError(t, err)
	require.NotNil(t, closeResp)
	require.Equal(t, deposit, closeResp.FinalBalance)
}

func TestMsgServer_CloseChannel_BeforeExpiration(t *testing.T) {
	k, ctx, bankKeeper, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	sender := newSettlementAddress()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))

	// Open channel
	openMsg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       recipient.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 1000,
	}
	openResp, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), openMsg)
	require.NoError(t, err)

	// Try to close before expiration
	closeMsg := &types.MsgCloseChannel{
		ChannelId: openResp.ChannelId,
		Closer:    sender.String(),
	}

	_, err = msgServer.CloseChannel(sdk.WrapSDKContext(ctx), closeMsg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrChannelNotExpired)
}

func TestMsgServer_CloseChannel_InvalidAddress(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := &types.MsgCloseChannel{
		ChannelId: 1,
		Closer:    "invalid",
	}

	_, err := msgServer.CloseChannel(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_ClaimChannel tests the ClaimChannel message handler
func TestMsgServer_ClaimChannel(t *testing.T) {
	k, ctx, bankKeeper, _, accountKeeper := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	senderKey, sender := newSettlementKeyPair()
	recipient := newSettlementAddress()
	deposit := sdk.NewCoin("ssusd", sdkmath.NewInt(1000000))

	bankKeeper.SetBalance(sender.String(), sdk.NewCoins(sdk.NewCoin("ssusd", sdkmath.NewInt(2000000))))
	accountKeeper.SetPubKey(sender, senderKey.PubKey())

	// Open channel
	openMsg := &types.MsgOpenChannel{
		Sender:          sender.String(),
		Recipient:       recipient.String(),
		Deposit:         deposit,
		ExpiresInBlocks: 1000,
	}
	openResp, err := msgServer.OpenChannel(sdk.WrapSDKContext(ctx), openMsg)
	require.NoError(t, err)

	// Claim from channel
	claimAmount := sdk.NewCoin("ssusd", sdkmath.NewInt(300000))
	sig := signChannelClaim(senderKey, openResp.ChannelId, recipient, claimAmount, 1)

	claimMsg := &types.MsgClaimChannel{
		ChannelId: openResp.ChannelId,
		Recipient: recipient.String(),
		Amount:    claimAmount,
		Nonce:     1,
		Signature: sig,
	}

	claimResp, err := msgServer.ClaimChannel(sdk.WrapSDKContext(ctx), claimMsg)
	require.NoError(t, err)
	require.NotNil(t, claimResp)
	require.Equal(t, claimAmount, claimResp.AmountClaimed)
}

func TestMsgServer_ClaimChannel_InvalidRecipient(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	msg := &types.MsgClaimChannel{
		ChannelId: 1,
		Recipient: "invalid",
		Amount:    sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		Nonce:     1,
		Signature: "abc",
	}

	_, err := msgServer.ClaimChannel(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
}

// TestMsgServer_RegisterMerchant tests the RegisterMerchant message handler
func TestMsgServer_RegisterMerchant(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	msg := &types.MsgRegisterMerchant{
		Authority:      merchant.String(),
		Merchant:       merchant.String(),
		Name:           "Test Merchant",
		FeeRateBps:     25,
		MinSettlement:  sdk.NewCoin("ssusd", sdkmath.NewInt(1000)),
		MaxSettlement:  sdk.NewCoin("ssusd", sdkmath.NewInt(10000000)),
		BatchEnabled:   true,
		BatchThreshold: sdk.NewCoin("ssusd", sdkmath.NewInt(100000)),
		WebhookUrl:     "https://example.com/webhook",
	}

	resp, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), msg)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Verify merchant was registered
	storedMerchant, found := k.GetMerchant(ctx, merchant.String())
	require.True(t, found)
	require.Equal(t, "Test Merchant", storedMerchant.Name)
	require.Equal(t, uint32(25), storedMerchant.FeeRateBps)
}

func TestMsgServer_RegisterMerchant_EmptyName(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	msg := &types.MsgRegisterMerchant{
		Authority:  merchant.String(),
		Merchant:   merchant.String(),
		Name:       "",
		FeeRateBps: 25,
	}

	resp, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), msg)
	// Empty name should be allowed but stored
	require.NoError(t, err)
	require.NotNil(t, resp)
}

func TestMsgServer_RegisterMerchant_Unauthorized(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	msg := &types.MsgRegisterMerchant{
		Authority:  newSettlementAddress().String(),
		Merchant:   merchant.String(),
		Name:       "Test Merchant",
		FeeRateBps: 25,
	}

	_, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

// TestMsgServer_UpdateMerchant tests the UpdateMerchant message handler
func TestMsgServer_UpdateMerchant(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	// Register merchant first
	registerMsg := &types.MsgRegisterMerchant{
		Authority:    merchant.String(),
		Merchant:     merchant.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25,
		BatchEnabled: true,
	}
	_, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), registerMsg)
	require.NoError(t, err)

	// Update merchant
	batchEnabled := false
	isActive := true
	updateMsg := &types.MsgUpdateMerchant{
		Authority:    merchant.String(),
		Merchant:     merchant.String(),
		Name:         "Updated Merchant",
		FeeRateBps:   30,
		BatchEnabled: &batchEnabled,
		IsActive:     &isActive,
	}

	resp, err := msgServer.UpdateMerchant(sdk.WrapSDKContext(ctx), updateMsg)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Verify updates
	storedMerchant, found := k.GetMerchant(ctx, merchant.String())
	require.True(t, found)
	require.Equal(t, "Updated Merchant", storedMerchant.Name)
	require.Equal(t, uint32(30), storedMerchant.FeeRateBps)
	require.False(t, storedMerchant.BatchEnabled)
}

func TestMsgServer_UpdateMerchant_Unauthorized(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	registerMsg := &types.MsgRegisterMerchant{
		Authority:    merchant.String(),
		Merchant:     merchant.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25,
		BatchEnabled: true,
	}
	_, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), registerMsg)
	require.NoError(t, err)

	updateMsg := &types.MsgUpdateMerchant{
		Authority:  newSettlementAddress().String(),
		Merchant:   merchant.String(),
		Name:       "Updated Merchant",
		FeeRateBps: 30,
	}

	_, err = msgServer.UpdateMerchant(sdk.WrapSDKContext(ctx), updateMsg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrUnauthorized)
}

func TestMsgServer_UpdateMerchant_NotFound(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	msg := &types.MsgUpdateMerchant{
		Authority:  merchant.String(),
		Merchant:   merchant.String(),
		Name:       "Updated Merchant",
		FeeRateBps: 30,
	}

	_, err := msgServer.UpdateMerchant(sdk.WrapSDKContext(ctx), msg)
	require.Error(t, err)
	require.ErrorIs(t, err, types.ErrMerchantNotFound)
}

func TestMsgServer_UpdateMerchant_PartialUpdate(t *testing.T) {
	k, ctx, _, _, _ := setupSettlementKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	merchant := newSettlementAddress()

	// Register merchant
	registerMsg := &types.MsgRegisterMerchant{
		Authority:    merchant.String(),
		Merchant:     merchant.String(),
		Name:         "Test Merchant",
		FeeRateBps:   25,
		BatchEnabled: true,
	}
	_, err := msgServer.RegisterMerchant(sdk.WrapSDKContext(ctx), registerMsg)
	require.NoError(t, err)

	// Update only name
	updateMsg := &types.MsgUpdateMerchant{
		Authority: merchant.String(),
		Merchant:  merchant.String(),
		Name:      "Updated Name Only",
	}

	resp, err := msgServer.UpdateMerchant(sdk.WrapSDKContext(ctx), updateMsg)
	require.NoError(t, err)
	require.NotNil(t, resp)

	// Verify only name changed
	storedMerchant, found := k.GetMerchant(ctx, merchant.String())
	require.True(t, found)
	require.Equal(t, "Updated Name Only", storedMerchant.Name)
	require.Equal(t, uint32(25), storedMerchant.FeeRateBps) // Unchanged
	require.True(t, storedMerchant.BatchEnabled)            // Unchanged
}
