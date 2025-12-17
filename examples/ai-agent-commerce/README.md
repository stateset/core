# AI Agent Commerce with `ssusd` (On-Chain)

This example shows an **AI buyer agent** checking out and paying an **AI seller agent** using the native `ssusd` stablecoin **on-chain**.

It demonstrates three on-chain flows:

1. **Instant checkout** (`x/settlement` `instant-checkout`)
2. **Escrow checkout + release** (`x/settlement` `instant-checkout --use-escrow` + `release-escrow`)
3. **Service purchase via payment intent** (`x/payments` `create` + seller `settle`)

## Quick Start (Local Single-Node Chain)

Prereqs:
- `jq`
- a built binary (recommended): `make build` (creates `./build/statesetd`)

Run:

```bash
bash examples/ai-agent-commerce/shell/agent-to-agent-checkout.sh
```

Keep the temporary chain state for inspection:

```bash
KEEP_HOME=1 bash examples/ai-agent-commerce/shell/agent-to-agent-checkout.sh
```

## What To Look At On-Chain

The script prints IDs and then queries state:

- Settlements: `statesetd query settlement settlement <id>`
- Payments: `statesetd query payments payment <id>`
- Balances: `statesetd query bank balances <addr>`

## Denomination

`ssusd` uses 6 decimals (base units):
- `1 ssUSD = 1_000_000ssusd`
