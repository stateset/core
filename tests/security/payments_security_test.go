package security_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/stateset/core/app/apptesting"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	paymentskeeper "github.com/stateset/core/x/payments/keeper"
	paymentstypes "github.com/stateset/core/x/payments/types"
)

type PaymentsSecurityTestSuite struct {
	apptesting.KeeperTestHelper

	denom string

	payer sdk.AccAddress
	payee sdk.AccAddress
	other sdk.AccAddress
}

func TestPaymentsSecurityTestSuite(t *testing.T) {
	suite.Run(t, new(PaymentsSecurityTestSuite))
}

func (suite *PaymentsSecurityTestSuite) SetupTest() {
	suite.KeeperTestHelper.Setup()

	suite.denom = "ssusd"
	suite.payer = suite.TestAccs[0]
	suite.payee = suite.TestAccs[1]
	suite.other = suite.TestAccs[2]

	suite.setCompliant(suite.payer)
	suite.setCompliant(suite.payee)
	suite.setCompliant(suite.other)

	suite.FundAcc(suite.payer, sdk.NewCoins(sdk.NewInt64Coin(suite.denom, 10_000)))
}

func (suite *PaymentsSecurityTestSuite) compliantProfile(addr sdk.AccAddress) compliancetypes.Profile {
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

func (suite *PaymentsSecurityTestSuite) setCompliant(addr sdk.AccAddress) {
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), suite.compliantProfile(addr))
}

func (suite *PaymentsSecurityTestSuite) setSanctioned(addr sdk.AccAddress) {
	profile := suite.compliantProfile(addr)
	profile.Sanction = true
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), profile)
}

func (suite *PaymentsSecurityTestSuite) setSuspended(addr sdk.AccAddress) {
	profile := suite.compliantProfile(addr)
	profile.Status = compliancetypes.StatusSuspended
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), profile)
}

func (suite *PaymentsSecurityTestSuite) TestSettlePayment_PayeeOnly() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-1",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.other.String(), resp.PaymentId))
	suite.Require().ErrorIs(err, paymentstypes.ErrNotAuthorized)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payer.String(), resp.PaymentId))
	suite.Require().ErrorIs(err, paymentstypes.ErrNotAuthorized)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().NoError(err)

	payment, found := suite.App.PaymentsKeeper.GetPayment(suite.Ctx, resp.PaymentId)
	suite.Require().True(found)
	suite.Require().Equal(paymentstypes.PaymentStatusSettled, payment.Status)

	suite.Require().Equal(int64(9_000), suite.App.BankKeeper.GetBalance(suite.Ctx, suite.payer, suite.denom).Amount.Int64())
	suite.Require().Equal(int64(1_000), suite.App.BankKeeper.GetBalance(suite.Ctx, suite.payee, suite.denom).Amount.Int64())
	moduleAddr := suite.App.AccountKeeper.GetModuleAddress(paymentstypes.ModuleAccountName)
	suite.Require().Equal(int64(0), suite.App.BankKeeper.GetBalance(suite.Ctx, moduleAddr, suite.denom).Amount.Int64())
}

func (suite *PaymentsSecurityTestSuite) TestCancelPayment_PayerOnly() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-2",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.CancelPayment(goCtx, paymentstypes.NewMsgCancelPayment(suite.payee.String(), resp.PaymentId, "test"))
	suite.Require().ErrorIs(err, paymentstypes.ErrNotAuthorized)

	_, err = msgServer.CancelPayment(goCtx, paymentstypes.NewMsgCancelPayment(suite.other.String(), resp.PaymentId, "test"))
	suite.Require().ErrorIs(err, paymentstypes.ErrNotAuthorized)

	_, err = msgServer.CancelPayment(goCtx, paymentstypes.NewMsgCancelPayment(suite.payer.String(), resp.PaymentId, "test"))
	suite.Require().NoError(err)

	payment, found := suite.App.PaymentsKeeper.GetPayment(suite.Ctx, resp.PaymentId)
	suite.Require().True(found)
	suite.Require().Equal(paymentstypes.PaymentStatusCancelled, payment.Status)

	suite.Require().Equal(int64(10_000), suite.App.BankKeeper.GetBalance(suite.Ctx, suite.payer, suite.denom).Amount.Int64())
	moduleAddr := suite.App.AccountKeeper.GetModuleAddress(paymentstypes.ModuleAccountName)
	suite.Require().Equal(int64(0), suite.App.BankKeeper.GetBalance(suite.Ctx, moduleAddr, suite.denom).Amount.Int64())
}

