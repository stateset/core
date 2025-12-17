# Orders Module

The Orders module provides comprehensive e-commerce order lifecycle management for the Stateset blockchain, enabling merchants and customers to create, track, and settle orders with built-in dispute resolution.

## Overview

The orders module handles:
- **Order Creation**: Customers create orders with merchants including line items and shipping info
- **Order Lifecycle**: Full state machine from pending through completion
- **Payment Integration**: Direct integration with settlement module for escrow or instant payments
- **Shipping Tracking**: Carrier and tracking number management
- **Dispute Resolution**: Built-in dispute system with evidence submission and authority resolution
- **Auto-Completion**: Automatic order completion after delivery window

## Features

### Order Lifecycle

Orders follow a strict state machine with the following states:

```
pending → confirmed → paid → shipped → delivered → completed
    ↓         ↓        ↓        ↓          ↓
cancelled  cancelled  refunded  disputed   disputed
                      disputed             refunded
```

**State Transitions:**
- `pending`: Order created, awaiting merchant confirmation
- `confirmed`: Merchant accepted, awaiting payment
- `paid`: Payment received (instant or escrowed)
- `shipped`: Merchant shipped, tracking available
- `delivered`: Package received by customer
- `completed`: Order finalized, escrow released
- `cancelled`: Order cancelled before payment
- `refunded`: Payment returned to customer
- `disputed`: Dispute opened, awaiting resolution

### Payment Options

Orders support two payment methods via the settlement module:

1. **Instant Transfer**: Immediate payment to merchant
   - Funds transfer directly on payment
   - Refunds require merchant action

2. **Escrow Payment**: Protected buyer transaction
   - Funds held in escrow until delivery
   - Customer releases on satisfaction
   - Auto-release after delivery window
   - Dispute-protected

### Dispute Resolution

Built-in dispute handling for order issues:
- Customer can open disputes on paid/shipped/delivered orders
- Evidence submission (URLs, descriptions)
- Authority-controlled resolution
- Automatic escrow handling on resolution

### Auto-Completion

Delivered orders auto-complete after configurable window:
- Default: 3 days after delivery
- Releases escrow automatically
- Can be disabled via parameters

## Messages

| Message | Description | Signer |
|---------|-------------|--------|
| `MsgCreateOrder` | Create new order with items | Customer |
| `MsgConfirmOrder` | Merchant confirms order | Merchant |
| `MsgPayOrder` | Pay for confirmed order | Customer |
| `MsgShipOrder` | Mark order as shipped with tracking | Merchant |
| `MsgDeliverOrder` | Mark order as delivered | Customer/Merchant |
| `MsgCompleteOrder` | Complete order, release escrow | Customer |
| `MsgCancelOrder` | Cancel unpaid order | Customer/Merchant |
| `MsgRefundOrder` | Refund paid order | Merchant |
| `MsgOpenDispute` | Open dispute on order | Customer |
| `MsgResolveDispute` | Resolve dispute | Authority |

## Queries

| Query | Description |
|-------|-------------|
| `Order` | Get order by ID |
| `Orders` | List orders with filters (customer, merchant, status) |
| `Params` | Get module parameters |

## Parameters

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `default_order_expiration` | int64 | 86400 | Order expiration in seconds (24h) |
| `default_escrow_expiration` | int64 | 604800 | Escrow duration in seconds (7d) |
| `dispute_window` | int64 | 1209600 | Dispute opening window (14d) |
| `min_order_amount` | Coin | 100ssusd | Minimum order amount |
| `max_order_amount` | Coin | 1000000000000ssusd | Maximum order amount ($1M) |
| `default_fee_rate_bps` | uint32 | 100 | Default fee rate (1%) |
| `stablecoin_denom` | string | ssusd | Payment denomination |
| `auto_complete_after_delivery` | bool | true | Enable auto-completion |
| `auto_complete_window` | int64 | 259200 | Auto-complete delay (3d) |

## Security

### Compliance Integration
- All order parties (customer, merchant) must pass compliance checks
- Orders blocked if either party is sanctioned/blocked
- Integration with compliance module for KYC/AML

### Authorization
- Only customers can create orders, pay, complete, and open disputes
- Only merchants can confirm, ship, and refund orders
- Either party can mark as delivered or cancel (before payment)
- Only authority can resolve disputes

### Input Validation
- All addresses validated via `AccAddressFromBech32`
- Amount validation (positive, within limits)
- Status transition validation (prevents invalid state changes)
- Order item validation (non-empty, valid prices)

### Settlement Integration
- Secure integration with settlement module
- Escrow funds held in module account
- Atomic operations for payment and refund

## Events

