# STST: StateSet Autonomous Commerce Network Token

The STST (StateSet Token) module powers the StateSet Protocol - an autonomous commerce network where AI agents execute workflows, manage transactions, and secure the ecosystem through a sophisticated tokenomics model.

## 🎯 Mission

STST is the economic engine of the StateSet Protocol, creating a self-sustaining ecosystem that incentivizes participation, secures the network, and drives long-term value accrual through direct utility and deflationary mechanisms.

## 🚀 The STST Economic Flywheel

The tokenomics create a virtuous cycle where network growth directly enhances token value:

1. **Demand Driver**: Brands and services fund autonomous workflows, creating foundational demand for STST
2. **Utility & Consumption**: AI Agents consume STST as execution fuel, with a portion burned to reduce supply
3. **Security & Yield**: Validators and stakers lock STST to secure the network and earn rewards
4. **Value Accrual**: Growing usage increases burn rate and staking locks, creating scarcity and value appreciation

## 💎 Core Token Details

| Attribute | Value |
|-----------|-------|
| **Token Name** | StateSet Token |
| **Ticker** | $STST |
| **Total Supply** | 1,000,000,000 STST (Fixed - No Inflation) |
| **Initial Circulating** | 50,000,000 STST (5% at TGE) |
| **Base Denomination** | `ustst` (micro-STST) |
| **Decimals** | 6 |
| **Core Functions** | Utility, Governance, Staking, Settlement |
| **Value Model** | Deflationary (Fee Burning) + Staking Rewards |

## 📊 Token Allocation & Vesting

| Category | Allocation | Amount (STST) | Vesting Schedule |
|----------|------------|---------------|------------------|
| **Protocol Treasury** | 30% | 300,000,000 | DAO-governed for ecosystem grants & liquidity |
| **Validator & Agent Rewards** | 25% | 250,000,000 | Linear release over 10 years via smart contract |
| **Team & Founders** | 15% | 150,000,000 | 12-month cliff + 36-month linear vesting |
| **Investors** | 15% | 150,000,000 | 6-12 month cliff + 24-36 month linear vesting |
| **Partner Ecosystem** | 10% | 100,000,000 | DAO-governed for integrations & partnerships |
| **Community & Airdrop** | 5% | 50,000,000 | Immediate for early adopters & community building |

## 🔥 Deflationary Mechanisms

### Fixed Supply Design
- **Total Supply**: 1 billion STST (no new tokens can ever be minted)
- **No Inflation**: Only pre-allocated validator rewards are distributed over 10 years
- **Deflationary Pressure**: Multiple burn mechanisms reduce circulating supply

### Burn Mechanisms
1. **Execution Fee Burn**: 50% of AI agent execution fees permanently burned
2. **Slashing Burns**: Portion of slashed validator tokens removed from circulation  
3. **Protocol Buybacks**: DAO can vote to burn tokens from protocol revenue

## ⚡ Utility & Staking Features

### AI Agent Operations
- **Execution Fuel**: STST powers all AI agent workflows and transactions
- **Gas Payments**: Agents pay execution fees in STST (0.001 STST base fee)
- **Network Access**: STST required for agent registration and operation

### Staking & Security
- **Validator Staking**: Secure the network and earn from 250M STST rewards pool
- **Delegated Staking**: Token holders delegate to validators for 12% annual rewards
- **Slashing Protection**: 5% slash for double signing, 1% for downtime
- **Unbonding Period**: 21-day security period for unstaking

### Governance Power
- **DAO Voting**: STST holders control protocol parameters and upgrades
- **Treasury Management**: Community controls 300M STST treasury allocation
- **Parameter Updates**: Adjust burn rates, fees, and reward mechanisms
- **Voting Period**: 7-day voting cycles for efficient governance

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

## 🔧 Usage Examples

### Agent Operations
```bash
# Execute an AI agent workflow
statesetd tx xss execute-agent \
  --agent-id="autonomous_commerce_agent_001" \
  --workflow-data=[encoded-workflow] \
  --execution-fee=1000ustst \
  --from=[executor-address]

# Query agent execution history
statesetd query xss agent-executions [agent-id]
```

### Staking & Rewards
```bash
# Delegate 100 STST to a validator
statesetd tx xss stake-tokens [validator-address] 100000000ustst --from [delegator]

# Undelegate 50 STST from a validator
statesetd tx xss unstake-tokens [validator-address] 50000000ustst --from [delegator]

# Withdraw staking rewards
statesetd tx xss withdraw-rewards [validator-address] --from [delegator]
```

### Deflationary Operations
```bash
# Burn tokens from treasury (DAO governance)
statesetd tx xss burn-tokens \
  --amount=1000000ustst \
  --reason="Protocol buyback and burn" \
  --from=[treasury-address]

# Query burned tokens history
statesetd query xss burned-tokens
```

### Validator Operations
```bash
# Create a new validator
statesetd tx xss create-validator \
  --amount=1000000ustst \
  --pubkey=[consensus-pubkey] \
  --moniker="StateSet Validator" \
  --commission-rate="0.05" \
  --commission-max-rate="0.10" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --from=[validator-operator]
```

### Governance & Queries
```bash
# Query module parameters
statesetd query xss params

# Query tokenomics data
statesetd query xss total-supply
statesetd query xss circulating-supply
statesetd query xss staked-supply

# Query validator information
statesetd query xss validators
statesetd query xss validator [validator-address]

# Query rewards and burns
statesetd query xss rewards [delegator-address]
statesetd query xss burn-rate
```

## 🛠️ Technical Parameters

The STST module supports the following configurable parameters:

| Parameter | Default Value | Description |
|-----------|---------------|-------------|
| `mint_denom` | `ustst` | Base denomination for STST tokens |
| `max_supply` | `1000000000000000` | Fixed maximum supply (1B STST) |
| `initial_supply` | `50000000000000` | Initial circulating supply (50M STST) |
| `staking_rewards_rate` | `0.12` | Annual staking rewards rate (12%) |
| `min_staking_amount` | `1000000` | Minimum staking amount (1 STST) |
| `unstaking_period` | `21 days` | Unbonding period duration |
| `slash_fraction_double_sign` | `0.05` | Slashing rate for double signing (5%) |
| `slash_fraction_downtime` | `0.01` | Slashing rate for downtime (1%) |
| `governance_voting_period` | `7 days` | DAO voting period |
| `burn_rate` | `0.50` | Percentage of fees burned (50%) |
| `agent_execution_fee` | `1000` | Base fee per agent execution (0.001 STST) |
| `validator_rewards_pool` | `250000000000000` | Validator rewards allocation (250M STST) |
| `treasury_address` | `` | DAO treasury address (governance controlled) |

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