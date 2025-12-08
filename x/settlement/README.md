# Settlement Module

The Settlement module provides comprehensive stablecoin payment infrastructure for the Stateset blockchain, enabling instant transfers, escrow settlements, batch payments, and payment channels.

## Overview

The settlement module handles:
- **Instant Transfers**: Immediate stablecoin transfers with settlement records
- **Escrow Settlements**: Time-locked funds with release/refund capabilities
- **Batch Settlements**: Aggregated payments to merchants with fee optimization
- **Payment Channels**: State channels for high-frequency micropayments
- **Merchant Management**: Configuration for merchants including custom fee rates

## Features

### Instant Transfers
Immediate settlement of stablecoin payments between parties with:
- Compliance checks for both sender and recipient
- Configurable fee rates (basis points)
- Settlement records for audit/reporting

### Escrow Settlements
Hold funds in escrow until conditions are met:
- Sender can release funds to recipient
- Recipient can initiate refund back to sender
- Automatic expiration handling (refunds to sender)
- Configurable expiration times

### Batch Settlements
Aggregate multiple payments for efficiency:
- Multiple senders to single merchant
- Reduced transaction overhead
- Bulk fee calculation
- Authority-controlled settlement execution

### Payment Channels
Off-chain payment capabilities:
- Open channels with deposit
- Claim funds with signed authorizations
- Replay protection via nonces
- Expiration-based closure

### Merchant Configuration
Custom settings per merchant:
- Fee rates (basis points, max 100%)
- Settlement limits (min/max)
- Batch settlement thresholds
- Webhook notifications (HTTPS required)

## Messages

| Message | Description |
|---------|-------------|
| `MsgInstantTransfer` | Execute immediate stablecoin transfer |
| `MsgCreateEscrow` | Create time-locked escrow settlement |
| `MsgReleaseEscrow` | Release escrowed funds to recipient |
| `MsgRefundEscrow` | Refund escrowed funds to sender |
| `MsgCreateBatch` | Create batch of settlements |
| `MsgSettleBatch` | Execute batch settlement |
| `MsgOpenChannel` | Open payment channel with deposit |
| `MsgCloseChannel` | Close channel after expiration |
| `MsgClaimChannel` | Claim funds from channel with signature |
| `MsgRegisterMerchant` | Register new merchant configuration |
| `MsgUpdateMerchant` | Update merchant settings |

## Queries

| Query | Description |
|-------|-------------|
| `Settlement` | Get settlement by ID |
| `Settlements` | List all settlements with pagination |
| `SettlementsByStatus` | Filter settlements by status |
| `Batch` | Get batch by ID |
| `Batches` | List all batches |
| `Channel` | Get payment channel by ID |
| `Channels` | List all channels |
| `ChannelsByParty` | Get channels for address |
| `Merchant` | Get merchant configuration |
| `Merchants` | List all merchants |
| `Params` | Get module parameters |

## Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `default_fee_rate_bps` | uint32 | Default fee rate in basis points |
| `min_settlement_amount` | Coin | Minimum allowed settlement |
| `max_settlement_amount` | Coin | Maximum allowed settlement |
| `default_escrow_expiration` | int64 | Default escrow duration (seconds) |
| `max_escrow_expiration` | int64 | Maximum escrow duration |
| `fee_collector` | string | Address to receive collected fees |

## Security

### Input Validation
- All addresses validated via `AccAddressFromBech32`
- Amount validation (positive, valid denomination)
- Fee rate bounds (0-10000 bps)
- Webhook URL validation (HTTPS only, no private IPs)

### Signature Verification
Payment channel claims require cryptographic signatures:
- Message format: `channel_claim:{channelId}:{recipient}:{amount}:{nonce}`
- Prevents unauthorized withdrawals
- Replay protection via incrementing nonces

### Compliance Integration
- Mandatory compliance checks for all parties
- Integration with compliance module for KYC/AML
- Transaction blocked if parties are sanctioned/blocked

### Webhook Security
- HTTPS required for all webhook URLs
- Private IP ranges blacklisted (localhost, 10.x, 172.x, 192.168.x)
- Link-local addresses blocked

## Events

| Event | Attributes |
|-------|------------|
| `settlement_created` | settlement_id, sender, recipient, amount, status |
| `settlement_completed` | settlement_id, recipient, amount, fee |
| `settlement_refunded` | settlement_id, sender, amount |
| `batch_created` | batch_id, merchant, amount |
| `batch_settled` | batch_id, merchant, amount, fee |
| `instant_transfer` | settlement_id, sender, recipient, amount, fee, reference |
| `channel_opened` | channel_id, sender, recipient, amount |
| `channel_closed` | channel_id, amount |
| `channel_updated` | channel_id, amount |
| `fee_collected` | amount, recipient |
| `escrow_expired` | settlement_id, sender, amount |
| `channel_expired` | channel_id, sender, recipient, amount |

## EndBlock Processing

The module processes the following in EndBlock:
1. **Expired Escrows**: Automatically refund to sender
2. **Expired Channels**: Emit events for closeable channels

## CLI Commands

### Transactions
```bash
# Instant transfer
statesetd tx settlement instant-transfer [recipient] [amount] [reference] --from [sender]

# Create escrow
statesetd tx settlement create-escrow [recipient] [amount] [expires-in] --from [sender]

# Release escrow
statesetd tx settlement release-escrow [settlement-id] --from [sender]

# Open channel
statesetd tx settlement open-channel [recipient] [deposit] [expires-in-blocks] --from [sender]
```

### Queries
```bash
# Get settlement
statesetd query settlement settlement [id]

# List settlements
statesetd query settlement settlements

# Get channel
statesetd query settlement channel [id]
```

## State

| Key | Value |
|-----|-------|
| `0x01{id}` | Settlement |
| `0x02{id}` | BatchSettlement |
| `0x03` | NextSettlementID |
| `0x04` | NextBatchID |
| `0x05{address}` | MerchantConfig |
| `0x06{id}` | PaymentChannel |
| `0x07` | NextChannelID |
| `0x08` | Params |

## Error Codes

| Code | Error |
|------|-------|
| 1 | Invalid settlement |
| 2 | Settlement not found |
| 3 | Insufficient funds |
| 4 | Unauthorized |
| 5 | Settlement completed |
| 6 | Settlement cancelled |
| 14 | Channel not found |
| 15 | Channel closed |
| 18 | Invalid nonce |
| 21 | Compliance check failed |
| 30 | Invalid signature |
| 32 | Invalid webhook URL |
| 33 | Webhook URL must use HTTPS |
