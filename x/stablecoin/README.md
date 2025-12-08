# Stablecoin Module

The Stablecoin module manages the native stablecoin (ssUSD) for the Stateset blockchain, including vault-based collateral management.

## Overview

The stablecoin module handles:
- **Vault Management**: Create and manage collateral vaults
- **Collateral Operations**: Deposit and withdraw collateral
- **Minting/Burning**: Mint and burn stablecoins
- **Reserve Management**: Track collateralization ratios

## Features

### Vault-Based Collateral
- Individual vaults per user
- Collateralized debt positions
- Over-collateralization requirements

### Stablecoin Operations
- Mint ssUSD against collateral
- Burn ssUSD to reduce debt
- Automatic reserve ratio tracking

### Oracle Integration
- Real-time collateral valuation
- Price-based liquidation triggers

## Messages

| Message | Description |
|---------|-------------|
| `MsgCreateVault` | Create new collateral vault |
| `MsgDeposit` | Add collateral to vault |
| `MsgWithdraw` | Remove collateral from vault |
| `MsgMint` | Mint stablecoins |
| `MsgBurn` | Burn stablecoins |
| `MsgLiquidate` | Liquidate undercollateralized vault |

## Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `min_collateral_ratio` | 150% | Minimum collateralization |
| `liquidation_ratio` | 120% | Liquidation trigger ratio |
| `stability_fee` | 2% | Annual stability fee |
| `liquidation_penalty` | 13% | Liquidation penalty |

## State

| Key | Value |
|-----|-------|
| `0x01{id}` | Vault |
| `0x02` | NextVaultID |
| `0x03` | TotalSupply |
| `0x04` | Params |

## Events

| Event | Attributes |
|-------|------------|
| `vault_created` | id, owner |
| `collateral_deposited` | vault_id, amount |
| `stablecoin_minted` | vault_id, amount |
| `stablecoin_burned` | vault_id, amount |
| `vault_liquidated` | vault_id, debt, penalty |
