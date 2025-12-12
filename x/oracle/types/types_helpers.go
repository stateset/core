package types

import (
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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

// IsStale checks if a price is stale based on the given threshold.
func (p Price) IsStale(currentTime time.Time, stalenessThreshold time.Duration) bool {
	if p.UpdatedAt.IsZero() {
		return true
	}
	return currentTime.Sub(p.UpdatedAt) > stalenessThreshold
}

// DefaultOracleConfig returns a default oracle configuration.
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

// DefaultOracleParams returns default oracle parameters.
func DefaultOracleParams() OracleParams {
	return OracleParams{
		DefaultMaxDeviationBps:    500,  // 5%
		DefaultStalenessThreshold: 3600, // 1 hour
		SlashFractionBps:          1000, // 10%
		MaxProviders:              10,
		PriceHistorySize:          100,
	}
}

// CalculateDeviation calculates the percentage deviation between two prices.
func CalculateDeviation(oldPrice, newPrice sdkmath.LegacyDec) sdkmath.LegacyDec {
	if oldPrice.IsZero() {
		return sdkmath.LegacyZeroDec()
	}
	diff := newPrice.Sub(oldPrice).Abs()
	return diff.Quo(oldPrice).Mul(sdkmath.LegacyNewDec(10000)) // Return in basis points
}
