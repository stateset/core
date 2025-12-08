# Payments Module

The Payments module provides payment intent lifecycle management for the Stateset blockchain.

## Overview

The payments module handles:
- **Payment Intent Creation**: Create pending payment records
- **Payment Settlement**: Complete payments to payees
- **Payment Cancellation**: Cancel pending payments
- **Payment Tracking**: Query payment status and history

## Features

### Payment Intent Lifecycle
1. **Create**: Payer creates payment intent, funds held in module
2. **Settle**: Payee claims funds, payment completed
3. **Cancel**: Payer cancels, funds returned

### Compliance Integration
- Mandatory compliance checks for payer and payee
- Blocked if parties are non-compliant

### Self-Payment Prevention
- Payer and payee must be different addresses

## Payment States

| State | Description |
|-------|-------------|
| `PENDING` | Payment created, awaiting settlement |
| `SETTLED` | Payment completed to payee |
| `CANCELLED` | Payment cancelled, funds returned |

## Messages

| Message | Description |
|---------|-------------|
| `MsgCreatePayment` | Create new payment intent |
| `MsgSettlePayment` | Settle payment to payee |
| `MsgCancelPayment` | Cancel pending payment |

## Queries

| Query | Description |
|-------|-------------|
| `Payment` | Get payment by ID |
| `Payments` | List all payments with pagination |
| `PaymentsByPayer` | Get payments for specific payer |
| `PaymentsByPayee` | Get payments for specific payee |
| `PaymentsByStatus` | Filter payments by status |

## State

| Key | Value |
|-----|-------|
| `0x01{id}` | PaymentIntent |
| `0x02` | NextPaymentID |

## Events

| Event | Attributes |
|-------|------------|
| `payment_created` | id, payer, payee, amount |
| `payment_settled` | id, payee, amount |
| `payment_cancelled` | id, payer, amount |