| Event | Attributes |
|-------|------------|
| `order_created` | order_id, customer, merchant, total_amount |
| `order_confirmed` | order_id, merchant |
| `order_paid` | order_id, customer, amount, settlement_id, use_escrow |
| `order_shipped` | order_id, merchant, carrier, tracking_number |
| `order_delivered` | order_id, delivered_by |
| `order_completed` | order_id, customer |
| `order_cancelled` | order_id, cancelled_by, reason |
| `order_refunded` | order_id, merchant, refund_amount, reason |
| `order_expired` | order_id, customer |
| `order_auto_completed` | order_id, customer |
| `dispute_opened` | dispute_id, order_id, customer, reason |
| `dispute_resolved` | dispute_id, order_id, resolution, to_customer |

## EndBlock Processing

The module processes the following in EndBlock:
1. **Expired Orders**: Auto-cancel pending/confirmed orders past expiration
2. **Auto-Complete**: Complete delivered orders after auto-complete window

## CLI Commands

### Transactions
```bash
# Create order
statesetd tx orders create-order [merchant] [items-json] [shipping-json] --from [customer]

# Confirm order (merchant)
statesetd tx orders confirm-order [order-id] --from [merchant]

# Pay for order
statesetd tx orders pay-order [order-id] [amount] --use-escrow --from [customer]

# Ship order
statesetd tx orders ship-order [order-id] [carrier] [tracking-number] --from [merchant]

# Mark delivered
statesetd tx orders deliver-order [order-id] --from [customer]

# Complete order
statesetd tx orders complete-order [order-id] --from [customer]

# Cancel order
statesetd tx orders cancel-order [order-id] [reason] --from [signer]

# Refund order
statesetd tx orders refund-order [order-id] [amount] [reason] --from [merchant]

# Open dispute
statesetd tx orders open-dispute [order-id] [reason] [description] --from [customer]
```

### Queries
```bash
# Get order
statesetd query orders order [id]

# List orders
statesetd query orders orders --customer [addr] --merchant [addr] --status [status]

# Get params
statesetd query orders params
```

## State

| Key | Value |
|-----|-------|
| `0x01{id}` | Order |
| `0x02{customer}{id}` | Order index by customer |
| `0x03{merchant}{id}` | Order index by merchant |
| `0x04{status}{id}` | Order index by status |
| `0x05` | NextOrderID |
| `0x06` | Params |
| `0x07{id}` | Dispute |

## Error Codes

| Code | Error | Description |
|------|-------|-------------|
| 2 | ErrOrderNotFound | Order does not exist |
| 3 | ErrInvalidOrder | Invalid order data |
| 4 | ErrUnauthorized | Caller not authorized |
| 5 | ErrInvalidStatus | Invalid order status |
| 6 | ErrInvalidTransition | Invalid status transition |
| 7 | ErrOrderExpired | Order has expired |
| 8 | ErrInsufficientFunds | Insufficient funds for operation |
| 9 | ErrPaymentFailed | Payment processing failed |
| 10 | ErrAlreadyPaid | Order already paid |
| 11 | ErrNotPaid | Order not yet paid |
| 12 | ErrCannotRefund | Order cannot be refunded |
| 13 | ErrCannotCancel | Order cannot be cancelled |
| 14 | ErrDisputeNotFound | Dispute does not exist |
| 15 | ErrDisputeAlreadyExists | Dispute already opened |
| 16 | ErrCannotDispute | Order cannot be disputed |
| 17 | ErrInvalidAmount | Invalid amount |
| 18 | ErrComplianceFailed | Compliance check failed |
| 19 | ErrSettlementFailed | Settlement operation failed |
| 20 | ErrOrderAlreadyCompleted | Order already completed |
| 21 | ErrEmptyItems | Order must have items |
| 22 | ErrInvalidMerchant | Invalid merchant address |
| 23 | ErrInvalidCustomer | Invalid customer address |

## Order Flow Example

```
1. Customer creates order with items
   → Order status: PENDING

2. Merchant confirms order
   → Order status: CONFIRMED

3. Customer pays with escrow
   → Funds held in settlement escrow
   → Order status: PAID

4. Merchant ships order
   → Tracking info recorded
   → Order status: SHIPPED

5. Customer confirms delivery
   → Order status: DELIVERED

6. Customer completes order (or auto-complete after 3 days)
   → Escrow released to merchant
   → Order status: COMPLETED
```

## Dispute Flow Example

```
1. Customer opens dispute on shipped order
   → Dispute created with evidence
   → Order status: DISPUTED

2. Authority reviews and resolves
   → If customer wins: escrow refunded
   → If merchant wins: escrow released
   → Dispute status: RESOLVED
```
