package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Params defines the parameters for the settlement module
type Params struct {
	// Default fee rate in basis points (e.g., 50 = 0.50%)
	DefaultFeeRateBps uint32 `json:"default_fee_rate_bps" yaml:"default_fee_rate_bps"`

	// Fee collector address
	FeeCollector string `json:"fee_collector" yaml:"fee_collector"`

	// Minimum settlement amount
	MinSettlementAmount sdk.Coin `json:"min_settlement_amount" yaml:"min_settlement_amount"`

	// Maximum settlement amount
	MaxSettlementAmount sdk.Coin `json:"max_settlement_amount" yaml:"max_settlement_amount"`

	// Default escrow expiration (in seconds)
	DefaultEscrowExpiration int64 `json:"default_escrow_expiration" yaml:"default_escrow_expiration"`

	// Maximum escrow expiration (in seconds)
	MaxEscrowExpiration int64 `json:"max_escrow_expiration" yaml:"max_escrow_expiration"`

	// Minimum channel expiration (in blocks)
	MinChannelExpiration int64 `json:"min_channel_expiration" yaml:"min_channel_expiration"`

	// Maximum channel expiration (in blocks)
	MaxChannelExpiration int64 `json:"max_channel_expiration" yaml:"max_channel_expiration"`

	// Maximum batch size
	MaxBatchSize uint32 `json:"max_batch_size" yaml:"max_batch_size"`

	// Maximum query limit for paginated queries
	MaxQueryLimit uint32 `json:"max_query_limit" yaml:"max_query_limit"`

	// Whether instant transfers are enabled
	InstantTransfersEnabled bool `json:"instant_transfers_enabled" yaml:"instant_transfers_enabled"`

	// Whether escrow settlements are enabled
	EscrowEnabled bool `json:"escrow_enabled" yaml:"escrow_enabled"`

	// Whether payment channels are enabled
	ChannelsEnabled bool `json:"channels_enabled" yaml:"channels_enabled"`
}

// ProtoMessage implements proto.Message
func (p *Params) ProtoMessage() {}

// Reset implements proto.Message
func (p *Params) Reset() { *p = Params{} }

// String implements proto.Message
func (p Params) String() string {
	return fmt.Sprintf("Params{feeRateBps:%d, feeCollector:%s, minSettlement:%s, maxSettlement:%s}",
		p.DefaultFeeRateBps, p.FeeCollector, p.MinSettlementAmount.String(), p.MaxSettlementAmount.String())
}

// DefaultParams returns default settlement parameters
func DefaultParams() Params {
	return Params{
		DefaultFeeRateBps:       50, // 0.50%
		FeeCollector:           "",  // Will use module account
		MinSettlementAmount:    sdk.NewCoin(StablecoinDenom, sdkmath.NewInt(1000)),       // 0.001 ssusd (assuming 6 decimals)
		MaxSettlementAmount:    sdk.NewCoin(StablecoinDenom, sdkmath.NewInt(1000000000000)), // 1M ssusd
		DefaultEscrowExpiration: 86400 * 7,  // 7 days
		MaxEscrowExpiration:     86400 * 30, // 30 days
		MinChannelExpiration:    100,        // ~10 minutes at 6s blocks
		MaxChannelExpiration:    518400,     // ~30 days at 6s blocks
		MaxBatchSize:            100,
		MaxQueryLimit:           100,
		InstantTransfersEnabled: true,
		EscrowEnabled:          true,
		ChannelsEnabled:        true,
	}
}

// Validate validates the params
func (p Params) Validate() error {
	if p.DefaultFeeRateBps > 10000 {
		return fmt.Errorf("default fee rate must be <= 10000 bps")
	}
	if !p.MinSettlementAmount.IsValid() {
		return fmt.Errorf("invalid min settlement amount")
	}
	if !p.MaxSettlementAmount.IsValid() {
		return fmt.Errorf("invalid max settlement amount")
	}
	if p.MinSettlementAmount.Amount.GT(p.MaxSettlementAmount.Amount) {
		return fmt.Errorf("min settlement must be <= max settlement")
	}
	if p.DefaultEscrowExpiration <= 0 {
		return fmt.Errorf("default escrow expiration must be positive")
	}
	if p.MaxEscrowExpiration < p.DefaultEscrowExpiration {
		return fmt.Errorf("max escrow expiration must be >= default")
	}
	if p.MinChannelExpiration <= 0 {
		return fmt.Errorf("min channel expiration must be positive")
	}
	if p.MaxChannelExpiration < p.MinChannelExpiration {
		return fmt.Errorf("max channel expiration must be >= min")
	}
	if p.MaxBatchSize == 0 {
		return fmt.Errorf("max batch size must be positive")
	}
	if p.MaxQueryLimit == 0 {
		return fmt.Errorf("max query limit must be positive")
	}
	return nil
}

// GenesisState defines the settlement module's genesis state
type GenesisState struct {
	Params           Params            `json:"params" yaml:"params"`
	Settlements      []Settlement      `json:"settlements" yaml:"settlements"`
	Batches          []BatchSettlement `json:"batches" yaml:"batches"`
	Channels         []PaymentChannel  `json:"channels" yaml:"channels"`
	Merchants        []MerchantConfig  `json:"merchants" yaml:"merchants"`
	NextSettlementId uint64            `json:"next_settlement_id" yaml:"next_settlement_id"`
	NextBatchId      uint64            `json:"next_batch_id" yaml:"next_batch_id"`
	NextChannelId    uint64            `json:"next_channel_id" yaml:"next_channel_id"`
}

// ProtoMessage implements proto.Message
func (gs *GenesisState) ProtoMessage() {}

// Reset implements proto.Message
func (gs *GenesisState) Reset() { *gs = GenesisState{} }

// String implements proto.Message
func (gs *GenesisState) String() string {
	return fmt.Sprintf("GenesisState{settlements:%d, batches:%d, channels:%d, merchants:%d}",
		len(gs.Settlements), len(gs.Batches), len(gs.Channels), len(gs.Merchants))
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		Settlements:      []Settlement{},
		Batches:          []BatchSettlement{},
		Channels:         []PaymentChannel{},
		Merchants:        []MerchantConfig{},
		NextSettlementId: 1,
		NextBatchId:      1,
		NextChannelId:    1,
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	settlementIds := make(map[uint64]bool)
	for _, s := range gs.Settlements {
		if settlementIds[s.Id] {
			return fmt.Errorf("duplicate settlement id: %d", s.Id)
		}
		settlementIds[s.Id] = true
	}

	batchIds := make(map[uint64]bool)
	for _, b := range gs.Batches {
		if batchIds[b.Id] {
			return fmt.Errorf("duplicate batch id: %d", b.Id)
		}
		batchIds[b.Id] = true
	}

	channelIds := make(map[uint64]bool)
	for _, c := range gs.Channels {
		if channelIds[c.Id] {
			return fmt.Errorf("duplicate channel id: %d", c.Id)
		}
		channelIds[c.Id] = true
	}

	merchantAddrs := make(map[string]bool)
	for _, m := range gs.Merchants {
		if merchantAddrs[m.Address] {
			return fmt.Errorf("duplicate merchant address: %s", m.Address)
		}
		merchantAddrs[m.Address] = true
	}

	return nil
}
