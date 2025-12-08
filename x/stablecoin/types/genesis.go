package types

type GenesisState struct {
	Params      Params  `json:"params" yaml:"params"`
	NextVaultId uint64  `json:"next_vault_id" yaml:"next_vault_id"`
	Vaults      []Vault `json:"vaults" yaml:"vaults"`
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:      DefaultParams(),
		NextVaultId: 1,
		Vaults:      []Vault{},
	}
}

func (gs GenesisState) Validate() error {
	if err := gs.Params.ValidateBasic(); err != nil {
		return err
	}
	for _, vault := range gs.Vaults {
		if err := vault.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
