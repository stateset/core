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
- Authority-based price updates

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
| `MsgUpdatePrice` | Update price for a denomination |
| `MsgRegisterProvider` | Register new oracle provider |
| `MsgRemoveProvider` | Remove oracle provider |

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
| `staleness_threshold` | 1 hour | Price staleness limit |
| `min_update_interval` | 1 minute | Minimum time between updates |
| `deviation_threshold_atom` | 5% | ATOM price deviation limit |
| `deviation_threshold_stable` | 1% | Stablecoin deviation limit |
| `daily_deviation_cap` | 20% | Maximum daily price change |

## State

| Key | Value |
|-----|-------|
| `0x01{denom}` | Price |
| `0x02{provider}` | ProviderInfo |
| `0x03{denom}` | OracleConfig |
| `0x04{denom}` | PriceHistory |

## Events

| Event | Attributes |
|-------|------------|
| `price_updated` | denom, price, timestamp |
| `provider_registered` | address, name |
| `provider_removed` | address |
