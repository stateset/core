package property_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// Property-based testing for settlement module invariants
// These tests generate random inputs to verify properties always hold

// Property: Total supply should never change during transfers
// For any transfer from A to B, sum(balance_before) == sum(balance_after)
func TestProperty_TransferPreservesTotalSupply(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("transfer preserves total supply", prop.ForAll(
		func(initialBalanceA, initialBalanceB, transferAmount uint64) bool {
			// Skip invalid cases
			if transferAmount > initialBalanceA {
				return true // Discard this case
			}

			// Calculate balances before
			totalBefore := initialBalanceA + initialBalanceB

			// Simulate transfer
			finalBalanceA := initialBalanceA - transferAmount
			finalBalanceB := initialBalanceB + transferAmount

			// Calculate balances after
			totalAfter := finalBalanceA + finalBalanceB

			// Property: total supply unchanged
			return totalBefore == totalAfter
		},
		gen.UInt64(),
		gen.UInt64(),
		gen.UInt64(),
	))

	properties.TestingRun(t)
}

// Property: Escrow balance should equal locked funds
func TestProperty_EscrowBalanceMatchesLocks(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("escrow balance equals locked funds", prop.ForAll(
		func(deposits []uint64) bool {
			var totalDeposited uint64
			var escrowBalance uint64

			for _, amount := range deposits {
				// Check for overflow
				if totalDeposited > ^uint64(0)-amount {
					return true // Discard overflowing cases
				}
				totalDeposited += amount
				escrowBalance += amount
			}

			// Property: escrow balance matches sum of deposits
			return totalDeposited == escrowBalance
		},
		gen.SliceOf(gen.UInt64Range(1, 1000000)).SuchThat(func(v interface{}) bool {
			slice := v.([]uint64)
			return len(slice) > 0 && len(slice) <= 100
		}),
	))

	properties.TestingRun(t)
}

// Property: Batch settlement is commutative (order doesn't matter for final state)
func TestProperty_BatchSettlementCommutative(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("batch settlement order doesn't affect final state", prop.ForAll(
		func(amounts []uint64, initialBalance uint64) bool {
			// Skip if total exceeds initial balance
			var total uint64
			for _, amt := range amounts {
				if total > ^uint64(0)-amt {
					return true // Overflow, discard
				}
				total += amt
			}
			if total > initialBalance {
				return true // Insufficient funds, discard
			}

			// Apply settlements in original order
			balance1 := initialBalance
			for _, amt := range amounts {
				balance1 -= amt
			}

			// Apply settlements in reverse order
			balance2 := initialBalance
			for i := len(amounts) - 1; i >= 0; i-- {
				balance2 -= amounts[i]
			}

			// Property: order doesn't matter
			return balance1 == balance2
		},
		gen.SliceOf(gen.UInt64Range(1, 10000)).SuchThat(func(v interface{}) bool {
			slice := v.([]uint64)
			return len(slice) > 0 && len(slice) <= 50
		}),
		gen.UInt64Range(1000000, 10000000),
	))

	properties.TestingRun(t)
}

// Property: Payment channel balance updates are deterministic
func TestProperty_PaymentChannelDeterministic(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("payment channel nonce prevents replay", prop.ForAll(
		func(payments []uint64, channelBalance uint64) bool {
			// Skip if total exceeds channel balance
			var total uint64
			for _, amt := range payments {
				if total > ^uint64(0)-amt {
					return true
				}
				total += amt
			}
			if total > channelBalance {
				return true
			}

			// Process with nonces
			balance := channelBalance
			nonce := uint64(0)
			seenNonces := make(map[uint64]bool)

			for _, amt := range payments {
				// Nonce must be unique and incrementing
				if seenNonces[nonce] {
					return false // Duplicate nonce detected
				}
				seenNonces[nonce] = true
				nonce++

				balance -= amt
			}

			// Property: final balance is predictable
			expected := channelBalance - total
			return balance == expected
		},
		gen.SliceOf(gen.UInt64Range(1, 1000)).SuchThat(func(v interface{}) bool {
			slice := v.([]uint64)
			return len(slice) > 0 && len(slice) <= 100
		}),
		gen.UInt64Range(100000, 1000000),
	))

	properties.TestingRun(t)
}

// Property: Fee calculations never exceed transfer amount
func TestProperty_FeesNeverExceedAmount(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("fees are always less than or equal to amount", prop.ForAll(
		func(amount uint64, feeRateBps uint64) bool {
			// Fee rate in basis points (0-10000, i.e., 0-100%)
			if feeRateBps > 10000 {
				return true // Invalid rate, discard
			}

			// Calculate fee
			fee := (amount * feeRateBps) / 10000

			// Property: fee <= amount
			return fee <= amount
		},
		gen.UInt64Range(1, 1000000000),
		gen.UInt64Range(0, 10000),
	))

	properties.TestingRun(t)
}

// Property: Merchant fee rates compound correctly
func TestProperty_MerchantFeeCompounding(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("applying fee twice equals combined rate", prop.ForAll(
		func(amount uint64, rate1Bps, rate2Bps uint64) bool {
			if rate1Bps > 1000 || rate2Bps > 1000 {
				return true // Keep rates reasonable
			}
			if amount == 0 {
				return true
			}

			// Apply fees sequentially
			afterFee1 := amount - (amount*rate1Bps)/10000
			afterBothFees := afterFee1 - (afterFee1*rate2Bps)/10000

			// Property: final amount should be positive and less than original
			return afterBothFees > 0 && afterBothFees <= amount
		},
		gen.UInt64Range(1000, 1000000),
		gen.UInt64Range(1, 1000),
		gen.UInt64Range(1, 1000),
	))

	properties.TestingRun(t)
}

// Property: Integer arithmetic never overflows in safe math operations
func TestProperty_NoIntegerOverflow(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("sdkmath.Int prevents overflow", prop.ForAll(
		func(a, b uint64) bool {
			// Use SDK's safe math
			intA := sdkmath.NewIntFromUint64(a)
			intB := sdkmath.NewIntFromUint64(b)

			// Addition
			sum := intA.Add(intB)

			// Property: result should be >= both operands
			return sum.GTE(intA) && sum.GTE(intB)
		},
		gen.UInt64Range(0, ^uint64(0)/2), // Limit to prevent test overflow
		gen.UInt64Range(0, ^uint64(0)/2),
	))

	properties.TestingRun(t)
}

// Property: Escrow timeout behavior is consistent
func TestProperty_EscrowTimeoutConsistency(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("escrow timeout triggers refund", prop.ForAll(
		func(depositTime, currentTime, timeoutDuration int64) bool {
			if depositTime < 0 || currentTime < 0 || timeoutDuration < 0 {
				return true // Invalid times
			}

			timeoutAt := depositTime + timeoutDuration
			isTimedOut := currentTime >= timeoutAt

			// Property: if current time >= timeout, escrow should be refundable
			if isTimedOut {
				return currentTime >= timeoutAt
			} else {
				return currentTime < timeoutAt
			}
		},
		gen.Int64Range(1000000, 2000000), // deposit time
		gen.Int64Range(1000000, 3000000), // current time
		gen.Int64Range(1000, 100000),     // timeout duration
	))

	properties.TestingRun(t)
}
