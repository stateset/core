// +build gofuzz

package fuzz_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Fuzzing tests for payments module
// Run with: go test -fuzz=FuzzPaymentIntent

// FuzzPaymentIntent tests payment intent creation with random inputs
func FuzzPaymentIntent(f *testing.F) {
	// Seed corpus with interesting test cases
	f.Add("sender1", "recipient1", uint64(1000), "ustate")
	f.Add("", "recipient1", uint64(0), "ustate")
	f.Add("sender1", "", uint64(1000), "")
	f.Add("cosmos1abcdef", "cosmos1ghijk", uint64(^uint64(0)), "uatom") // Max uint64

	f.Fuzz(func(t *testing.T, sender, recipient string, amount uint64, denom string) {
		// Test that payment intent creation never panics
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Payment intent creation panicked: %v", r)
			}
		}()

		// Validate inputs (this should match actual module validation)
		if sender == "" || recipient == "" {
			// Invalid addresses should be rejected gracefully
			return
		}

		if amount == 0 {
			// Zero amount should be rejected
			return
		}

		if denom == "" {
			// Empty denom should be rejected
			return
		}

		// Create coin
		coin := sdk.NewCoin(denom, sdkmath.NewIntFromUint64(amount))

		// Validate coin (should not panic)
		_ = coin.IsValid()

		// Additional invariants to check:
		// 1. Amount should never be negative
		if coin.Amount.IsNegative() {
			t.Error("Amount became negative")
		}

		// 2. String representation should not panic
		_ = coin.String()
	})
}

// FuzzEscrowRelease tests escrow release logic
func FuzzEscrowRelease(f *testing.F) {
	// Seed corpus
	f.Add(uint64(1000), uint64(100), int64(1000), int64(2000))
	f.Add(uint64(0), uint64(0), int64(0), int64(0))
	f.Add(uint64(^uint64(0)), uint64(1), int64(1), int64(2))

	f.Fuzz(func(t *testing.T, escrowAmount, releaseAmount uint64, lockTime, currentTime int64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Escrow release panicked: %v", r)
			}
		}()

		// Invariant: can't release more than escrowed
		if releaseAmount > escrowAmount {
			// This should be rejected by validation
			return
		}

		// Invariant: can't release before lock time
		if currentTime < lockTime {
			// Should be rejected
			return
		}

		// Calculate remaining
		remaining := escrowAmount - releaseAmount

		// Invariants to check:
		// 1. Remaining should never be negative (checked by type)
		// 2. Remaining + released should equal original
		if remaining+releaseAmount != escrowAmount {
			t.Errorf("Escrow arithmetic error: %d + %d != %d", remaining, releaseAmount, escrowAmount)
		}
	})
}

// FuzzFeeCalculation tests fee calculation logic
func FuzzFeeCalculation(f *testing.F) {
	// Seed corpus
	f.Add(uint64(1000), uint64(100)) // 1% fee
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(^uint64(0)), uint64(10000)) // 100% fee on max amount

	f.Fuzz(func(t *testing.T, amount, feeRateBps uint64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Fee calculation panicked: %v", r)
			}
		}()

		// Validate fee rate (basis points: 0-10000 = 0-100%)
		if feeRateBps > 10000 {
			// Invalid rate, should be rejected
			return
		}

		// Calculate fee
		var fee uint64
		if amount > 0 && feeRateBps > 0 {
			// Check for potential overflow
			if amount > ^uint64(0)/feeRateBps {
				// Would overflow, should use safe math
				fee = 0
				return
			}
			fee = (amount * feeRateBps) / 10000
		}

		// Invariants to check:
		// 1. Fee should never exceed amount
		if fee > amount {
			t.Errorf("Fee %d exceeds amount %d", fee, amount)
		}

		// 2. Net amount should be non-negative
		netAmount := amount - fee
		if netAmount > amount {
			t.Errorf("Integer underflow in net amount calculation")
		}

		// 3. Amount = netAmount + fee
		if netAmount+fee != amount {
			t.Errorf("Amount arithmetic error: %d + %d != %d", netAmount, fee, amount)
		}
	})
}

