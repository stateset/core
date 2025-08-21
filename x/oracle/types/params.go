package types

import (
	"fmt"
	"time"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinProviders           = []byte("MinProviders")
	KeyUpdateInterval         = []byte("UpdateInterval")
	KeyMaxPriceAge            = []byte("MaxPriceAge")
	KeyPriceDeviationThreshold = []byte("PriceDeviationThreshold")
	KeyRewardAmount           = []byte("RewardAmount")
	KeySlashAmount            = []byte("SlashAmount")
	KeyMinSubmissionCount     = []byte("MinSubmissionCount")
	KeyHistoryRetention       = []byte("HistoryRetention")
	KeyEmergencyThreshold     = []byte("EmergencyThreshold")
)

// Default parameter values
const (
	DefaultMinProviders           = uint32(3)
	DefaultUpdateInterval         = time.Minute * 5
	DefaultMaxPriceAge            = time.Minute * 10
	DefaultMinSubmissionCount     = uint32(10)
	DefaultHistoryRetention      = time.Hour * 24 * 7 // 7 days
)

var (
	DefaultPriceDeviationThreshold = sdk.NewDecWithPrec(5, 2)  // 5%
	DefaultRewardAmount            = sdk.NewInt(100)           // 100 tokens
	DefaultSlashAmount             = sdk.NewInt(1000)          // 1000 tokens
	DefaultEmergencyThreshold      = sdk.NewDecWithPrec(20, 2) // 20%
)

// Params defines the parameters for the oracle module
type Params struct {
	// MinProviders is the minimum number of oracle providers required for a valid price
	MinProviders uint32 `json:"min_providers"`
	
	// UpdateInterval is the expected interval between price updates
	UpdateInterval time.Duration `json:"update_interval"`
	
	// MaxPriceAge is the maximum age of a price before it's considered stale
	MaxPriceAge time.Duration `json:"max_price_age"`
	
	// PriceDeviationThreshold is the maximum allowed deviation from the median price
	PriceDeviationThreshold sdk.Dec `json:"price_deviation_threshold"`
	
	// RewardAmount is the amount rewarded for valid price submissions
	RewardAmount sdk.Int `json:"reward_amount"`
	
	// SlashAmount is the amount slashed for invalid submissions
	SlashAmount sdk.Int `json:"slash_amount"`
	
	// MinSubmissionCount is the minimum number of submissions required from each provider
	MinSubmissionCount uint32 `json:"min_submission_count"`
	
	// HistoryRetention is how long to keep price history
	HistoryRetention time.Duration `json:"history_retention"`
	
	// EmergencyThreshold is the deviation threshold that triggers emergency mode
	EmergencyThreshold sdk.Dec `json:"emergency_threshold"`
}

// NewParams creates a new Params instance
func NewParams(
	minProviders uint32,
	updateInterval time.Duration,
	maxPriceAge time.Duration,
	priceDeviationThreshold sdk.Dec,
	rewardAmount sdk.Int,
	slashAmount sdk.Int,
	minSubmissionCount uint32,
	historyRetention time.Duration,
	emergencyThreshold sdk.Dec,
) Params {
	return Params{
		MinProviders:            minProviders,
		UpdateInterval:          updateInterval,
		MaxPriceAge:             maxPriceAge,
		PriceDeviationThreshold: priceDeviationThreshold,
		RewardAmount:            rewardAmount,
		SlashAmount:             slashAmount,
		MinSubmissionCount:      minSubmissionCount,
		HistoryRetention:        historyRetention,
		EmergencyThreshold:      emergencyThreshold,
	}
}

// DefaultParams returns default oracle parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMinProviders,
		DefaultUpdateInterval,
		DefaultMaxPriceAge,
		DefaultPriceDeviationThreshold,
		DefaultRewardAmount,
		DefaultSlashAmount,
		DefaultMinSubmissionCount,
		DefaultHistoryRetention,
		DefaultEmergencyThreshold,
	)
}

// ParamSetPairs returns the parameter set pairs
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinProviders, &p.MinProviders, validateMinProviders),
		paramtypes.NewParamSetPair(KeyUpdateInterval, &p.UpdateInterval, validateUpdateInterval),
		paramtypes.NewParamSetPair(KeyMaxPriceAge, &p.MaxPriceAge, validateMaxPriceAge),
		paramtypes.NewParamSetPair(KeyPriceDeviationThreshold, &p.PriceDeviationThreshold, validatePriceDeviationThreshold),
		paramtypes.NewParamSetPair(KeyRewardAmount, &p.RewardAmount, validateRewardAmount),
		paramtypes.NewParamSetPair(KeySlashAmount, &p.SlashAmount, validateSlashAmount),
		paramtypes.NewParamSetPair(KeyMinSubmissionCount, &p.MinSubmissionCount, validateMinSubmissionCount),
		paramtypes.NewParamSetPair(KeyHistoryRetention, &p.HistoryRetention, validateHistoryRetention),
		paramtypes.NewParamSetPair(KeyEmergencyThreshold, &p.EmergencyThreshold, validateEmergencyThreshold),
	}
}

// Validate validates the parameter set
func (p Params) Validate() error {
	if err := validateMinProviders(p.MinProviders); err != nil {
		return err
	}
	if err := validateUpdateInterval(p.UpdateInterval); err != nil {
		return err
	}
	if err := validateMaxPriceAge(p.MaxPriceAge); err != nil {
		return err
	}
	if err := validatePriceDeviationThreshold(p.PriceDeviationThreshold); err != nil {
		return err
	}
	if err := validateRewardAmount(p.RewardAmount); err != nil {
		return err
	}
	if err := validateSlashAmount(p.SlashAmount); err != nil {
		return err
	}
	if err := validateMinSubmissionCount(p.MinSubmissionCount); err != nil {
		return err
	}
	if err := validateHistoryRetention(p.HistoryRetention); err != nil {
		return err
	}
	if err := validateEmergencyThreshold(p.EmergencyThreshold); err != nil {
		return err
	}
	return nil
}

func validateMinProviders(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == 0 {
		return fmt.Errorf("minimum providers must be greater than 0")
	}
	return nil
}

func validateUpdateInterval(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("update interval must be positive")
	}
	return nil
}

func validateMaxPriceAge(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("max price age must be positive")
	}
	return nil
}

func validatePriceDeviationThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("price deviation threshold must be between 0 and 1")
	}
	return nil
}

func validateRewardAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("reward amount cannot be negative")
	}
	return nil
}

func validateSlashAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() {
		return fmt.Errorf("slash amount cannot be negative")
	}
	return nil
}

func validateMinSubmissionCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v == 0 {
		return fmt.Errorf("minimum submission count must be greater than 0")
	}
	return nil
}

func validateHistoryRetention(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("history retention must be positive")
	}
	return nil
}

func validateEmergencyThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("emergency threshold must be between 0 and 1")
	}
	return nil
}

// ParamKeyTable returns the parameter key table
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}