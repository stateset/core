package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	circuitkeeper "github.com/stateset/core/x/circuit/keeper"
	circuittypes "github.com/stateset/core/x/circuit/types"
)

type MsgServerTestSuite struct {
	suite.Suite
	keeper    circuitkeeper.Keeper
	ctx       sdk.Context
	msgServer circuittypes.MsgServer
	authority string
}

func TestMsgServerTestSuite(t *testing.T) {
	suite.Run(t, new(MsgServerTestSuite))
}

func (s *MsgServerTestSuite) SetupTest() {
	storeKey := storetypes.NewKVStoreKey(circuittypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	s.authority = "stateset1authority"
	s.keeper = circuitkeeper.NewKeeper(cdc, storeKey, s.authority)
	s.ctx = sdk.NewContext(stateStore, cmtproto.Header{
		Height:  1,
		ChainID: "stateset-test",
		Time:    time.Now(),
	}, false, log.NewNopLogger())

	s.msgServer = circuitkeeper.NewMsgServerImpl(s.keeper)
}

func (s *MsgServerTestSuite) TestMsgPauseSystem() {
	testCases := []struct {
		name      string
		msg       *circuittypes.MsgPauseSystem
		expectErr bool
		errType   error
	}{
		{
			name: "valid pause by authority",
			msg: &circuittypes.MsgPauseSystem{
				Authority:       s.authority,
				Reason:          "emergency maintenance",
				DurationSeconds: 3600,
			},
			expectErr: false,
		},
		{
			name: "invalid pause by non-authority",
			msg: &circuittypes.MsgPauseSystem{
				Authority:       "stateset1random",
				Reason:          "unauthorized pause",
				DurationSeconds: 3600,
			},
			expectErr: true,
			errType:   circuittypes.ErrUnauthorized,
		},
		{
			name: "pause with zero duration (indefinite)",
			msg: &circuittypes.MsgPauseSystem{
				Authority:       s.authority,
				Reason:          "indefinite pause",
				DurationSeconds: 0,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Reset state before each test
			if s.keeper.IsGloballyPaused(s.ctx) {
				err := s.keeper.ResumeSystem(s.ctx, s.authority)
				s.Require().NoError(err)
			}

			_, err := s.msgServer.PauseSystem(s.ctx, tc.msg)
			if tc.expectErr {
				s.Require().Error(err)
				if tc.errType != nil {
					s.Require().ErrorIs(err, tc.errType)
				}
			} else {
				s.Require().NoError(err)
				s.Require().True(s.keeper.IsGloballyPaused(s.ctx))
			}
		})
	}
}

func (s *MsgServerTestSuite) TestMsgResumeSystem() {
	// First pause the system
	err := s.keeper.PauseSystem(s.ctx, s.authority, "test pause", 0)
	s.Require().NoError(err)
	s.Require().True(s.keeper.IsGloballyPaused(s.ctx))

	testCases := []struct {
		name      string
		msg       *circuittypes.MsgResumeSystem
		expectErr bool
		errType   error
	}{
		{
			name: "invalid resume by non-authority",
			msg: &circuittypes.MsgResumeSystem{
				Authority: "stateset1random",
			},
			expectErr: true,
			errType:   circuittypes.ErrUnauthorized,
		},
		{
			name: "valid resume by authority",
			msg: &circuittypes.MsgResumeSystem{
				Authority: s.authority,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			_, err := s.msgServer.ResumeSystem(s.ctx, tc.msg)
			if tc.expectErr {
				s.Require().Error(err)
				if tc.errType != nil {
					s.Require().ErrorIs(err, tc.errType)
				}
			} else {
				s.Require().NoError(err)
				s.Require().False(s.keeper.IsGloballyPaused(s.ctx))
			}
		})
	}
}

func (s *MsgServerTestSuite) TestMsgTripCircuit() {
	testCases := []struct {
		name            string
		msg             *circuittypes.MsgTripCircuit
		expectErr       bool
		errType         error
		disabledMsgType string
	}{
		{
			name: "valid trip by authority",
			msg: &circuittypes.MsgTripCircuit{
				Authority:  s.authority,
				ModuleName: "stablecoin",
				Reason:     "testing circuit trip",
			},
			expectErr: false,
		},
		{
			name: "invalid trip by non-authority",
			msg: &circuittypes.MsgTripCircuit{
				Authority:  "stateset1random",
				ModuleName: "settlement",
				Reason:     "unauthorized trip",
			},
			expectErr: true,
			errType:   circuittypes.ErrUnauthorized,
		},
		{
			name: "trip with specific disabled messages",
			msg: &circuittypes.MsgTripCircuit{
				Authority:       s.authority,
				ModuleName:      "payments",
				Reason:          "disable specific operations",
				DisableMessages: []string{"/stateset.payments.v1.MsgCreatePayment"},
			},
			expectErr:       false,
			disabledMsgType: "/stateset.payments.v1.MsgCreatePayment",
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			// Reset circuit before each test
			if s.keeper.IsModuleCircuitOpen(s.ctx, tc.msg.ModuleName) {
				_ = s.keeper.ResetCircuit(s.ctx, tc.msg.ModuleName, s.authority)
			}

			_, err := s.msgServer.TripCircuit(s.ctx, tc.msg)
			if tc.expectErr {
				s.Require().Error(err)
				if tc.errType != nil {
					s.Require().ErrorIs(err, tc.errType)
				}
			} else {
				s.Require().NoError(err)
				s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, tc.msg.ModuleName))

				if tc.disabledMsgType != "" {
					s.Require().True(s.keeper.IsMessageDisabled(s.ctx, tc.msg.ModuleName, tc.disabledMsgType))
				}
			}
		})
	}
}

