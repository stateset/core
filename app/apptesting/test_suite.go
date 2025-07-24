package apptesting

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/types/signing"
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
		tmdb.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		app.MakeEncodingConfig(app.ModuleBasics),
		nil,
	)
	s.Ctx = s.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stateset-1", Time: time.Now().UTC()})
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
		tmdb.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		app.DefaultNodeHome,
		5,
		app.MakeEncodingConfig(app.ModuleBasics),
		nil,
	)
	s.Ctx = s.App.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stateset-1", Time: time.Now().UTC()})
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
	err := testutil.FundAccount(s.App.BankKeeper, s.Ctx, acc, amounts)
	s.Require().NoError(err)
}

// FundModuleAcc funds target modules with specified amount.
func (s *KeeperTestHelper) FundModuleAcc(moduleName string, amounts sdk.Coins) {
	err := testutil.FundModuleAccount(s.App.BankKeeper, s.Ctx, moduleName, amounts)
	s.Require().NoError(err)
}

func (s *KeeperTestHelper) BuildTx(
	txBuilder sdk.TxBuilder,
	msgs []sdk.Msg,
	sigV2 signing.SignatureV2,
	memo string,
	txFee sdk.Coins,
	gasLimit uint64,
) sdk.Tx {
	err := txBuilder.SetMsgs(msgs...)
	s.Require().NoError(err)

	err = txBuilder.SetSignatures(sigV2)
	s.Require().NoError(err)

	txBuilder.SetMemo(memo)
	txBuilder.SetFeeAmount(txFee)
	txBuilder.SetGasLimit(gasLimit)

	return txBuilder.GetTx()
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

func (s *KeeperTestHelper) AllocateRewardsToValidator(valAddr sdk.ValAddress, rewardAmt sdk.Int) {
	validator := s.App.StakingKeeper.GetValidator(s.Ctx, valAddr)
	
	// allocate reward tokens to distribution module
	coins := sdk.Coins{sdk.NewCoin(s.App.StakingKeeper.BondDenom(s.Ctx), rewardAmt)}
	err := testutil.FundModuleAccount(s.App.BankKeeper, s.Ctx, "distribution", coins)
	s.Require().NoError(err)

	// allocate rewards to validator
	s.App.DistrKeeper.AllocateTokensToValidator(s.Ctx, validator, sdk.DecCoinsFromCoins(coins...))
}

func (s *KeeperTestHelper) BeginNewBlock() {
	valAddr := []byte("validator")
	validator := abci.Validator{
		Address: valAddr,
		Power:   100,
	}

	s.Ctx = s.Ctx.WithBlockHeight(s.Ctx.BlockHeight() + 1).
		WithBlockTime(s.Ctx.BlockTime().Add(time.Second))
	s.App.BeginBlocker(s.Ctx, abci.RequestBeginBlock{
		Header: tmproto.Header{Height: s.Ctx.BlockHeight(), Time: s.Ctx.BlockTime()},
		LastCommitInfo: abci.LastCommitInfo{
			Votes: []abci.VoteInfo{{
				Validator:       validator,
				SignedLastBlock: true,
			}},
		},
	})
}

func (s *KeeperTestHelper) EndBlock() {
	s.App.EndBlocker(s.Ctx, abci.RequestEndBlock{Height: s.Ctx.BlockHeight()})
}

func (s *KeeperTestHelper) Commit() {
	s.App.Commit()
}