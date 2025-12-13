# Stateset Core Tokenomics

## Overview

Stateset Core operates with a dual-token model designed for enterprise commerce and settlement:

| Token | Symbol | Purpose |
|-------|--------|---------|
| **STATE** | `ustate` | Native staking/governance token |
| **ssUSD** | `ssusd` | Collateral-backed stablecoin |

---

## STATE Token Economics

### Supply Parameters

| Parameter | Value |
|-----------|-------|
| **Initial Supply** | 1,000,000,000 STATE |
| **Max Supply** | Uncapped (inflationary) |
| **Inflation Rate** | 7-20% annually (dynamic) |
| **Target Bonded Ratio** | 67% |
| **Minimum Inflation** | 7% |
| **Maximum Inflation** | 20% |

### Inflation Mechanism

The inflation rate dynamically adjusts based on the bonded ratio:

```
If bondedRatio < 67%:
    inflation increases toward 20% (incentivize staking)
If bondedRatio > 67%:
    inflation decreases toward 7% (reduce dilution)
```

### Token Distribution (Genesis)

| Allocation | Percentage | Tokens | Vesting |
|------------|------------|--------|---------|
| **Ecosystem & Grants** | 30% | 300,000,000 | 4-year linear |
| **Team & Advisors** | 15% | 150,000,000 | 4-year with 1-year cliff |
| **Foundation Reserve** | 20% | 200,000,000 | Treasury-controlled |
| **Community Sale** | 10% | 100,000,000 | Unlocked at TGE |
| **Strategic Partners** | 10% | 100,000,000 | 2-year linear |
| **Liquidity Provision** | 10% | 100,000,000 | Unlocked at TGE |
| **Airdrop** | 5% | 50,000,000 | Unlocked at TGE |

### Staking Rewards

| Parameter | Value |
|-----------|-------|
| **Unbonding Period** | 21 days |
| **Minimum Commission** | 5% |
| **Maximum Validators** | 150 |
| **Slash Fraction (Downtime)** | 0.01% |
| **Slash Fraction (Double Sign)** | 5% |
| **Signed Blocks Window** | 10,000 blocks |

---

## ssUSD Stablecoin Economics

ssUSD is a **100%+ reserve-backed stablecoin** backed by **US Treasury Notes (T-Notes)**. On-chain mint/redeem (Path B) is backed by a tokenized T-Note reserve asset (default denom: `ustn`). Unlike algorithmic or crypto-collateralized stablecoins, ssUSD maintains strict backing with real-world reserves and enforces safety buffers via haircuts, fees, allocation caps, and daily limits.

### Reserve Backing Model

| Reserve Type | Description | Max Allocation |
|--------------|-------------|----------------|
| **US Treasury Notes (Tokenized)** | Tokenized T-Notes on-chain (`ustn`) | 100% |

### Approved Tokenized Treasury Tokens

| Token | Issuer | Underlying | Haircut | Max Allocation |
|-------|--------|------------|---------|----------------|
| **USTN** | OpenEden | US Treasury Notes (T-Notes) | 0.5% | 100% |

### Reserve Parameters

| Parameter | Value |
|-----------|-------|
| **Minimum Reserve Ratio** | 100% (fully backed) |
| **Target Reserve Ratio** | 102% (2% buffer) |
| **Mint Fee** | 0.1% |
| **Redemption Fee** | 0.1% |
| **Minimum Mint** | 100 ssUSD |
| **Minimum Redeem** | 100 ssUSD |
| **Daily Mint Limit** | 100,000,000 ssUSD |
| **Daily Redeem Limit** | 100,000,000 ssUSD |
| **Redemption Delay** | Instant (for tokenized) |
| **KYC Required** | Yes |

### How Minting Works

```
Mint ssUSD with Tokenized Treasuries:
1. User deposits approved tokenized Treasury Notes (`ustn`)
2. Oracle provides real-time price for `ustn`
3. Haircut applied for safety (0.5% default for `ustn`)
4. Mint fee (0.1%) deducted
5. ssUSD minted to user 1:1 with adjusted USD value
6. Deposited tokens held by stablecoin module as reserves

Example:
- Deposit: 10,000 USTN (valued at $10,000)
- Haircut (0.5%): -$50
- Net Value: $9,950
- Mint Fee (0.1%): -$9.95
- ssUSD Received: 9,940.05 ssUSD
```

### How Redemption Works

