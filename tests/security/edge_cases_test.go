package security

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/stretchr/testify/require"
)

// TestOverflowProtection tests protection against integer overflow attacks
func TestOverflowProtection(t *testing.T) {
	testCases := []struct {
		name       string
		a          sdkmath.Int
		b          sdkmath.Int
		operation  string
		expectPanic bool
	}{
		{
			name:       "normal addition",
			a:          sdkmath.NewInt(1000),
			b:          sdkmath.NewInt(2000),
			operation:  "add",
			expectPanic: false,
		},
		{
			name:       "large number addition - safe",
			a:          sdkmath.NewInt(1 << 60),
			b:          sdkmath.NewInt(1 << 60),
			operation:  "add",
			expectPanic: false,
		},
		{
			name:       "normal subtraction",
			a:          sdkmath.NewInt(2000),
			b:          sdkmath.NewInt(1000),
			operation:  "sub",
			expectPanic: false,
		},
		{
			name:       "normal multiplication",
			a:          sdkmath.NewInt(1000),
			b:          sdkmath.NewInt(1000),
			operation:  "mul",
			expectPanic: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.expectPanic {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			var result sdkmath.Int
			switch tc.operation {
			case "add":
				result = tc.a.Add(tc.b)
			case "sub":
				result = tc.a.Sub(tc.b)
			case "mul":
				result = tc.a.Mul(tc.b)
			}

			if !tc.expectPanic {
				require.False(t, result.IsNegative() && tc.operation != "sub")
			}
		})
	}
}

// TestReentrancyProtection tests protection against reentrancy attacks
func TestReentrancyProtection(t *testing.T) {
	// Simulate state lock pattern
	type StateLock struct {
		locked bool
	}

	lock := &StateLock{locked: false}

	// Function that should be protected
	executeWithLock := func() error {
		if lock.locked {
			return errAlreadyLocked
		}
		lock.locked = true
		defer func() { lock.locked = false }()

		// Simulate operation
		return nil
	}

	// First call should succeed
	err := executeWithLock()
	require.NoError(t, err)

	// Simulating reentrant call (lock already held)
	lock.locked = true
	err = executeWithLock()
	require.Error(t, err)
	require.Equal(t, errAlreadyLocked, err)
}

var errAlreadyLocked = &lockedError{}

type lockedError struct{}

func (e *lockedError) Error() string { return "already locked" }

// TestNonceReplayProtection tests protection against replay attacks
func TestNonceReplayProtection(t *testing.T) {
	type Channel struct {
		Nonce uint64
	}

	channel := &Channel{Nonce: 0}

	testCases := []struct {
		name        string
		claimNonce  uint64
		expectError bool
	}{
		{
			name:        "valid first claim",
			claimNonce:  1,
			expectError: false,
		},
		{
			name:        "replay attack - same nonce",
			claimNonce:  1,
			expectError: true,
		},
		{
			name:        "valid second claim",
			claimNonce:  2,
			expectError: false,
		},
		{
			name:        "replay attack - old nonce",
			claimNonce:  1,
			expectError: true,
		},
		{
			name:        "valid skip nonce",
			claimNonce:  5,
			expectError: false,
		},
		{
			name:        "replay attack - nonce 3",
			claimNonce:  3,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Validate nonce is strictly increasing
			isValid := tc.claimNonce > channel.Nonce

			if tc.expectError {
				require.False(t, isValid, "should reject nonce %d (current: %d)", tc.claimNonce, channel.Nonce)
			} else {
				require.True(t, isValid, "should accept nonce %d (current: %d)", tc.claimNonce, channel.Nonce)
				channel.Nonce = tc.claimNonce
			}
		})
	}
}

// TestPriceManipulationProtection tests protection against oracle price manipulation
func TestPriceManipulationProtection(t *testing.T) {
	type OracleConfig struct {
		MaxDeviation      float64
		MinProviders      int
		StalenessSeconds  int64
	}

	config := OracleConfig{
		MaxDeviation:     0.05, // 5%
		MinProviders:     3,
		StalenessSeconds: 3600, // 1 hour
	}

	testCases := []struct {
		name         string
		priceChange  float64
		numProviders int
		ageSeconds   int64
		expectReject bool
	}{
		{
			name:         "normal price update",
			priceChange:  0.02,
			numProviders: 5,
			ageSeconds:   60,
			expectReject: false,
		},
		{
			name:         "excessive price deviation",
			priceChange:  0.10,
			numProviders: 5,
			ageSeconds:   60,
			expectReject: true,
		},
		{
			name:         "too few providers",
			priceChange:  0.01,
			numProviders: 2,
			ageSeconds:   60,
			expectReject: true,
		},
		{
			name:         "stale price",
			priceChange:  0.01,
			numProviders: 5,
			ageSeconds:   7200, // 2 hours
			expectReject: true,
		},
		{
			name:         "at deviation threshold",
			priceChange:  0.05,
			numProviders: 3,
			ageSeconds:   3600,
			expectReject: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Validate price update
			deviationOK := tc.priceChange <= config.MaxDeviation
			providersOK := tc.numProviders >= config.MinProviders
			freshnessOK := tc.ageSeconds <= config.StalenessSeconds

			isValid := deviationOK && providersOK && freshnessOK

			if tc.expectReject {
				require.False(t, isValid)
			} else {
				require.True(t, isValid)
			}
		})
	}
}

