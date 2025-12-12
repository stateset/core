package types

import sdkmath "cosmossdk.io/math"

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:     DefaultParams(),
		BaseFee:    DefaultInitialBaseFee,
		FeeHistory: []FeeHistoryEntry{},
	}
}

// Validate performs basic genesis validation.
func (gs GenesisState) Validate() error {
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	if gs.BaseFee.IsNegative() {
		return ErrInvalidBaseFee
	}
	return nil
}

// DefaultInitialBaseFee is the base fee used at chain start.
var DefaultInitialBaseFee = DefaultMinBaseFee

// DefaultMinBaseFee is a sane lower bound for base fee.
var DefaultMinBaseFee = sdkmath.LegacyMustNewDecFromStr("0.025")
