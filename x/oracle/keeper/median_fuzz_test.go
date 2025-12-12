package keeper

import (
	"testing"

	sdkmath "cosmossdk.io/math"

	"github.com/stateset/core/x/oracle/types"
)

func FuzzMedianPendingPrice(f *testing.F) {
	f.Add(int64(100), int64(110))
	f.Add(int64(1), int64(2))

	f.Fuzz(func(t *testing.T, a, b int64) {
		if a == 0 {
			a = 1
		}
		if b == 0 {
			b = 1
		}
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}

		pending := []types.PendingPrice{
			{Denom: "uusdc", Amount: sdkmath.LegacyNewDec(a), Provider: "provider1"},
			{Denom: "uusdc", Amount: sdkmath.LegacyNewDec(b), Provider: "provider2"},
		}

		median, _ := medianPendingPrice(pending)

		lo := pending[0].Amount
		hi := pending[1].Amount
		if hi.LT(lo) {
			lo, hi = hi, lo
		}

		if median.LT(lo) || median.GT(hi) {
			t.Fatalf("median %s out of range [%s, %s]", median.String(), lo.String(), hi.String())
		}
	})
}
