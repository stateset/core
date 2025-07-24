# Stablecoins Module

The Stablecoins module provides comprehensive functionality for creating, managing, and operating stablecoins on the Stateset blockchain. This module is designed to support various types of stablecoins including USD-pegged, EUR-pegged, and other asset-backed stablecoins with advanced features like access controls, reserve management, and price stability mechanisms.

## Features

### Core Functionality
- **Stablecoin Creation**: Create new stablecoins with customizable parameters
- **Minting & Burning**: Controlled token issuance and destruction
- **Access Control**: Whitelist/blacklist functionality for compliance
- **Price Management**: Oracle integration for price data
- **Reserve Management**: Collateral and reserve tracking
- **Pause Controls**: Emergency pause functionality for operations

### Advanced Features
- **Multi-Asset Pegging**: Support for USD, EUR, BTC, and other asset pegs
- **Stability Mechanisms**: Collateralized, algorithmic, and hybrid approaches
- **Fee Management**: Configurable fees for operations
- **Governance Integration**: Parameter updates via governance
- **Analytics**: Comprehensive statistics and reporting

## Architecture

### Key Components

1. **Stablecoin Entity**: Core stablecoin configuration and metadata
2. **PegInfo**: Pegging mechanism and target asset information
3. **ReserveInfo**: Collateral and reserve asset management
4. **AccessControl**: Whitelist/blacklist and KYC requirements
5. **FeeInfo**: Fee structure for operations
6. **PriceData**: Oracle price feed integration

### Module Structure

```
x/stablecoins/
├── types/          # Message types, errors, and keys
├── keeper/         # Business logic and state management  
├── client/cli/     # CLI commands
├── genesis.go      # Genesis state handling
└── module.go       # Module definition and app integration
```

## Usage

### Creating a Stablecoin

```bash
# Create a USD-pegged stablecoin
statesetd tx stablecoins create-stablecoin \
  usdx \
  "StateSet USD" \
  "USDX" \
  6 \
  1000000000000000 \
  --target-asset="USD" \
  --target-price="1.0" \
  --description="USD-pegged stablecoin" \
  --from=issuer
```

### Minting Tokens

```bash
# Mint 1000 USDX tokens
statesetd tx stablecoins mint-stablecoin \
  usdx \
  1000000000 \
  stateset1recipient... \
  --from=issuer
```

### Burning Tokens

```bash
# Burn 500 USDX tokens
statesetd tx stablecoins burn-stablecoin \
  usdx \
  500000000 \
  --from=holder
```

### Access Control

```bash
# Whitelist an address
statesetd tx stablecoins whitelist-address \
  usdx \
  stateset1address... \
  --from=admin

# Blacklist an address
statesetd tx stablecoins blacklist-address \
  usdx \
  stateset1badactor... \
  "Compliance violation" \
  --from=admin
```

### Emergency Controls

```bash
# Pause minting operations
statesetd tx stablecoins pause-stablecoin \
  usdx \
  mint \
  "Emergency maintenance" \
  --from=admin

# Unpause operations
statesetd tx stablecoins unpause-stablecoin \
  usdx \
  mint \
  --from=admin
```

## Querying

### Basic Queries

```bash
# Get stablecoin details
statesetd query stablecoins show-stablecoin usdx

# List all stablecoins
statesetd query stablecoins list-stablecoin

# Get supply information
statesetd query stablecoins supply usdx

# Get ecosystem statistics
statesetd query stablecoins stats
```

### Access Control Queries

```bash
# Check if address is whitelisted
statesetd query stablecoins is-whitelisted usdx stateset1address...

# Check if address is blacklisted
statesetd query stablecoins is-blacklisted usdx stateset1address...
```

### Price and Reserve Data

```bash
# Get price data
statesetd query stablecoins price-data usdx

# Get reserve information
statesetd query stablecoins reserve-info usdx
```

## REST API

The module exposes a comprehensive REST API for integration with web applications and services.

### Stablecoin Operations

```
POST /stateset/stablecoins/v1/stablecoins          # Create stablecoin
GET  /stateset/stablecoins/v1/stablecoin/{denom}    # Get stablecoin
GET  /stateset/stablecoins/v1/stablecoins           # List stablecoins
GET  /stateset/stablecoins/v1/supply/{denom}        # Get supply info
GET  /stateset/stablecoins/v1/price/{denom}         # Get price data
GET  /stateset/stablecoins/v1/stats                 # Get statistics
```

### Access Control

```
GET /stateset/stablecoins/v1/whitelist/{denom}/{address}  # Check whitelist
GET /stateset/stablecoins/v1/blacklist/{denom}/{address}  # Check blacklist
```

## Configuration

### Module Parameters

