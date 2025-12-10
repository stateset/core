# Fee Market & Gas Oracle

This document defines the fee market and gas oracle design to deliver predictable fees, congestion pricing, and observable gas price guidance for Stateset. It targets an EIP-1559–style base fee plus priority fee model with Cosmos SDK integration.

## Goals
- Dynamic base fee that rises under congestion and falls when blocks are under-utilized.
- Clear separation of `base_fee` (burned or protocol revenue) and `priority_fee` (tip to validators).
- Gas oracle endpoint that surfaces current `base_fee`, suggested `priority_fee`, and min gas price guidance.
- Safety rails: min/max base fee, per-block change bounds, pausable circuit breaker.

## Components

### Parameters
- `enabled` (bool): toggle fee market logic.
- `base_fee_change_denominator` (uint64): dampens per-block change (typical: 8–16).
- `elasticity_multiplier` (uint64): expands virtual block gas target (typical: 2).
- `target_gas` (uint64): gas target per block before elasticity (if zero, default to `max_block_gas / elasticity_multiplier`).
- `min_base_fee` (Dec): floor for base fee (default: current `min_gas_price` 0.025 STATE).
- `max_base_fee` (Dec): optional ceiling (0 means no ceiling).
- `initial_base_fee` (Dec): bootstrap value used at genesis if no prior base fee exists.
- `priority_fee_floor` (Dec): minimum tip suggested by oracle.
- `max_fee_history` (uint32): how many blocks to keep in history for oracle percentiles.

### State
- `base_fee` (Dec): stored per block and updated in EndBlock.
- `fee_history`: rolling window of recent block gas used and priority fees to compute oracle outputs.

### Base Fee Update (EndBlock)
Inputs: `gas_used`, `target_gas`, `base_fee_change_denominator`, `elasticity_multiplier`.

1. Compute `target = target_gas * elasticity_multiplier` (or `max_block_gas / elasticity_multiplier` if target is unset).
2. Let `delta = gas_used - target`.
3. `change = base_fee * delta / target / base_fee_change_denominator`.
4. `next_base_fee = clamp(base_fee + change, min_base_fee, max_base_fee_if_set)`, never negative.
5. Persist `next_base_fee` and append `(gas_used, avg_priority_fee)` to `fee_history`.

### Transaction Fee Check (Ante)
- Required fee: `required = base_fee * gas_limit + priority_fee * gas_limit`.
- If fee market is enabled, reject tx if supplied fee < `base_fee * gas_limit` (priority fee may be zero but will result in slower inclusion).
- Priority fee is user-selected; oracle only suggests a value.

### Gas Oracle
Exposes:
- `current_base_fee`: latest base fee.
- `recommended_priority_fee`: percentile-based suggestion (e.g., 50th/75th/90th percentile of last N blocks).
- `min_gas_price`: `current_base_fee` + `priority_fee_floor` (for UX and relayers).
- `fee_history`: optional recent history for wallets/indexers.

### Safety & Operations
- Circuit breaker flag: pause base fee updates (freeze at last base fee) if enabled.
- Governance/authority can update params via `MsgUpdateParams`.
- Limits: cap per-block change via `base_fee_change_denominator` and `max_base_fee`; `min_base_fee` prevents fee collapse.
- Telemetry: export gauges for base fee, priority fee percentiles, gas used vs. target, fee revenue.

## Integration Steps (follow-up)
1) Add `x/feemarket` module wiring in `app/app.go`:
   - Store keys, keeper, module registration, `BeginBlocker/EndBlocker` hook (base fee update in EndBlock).
   - Add ante decorator to enforce `base_fee` requirement and record priority fee tips.
2) CLI/GRPC:
   - Queries: `base_fee`, `oracle/suggested_priority_fee`, `params`, `fee_history`.
   - Tx: `update-params` (authority).
3) Genesis:
   - Include params and `initial_base_fee` in genesis; default to 0.025 STATE.
4) Observability:
   - Add metrics to the existing `x/metrics` module or Prometheus exporter (base fee, gas used, target, percentiles).
5) Testing:
   - Unit tests for base fee math (edge cases: empty block, max change, floor/ceiling).
   - Ante handler tests for fee enforcement.
   - Integration test with varying block gas to see convergence to target.

## Defaults (proposed)
- `enabled`: true
- `base_fee_change_denominator`: 8
- `elasticity_multiplier`: 2
- `target_gas`: 10,000,000 (adjust to match block gas limit)
- `initial_base_fee`: 0.025 STATE
- `min_base_fee`: 0.025 STATE
- `max_base_fee`: 0 (unset)
- `priority_fee_floor`: 0.0001 STATE
- `max_fee_history`: 100
