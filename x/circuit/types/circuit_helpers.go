package types

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
)

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		Authorities:             []string{},
		DefaultFailureThreshold: 5,
		DefaultRecoveryPeriod:   300,   // 5 minutes
		MaxPauseDuration:        86400, // 24 hours
		RateLimits: []RateLimitConfig{
			{
				Name:          "global_tx",
				MaxRequests:   1000,
				WindowSeconds: 60,
				PerAddress:    false,
				Enabled:       true,
			},
			{
				Name:          "per_address_tx",
				MaxRequests:   100,
				WindowSeconds: 60,
				PerAddress:    true,
				Enabled:       true,
			},
			{
				Name:          "stablecoin_mint",
				MaxRequests:   10,
				WindowSeconds: 60,
				PerAddress:    true,
				Enabled:       true,
				MessageTypes:  []string{"/stateset.stablecoin.Msg/MintStablecoin"},
			},
			{
				Name:          "large_settlement",
				MaxRequests:   5,
				WindowSeconds: 300,
				PerAddress:    true,
				Enabled:       true,
				MessageTypes:  []string{"/stateset.settlement.Msg/InstantTransfer"},
			},
		},
	}
}

// Validate validates the params.
func (p Params) Validate() error {
	if p.DefaultFailureThreshold == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "failure threshold must be positive")
	}
	if p.DefaultRecoveryPeriod <= 0 {
		return errorsmod.Wrap(ErrInvalidParams, "recovery period must be positive")
	}
	if p.MaxPauseDuration <= 0 {
		return errorsmod.Wrap(ErrInvalidParams, "max pause duration must be positive")
	}
	return nil
}

// DefaultOracleDeviationConfigs returns default oracle deviation configurations.
func DefaultOracleDeviationConfigs() []OracleDeviationConfig {
	return []OracleDeviationConfig{
		{
			Denom:                "uatom",
			MaxDeviationBps:      500,  // 5% max per update
			MaxDailyDeviationBps: 2000, // 20% max per day
			StalenessThreshold:   3600, // 1 hour
			MinUpdateInterval:    60,   // 1 minute
		},
		{
			Denom:                "uusdc",
			MaxDeviationBps:      100,  // 1% max (stablecoin)
			MaxDailyDeviationBps: 200,  // 2% max per day
			StalenessThreshold:   3600, // 1 hour
			MinUpdateInterval:    60,   // 1 minute
		},
	}
}

// DefaultLiquidationSurgeProtection returns default surge protection settings.
func DefaultLiquidationSurgeProtection() LiquidationSurgeProtection {
	return LiquidationSurgeProtection{
		MaxLiquidationsPerBlock:  10,
		MaxLiquidationValue:      sdkmath.NewInt(1_000_000_000_000), // 1M units
		CooldownBlocks:           5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          0,
	}
}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                 DefaultParams(),
		CircuitState:           CircuitState{GlobalPaused: false},
		ModuleCircuits:         []ModuleCircuitState{},
		RateLimitStates:        []RateLimitState{},
		OracleDeviationConfigs: DefaultOracleDeviationConfigs(),
		LiquidationProtection:  DefaultLiquidationSurgeProtection(),
	}
}

// Validate performs basic validation of the genesis state.
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// JSON marshal helpers (KV store serialization).
func MarshalCircuitState(cs CircuitState) ([]byte, error) {
	return json.Marshal(cs)
}

func UnmarshalCircuitState(bz []byte) (CircuitState, error) {
	var cs CircuitState
	err := json.Unmarshal(bz, &cs)
	return cs, err
}

func MarshalModuleCircuitState(mcs ModuleCircuitState) ([]byte, error) {
	return json.Marshal(mcs)
}

func UnmarshalModuleCircuitState(bz []byte) (ModuleCircuitState, error) {
	var mcs ModuleCircuitState
	err := json.Unmarshal(bz, &mcs)
	return mcs, err
}

func MarshalRateLimitState(rls RateLimitState) ([]byte, error) {
	return json.Marshal(rls)
}

func UnmarshalRateLimitState(bz []byte) (RateLimitState, error) {
	var rls RateLimitState
	err := json.Unmarshal(bz, &rls)
	return rls, err
}
