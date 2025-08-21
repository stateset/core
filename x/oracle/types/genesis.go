package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:           DefaultParams(),
		OracleProviders:  []OracleProvider{},
		PriceFeeds:       []PriceFeed{},
		AggregatedPrices: []AggregatedPrice{},
		PriceHistories:   []PriceHistory{},
	}
}

// GenesisState defines the oracle module's genesis state
type GenesisState struct {
	Params           Params            `json:"params"`
	OracleProviders  []OracleProvider  `json:"oracle_providers"`
	PriceFeeds       []PriceFeed       `json:"price_feeds"`
	AggregatedPrices []AggregatedPrice `json:"aggregated_prices"`
	PriceHistories   []PriceHistory    `json:"price_histories"`
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	// Validate params
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}
	
	// Validate oracle providers
	providerMap := make(map[string]bool)
	for _, provider := range gs.OracleProviders {
		if err := provider.Validate(); err != nil {
			return fmt.Errorf("invalid oracle provider %s: %w", provider.Address, err)
		}
		if providerMap[provider.Address] {
			return fmt.Errorf("duplicate oracle provider: %s", provider.Address)
		}
		providerMap[provider.Address] = true
	}
	
	// Validate price feeds
	feedMap := make(map[string]bool)
	for _, feed := range gs.PriceFeeds {
		if err := feed.Validate(); err != nil {
			return fmt.Errorf("invalid price feed %s: %w", feed.FeedID, err)
		}
		if feedMap[feed.FeedID] {
			return fmt.Errorf("duplicate price feed: %s", feed.FeedID)
		}
		feedMap[feed.FeedID] = true
		
		// Check that provider exists
		if !providerMap[feed.Provider] {
			return fmt.Errorf("price feed %s references unknown provider: %s", feed.FeedID, feed.Provider)
		}
	}
	
	// Validate aggregated prices
	assetMap := make(map[string]bool)
	for _, price := range gs.AggregatedPrices {
		if price.Asset == "" {
			return fmt.Errorf("aggregated price has empty asset")
		}
		if assetMap[price.Asset] {
			return fmt.Errorf("duplicate aggregated price for asset: %s", price.Asset)
		}
		assetMap[price.Asset] = true
		
		if price.Price.IsNegative() {
			return fmt.Errorf("aggregated price for %s is negative", price.Asset)
		}
		if price.NumProviders == 0 {
			return fmt.Errorf("aggregated price for %s has zero providers", price.Asset)
		}
	}
	
	// Validate price histories
	historyMap := make(map[string]bool)
	for _, history := range gs.PriceHistories {
		if history.Asset == "" {
			return fmt.Errorf("price history has empty asset")
		}
		if historyMap[history.Asset] {
			return fmt.Errorf("duplicate price history for asset: %s", history.Asset)
		}
		historyMap[history.Asset] = true
		
		// Validate price points
		for i, point := range history.Prices {
			if point.Price.IsNegative() {
				return fmt.Errorf("price history for %s has negative price at index %d", history.Asset, i)
			}
		}
	}
	
	return nil
}