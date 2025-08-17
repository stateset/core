package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMinCollateralRatio    = []byte("MinCollateralRatio")
	KeyLiquidationThreshold  = []byte("LiquidationThreshold")
	KeyLiquidationPenalty    = []byte("LiquidationPenalty")
	KeyStabilityFee          = []byte("StabilityFee")
	KeyMintingFee            = []byte("MintingFee")
	KeyRedemptionFee         = []byte("RedemptionFee")
	KeyMaxDebtPerUser        = []byte("MaxDebtPerUser")
	KeyGlobalDebtCeiling     = []byte("GlobalDebtCeiling")
	KeyOracleUpdateFrequency = []byte("OracleUpdateFrequency")
	KeyAuctionDuration       = []byte("AuctionDuration")
	KeyAuctionPriceDecay     = []byte("AuctionPriceDecay")
	KeyMinAgentReputation    = []byte("MinAgentReputation")
	KeyAgentFeeDiscount      = []byte("AgentFeeDiscount")
	KeyEmergencyShutdown     = []byte("EmergencyShutdown")
	KeyAcceptedCollateral    = []byte("AcceptedCollateral")
	KeyOracleWhitelist       = []byte("OracleWhitelist")
)

var (
	DefaultMinCollateralRatio    = sdk.NewDecWithPrec(150, 2) // 1.5 or 150%
	DefaultLiquidationThreshold  = sdk.NewDecWithPrec(130, 2) // 1.3 or 130%
	DefaultLiquidationPenalty    = sdk.NewDecWithPrec(10, 2)  // 0.1 or 10%
	DefaultStabilityFee          = sdk.NewDecWithPrec(2, 2)   // 0.02 or 2% annual
	DefaultMintingFee            = sdk.NewDecWithPrec(1, 3)   // 0.001 or 0.1%
	DefaultRedemptionFee         = sdk.NewDecWithPrec(1, 3)   // 0.001 or 0.1%
	DefaultMaxDebtPerUser        = sdk.NewInt(10000000000000) // 10M ssUSD
	DefaultGlobalDebtCeiling     = sdk.NewInt(1000000000000000) // 1B ssUSD
	DefaultOracleUpdateFrequency = uint64(60)                 // 60 seconds
	DefaultAuctionDuration       = uint64(21600)              // 6 hours
	DefaultAuctionPriceDecay     = sdk.NewDecWithPrec(1, 4)   // 0.0001 per second
	DefaultMinAgentReputation    = sdk.NewDec(50)             // Minimum reputation score
	DefaultAgentFeeDiscount      = sdk.NewDecWithPrec(50, 2)  // 50% discount
	DefaultEmergencyShutdown     = false
	DefaultAcceptedCollateral    = []string{"stst"}
	DefaultOracleWhitelist       = []string{}
)

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(
	minCollateralRatio sdk.Dec,
	liquidationThreshold sdk.Dec,
	liquidationPenalty sdk.Dec,
	stabilityFee sdk.Dec,
	mintingFee sdk.Dec,
	redemptionFee sdk.Dec,
	maxDebtPerUser sdk.Int,
	globalDebtCeiling sdk.Int,
	oracleUpdateFrequency uint64,
	auctionDuration uint64,
	auctionPriceDecay sdk.Dec,
	minAgentReputation sdk.Dec,
	agentFeeDiscount sdk.Dec,
	emergencyShutdown bool,
	acceptedCollateral []string,
	oracleWhitelist []string,
) Params {
	return Params{
		MinCollateralRatio:    minCollateralRatio,
		LiquidationThreshold:  liquidationThreshold,
		LiquidationPenalty:    liquidationPenalty,
		StabilityFee:          stabilityFee,
		MintingFee:            mintingFee,
		RedemptionFee:         redemptionFee,
		MaxDebtPerUser:        maxDebtPerUser,
		GlobalDebtCeiling:     globalDebtCeiling,
		OracleUpdateFrequency: oracleUpdateFrequency,
		AuctionDuration:       auctionDuration,
		AuctionPriceDecay:     auctionPriceDecay,
		MinAgentReputation:    minAgentReputation,
		AgentFeeDiscount:      agentFeeDiscount,
		EmergencyShutdown:     emergencyShutdown,
		AcceptedCollateral:    acceptedCollateral,
		OracleWhitelist:       oracleWhitelist,
	}
}

