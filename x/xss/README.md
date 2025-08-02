# STST Staking Token Module

The STST (Stateset Staking Token) module is a comprehensive staking token implementation for the Stateset blockchain network. This module provides secure staking functionality, validator management, and reward distribution mechanisms.

## Overview

STST is the native staking token of the Stateset blockchain, designed to:

- **Secure the Network**: Token holders can stake STST to validators to help secure the blockchain
- **Earn Rewards**: Stakers earn rewards proportional to their stake and the performance of chosen validators
- **Participate in Governance**: STST holders can participate in on-chain governance decisions
- **Validator Operations**: Run validator nodes and earn commission from delegated stake

## Token Details

- **Symbol**: STST
- **Base Denomination**: `ustst` (micro-STST)
- **Display Denomination**: `STST`
- **Decimals**: 6
- **Maximum Supply**: 1,000,000,000 STST (1 billion)
- **Initial Supply**: 100,000,000 STST (100 million)

## Key Features

### Staking Operations
- **Delegate**: Stake STST tokens to validators
- **Undelegate**: Initiate unstaking with a 21-day unbonding period
- **Redelegate**: Move stake between validators instantly
- **Withdraw Rewards**: Claim accumulated staking rewards

### Validator Management
- **Create Validator**: Set up a new validator node
- **Edit Validator**: Update validator metadata and commission
- **Unjail Validator**: Restore a slashed validator to active status

### Security Features
- **Slashing Protection**: Validators are slashed for misbehavior
  - Double signing: 5% slash
  - Downtime: 1% slash
- **Minimum Stake**: 1 STST minimum staking amount
- **Unbonding Period**: 21-day security unbonding period

### Reward Distribution
- **Annual Staking Rate**: 8% annual rewards for stakers
- **Commission System**: Validators can set commission rates
- **Automatic Distribution**: Rewards are distributed each block

## Module Architecture

The STST module is built using the Cosmos SDK and includes:

### Core Components
- **Keeper**: Core business logic for staking operations
- **Types**: Message types, parameters, and data structures
- **Client**: CLI and REST API interfaces
- **Genesis**: Initial state configuration

### Protobuf Definitions
- `xss.proto`: Core data structures and parameters
- `tx.proto`: Transaction message definitions
- `query.proto`: Query service definitions
- `genesis.proto`: Genesis state structure

### Key Files
```
x/xss/
├── keeper/           # Core keeper implementation
├── types/            # Type definitions and interfaces
│   ├── keys.go       # Store keys and constants
│   ├── params.go     # Parameter definitions and validation
│   ├── codec.go      # Message codec registration
│   ├── events.go     # Event type definitions
│   └── expected_keepers.go  # External keeper interfaces
├── client/           # CLI and REST interfaces
└── proto/            # Protobuf definitions
```

## Usage Examples

### Staking Tokens
```bash
# Delegate 100 STST to a validator
statesetd tx xss stake-tokens [validator-address] 100000000ustst --from [delegator]

# Undelegate 50 STST from a validator
statesetd tx xss unstake-tokens [validator-address] 50000000ustst --from [delegator]

# Withdraw staking rewards
statesetd tx xss withdraw-rewards [validator-address] --from [delegator]
```

### Validator Operations
```bash
# Create a new validator
statesetd tx xss create-validator \
  --amount=1000000ustst \
  --pubkey=[consensus-pubkey] \
  --moniker="My Validator" \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --from=[validator-operator]

# Edit validator description
statesetd tx xss edit-validator \
  --moniker="Updated Validator Name" \
  --website="https://validator.example.com" \
  --details="Updated validator description" \
  --from=[validator-operator]
```

### Queries
```bash
# Query module parameters
statesetd query xss params

# Query total supply
statesetd query xss total-supply

# Query all validators
statesetd query xss validators

# Query validator details
statesetd query xss validator [validator-address]

# Query delegator rewards
statesetd query xss rewards [delegator-address]
```

## Parameters

The STST module supports the following configurable parameters:

| Parameter | Default Value | Description |
|-----------|---------------|-------------|
| `mint_denom` | `ustst` | Base denomination for minting |
| `max_supply` | `1000000000000000` | Maximum token supply (1B STST) |
| `initial_supply` | `100000000000000` | Initial token supply (100M STST) |
| `staking_rewards_rate` | `0.08` | Annual staking rewards rate (8%) |
| `min_staking_amount` | `1000000` | Minimum staking amount (1 STST) |
| `unstaking_period` | `21 days` | Unbonding period duration |
| `slash_fraction_double_sign` | `0.05` | Slashing rate for double signing (5%) |
| `slash_fraction_downtime` | `0.01` | Slashing rate for downtime (1%) |
| `governance_voting_period` | `14 days` | Governance proposal voting period |

## Integration

To integrate the STST module into your Cosmos SDK application:

1. **Import the module** in your `app.go`
2. **Register keepers** and wire dependencies
3. **Add to module manager** for genesis and upgrades
4. **Configure parameters** in genesis state

## Security Considerations

- **Validator Security**: Validators must maintain high uptime and avoid double signing
- **Key Management**: Secure private key storage for validator operations
- **Slashing Risk**: Understand slashing conditions before staking
- **Unbonding Period**: Factor in 21-day unbonding when planning liquidity

## Development

The STST module follows Cosmos SDK best practices:

- **Testing**: Comprehensive unit and integration tests
- **Documentation**: Detailed API and usage documentation
- **Upgrades**: Support for on-chain parameter updates
- **Events**: Rich event emission for indexing and monitoring

## Support

For technical support, questions, or contributions:

- **Documentation**: See `/docs` directory
- **Issues**: Submit issues via GitHub
- **Community**: Join the Stateset community channels

---

**Note**: This module is designed for the Stateset blockchain network and implements STST as the primary staking token for network security and governance.