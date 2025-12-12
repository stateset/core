package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultParams returns default settlement parameters.
func DefaultParams() Params {
	return Params{
		DefaultFeeRateBps:       50,                                                              // 0.50%
		FeeCollector:            "",                                                              // Will use module account
		MinSettlementAmount:     sdk.NewCoin(StablecoinDenom, sdkmath.NewInt(1000)),              // 0.001 ssusd (assuming 6 decimals)
		MaxSettlementAmount:     sdk.NewCoin(StablecoinDenom, sdkmath.NewInt(1_000_000_000_000)), // 1M ssusd
		DefaultEscrowExpiration: 86400 * 7,                                                       // 7 days
		MaxEscrowExpiration:     86400 * 30,                                                      // 30 days
		MinChannelExpiration:    100,                                                             // ~10 minutes at 6s blocks
		MaxChannelExpiration:    518400,                                                          // ~30 days at 6s blocks
		MaxBatchSize:            100,
		MaxQueryLimit:           100,
		InstantTransfersEnabled: true,
		EscrowEnabled:           true,
		ChannelsEnabled:         true,
	}
}

// Validate validates the params.
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

// DefaultGenesis returns the default genesis state.
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

// Validate performs basic genesis state validation.
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