func (suite *PaymentsSecurityTestSuite) TestPaymentStatus_NoDoubleSettle() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-3",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().NoError(err)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().ErrorIs(err, paymentstypes.ErrPaymentCompleted)

	suite.Require().Equal(int64(1_000), suite.App.BankKeeper.GetBalance(suite.Ctx, suite.payee, suite.denom).Amount.Int64())
}

func (suite *PaymentsSecurityTestSuite) TestPaymentStatus_NoCancelAfterSettle() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-4",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().NoError(err)

	_, err = msgServer.CancelPayment(goCtx, paymentstypes.NewMsgCancelPayment(suite.payer.String(), resp.PaymentId, "test"))
	suite.Require().ErrorIs(err, paymentstypes.ErrPaymentCompleted)
}

func (suite *PaymentsSecurityTestSuite) TestPaymentStatus_NoSettleAfterCancel() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-5",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	_, err = msgServer.CancelPayment(goCtx, paymentstypes.NewMsgCancelPayment(suite.payer.String(), resp.PaymentId, "test"))
	suite.Require().NoError(err)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().ErrorIs(err, paymentstypes.ErrPaymentCancelled)
}

func (suite *PaymentsSecurityTestSuite) TestEscrow_InsufficientBalance() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 20_000),
		"invoice-too-big",
	)
	_, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().ErrorIs(err, paymentstypes.ErrInsufficientBalance)
}

func (suite *PaymentsSecurityTestSuite) TestCompliance_BothPartiesChecked() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	suite.setSanctioned(suite.payee)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-sanctioned",
	)
	_, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().ErrorIs(err, compliancetypes.ErrSanctionedAddress)
}

func (suite *PaymentsSecurityTestSuite) TestCompliance_StatusChangeBlocksSettlement() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	create := paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-status-change",
	)
	resp, err := msgServer.CreatePayment(goCtx, create)
	suite.Require().NoError(err)

	suite.setSuspended(suite.payee)

	_, err = msgServer.SettlePayment(goCtx, paymentstypes.NewMsgSettlePayment(suite.payee.String(), resp.PaymentId))
	suite.Require().ErrorIs(err, compliancetypes.ErrComplianceBlocked)
}

func (suite *PaymentsSecurityTestSuite) TestPayment_ValidAddresses() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	_, err := msgServer.CreatePayment(goCtx, paymentstypes.NewMsgCreatePayment(
		"not-an-address",
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invalid",
	))
	suite.Require().ErrorIs(err, paymentstypes.ErrInvalidPayment)

	_, err = msgServer.CreatePayment(goCtx, paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		"not-an-address",
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invalid",
	))
	suite.Require().ErrorIs(err, paymentstypes.ErrInvalidPayment)
}

func (suite *PaymentsSecurityTestSuite) TestPayment_DifferentParties() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	_, err := msgServer.CreatePayment(goCtx, paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payer.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"self",
	))
	suite.Require().ErrorIs(err, paymentstypes.ErrInvalidPayment)
}

func (suite *PaymentsSecurityTestSuite) TestPaymentID_Unique() {
	msgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	resp1, err := msgServer.CreatePayment(goCtx, paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-1",
	))
	suite.Require().NoError(err)

	resp2, err := msgServer.CreatePayment(goCtx, paymentstypes.NewMsgCreatePayment(
		suite.payer.String(),
		suite.payee.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"invoice-2",
	))
	suite.Require().NoError(err)

	suite.Require().NotEqual(resp1.PaymentId, resp2.PaymentId)
	suite.Require().Equal(resp1.PaymentId+1, resp2.PaymentId)
}