func (s *MsgServerTestSuite) TestMsgResetCircuit() {
	// First trip a circuit
	moduleName := "oracle"
	err := s.keeper.TripCircuit(s.ctx, moduleName, "test trip", s.authority, nil)
	s.Require().NoError(err)
	s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, moduleName))

	testCases := []struct {
		name      string
		msg       *circuittypes.MsgResetCircuit
		expectErr bool
		errType   error
	}{
		{
			name: "invalid reset by non-authority",
			msg: &circuittypes.MsgResetCircuit{
				Authority:  "stateset1random",
				ModuleName: moduleName,
			},
			expectErr: true,
			errType:   circuittypes.ErrUnauthorized,
		},
		{
			name: "valid reset by authority",
			msg: &circuittypes.MsgResetCircuit{
				Authority:  s.authority,
				ModuleName: moduleName,
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			_, err := s.msgServer.ResetCircuit(s.ctx, tc.msg)
			if tc.expectErr {
				s.Require().Error(err)
				if tc.errType != nil {
					s.Require().ErrorIs(err, tc.errType)
				}
			} else {
				s.Require().NoError(err)
				s.Require().False(s.keeper.IsModuleCircuitOpen(s.ctx, moduleName))
			}
		})
	}
}

func (s *MsgServerTestSuite) TestMsgUpdateParams() {
	testCases := []struct {
		name      string
		msg       *circuittypes.MsgUpdateParams
		expectErr bool
		errType   error
	}{
		{
			name: "valid update by authority",
			msg: &circuittypes.MsgUpdateParams{
				Authority: s.authority,
				Params: circuittypes.Params{
					DefaultFailureThreshold: 10,
					DefaultRecoveryPeriod:   120,
					MaxPauseDuration:        7200,
					Authorities:             []string{s.authority},
					RateLimits:              []circuittypes.RateLimitConfig{},
				},
			},
			expectErr: false,
		},
		{
			name: "invalid update by non-authority",
			msg: &circuittypes.MsgUpdateParams{
				Authority: "stateset1random",
				Params:    circuittypes.DefaultParams(),
			},
			expectErr: true,
			errType:   circuittypes.ErrUnauthorized,
		},
		{
			name: "invalid params - zero failure threshold",
			msg: &circuittypes.MsgUpdateParams{
				Authority: s.authority,
				Params: circuittypes.Params{
					DefaultFailureThreshold: 0, // Invalid
					DefaultRecoveryPeriod:   60,
					MaxPauseDuration:        3600,
				},
			},
			expectErr: true,
		},
		{
			name: "invalid params - zero recovery period",
			msg: &circuittypes.MsgUpdateParams{
				Authority: s.authority,
				Params: circuittypes.Params{
					DefaultFailureThreshold: 5,
					DefaultRecoveryPeriod:   0, // Invalid
					MaxPauseDuration:        3600,
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			_, err := s.msgServer.UpdateParams(s.ctx, tc.msg)
			if tc.expectErr {
				s.Require().Error(err)
				if tc.errType != nil {
					s.Require().ErrorIs(err, tc.errType)
				}
			} else {
				s.Require().NoError(err)
				params := s.keeper.GetParams(s.ctx)
				s.Require().Equal(tc.msg.Params.DefaultFailureThreshold, params.DefaultFailureThreshold)
			}
		})
	}
}

func (s *MsgServerTestSuite) TestPauseResumeCycle() {
	// Test full pause/resume cycle
	pauseMsg := &circuittypes.MsgPauseSystem{
		Authority:       s.authority,
		Reason:          "cycle test",
		DurationSeconds: 3600,
	}

	// Pause
	_, err := s.msgServer.PauseSystem(s.ctx, pauseMsg)
	s.Require().NoError(err)
	s.Require().True(s.keeper.IsGloballyPaused(s.ctx))

	// Try to pause again (should fail)
	_, err = s.msgServer.PauseSystem(s.ctx, pauseMsg)
	s.Require().Error(err)
	s.Require().ErrorIs(err, circuittypes.ErrAlreadyPaused)

	// Resume
	resumeMsg := &circuittypes.MsgResumeSystem{
		Authority: s.authority,
	}
	_, err = s.msgServer.ResumeSystem(s.ctx, resumeMsg)
	s.Require().NoError(err)
	s.Require().False(s.keeper.IsGloballyPaused(s.ctx))

	// Try to resume again (should fail)
	_, err = s.msgServer.ResumeSystem(s.ctx, resumeMsg)
	s.Require().Error(err)
	s.Require().ErrorIs(err, circuittypes.ErrNotPaused)
}

func (s *MsgServerTestSuite) TestTripResetCycle() {
	moduleName := "compliance"

	// Trip
	tripMsg := &circuittypes.MsgTripCircuit{
		Authority:  s.authority,
		ModuleName: moduleName,
		Reason:     "trip test",
	}
	_, err := s.msgServer.TripCircuit(s.ctx, tripMsg)
	s.Require().NoError(err)
	s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, moduleName))

	// Reset
	resetMsg := &circuittypes.MsgResetCircuit{
		Authority:  s.authority,
		ModuleName: moduleName,
	}
	_, err = s.msgServer.ResetCircuit(s.ctx, resetMsg)
	s.Require().NoError(err)
	s.Require().False(s.keeper.IsModuleCircuitOpen(s.ctx, moduleName))
}

