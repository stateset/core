package security_test

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/stateset/core/app/apptesting"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
)

type SettlementSecurityTestSuite struct {
	apptesting.KeeperTestHelper

	denom string

	sender    sdk.AccAddress
	recipient sdk.AccAddress
	other     sdk.AccAddress

	channelSenderPriv *secp256k1.PrivKey
	channelSender     sdk.AccAddress
}

func TestSettlementSecurityTestSuite(t *testing.T) {
	suite.Run(t, new(SettlementSecurityTestSuite))
}

func (suite *SettlementSecurityTestSuite) SetupTest() {
	suite.KeeperTestHelper.Setup()

	suite.denom = settlementtypes.StablecoinDenom
	suite.sender = suite.TestAccs[0]
	suite.recipient = suite.TestAccs[1]
	suite.other = suite.TestAccs[2]

	suite.setCompliant(suite.sender)
	suite.setCompliant(suite.recipient)
	suite.setCompliant(suite.other)

	suite.FundAcc(suite.sender, sdk.NewCoins(sdk.NewInt64Coin(suite.denom, 1_000_000)))

	// Dedicated keypair for channel signature tests.
	priv := secp256k1.GenPrivKey()
	suite.channelSenderPriv = priv
	suite.channelSender = sdk.AccAddress(priv.PubKey().Address())
	suite.setCompliant(suite.channelSender)
	suite.FundAcc(suite.channelSender, sdk.NewCoins(sdk.NewInt64Coin(suite.denom, 1_000_000)))

	// Ensure the sender account has a pubkey set so ClaimChannel can verify signatures.
	acc := suite.App.AccountKeeper.GetAccount(suite.Ctx, suite.channelSender)
	suite.Require().NotNil(acc)
	acc.SetPubKey(priv.PubKey())
	suite.App.AccountKeeper.SetAccount(suite.Ctx, acc)
}

func (suite *SettlementSecurityTestSuite) compliantProfile(addr sdk.AccAddress) compliancetypes.Profile {
	now := suite.Ctx.BlockTime()
	return compliancetypes.Profile{
		Address:        addr.String(),
		KYCLevel:       compliancetypes.KYCStandard,
		Risk:           compliancetypes.RiskLow,
		Status:         compliancetypes.StatusActive,
		Sanction:       false,
		Jurisdiction:   "US",
		VerifiedAt:     now,
		ExpiresAt:      now.Add(365 * 24 * time.Hour),
		LastLimitReset: now,
	}
}

func (suite *SettlementSecurityTestSuite) setCompliant(addr sdk.AccAddress) {
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), suite.compliantProfile(addr))
}

func (suite *SettlementSecurityTestSuite) TestReleaseEscrow_Unauthorized() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := settlementtypes.NewMsgCreateEscrow(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		"escrow-1",
		"meta",
		0, // use default
	)
	resp, err := msgServer.CreateEscrow(goCtx, create)
	suite.Require().NoError(err)

	settlement, found := suite.App.SettlementKeeper.GetSettlement(suite.Ctx, resp.SettlementId)
	suite.Require().True(found)
	suite.Require().Equal(settlementtypes.SettlementStatusPending, settlement.Status)

	_, err = msgServer.ReleaseEscrow(goCtx, settlementtypes.NewMsgReleaseEscrow(suite.other.String(), resp.SettlementId))
	suite.Require().ErrorIs(err, settlementtypes.ErrUnauthorized)

	_, err = msgServer.ReleaseEscrow(goCtx, settlementtypes.NewMsgReleaseEscrow(suite.sender.String(), resp.SettlementId))
	suite.Require().NoError(err)

	_, err = msgServer.ReleaseEscrow(goCtx, settlementtypes.NewMsgReleaseEscrow(suite.sender.String(), resp.SettlementId))
	suite.Require().ErrorIs(err, settlementtypes.ErrSettlementCompleted)

	moduleAddr := suite.App.AccountKeeper.GetModuleAddress(settlementtypes.ModuleAccountName)
	suite.Require().Equal(settlement.Fee.Amount.Int64(), suite.App.BankKeeper.GetBalance(suite.Ctx, moduleAddr, suite.denom).Amount.Int64())
	suite.Require().Equal(settlement.NetAmount.Amount.Int64(), suite.App.BankKeeper.GetBalance(suite.Ctx, suite.recipient, suite.denom).Amount.Int64())
}

