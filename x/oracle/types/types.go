package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Price represents the current oracle price for a denom.
type Price struct {
	Denom       string            `json:"denom" yaml:"denom"`
	Amount      sdkmath.LegacyDec `json:"amount" yaml:"amount"`
	LastUpdater string            `json:"last_updater" yaml:"last_updater"`
	LastHeight  int64             `json:"last_height" yaml:"last_height"`
	UpdatedAt   time.Time         `json:"updated_at" yaml:"updated_at"`
}

// ValidateBasic performs stateless validation of the price.
func (p Price) ValidateBasic() error {
	if len(p.Denom) == 0 {
		return errorsmod.Wrap(ErrInvalidDenom, "denom cannot be empty")
	}
	if !p.Amount.IsPositive() {
		return errorsmod.Wrap(ErrInvalidPrice, "price must be positive")
	}
	if p.LastUpdater != "" {
		if _, err := sdk.AccAddressFromBech32(p.LastUpdater); err != nil {
			return errorsmod.Wrap(ErrInvalidAuthority, err.Error())
		}
	}
	return nil
}

// IsStale checks if a price is stale based on the given threshold
func (p Price) IsStale(currentTime time.Time, stalenessThreshold time.Duration) bool {
	if p.UpdatedAt.IsZero() {
		return true
	}
	return currentTime.Sub(p.UpdatedAt) > stalenessThreshold
}

// OracleConfig holds configuration for oracle price updates
type OracleConfig struct {
	// Denom is the asset denom
	Denom string `json:"denom"`
	// MaxDeviationBps is the maximum allowed deviation per update (basis points)
	MaxDeviationBps uint64 `json:"max_deviation_bps"`
	// StalenessThresholdSeconds is when a price is considered stale
	StalenessThresholdSeconds int64 `json:"staleness_threshold_seconds"`
	// MinUpdateIntervalSeconds is the minimum time between updates
	MinUpdateIntervalSeconds int64 `json:"min_update_interval_seconds"`
	// RequiredConfirmations is the number of oracle confirmations needed
	RequiredConfirmations uint32 `json:"required_confirmations"`
	// Enabled indicates if this config is active
	Enabled bool `json:"enabled"`
}

// DefaultOracleConfig returns a default oracle configuration
func DefaultOracleConfig(denom string) OracleConfig {
	return OracleConfig{
		Denom:                     denom,
		MaxDeviationBps:           500,  // 5%
		StalenessThresholdSeconds: 3600, // 1 hour
		MinUpdateIntervalSeconds:  60,   // 1 minute
		RequiredConfirmations:     1,
		Enabled:                   true,
	}
}

// OracleProvider represents a registered oracle provider
type OracleProvider struct {
	// Address is the provider's address
	Address string `json:"address"`
	// Name is a human-readable name
	Name string `json:"name"`
	// Weight is the provider's voting weight (for aggregation)
	Weight uint64 `json:"weight"`
	// IsActive indicates if the provider can submit prices
	IsActive bool `json:"is_active"`
	// TotalSubmissions is the total number of price submissions
	TotalSubmissions uint64 `json:"total_submissions"`
	// SuccessfulSubmissions is the number of accepted submissions
	SuccessfulSubmissions uint64 `json:"successful_submissions"`
	// Slashed indicates if the provider has been slashed
	Slashed bool `json:"slashed"`
	// SlashCount is the number of times slashed
	SlashCount uint32 `json:"slash_count"`
}

// PendingPrice represents a price awaiting confirmation
type PendingPrice struct {
	Denom         string            `json:"denom"`
	Amount        sdkmath.LegacyDec `json:"amount"`
	Provider      string            `json:"provider"`
	SubmittedAt   time.Time         `json:"submitted_at"`
	Confirmations []string          `json:"confirmations"`
}

// PriceHistory stores historical price data for a denom
type PriceHistory struct {
	Denom   string       `json:"denom"`
	Prices  []PricePoint `json:"prices"`
	MaxSize uint32       `json:"max_size"`
}

// PricePoint is a single historical price entry
type PricePoint struct {
	Amount    sdkmath.LegacyDec `json:"amount"`
	Timestamp time.Time         `json:"timestamp"`
	Height    int64             `json:"height"`
}

// CalculateDeviation calculates the percentage deviation between two prices
func CalculateDeviation(oldPrice, newPrice sdkmath.LegacyDec) sdkmath.LegacyDec {
	if oldPrice.IsZero() {
		return sdkmath.LegacyZeroDec()
	}
	diff := newPrice.Sub(oldPrice).Abs()
	return diff.Quo(oldPrice).Mul(sdkmath.LegacyNewDec(10000)) // Return in basis points
}

// OracleParams holds module parameters
type OracleParams struct {
	// DefaultMaxDeviationBps is the default max deviation
	DefaultMaxDeviationBps uint64 `json:"default_max_deviation_bps"`
	// DefaultStalenessThreshold in seconds
	DefaultStalenessThreshold int64 `json:"default_staleness_threshold"`
	// SlashFractionBps is the slash amount in basis points
	SlashFractionBps uint64 `json:"slash_fraction_bps"`
	// MaxProviders is the maximum number of oracle providers
	MaxProviders uint32 `json:"max_providers"`
	// PriceHistorySize is how many price points to keep
	PriceHistorySize uint32 `json:"price_history_size"`
}

// DefaultOracleParams returns default oracle parameters
func DefaultOracleParams() OracleParams {
	return OracleParams{
		DefaultMaxDeviationBps:    500,   // 5%
		DefaultStalenessThreshold: 3600,  // 1 hour
		SlashFractionBps:          1000,  // 10%
		MaxProviders:              10,
		PriceHistorySize:          100,
	}
}
