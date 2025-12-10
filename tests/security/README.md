# Security Test Suite

This directory contains comprehensive security tests for all financial modules in Stateset Core.

## Overview

The security test suite validates that financial modules (settlement, stablecoin, payments, oracle) are protected against common vulnerabilities and attack vectors.

## Test Files

### Module-Specific Tests

- **`settlement_security_test.go`** - Settlement module security tests
  - Authorization controls (escrow release, refund, channel operations)
  - Reentrancy protection
  - Input validation (amounts, addresses, signatures)
  - State consistency (escrow expiration, status transitions)
  - Edge cases (zero amounts, self-transfers, batch limits)

- **`stablecoin_security_test.go`** - Stablecoin module security tests
  - Collateralization enforcement
  - Authorization controls (vault operations, reserve params)
  - Oracle manipulation resistance
  - Reserve-backed system security (daily limits, reserve ratio)
  - Liquidation safety
  - Financial invariants
  - Pause mechanisms

- **`payments_security_test.go`** - Payments module security tests
  - Authorization controls (settlement, cancellation)
  - State transition safety (double-settlement prevention)
  - Escrow integrity
  - Compliance enforcement
  - Input validation
  - Module balance consistency

### Cross-Module Tests

- **`integration_security_test.go`** - Cross-module security tests
  - Oracle dependency security
  - Compliance integration across modules
  - Module account isolation
  - Cross-module fund flows
  - Gas and DoS protection
  - Emergency pause mechanisms
  - Attack simulations (oracle manipulation, governance attacks)

## Running Tests

### Run All Security Tests

```bash
go test ./tests/security/... -v
```

### Run Specific Module Tests

```bash
# Settlement security tests
go test ./tests/security/ -run TestSettlementSecurityTestSuite -v

# Stablecoin security tests
go test ./tests/security/ -run TestStablecoinSecurityTestSuite -v

# Payments security tests
go test ./tests/security/ -run TestPaymentsSecurityTestSuite -v

# Integration security tests
go test ./tests/security/ -run TestIntegrationSecurityTestSuite -v
```

### Run Specific Test Cases

```bash
# Run specific authorization test
go test ./tests/security/ -run TestReleaseEscrow_Unauthorized -v

# Run all reentrancy tests
go test ./tests/security/ -run Reentrancy -v

# Run all input validation tests
go test ./tests/security/ -run Validation -v
```

## Test Structure

Each test file follows this structure:

```go
type ModuleSecurityTestSuite struct {
    suite.Suite
    keeper keeper.Keeper
    ctx    sdk.Context
    addrs  []sdk.AccAddress
}
```

Tests are organized into categories:

1. **Authorization Tests** - Verify only authorized parties can perform operations
2. **Reentrancy Tests** - Verify reentrancy protection (documentary for Cosmos)
3. **Input Validation Tests** - Verify all inputs are validated before use
4. **Overflow/Underflow Tests** - Verify safe arithmetic operations
5. **State Consistency Tests** - Verify state transitions are valid
6. **Compliance Tests** - Verify compliance is enforced
7. **Rate Limiting Tests** - Verify rate limits and minimums
8. **Edge Case Tests** - Verify boundary conditions and corner cases
9. **Invariant Tests** - Verify financial invariants always hold
10. **Documentation Tests** - Document security models and assumptions

## Security Test Checklist

When adding new features, ensure these security aspects are tested:

- [ ] **Authorization**: Can unauthorized users perform the operation?
- [ ] **Input Validation**: Are all inputs validated (addresses, amounts, etc.)?
- [ ] **State Consistency**: Can the operation leave state in invalid condition?
- [ ] **Overflow/Underflow**: Are arithmetic operations safe?
- [ ] **Reentrancy**: Is the operation protected against reentrancy? (documentary)
- [ ] **Compliance**: Is compliance enforced?
- [ ] **Rate Limiting**: Are there appropriate limits?
- [ ] **Edge Cases**: What happens with zero, max, negative values?
- [ ] **Invariants**: Do financial invariants still hold?
- [ ] **Error Handling**: Are errors handled safely?

