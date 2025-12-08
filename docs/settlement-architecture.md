# Stateset Settlement Chain Architecture

## Vision

Design and operate a Stateset blockchain network that can securely settle billions of dollars in USD–denominated stablecoin transactions. The chain must combine institutional-grade risk controls, deterministic settlement guarantees, rich compliance tooling, and extensibility for intelligent commerce agents.

---

## Platform Foundations

- **Cosmos SDK Base**: Target SDK v0.52+ (ABCIDev), CometBFT 0.39+, ADR-054 wiring, ADR-068 keyring, ABCI++ for low-latency finality.
- **Consensus & Networking**: 3s block times, BFT validator set ≥50 with geographical and jurisdictional dispersion; gossip tuned for 5k TPS.
- **State Management**: LevelDB/Pe… writes optimized with IAVL fast sync and state pruning strategy; parallel tx execution roadmap using mempool prioritisation.

---

## Core Modules

| Module            | Purpose                                                                        | Key Features                                                                                         |
|-------------------|--------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `x/stablecoin`    | Issuance/redemption and risk controls for ssUSD & future denominations        | Multi-collateral vaults, stability fees, liquidation auctions, oracle-driven mark-to-market          |
| `x/payments`      | High-throughput settlement rails for B2B/B2C commerce                         | Netting engine, multi-leg settlements, programmable routing, circuit breaker for anomalous flows     |
| `x/treasury`      | Reserve proofs and fiat off-ramp integration                                  | Proof-of-reserves attestations, omnibus bank account reconciliation, fiat mint/burn pipelines        |
| `x/compliance`    | KYC/AML enforcement and reporting                                             | Travel-rule exchange, sanctions/KYC checks, suspicious-activity detection, regulator data exports    |
| `x/oracle`        | Aggregates verified price feeds for collateral and FX                         | Pulls from Pyth, Chainlink, and institutional APIs; medianized feed persistence with slashable oracles|
| `x/agent`         | Intelligent commerce agents orchestration                                     | Agent registration, policy controls, delegated execution limits, transaction intent attestations     |

Additional modules: `x/liquidity` (AMMs for treasury rebalancing), `x/risk` (stress testing, VaR metrics), `x/analytics` (BI dashboards).

---

## Stablecoin Engine (`x/stablecoin`)

### Vault Structure
- **Vault**: `id`, `owner`, `collateral_denom`, `collateral_amount`, `debt_amount`, `last_interest_accrual_height`.
- **Collateral Types**: Configurable params (min collateral ratio, stability fee, liquidation penalty, debt ceiling).
- **Accounting**: Use `sdk.Dec` for ratios, `sdk.Int` for supply.

### Flows
1. **Deposit Collateral** (`MsgDepositCollateral`)
   - Transfer tokens to module account via `bank` keeper.
   - Update vault collateral amount.
2. **Mint Stablecoin** (`MsgMintStablecoin`)
   - Enforce collateral ratio using marked price from oracle.
   - Mint ssUSD to owner (module account as minter).
3. **Repay & Withdraw** (`MsgRepayStablecoin`, `MsgWithdrawCollateral`)
   - Burn ssUSD, adjust debt, release collateral if ratios healthy.
4. **Liquidation** (`MsgTriggerLiquidation`)
   - Anyone can flag undercollateralized vaults.
   - Create `Auction` record consumed by `x/auction`.
5. **Global Shutdown / Circuit Breaker**
   - Governance flag halts mint/burn; ensures emergency unwind.

### Data Storage (collections)
```go
vaults := collections.NewMap[VaultID, Vault](...),
collateralParams := collections.NewMap[Denom, CollateralParam](...),
oracles := collections.NewMap[Denom, PriceFeed](...)
```

### Inter-module dependencies
- `bank` for token transfers & minting.
- `oracle` for price data.
- `treasury` for off-chain reserve proofs (fiat-backed wrappers).
- `compliance` for address risk scoring gating operations.

---

## Payments & Settlement (`x/payments`)

### Features
- **Netting Engine**: Batch settle obligations across merchants, reducing on-chain load.
- **Escrow Flows**: Multi-signature release with compliance hooks.
- **Payment Intents**: Metadata describing invoice, counterparties, compliance status.
- **Bulk Settlement**: `MsgSettleBatch` enabling large batches validated off-chain then committed on-chain.
- **Circuit Breakers**: Limits per merchant / per denom; integrates with `x/risk`.

### Data Structures
```go
PaymentIntent { id, payer, payee, amount, settlement_window, metadata }
Batch { id, intents[], status, netted_transfers[] }
```

---

## Compliance Layer (`x/compliance`)

### Capabilities
- **Party Registry**: Stores KYC levels, jurisdiction tags, sanctions status.
- **Transaction Screening**: Hooks to evaluate intents before execution.
- **Reporting**: Scheduled generation of SARs, travel-rule data.
- **Policy Management**: Governance sets thresholds; compliance officers operate via privileged accounts.

### Execution Path
1. `x/payments` and `x/stablecoin` call `complianceKeeper.AssertCompliant(ctx, policy, actors, payload)`.
2. Keeper checks cached risk scores, optionally queries external oracles via ICA/ICQ.
3. On failure, tx aborted with descriptive error.

