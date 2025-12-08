package types

import (
	"encoding/json"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
)

// CircuitState represents the global circuit breaker state
type CircuitState struct {
	// GlobalPaused indicates if all operations are paused
	GlobalPaused bool `json:"global_paused"`
	// PausedAt is the timestamp when the circuit was paused
	PausedAt time.Time `json:"paused_at,omitempty"`
	// PausedBy is the authority that paused the circuit
	PausedBy string `json:"paused_by,omitempty"`
	// Reason for the pause
	Reason string `json:"reason,omitempty"`
	// AutoResumeAt is the optional time for automatic resume
	AutoResumeAt time.Time `json:"auto_resume_at,omitempty"`
}

// ModuleCircuitState represents circuit breaker state for a specific module
type ModuleCircuitState struct {
	// ModuleName is the name of the module
	ModuleName string `json:"module_name"`
	// Status is the current circuit status
	Status CircuitStatus `json:"status"`
	// TrippedAt is when the circuit was tripped
	TrippedAt time.Time `json:"tripped_at,omitempty"`
	// TrippedBy is who/what tripped the circuit
	TrippedBy string `json:"tripped_by,omitempty"`
	// Reason for the trip
	Reason string `json:"reason,omitempty"`
	// FailureCount tracks consecutive failures
	FailureCount uint64 `json:"failure_count"`
	// FailureThreshold is the number of failures to trip
	FailureThreshold uint64 `json:"failure_threshold"`
	// RecoveryTime is when the circuit can attempt half-open
	RecoveryTime time.Time `json:"recovery_time,omitempty"`
	// DisabledMessages are specific message types that are disabled
	DisabledMessages []string `json:"disabled_messages,omitempty"`
}

// RateLimitConfig defines rate limiting configuration
type RateLimitConfig struct {
	// Name is a unique identifier for this rate limit
	Name string `json:"name"`
	// MaxRequests is the maximum number of requests allowed
	MaxRequests uint64 `json:"max_requests"`
	// WindowSeconds is the time window in seconds
	WindowSeconds int64 `json:"window_seconds"`
	// PerAddress if true, limits are per-address; if false, global
	PerAddress bool `json:"per_address"`
	// Enabled indicates if this rate limit is active
	Enabled bool `json:"enabled"`
	// MessageTypes that this rate limit applies to (empty = all)
	MessageTypes []string `json:"message_types,omitempty"`
}

// RateLimitState tracks current rate limit usage
type RateLimitState struct {
	// ConfigName references the RateLimitConfig
	ConfigName string `json:"config_name"`
	// Address is the address being tracked (empty for global limits)
	Address string `json:"address,omitempty"`
	// RequestCount is the current request count in the window
	RequestCount uint64 `json:"request_count"`
	// WindowStart is when the current window started
	WindowStart time.Time `json:"window_start"`
}

// Params defines the circuit breaker module parameters
type Params struct {
	// Authorities are the addresses that can control circuit breakers
	Authorities []string `json:"authorities"`
	// DefaultFailureThreshold is the default number of failures to trip a circuit
	DefaultFailureThreshold uint64 `json:"default_failure_threshold"`
	// DefaultRecoveryPeriod is the default recovery period in seconds
	DefaultRecoveryPeriod int64 `json:"default_recovery_period"`
	// MaxPauseDuration is the maximum duration for a pause in seconds
	MaxPauseDuration int64 `json:"max_pause_duration"`
	// RateLimits are the configured rate limits
	RateLimits []RateLimitConfig `json:"rate_limits"`
}

// DefaultParams returns default module parameters
func DefaultParams() Params {
	return Params{
		Authorities:             []string{},
		DefaultFailureThreshold: 5,
		DefaultRecoveryPeriod:   300, // 5 minutes
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
				MessageTypes:  []string{"/stateset.stablecoin.v1.MsgMintStablecoin"},
			},
			{
				Name:          "large_settlement",
				MaxRequests:   5,
				WindowSeconds: 300,
				PerAddress:    true,
				Enabled:       true,
				MessageTypes:  []string{"/stateset.settlement.v1.MsgInstantTransfer"},
			},
		},
	}
}

// Validate validates the params
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

