package types

import (
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

// Parameter store keys
var (
	KeyMintDenom                = []byte("MintDenom")
	KeyMaxSupply               = []byte("MaxSupply")
	KeyInitialSupply           = []byte("InitialSupply")
	KeyStakingRewardsRate      = []byte("StakingRewardsRate")
	KeyMinStakingAmount        = []byte("MinStakingAmount")
	KeyUnstakingPeriod         = []byte("UnstakingPeriod")
	KeySlashFractionDoubleSign = []byte("SlashFractionDoubleSign")
	KeySlashFractionDowntime   = []byte("SlashFractionDowntime")
	KeyGovernanceVotingPeriod  = []byte("GovernanceVotingPeriod")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Params defines the parameters for the XSS module
type Params struct {
	MintDenom                string        `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty" yaml:"mint_denom"`
	MaxSupply               math.Int      `protobuf:"bytes,2,opt,name=max_supply,json=maxSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_supply" yaml:"max_supply"`
	InitialSupply           math.Int      `protobuf:"bytes,3,opt,name=initial_supply,json=initialSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"initial_supply" yaml:"initial_supply"`
	StakingRewardsRate      math.LegacyDec `protobuf:"bytes,4,opt,name=staking_rewards_rate,json=stakingRewardsRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking_rewards_rate" yaml:"staking_rewards_rate"`
	MinStakingAmount        math.Int      `protobuf:"bytes,5,opt,name=min_staking_amount,json=minStakingAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_staking_amount" yaml:"min_staking_amount"`
	UnstakingPeriod         time.Duration `protobuf:"bytes,6,opt,name=unstaking_period,json=unstakingPeriod,proto3,stdduration" json:"unstaking_period" yaml:"unstaking_period"`
	SlashFractionDoubleSign math.LegacyDec `protobuf:"bytes,7,opt,name=slash_fraction_double_sign,json=slashFractionDoubleSign,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_double_sign" yaml:"slash_fraction_double_sign"`
	SlashFractionDowntime   math.LegacyDec `protobuf:"bytes,8,opt,name=slash_fraction_downtime,json=slashFractionDowntime,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction_downtime" yaml:"slash_fraction_downtime"`
	GovernanceVotingPeriod  time.Duration `protobuf:"bytes,9,opt,name=governance_voting_period,json=governanceVotingPeriod,proto3,stdduration" json:"governance_voting_period" yaml:"governance_voting_period"`
}

// NewParams creates a new Params instance
func NewParams(
	mintDenom string,
	maxSupply, initialSupply, minStakingAmount math.Int,
	stakingRewardsRate, slashFractionDoubleSign, slashFractionDowntime math.LegacyDec,
	unstakingPeriod, governanceVotingPeriod time.Duration,
) Params {
	return Params{
		MintDenom:                mintDenom,
		MaxSupply:               maxSupply,
		InitialSupply:           initialSupply,
		StakingRewardsRate:      stakingRewardsRate,
		MinStakingAmount:        minStakingAmount,
		UnstakingPeriod:         unstakingPeriod,
		SlashFractionDoubleSign: slashFractionDoubleSign,
		SlashFractionDowntime:   slashFractionDowntime,
		GovernanceVotingPeriod:  governanceVotingPeriod,
	}
}

// DefaultParams returns the default parameters for the XSS module
func DefaultParams() Params {
	return NewParams(
		XSSDenom,                                          // mint_denom
		math.NewInt(1_000_000_000_000_000),               // max_supply: 1 billion XSS (with 6 decimals)
		math.NewInt(100_000_000_000_000),                 // initial_supply: 100 million XSS
		math.NewInt(1_000_000),                           // min_staking_amount: 1 XSS
		math.LegacyNewDecWithPrec(8, 2),                  // staking_rewards_rate: 8% annual
		math.LegacyNewDecWithPrec(5, 2),                  // slash_fraction_double_sign: 5%
		math.LegacyNewDecWithPrec(1, 2),                  // slash_fraction_downtime: 1%
		time.Hour*24*21,                                  // unstaking_period: 21 days
		time.Hour*24*14,                                  // governance_voting_period: 14 days
	)
}

// ParamSetPairs implements the ParamSet interface and returns all key/value pairs
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMintDenom, &p.MintDenom, validateMintDenom),
		paramtypes.NewParamSetPair(KeyMaxSupply, &p.MaxSupply, validateMaxSupply),
		paramtypes.NewParamSetPair(KeyInitialSupply, &p.InitialSupply, validateInitialSupply),
		paramtypes.NewParamSetPair(KeyStakingRewardsRate, &p.StakingRewardsRate, validateStakingRewardsRate),
		paramtypes.NewParamSetPair(KeyMinStakingAmount, &p.MinStakingAmount, validateMinStakingAmount),
		paramtypes.NewParamSetPair(KeyUnstakingPeriod, &p.UnstakingPeriod, validateUnstakingPeriod),
		paramtypes.NewParamSetPair(KeySlashFractionDoubleSign, &p.SlashFractionDoubleSign, validateSlashFractionDoubleSign),
		paramtypes.NewParamSetPair(KeySlashFractionDowntime, &p.SlashFractionDowntime, validateSlashFractionDowntime),
		paramtypes.NewParamSetPair(KeyGovernanceVotingPeriod, &p.GovernanceVotingPeriod, validateGovernanceVotingPeriod),
	}
}