func DefaultParams() Params {
	return NewParams(
		DefaultMinCollateralRatio,
		DefaultLiquidationThreshold,
		DefaultLiquidationPenalty,
		DefaultStabilityFee,
		DefaultMintingFee,
		DefaultRedemptionFee,
		DefaultMaxDebtPerUser,
		DefaultGlobalDebtCeiling,
		DefaultOracleUpdateFrequency,
		DefaultAuctionDuration,
		DefaultAuctionPriceDecay,
		DefaultMinAgentReputation,
		DefaultAgentFeeDiscount,
		DefaultEmergencyShutdown,
		DefaultAcceptedCollateral,
		DefaultOracleWhitelist,
	)
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinCollateralRatio, &p.MinCollateralRatio, validateMinCollateralRatio),
		paramtypes.NewParamSetPair(KeyLiquidationThreshold, &p.LiquidationThreshold, validateLiquidationThreshold),
		paramtypes.NewParamSetPair(KeyLiquidationPenalty, &p.LiquidationPenalty, validateLiquidationPenalty),
		paramtypes.NewParamSetPair(KeyStabilityFee, &p.StabilityFee, validateStabilityFee),
		paramtypes.NewParamSetPair(KeyMintingFee, &p.MintingFee, validateMintingFee),
		paramtypes.NewParamSetPair(KeyRedemptionFee, &p.RedemptionFee, validateRedemptionFee),
		paramtypes.NewParamSetPair(KeyMaxDebtPerUser, &p.MaxDebtPerUser, validateMaxDebtPerUser),
		paramtypes.NewParamSetPair(KeyGlobalDebtCeiling, &p.GlobalDebtCeiling, validateGlobalDebtCeiling),
		paramtypes.NewParamSetPair(KeyOracleUpdateFrequency, &p.OracleUpdateFrequency, validateOracleUpdateFrequency),
		paramtypes.NewParamSetPair(KeyAuctionDuration, &p.AuctionDuration, validateAuctionDuration),
		paramtypes.NewParamSetPair(KeyAuctionPriceDecay, &p.AuctionPriceDecay, validateAuctionPriceDecay),
		paramtypes.NewParamSetPair(KeyMinAgentReputation, &p.MinAgentReputation, validateMinAgentReputation),
		paramtypes.NewParamSetPair(KeyAgentFeeDiscount, &p.AgentFeeDiscount, validateAgentFeeDiscount),
		paramtypes.NewParamSetPair(KeyEmergencyShutdown, &p.EmergencyShutdown, validateEmergencyShutdown),
		paramtypes.NewParamSetPair(KeyAcceptedCollateral, &p.AcceptedCollateral, validateAcceptedCollateral),
		paramtypes.NewParamSetPair(KeyOracleWhitelist, &p.OracleWhitelist, validateOracleWhitelist),
	}
}

func (p Params) Validate() error {
	if err := validateMinCollateralRatio(p.MinCollateralRatio); err != nil {
		return err
	}
	if err := validateLiquidationThreshold(p.LiquidationThreshold); err != nil {
		return err
	}
	if err := validateLiquidationPenalty(p.LiquidationPenalty); err != nil {
		return err
	}
	if err := validateStabilityFee(p.StabilityFee); err != nil {
		return err
	}
	if err := validateMintingFee(p.MintingFee); err != nil {
		return err
	}
	if err := validateRedemptionFee(p.RedemptionFee); err != nil {
		return err
	}
	if err := validateMaxDebtPerUser(p.MaxDebtPerUser); err != nil {
		return err
	}
	if err := validateGlobalDebtCeiling(p.GlobalDebtCeiling); err != nil {
		return err
	}
	if err := validateOracleUpdateFrequency(p.OracleUpdateFrequency); err != nil {
		return err
	}
	if err := validateAuctionDuration(p.AuctionDuration); err != nil {
		return err
	}
	if err := validateAuctionPriceDecay(p.AuctionPriceDecay); err != nil {
		return err
	}
	if err := validateMinAgentReputation(p.MinAgentReputation); err != nil {
		return err
	}
	if err := validateAgentFeeDiscount(p.AgentFeeDiscount); err != nil {
		return err
	}
	if err := validateEmergencyShutdown(p.EmergencyShutdown); err != nil {
		return err
	}
	if err := validateAcceptedCollateral(p.AcceptedCollateral); err != nil {
		return err
	}
	if err := validateOracleWhitelist(p.OracleWhitelist); err != nil {
		return err
	}
	return nil
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMinCollateralRatio(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LTE(sdk.OneDec()) {
		return fmt.Errorf("minimum collateral ratio must be greater than 1: %s", v)
	}

	return nil
}

func validateLiquidationThreshold(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LTE(sdk.OneDec()) {
		return fmt.Errorf("liquidation threshold must be greater than 1: %s", v)
	}

	return nil
}

func validateLiquidationPenalty(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("liquidation penalty must be between 0 and 1: %s", v)
	}

	return nil
}

func validateStabilityFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("stability fee must be between 0 and 1: %s", v)
	}

	return nil
}

func validateMintingFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("minting fee must be between 0 and 1: %s", v)
	}

	return nil
}

func validateRedemptionFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("redemption fee must be between 0 and 1: %s", v)
	}

	return nil
}

func validateMaxDebtPerUser(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("max debt per user cannot be negative: %s", v)
	}

	return nil
}

func validateGlobalDebtCeiling(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("global debt ceiling cannot be negative: %s", v)
	}

	return nil
}

func validateOracleUpdateFrequency(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("oracle update frequency must be greater than 0")
	}

	return nil
}

func validateAuctionDuration(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("auction duration must be greater than 0")
	}

	return nil
}

func validateAuctionPriceDecay(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("auction price decay must be between 0 and 1: %s", v)
	}

	return nil
}

func validateMinAgentReputation(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum agent reputation cannot be negative: %s", v)
	}

	return nil
}

func validateAgentFeeDiscount(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("agent fee discount must be between 0 and 1: %s", v)
	}

	return nil
}

func validateEmergencyShutdown(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateAcceptedCollateral(i interface{}) error {
	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("accepted collateral list cannot be empty")
	}

	return nil
}

func validateOracleWhitelist(i interface{}) error {
	_, ok := i.([]string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}