package types

// GenesisState defines module genesis configuration.
type GenesisState struct {
	Authority string    `json:"authority" yaml:"authority"`
	Profiles  []Profile `json:"profiles" yaml:"profiles"`
}

// DefaultGenesis returns default state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Authority: "",
		Profiles:  []Profile{},
	}
}

// Validate ensures the genesis state is valid.
func (gs GenesisState) Validate() error {
	for _, profile := range gs.Profiles {
		if err := profile.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