```
Redeem ssUSD for Reserves:
1. User submits ssUSD with desired output token (default: `ustn`)
2. System checks reserve availability (net of amounts locked by pending redemptions)
3. ssUSD is burned immediately at request time and the output amount is computed using the current oracle price (net of redemption fee)
4. The computed output amount is locked until execution so later redemptions cannot overbook reserves
5. If `redemption_delay > 0`, anyone can execute after the delay; execution transfers the locked output tokens to the user

Example:
- Submit: 10,000 ssUSD
- Redeem Fee (0.1%): -$10
- Net Redemption: $9,990
- Output: 9,990 USTN (at request-time oracle price)
```

### Proof of Reserves

ssUSD implements transparent proof-of-reserves through:

1. **On-Chain Reserves**: Tokenized treasuries held in module account, queryable in real-time
2. **Off-Chain Attestations**: Regular attestations from approved auditors
3. **Reserve Ratio Tracking**: Continuous monitoring of total reserves vs supply
4. **Public Dashboard**: Real-time reserve composition visibility

### Off-Chain Reserve Attestation

The protocol supports recording off-chain attestations for transparency. For the current Treasury-Notes-only model, attesters should report non-zero values in the T-Notes fields and keep other categories at zero unless governance expands supported reserve types.

| Field | Description |
|-------|-------------|
| **Total Cash** | USD held at custodian |
| **Total T-Bills** | US Treasury Bills value |
| **Total T-Notes** | US Treasury Notes value |
| **Total T-Bonds** | US Treasury Bonds value |
| **Custodian Name** | Primary custodian identity |
| **Audit Firm** | Independent auditor |
| **Report Date** | Attestation date |
| **Attestation Hash** | Document hash for verification |

---

## Legacy Vault Model (Deprecated)

The following vault-based model is deprecated but remains available for backwards compatibility:

### Collateral Types (Legacy)

| Collateral | Enabled | LTV | Liquidation | Stability Fee |
|------------|---------|-----|-------------|---------------|
| STATE | Deprecated | 66% | 130% | 2.0% |
| ATOM | Deprecated | 75% | 125% | 1.5% |

**Note**: The crypto-collateralized vault model has been deprecated in favor of the reserve-backed model for improved stability and regulatory compliance.

---

## Fee Distribution

All transaction fees collected by the network are distributed as follows:

| Recipient | Percentage | Purpose |
|-----------|------------|---------|
| **Validators** | 50% | Block production rewards |
| **Community Pool** | 25% | Governance-controlled grants |
| **Burn** | 25% | Deflationary pressure |

### Fee Types

| Operation | Fee |
|-----------|-----|
| Transfer | 0.001 STATE |
| Staking | 0.01 STATE |
| Vault Creation | 0.1 STATE |
| Settlement | 0.05% of amount |
| Compliance Check | 0.001 STATE |

---

## Treasury Management

### Budget Categories

| Category | Purpose | Period Limit |
|----------|---------|--------------|
| **Development** | Core protocol development | 10% of treasury/month |
| **Marketing** | Brand and awareness | 5% of treasury/month |
| **Grants** | Ecosystem grants | 15% of treasury/month |
| **Security** | Audits and bug bounties | 10% of treasury/month |
| **Operations** | Operational expenses | 5% of treasury/month |
| **Infrastructure** | Node operation | 5% of treasury/month |
| **Reserve** | Emergency reserve | No limit |

### Timelock Parameters

| Parameter | Value |
|-----------|-------|
| **Minimum Timelock** | 24 hours |
| **Maximum Timelock** | 30 days |
| **Proposal Expiry** | 7 days after timelock |
| **Max Pending Proposals** | 100 |

---

## Governance

### Proposal Types

| Type | Deposit | Voting Period | Quorum | Threshold |
|------|---------|---------------|--------|-----------|
| **Text** | 1,000 STATE | 7 days | 33.4% | 50% |
| **Parameter Change** | 5,000 STATE | 14 days | 40% | 66.7% |
| **Software Upgrade** | 10,000 STATE | 21 days | 50% | 66.7% |
| **Treasury Spend** | 5,000 STATE | 7 days | 33.4% | 50% |
| **Emergency** | 50,000 STATE | 3 days | 66.7% | 75% |

### Voting Power

Voting power is determined by staked STATE tokens:

- 1 STATE = 1 vote
- Delegated tokens inherit delegator's vote unless overridden
- Validators vote with full delegation weight

---

## Economic Security

### Circuit Breaker Triggers

| Trigger | Threshold | Action |
|---------|-----------|--------|
| **Price Deviation** | >5% in 1 block | Pause oracle updates |
| **Reserve Ratio** | <100% | Pause minting |
| **Daily Mint Limit** | >100M ssUSD/day | Pause minting |
| **Daily Redeem Limit** | >100M ssUSD/day | Rate limit redemptions |
| **Redemption Surge** | >10% TVL/hour | Pause redemptions |

