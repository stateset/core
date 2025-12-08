# Compliance Module

The Compliance module provides KYC/AML compliance checking and transaction limits for the Stateset blockchain.

## Overview

The compliance module handles:
- **Compliance Profiles**: Store compliance status per address
- **Sanction Lists**: Manage blocked/sanctioned addresses
- **Transaction Limits**: Enforce daily/monthly transaction limits
- **Jurisdiction Controls**: Block transactions from restricted jurisdictions

## Features

### Compliance Profiles
- Per-address compliance status
- Verification expiration tracking
- Risk level classification
- Document verification status

### Sanction Management
- Sanction list maintenance
- Real-time sanction checking
- Bulk sanction updates

### Transaction Limits
- Daily transaction limits
- Monthly transaction limits
- Per-denomination limits
- Automatic limit reset

### Jurisdiction Controls
- Country-based blocking
- Regional restrictions
- Configurable jurisdiction rules

## Compliance States

| State | Description |
|-------|-------------|
| `COMPLIANT` | Address passed all compliance checks |
| `PENDING` | Verification in progress |
| `BLOCKED` | Address blocked (temporary) |
| `SANCTIONED` | Address on sanction list (permanent) |
| `EXPIRED` | Compliance verification expired |

## Messages

| Message | Description |
|---------|-------------|
| `MsgSetProfile` | Set compliance profile for address |
| `MsgUpdateProfile` | Update existing profile |
| `MsgAddToSanctionList` | Add address to sanction list |
| `MsgRemoveFromSanctionList` | Remove from sanction list |
| `MsgBlockJurisdiction` | Block a jurisdiction |
| `MsgUnblockJurisdiction` | Unblock a jurisdiction |

## Queries

| Query | Description |
|-------|-------------|
| `Profile` | Get compliance profile |
| `IsCompliant` | Check if address is compliant |
| `SanctionList` | Get sanction list |
| `BlockedJurisdictions` | Get blocked jurisdictions |
| `TransactionLimits` | Get transaction limits for address |

## Integration

Other modules integrate with compliance via:

```go
// Check compliance before transaction
if err := compKeeper.AssertCompliant(ctx, address); err != nil {
    return err
}
```

## Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `default_daily_limit` | 10000 | Default daily transaction limit |
| `default_monthly_limit` | 100000 | Default monthly limit |
| `verification_expiry_days` | 365 | Days until verification expires |

## State

| Key | Value |
|-----|-------|
| `0x01{address}` | ComplianceProfile |
| `0x02{address}` | SanctionEntry |
| `0x03{jurisdiction}` | JurisdictionBlock |
| `0x04` | Params |

## Events

| Event | Attributes |
|-------|------------|
| `profile_updated` | address, status |
| `sanction_added` | address, reason |
| `sanction_removed` | address |
| `jurisdiction_blocked` | code |
| `limit_exceeded` | address, limit_type |
