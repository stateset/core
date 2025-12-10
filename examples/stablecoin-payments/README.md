# ssUSD Stablecoin Payments Examples

Complete examples for integrating ssUSD stablecoin payments into your applications on the Stateset Commerce Network.

## Quick Start

```bash
# Run the quick start demo
./shell/quick-start.sh
```

This will:
1. Create test accounts (Alice and Bob)
2. Fund Alice with ssUSD
3. Send your first ssUSD payment to Bob
4. Show you what to do next

## Example Categories

### Shell Scripts (`shell/`)

Interactive bash scripts demonstrating core payment flows:

| Script | Description |
|--------|-------------|
| `quick-start.sh` | 5-minute intro to ssUSD payments |

Additional scripts in `/scripts/`:
- `ssusd-payment-demo.sh` - Full payment demo with multiple scenarios
- `demo-order-processing.sh` - E-commerce order flow with ssUSD
- `test-ssusd-transfers.sh` - Comprehensive transfer testing

### TypeScript SDK (`typescript/`)

Production-ready examples using the Stateset TypeScript SDK:

| Example | Description | Use Case |
|---------|-------------|----------|
| `basic-payment.ts` | Simple instant transfers | Peer-to-peer payments |
| `escrow-payment.ts` | Escrow-based transactions | Marketplace purchases |
| `batch-payments.ts` | Multi-recipient payments | Payroll, revenue splits |
| `payment-channel.ts` | Off-chain micropayments | Streaming, gaming |
| `merchant-integration.ts` | Full e-commerce flow | Online stores |

#### Running TypeScript Examples

```bash
# Install dependencies
cd sdk/typescript
npm install

# Set environment variables
export STATESET_RPC="http://localhost:26657"
export STATESET_MNEMONIC="your wallet mnemonic here"

# Run an example
npx ts-node ../../examples/stablecoin-payments/typescript/basic-payment.ts
```

### Python SDK (`python/`)

Python implementations for backend services and automation:

| Example | Description | Use Case |
|---------|-------------|----------|
| `basic_payment.py` | Simple instant transfers | Scripting, automation |
| `escrow_payment.py` | Escrow-based transactions | Backend services |
| `batch_payments.py` | Multi-recipient payments | Batch processing |

#### Running Python Examples

```bash
# Set environment variables
export STATESET_CHAIN_ID="stateset-1"
export STATESET_NODE="tcp://localhost:26657"

# Run an example
python examples/stablecoin-payments/python/basic_payment.py
```

## Payment Patterns

### 1. Instant Transfer

Direct peer-to-peer payment, settled immediately on-chain.

```typescript
// TypeScript
await client.settlement.instantTransfer({
  recipient: 'stateset1...',
  amount: { denom: 'ssusd', amount: '100000000' }, // 100 ssUSD
  memo: 'Payment for services',
});
```

```bash
# CLI
statesetd tx bank send alice bob 100000000ssusd --keyring-backend=test --yes
```

**Best for:** Tips, small purchases, known counterparties

### 2. Escrow Payment

Funds held by the protocol until conditions are met.

```typescript
// TypeScript
const result = await client.settlement.createEscrow({
  recipient: sellerAddress,
  amount: { denom: 'ssusd', amount: '500000000' },
  releaseTime: new Date(Date.now() + 7 * 24 * 60 * 60 * 1000),
  conditions: JSON.stringify({ orderId: 'ORD-123' }),
});

// After delivery confirmation
await client.settlement.releaseEscrow(result.escrowId);
```

**Best for:** Marketplace transactions, freelance work, any buyer-seller scenario

### 3. Batch Payment

Multiple payments in a single transaction for efficiency.

```typescript
// TypeScript
await client.settlement.batchPayment({
  payments: [
    { recipient: 'stateset1alice...', amount: { denom: 'ssusd', amount: '5000000000' } },
    { recipient: 'stateset1bob...', amount: { denom: 'ssusd', amount: '4500000000' } },
    { recipient: 'stateset1carol...', amount: { denom: 'ssusd', amount: '4000000000' } },
  ],
  memo: 'Payroll - December 2024',
});
```

**Best for:** Payroll, affiliate payouts, revenue distribution, mass refunds

### 4. Payment Channel

Off-chain payments with on-chain settlement. Zero gas for individual payments.

```typescript
// Open channel with deposit
const channel = await client.settlement.openChannel(
  recipientAddress,
  { denom: 'ssusd', amount: '10000000' },
  expiresAt,
);

// Send many off-chain payments (no gas!)
// ...

// Close channel and settle final balance
await client.settlement.closeChannel(channelId, finalAmount, signature);
```

**Best for:** Streaming payments, micropayments, gaming, IoT

## ssUSD Denomination

ssUSD uses 6 decimal places:

| Human Amount | Base Units |
|--------------|------------|
| 1 ssUSD | 1,000,000 |
| 100 ssUSD | 100,000,000 |
| 0.01 ssUSD | 10,000 |
| 0.000001 ssUSD | 1 |

Conversion helpers:

```typescript
// TypeScript
const toBaseUnits = (amount: number): string =>
  Math.floor(amount * 1_000_000).toString();

const fromBaseUnits = (amount: string): number =>
  parseInt(amount) / 1_000_000;
```

```python
# Python
def to_base_units(amount: float) -> int:
    return int(amount * 1_000_000)

def from_base_units(amount: int) -> float:
    return amount / 1_000_000
```

## Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `STATESET_RPC` | RPC endpoint | `http://localhost:26657` |
| `STATESET_CHAIN_ID` | Chain identifier | `stateset-1` |
| `STATESET_MNEMONIC` | Wallet mnemonic | Required for signing |
| `STATESET_KEYRING` | Keyring backend | `test` |
| `STATESET_BINARY` | CLI binary path | `statesetd` |

### Network Endpoints

| Network | RPC Endpoint |
|---------|--------------|
| Local | `http://localhost:26657` |
| Testnet | `https://rpc.testnet.stateset.network` |
| Mainnet | `https://rpc.stateset.network` |

## Error Handling

Common errors and solutions:

```typescript
try {
  await client.settlement.instantTransfer({ ... });
} catch (error) {
  if (error.message.includes('insufficient funds')) {
    // Handle: Check balance before sending
  } else if (error.message.includes('account not found')) {
    // Handle: Recipient doesn't exist on chain
  } else if (error.message.includes('out of gas')) {
    // Handle: Increase gas limit
  }
}
```

## Security Best Practices

1. **Never hardcode mnemonics** - Use environment variables or secure vaults
2. **Validate addresses** - Check recipient addresses before sending
3. **Use escrow for untrusted parties** - Never send direct payments to strangers
4. **Set appropriate timeouts** - Don't lock funds in escrow indefinitely
5. **Test on testnet first** - Always test payment flows before mainnet

## Related Documentation

- [Settlement Architecture](/docs/settlement-architecture.md)
- [CLI Guide](/docs/cli-guide.md)
- [TypeScript SDK](/sdk/typescript/)
- [Security Architecture](/docs/security_architecture.md)

## Support

- Issues: https://github.com/stateset/core/issues
- Documentation: https://docs.stateset.network
- Discord: https://discord.gg/stateset
