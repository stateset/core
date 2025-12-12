# Integration Tests

This directory contains comprehensive integration tests for the Stateset blockchain that test cross-module workflows and interactions.

## Overview

Integration tests verify that multiple modules work together correctly by testing complete end-to-end workflows. Unlike unit tests that test individual components in isolation, integration tests ensure that the entire system functions as expected when modules interact.

## Test Suites

### 1. E-Commerce Flow (`ecommerce_flow_test.go`)

Tests the complete payment processing workflow from customer purchase to merchant receiving funds.

**Workflow Tested:**
1. User initiates payment
2. Compliance checks for both parties
3. Payment intent created and funds escrowed
4. Settlement executes
5. Merchant receives funds (minus fees)
6. Events emitted and state updated

**Test Cases:**
- `TestSuccessfulECommerceFlow` - Happy path with compliant users
- `TestBlockedUserCannotPay` - Sanctioned users blocked
- `TestInsufficientFunds` - Insufficient balance handling
- `TestPaymentToBlockedMerchant` - Payments to sanctioned merchants blocked
- `TestEscrowSettlement` - Escrow-based payment flow
- `TestBatchSettlement` - Multiple payments processed in batch
- `TestStateChangesAreConsistent` - Verify accounting consistency

**Key Features Tested:**
- Payment intent creation
- Compliance integration
- Escrow functionality
- Fee calculation and collection
- Event emission
- State consistency

### 2. Stablecoin Lifecycle (`stablecoin_lifecycle_test.go`)

Tests the complete lifecycle of collateralized stablecoin positions.

**Workflow Tested:**
1. User deposits collateral
2. Mints ssUSD against collateral
3. Oracle price changes
4. Liquidation check triggered
5. Undercollateralized vaults liquidated
6. Collateral distributed to liquidators

**Test Cases:**
- `TestCompleteStablecoinLifecycle` - Full lifecycle from creation to liquidation
- `TestVaultWithdrawCollateral` - Withdrawing while maintaining health
- `TestRepayDebt` - Repaying stablecoin debt
- `TestDepositMoreCollateral` - Adding collateral to existing vault
- `TestCannotMintOverDebtLimit` - Debt limit enforcement
- `TestLiquidationFailsForHealthyVault` - Healthy vaults cannot be liquidated
- `TestMultipleVaults` - Managing multiple vaults independently

**Key Features Tested:**
- Vault creation and management
- Collateralization ratio calculation
- Oracle price integration
- Liquidation mechanics
- Debt limits
- Collateral operations

### 3. AI Agent Transaction (`ai_agent_transaction_test.go`)

Tests autonomous AI agent service delivery and payment workflows.

**Workflow Tested:**
1. User requests AI service
2. Payment intent created
3. Funds held in escrow
4. AI agent delivers service
5. User confirms or auto-confirms
6. Payment released to agent

**Test Cases:**
- `TestSuccessfulAIAgentTransaction` - Complete service delivery flow
- `TestAIAgentServiceRefund` - Refund for unsatisfactory service
- `TestEscrowExpiration` - Auto-refund on timeout
- `TestMultipleAIAgentTransactions` - Concurrent service requests
- `TestAIAgentNonCompliantUser` - Compliance checks for AI services
- `TestPaymentChannelForStreamingAI` - Payment channels for streaming services
- `TestAIAgentComplianceRecording` - Transaction recording for compliance

**Key Features Tested:**
- Escrow for service delivery
- Refund mechanisms
- Expiration handling
- Payment channels
- Compliance integration
- Multi-service handling

### 4. Circuit Breaker (`circuit_breaker_test.go`)

Tests system protection mechanisms that detect anomalies and prevent cascading failures.

**Workflow Tested:**
1. Anomaly detected
2. Circuit breaker triggered
3. Module operations blocked
4. Recovery period enforced
5. Circuit reset and operations resume

**Test Cases:**
- `TestGlobalSystemPause` - Emergency pause of entire system
- `TestModuleCircuitTrip` - Pause specific modules
- `TestAutoCircuitTripOnFailures` - Automatic trip on threshold
- `TestCircuitAutoRecovery` - Automatic recovery after period
- `TestOraclePriceDeviationTriggersCircuit` - Price deviation protection
- `TestLiquidationSurgeProtection` - Limit liquidations per block
- `TestRateLimiting` - Rate limit high-frequency operations
- `TestSelectiveMessageDisabling` - Disable specific message types
- `TestUnauthorizedCircuitControl` - Authorization enforcement
- `TestMultipleModuleCircuits` - Independent circuit states

**Key Features Tested:**
- Global system pause
- Module-level circuit breakers
- Automatic trip on failures
- Recovery mechanisms
- Rate limiting
- Liquidation surge protection
- Oracle deviation protection

### 5. Cross-Module Compliance (`cross_module_compliance_test.go`)

Tests that compliance checks work consistently across all modules.

