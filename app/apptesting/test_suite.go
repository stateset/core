package apptesting

import (
	"fmt"
	"testing"
	"time"

	"cosmossdk.io/log"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/stateset/core/app"
)

type KeeperTestHelper struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

func (s *KeeperTestHelper) Setup() {
	s.App = app.New(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		app.MakeEncodingConfig(app.ModuleBasics),
		nil,
	)
	s.Ctx = s.App.BaseApp.NewContext(false)
	s.Ctx = s.Ctx.WithBlockHeight(1).WithChainID("stateset-1").WithBlockTime(time.Now().UTC())
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}

	s.TestAccs = CreateRandomAccounts(3)
}

func (s *KeeperTestHelper) SetupTestForInitGenesis() {
	// Reset app state
	s.App = app.New(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		app.MakeEncodingConfig(app.ModuleBasics),
		nil,
	)
	s.Ctx = s.App.BaseApp.NewContext(false)
	s.Ctx = s.Ctx.WithBlockHeight(1).WithChainID("stateset-1").WithBlockTime(time.Now().UTC())
}

// CreateRandomAccounts is a function return a list of randomly generated AccAddresses
func CreateRandomAccounts(numAccts int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, numAccts)
	for i := 0; i < numAccts; i++ {
		pk := secp256k1.GenPrivKey().PubKey()
		testAddrs[i] = sdk.AccAddress(pk.Address())
	}

	return testAddrs
}

// FundAcc funds target address with specified amount.
func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	err := FundAccount(s.App.BankKeeper, s.Ctx, acc, amounts)
	s.Require().NoError(err)
}

// FundModuleAcc funds target modules with specified amount.
func (s *KeeperTestHelper) FundModuleAcc(moduleName string, amounts sdk.Coins) {
	err := FundModuleAccount(s.App.BankKeeper, s.Ctx, moduleName, amounts)
	s.Require().NoError(err)
}

// FundAccount mints coins and sends them to the given account.
func FundAccount(bk bankkeeper.Keeper, ctx sdk.Context, addr sdk.AccAddress, amounts sdk.Coins) error {
	if err := bk.MintCoins(ctx, banktypes.ModuleName, amounts); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToAccount(ctx, banktypes.ModuleName, addr, amounts)
}

// FundModuleAccount mints coins and sends them to the given module account.
func FundModuleAccount(bk bankkeeper.Keeper, ctx sdk.Context, recipientMod string, amounts sdk.Coins) error {
	if err := bk.MintCoins(ctx, banktypes.ModuleName, amounts); err != nil {
		return err
	}
	return bk.SendCoinsFromModuleToModule(ctx, banktypes.ModuleName, recipientMod, amounts)
}

func (s *KeeperTestHelper) RunMsg(msg sdk.Msg) (*sdk.Result, error) {
	// Create a new context for each message to simulate a new block
	ctx := s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + 1)

	msgServer := s.App.MsgServiceRouter()
	handler := msgServer.Handler(msg)
	if handler == nil {
		return nil, fmt.Errorf("handler not found for message %T", msg)
	}

	return handler(ctx, msg)
}

func (s *KeeperTestHelper) AllocateRewardsToValidator(valAddr sdk.ValAddress, rewardAmt sdkmath.Int) {
	validator, err := s.App.StakingKeeper.GetValidator(s.Ctx, valAddr)
	s.Require().NoError(err)

	// allocate reward tokens to distribution module
	bondDenom, err := s.App.StakingKeeper.BondDenom(s.Ctx)
	s.Require().NoError(err)
	coins := sdk.Coins{sdk.NewCoin(bondDenom, rewardAmt)}
	err = FundModuleAccount(s.App.BankKeeper, s.Ctx, "distribution", coins)
	s.Require().NoError(err)

	// allocate rewards to validator
	err = s.App.DistrKeeper.AllocateTokensToValidator(s.Ctx, validator, sdk.NewDecCoinsFromCoins(coins...))
	s.Require().NoError(err)
}

func (s *KeeperTestHelper) BeginNewBlock() {
	s.Ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + 1).
		WithBlockTime(s.Ctx.BlockTime().Add(time.Second))
	_, err := s.App.BeginBlocker(s.Ctx)
	s.Require().NoError(err)
}

func (s *KeeperTestHelper) EndBlock() {
	_, err := s.App.EndBlocker(s.Ctx)
	s.Require().NoError(err)
}

func (s *KeeperTestHelper) Commit() {
	_, err := s.App.Commit()
	s.Require().NoError(err)
}

// Ensure this is used by tests
var _ = testing.T{}
