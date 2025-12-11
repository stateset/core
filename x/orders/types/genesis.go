package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	stablecointypes "github.com/stateset/core/x/stablecoin/types"
)

// Params defines the configurable parameters for the orders module.
type Params struct {
	// DefaultOrderExpiration is the default time until an unpaid order expires (in seconds).
	DefaultOrderExpiration int64 `json:"default_order_expiration"`
	// DefaultEscrowExpiration is the default escrow release time after delivery (in seconds).
	DefaultEscrowExpiration int64 `json:"default_escrow_expiration"`
	// DisputeWindow is the time window during which a customer can open a dispute (in seconds).
	DisputeWindow int64 `json:"dispute_window"`
	// MinOrderAmount is the minimum order amount allowed.
	MinOrderAmount sdk.Coin `json:"min_order_amount"`
	// MaxOrderAmount is the maximum order amount allowed.
	MaxOrderAmount sdk.Coin `json:"max_order_amount"`
	// DefaultFeeRateBps is the default fee rate in basis points (100 = 1%).
	DefaultFeeRateBps uint32 `json:"default_fee_rate_bps"`
	// StablecoinDenom is the accepted stablecoin denomination for payments.
	StablecoinDenom string `json:"stablecoin_denom"`
	// AutoCompleteAfterDelivery if true, automatically complete orders after delivery + window.
	AutoCompleteAfterDelivery bool `json:"auto_complete_after_delivery"`
	// AutoCompleteWindow is the time window after delivery before auto-completion (in seconds).
	AutoCompleteWindow int64 `json:"auto_complete_window"`
}

func (p *Params) Reset()         { *p = Params{} }
func (p *Params) String() string { return "Params" }
func (*Params) ProtoMessage()    {}

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
