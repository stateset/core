# Oracle Module

## Overview

The Oracle module provides decentralized price feed infrastructure for the Stateset blockchain, enabling secure and reliable price data for stablecoins, commerce operations, and DeFi applications. This module implements a multi-provider oracle system with reputation tracking, deviation monitoring, and emergency controls.

## Key Features

### Core Functionality
- **Multi-Provider System**: Support for multiple independent oracle providers
- **Price Aggregation**: Weighted average, median, and confidence calculations
- **Reputation System**: Dynamic reputation tracking for oracle performance
- **Deviation Monitoring**: Automatic detection and penalization of outliers
- **Historical Data**: Complete price history tracking for analytics
- **Emergency Controls**: Circuit breakers for extreme market conditions

### Security Features
- **Provider Authentication**: Only registered providers can submit prices
- **Stake & Slash**: Economic incentives for accurate submissions
- **Deviation Thresholds**: Configurable limits for price variations
- **Minimum Providers**: Required number of providers for valid prices
- **Staleness Detection**: Automatic detection of outdated prices

## Architecture

```
x/oracle/
├── types/           # Core types and message definitions
│   ├── oracle.go    # Price feed and provider structures
│   ├── messages.go  # Transaction message types
│   ├── errors.go    # Error definitions
│   ├── params.go    # Module parameters
│   └── genesis.go   # Genesis state
├── keeper/          # Business logic
│   ├── keeper.go    # Core keeper implementation
│   └── msg_server.go # Message handlers
├── client/cli/      # CLI commands
└── module.go        # Module integration
```

## Usage

### Registering an Oracle Provider

```bash
# Register a new oracle provider (requires authority)
statesetd tx oracle register-oracle \
  "Chainlink" \
  stateset1provider... \
  --priority=10 \
  --min-submissions=100 \
  --max-deviation=0.05 \
  --from=authority
```

### Submitting Price Feeds

```bash
# Submit a price feed as an oracle provider
statesetd tx oracle submit-price \
  "USD" \
  "1.00" \
  --confidence=0.99 \
  --volume=1000000 \
  --expiry="2024-01-01T00:00:00Z" \
  --from=provider
```

### Querying Prices

```bash
# Get current price for an asset
statesetd query oracle price USD

# Get specific price feed
statesetd query oracle price-feed feed-123

# List all oracle providers
statesetd query oracle oracles --active-only

# Get price history
statesetd query oracle price-history USD --limit=100

# Get module parameters
statesetd query oracle params
```

## Price Aggregation

The module uses sophisticated price aggregation algorithms:

1. **Weighted Average**: Considers provider reputation and confidence
2. **Median Calculation**: Provides robust central tendency
3. **Standard Deviation**: Measures price consensus
4. **Confidence Score**: Indicates reliability of aggregated price

### Aggregation Formula

```
WeightedPrice = Σ(Price × Reputation × Confidence) / Σ(Reputation × Confidence)
```

## Reputation System

Oracle providers maintain a reputation score (0.0 to 1.0) that affects:
- Weight in price aggregation
- Reward amounts
- Slash penalties
- Priority in updates

### Reputation Updates
- **Good Submission** (<1% deviation): +0.001 reputation
- **Excessive Deviation**: -0.01 reputation
- **Inactivity**: -0.005 reputation per period

## Parameters

| Parameter | Description | Default |
|-----------|-------------|---------|
| `min_providers` | Minimum oracle providers for valid price | 3 |
| `update_interval` | Expected update frequency | 5 minutes |
| `max_price_age` | Maximum age before price is stale | 10 minutes |
| `price_deviation_threshold` | Maximum allowed deviation | 5% |
| `reward_amount` | Tokens rewarded per valid submission | 100 STATE |
| `slash_amount` | Tokens slashed for invalid submission | 1000 STATE |
| `emergency_threshold` | Deviation triggering emergency mode | 20% |

## Integration with Other Modules

### Stablecoins Module
```go
// Get USD price for stablecoin operations
price, err := oracleKeeper.GetPrice(ctx, "USD")
if err != nil {
    return err
}
```

### Commerce Module
```go
// Get product pricing in multiple currencies
usdPrice, _ := oracleKeeper.GetPrice(ctx, "USD")
eurPrice, _ := oracleKeeper.GetPrice(ctx, "EUR")
```

### DeFi Modules
```go
// Get collateral values for lending
btcPrice, _ := oracleKeeper.GetPrice(ctx, "BTC")
ethPrice, _ := oracleKeeper.GetPrice(ctx, "ETH")
```

## Events

The module emits the following events:

| Event | Description | Attributes |
|-------|-------------|------------|
| `price_feed_submitted` | New price feed received | feed_id, asset, price, provider |
| `price_updated` | Aggregated price updated | asset, price, median, confidence |
| `oracle_registered` | New oracle provider | name, address, priority |
| `provider_inactive` | Provider hasn't submitted | provider, last_update |
| `stale_price_detected` | Price exceeds max age | asset, last_update, age |
| `emergency_threshold_exceeded` | High volatility detected | asset, standard_dev |

## Emergency Procedures

### Circuit Breaker Activation
When price deviation exceeds emergency threshold:
1. Event emitted for monitoring
2. Governance proposal can pause module
3. Manual intervention may be required

### Recovery Process
1. Identify cause of deviation
2. Remove/penalize bad actors
3. Re-establish price consensus
4. Resume normal operations

## Testing

### Unit Tests
```bash
go test ./x/oracle/keeper/...
go test ./x/oracle/types/...
```

### Integration Tests
```bash
# Start local testnet
./scripts/run-stateset.sh

# Register test oracles
./scripts/test-oracle-registration.sh

# Submit test prices
./scripts/test-price-submission.sh
```

## Security Considerations

1. **Sybil Resistance**: Provider registration requires governance approval
2. **Economic Security**: Stake requirements and slashing mechanisms
3. **Data Integrity**: Cryptographic signatures on all submissions
4. **Availability**: Multiple providers ensure redundancy
5. **Manipulation Resistance**: Median and deviation checks

## Future Enhancements

- **Cross-Chain Oracles**: IBC-based price feed aggregation
- **Chainlink Integration**: Direct Chainlink price feed support
- **Band Protocol Support**: Alternative oracle provider
- **AI-Based Filtering**: Machine learning for anomaly detection
- **Prediction Markets**: Crowd-sourced price discovery

## Resources

- [Oracle Design Doc](docs/oracle-design.md)
- [Integration Guide](docs/oracle-integration.md)
- [Security Audit Report](docs/oracle-audit.pdf)
- [Performance Benchmarks](docs/oracle-performance.md)

## Support

For issues or questions:
- GitHub: [stateset/core/issues](https://github.com/stateset/core/issues)
- Discord: [discord.gg/stateset](https://discord.gg/stateset)
- Documentation: [docs.stateset.com/oracle](https://docs.stateset.com/oracle)