// OracleDeviationConfig defines price deviation thresholds
type OracleDeviationConfig struct {
	// Denom is the asset denom
	Denom string `json:"denom"`
	// MaxDeviationPercent is the maximum allowed deviation per update (in basis points, 100 = 1%)
	MaxDeviationBps uint64 `json:"max_deviation_bps"`
	// MaxDailyDeviationBps is the maximum daily deviation allowed
	MaxDailyDeviationBps uint64 `json:"max_daily_deviation_bps"`
	// StalenessThreshold is the number of seconds after which a price is stale
	StalenessThreshold int64 `json:"staleness_threshold"`
	// MinUpdateInterval is the minimum seconds between price updates
	MinUpdateInterval int64 `json:"min_update_interval"`
}

// DefaultOracleDeviationConfigs returns default oracle deviation configurations
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

// LiquidationSurgeProtection defines protection against liquidation cascades
type LiquidationSurgeProtection struct {
	// MaxLiquidationsPerBlock limits liquidations per block
	MaxLiquidationsPerBlock uint64 `json:"max_liquidations_per_block"`
	// MaxLiquidationValue limits total value liquidated per block (in base units)
	MaxLiquidationValue sdkmath.Int `json:"max_liquidation_value"`
	// CooldownBlocks is the number of blocks to wait after hitting limits
	CooldownBlocks uint64 `json:"cooldown_blocks"`
	// CurrentBlockLiquidations tracks current block liquidations
	CurrentBlockLiquidations uint64 `json:"current_block_liquidations"`
	// CurrentBlockValue tracks current block liquidation value
	CurrentBlockValue sdkmath.Int `json:"current_block_value"`
	// LastResetHeight is the last block height when counters were reset
	LastResetHeight int64 `json:"last_reset_height"`
}

// DefaultLiquidationSurgeProtection returns default surge protection settings
func DefaultLiquidationSurgeProtection() LiquidationSurgeProtection {
	return LiquidationSurgeProtection{
		MaxLiquidationsPerBlock: 10,
		MaxLiquidationValue:     sdkmath.NewInt(1_000_000_000_000), // 1M units
		CooldownBlocks:          5,
		CurrentBlockLiquidations: 0,
		CurrentBlockValue:        sdkmath.ZeroInt(),
		LastResetHeight:          0,
	}
}

// GenesisState defines the circuit module's genesis state
type GenesisState struct {
	Params                  Params                     `json:"params"`
	CircuitState            CircuitState               `json:"circuit_state"`
	ModuleCircuits          []ModuleCircuitState       `json:"module_circuits"`
	RateLimitStates         []RateLimitState           `json:"rate_limit_states"`
	OracleDeviationConfigs  []OracleDeviationConfig    `json:"oracle_deviation_configs"`
	LiquidationProtection   LiquidationSurgeProtection `json:"liquidation_protection"`
}

// Proto message methods for GenesisState
func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return "" }
func (m *GenesisState) ProtoMessage()  {}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:                  DefaultParams(),
		CircuitState:            CircuitState{GlobalPaused: false},
		ModuleCircuits:          []ModuleCircuitState{},
		RateLimitStates:         []RateLimitState{},
		OracleDeviationConfigs:  DefaultOracleDeviationConfigs(),
		LiquidationProtection:   DefaultLiquidationSurgeProtection(),
	}
}

// Validate performs basic validation of the genesis state
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}

// Marshal helpers
func (cs CircuitState) Marshal() ([]byte, error) {
	return json.Marshal(cs)
}

func UnmarshalCircuitState(bz []byte) (CircuitState, error) {
	var cs CircuitState
	err := json.Unmarshal(bz, &cs)
	return cs, err
}

func (mcs ModuleCircuitState) Marshal() ([]byte, error) {
	return json.Marshal(mcs)
}

func UnmarshalModuleCircuitState(bz []byte) (ModuleCircuitState, error) {
	var mcs ModuleCircuitState
	err := json.Unmarshal(bz, &mcs)
	return mcs, err
}

func (rls RateLimitState) Marshal() ([]byte, error) {
	return json.Marshal(rls)
}

func UnmarshalRateLimitState(bz []byte) (RateLimitState, error) {
	var rls RateLimitState
	err := json.Unmarshal(bz, &rls)
	return rls, err
}
