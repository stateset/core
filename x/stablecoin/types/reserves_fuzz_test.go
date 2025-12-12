package types

import (
	"testing"

	sdkmath "cosmossdk.io/math"
)

func FuzzReserve_GetReserveRatio(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(1000), int64(500))

	f.Fuzz(func(t *testing.T, totalValue, totalMinted int64) {
		if totalValue < 0 {
			totalValue = -totalValue
		}
		if totalMinted < 0 {
			totalMinted = -totalMinted
		}

		r := Reserve{
			TotalValue:  sdkmath.NewInt(totalValue),
			TotalMinted: sdkmath.NewInt(totalMinted),
		}

		ratio := r.GetReserveRatio()

		if totalMinted == 0 && ratio != 10000 {
			t.Fatalf("expected 10000 bps for zero minted, got %d", ratio)
		}
		if ratio > 100000 {
			t.Fatalf("reserve ratio capped at 100000 bps, got %d", ratio)
		}
	})
}
