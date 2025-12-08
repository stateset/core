# CLI Guide

This guide covers the core `statesetd` transaction and query commands now available for the rebuilt modules. Every command assumes you have a funded key in your local keyring (for example `--from <name>`).

---

## Compliance Module

### Transactions

```bash
statesetd tx compliance upsert-profile <address> <kyc-level> <risk-level> \
  --metadata "optional note" --sanction=false --from <admin>

statesetd tx compliance set-sanction <address> <true|false> \
  --reason "brief explanation" --from <admin>
```

### Queries

```bash
statesetd query compliance profile <address>
```

---

## Payments Module

### Transactions

```bash
statesetd tx payments create <payee-address> <amount> \
  --metadata "invoice-42" --from <payer>

statesetd tx payments settle <payment-id> --from <payee>

statesetd tx payments cancel <payment-id> \
  --reason "duplicate invoice" --from <payer>
```

### Queries

```bash
statesetd query payments payment <payment-id>
```

---

## Stablecoin Module

### Transactions

```bash
# Create a vault and optionally mint debt in one step.
statesetd tx stablecoin create-vault <collateral> \
  --debt <optional-debt> --from <owner>

statesetd tx stablecoin deposit <vault-id> <collateral> --from <owner>

statesetd tx stablecoin withdraw <vault-id> <collateral> --from <owner>

statesetd tx stablecoin mint <vault-id> <amount> --from <owner>

statesetd tx stablecoin repay <vault-id> <amount> --from <owner>

statesetd tx stablecoin liquidate <vault-id> --from <liquidator>
```

Liquidation automatically deducts the vault's entire `ssusd` debt from the calling
account before collateral is released, so be sure the liquidator holds enough
stablecoins to cover the position.

### Queries

```bash
statesetd query stablecoin vault <vault-id>
```

---

## Oracle Module

### Transactions

```bash
statesetd tx oracle update-price <denom> <price-decimal> --from <authority>
```

### Queries

```bash
statesetd query oracle price <denom>
```

---

## Treasury Module

### Transactions

```bash
statesetd tx treasury record-reserve <total-supply-coin> <fiat-reserve-coin> \
  --other-reserves "10eth,50000usdc" \
  --metadata "March 2025 audit" \
  --timestamp "2025-03-15T00:00:00Z" \
  --from <treasury-admin>
```

### Queries

```bash
statesetd query treasury snapshot <snapshot-id>
```

---

### Tips

- All transaction commands support standard Cosmos SDK flags (fees, gas, memo, etc.).
- Query commands accept `--height` to inspect historical state.
- Use `--help` on any subcommand for flag details.

Keep this guide handy while operating the chain to avoid hunting through source files for command syntax.