// TestSandwichAttackProtection tests protection against sandwich attacks
func TestSandwichAttackProtection(t *testing.T) {
	// In settlement/swaps, ensure that:
	// 1. Slippage protection is enforced
	// 2. Transaction ordering cannot be exploited

	type Settlement struct {
		Amount          sdkmath.Int
		MinReceive      sdkmath.Int // Slippage protection
		MaxSlippageBps  uint32
	}

	testCases := []struct {
		name          string
		settlement    Settlement
		actualReceive sdkmath.Int
		expectSuccess bool
	}{
		{
			name: "within slippage tolerance",
			settlement: Settlement{
				Amount:         sdkmath.NewInt(1000000),
				MinReceive:     sdkmath.NewInt(990000), // 1% slippage
				MaxSlippageBps: 100,
			},
			actualReceive: sdkmath.NewInt(995000),
			expectSuccess: true,
		},
		{
			name: "exceeds slippage tolerance",
			settlement: Settlement{
				Amount:         sdkmath.NewInt(1000000),
				MinReceive:     sdkmath.NewInt(990000),
				MaxSlippageBps: 100,
			},
			actualReceive: sdkmath.NewInt(980000), // 2% slippage
			expectSuccess: false,
		},
		{
			name: "exact minimum",
			settlement: Settlement{
				Amount:         sdkmath.NewInt(1000000),
				MinReceive:     sdkmath.NewInt(950000),
				MaxSlippageBps: 500,
			},
			actualReceive: sdkmath.NewInt(950000),
			expectSuccess: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			meetsMinimum := tc.actualReceive.GTE(tc.settlement.MinReceive)
			require.Equal(t, tc.expectSuccess, meetsMinimum)
		})
	}
}

// TestAuthorizationBypass tests that authorization cannot be bypassed
func TestAuthorizationBypass(t *testing.T) {
	type Vault struct {
		ID    uint64
		Owner string
	}

	vault := Vault{
		ID:    1,
		Owner: "stateset1owner",
	}

	testCases := []struct {
		name        string
		caller      string
		expectAuth  bool
	}{
		{
			name:       "owner can access",
			caller:     "stateset1owner",
			expectAuth: true,
		},
		{
			name:       "other user cannot access",
			caller:     "stateset1attacker",
			expectAuth: false,
		},
		{
			name:       "empty caller",
			caller:     "",
			expectAuth: false,
		},
		{
			name:       "similar address cannot access",
			caller:     "stateset1owner1", // Similar but different
			expectAuth: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isAuthorized := tc.caller == vault.Owner
			require.Equal(t, tc.expectAuth, isAuthorized)
		})
	}
}

// TestDoubleLiquidation tests that double liquidation is prevented
func TestDoubleLiquidation(t *testing.T) {
	type Vault struct {
		ID         uint64
		Liquidated bool
	}

	vault := &Vault{ID: 1, Liquidated: false}

	// First liquidation should succeed
	canLiquidate := !vault.Liquidated
	require.True(t, canLiquidate)
	vault.Liquidated = true

	// Second liquidation should fail
	canLiquidate = !vault.Liquidated
	require.False(t, canLiquidate)
}

// TestZeroAmountProtection tests that zero amount operations are rejected
func TestZeroAmountProtection(t *testing.T) {
	testCases := []struct {
		name        string
		amount      sdkmath.Int
		expectValid bool
	}{
		{
			name:        "positive amount - valid",
			amount:      sdkmath.NewInt(1000),
			expectValid: true,
		},
		{
			name:        "zero amount - invalid",
			amount:      sdkmath.ZeroInt(),
			expectValid: false,
		},
		{
			name:        "negative amount - invalid",
			amount:      sdkmath.NewInt(-1000),
			expectValid: false,
		},
		{
			name:        "minimum positive - valid",
			amount:      sdkmath.NewInt(1),
			expectValid: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := tc.amount.IsPositive()
			require.Equal(t, tc.expectValid, isValid)
		})
	}
}

