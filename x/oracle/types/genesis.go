package types

import errorsmod "cosmossdk.io/errors"

// DefaultAuthority identifies the default address allowed to update prices (governance module account).
const DefaultAuthority = ""

// GenesisState defines the oracle module genesis data.
type GenesisState struct {
	Authority     string           `json:"authority" yaml:"authority"`
	Params        OracleParams     `json:"params" yaml:"params"`
	Prices        []Price          `json:"prices" yaml:"prices"`
	Configs       []OracleConfig   `json:"configs" yaml:"configs"`
	Providers     []OracleProvider `json:"providers" yaml:"providers"`
	PendingPrices []PendingPrice   `json:"pending_prices" yaml:"pending_prices"`
}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Authority:     DefaultAuthority,
		Params:        DefaultOracleParams(),
		Prices:        []Price{},
		Configs:       []OracleConfig{},
		Providers:     []OracleProvider{},
		PendingPrices: []PendingPrice{},
	}
}

// Validate performs basic genesis state validation.
func (gs GenesisState) Validate() error {
	if gs.Params.PriceHistorySize == 0 {
		return errorsmod.Wrap(ErrInvalidParams, "price history size must be positive")
	}
	for _, price := range gs.Prices {
		if err := price.ValidateBasic(); err != nil {
			return err
		}
	}
	for _, cfg := range gs.Configs {
		if cfg.Denom == "" {
			return errorsmod.Wrap(ErrInvalidDenom, "config denom cannot be empty")
		}
		if cfg.RequiredConfirmations == 0 {
			return errorsmod.Wrap(ErrInvalidParams, "required confirmations must be >= 1")
		}
	}
	for _, p := range gs.Providers {
		if p.Address == "" {
			return errorsmod.Wrap(ErrInvalidAuthority, "provider address cannot be empty")
		}
	}
	for _, pp := range gs.PendingPrices {
		if pp.Denom == "" {
			return errorsmod.Wrap(ErrInvalidDenom, "pending price denom cannot be empty")
		}
		if !pp.Amount.IsPositive() {
			return errorsmod.Wrap(ErrInvalidPrice, "pending price must be positive")
		}
		if pp.Provider == "" {
			return errorsmod.Wrap(ErrInvalidAuthority, "pending provider cannot be empty")
		}
	}
	return nil
}
