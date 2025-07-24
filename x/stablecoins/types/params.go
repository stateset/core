package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxStablecoins        = []byte("MaxStablecoins")
	DefaultMaxStablecoins    = uint64(100)
	KeyMinInitialSupply      = []byte("MinInitialSupply")
	DefaultMinInitialSupply  = "1000000" // 1M in smallest unit
	KeyMaxInitialSupply      = []byte("MaxInitialSupply")
	DefaultMaxInitialSupply  = "1000000000000000" // 1 quadrillion in smallest unit
	KeyCreationFee           = []byte("CreationFee")
	DefaultCreationFee       = "100000000" // 100 STATE
	KeyMinReserveRatio       = []byte("MinReserveRatio")
	DefaultMinReserveRatio   = "1.00" // 100% collateralized minimum
	KeyMaxFeePercentage      = []byte("MaxFeePercentage")
	DefaultMaxFeePercentage  = "0.05" // 5% maximum fee
)

// NewParams creates a new Params instance
func NewParams(
	maxStablecoins uint64,
	minInitialSupply string,
	maxInitialSupply string,
	creationFee string,
	minReserveRatio string,
	maxFeePercentage string,
) Params {
	return Params{
		MaxStablecoins:   maxStablecoins,
		MinInitialSupply: minInitialSupply,
		MaxInitialSupply: maxInitialSupply,
		CreationFee:      creationFee,
		MinReserveRatio:  minReserveRatio,
		MaxFeePercentage: maxFeePercentage,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxStablecoins,
		DefaultMinInitialSupply,
		DefaultMaxInitialSupply,
		DefaultCreationFee,
		DefaultMinReserveRatio,
		DefaultMaxFeePercentage,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxStablecoins, &p.MaxStablecoins, validateMaxStablecoins),
		paramtypes.NewParamSetPair(KeyMinInitialSupply, &p.MinInitialSupply, validateMinInitialSupply),
		paramtypes.NewParamSetPair(KeyMaxInitialSupply, &p.MaxInitialSupply, validateMaxInitialSupply),
		paramtypes.NewParamSetPair(KeyCreationFee, &p.CreationFee, validateCreationFee),
		paramtypes.NewParamSetPair(KeyMinReserveRatio, &p.MinReserveRatio, validateMinReserveRatio),
		paramtypes.NewParamSetPair(KeyMaxFeePercentage, &p.MaxFeePercentage, validateMaxFeePercentage),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxStablecoins(p.MaxStablecoins); err != nil {
		return err
	}
	if err := validateMinInitialSupply(p.MinInitialSupply); err != nil {
		return err
	}
	if err := validateMaxInitialSupply(p.MaxInitialSupply); err != nil {
		return err
	}
	if err := validateCreationFee(p.CreationFee); err != nil {
		return err
	}
	if err := validateMinReserveRatio(p.MinReserveRatio); err != nil {
		return err
	}
	if err := validateMaxFeePercentage(p.MaxFeePercentage); err != nil {
		return err
	}
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateMaxStablecoins(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max stablecoins must be positive: %d", v)
	}

	return nil
}

func validateMinInitialSupply(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Additional validation can be added here to check if it's a valid number
	return nil
}

func validateMaxInitialSupply(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Additional validation can be added here to check if it's a valid number
	return nil
}

func validateCreationFee(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Additional validation can be added here to check if it's a valid number
	return nil
}

func validateMinReserveRatio(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Additional validation can be added here to check if it's a valid decimal
	return nil
}

func validateMaxFeePercentage(i interface{}) error {
	_, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// Additional validation can be added here to check if it's a valid percentage
	return nil
}

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}