| Parameter | Description | Default |
|-----------|-------------|---------|
| `max_stablecoins` | Maximum number of stablecoins | 100 |
| `min_initial_supply` | Minimum initial supply | 1,000,000 |
| `max_initial_supply` | Maximum initial supply | 1 quadrillion |
| `creation_fee` | Fee to create stablecoin | 100 STATE |
| `min_reserve_ratio` | Minimum reserve ratio | 100% |
| `max_fee_percentage` | Maximum fee percentage | 5% |

### Stablecoin Configuration

```json
{
  "denom": "usdx",
  "name": "StateSet USD",
  "symbol": "USDX",
  "decimals": 6,
  "max_supply": "1000000000000000",
  "peg_info": {
    "target_asset": "USD",
    "target_price": "1.0",
    "price_tolerance": "0.01"
  },
  "stability_mechanism": "collateralized",
  "access_control": {
    "whitelist_enabled": false,
    "blacklist_enabled": true
  }
}
```

## Integration

### App Integration

Add the module to your application in `app/app.go`:

```go
import (
    stablecoinskeeper "github.com/stateset/core/x/stablecoins/keeper"
    stablecoinstypes "github.com/stateset/core/x/stablecoins/types"
    stablecoinsmodule "github.com/stateset/core/x/stablecoins"
)

// Add to keepers
app.StablecoinsKeeper = stablecoinskeeper.NewKeeper(
    appCodec,
    keys[stablecoinstypes.StoreKey],
    keys[stablecoinstypes.MemStoreKey],
    app.GetSubspace(stablecoinstypes.ModuleName),
    app.AccountKeeper,
    app.BankKeeper,
)

// Add to module manager
app.mm = module.NewManager(
    // other modules...
    stablecoinsmodule.NewAppModule(appCodec, app.StablecoinsKeeper, app.AccountKeeper, app.BankKeeper),
)
```

### Genesis Configuration

```json
{
  "stablecoins": {
    "params": {
      "max_stablecoins": 100,
      "min_initial_supply": "1000000",
      "max_initial_supply": "1000000000000000",
      "creation_fee": "100000000",
      "min_reserve_ratio": "1.00",
      "max_fee_percentage": "0.05"
    },
    "stablecoin_list": [],
    "stablecoin_count": 0
  }
}
```

## Security Features

### Access Controls
- **Issuer Authorization**: Only designated issuers can mint tokens
- **Admin Controls**: Separate admin role for configuration updates
- **Whitelist/Blacklist**: Compliance-friendly address filtering
- **Pause Mechanisms**: Emergency stops for critical operations

### Validation
- **Amount Validation**: Prevents overflow and invalid amounts
- **Supply Limits**: Enforces maximum supply constraints
- **Authorization Checks**: Verifies permissions for all operations
- **Input Sanitization**: Validates all user inputs

### Audit Trail
- **Event Emission**: All operations emit trackable events
- **State Transitions**: Complete audit trail of state changes
- **Error Logging**: Comprehensive error tracking

## Compliance Features

### KYC Integration
- Configurable KYC requirements per stablecoin
- Address verification and validation
- Compliance status tracking

### Regulatory Controls
- **Geographic Restrictions**: Region-based access controls
- **Transaction Limits**: Configurable transaction size limits
- **Reporting**: Built-in compliance reporting tools

## Performance Considerations

### Scalability
- Efficient key-value storage design
- Paginated queries for large datasets
- Optimized iteration patterns

### Gas Optimization
- Minimal storage operations
- Efficient validation logic
- Batched operations where possible

## Error Handling

Common error scenarios and their handling:

| Error | Description | Resolution |
|-------|-------------|------------|
| `ErrStablecoinAlreadyExists` | Denom already in use | Choose different denom |
| `ErrUnauthorized` | Insufficient permissions | Check issuer/admin roles |
| `ErrExceedsMaxSupply` | Minting exceeds limit | Increase max supply or reduce amount |
| `ErrAddressBlacklisted` | Address is blacklisted | Remove from blacklist first |

## Future Enhancements

### Planned Features
- **Cross-chain Integration**: IBC support for inter-chain transfers
- **Advanced Oracles**: Multiple oracle source aggregation
- **Algorithmic Stability**: Advanced algorithmic rebalancing
- **Yield Generation**: Staking rewards for stablecoin holders
- **Insurance Integration**: DeFi insurance protocol integration

### Governance Proposals
- Parameter updates via governance
- Emergency pause proposals
- New stablecoin approvals
- Reserve requirement changes

## Examples

See the `/examples` directory for:
- Basic stablecoin creation
- Multi-asset reserve setup
- Oracle integration patterns
- Compliance implementation guides

## Support

For support and questions:
- GitHub Issues: [stateset/core/issues](https://github.com/stateset/core/issues)
- Documentation: [docs.stateset.com](https://docs.stateset.com)
- Community: [discord.gg/stateset](https://discord.gg/stateset)