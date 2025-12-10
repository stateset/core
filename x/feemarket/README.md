# Fee Market Module

Implements an EIP-1559–style fee market with dynamic `base_fee` and a simple gas oracle for priority fee suggestions.

## Key Features
- Dynamic base fee updated in `EndBlock` based on gas used vs. target.
- Parameters for change rate, elasticity, min/max base fee, and history length.
- Gas oracle inputs (latest gas used, base fee) for wallet/relayer guidance.
- Authority-controlled params (governance/upgrade handler).

## State
- `params`: fee market parameters.
- `base_fee`: current base fee per gas.
- `latest_gas`: last block gas used (for quick oracle responses).

## Next Steps (wiring)
- Register module in `app/app.go` with keeper and store keys.
- Add ante decorator to enforce `base_fee` per gas.
- Expose queries (base fee, suggested priority fee, params) and an `update-params` tx.
- Add metrics (base fee, gas used, target) to Prometheus/`x/metrics`.
