# Stablecoin Payments for Orders - Complete Guide

This guide explains how to use stablecoin payments for orders in the Stateset blockchain platform.

## Overview

The stablecoin payment system allows customers to pay for orders using supported stablecoins with optional escrow functionality. The system provides:

- Direct stablecoin payments to merchants
- Escrow-based payments for secure transactions
- Multi-confirmation payment verification
- Automated refund processing
- Comprehensive payment tracking

## Features

### 🪙 Supported Payment Methods
- Direct stablecoin transfers
- Escrow-secured payments
- Multi-signature release mechanisms
- Automated timeout handling

### 🔒 Security Features
- Address whitelist/blacklist validation
- Minimum/maximum amount controls
- Confirmation requirements
- Escrow timeout protection

### 💰 Refund Capabilities
- Full and partial refunds
- Automatic escrow refunds
- Direct merchant refunds
- Comprehensive audit trails

## Prerequisites

Before using stablecoin payments, ensure:

1. **Stablecoins are configured**: The stablecoin denomination must be registered in the stablecoins module
2. **Accounts have balances**: Customer accounts must have sufficient stablecoin balances
3. **Addresses are validated**: If whitelisting is enabled, addresses must be whitelisted

## Basic Usage

### 1. Direct Payment

Pay for an order directly without escrow:

```bash
# Direct payment example
statesed tx orders stablecoin pay-with-stablecoin \
  ORDER-123 \
  uusdc \
  1000000 \
  cosmos1customer... \
  cosmos1merchant... \
  1.0 \
  --from customer \
  --chain-id stateset-1
```

**Parameters:**
- `ORDER-123`: Order ID
- `uusdc`: Stablecoin denomination
- `1000000`: Payment amount (1 USDC with 6 decimals)
- `cosmos1customer...`: Customer's address
- `cosmos1merchant...`: Merchant's address
- `1.0`: Exchange rate (USD to stablecoin)

### 2. Escrow Payment

Use escrow for secure payments:

```bash
# Escrow payment example
statesed tx orders stablecoin pay-with-stablecoin \
  ORDER-123 \
  uusdc \
  1000000 \
  cosmos1customer... \
  cosmos1merchant... \
  1.0 \
  --use-escrow \
  --confirmations-required 6 \
  --escrow-timeout-hours 168 \
  --from customer \
  --chain-id stateset-1
```

**Additional Flags:**
- `--use-escrow`: Enable escrow protection
- `--confirmations-required 6`: Require 6 confirmations
- `--escrow-timeout-hours 168`: 7-day timeout

### 3. Confirm Payment

Merchant confirms payment after verifications:

```bash
# Confirm payment
statesed tx orders stablecoin confirm-stablecoin-payment \
  ORDER-123 \
  6 \
  1234567 \
  --from merchant \
  --chain-id stateset-1
```

**Parameters:**
- `ORDER-123`: Order ID
- `6`: Confirmation count achieved
- `1234567`: Block height of confirmation

### 4. Release Escrow

Release escrowed funds to merchant:

```bash
# Release escrow (can be called by customer or merchant)
statesed tx orders stablecoin release-escrow \
  ORDER-123 \
  --from customer \
  --chain-id stateset-1
```

### 5. Process Refund

Merchant processes a refund:

```bash
# Full refund
statesed tx orders stablecoin refund-stablecoin-payment \
  ORDER-123 \
  cosmos1customer... \
  1000000 \
  "Order cancelled by customer" \
  --from merchant \
  --chain-id stateset-1

# Partial refund
statesed tx orders stablecoin refund-stablecoin-payment \
  ORDER-123 \
  cosmos1customer... \
  500000 \
  "Partial return - damaged item" \
  --partial \
  --from merchant \
  --chain-id stateset-1
```

## Advanced Features

### Exchange Rate Handling

The system supports dynamic exchange rates for accurate pricing:

```go
// Example: 1 USD = 0.998 USDC
exchangeRate := sdk.MustNewDecFromStr("0.998")
```

### Timeout Management

Escrow timeouts provide automatic protection:

- **Default**: 30 days
- **Configurable**: 1 hour to 1 year
- **Automatic refund**: On timeout expiry

### Multi-Confirmation Security

Configure confirmation requirements based on payment amount:

- **Small payments**: 1-3 confirmations
- **Medium payments**: 3-6 confirmations
- **Large payments**: 6+ confirmations

## Integration Examples

### E-commerce Integration

```javascript
// Frontend payment processing
async function processStablecoinPayment(orderDetails) {
  const txMsg = {
    type: "orders/MsgPayWithStablecoin",
    value: {
      creator: customerAddress,
      order_id: orderDetails.id,
      stablecoin_denom: "uusdc",
      stablecoin_amount: orderDetails.total.toString(),
      customer_address: customerAddress,
      merchant_address: merchantAddress,
      exchange_rate: await getExchangeRate("USD", "USDC"),
      use_escrow: orderDetails.requiresEscrow,
      confirmations_required: getRequiredConfirmations(orderDetails.total),
      escrow_timeout: getEscrowTimeout(orderDetails.category)
    }
  };
  
  return await broadcastTx(txMsg);
}
```

### Backend Order Processing

