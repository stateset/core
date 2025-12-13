package security_test

import (
	"testing"
	"time"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/stateset/core/app/apptesting"
	compliancetypes "github.com/stateset/core/x/compliance/types"
	oracletypes "github.com/stateset/core/x/oracle/types"
	stablecoinkeeper "github.com/stateset/core/x/stablecoin/keeper"
	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

type StablecoinSecurityTestSuite struct {
	apptesting.KeeperTestHelper

	authority string

	owner      sdk.AccAddress
	attacker   sdk.AccAddress
	liquidator sdk.AccAddress
}

func TestStablecoinSecurityTestSuite(t *testing.T) {
	suite.Run(t, new(StablecoinSecurityTestSuite))
}

func (suite *StablecoinSecurityTestSuite) SetupTest() {
	suite.KeeperTestHelper.Setup()

	suite.authority = suite.App.StablecoinKeeper.GetAuthority()

	suite.owner = suite.TestAccs[0]
	suite.attacker = suite.TestAccs[1]
	suite.liquidator = suite.TestAccs[2]

	suite.setCompliant(suite.owner)
	suite.setCompliant(suite.attacker)
	suite.setCompliant(suite.liquidator)

	// Enable vault minting for vault security tests.
	params := stablecointypes.DefaultParams()
	params.VaultMintingEnabled = true
	suite.App.StablecoinKeeper.SetParams(suite.Ctx, params)

	// Set a fresh oracle price for the default collateral denom.
	suite.setOraclePrice("stst", sdkmath.LegacyOneDec(), suite.Ctx.BlockTime())
	// Reserve-backed ssUSD uses USTN (tokenized US Treasury Notes) by default.
	suite.setOraclePrice("ustn", sdkmath.LegacyOneDec(), suite.Ctx.BlockTime())

	suite.FundAcc(suite.owner, sdk.NewCoins(sdk.NewInt64Coin("stst", 1_000_000)))
	suite.FundAcc(suite.liquidator, sdk.NewCoins(sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 1_000_000)))

	// Make reserve-backed tests easier to reason about.
	reserveParams := stablecointypes.DefaultReserveParams()
	reserveParams.MintFeeBps = 0
	reserveParams.RedeemFeeBps = 0
	reserveParams.MinMintAmount = sdkmath.NewInt(1)
	reserveParams.MinRedeemAmount = sdkmath.NewInt(1)
	reserveParams.MaxDailyMint = sdkmath.NewInt(1_000_000_000)
	reserveParams.MaxDailyRedeem = sdkmath.NewInt(1_000_000_000)
	reserveParams.RedemptionDelay = time.Hour
	suite.Require().NoError(suite.App.StablecoinKeeper.SetReserveParams(suite.Ctx, reserveParams))
}

func (suite *StablecoinSecurityTestSuite) compliantProfile(addr sdk.AccAddress) compliancetypes.Profile {
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

func (suite *StablecoinSecurityTestSuite) setCompliant(addr sdk.AccAddress) {
	suite.App.ComplianceKeeper.SetProfile(sdk.WrapSDKContext(suite.Ctx), suite.compliantProfile(addr))
}

func (suite *StablecoinSecurityTestSuite) setOraclePrice(denom string, amount sdkmath.LegacyDec, updatedAt time.Time) {
	suite.App.OracleKeeper.SetPrice(sdk.WrapSDKContext(suite.Ctx), oracletypes.Price{
		Denom:      denom,
		Amount:     amount,
		LastHeight: suite.Ctx.BlockHeight(),
		UpdatedAt:  updatedAt,
	})
}

func (suite *StablecoinSecurityTestSuite) TestVaultLiquidation_OnlyUndercollateralized() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	createResp, err := msgServer.CreateVault(goCtx, &stablecointypes.MsgCreateVault{
		Owner:      suite.owner.String(),
		Collateral: sdk.NewInt64Coin("stst", 200),
		Debt:       sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 100),
	})
	suite.Require().NoError(err)

	// Healthy vault should not be liquidatable.
	_, err = msgServer.LiquidateVault(goCtx, &stablecointypes.MsgLiquidateVault{
		Liquidator: suite.liquidator.String(),
		VaultId:    createResp.VaultId,
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrVaultHealthy)

	// Drop price to push vault below liquidation ratio and allow liquidation.
	suite.setOraclePrice("stst", sdkmath.LegacyMustNewDecFromStr("0.5"), suite.Ctx.BlockTime())

	_, err = msgServer.LiquidateVault(goCtx, &stablecointypes.MsgLiquidateVault{
		Liquidator: suite.liquidator.String(),
		VaultId:    createResp.VaultId,
	})
	suite.Require().NoError(err)

	_, found := suite.App.StablecoinKeeper.GetVault(suite.Ctx, createResp.VaultId)
	suite.Require().False(found)
}

func (suite *StablecoinSecurityTestSuite) TestVaultLiquidation_FailsWhenPriceStale() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	createResp, err := msgServer.CreateVault(goCtx, &stablecointypes.MsgCreateVault{
		Owner:      suite.owner.String(),
		Collateral: sdk.NewInt64Coin("stst", 200),
		Debt:       sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 100),
	})
	suite.Require().NoError(err)

	// Make the oracle price stale.
	suite.setOraclePrice("stst", sdkmath.LegacyOneDec(), suite.Ctx.BlockTime().Add(-2*time.Hour))

	_, err = msgServer.LiquidateVault(goCtx, &stablecointypes.MsgLiquidateVault{
		Liquidator: suite.liquidator.String(),
		VaultId:    createResp.VaultId,
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrPriceStale)
}