## Current Status

### Implemented Tests

The test files contain comprehensive test templates that document all security test scenarios. These include:

- ✅ Authorization test templates for all critical operations
- ✅ Reentrancy protection documentation (Cosmos SDK provides this)
- ✅ Input validation test templates
- ✅ State consistency test templates
- ✅ Invariant test templates
- ✅ Edge case test templates
- ✅ Attack simulation templates
- ✅ Security model documentation

### Pending Implementation

The test templates use `suite.T().Skip("Requires full keeper setup")` as placeholders. To fully implement:

1. **Setup Test Fixtures**: Create proper keeper initialization for each module
2. **Mock Dependencies**: Set up mock bank, oracle, and compliance keepers
3. **Implement Test Logic**: Fill in the TODO sections with actual test code
4. **Add Assertions**: Verify expected behavior with require/assert statements
5. **Run and Validate**: Execute tests and verify they catch security issues

## Integration with CI/CD

These security tests should be integrated into the CI/CD pipeline:

1. **Pre-commit hooks**: Run critical security tests before commits
2. **Pull request checks**: Run full security test suite on PRs
3. **Nightly builds**: Run extended tests including fuzzing
4. **Pre-deployment**: Run all security tests before mainnet deployment

## Security Test Best Practices

### 1. Test Attack Scenarios

Don't just test happy paths. Actively try to break the system:

```go
// Good: Tests unauthorized access attempt
func TestUnauthorizedAccess() {
    // Attacker tries to release someone else's escrow
    err := keeper.ReleaseEscrow(ctx, settlementId, attackerAddr)
    require.Error(t, err)
    require.ErrorIs(t, err, types.ErrUnauthorized)
}
```

### 2. Document Threat Model

Each test should document what threat it's testing:

```go
// SECURITY: Prevent unauthorized release of escrowed funds
// Threat: Attacker attempts to release another user's escrowed funds
func TestReleaseEscrow_Unauthorized() {
    // Test implementation
}
```

### 3. Verify Complete Prevention

Ensure the vulnerability is completely prevented, not just harder:

```go
// Not enough: Check error is returned
require.Error(t, err)

// Better: Check specific error type
require.ErrorIs(t, err, types.ErrUnauthorized)

// Best: Also verify state unchanged
settlement, _ := keeper.GetSettlement(ctx, id)
require.Equal(t, types.SettlementStatusPending, settlement.Status)
require.True(t, balance.Equal(initialBalance)) // Funds still escrowed
```

### 4. Test Edge Cases

Test boundary conditions:

```go
// Test with zero amount
// Test with max amount
// Test with negative amount (should fail at validation)
// Test with amounts just below/above limits
// Test with empty addresses
// Test with invalid addresses
```

### 5. Test Invariants

After any operation, verify invariants still hold:

```go
// After settlement, verify module balance decreased correctly
moduleBalance := bankKeeper.GetBalance(ctx, moduleAddr, denom)
expectedBalance := initialModuleBalance.Sub(settlement.NetAmount)
require.True(t, moduleBalance.Equal(expectedBalance))
```

## Related Documentation

- [Security Architecture](../../docs/security_architecture.md) - Overall security design
- [Threat Model](../../docs/security_architecture.md#threat-model) - Detailed threat analysis
- [Invariants](../../x/stablecoin/keeper/invariants.go) - Financial invariants implementation

## Contributing

When adding new security tests:

1. Follow the existing test structure and naming conventions
2. Document the threat model in comments
3. Test both positive (authorized) and negative (unauthorized) cases
4. Verify state consistency after operations
5. Add test to this README if it's a new category
6. Update security documentation if needed

## Questions or Issues

If you discover a security issue:

1. **DO NOT** open a public GitHub issue
2. Contact the security team privately: security@stateset.io
3. Provide detailed reproduction steps
4. Wait for confirmation before disclosure

## License

Copyright Stateset. All rights reserved.