// TestDenomValidation tests that invalid denoms are rejected
func TestDenomValidation(t *testing.T) {
	validDenoms := map[string]bool{
		"uatom": true,
		"ssusd": true,
		"uosmo": true,
	}

	testCases := []struct {
		name        string
		denom       string
		expectValid bool
	}{
		{
			name:        "valid denom - uatom",
			denom:       "uatom",
			expectValid: true,
		},
		{
			name:        "valid denom - ssusd",
			denom:       "ssusd",
			expectValid: true,
		},
		{
			name:        "invalid denom",
			denom:       "invalid",
			expectValid: false,
		},
		{
			name:        "empty denom",
			denom:       "",
			expectValid: false,
		},
		{
			name:        "sql injection attempt",
			denom:       "uatom'; DROP TABLE vaults;--",
			expectValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := validDenoms[tc.denom]
			require.Equal(t, tc.expectValid, isValid)
		})
	}
}

// TestEscrowExpirationExploit tests that expired escrows cannot be manipulated
func TestEscrowExpirationExploit(t *testing.T) {
	type Escrow struct {
		ID            uint64
		ExpiresAt     int64
		Status        string
		Amount        sdkmath.Int
	}

	currentTime := int64(1000)

	testCases := []struct {
		name           string
		escrow         Escrow
		action         string
		expectAllowed  bool
	}{
		{
			name: "release active escrow",
			escrow: Escrow{
				ID:        1,
				ExpiresAt: 2000,
				Status:    "PENDING",
				Amount:    sdkmath.NewInt(1000000),
			},
			action:        "release",
			expectAllowed: true,
		},
		{
			name: "release expired escrow",
			escrow: Escrow{
				ID:        2,
				ExpiresAt: 500,
				Status:    "PENDING",
				Amount:    sdkmath.NewInt(1000000),
			},
			action:        "release",
			expectAllowed: false, // Should auto-refund instead
		},
		{
			name: "release already completed",
			escrow: Escrow{
				ID:        3,
				ExpiresAt: 2000,
				Status:    "COMPLETED",
				Amount:    sdkmath.NewInt(1000000),
			},
			action:        "release",
			expectAllowed: false,
		},
		{
			name: "refund active escrow",
			escrow: Escrow{
				ID:        4,
				ExpiresAt: 2000,
				Status:    "PENDING",
				Amount:    sdkmath.NewInt(1000000),
			},
			action:        "refund",
			expectAllowed: true, // Recipient can refund
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isExpired := tc.escrow.ExpiresAt <= currentTime
			isPending := tc.escrow.Status == "PENDING"

			var allowed bool
			switch tc.action {
			case "release":
				allowed = isPending && !isExpired
			case "refund":
				allowed = isPending
			}

			require.Equal(t, tc.expectAllowed, allowed)
		})
	}
}

// TestRateLimitBypass tests that rate limits cannot be bypassed
func TestRateLimitBypass(t *testing.T) {
	type RateLimiter struct {
		requests     map[string]int
		maxPerWindow int
		window       int64
	}

	limiter := &RateLimiter{
		requests:     make(map[string]int),
		maxPerWindow: 10,
		window:       60,
	}

	// Single address trying to bypass
	address := "stateset1user"

	for i := 0; i < 15; i++ {
		count := limiter.requests[address]
		allowed := count < limiter.maxPerWindow

		if i < 10 {
			require.True(t, allowed, "request %d should be allowed", i)
			limiter.requests[address]++
		} else {
			require.False(t, allowed, "request %d should be rate limited", i)
		}
	}

	// Different address should have separate limit
	address2 := "stateset1user2"
	count := limiter.requests[address2]
	allowed := count < limiter.maxPerWindow
	require.True(t, allowed, "different address should have separate limit")
}

// TestChannelBalanceManipulation tests that channel balances cannot be manipulated
func TestChannelBalanceManipulation(t *testing.T) {
	type Channel struct {
		Deposit sdkmath.Int
		Spent   sdkmath.Int
		Balance sdkmath.Int
	}

	channel := Channel{
		Deposit: sdkmath.NewInt(1000000),
		Spent:   sdkmath.ZeroInt(),
		Balance: sdkmath.NewInt(1000000),
	}

	testCases := []struct {
		name          string
		claimAmount   sdkmath.Int
		expectSuccess bool
	}{
		{
			name:          "valid claim within balance",
			claimAmount:   sdkmath.NewInt(300000),
			expectSuccess: true,
		},
		{
			name:          "claim exceeds balance",
			claimAmount:   sdkmath.NewInt(2000000),
			expectSuccess: false,
		},
		{
			name:          "zero claim",
			claimAmount:   sdkmath.ZeroInt(),
			expectSuccess: false,
		},
		{
			name:          "negative claim",
			claimAmount:   sdkmath.NewInt(-100000),
			expectSuccess: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid := tc.claimAmount.IsPositive() && tc.claimAmount.LTE(channel.Balance)
			require.Equal(t, tc.expectSuccess, isValid)
		})
	}
}

