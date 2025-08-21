package types

import (
	"fmt"
	"time"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PriceFeed represents a price data point from an oracle
type PriceFeed struct {
	FeedID      string    `json:"feed_id"`
	Asset       string    `json:"asset"`
	Price       sdk.Dec   `json:"price"`
	Provider    string    `json:"provider"`
	Timestamp   time.Time `json:"timestamp"`
	Expiry      time.Time `json:"expiry"`
	Confidence  sdk.Dec   `json:"confidence"`
	Volume      sdk.Int   `json:"volume"`
}

// OracleProvider represents an authorized oracle data provider
type OracleProvider struct {
	Name          string   `json:"name"`
	Address       string   `json:"address"`
	Active        bool     `json:"active"`
	Priority      uint32   `json:"priority"`
	MinSubmissions uint32  `json:"min_submissions"`
	MaxDeviation  sdk.Dec  `json:"max_deviation"`
	LastUpdate    time.Time `json:"last_update"`
	Reputation    sdk.Dec  `json:"reputation"`
}

// AggregatedPrice represents the consensus price from multiple oracles
type AggregatedPrice struct {
	Asset            string    `json:"asset"`
	Price            sdk.Dec   `json:"price"`
	MedianPrice      sdk.Dec   `json:"median_price"`
	WeightedAvgPrice sdk.Dec   `json:"weighted_avg_price"`
	StandardDev      sdk.Dec   `json:"standard_dev"`
	Confidence       sdk.Dec   `json:"confidence"`
	NumProviders     uint32    `json:"num_providers"`
	LastUpdate       time.Time `json:"last_update"`
	NextUpdate       time.Time `json:"next_update"`
}

// PriceHistory tracks historical price data
type PriceHistory struct {
	Asset     string      `json:"asset"`
	Prices    []PricePoint `json:"prices"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// PricePoint represents a single price data point in history
type PricePoint struct {
	Price     sdk.Dec   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	Volume    sdk.Int   `json:"volume"`
}

// ValidatePriceFeed validates a price feed submission
func (pf PriceFeed) Validate() error {
	if pf.FeedID == "" {
		return fmt.Errorf("feed ID cannot be empty")
	}
	if pf.Asset == "" {
		return fmt.Errorf("asset cannot be empty")
	}
	if pf.Price.IsNegative() {
		return fmt.Errorf("price cannot be negative")
	}
	if pf.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}
	if pf.Timestamp.After(pf.Expiry) {
		return fmt.Errorf("timestamp cannot be after expiry")
	}
	if pf.Confidence.LT(sdk.ZeroDec()) || pf.Confidence.GT(sdk.OneDec()) {
		return fmt.Errorf("confidence must be between 0 and 1")
	}
	return nil
}

// IsExpired checks if a price feed has expired
func (pf PriceFeed) IsExpired(currentTime time.Time) bool {
	return currentTime.After(pf.Expiry)
}

// ValidateOracleProvider validates an oracle provider
func (op OracleProvider) Validate() error {
	if op.Name == "" {
		return fmt.Errorf("provider name cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(op.Address); err != nil {
		return fmt.Errorf("invalid provider address: %w", err)
	}
	if op.MinSubmissions == 0 {
		return fmt.Errorf("minimum submissions must be greater than 0")
	}
	if op.MaxDeviation.IsNegative() || op.MaxDeviation.GT(sdk.OneDec()) {
		return fmt.Errorf("max deviation must be between 0 and 1")
	}
	if op.Reputation.LT(sdk.ZeroDec()) || op.Reputation.GT(sdk.OneDec()) {
		return fmt.Errorf("reputation must be between 0 and 1")
	}
	return nil
}

// CalculateWeightedAverage calculates weighted average price based on provider reputation
func CalculateWeightedAverage(feeds []PriceFeed, providers map[string]OracleProvider) sdk.Dec {
	if len(feeds) == 0 {
		return sdk.ZeroDec()
	}
	
	totalWeight := sdk.ZeroDec()
	weightedSum := sdk.ZeroDec()
	
	for _, feed := range feeds {
		provider, exists := providers[feed.Provider]
		if !exists || !provider.Active {
			continue
		}
		
		weight := provider.Reputation.Mul(feed.Confidence)
		weightedSum = weightedSum.Add(feed.Price.Mul(weight))
		totalWeight = totalWeight.Add(weight)
	}
	
	if totalWeight.IsZero() {
		return sdk.ZeroDec()
	}
	
	return weightedSum.Quo(totalWeight)
}

// CalculateMedianPrice calculates the median price from multiple feeds
func CalculateMedianPrice(feeds []PriceFeed) sdk.Dec {
	if len(feeds) == 0 {
		return sdk.ZeroDec()
	}
	
	prices := make([]sdk.Dec, len(feeds))
	for i, feed := range feeds {
		prices[i] = feed.Price
	}
	
	// Sort prices
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i].GT(prices[j]) {
				prices[i], prices[j] = prices[j], prices[i]
			}
		}
	}
	
	// Calculate median
	n := len(prices)
	if n%2 == 0 {
		return prices[n/2-1].Add(prices[n/2]).Quo(sdk.NewDec(2))
	}
	return prices[n/2]
}

// CalculateStandardDeviation calculates the standard deviation of prices
func CalculateStandardDeviation(feeds []PriceFeed, mean sdk.Dec) sdk.Dec {
	if len(feeds) <= 1 {
		return sdk.ZeroDec()
	}
	
	sumSquaredDiff := sdk.ZeroDec()
	for _, feed := range feeds {
		diff := feed.Price.Sub(mean)
		sumSquaredDiff = sumSquaredDiff.Add(diff.Mul(diff))
	}
	
	variance := sumSquaredDiff.Quo(sdk.NewDec(int64(len(feeds))))
	
	// Simple square root approximation for standard deviation
	// Using Newton's method for better precision
	if variance.IsZero() {
		return sdk.ZeroDec()
	}
	
	// Initial guess
	stdDev := variance.Quo(sdk.NewDec(2))
	for i := 0; i < 10; i++ {
		stdDev = stdDev.Add(variance.Quo(stdDev)).Quo(sdk.NewDec(2))
	}
	
	return stdDev
}