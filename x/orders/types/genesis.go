package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Orders: []Order{},
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in orders
	orderIndexMap := make(map[string]struct{})

	for _, elem := range gs.Orders {
		index := string(OrderKey)
		if _, ok := orderIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for order")
		}
		orderIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}

// GenesisState defines the orders module's genesis state.
type GenesisState struct {
	Params Params  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Orders []Order `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders"`
}