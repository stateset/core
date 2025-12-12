# Oracle Module

The Oracle module provides decentralized price feed management and validation for the Stateset blockchain.

## Overview

The oracle module handles:
- **Price Storage**: Store and retrieve asset prices
- **Provider Management**: Register and manage oracle providers
- **Price Validation**: Validate prices against thresholds
- **Price History**: Maintain historical price data

## Features

### Price Feeds
- Multi-denomination price support
- Timestamp-based staleness detection
- Historical price tracking
- Authorized-provider price updates (governance authority or registered providers)
- Optional multi-provider confirmation with median aggregation (`required_confirmations > 1`)

### Provider Management
- Provider registration and deregistration
- Success/failure tracking
- Provider performance metrics

### Price Validation
- Positive price enforcement
- Deviation threshold checking
- Staleness detection

## Messages

| Message | Description |
|---------|-------------|
| `MsgUpdatePrice` | Submit a price update for a denomination (validated and optionally aggregated) |

Provider and config management are currently performed via genesis/governance upgrades. `MsgRegisterProvider` / `MsgRemoveProvider` are roadmap items.

## Queries

| Query | Description |
|-------|-------------|
| `Price` | Get current price for denomination |
| `Prices` | Get all current prices |
| `Provider` | Get provider information |
| `Providers` | List all providers |
| `PriceHistory` | Get historical prices |

## Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `default_max_deviation_bps` | 500 | Default max deviation per update |
| `default_staleness_threshold` | 3600s | Default staleness threshold |
| `slash_fraction_bps` | 1000 | Oracle slashing fraction (bps) |
| `max_providers` | 10 | Maximum registered providers |
| `price_history_size` | 100 | History points kept per denom |

Per-denom `OracleConfig` overrides defaults:
`max_deviation_bps`, `staleness_threshold_seconds`, `min_update_interval_seconds`, `required_confirmations`, `enabled`.

## State

| Key | Value |
|-----|-------|
| `0x01{denom}` | Price |
| `0x02{provider}` | ProviderInfo |
| `0x03{denom}` | OracleConfig |
| `0x04{denom}` | PriceHistory |
| `pending:{denom}` | Pending submissions for aggregation |

## Events

| Event | Attributes |
|-------|------------|
| `price_updated` | denom, price, timestamp |
| `provider_registered` | address, name |
| `provider_removed` | address |
