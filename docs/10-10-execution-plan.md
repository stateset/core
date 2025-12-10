# 10/10 Execution Plan

Fastest path to raise Stateset to a 10/10 across security, decentralization, performance, developer ecosystem/UX, cost, governance, and observability. Each track includes the target outcome and immediate actions to start execution.

## Cross-Track Prerequisites
- Upgrade local and CI toolchains to Go 1.23+ (go.mod uses `toolchain go1.23.4`), then re-run `go test ./...`.
- Turn on nightly builds: `go test ./...` with `-count=1`, coverage upload, and lint/static analysis.
- Freeze secrets: no validator keys or `~/.statesetd` data in commits; keep configs in `config/` templates.

## Security & Safety
- Target: External audit coverage of critical modules, >80% test coverage on safety-critical code, CosmWasm re-enabled safely.
- Immediate actions:
  - Scope an external audit per `docs/SECURITY.md` (stablecoins, oracle, compliance, governance, circuit breakers).
  - Add fuzz/property tests for liquidation bounds, oracle deviation/staleness, and compliance gates (start with `x/stablecoins`, `x/commerce`, `x/compliance`).
  - Re-enable CosmWasm when a v0.53-compatible wasmd lands; run migration tests and circuit-breaker regressions before enabling in `app/app.go`.
  - Stand up a public bug bounty after audit handoff.

## Performance & Scalability
- Target: >1k TPS sustained, <6s block time, <10s finality under load.
- Immediate actions:
  - Stand up a public testnet with telemetry (Prometheus/Grafana) and run load via tm-load-test or custom tx generator; publish dashboards.
  - Tune CometBFT/SDK params (mempool size, p2p limits, block gas/tx limits, pruning) based on load data.
  - Add perf regression jobs (targeted module load tests) to CI on tagged builds.

## Decentralization & Validator Program
- Target: 100+ independent validators, >95% participation, healthy peer diversity.
- Immediate actions:
  - Publish validator onboarding/runbook (hardware, config, monitoring) and invite operators for testnet.
  - Set and document `min_commission`, slashing params, and peer set targets; monitor with alerts on participation and voting power concentration.
  - Run testnet incentive program and publish a live validator dashboard.

## Economics & Cost (Fees/Stablecoins)
- Target: Predictable low fees and transparent stablecoin economics.
- Immediate actions:
  - Implement and test the fee market module (EIP-1559-style base fee + priority fee); wire a gas price oracle and expose fee metrics.
  - Validate stablecoin fees/penalties match `MAINNET_READINESS.md`; add simulations for liquidation flows and stability pool health.
  - Publish observed fee stats from testnet alongside min gas price recommendations.

## Governance & Upgrade Safety
- Target: Safe, transparent upgrades with accountable controls.
- Immediate actions:
  - Add governance timelocks and emergency controls (module-level circuit breakers already described) with regression tests.
  - Exercise an end-to-end software upgrade on testnet (proposal → vote → halt → migrate → resume) and document the runbook.
  - Enable emergency multisig for critical actions; document thresholds and custody guidance.

## Developer Experience & UX
- Target: Simple local dev and clear user flows.
- Immediate actions:
  - Expand `docs/cli-guide.md` with end-to-end flows (mint/burn ssUSD, agreement lifecycle, compliance gating).
  - Publish SDK snippets/examples (Go/TypeScript) for common tx/query patterns; add fixtures to `make dev` for faster demos.
  - Add explorer/wallet integration guidance and UX notes for KYC/AML flows.

## Observability & Ops
- Target: 24/7 visibility with actionable alerts.
- Immediate actions:
  - Deploy Prometheus/Grafana with dashboards for block timing, mempool, fees, oracle freshness, liquidations, compliance blocks.
  - Set alerting thresholds (from `docs/SECURITY.md`) and on-call runbooks; verify log aggregation and backups.
  - Add SLOs (availability, finality, API latency) and report them for each testnet release.

## Timeline (suggested)
- Weeks 0-2: Go toolchain upgrade, nightly CI, audit scoping, initial fuzz tests, public testnet bring-up with telemetry.
- Weeks 3-6: Fee market implementation, perf tuning with published metrics, governance timelock + emergency controls, SDK/examples.
- Weeks 7-10: External audit + remediation, bug bounty launch, validator incentive phase, end-to-end upgrade drill.
- Weeks 11-12: Freeze for mainnet readiness with published KPIs, SLOs, and operator guides.
