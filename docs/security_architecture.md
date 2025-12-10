# Security Architecture Documentation

This document describes the security architecture of the Stateset Core financial modules, including threat models, security controls, and testing strategies.

## Table of Contents

1. [Executive Summary](#executive-summary)
2. [Security Principles](#security-principles)
3. [Module Security Analysis](#module-security-analysis)
4. [Threat Model](#threat-model)
5. [Security Controls](#security-controls)
6. [Testing Strategy](#testing-strategy)
7. [Incident Response](#incident-response)
8. [Audit Recommendations](#audit-recommendations)

## Executive Summary

The Stateset Core blockchain implements a comprehensive suite of financial modules including settlement, stablecoin (both collateralized and reserve-backed), payments, and oracle systems. These modules handle significant value and require rigorous security analysis.

**Key Security Properties:**
- **No Fund Loss**: All funds are accounted for in module accounts or user balances
- **No Unauthorized Access**: Strict authorization checks on all operations
- **No Invalid State**: State machines prevent illegal transitions
- **No Accounting Errors**: Runtime invariants detect discrepancies
- **Manipulation Resistant**: Oracle protections and rate limits prevent exploits

## Security Principles

### Defense in Depth

The system implements multiple layers of security:

1. **Input Validation**: All user inputs validated before processing
2. **Authorization**: Every operation checks caller permissions
3. **Business Logic**: Core security rules (collateralization, reserves)
4. **Invariants**: Runtime checks that verify accounting correctness
5. **Module Isolation**: Modules cannot directly access each other's funds
6. **SDK Guarantees**: Cosmos SDK provides atomic transactions and prevents reentrancy

### Principle of Least Privilege

- Each module account has only necessary permissions
- Users can only modify their own positions
- Authority/governance required for system-wide changes
- Module-to-module interactions minimized

### Fail-Safe Defaults

- Operations fail if compliance check fails
- Operations fail if price data unavailable or stale
- Undercollateralized positions can be liquidated
- Emergency pause mechanisms available for critical modules

## Module Security Analysis

### Settlement Module

**Purpose**: Handles instant transfers, escrow settlements, batch payments, and payment channels for stablecoin.

#### Security Properties

1. **Authorization**
   - Only sender can release escrow
   - Only recipient can refund escrow
   - Only sender can close expired channels
   - Only authority can settle batches

2. **Escrow Safety**
   - Funds locked in module account upon creation
   - Automatic refund for expired escrows
   - Status prevents double-settlement or double-refund

3. **Channel Security**
   - Signature verification prevents unauthorized claims
   - Nonce validation prevents replay attacks
   - Expiration prevents indefinite locks
   - Sender cannot close before expiration (protects recipient)

4. **Compliance Integration**
   - All parties checked for compliance
   - Batch operations verify all participants
   - Non-compliant parties cannot transact

#### Key Vulnerabilities Mitigated

| Threat | Mitigation |
|--------|------------|
| Unauthorized escrow release | Only sender address can release |
| Replay attack on channels | Nonce must strictly increase |
| Funds locked forever | Automatic expiration refund |
| Signature forgery | Cryptographic verification |
| Non-compliant transactions | Compliance checks enforced |

#### Code Locations

- **Keeper**: `/x/settlement/keeper/keeper.go`
- **Message Server**: `/x/settlement/keeper/msg_server.go`
- **Security Tests**: `/tests/security/settlement_security_test.go`

### Stablecoin Module

**Purpose**: Manages two parallel stablecoin systems:
1. Collateralized stablecoin (vaults with overcollateralization)
2. Reserve-backed stablecoin (1:1 backed by USD reserves)

#### Security Properties

1. **Collateralization** (Vault System)
   - Minimum collateral ratio enforced (e.g., 150%)
   - Oracle price used for valuation
   - Liquidation mechanism for undercollateralized vaults
   - Debt limits per vault and globally

2. **Reserve Backing** (Reserve System)
   - Reserve ratio must stay >= 100%
   - Daily mint and redeem limits
   - Minimum amounts to prevent dust spam
   - Attestation system for reserve verification

3. **Authorization**
   - Only vault owner can modify vault
   - Only authority can update reserve parameters
   - Only approved attesters can record attestations

4. **Oracle Dependency**
   - Staleness checks on all price reads
   - Operations fail if price unavailable
   - Deviation limits protect against manipulation

#### Key Vulnerabilities Mitigated

| Threat | Mitigation |
|--------|------------|
| Undercollateralized minting | Collateralization check before mint |
| Oracle manipulation | Deviation limits, staleness checks |
| Unlimited minting | Debt limits, reserve ratio requirements |
| Bad debt accumulation | Liquidation mechanism with incentives |
| Bank run on reserves | Daily redemption limits, pause mechanism |
| Unauthorized operations | Owner/authority checks on all operations |
| Accounting errors | Invariants verify supply and reserves |

#### Critical Invariants

1. **Reserve Backing Invariant**: `reserve.TotalValue >= reserve.TotalMinted * MinReserveRatioBps / 10000`
2. **Supply Match Invariant**: `TrackedSupply == BankModuleSupply`
3. **Vault Collateralization Invariant**: `CollateralValue >= Debt * LiquidationRatio` (for all vaults)
4. **Deposit Consistency Invariant**: `Sum(ActiveDeposits) <= Reserve.TotalDeposited`

#### Code Locations

- **Keeper**: `/x/stablecoin/keeper/keeper.go`
- **Reserve Logic**: `/x/stablecoin/keeper/reserve.go`
- **Invariants**: `/x/stablecoin/keeper/invariants.go`
- **Message Server**: `/x/stablecoin/keeper/msg_server.go`
- **Security Tests**: `/tests/security/stablecoin_security_test.go`

### Payments Module

**Purpose**: Simple escrow-based payments where payer locks funds and payee claims them.

#### Security Properties

1. **Escrow Model**
   - Funds transferred to module account on creation
   - Held until settlement or cancellation
   - No third-party access

2. **Authorization**
   - Only payee can settle
   - Only payer can cancel
   - Address comparison for verification

3. **Status Machine**
   - Pending → Settled (by payee)
   - Pending → Cancelled (by payer)
   - No transitions from Settled or Cancelled
   - Prevents double-spend

#### Key Vulnerabilities Mitigated

| Threat | Mitigation |
|--------|------------|
| Unauthorized settlement | Only payee address can settle |
| Double settlement | Status check prevents re-settlement |
| Insufficient balance | Balance checked before escrow creation |
| Invalid addresses | Address validation prevents fund loss |
| Self-payments | Payer and payee must differ |

#### Code Locations

- **Keeper**: `/x/payments/keeper/keeper.go`
- **Message Server**: `/x/payments/keeper/msg_server.go`
- **Security Tests**: `/tests/security/payments_security_test.go`

### Oracle Module

**Purpose**: Provides price feeds for collateral valuation in stablecoin module.

#### Security Properties

1. **Price Validation**
   - Maximum deviation limits (e.g., 5% per update)
   - Minimum update intervals (prevents spam)
   - Staleness thresholds (old prices rejected)
   - Authority or authorized provider required

2. **Provider Management**
   - Providers can be authorized/deauthorized
   - Slashing for bad behavior (low success rate)
   - Provider statistics tracked

3. **Price History**
   - Maintains historical price data
   - Size-limited to prevent bloat
   - Used for analysis and validation

#### Key Vulnerabilities Mitigated

| Threat | Mitigation |
|--------|------------|
| Price manipulation | Deviation limits reject large changes |
| Flash crash exploitation | Staleness checks prevent using old prices |
| Spam attacks | Minimum update intervals enforced |
| Unauthorized updates | Authority/provider authorization required |
| Provider misbehavior | Automatic slashing for failures |

#### Code Locations

- **Keeper**: `/x/oracle/keeper/keeper.go`
- **Message Server**: `/x/oracle/keeper/msg_server.go`
- **Security Tests**: `/x/oracle/keeper/oracle_security_test.go`

## Threat Model

### Attacker Capabilities

We consider attackers with the following capabilities:

1. **User-Level Attacker**
   - Can create accounts and submit transactions
   - Limited by transaction fees and gas limits
   - Cannot directly modify state

2. **Whale Attacker**
   - Has significant capital
   - Can attempt market manipulation
   - Can participate in governance
   - Still bound by protocol rules

3. **Oracle Provider Attacker**
   - Can submit price updates (if authorized)
   - Limited by deviation and staleness checks
   - Can be slashed for misbehavior

4. **Governance Attacker**
   - Controls governance (extreme scenario)
   - Can update parameters
   - Bounded by parameter validation
   - Community can fork chain

### Attack Scenarios

#### 1. Oracle Manipulation Attack

**Scenario**: Attacker attempts to manipulate oracle price to liquidate healthy vaults or mint excess stablecoin.

**Attack Steps**:
1. Attacker becomes authorized oracle provider
2. Submits price with large deviation (e.g., 50% drop)
3. Attempts to liquidate vaults using manipulated price

**Mitigations**:
- ✅ Deviation limits reject large price changes (e.g., max 5% per update)
- ✅ Staleness checks prevent using old manipulated prices
- ✅ Provider slashed for failed updates
- ✅ Multiple providers can be configured for consensus

**Result**: Attack prevented by oracle security controls.

#### 2. Reentrancy Attack

**Scenario**: Attacker attempts to re-enter module during fund transfer to drain escrow.

**Attack Steps**:
1. Attacker creates malicious contract (in theory)
2. Triggers escrow release
3. Attempts to re-enter settlement module

**Mitigations**:
- ✅ Cosmos SDK does not support external calls during execution
- ✅ No callback mechanisms exist
- ✅ State updates before transfers
- ✅ Status checks prevent re-execution

**Result**: Attack not possible in Cosmos architecture.

#### 3. Flash Loan Arbitrage

**Scenario**: Attacker uses flash loan to manipulate price, profit, and repay in single transaction.

**Attack Steps**:
1. Attacker borrows large amount via flash loan
2. Manipulates oracle or collateral price
3. Liquidates vaults or mints excess
4. Repays flash loan, keeps profit

**Mitigations**:
- ✅ Cosmos does not support atomic multi-message calls
- ✅ Each transaction independent
- ✅ Cannot borrow and manipulate in single block
- ✅ Oracle prices not manipulable within single transaction

**Result**: Attack not possible in Cosmos architecture.

#### 4. Reserve Bank Run

**Scenario**: Sudden loss of confidence causes mass redemption requests, draining reserves.

**Attack Steps**:
1. FUD campaign causes panic
2. Many users request redemptions simultaneously
3. Reserves depleted before all requests processed

**Mitigations**:
- ✅ Daily redemption limits prevent rapid drain
- ✅ Redemption delay provides time for response
- ✅ Emergency pause can stop new redemptions
- ✅ Reserve ratio > 100% provides buffer

**Result**: Attack mitigated by rate limits and operational controls.

#### 5. Governance Parameter Attack

**Scenario**: Attacker gains governance control and sets dangerous parameters.

**Attack Steps**:
1. Attacker accumulates voting power
2. Proposes reducing collateral ratio to 10%
3. Mass minting of undercollateralized stablecoin

**Mitigations**:
- ✅ Parameter validation rejects extreme values
- ✅ Time delays on critical changes
- ✅ Community monitoring of proposals
- ✅ Community can fork chain if compromised

**Result**: Attack limited by validation; requires community response.

#### 6. Compliance Bypass

**Scenario**: Non-compliant user attempts to transact by exploiting different modules.

**Attack Steps**:
1. User flagged as non-compliant
2. Attempts to create vault (blocked)
3. Attempts to use settlement (blocked)
4. Attempts to use payments (blocked)

**Mitigations**:
- ✅ All financial modules check compliance
- ✅ Compliance checked at operation time (not just account creation)
- ✅ No way to move significant value without passing compliance

**Result**: Attack prevented by universal compliance enforcement.

## Security Controls

### Input Validation

All user inputs are validated before processing:

| Input | Validation |
|-------|------------|
| Addresses | Bech32 format, non-empty, valid |
| Amounts | Positive, within min/max bounds |
| Denoms | Whitelisted collateral types |
| Expirations | Within reasonable bounds |
| Nonces | Strictly increasing |
| Signatures | Cryptographically verified |

### Authorization Controls

| Operation | Authorization Requirement |
|-----------|---------------------------|
| Release escrow | Must be original sender |
| Refund escrow | Must be designated recipient |
| Settle payment | Must be designated payee |
| Cancel payment | Must be original payer |
| Modify vault | Must be vault owner |
| Update oracle price | Must be authority or authorized provider |
| Update parameters | Must be authority (governance) |
| Settle batch | Must be authority |

### Rate Limiting

| Operation | Limit |
|-----------|-------|
| Settlement amount | Min/max per transaction |
| Batch size | Maximum number of settlements |
| Oracle updates | Minimum time interval between updates |
| Stablecoin minting | Daily mint limit |
| Stablecoin redemption | Daily redeem limit |
| Reserve deposits | Minimum amount to prevent dust |

### Oracle Protection

| Control | Purpose |
|---------|---------|
| Deviation limits | Reject large price changes (e.g., 5% max) |
| Staleness checks | Reject old prices (e.g., 1 hour) |
| Update intervals | Prevent spam (e.g., 60 seconds minimum) |
| Provider authorization | Only approved providers can update |
| Provider slashing | Penalize bad data submissions |

### Pause Mechanisms

| Module | Pause Capability |
|--------|------------------|
| Stablecoin | Can pause minting and redemption separately |
| Settlement | Parameters can disable features |
| Payments | Handled via compliance module |
| Oracle | Authority-only updates provide control |

## Testing Strategy

### Test Levels

1. **Unit Tests**
   - Test individual functions
   - Mock dependencies
   - Cover edge cases and error paths
   - Located in: `x/*/keeper/*_test.go`

2. **Integration Tests**
   - Test module interactions
   - Real keeper instances
   - End-to-end flows
   - Located in: `tests/integration/`

3. **Security Tests**
   - Test attack scenarios
   - Authorization bypass attempts
   - Overflow/underflow tests
   - Reentrancy tests (documentary)
   - Located in: `tests/security/`

4. **Invariant Tests**
   - Runtime accounting checks
   - Verify financial invariants
   - Catch accounting errors
   - Located in: `x/stablecoin/keeper/invariants.go`

5. **Fuzz Testing** (Recommended)
   - Random input generation
   - Discover unexpected edge cases
   - High-volume testing
   - **TODO**: Implement fuzzing framework

### Security Test Coverage

Each financial module has dedicated security test suites:

- ✅ `tests/security/settlement_security_test.go` - Settlement module security tests
- ✅ `tests/security/stablecoin_security_test.go` - Stablecoin module security tests
- ✅ `tests/security/payments_security_test.go` - Payments module security tests
- ✅ `tests/security/integration_security_test.go` - Cross-module security tests
- ✅ `x/oracle/keeper/oracle_security_test.go` - Oracle module security tests

### Critical Test Scenarios

Must be tested for each module:

- [ ] Authorization bypass attempts
- [ ] Double-spend/double-settlement
- [ ] Insufficient balance handling
- [ ] Invalid address handling
- [ ] Zero/negative amount rejection
- [ ] Overflow/underflow in calculations
- [ ] State consistency after errors
- [ ] Reentrancy protection (documentary)
- [ ] Compliance enforcement
- [ ] Rate limit enforcement
- [ ] Invariant validation

### Continuous Security Testing

1. **Pre-Commit Hooks**: Run security tests before each commit
2. **CI/CD Pipeline**: Run full test suite on pull requests
3. **Nightly Builds**: Extended test runs including fuzz testing
4. **Testnet Deployment**: Real-world testing before mainnet
5. **Bug Bounty Program**: Incentivize external security research

## Incident Response

### Detection

Monitor for:
- Invariant violations (halt chain if detected)
- Unusual price movements
- Large withdrawals/redemptions
- Failed compliance checks
- Module balance discrepancies
- Parameter changes

### Response Procedures

#### Level 1: Informational
- Unusual activity detected
- No immediate threat
- **Action**: Monitor, investigate

#### Level 2: Warning
- Potential security issue
- No active exploit
- **Action**: Investigate, prepare response

#### Level 3: Critical
- Active exploit or imminent threat
- Funds at risk
- **Action**:
  1. Engage emergency pause mechanisms
  2. Alert validators and community
  3. Analyze exploit
  4. Prepare patch
  5. Coordinate upgrade

#### Level 4: Catastrophic
- Significant funds lost or at immediate risk
- **Action**:
  1. Emergency chain halt if necessary
  2. Community coordination
  3. Forensic analysis
  4. Recovery plan
  5. Governance vote on response

### Emergency Contacts

- **Core Team**: [Contact information]
- **Validators**: [Communication channels]
- **Security Auditor**: [Contact information]
- **Community**: [Discord/Telegram]

### Post-Incident

1. **Root Cause Analysis**: Document what happened and why
2. **Fix Implementation**: Patch vulnerability
3. **Testing**: Verify fix works and doesn't break other functionality
4. **Deployment**: Coordinate upgrade with validators
5. **Post-Mortem**: Public report of incident and response
6. **Prevention**: Update tests and documentation to prevent recurrence

## Audit Recommendations

### Pre-Audit Preparation

- [ ] Complete all security tests
- [ ] Document all security assumptions
- [ ] Provide threat model
- [ ] Explain all authorization checks
- [ ] Document invariants and their criticality
- [ ] Provide test coverage reports
- [ ] Document known issues/limitations

### Audit Focus Areas

1. **Critical Path Security**
   - Fund escrow mechanisms
   - Collateralization checks
   - Oracle price validation
   - Liquidation logic
   - Reserve ratio maintenance

2. **Authorization**
   - All permission checks
   - Module account permissions
   - Governance controls
   - Authority verification

3. **Integer Safety**
   - Overflow/underflow risks
   - Fee calculations
   - Collateral value calculations
   - Debt accumulation

4. **State Consistency**
   - Invariants
   - Status transitions
   - Genesis export/import
   - Upgrade safety

5. **Oracle Security**
   - Deviation limits
   - Staleness checks
   - Provider authorization
   - Price manipulation resistance

### Auditor Questions to Address

1. What happens if oracle price unavailable?
2. Can liquidation be front-run? (Yes, but fair game)
3. What if governance is compromised?
4. How are reserve attestations verified?
5. What prevents dust spam?
6. How is compliance enforced across modules?
7. What are the pause mechanisms?
8. How are upgrades handled?

## Security Assumptions

### Trusted Components

1. **Cosmos SDK**: Correctly implements bank module, no reentrancy
2. **Validators**: Honest majority (Byzantine fault tolerance)
3. **Governance**: Acts in good faith (or community forks)
4. **Oracle Providers**: Provide accurate data (with verification)
5. **Compliance Module**: Correctly identifies participants
6. **Reserve Attesters**: Honest reporting of reserves

### Limitations

1. **Oracle Dependency**: Stablecoin system critically depends on accurate price feeds
2. **Governance Trust**: Governance can update parameters, community must monitor
3. **Liquidation Timing**: Liquidations depend on external actors to execute
4. **Reserve Verification**: Off-chain reserves must be verified by attesters
5. **Compliance Accuracy**: Compliance checks depend on external verification

## Security Checklist

Use this checklist for security reviews:

### Code Review
- [ ] All user inputs validated
- [ ] All addresses checked before use
- [ ] All amounts verified positive and within bounds
- [ ] Authorization checks on all state-modifying operations
- [ ] Status checks prevent invalid transitions
- [ ] Overflow protection in calculations
- [ ] Error handling doesn't leak information
- [ ] No hardcoded secrets or credentials

### Testing
- [ ] Unit tests for all functions
- [ ] Integration tests for cross-module flows
- [ ] Security tests for attack scenarios
- [ ] Invariant tests for accounting
- [ ] Edge cases and boundary conditions tested
- [ ] Error paths tested
- [ ] Genesis export/import tested

### Documentation
- [ ] Security assumptions documented
- [ ] Threat model described
- [ ] Authorization requirements clear
- [ ] Invariants documented
- [ ] Emergency procedures defined
- [ ] Audit recommendations provided

### Deployment
- [ ] Testnet deployment and testing
- [ ] Parameter validation in production config
- [ ] Monitoring and alerting configured
- [ ] Incident response plan in place
- [ ] Emergency contact list current
- [ ] Bug bounty program active

## Conclusion

The Stateset Core financial modules implement defense-in-depth security with multiple layers of protection:

1. **Input validation** prevents invalid data from entering the system
2. **Authorization checks** ensure only authorized parties can perform operations
3. **Business logic** enforces financial rules (collateralization, reserves, etc.)
4. **Invariants** catch any accounting errors at runtime
5. **Module isolation** prevents unauthorized inter-module access
6. **SDK guarantees** provide atomic transactions and prevent reentrancy

The system is designed to fail safely: operations fail rather than compromise security. Emergency pause mechanisms provide a last line of defense.

Ongoing security requires:
- Continuous testing and monitoring
- Regular security audits
- Active bug bounty program
- Community vigilance
- Rapid incident response capability

This document should be updated as the system evolves and new threats are identified.

---

**Document Version**: 1.0
**Last Updated**: 2025-12-10
**Next Review**: Q1 2026
**Maintained By**: Stateset Security Team
