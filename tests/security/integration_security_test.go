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
	settlementkeeper "github.com/stateset/core/x/settlement/keeper"
	settlementtypes "github.com/stateset/core/x/settlement/types"
	stablecoinkeeper "github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

type IntegrationSecurityTestSuite struct {
	apptesting.KeeperTestHelper

	denom string

	alice sdk.AccAddress
	bob   sdk.AccAddress
}

func TestIntegrationSecurityTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationSecurityTestSuite))
}

func (suite *IntegrationSecurityTestSuite) SetupTest() {
	suite.KeeperTestHelper.Setup()

	suite.denom = settlementtypes.StablecoinDenom
	suite.alice = suite.TestAccs[0]
	suite.bob = suite.TestAccs[1]

	suite.setCompliant(suite.alice)
	suite.setCompliant(suite.bob)

	suite.FundAcc(suite.alice, sdk.NewCoins(sdk.NewInt64Coin(suite.denom, 100_000)))
	suite.FundAcc(suite.alice, sdk.NewCoins(sdk.NewInt64Coin("usdc", 100_000)))
}

func (suite *IntegrationSecurityTestSuite) compliantProfile(addr sdk.AccAddress) compliancetypes.Profile {
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

func (suite *IntegrationSecurityTestSuite) setCompliant(addr sdk.AccAddress) {
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), suite.compliantProfile(addr))
}

func (suite *IntegrationSecurityTestSuite) setSuspended(addr sdk.AccAddress) {
	profile := suite.compliantProfile(addr)
	profile.Status = compliancetypes.StatusSuspended
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), profile)
}

func (suite *IntegrationSecurityTestSuite) TestSettlement_DoesNotRequireOracle() {
	// Settlement should not require oracle prices to move a fixed-denom stablecoin.
	msgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	_, err := msgServer.InstantTransfer(goCtx, settlementtypes.NewMsgInstantTransfer(
		suite.alice.String(),
		suite.bob.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"oracle-independence",
		"",
	))
	suite.Require().NoError(err)
}

func (suite *IntegrationSecurityTestSuite) TestComplianceBlocksAcrossModules() {
	suite.setSuspended(suite.bob)

	// Payments: compliance error should propagate.
	paymentsMsgServer := paymentskeeper.NewMsgServerImpl(suite.App.PaymentsKeeper)
	_, err := paymentsMsgServer.CreatePayment(sdk.WrapSDKContext(suite.Ctx), paymentstypes.NewMsgCreatePayment(
		suite.alice.String(),
		suite.bob.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"blocked-payee",
	))
	suite.Require().ErrorIs(err, compliancetypes.ErrComplianceBlocked)

	// Settlement: compliance failure is surfaced as a settlement error.
	settlementMsgServer := settlementkeeper.NewMsgServerImpl(suite.App.SettlementKeeper)
	_, err = settlementMsgServer.InstantTransfer(sdk.WrapSDKContext(suite.Ctx), settlementtypes.NewMsgInstantTransfer(
		suite.alice.String(),
		suite.bob.String(),
		sdk.NewInt64Coin(suite.denom, 1_000),
		"blocked-recipient",
		"",
	))
	suite.Require().ErrorIs(err, settlementtypes.ErrComplianceCheckFailed)

	// Stablecoin reserve deposit: compliance gated when KYC is required.
	stablecoinMsgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	_, err = stablecoinMsgServer.DepositReserve(sdk.WrapSDKContext(suite.Ctx), &stablecointypes.MsgDepositReserve{
		Depositor: suite.bob.String(),
		Amount:    sdk.NewInt64Coin("usdc", 1),
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrKYCRequired)
}