---

## Treasury & Proof-of-Reserves (`x/treasury`)

- **Fiat Mint/Burn Requests**: Create/approve flows referencing off-chain banking rails.
- **Reserve Snapshots**: Daily attestations (auditor signed) anchored in chain state.
- **Omnibus Reconciliation**: Track bank balances vs issued supply; alert on mismatches.
- **Integration Hooks**: Provide data for `x/risk` metrics and compliance audits.

---

## Oracle Aggregation (`x/oracle`)

1. Collect price feeds via IBC (e.g., Pyth, UMA) and direct REST connectors.
2. Medianize and post to KV store with timestamp.
3. Slash oracle providers submitting bad data (requires staking bond).
4. Provide streaming service for off-chain risk engines. 

---

## Governance & Risk Management

- **Parameter DAO**: On-chain governance manages risk params; heavy changes gated by multi-sig council.
- **Emergency Response**: Liquidity committee with fast-track authority to pause modules.
- **Risk Engine**:
  - Monitor LTV distribution, pending liquidations, liquidity coverage ratio.
  - Run stress tests via simulation orchestrators (off-chain) with results anchored on-chain.

---

## Infrastructure & DevOps

- **Build**: `make build`, `make proto`, `make test`, `make test-sim`.
- **Testing**:
  - Unit tests per module.
  - Simulation tests for liquidation scenarios & high-volume settlement.
  - Fuzzing of Msg/Query.
  - Load tests with realistic payment batches.
- **Observability**:
  - Prometheus metrics for settlement latency, vault health, compliance flags.
  - OpenTelemetry tracing for payment flows.
  - SLAs: 99.95% uptime, <10s finality, <0.1% failed settlements.
- **Security**:
  - Formal verification roadmap for vault math.
  - Regular external audits (app + contracts).
  - Bug bounty focused on liquidation / reserve discrepancy vulnerabilities.

---

## Delivery Roadmap

1. **Foundation Sprint (Weeks 1-4)**  
   - Upgrade to Cosmos SDK 0.52.  
   - Implement scaffolding for `x/stablecoin`, `x/payments`, `x/compliance`, `x/treasury`.  
   - Establish module accounts, params, genesis defaults.

2. **Stablecoin Core (Weeks 5-10)**  
   - Complete vault lifecycle (deposit, mint, repay, withdraw).  
   - Implement oracle module integration.  
   - Add liquidation queue and auctions (baseline).

3. **Compliance & Treasury Integration (Weeks 8-12)**  
   - Add KYC registry, travel-rule messaging.  
   - Implement fiat proof-of-reserve ingest pipeline.  
   - Build admin dashboards (off-chain) using gRPC queries.

4. **Payments & Scaling (Weeks 11-16)**  
   - Netting engine, batch settlement, escrow flows.  
   - Circuit breakers + risk metrics.  
   - Performance tuning with >5k TPS load testing.

5. **Hardening & Launch Prep (Weeks 16-20)**  
   - Security audits, chaos testing, observability hardening.  
   - Public incentivized testnet, key management drills.  
   - Governance bootstrap and final go/no-go for mainnet.

---

## Next Actions

1. Confirm target SDK/CometBFT versions and dependency lock.  
2. Scaffold modules (Starport / Ignite or manual) with proto definitions.  
3. Implement `x/stablecoin` vault data structures and CLI/REST/GRPC endpoints.  
4. Wire new modules into `app/app.go` with module accounts & invariants.  
5. Develop simulation tests covering collateral stress cases.  
6. Stand up CI pipeline (lint, unit, simulation, integration).  
7. Engage compliance and treasury partners for schema integration.  
8. Draft governance charter and emergency response playbooks.

---

## Considerations for Multi-Billion Scale

- **Regulatory**: Support for segregated ledgers per jurisdiction, on-chain attestation for compliance.  
- **Interoperability**: Native IBC channels to major Cosmos zones, bridged USDC via CCTP, Ethereum interoperability with ICS-20 wrappers.  
- **Resilience**: Multi-region validator infrastructure with managed failover, hardware security modules for keys.  
- **Upgradability**: Use `x/upgrade` smooth migrations, feature flags for staged rollouts.  
- **Data Privacy**: Explore shielded transactions for sensitive B2B flows (zk proofs on roadmap).

---

## Appendix: Module Accounts

| Module         | Account                               | Permissions                    |
|----------------|----------------------------------------|--------------------------------|
| Stablecoin     | `stablecoin_macc`                     | `Minter`, `Burner`             |
| Payments       | `payments_macc`                       | `Minter`, `Burner`, `Staking` (for fee routing) |
| Treasury       | `treasury_macc`                       | `Minter`, `Burner`             |
| Compliance     | `compliance_macc` (escrowed penalties)| `Burner`                       |
| Oracle         | `oracle_macc`                         | `None` (fee distribution only) |

---

This document sets the architectural baseline. Next iterations will add protobuf schemas, keeper implementations, and integration tests to realize the blueprint.