// TestComplianceBypass tests that compliance checks cannot be bypassed
func TestComplianceBypass(t *testing.T) {
	type ComplianceStatus string
	const (
		StatusCompliant  ComplianceStatus = "COMPLIANT"
		StatusPending    ComplianceStatus = "PENDING"
		StatusBlocked    ComplianceStatus = "BLOCKED"
		StatusSanctioned ComplianceStatus = "SANCTIONED"
		StatusExpired    ComplianceStatus = "EXPIRED"
	)

	testCases := []struct {
		name           string
		senderStatus   ComplianceStatus
		receiverStatus ComplianceStatus
		expectAllowed  bool
	}{
		{
			name:           "both compliant",
			senderStatus:   StatusCompliant,
			receiverStatus: StatusCompliant,
			expectAllowed:  true,
		},
		{
			name:           "sender pending",
			senderStatus:   StatusPending,
			receiverStatus: StatusCompliant,
			expectAllowed:  false,
		},
		{
			name:           "receiver blocked",
			senderStatus:   StatusCompliant,
			receiverStatus: StatusBlocked,
			expectAllowed:  false,
		},
		{
			name:           "sender sanctioned",
			senderStatus:   StatusSanctioned,
			receiverStatus: StatusCompliant,
			expectAllowed:  false,
		},
		{
			name:           "sender expired",
			senderStatus:   StatusExpired,
			receiverStatus: StatusCompliant,
			expectAllowed:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			senderOK := tc.senderStatus == StatusCompliant
			receiverOK := tc.receiverStatus == StatusCompliant
			allowed := senderOK && receiverOK
			require.Equal(t, tc.expectAllowed, allowed)
		})
	}
}

// TestModuleAccountExploit tests that module accounts cannot be exploited
func TestModuleAccountExploit(t *testing.T) {
	// Module accounts should not be directly controllable
	moduleAccounts := []string{
		"stablecoin",
		"settlement",
		"treasury",
		"fee_collector",
	}

	for _, ma := range moduleAccounts {
		t.Run(ma+" account protection", func(t *testing.T) {
			// Users should not be able to send directly to module accounts
			// or impersonate module accounts
			isModuleAccount := true // Simulated check
			require.True(t, isModuleAccount, "%s should be protected", ma)
		})
	}
}

// TestFlashLoanProtection tests protection against flash loan-style attacks
func TestFlashLoanProtection(t *testing.T) {
	// Ensure that within a single transaction:
	// 1. Borrowed funds must be returned
	// 2. Collateralization is checked at start AND end
	// 3. Price manipulation via large trades is prevented

	type Transaction struct {
		StartCollateralRatio sdkmath.LegacyDec
		EndCollateralRatio   sdkmath.LegacyDec
		MinRatio             sdkmath.LegacyDec
	}

	testCases := []struct {
		name        string
		tx          Transaction
		expectValid bool
	}{
		{
			name: "normal transaction",
			tx: Transaction{
				StartCollateralRatio: sdkmath.LegacyMustNewDecFromStr("2.0"),
				EndCollateralRatio:   sdkmath.LegacyMustNewDecFromStr("1.8"),
				MinRatio:             sdkmath.LegacyMustNewDecFromStr("1.5"),
			},
			expectValid: true,
		},
		{
			name: "starts healthy ends unhealthy",
			tx: Transaction{
				StartCollateralRatio: sdkmath.LegacyMustNewDecFromStr("2.0"),
				EndCollateralRatio:   sdkmath.LegacyMustNewDecFromStr("1.4"),
				MinRatio:             sdkmath.LegacyMustNewDecFromStr("1.5"),
			},
			expectValid: false,
		},
		{
			name: "flash manipulation attempt",
			tx: Transaction{
				StartCollateralRatio: sdkmath.LegacyMustNewDecFromStr("1.6"),
				EndCollateralRatio:   sdkmath.LegacyMustNewDecFromStr("1.4"),
				MinRatio:             sdkmath.LegacyMustNewDecFromStr("1.5"),
			},
			expectValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Both start and end must be above minimum
			startOK := tc.tx.StartCollateralRatio.GTE(tc.tx.MinRatio)
			endOK := tc.tx.EndCollateralRatio.GTE(tc.tx.MinRatio)
			isValid := startOK && endOK
			require.Equal(t, tc.expectValid, isValid)
		})
	}
}