func (s *MsgServerTestSuite) TestMultipleModuleCircuits() {
	modules := []string{"stablecoin", "settlement", "payments", "oracle"}

	// Trip all circuits
	for _, mod := range modules {
		tripMsg := &circuittypes.MsgTripCircuit{
			Authority:  s.authority,
			ModuleName: mod,
			Reason:     "multi-module test",
		}
		_, err := s.msgServer.TripCircuit(s.ctx, tripMsg)
		s.Require().NoError(err)
		s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, mod))
	}

	// Verify all are open
	for _, mod := range modules {
		s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, mod))
	}

	// Reset only some
	for i, mod := range modules[:2] {
		resetMsg := &circuittypes.MsgResetCircuit{
			Authority:  s.authority,
			ModuleName: mod,
		}
		_, err := s.msgServer.ResetCircuit(s.ctx, resetMsg)
		s.Require().NoError(err)
		s.Require().False(s.keeper.IsModuleCircuitOpen(s.ctx, mod), "module %d should be closed", i)
	}

	// Verify partial reset
	s.Require().False(s.keeper.IsModuleCircuitOpen(s.ctx, modules[0]))
	s.Require().False(s.keeper.IsModuleCircuitOpen(s.ctx, modules[1]))
	s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, modules[2]))
	s.Require().True(s.keeper.IsModuleCircuitOpen(s.ctx, modules[3]))
}
