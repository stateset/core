package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

// DefaultParams returns the default parameters for the orders module.
func DefaultParams() Params {
	return Params{
		DefaultOrderExpiration:    86400,                                                                       // 24 hours
		DefaultEscrowExpiration:   604800,                                                                      // 7 days
		DisputeWindow:             1209600,                                                                     // 14 days
		MinOrderAmount:            sdk.NewCoin(stablecointypes.StablecoinDenom, sdkmath.NewInt(100)),           // $0.0001
		MaxOrderAmount:            sdk.NewCoin(stablecointypes.StablecoinDenom, sdkmath.NewInt(1000000000000)), // $1M
		DefaultFeeRateBps:         100,                                                                         // 1%
		StablecoinDenom:           stablecointypes.StablecoinDenom,
		AutoCompleteAfterDelivery: true,
		AutoCompleteWindow:        259200, // 3 days after delivery
	}
}

// Validate performs basic validation of the params.
func (p Params) Validate() error {
	if p.DefaultOrderExpiration <= 0 {
		return ErrInvalidOrder
	}
	if p.DefaultEscrowExpiration <= 0 {
		return ErrInvalidOrder
	}
	if p.DisputeWindow <= 0 {
		return ErrInvalidOrder
	}
	if p.DefaultFeeRateBps > 10000 {
		return ErrInvalidAmount
	}
	return nil
}

// GenesisState defines the orders module's genesis state.
type GenesisState struct {
	Params        Params    `json:"params"`
	Orders        []Order   `json:"orders"`
	Disputes      []Dispute `json:"disputes"`
	NextOrderId   uint64    `json:"next_order_id"`
	NextDisputeId uint64    `json:"next_dispute_id"`
}

func (gs *GenesisState) Reset()         { *gs = GenesisState{} }
func (gs *GenesisState) String() string { return "GenesisState" }
func (*GenesisState) ProtoMessage()     {}

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:        DefaultParams(),
		Orders:        []Order{},
		Disputes:      []Dispute{},
		NextOrderId:   1,
		NextDisputeId: 1,
	}
}

// Validate performs basic validation of the genesis state.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	return nil
}
