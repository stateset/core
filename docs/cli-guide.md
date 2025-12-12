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

## Settlement Module

### Transactions

```bash
# Instant transfer
statesetd tx settlement instant-transfer <recipient> <amount> <reference> --from <sender>

# Escrow flow
statesetd tx settlement create-escrow <recipient> <amount> <expires-in> --from <sender>
statesetd tx settlement release-escrow <settlement-id> --from <sender>
statesetd tx settlement refund-escrow <settlement-id> --from <recipient>

# Payment channels
statesetd tx settlement open-channel <recipient> <deposit> <expires-in-blocks> --from <sender>
statesetd tx settlement claim-channel <channel-id> <amount> <nonce> <signature> --from <recipient>
statesetd tx settlement close-channel <channel-id> --from <sender>

# Merchant config
statesetd tx settlement register-merchant <name> --fee-rate-bps 25 --batch-enabled \
  --min-settlement 10ssusd --max-settlement 100000ssusd \
  --webhook-url https://example.com/webhook --from <merchant>
statesetd tx settlement update-merchant <merchant-address> --name "New Name" --from <merchant>

# Batch settlements (authority-only; comma-separated lists)
statesetd tx settlement create-batch <merchant> <sender1,sender2> <100ssusd,200ssusd> <ref1,ref2> --from <authority>
statesetd tx settlement settle-batch <batch-id> --from <authority>
```

### Queries

```bash
statesetd query settlement settlement <id>
statesetd query settlement batch <id>
statesetd query settlement channel <id>
statesetd query settlement merchant <merchant-address>
statesetd query settlement params
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

Vault-based minting is disabled by default (`vault_minting_enabled=false`) to keep ssUSD strictly reserve‑backed. Governance must enable it before the vault commands will succeed.

Liquidation automatically deducts the vault's entire `ssusd` debt from the calling
account before collateral is released, so be sure the liquidator holds enough
stablecoin to cover the position.

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
