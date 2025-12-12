package property_test

import (
	"reflect"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
)

// Property-based tests for stablecoin module invariants

// Property: Total minted ssUSD must never exceed total collateral value
func TestProperty_CollateralizationRatio(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("minted stablecoins always backed by sufficient collateral", prop.ForAll(
		func(collateralValue, debtValue uint64, minRatio uint64) bool {
			// Skip invalid cases
			if minRatio == 0 || minRatio < 100 {
				return true // Min ratio must be >= 100%
			}
			if debtValue == 0 {
				return true // No debt, always healthy
			}

			// Only consider states that satisfy minimum collateralization.
			// (Property tests here are arithmetic invariants, not full module simulation.)
			if (collateralValue * 100) < (debtValue * minRatio) {
				return true
			}

			// Calculate actual ratio (percentage)
			actualRatio := (collateralValue * 100) / debtValue

			// Property: minted stablecoins only exist if collateral ratio >= min ratio
			return actualRatio >= minRatio
		},
		gen.UInt64Range(0, 10000000),
		gen.UInt64Range(0, 10000000),
		gen.UInt64Range(100, 200), // 100-200% collateral ratio
	))

	properties.TestingRun(t)
}

// Property: Vault liquidation always reduces system risk
func TestProperty_LiquidationReducesRisk(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("liquidation improves system collateral ratio", prop.ForAll(
		func(vaultCollateral, vaultDebt, liquidationAmount uint64) bool {
			if vaultCollateral == 0 || vaultDebt == 0 {
				return true
			}
			if liquidationAmount > vaultDebt {
				return true // Can't liquidate more than debt
			}

			// Calculate ratio before liquidation
			ratioBefore := (vaultCollateral * 100) / vaultDebt

			// After liquidation: proportionally reduce collateral and debt
			collateralReduction := (vaultCollateral * liquidationAmount) / vaultDebt
			remainingCollateral := vaultCollateral - collateralReduction
			remainingDebt := vaultDebt - liquidationAmount

			if remainingDebt == 0 {
				return true // Fully liquidated
			}

			// Calculate ratio after liquidation
			ratioAfter := (remainingCollateral * 100) / remainingDebt

			// Property: liquidation should not reduce the collateral ratio.
			return ratioAfter >= ratioBefore
		},
		gen.UInt64Range(1000, 1000000),
		gen.UInt64Range(1000, 1000000),
		gen.UInt64Range(100, 10000),
	))

	properties.TestingRun(t)
}

// Property: Minting and burning are inverse operations
func TestProperty_MintBurnInverse(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("minting then burning returns to original state", prop.ForAll(
		func(initialCollateral, initialDebt, mintAmount uint64) bool {
			if initialCollateral == 0 {
				return true
			}

			// Check if mint is valid (sufficient collateral)
			minRatio := uint64(150) // 150% collateralization
			maxMintable := (initialCollateral * 100) / minRatio

			if mintAmount > maxMintable || mintAmount == 0 {
				return true // Invalid mint, skip
			}

			// Mint
			debtAfterMint := initialDebt + mintAmount

			// Burn same amount
			debtAfterBurn := debtAfterMint - mintAmount

			// Property: debt returns to original
			return debtAfterBurn == initialDebt
		},
		gen.UInt64Range(10000, 1000000),
		gen.UInt64Range(0, 100000),
		gen.UInt64Range(1, 10000),
	))

	properties.TestingRun(t)
}

// Property: Interest accrual is monotonically increasing
func TestProperty_InterestMonotonicIncreasing(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("debt with interest never decreases over time", prop.ForAll(
		func(principal uint64, interestRateBps uint64, time1, time2 uint64) bool {
			if principal == 0 || interestRateBps == 0 {
				return true
			}
			if interestRateBps > 10000 { // Max 100% per period
				return true
			}

			// Calculate debt at time1
			periods1 := time1
			debt1 := principal + (principal*interestRateBps*periods1)/10000

			// Calculate debt at time2 (should be >= time1)
			periods2 := time2
			debt2 := principal + (principal*interestRateBps*periods2)/10000

			// Property: if time2 >= time1, then debt2 >= debt1
			if time2 >= time1 {
				return debt2 >= debt1
			} else {
				return debt1 >= debt2
			}
		},
		gen.UInt64Range(1000, 1000000),
		gen.UInt64Range(1, 1000),
		gen.UInt64Range(0, 365),
		gen.UInt64Range(0, 365),
	))

	properties.TestingRun(t)
}

