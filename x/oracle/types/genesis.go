package types

// DefaultAuthority identifies the default address allowed to update prices (governance module account).
const DefaultAuthority = ""

// GenesisState defines the oracle module genesis data.
type GenesisState struct {
	Authority string  `json:"authority" yaml:"authority"`
	Prices    []Price `json:"prices" yaml:"prices"`
}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Authority: DefaultAuthority,
		Prices:    []Price{},
	}
}

// Validate performs basic genesis state validation.
func (gs GenesisState) Validate() error {
	for _, price := range gs.Prices {
		if err := price.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
