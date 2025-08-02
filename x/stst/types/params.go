package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter values based on STST tokenomics
var (
	// DefaultTotalSupply is 1 billion STST tokens
	DefaultTotalSupply = sdk.NewIntFromUint64(1_000_000_000_000_000) // 1B with 6 decimal places

	// DefaultTokenDenom is the default token denomination
	DefaultTokenDenom = "stst"

	// DefaultStakingRewardsRate is the annual staking rewards rate (10%)
	DefaultStakingRewardsRate = sdk.NewDecWithPrec(10, 2) // 10%

	// DefaultFeeBurnRate is the percentage of transaction fees to burn (25%)
	DefaultFeeBurnRate = sdk.NewDecWithPrec(25, 2) // 25%

	// DefaultMinStakingAmount is the minimum amount required for staking (1000 STST)
	DefaultMinStakingAmount = sdk.NewIntFromUint64(1_000_000_000) // 1000 STST with 6 decimal places

	// DefaultGovernanceVotingPeriod is the duration for governance proposals (7 days)
	DefaultGovernanceVotingPeriod = time.Hour * 24 * 7 // 7 days

	// DefaultSlashingRate is the percentage of staked tokens slashed for malicious behavior (5%)
	DefaultSlashingRate = sdk.NewDecWithPrec(5, 2) // 5%

	// DefaultBurnFromSlashRate is the percentage of slashed tokens to burn (50%)
	DefaultBurnFromSlashRate = sdk.NewDecWithPrec(50, 2) // 50%
)

// Parameter store keys
var (
	KeyTotalSupply             = []byte("TotalSupply")
	KeyTokenDenom              = []byte("TokenDenom")
	KeyStakingRewardsRate      = []byte("StakingRewardsRate")
	KeyFeeBurnRate             = []byte("FeeBurnRate")
	KeyMinStakingAmount        = []byte("MinStakingAmount")
	KeyGovernanceVotingPeriod  = []byte("GovernanceVotingPeriod")
	KeySlashingRate            = []byte("SlashingRate")
	KeyBurnFromSlashRate       = []byte("BurnFromSlashRate")
)

// ParamKeyTable returns the parameter key table for the STST module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	totalSupply sdk.Int,
	tokenDenom string,
	stakingRewardsRate sdk.Dec,
	feeBurnRate sdk.Dec,
	minStakingAmount sdk.Int,
	governanceVotingPeriod time.Duration,
	slashingRate sdk.Dec,
	burnFromSlashRate sdk.Dec,
) Params {
	return Params{
		TotalSupply:            totalSupply,
		TokenDenom:             tokenDenom,
		StakingRewardsRate:     stakingRewardsRate,
		FeeBurnRate:            feeBurnRate,
		MinStakingAmount:       minStakingAmount,
		GovernanceVotingPeriod: governanceVotingPeriod,
		SlashingRate:           slashingRate,
		BurnFromSlashRate:      burnFromSlashRate,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultTotalSupply,
		DefaultTokenDenom,
		DefaultStakingRewardsRate,
		DefaultFeeBurnRate,
		DefaultMinStakingAmount,
		DefaultGovernanceVotingPeriod,
		DefaultSlashingRate,
		DefaultBurnFromSlashRate,
	)
}

// ParamSetPairs implements the ParamSet interface
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyTotalSupply, &p.TotalSupply, validateTotalSupply),
		paramtypes.NewParamSetPair(KeyTokenDenom, &p.TokenDenom, validateTokenDenom),
		paramtypes.NewParamSetPair(KeyStakingRewardsRate, &p.StakingRewardsRate, validateStakingRewardsRate),
		paramtypes.NewParamSetPair(KeyFeeBurnRate, &p.FeeBurnRate, validateFeeBurnRate),
		paramtypes.NewParamSetPair(KeyMinStakingAmount, &p.MinStakingAmount, validateMinStakingAmount),
		paramtypes.NewParamSetPair(KeyGovernanceVotingPeriod, &p.GovernanceVotingPeriod, validateGovernanceVotingPeriod),
		paramtypes.NewParamSetPair(KeySlashingRate, &p.SlashingRate, validateSlashingRate),
		paramtypes.NewParamSetPair(KeyBurnFromSlashRate, &p.BurnFromSlashRate, validateBurnFromSlashRate),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateTotalSupply(p.TotalSupply); err != nil {
		return err
	}
	if err := validateTokenDenom(p.TokenDenom); err != nil {
		return err
	}
	if err := validateStakingRewardsRate(p.StakingRewardsRate); err != nil {
		return err
	}
	if err := validateFeeBurnRate(p.FeeBurnRate); err != nil {
		return err
	}
	if err := validateMinStakingAmount(p.MinStakingAmount); err != nil {
		return err
	}
	if err := validateGovernanceVotingPeriod(p.GovernanceVotingPeriod); err != nil {
		return err
	}
	if err := validateSlashingRate(p.SlashingRate); err != nil {
		return err
	}
	if err := validateBurnFromSlashRate(p.BurnFromSlashRate); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface
func (p Params) String() string {
	return fmt.Sprintf(`STST Params:
  Total Supply:              %s
  Token Denom:               %s
  Staking Rewards Rate:      %s
  Fee Burn Rate:             %s
  Min Staking Amount:        %s
  Governance Voting Period:  %s
  Slashing Rate:             %s
  Burn From Slash Rate:      %s`,
		p.TotalSupply,
		p.TokenDenom,
		p.StakingRewardsRate,
		p.FeeBurnRate,
		p.MinStakingAmount,
		p.GovernanceVotingPeriod,
		p.SlashingRate,
		p.BurnFromSlashRate,
	)
}

// Validation functions

func validateTotalSupply(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() || v.IsZero() {
		return fmt.Errorf("total supply must be positive: %s", v)
	}

	// Ensure total supply is exactly 1 billion tokens (with 6 decimal places)
	expectedSupply := sdk.NewIntFromUint64(1_000_000_000_000_000)
	if !v.Equal(expectedSupply) {
		return fmt.Errorf("total supply must be exactly 1 billion STST tokens: got %s, expected %s", v, expectedSupply)
	}

	return nil
}

func validateTokenDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("token denom cannot be empty")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return fmt.Errorf("invalid token denom: %w", err)
	}

	return nil
}

func validateStakingRewardsRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("staking rewards rate must be non-negative: %s", v)
	}

	// Maximum 100% rewards rate
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("staking rewards rate cannot exceed 100%%: %s", v)
	}

	return nil
}

func validateFeeBurnRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("fee burn rate must be non-negative: %s", v)
	}

	// Maximum 100% burn rate
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("fee burn rate cannot exceed 100%%: %s", v)
	}

	return nil
}

func validateMinStakingAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("minimum staking amount must be non-negative: %s", v)
	}

	return nil
}

func validateGovernanceVotingPeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("governance voting period must be positive: %s", v)
	}

	// Minimum 1 hour, maximum 30 days
	if v < time.Hour {
		return fmt.Errorf("governance voting period cannot be less than 1 hour: %s", v)
	}

	if v > time.Hour*24*30 {
		return fmt.Errorf("governance voting period cannot exceed 30 days: %s", v)
	}

	return nil
}

func validateSlashingRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("slashing rate must be non-negative: %s", v)
	}

	// Maximum 100% slashing rate
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("slashing rate cannot exceed 100%%: %s", v)
	}

	return nil
}

func validateBurnFromSlashRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("burn from slash rate must be non-negative: %s", v)
	}

	// Maximum 100% burn rate
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("burn from slash rate cannot exceed 100%%: %s", v)
	}

	return nil
}