```go
// Golang service integration
func (s *OrderService) ProcessStablecoinPayment(
    ctx context.Context, 
    request *PaymentRequest,
) (*PaymentResponse, error) {
    // Validate payment request
    if err := s.validatePaymentRequest(request); err != nil {
        return nil, err
    }
    
    // Create payment message
    msg := &types.MsgPayWithStablecoin{
        Creator:               request.CustomerAddress,
        OrderId:               request.OrderID,
        StablecoinDenom:       request.StablecoinDenom,
        StablecoinAmount:      request.Amount,
        CustomerAddress:       request.CustomerAddress,
        MerchantAddress:       s.merchantAddress,
        ExchangeRate:          request.ExchangeRate,
        UseEscrow:             request.UseEscrow,
        ConfirmationsRequired: request.ConfirmationsRequired,
        EscrowTimeout:         request.EscrowTimeout,
    }
    
    // Submit transaction
    txResponse, err := s.txClient.BroadcastTx(ctx, msg)
    if err != nil {
        return nil, fmt.Errorf("payment failed: %w", err)
    }
    
    return &PaymentResponse{
        TxHash:    txResponse.TxHash,
        Success:   true,
        Message:   "Payment processed successfully",
        Timestamp: time.Now(),
    }, nil
}
```

## Error Handling

### Common Errors and Solutions

1. **Invalid Stablecoin**: `ErrInvalidStablecoin`
   - **Cause**: Unsupported stablecoin denomination
   - **Solution**: Use registered stablecoins only

2. **Insufficient Balance**: `ErrInsufficientBalance`
   - **Cause**: Customer lacks sufficient funds
   - **Solution**: Ensure adequate balance before payment

3. **Order Already Paid**: `ErrOrderAlreadyPaid`
   - **Cause**: Duplicate payment attempt
   - **Solution**: Check order status before payment

4. **Escrow Timeout**: `ErrEscrowTimeout`
   - **Cause**: Escrow period expired
   - **Solution**: Process refund or extend timeout

5. **Payment Failed**: `ErrPaymentFailed`
   - **Cause**: Various blockchain-level issues
   - **Solution**: Retry with proper gas/fees

### Error Response Format

```json
{
  "error": {
    "code": 1202,
    "message": "payment transaction failed",
    "details": "insufficient gas provided"
  }
}
```

## Best Practices

### For Merchants

1. **Set appropriate confirmations**: Higher amounts need more confirmations
2. **Use escrow for large orders**: Provides customer confidence
3. **Monitor payment status**: Implement proper event listening
4. **Handle refunds promptly**: Maintain customer satisfaction

### For Customers

1. **Verify order details**: Ensure amounts and addresses are correct
2. **Use escrow when uncertain**: Provides payment protection
3. **Monitor confirmation status**: Track payment progress
4. **Keep transaction records**: Save payment proofs

### For Developers

1. **Implement proper error handling**: Handle all error scenarios
2. **Use event-driven updates**: Listen for payment events
3. **Validate inputs thoroughly**: Prevent invalid transactions
4. **Test thoroughly**: Use testnet for development

## Events and Monitoring

### Payment Events

The system emits various events for monitoring:

```go
// Payment processed event
"stablecoin_payment_processed" {
    "order_id": "ORDER-123",
    "customer": "cosmos1abc...",
    "merchant": "cosmos1def...",
    "stablecoin_denom": "uusdc",
    "amount": "1000000",
    "use_escrow": "true",
    "tx_hash": "ABC123..."
}

// Payment confirmed event
"stablecoin_payment_confirmed" {
    "order_id": "ORDER-123",
    "confirmation_count": "6",
    "block_height": "1234567"
}

// Escrow released event
"escrow_released" {
    "order_id": "ORDER-123",
    "merchant": "cosmos1def...",
    "amount": "1000000"
}
```

### Monitoring Dashboard

Implement monitoring for:
- Payment success rates
- Average confirmation times
- Escrow release patterns
- Refund frequencies
- Error distributions

## Security Considerations

### Address Validation
- Verify customer and merchant addresses
- Implement address blacklist checking
- Use whitelist for restricted stablecoins

### Amount Validation
- Check minimum/maximum limits
- Validate against order totals
- Implement rate limiting

### Escrow Security
- Set reasonable timeout periods
- Monitor escrow balances
- Implement automated release triggers

## Testing

### Unit Tests

```go
func TestPayWithStablecoin(t *testing.T) {
    // Test direct payment
    msg := &types.MsgPayWithStablecoin{
        Creator:               customerAddr,
        OrderId:               "TEST-001",
        StablecoinDenom:       "uusdc",
        StablecoinAmount:      sdk.NewInt(1000000),
        CustomerAddress:       customerAddr,
        MerchantAddress:       merchantAddr,
        ExchangeRate:          sdk.OneDec(),
        UseEscrow:             false,
        ConfirmationsRequired: 1,
    }
    
    response, err := msgServer.PayWithStablecoin(ctx, msg)
    require.NoError(t, err)
    require.True(t, response.Success)
}
```

### Integration Tests

Test complete payment flows:
1. Order creation
2. Payment processing
3. Confirmation handling
4. Escrow release
5. Refund processing

## Troubleshooting

### Common Issues

1. **Transaction not found**
   - Check transaction hash
   - Verify network connectivity
   - Confirm correct chain ID

2. **Payment stuck in pending**
   - Check confirmation requirements
   - Verify network congestion
   - Monitor block production

3. **Escrow not releasing**
   - Check timeout status
   - Verify authorization
   - Confirm order completion

### Support Resources

- **Documentation**: Full API reference
- **Examples**: Sample implementations
- **Community**: Developer support channels
- **Testing**: Testnet environments

## Conclusion

The stablecoin payment system provides a robust, secure, and flexible solution for processing cryptocurrency payments in e-commerce applications. By following this guide, developers can implement comprehensive payment functionality with proper error handling, security measures, and monitoring capabilities.

For additional support or questions, refer to the full API documentation or reach out to the development community.