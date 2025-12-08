package types

type GenesisState struct {
	Authority string            `json:"authority" yaml:"authority"`
	NextID    uint64            `json:"next_id" yaml:"next_id"`
	Snapshots []ReserveSnapshot `json:"snapshots" yaml:"snapshots"`
}

func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Authority: "",
		NextID:    1,
		Snapshots: []ReserveSnapshot{},
	}
}

func (gs GenesisState) Validate() error {
	for _, snapshot := range gs.Snapshots {
		if err := snapshot.ValidateBasic(); err != nil {
			return err
		}
	}
	return nil
}