// Validate validates all parameters
func (p Params) Validate() error {
	if err := validateMintDenom(p.MintDenom); err != nil {
		return err
	}
	if err := validateMaxSupply(p.MaxSupply); err != nil {
		return err
	}
	if err := validateInitialSupply(p.InitialSupply); err != nil {
		return err
	}
	if err := validateStakingRewardsRate(p.StakingRewardsRate); err != nil {
		return err
	}
	if err := validateMinStakingAmount(p.MinStakingAmount); err != nil {
		return err
	}
	if err := validateUnstakingPeriod(p.UnstakingPeriod); err != nil {
		return err
	}
	if err := validateSlashFractionDoubleSign(p.SlashFractionDoubleSign); err != nil {
		return err
	}
	if err := validateSlashFractionDowntime(p.SlashFractionDowntime); err != nil {
		return err
	}
	if err := validateGovernanceVotingPeriod(p.GovernanceVotingPeriod); err != nil {
		return err
	}
	
	// Additional validation: initial supply should not exceed max supply
	if p.InitialSupply.GT(p.MaxSupply) {
		return fmt.Errorf("initial supply (%s) cannot exceed max supply (%s)", p.InitialSupply, p.MaxSupply)
	}
	
	return nil
}

// String implements the Stringer interface
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMintDenom validates the mint denomination
func validateMintDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return fmt.Errorf("mint denom cannot be blank")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return fmt.Errorf("invalid mint denom: %w", err)
	}

	return nil
}

// validateMaxSupply validates the maximum supply
func validateMaxSupply(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() || v.IsZero() {
		return fmt.Errorf("max supply must be positive: %s", v)
	}

	return nil
}

// validateInitialSupply validates the initial supply
func validateInitialSupply(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() || v.IsZero() {
		return fmt.Errorf("initial supply must be positive: %s", v)
	}

	return nil
}

// validateStakingRewardsRate validates the staking rewards rate
func validateStakingRewardsRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("staking rewards rate must be non-negative: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("staking rewards rate cannot exceed 100%%: %s", v)
	}

	return nil
}

// validateMinStakingAmount validates the minimum staking amount
func validateMinStakingAmount(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("min staking amount must be non-negative: %s", v)
	}

	return nil
}

// validateUnstakingPeriod validates the unstaking period
func validateUnstakingPeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("unstaking period must be positive: %s", v)
	}

	// Maximum unstaking period of 180 days
	if v > time.Hour*24*180 {
		return fmt.Errorf("unstaking period cannot exceed 180 days: %s", v)
	}

	return nil
}

// validateSlashFractionDoubleSign validates the slash fraction for double signing
func validateSlashFractionDoubleSign(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("slash fraction double sign must be non-negative: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("slash fraction double sign cannot exceed 100%%: %s", v)
	}

	return nil
}

// validateSlashFractionDowntime validates the slash fraction for downtime
func validateSlashFractionDowntime(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("slash fraction downtime must be non-negative: %s", v)
	}

	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("slash fraction downtime cannot exceed 100%%: %s", v)
	}

	return nil
}

// validateGovernanceVotingPeriod validates the governance voting period
func validateGovernanceVotingPeriod(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("governance voting period must be positive: %s", v)
	}

	// Maximum voting period of 30 days
	if v > time.Hour*24*30 {
		return fmt.Errorf("governance voting period cannot exceed 30 days: %s", v)
	}

	return nil
}