func (suite *StablecoinSecurityTestSuite) TestVaultOperations_OwnerOnly() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	createResp, err := msgServer.CreateVault(goCtx, &stablecointypes.MsgCreateVault{
		Owner:      suite.owner.String(),
		Collateral: sdk.NewInt64Coin("stst", 500),
		Debt:       sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 0),
	})
	suite.Require().NoError(err)

	_, err = msgServer.MintStablecoin(goCtx, &stablecointypes.MsgMintStablecoin{
		Owner:   suite.attacker.String(),
		VaultId: createResp.VaultId,
		Amount:  sdk.NewInt64Coin(stablecointypes.StablecoinDenom, 1),
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrUnauthorized)

	_, err = msgServer.WithdrawCollateral(goCtx, &stablecointypes.MsgWithdrawCollateral{
		Owner:      suite.attacker.String(),
		VaultId:    createResp.VaultId,
		Collateral: sdk.NewInt64Coin("stst", 1),
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrUnauthorized)
}

func (suite *StablecoinSecurityTestSuite) TestReserveParams_AuthorityOnly() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	newParams := stablecointypes.DefaultReserveParams()
	newParams.RequireKyc = false

	_, err := msgServer.UpdateReserveParams(goCtx, &stablecointypes.MsgUpdateReserveParams{
		Authority: suite.owner.String(),
		Params:    newParams,
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrUnauthorized)

	_, err = msgServer.UpdateReserveParams(goCtx, &stablecointypes.MsgUpdateReserveParams{
		Authority: suite.authority,
		Params:    newParams,
	})
	suite.Require().NoError(err)
}

func (suite *StablecoinSecurityTestSuite) TestAttestation_ApprovedAttesterOnly() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	_, err := msgServer.RecordAttestation(goCtx, &stablecointypes.MsgRecordAttestation{
		Attester:      suite.owner.String(),
		TotalCash:     "0",
		TotalTbills:   "0",
		TotalTnotes:   "0",
		TotalTbonds:   "0",
		TotalRepos:    "0",
		TotalMmf:      "0",
		TotalValue:    "1",
		CustodianName: "custodian",
		AuditFirm:     "auditor",
		ReportDate:    "2025-01-01",
		Hash:          "hash",
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrInvalidAttester)

	_, err = msgServer.SetApprovedAttester(goCtx, &stablecointypes.MsgSetApprovedAttester{
		Authority: suite.authority,
		Attester:  suite.owner.String(),
		Approved:  true,
	})
	suite.Require().NoError(err)

	_, err = msgServer.RecordAttestation(goCtx, &stablecointypes.MsgRecordAttestation{
		Attester:      suite.owner.String(),
		TotalCash:     "0",
		TotalTbills:   "0",
		TotalTnotes:   "0",
		TotalTbonds:   "0",
		TotalRepos:    "0",
		TotalMmf:      "0",
		TotalValue:    "1",
		CustodianName: "custodian",
		AuditFirm:     "auditor",
		ReportDate:    "2025-01-01",
		Hash:          "hash",
	})
	suite.Require().NoError(err)
}

func (suite *StablecoinSecurityTestSuite) TestRedemptionDelay_Enforced() {
	msgServer := stablecoinkeeper.NewMsgServerImpl(suite.App.StablecoinKeeper)

	// Fund reserve depositor with tokenized Treasury Note reserves.
	suite.FundAcc(suite.owner, sdk.NewCoins(sdk.NewInt64Coin("ustn", 10_000)))

	_, err := msgServer.DepositReserve(sdk.WrapSDKContext(suite.Ctx), &stablecointypes.MsgDepositReserve{
		Depositor: suite.owner.String(),
		Amount:    sdk.NewInt64Coin("ustn", 10_000),
	})
	suite.Require().NoError(err)

	reqResp, err := msgServer.RequestRedemption(sdk.WrapSDKContext(suite.Ctx), &stablecointypes.MsgRequestRedemption{
		Requester:   suite.owner.String(),
		SsusdAmount: "1000",
		OutputDenom: "ustn",
	})
	suite.Require().NoError(err)

	_, err = msgServer.ExecuteRedemption(sdk.WrapSDKContext(suite.Ctx), &stablecointypes.MsgExecuteRedemption{
		Executor:     suite.attacker.String(),
		RedemptionId: reqResp.RedemptionId,
	})
	suite.Require().ErrorIs(err, stablecointypes.ErrRedemptionNotReady)

	// Advance time beyond delay.
	suite.Ctx = suite.Ctx.WithBlockTime(suite.Ctx.BlockTime().Add(2 * time.Hour))
	_, err = msgServer.ExecuteRedemption(sdk.WrapSDKContext(suite.Ctx), &stablecointypes.MsgExecuteRedemption{
		Executor:     suite.attacker.String(),
		RedemptionId: reqResp.RedemptionId,
	})
	suite.Require().NoError(err)
}