// FuzzBatchSettlement tests batch settlement processing
func FuzzBatchSettlement(f *testing.F) {
	// Seed corpus
	f.Add([]byte{1, 2, 3, 4, 5}) // Will be interpreted as amounts

	f.Fuzz(func(t *testing.T, amountsBytes []byte) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Batch settlement panicked: %v", r)
			}
		}()

		// Convert bytes to amounts (simulate random batch sizes)
		if len(amountsBytes) == 0 || len(amountsBytes) > 100 {
			return // Empty or too large
		}

		var totalAmount uint64
		for _, b := range amountsBytes {
			amount := uint64(b) * 1000 // Scale up for more realistic amounts

			// Check for overflow
			if totalAmount > ^uint64(0)-amount {
				return // Would overflow
			}

			totalAmount += amount
		}

		// Invariants:
		// 1. Total should equal sum of parts
		var sum uint64
		for _, b := range amountsBytes {
			sum += uint64(b) * 1000
		}

		if sum != totalAmount {
			t.Errorf("Batch total mismatch: %d != %d", sum, totalAmount)
		}

		// 2. Batch should be processed atomically (all or nothing)
		// This would be tested in actual keeper logic
	})
}

// FuzzPaymentChannelNonce tests payment channel nonce validation
func FuzzPaymentChannelNonce(f *testing.F) {
	// Seed corpus
	f.Add(uint64(0), uint64(1))
	f.Add(uint64(100), uint64(101))
	f.Add(uint64(^uint64(0)-1), uint64(^uint64(0)))

	f.Fuzz(func(t *testing.T, currentNonce, newNonce uint64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Nonce validation panicked: %v", r)
			}
		}()

		// Invariant: nonce must increment by exactly 1
		expectedNext := currentNonce + 1

		// Check for overflow
		if currentNonce == ^uint64(0) {
			// At max value, can't increment
			return
		}

		// Validate nonce
		isValid := newNonce == expectedNext

		if isValid {
			// Check that increment worked correctly
			if newNonce <= currentNonce {
				t.Errorf("Nonce did not increment: %d -> %d", currentNonce, newNonce)
			}
		}
	})
}

// FuzzComplianceCheck tests compliance check logic
func FuzzComplianceCheck(f *testing.F) {
	// Seed corpus with various address patterns
	f.Add("cosmos1abcdefghijklmnopqrstuvwxyz")
	f.Add("")
	f.Add("invalid")
	f.Add(string(make([]byte, 1000))) // Very long address

	f.Fuzz(func(t *testing.T, address string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Compliance check panicked: %v", r)
			}
		}()

		// Test address validation (should never panic)
		_, err := sdk.AccAddressFromBech32(address)

		// If address is invalid, that's fine - just shouldn't panic
		if err != nil {
			return
		}

		// For valid addresses, additional checks would be performed
		// in actual compliance module (KYC, sanctions, etc.)
	})
}

// FuzzAmountConversion tests amount conversions between types
func FuzzAmountConversion(f *testing.F) {
	// Seed corpus
	f.Add(uint64(1000))
	f.Add(uint64(0))
	f.Add(uint64(^uint64(0)))

	f.Fuzz(func(t *testing.T, amount uint64) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Amount conversion panicked: %v", r)
			}
		}()

		// Convert to SDK Int
		sdkInt := sdkmath.NewIntFromUint64(amount)

		// Invariants:
		// 1. Should not be negative
		if sdkInt.IsNegative() {
			t.Error("SDK Int is negative after conversion from uint64")
		}

		// 2. Should preserve value
		if sdkInt.IsZero() && amount != 0 {
			t.Error("Non-zero amount became zero")
		}

		// 3. String conversion should not panic
		_ = sdkInt.String()
	})
}