func (suite *SettlementSecurityTestSuite) TestRefundEscrow_Unauthorized() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := settlementtypes.NewMsgCreateEscrow(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		"escrow-2",
		"meta",
		0,
	)
	resp, err := msgServer.CreateEscrow(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.RefundEscrow(goCtx, settlementtypes.NewMsgRefundEscrow(suite.sender.String(), resp.SettlementId, "sender cannot refund"))
	suite.Require().ErrorIs(err, settlementtypes.ErrUnauthorized)

	_, err = msgServer.RefundEscrow(goCtx, settlementtypes.NewMsgRefundEscrow(suite.other.String(), resp.SettlementId, "random cannot refund"))
	suite.Require().ErrorIs(err, settlementtypes.ErrUnauthorized)

	_, err = msgServer.RefundEscrow(goCtx, settlementtypes.NewMsgRefundEscrow(suite.recipient.String(), resp.SettlementId, "refund"))
	suite.Require().NoError(err)

	settlement, found := suite.App.SettlementKeeper.GetSettlement(suite.Ctx, resp.SettlementId)
	suite.Require().True(found)
	suite.Require().Equal(settlementtypes.SettlementStatusRefunded, settlement.Status)

	moduleAddr := suite.App.AccountKeeper.GetModuleAddress(settlementtypes.ModuleAccountName)
	suite.Require().Equal(int64(0), suite.App.BankKeeper.GetBalance(suite.Ctx, moduleAddr, suite.denom).Amount.Int64())
}

func (suite *SettlementSecurityTestSuite) TestInstantTransfer_AmountValidation() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	_, err := msgServer.InstantTransfer(goCtx, settlementtypes.NewMsgInstantTransfer(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 999),
		"too-small",
		"",
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrSettlementTooSmall)

	resp, err := msgServer.InstantTransfer(goCtx, settlementtypes.NewMsgInstantTransfer(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"min-ok",
		"",
	))
	suite.Require().NoError(err)
	suite.Require().Greater(resp.SettlementId, uint64(0))
}

func (suite *SettlementSecurityTestSuite) TestEscrow_ExpirationValidation() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	// 0 uses default
	_, err := msgServer.CreateEscrow(goCtx, settlementtypes.NewMsgCreateEscrow(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		"exp-default",
		"",
		0,
	))
	suite.Require().NoError(err)

	// > max should fail
	max := suite.App.SettlementKeeper.GetParams(suite.Ctx).MaxEscrowExpiration
	_, err = msgServer.CreateEscrow(goCtx, settlementtypes.NewMsgCreateEscrow(
		suite.sender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		"exp-too-long",
		"",
		time.Duration(max+1)*time.Second,
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrInvalidEscrowExpiration)
}

func (suite *SettlementSecurityTestSuite) TestChannel_ClaimNonceAndSignature() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)

	// Open a channel with min expiration.
	params := suite.App.SettlementKeeper.GetParams(suite.Ctx)
	open := settlementtypes.NewMsgOpenChannel(
		suite.channelSender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		params.MinChannelExpiration,
	)
	resp, err := msgServer.OpenChannel(sdk.WrapSDKContext(suite.Ctx), open)
	suite.Require().NoError(err)

	claimAmt := sdk.NewInt64Coin(suite.denom, 1_000)
	nonce := uint64(1)
	sig := suite.signChannelClaim(resp.ChannelId, suite.recipient, claimAmt, nonce)

	_, err = msgServer.ClaimChannel(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgClaimChannel(
		suite.recipient.String(),
		resp.ChannelId,
		claimAmt,
		nonce,
		sig,
	))
	suite.Require().NoError(err)

	// Replaying the same nonce should fail.
	_, err = msgServer.ClaimChannel(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgClaimChannel(
		suite.recipient.String(),
		resp.ChannelId,
		claimAmt,
		nonce,
		sig,
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrInvalidNonce)
}

func (suite *SettlementSecurityTestSuite) TestChannel_CloseAuthorizationAndExpiry() {
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)

	params := suite.App.SettlementKeeper.GetParams(suite.Ctx)
	open := settlementtypes.NewMsgOpenChannel(
		suite.channelSender.String(),
		suite.recipient.String(),
		sdk.NewInt64Coin(suite.denom, 10_000),
		params.MinChannelExpiration,
	)
	resp, err := msgServer.OpenChannel(sdk.WrapSDKContext(suite.Ctx), open)
	suite.Require().NoError(err)

	// Recipient cannot close.
	_, err = msgServer.CloseChannel(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgCloseChannel(
		suite.recipient.String(),
		resp.ChannelId,
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrUnauthorized)

	// Sender cannot close before expiration.
	_, err = msgServer.CloseChannel(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgCloseChannel(
		suite.channelSender.String(),
		resp.ChannelId,
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrChannelNotExpired)

	// Advance height to expiration and close.
	channel, found := suite.App.SettlementKeeper.GetChannel(suite.Ctx, resp.ChannelId)
	suite.Require().True(found)

	suite.Ctx = suite.Ctx.WithBlockHeight(channel.ExpiresAtHeight)
	_, err = msgServer.CloseChannel(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgCloseChannel(
		suite.channelSender.String(),
		resp.ChannelId,
	))
	suite.Require().NoError(err)
}

func (suite *SettlementSecurityTestSuite) signChannelClaim(channelID uint64, recipient sdk.AccAddress, amount sdk.Coin, nonce uint64) string {
	msg := fmt.Sprintf("channel_claim:%d:%s:%s:%d", channelID, recipient.String(), amount.String(), nonce)
	sig, err := suite.channelSenderPriv.Sign([]byte(msg))
	suite.Require().NoError(err)
	return hex.EncodeToString(sig)
}