### Oracle Security

| Parameter | Value |
|-----------|-------|
| **Minimum Providers** | 3 |
| **Price Aggregation** | Median |
| **Staleness Threshold** | 1 hour |
| **Max Deviation** | 5% |
| **Provider Slashing** | 1% stake per violation |

---

## Economic Projections

### Year 1-5 Projections

| Year | Inflation | New Supply | Total Supply | Bonded % |
|------|-----------|------------|--------------|----------|
| 1 | 15% | 150M | 1,150M | 55% |
| 2 | 12% | 138M | 1,288M | 62% |
| 3 | 9% | 116M | 1,404M | 67% |
| 4 | 7% | 98M | 1,502M | 70% |
| 5 | 7% | 105M | 1,607M | 68% |

### Fee Revenue Projections

| Year | Daily Tx | Avg Fee | Daily Revenue | Annual Revenue |
|------|----------|---------|---------------|----------------|
| 1 | 100K | 0.01 | 1,000 STATE | 365K STATE |
| 2 | 500K | 0.01 | 5,000 STATE | 1.8M STATE |
| 3 | 2M | 0.01 | 20,000 STATE | 7.3M STATE |
| 5 | 10M | 0.01 | 100,000 STATE | 36.5M STATE |

---

## Incentive Alignment

### Validators
- Earn block rewards + transaction fees
- Slashed for downtime/misbehavior
- Commission on delegator rewards

### Stakers
- Earn proportional inflation rewards
- Governance voting rights
- 21-day unbonding (security)

### ssUSD Users
- Stable medium of exchange
- Pay stability fees (protocol revenue)
- Liquidation risk (market discipline)

### Developers
- Grant funding from treasury
- Protocol fee share for dApps
- Ecosystem incentive programs

---

## Monetary Policy

### Deflationary Mechanisms
1. **Fee Burning**: 25% of all fees burned
2. **Stability Fees**: Collected and partially burned
3. **Liquidation Penalties**: Partially burned

### Inflationary Mechanisms
1. **Staking Rewards**: 7-20% annual inflation
2. **Community Incentives**: Treasury-funded grants

### Equilibrium Target
The protocol targets a long-term equilibrium where:
- Inflation = Burn rate + Economic growth
- Bonded ratio ≈ 67%
- ssUSD peg maintained at $1.00 ± 0.5%

---

## Parameter Governance

All economic parameters can be modified through governance:

| Parameter | Current | Range | Governance Type |
|-----------|---------|-------|-----------------|
| Min Collateral Ratio | 150% | 120-200% | Parameter Change |
| Stability Fee | 2% | 0-10% | Parameter Change |
| Inflation Bounds | 7-20% | 0-50% | Software Upgrade |
| Validator Count | 150 | 50-300 | Software Upgrade |
| Unbonding Period | 21 days | 7-42 days | Parameter Change |

---

## Appendix: Token Addresses

| Network | Token | Address |
|---------|-------|---------|
| Mainnet | STATE | `ustate` (native) |
| Mainnet | ssUSD | `ssusd` (native) |
| Testnet | STATE | `ustate` (native) |
| Testnet | ssUSD | `ssusd` (native) |

## Appendix: Key Contracts

| Module | Function | Gas Estimate |
|--------|----------|--------------|
| Stablecoin | DepositReserve | 150,000 |
| Stablecoin | RequestRedemption | 120,000 |
| Stablecoin | ExecuteRedemption | 100,000 |
| Stablecoin | RecordAttestation | 80,000 |
| Settlement | InstantTransfer | 80,000 |
| Treasury | ProposeSpend | 100,000 |
| Treasury | ExecuteSpend | 150,000 |
| Oracle | SubmitPrice | 50,000 |

### Reserve Message Types

| Message | Description |
|---------|-------------|
| `MsgDepositReserve` | Deposit tokenized US Treasury Notes (`ustn`) to mint ssUSD |
| `MsgRequestRedemption` | Request redemption of ssUSD for reserves |
| `MsgExecuteRedemption` | Execute a pending redemption |
| `MsgCancelRedemption` | Cancel a pending redemption (authority only) |
| `MsgUpdateReserveParams` | Update reserve parameters (governance) |
| `MsgRecordAttestation` | Record off-chain reserve attestation |
| `MsgSetApprovedAttester` | Approve/revoke reserve attesters |