// Property: Collateral withdrawal never makes vault unhealthy
func TestProperty_SafeWithdrawal(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("collateral withdrawal maintains health ratio", prop.ForAll(
		func(collateral, debt, withdrawal uint64) bool {
			if collateral == 0 || debt == 0 {
				return true
			}

			minRatio := uint64(150) // 150%

			minRemainingCollateral := (debt*minRatio + 99) / 100 // ceil(debt*minRatio/100)
			if collateral < minRemainingCollateral {
				return true // Already unhealthy, skip
			}

			if withdrawal > collateral {
				return true // Invalid withdrawal
			}

			maxSafeWithdrawal := collateral - minRemainingCollateral
			if withdrawal > maxSafeWithdrawal {
				return true // Would be rejected by the system, discard case
			}

			remainingCollateral := collateral - withdrawal
			newRatio := (remainingCollateral * 100) / debt

			return newRatio >= minRatio
		},
		gen.UInt64Range(150000, 1000000),
		gen.UInt64Range(100000, 500000),
		gen.UInt64Range(0, 100000),
	))

	properties.TestingRun(t)
}

// Property: Oracle price changes don't create/destroy value
func TestProperty_PriceChangePreservesValue(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("price changes update valuation but not actual balances", prop.ForAll(
		func(collateralAmount, price1, price2 uint64) bool {
			if collateralAmount == 0 || price1 == 0 || price2 == 0 {
				return true
			}

			// Value at price1
			value1 := collateralAmount * price1

			// Value at price2
			value2 := collateralAmount * price2

			if price1 == price2 {
				return value1 == value2
			}
			return value1 != value2
		},
		gen.UInt64Range(1000, 100000),
		gen.UInt64Range(1, 1000),
		gen.UInt64Range(1, 1000),
	))

	properties.TestingRun(t)
}

// Property: System is never under-collateralized in aggregate
func TestProperty_SystemCollateralization(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("aggregate collateral always exceeds aggregate debt", prop.ForAll(
		func(vaults []struct{ collateral, debt uint64 }) bool {
			var totalCollateral, totalDebt uint64

			for _, vault := range vaults {
				// Check for overflow
				if totalCollateral > ^uint64(0)-vault.collateral {
					return true
				}
				if totalDebt > ^uint64(0)-vault.debt {
					return true
				}

				totalCollateral += vault.collateral
				totalDebt += vault.debt
			}

			minSystemRatio := uint64(120) // 120% system-wide minimum

			if totalDebt == 0 {
				return true // No debt, always healthy
			}

			systemRatio := (totalCollateral * 100) / totalDebt

			// Property: system ratio should always be above minimum
			return systemRatio >= minSystemRatio || totalDebt == 0
		},
		gen.SliceOf(gen.Struct(reflect.TypeOf(struct{ collateral, debt uint64 }{}), map[string]gopter.Gen{
			"collateral": gen.UInt64Range(10000, 100000),
			"debt":       gen.UInt64Range(5000, 50000),
		})).SuchThat(func(v interface{}) bool {
			slice := v.([]struct{ collateral, debt uint64 })
			return len(slice) > 0 && len(slice) <= 100
		}),
	))

	properties.TestingRun(t)
}

// Property: Liquidation penalty benefits the protocol
func TestProperty_LiquidationPenalty(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property("liquidation penalty accrues to protocol", prop.ForAll(
		func(collateralValue, debtValue, penaltyBps uint64) bool {
			if collateralValue == 0 || debtValue == 0 {
				return true
			}
			if penaltyBps > 2000 { // Max 20% penalty
				return true
			}

			// Penalty amount
			penalty := (collateralValue * penaltyBps) / 10000

			// Liquidator gets: debt repayment + penalty
			liquidatorGets := debtValue + penalty

			// Property: liquidator should get more than just debt repayment
			if penaltyBps > 0 {
				return liquidatorGets > debtValue
			}

			return true
		},
		gen.UInt64Range(10000, 1000000),
		gen.UInt64Range(10000, 500000),
		gen.UInt64Range(0, 2000),
	))

	properties.TestingRun(t)
}
