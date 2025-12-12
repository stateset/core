# Stateset Modules

This repository contains the following custom Cosmos SDK modules:

- `x/stablecoin`: `ssUSD` stablecoin engine (reserve‑backed by default, optional CDP vaults gated by governance).
- `x/settlement`: Stablecoin settlement rails (instant transfers, escrow, batch settlement, payment channels).
- `x/payments`: Payment intent lifecycle for commerce flows.
- `x/orders`: On‑chain order lifecycle management.
- `x/compliance`: KYC/AML profiles, sanctions, limits, jurisdiction rules.
- `x/treasury`: Protocol treasury and reserve attestation state.
- `x/oracle`: Price feeds with staleness/deviation checks.
- `x/feemarket`: EIP‑1559 style fee market and gas oracle.
- `x/circuit`: Circuit breakers, global pause, rate limiting, liquidation surge protection.
- `x/metrics`: On‑chain metrics and Prometheus export.
- `x/zkpverify`: On‑chain zero‑knowledge proof verification.

Legacy business modules referenced in older docs (agreement, invoice, loan, etc.) are not part of this repo; see `docs/10-10-execution-plan.md` and `IMPROVEMENT_ROADMAP.md` for the current roadmap.
