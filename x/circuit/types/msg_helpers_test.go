package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/stateset/core/x/circuit/types"
)

func TestMsgPauseSystem_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgPauseSystem
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgPauseSystem{
				Authority:       validAuthority,
				Reason:          "Emergency maintenance",
				DurationSeconds: 3600,
			},
			expectErr: false,
		},
		{
			name: "valid with zero duration (indefinite)",
			msg: &types.MsgPauseSystem{
				Authority:       validAuthority,
				Reason:          "Emergency maintenance",
				DurationSeconds: 0,
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgPauseSystem{
				Authority:       "invalid",
				Reason:          "Emergency maintenance",
				DurationSeconds: 3600,
			},
			expectErr: true,
		},
		{
			name: "empty reason",
			msg: &types.MsgPauseSystem{
				Authority:       validAuthority,
				Reason:          "",
				DurationSeconds: 3600,
			},
			expectErr: true,
		},
		{
			name: "negative duration",
			msg: &types.MsgPauseSystem{
				Authority:       validAuthority,
				Reason:          "Emergency maintenance",
				DurationSeconds: -1,
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgResumeSystem_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgResumeSystem
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgResumeSystem{
				Authority: validAuthority,
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgResumeSystem{
				Authority: "invalid",
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgTripCircuit_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgTripCircuit
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgTripCircuit{
				Authority:       validAuthority,
				ModuleName:      "stablecoin",
				Reason:          "Oracle failure",
				DisableMessages: []string{"/stateset.stablecoin.Msg/MintStablecoin"},
			},
			expectErr: false,
		},
		{
			name: "valid without disable messages",
			msg: &types.MsgTripCircuit{
				Authority:  validAuthority,
				ModuleName: "stablecoin",
				Reason:     "Oracle failure",
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgTripCircuit{
				Authority:  "invalid",
				ModuleName: "stablecoin",
				Reason:     "Oracle failure",
			},
			expectErr: true,
		},
		{
			name: "empty module name",
			msg: &types.MsgTripCircuit{
				Authority:  validAuthority,
				ModuleName: "",
				Reason:     "Oracle failure",
			},
			expectErr: true,
		},
		{
			name: "empty reason",
			msg: &types.MsgTripCircuit{
				Authority:  validAuthority,
				ModuleName: "stablecoin",
				Reason:     "",
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgResetCircuit_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgResetCircuit
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgResetCircuit{
				Authority:  validAuthority,
				ModuleName: "stablecoin",
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgResetCircuit{
				Authority:  "invalid",
				ModuleName: "stablecoin",
			},
			expectErr: true,
		},
		{
			name: "empty module name",
			msg: &types.MsgResetCircuit{
				Authority:  validAuthority,
				ModuleName: "",
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	validAuthority := sdk.AccAddress("authority___________").String()

	tests := []struct {
		name      string
		msg       *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name: "valid message",
			msg: &types.MsgUpdateParams{
				Authority: validAuthority,
				Params:    types.DefaultParams(),
			},
			expectErr: false,
		},
		{
			name: "invalid authority",
			msg: &types.MsgUpdateParams{
				Authority: "invalid",
				Params:    types.DefaultParams(),
			},
			expectErr: true,
		},
		{
			name: "invalid params - zero failure threshold",
			msg: &types.MsgUpdateParams{
				Authority: validAuthority,
				Params: types.Params{
					DefaultFailureThreshold: 0,
					DefaultRecoveryPeriod:   300,
					MaxPauseDuration:        86400,
				},
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDefaultParams(t *testing.T) {
	params := types.DefaultParams()

	require.Equal(t, uint64(5), params.DefaultFailureThreshold)
	require.Equal(t, int64(300), params.DefaultRecoveryPeriod)
	require.Equal(t, int64(86400), params.MaxPauseDuration)
	require.Len(t, params.RateLimits, 4)

	// Validate that default params are valid
	require.NoError(t, params.Validate())
}

func TestParams_Validate(t *testing.T) {
	tests := []struct {
		name      string
		params    types.Params
		expectErr bool
	}{
		{
			name:      "valid default params",
			params:    types.DefaultParams(),
			expectErr: false,
		},
		{
			name: "zero failure threshold",
			params: types.Params{
				DefaultFailureThreshold: 0,
				DefaultRecoveryPeriod:   300,
				MaxPauseDuration:        86400,
			},
			expectErr: true,
		},
		{
			name: "zero recovery period",
			params: types.Params{
				DefaultFailureThreshold: 5,
				DefaultRecoveryPeriod:   0,
				MaxPauseDuration:        86400,
			},
			expectErr: true,
		},
		{
			name: "negative recovery period",
			params: types.Params{
				DefaultFailureThreshold: 5,
				DefaultRecoveryPeriod:   -1,
				MaxPauseDuration:        86400,
			},
			expectErr: true,
		},
		{
			name: "zero max pause duration",
			params: types.Params{
				DefaultFailureThreshold: 5,
				DefaultRecoveryPeriod:   300,
				MaxPauseDuration:        0,
			},
			expectErr: true,
		},
		{
			name: "negative max pause duration",
			params: types.Params{
				DefaultFailureThreshold: 5,
				DefaultRecoveryPeriod:   300,
				MaxPauseDuration:        -1,
			},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDefaultOracleDeviationConfigs(t *testing.T) {
	configs := types.DefaultOracleDeviationConfigs()

	require.Len(t, configs, 2)

	// Check uatom config
	atomConfig := configs[0]
	require.Equal(t, "uatom", atomConfig.Denom)
	require.Equal(t, uint64(500), atomConfig.MaxDeviationBps)
	require.Equal(t, uint64(2000), atomConfig.MaxDailyDeviationBps)

	// Check uusdc config
	usdcConfig := configs[1]
	require.Equal(t, "uusdc", usdcConfig.Denom)
	require.Equal(t, uint64(100), usdcConfig.MaxDeviationBps)
	require.Equal(t, uint64(200), usdcConfig.MaxDailyDeviationBps)
}

func TestDefaultLiquidationSurgeProtection(t *testing.T) {
	protection := types.DefaultLiquidationSurgeProtection()

	require.Equal(t, uint64(10), protection.MaxLiquidationsPerBlock)
	require.Equal(t, uint64(5), protection.CooldownBlocks)
	require.True(t, protection.CurrentBlockValue.IsZero())
}

func TestDefaultGenesis(t *testing.T) {
	genesis := types.DefaultGenesis()

	require.NotNil(t, genesis)
	require.False(t, genesis.CircuitState.GlobalPaused)
	require.Empty(t, genesis.ModuleCircuits)
	require.Empty(t, genesis.RateLimitStates)
	require.NoError(t, genesis.Validate())
}

func TestMarshalUnmarshalCircuitState(t *testing.T) {
	original := types.CircuitState{
		GlobalPaused: true,
		Reason:       "test",
	}

	// Marshal
	bz, err := types.MarshalCircuitState(original)
	require.NoError(t, err)
	require.NotEmpty(t, bz)

	// Unmarshal
	restored, err := types.UnmarshalCircuitState(bz)
	require.NoError(t, err)
	require.Equal(t, original.GlobalPaused, restored.GlobalPaused)
	require.Equal(t, original.Reason, restored.Reason)
}

func TestMarshalUnmarshalModuleCircuitState(t *testing.T) {
	original := types.ModuleCircuitState{
		ModuleName: "stablecoin",
		Status:     types.CircuitOpen,
		Reason:     "oracle failure",
	}

	// Marshal
	bz, err := types.MarshalModuleCircuitState(original)
	require.NoError(t, err)
	require.NotEmpty(t, bz)

	// Unmarshal
	restored, err := types.UnmarshalModuleCircuitState(bz)
	require.NoError(t, err)
	require.Equal(t, original.ModuleName, restored.ModuleName)
	require.Equal(t, original.Status, restored.Status)
	require.Equal(t, original.Reason, restored.Reason)
}

func TestMarshalUnmarshalRateLimitState(t *testing.T) {
	now := time.Now()
	original := types.RateLimitState{
		ConfigName:   "test_limit",
		RequestCount: 50,
		WindowStart:  now,
	}

	// Marshal
	bz, err := types.MarshalRateLimitState(original)
	require.NoError(t, err)
	require.NotEmpty(t, bz)

	// Unmarshal
	restored, err := types.UnmarshalRateLimitState(bz)
	require.NoError(t, err)
	require.Equal(t, original.ConfigName, restored.ConfigName)
	require.Equal(t, original.RequestCount, restored.RequestCount)
}

func TestMsgPauseSystem_GetSigners(t *testing.T) {
	authority := sdk.AccAddress("authority___________")
	msg := types.MsgPauseSystem{
		Authority: authority.String(),
	}
	signers := msg.GetSigners()
	require.Len(t, signers, 1)
	require.Equal(t, authority, signers[0])
}

func TestNewMsgPauseSystem(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	reason := "Emergency maintenance"
	duration := int64(3600)

	msg := types.NewMsgPauseSystem(authority, reason, duration)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, reason, msg.Reason)
	require.Equal(t, duration, msg.DurationSeconds)
}

func TestNewMsgTripCircuit(t *testing.T) {
	authority := sdk.AccAddress("authority___________").String()
	moduleName := "stablecoin"
	reason := "Oracle failure"
	disableMessages := []string{"/stateset.stablecoin.Msg/MintStablecoin"}

	msg := types.NewMsgTripCircuit(authority, moduleName, reason, disableMessages)

	require.Equal(t, authority, msg.Authority)
	require.Equal(t, moduleName, msg.ModuleName)
	require.Equal(t, reason, msg.Reason)
	require.Equal(t, disableMessages, msg.DisableMessages)
}