**Workflow Tested:**
1. User initiates operation
2. Compliance profile checked
3. KYC level verified
4. Transaction limits checked
5. Operation allowed or blocked
6. Usage recorded

**Test Cases:**
- `TestPaymentComplianceIntegration` - Payment module compliance
- `TestSettlementComplianceIntegration` - Settlement module compliance
- `TestTransactionLimitsEnforcement` - Daily/monthly limits
- `TestComplianceProfileManagement` - Profile updates and effects
- `TestAuditTrail` - Audit logging
- `TestBatchComplianceChecks` - Batch operation compliance
- `TestComplianceEventEmission` - Event emission
- `TestCrossModuleComplianceConsistency` - Consistency across modules
- `TestComplianceRecordTransaction` - Transaction recording
- `TestComplianceFailurePreventsSettlement` - Compliance blocks operations

**Key Features Tested:**
- KYC level verification
- Sanction list checking
- Profile status enforcement
- Transaction limits
- Audit trail
- Profile management
- Cross-module consistency

## Running Tests

### Run All Integration Tests
```bash
cd /home/dom/core
go test -tags=integration ./tests/integration/... -v
```

### Run Specific Test Suite
```bash
# E-Commerce flow
go test -tags=integration ./tests/integration -run TestECommerceFlowTestSuite -v

# Stablecoin lifecycle
go test -tags=integration ./tests/integration -run TestStablecoinLifecycleTestSuite -v

# AI Agent transactions
go test -tags=integration ./tests/integration -run TestAIAgentTransactionTestSuite -v

# Circuit breaker
go test -tags=integration ./tests/integration -run TestCircuitBreakerTestSuite -v

# Cross-module compliance
go test -tags=integration ./tests/integration -run TestCrossModuleComplianceTestSuite -v
```

### Run Specific Test Case
```bash
go test -tags=integration ./tests/integration -run TestECommerceFlowTestSuite/TestSuccessfulECommerceFlow -v
```

### Run with Coverage
```bash
go test -tags=integration ./tests/integration/... -v -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Structure

Each test suite follows a consistent pattern:

```go
type TestSuite struct {
    suite.Suite
    ctx            sdk.Context
    cdc            codec.Codec
    // Keepers
    accountKeeper    authkeeper.AccountKeeper
    bankKeeper       bankkeeper.Keeper
    // ... other keepers
    // Test accounts
    authority  sdk.AccAddress
    user1      sdk.AccAddress
    // ... other accounts
}

func (s *TestSuite) SetupTest() {
    // Initialize keepers
    // Setup test accounts
    // Setup initial state
}

func (s *TestSuite) TestScenario() {
    // Arrange - setup test data
    // Act - execute operations
    // Assert - verify results
}
```

## Test Best Practices

1. **Isolation**: Each test is independent and doesn't depend on other tests
2. **Setup**: `SetupTest()` creates fresh state for each test
3. **Clear Names**: Test names clearly describe what is being tested
4. **Documentation**: Tests include comments explaining the scenario
5. **Comprehensive**: Tests cover both happy path and error scenarios
6. **Assertions**: Clear assertions with helpful error messages
7. **Events**: Verify that proper events are emitted
8. **State**: Verify state changes are correct and consistent

## Error Scenarios Tested

Each test suite includes error scenario tests:

- **Compliance Failures**: Sanctioned users, expired KYC, suspended accounts
- **Insufficient Funds**: Transactions with insufficient balance
- **Authorization**: Unauthorized operations
- **Limits**: Transaction limits exceeded
- **State Violations**: Invalid state transitions
- **Timeouts**: Expiration handling
- **Liquidations**: Undercollateralized positions
- **Circuit Breakers**: Operations blocked during circuit trip

## State Verification

Tests verify:
- **Account Balances**: Before and after operations
- **Module State**: Proper state updates
- **Events**: Correct events emitted with proper attributes
- **Audit Trails**: Compliance audit logs
- **Consistency**: Totals and accounting remain consistent

## Production Readiness

These tests are production-quality and designed to:
- Catch integration issues before deployment
- Verify cross-module functionality
- Ensure compliance requirements are met
- Validate circuit breaker protection
- Test real-world workflows
- Provide confidence in system reliability

## Adding New Tests

When adding new integration tests:

1. Follow the existing test suite pattern
2. Create comprehensive test accounts with different compliance states
3. Test both success and failure paths
4. Verify state changes at each step
5. Check event emissions
6. Document the workflow being tested
7. Add comments explaining complex scenarios
8. Run all tests to ensure no regressions

## Continuous Integration

These tests should be run:
- Before each commit
- In CI/CD pipelines
- Before releases
- After any module changes
- During security audits

## Maintenance

- Keep tests updated with module changes
- Add tests for new features
- Update test data as requirements change
- Review and refactor tests periodically
- Document any test-specific configurations
