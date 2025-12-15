package types

import (
	"encoding/json"
	"time"

	sdkmath "cosmossdk.io/math"
)

// PaymentRoute defines the path and cost for a payment execution.
type PaymentRoute struct {
	Hops        []string           `json:"hops"`         // Sequence of intermediary addresses (bech32)
	TotalFee    sdkmath.LegacyDec  `json:"total_fee"`    // Cumulative fees estimated
	Probability sdkmath.LegacyDec  `json:"probability"`  // Estimated success probability (0-1)
	Latency     time.Duration      `json:"latency"`      // Expected duration
}

// NewPaymentRoute creates a new route.
func NewPaymentRoute(hops []string, fee sdkmath.LegacyDec, prob sdkmath.LegacyDec, lat time.Duration) PaymentRoute {
	return PaymentRoute{
		Hops:        hops,
		TotalFee:    fee,
		Probability: prob,
		Latency:     lat,
	}
}

// Validate checks if the route is valid.
func (r PaymentRoute) Validate() error {
	if len(r.Hops) == 0 {
		return nil // Direct route is valid (empty hops)
	}
	if r.TotalFee.IsNegative() {
		return ErrInvalidAmount // Reuse or define new error
	}
	if r.Probability.GT(sdkmath.LegacyOneDec()) || r.Probability.IsNegative() {
		return ErrInvalidAmount
	}
	return nil
}

// String returns a JSON representation of the route.
func (r PaymentRoute) String() string {
	bz, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(bz)
}
