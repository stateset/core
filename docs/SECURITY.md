# Stateset Core Security Architecture

## Overview

This document describes the security mechanisms implemented in Stateset Core to protect against common attack vectors and ensure the safety of user funds.

## Table of Contents

1. [Circuit Breaker System](#circuit-breaker-system)
2. [Rate Limiting](#rate-limiting)
3. [Oracle Security](#oracle-security)
4. [Liquidation Protection](#liquidation-protection)
5. [Compliance Integration](#compliance-integration)
6. [Access Control](#access-control)
7. [Emergency Procedures](#emergency-procedures)
8. [Threat Model](#threat-model)

---

## Circuit Breaker System

### Global Pause

The system includes a global pause mechanism that can halt all operations in case of emergency.

**Features:**
- Instant activation by authorized parties
- Optional auto-resume after specified duration
- Maximum pause duration limit (default: 24 hours)
- Full audit trail with timestamps and reasons

**Usage:**
```bash
# Pause system
statesetd tx circuit pause-system \
  --reason "Security incident detected" \
  --duration 3600 \
  --from authority

# Resume system
statesetd tx circuit resume-system --from authority
```

### Module-Level Circuit Breakers

Each module can have its circuit breaker tripped independently.

**States:**
- `CLOSED`: Normal operation
- `OPEN`: Operations blocked
- `HALF_OPEN`: Recovery mode, limited operations

**Automatic Tripping:**
- Configurable failure threshold (default: 5 consecutive failures)
- Automatic reset after recovery period (default: 5 minutes)
- Specific message types can be disabled

**Manual Control:**
```bash
# Trip circuit for stablecoin module
statesetd tx circuit trip-circuit stablecoin \
  --reason "Investigating anomaly" \
  --disable-messages "/stateset.stablecoin.v1.MsgLiquidateVault" \
  --from authority

# Reset circuit
statesetd tx circuit reset-circuit stablecoin --from authority
```

---

## Rate Limiting

### Per-Address Limits

Configurable rate limits per address to prevent spam and abuse.

**Default Limits:**
| Operation | Max Requests | Window |
|-----------|-------------|--------|
| All transactions | 100 | 60s |
| Stablecoin mints | 10 | 60s |
| Large settlements | 5 | 300s |

### Global Limits

System-wide rate limits to prevent network saturation.

**Default:** 1000 transactions per 60 seconds

### Message Type Filtering

Rate limits can be applied to specific message types:
```json
{
  "name": "mint_limit",
  "max_requests": 10,
  "window_seconds": 60,
  "per_address": true,
  "message_types": ["/stateset.stablecoin.v1.MsgMintStablecoin"]
}
```

---

## Oracle Security

### Price Deviation Protection

Prevents oracle manipulation by limiting price changes per update.

**Configuration:**
```json
{
  "denom": "uatom",
  "max_deviation_bps": 500,        // 5% max per update
  "staleness_threshold_seconds": 3600,
  "min_update_interval_seconds": 60
}
```

**Protection Mechanisms:**
1. **Deviation Limits**: Rejects price updates that deviate more than configured threshold
2. **Update Frequency**: Prevents rapid-fire price updates
3. **Staleness Checks**: Operations fail if price is older than threshold

### Multi-Provider System

Support for multiple oracle providers with weighted voting:

**Provider Features:**
- Registration with governance approval
- Success rate tracking
- Automatic slashing for poor performance
- Manual slashing for malicious behavior

**Slashing Criteria:**
- Success rate below 50% after 10+ submissions
- Manual slash by governance for detected manipulation

### Price History

Maintains historical price data for:
- Audit trails
- Anomaly detection
- TWAP calculations (future)

---

## Liquidation Protection

### Surge Protection

Prevents liquidation cascades that could destabilize the system.

**Limits:**
- Maximum liquidations per block: 10
- Maximum liquidation value per block: 1,000,000 units
- Cooldown blocks after hitting limits: 5

### Liquidation Process

1. Vault health check (collateral ratio)
2. Surge protection check
3. Liquidator pays outstanding debt
4. Collateral transferred to liquidator
5. Liquidation recorded for analytics

### Incentive Structure

| Parameter | Value | Purpose |
|-----------|-------|---------|
| Liquidation Ratio | 150% | Minimum collateral before liquidation |
| Stability Fee | 2% annual | Cost of maintaining debt |
| Liquidation Penalty | 10% | Incentive for liquidators |

---

## Compliance Integration

### KYC/AML Profiles

All addresses must have a compliance profile for certain operations.

**KYC Levels:**
- `NONE`: No verification, limited operations
- `BASIC`: Basic identity verification, standard limits
- `STANDARD`: Full verification, higher limits
- `ENHANCED`: Enhanced due diligence, highest limits

**Profile Status:**
- `PENDING`: Awaiting verification
- `ACTIVE`: Fully operational
- `SUSPENDED`: Temporarily blocked
- `REJECTED`: Verification failed
- `EXPIRED`: Needs renewal

### Risk Assessment

**Risk Levels:**
- `LOW`: Standard monitoring
- `MEDIUM`: Enhanced monitoring
- `HIGH`: Enhanced due diligence required

**High-Risk Indicators:**
- Location in high-risk jurisdiction (AF, BY, MM, VE, YE)
- Business type flagged for enhanced scrutiny
- Transaction patterns triggering alerts

### Jurisdiction Controls

**Blocked Jurisdictions:**
- KP (North Korea)
- IR (Iran)
- SY (Syria)
- CU (Cuba)
- RU (Russia)

**High-Risk Jurisdictions:**
Require enhanced KYC for transactions exceeding thresholds.

### Transaction Limits

Daily and monthly limits based on KYC level:

| KYC Level | Daily Limit | Monthly Limit |
|-----------|-------------|---------------|
| NONE | $1,000 | $5,000 |
| BASIC | $10,000 | $50,000 |
| STANDARD | $100,000 | $500,000 |
| ENHANCED | Unlimited | Unlimited |

Limits are automatically reset:
- Daily: Every 24 hours from last reset
- Monthly: First transaction of new month

### Transaction Screening

Operations checked against compliance:
- Settlement instant transfers
- Escrow creation
- Batch settlements
- Payment channel operations
- Vault creation

### Sanctions Checking

Sanctioned addresses are blocked from:
- All stablecoin operations
- Settlement operations
- Payment channel participation

### Audit Logging

All profile changes are logged:
- Timestamp
- Actor (who made the change)
- Action type
- Old status → New status
- Reason (if provided)

Maximum 100 audit entries retained per profile.

---

## Access Control

### Authority-Based

**Module Authorities:**
- Oracle: Can update prices, manage providers
- Settlement: Can settle batches
- Circuit: Can pause/resume system
- Compliance: Can manage profiles

### Owner-Based

Vault operations require owner verification:
- Deposit collateral
- Withdraw collateral
- Mint stablecoin
- Repay debt

### Permissionless Operations

Certain operations are intentionally permissionless:
- Vault liquidation (incentivized by profit)
- Escrow release (by sender only)
- Channel claims (by recipient with valid nonce)

---

## Emergency Procedures

### Severity Levels

| Level | Response | Authority |
|-------|----------|-----------|
| Low | Monitor | Operations team |
| Medium | Trip affected module | Security team |
| High | Pause affected modules | Governance |
| Critical | Global pause | Emergency multisig |

### Response Playbook

1. **Detection**
   - Monitor alerts trigger
   - Community reports
   - Audit findings

2. **Assessment**
   - Determine severity
   - Identify affected modules
   - Estimate impact

3. **Containment**
   - Trip relevant circuits
   - Pause if critical
   - Disable specific messages

4. **Resolution**
   - Deploy fix
   - Test in staging
   - Governance approval

5. **Recovery**
   - Reset circuits
   - Resume operations
   - Post-mortem analysis

### Emergency Contacts

Emergency multisig requires 3-of-5 signatures for:
- Global pause activation
- Emergency parameter changes
- Emergency upgrades

---

## Threat Model

### Oracle Manipulation

**Threat:** Attacker controls oracle and sets manipulated prices

**Mitigations:**
- Price deviation limits
- Multi-provider system
- Provider slashing
- Price history tracking

### Liquidation Cascade

**Threat:** Mass liquidations destabilize system

**Mitigations:**
- Per-block liquidation limits
- Value-based limits
- Cooldown periods

### Denial of Service

**Threat:** Spam transactions overwhelm network

**Mitigations:**
- Gas-based costs
- Rate limiting (per-address and global)
- Circuit breakers

### Replay Attacks

**Threat:** Reusing signed messages

**Mitigations:**
- Nonce-based payment channels
- Sequence numbers in transactions
- Unique settlement IDs

### Flash Loan Attacks

**Threat:** Borrowing large amounts to manipulate state

**Mitigations:**
- Block-scoped operations
- Atomic transaction validation
- Compliance checks

### Insider Threats

**Threat:** Malicious authority abuse

**Mitigations:**
- Governance oversight
- Action logging
- Multi-sig requirements
- Time-locked operations

---

## Security Audit Checklist

### Pre-Audit Requirements

- [ ] All tests passing
- [ ] Test coverage > 80%
- [ ] Static analysis clean
- [ ] Dependencies audited
- [ ] Documentation complete

### Audit Scope

1. **Smart Contract Logic**
   - Stablecoin vault mechanics
   - Liquidation logic
   - Settlement flows

2. **Oracle System**
   - Price validation
   - Provider management
   - Staleness handling

3. **Access Control**
   - Authority checks
   - Owner validation
   - Compliance integration

4. **Economic Security**
   - Incentive alignment
   - Attack profitability
   - Edge cases

### Post-Audit Actions

- [ ] Address all critical findings
- [ ] Address all high findings
- [ ] Document accepted risks
- [ ] Implement monitoring
- [ ] Update documentation

---

## Monitoring Recommendations

### Key Metrics

1. **System Health**
   - Circuit breaker states
   - Rate limit hits
   - Error rates

2. **Oracle Health**
   - Price freshness
   - Deviation alerts
   - Provider performance

3. **Economic Health**
   - Total collateral
   - Total debt
   - Collateralization ratios
   - Liquidation volume

4. **Security Events**
   - Failed authorization attempts
   - Compliance blocks
   - Unusual patterns

### Alert Thresholds

| Metric | Warning | Critical |
|--------|---------|----------|
| Rate limit hits | 50% capacity | 80% capacity |
| Price staleness | 30 min | 1 hour |
| Liquidation volume | 500k/hour | 1M/hour |
| Circuit trips | Any automatic | Multiple |

---

## CosmWasm Smart Contracts

### Current Status

CosmWasm integration is currently **disabled** pending a compatible wasmd release for Cosmos SDK v0.53.x.

**Reason:** The wasmd module does not yet have a stable release compatible with Cosmos SDK v0.53.4. Development is underway ([see PR #2319](https://github.com/CosmWasm/wasmd/pull/2319)).

### Re-enablement Plan

Once a compatible wasmd version is released:

1. **Update go.mod** with compatible wasmd version
2. **Uncomment imports** in app/app.go
3. **Re-enable wasm keeper** initialization
4. **Run migration tests** for any state changes
5. **Security audit** of CosmWasm configuration

### CosmWasm Security Considerations

When re-enabled, the following security measures apply:

**Contract Deployment:**
- Governance-gated uploads (optional)
- Code pinning for frequently used contracts
- Gas limits for instantiation

**Execution Safety:**
- Gas metering prevents infinite loops
- Deterministic execution across all nodes
- Sandboxed WebAssembly environment

**Access Control:**
- Contract admin can migrate/update
- Governance can freeze contracts
- Circuit breakers apply to wasm module

---

## Version History

| Version | Date | Changes |
|---------|------|---------|
| 1.0.0 | 2024-01 | Initial security architecture |
| 1.1.0 | 2024-03 | Added circuit breakers |
| 1.2.0 | 2024-06 | Enhanced oracle security |
| 2.0.0 | 2024-12 | Major security overhaul |
| 2.1.0 | 2025-12 | Added CosmWasm status section |
| 2.2.0 | 2025-12 | Enhanced compliance with KYC/AML workflow, jurisdiction controls, transaction limits, audit logging |

---

## Contact

For security concerns, contact: security@stateset.network

For responsible disclosure, please encrypt communications using our PGP key.
