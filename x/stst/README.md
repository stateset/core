# STST Token Module

The STST (StateSet Token) module implements the native staking token for the StateSet Blockchain. This module provides comprehensive tokenomics functionality including staking, governance, vesting, and deflationary mechanisms.

## Overview

STST is designed as the economic engine of the StateSet Protocol with the following key characteristics:

- **Fixed Supply**: 1 billion STST tokens (no inflation)
- **Deflationary Mechanism**: Transaction fee burning
- **Staking Rewards**: Distributed to validators and delegators
- **Governance**: DAO voting based on staked tokens
- **Vesting Schedules**: Controlled token release for different stakeholder groups

## Token Distribution

| Category | Allocation | Amount (STST) | Vesting Schedule |
|----------|------------|---------------|------------------|
| Protocol Treasury | 30% | 300M | DAO-governed |
| Validator & Agent Rewards | 25% | 250M | 10 years linear |
| Team & Founders | 15% | 150M | 12m cliff + 36m linear |
| Investors | 15% | 150M | 6m cliff + 24m linear |
| Partner Ecosystem | 10% | 100M | 3m cliff + 24m linear |
| Community & Airdrop | 5% | 50M | Immediate/12m linear |

## Key Features

### 1. Staking
- Stake STST tokens to validators
- Earn staking rewards
- 21-day unbonding period
- Minimum staking amount: 1,000 STST

### 2. Governance
- Submit proposals for protocol changes
- Vote on proposals with staked tokens
- Proposal types: parameter changes, software upgrades, treasury spending
- Voting period: 7 days (configurable)

### 3. Fee Burning
- 25% of transaction fees are burned (configurable)
- Creates deflationary pressure
- Historical burn rate tracking

### 4. Vesting
- Linear vesting schedules for different categories
- Cliff periods before vesting begins
- Claimable tokens after vesting

### 5. Slashing
- 5% slashing rate for malicious behavior
- 50% of slashed tokens are burned

## Messages

### Staking
- `MsgStakeTokens`: Stake tokens to a validator
- `MsgUnstakeTokens`: Begin unstaking process
- `MsgClaimStakingRewards`: Claim accumulated rewards

### Governance
- `MsgSubmitProposal`: Submit a governance proposal
- `MsgVote`: Vote on a proposal

### Vesting
- `MsgClaimVestedTokens`: Claim vested tokens

### Admin
- `MsgUpdateParams`: Update module parameters (governance only)

## Queries

- `Params`: Get module parameters
- `StakingState`: Get current staking statistics
- `FeeBurnState`: Get fee burning statistics
- `StakerInfo`: Get staker information
- `Proposals`: List governance proposals
- `VestingSchedules`: List vesting schedules
- `TokenSupply`: Get token supply information

## Parameters

- `total_supply`: Total token supply (1B STST)
- `token_denom`: Token denomination ("stst")
- `staking_rewards_rate`: Annual staking rewards rate (10%)
- `fee_burn_rate`: Fee burning percentage (25%)
- `min_staking_amount`: Minimum staking amount (1,000 STST)
- `governance_voting_period`: Voting period duration (7 days)
- `slashing_rate`: Slashing percentage (5%)
- `burn_from_slash_rate`: Burn percentage of slashed tokens (50%)

## Economic Flywheel

1. **Demand**: Brands fund autonomous workflows, creating STST demand
2. **Utility**: AI Agents consume STST as execution fuel
3. **Burning**: Portion of fees burned, reducing supply
4. **Staking**: Validators stake STST for network security
5. **Rewards**: Stakers earn rewards from network fees
6. **Value Accrual**: Increasing scarcity + demand = positive price pressure

## Implementation Notes

- Built on Cosmos SDK
- Uses protobuf for message definitions
- Implements standard Cosmos module interfaces
- Includes comprehensive invariants for safety
- Supports governance parameter updates

## Security Considerations

- Fixed supply prevents inflation attacks
- Slashing mechanism deters malicious behavior
- Vesting schedules prevent token dumps
- Module permissions restrict sensitive operations
- Invariants ensure system consistency

## Development Status

The module is currently in development with the following components completed:

✅ Core module structure
✅ Protobuf definitions
✅ Parameter management
✅ Genesis state handling
✅ Basic staking functionality
✅ Vesting schedule framework
✅ Fee burning mechanism
✅ Governance framework
✅ Error handling
✅ Invariants

🚧 Pending implementation:
- Complete governance logic
- Staking rewards distribution
- CLI commands
- Full query implementations
- Integration tests
- Protobuf code generation

## Usage Example

```go
// Stake tokens
msg := &types.MsgStakeTokens{
    Staker:           "cosmos1...",
    ValidatorAddress: "cosmosvaloper1...",
    Amount:           sdk.NewCoin("stst", sdk.NewInt(1000000000)), // 1000 STST
}

// Submit governance proposal
proposal := &types.MsgSubmitProposal{
    Proposer:     "cosmos1...",
    Title:        "Increase Fee Burn Rate",
    Description:  "Proposal to increase fee burn rate to 30%",
    ProposalType: types.PROPOSAL_TYPE_PARAMETER_CHANGE,
    InitialDeposit: sdk.NewCoins(sdk.NewCoin("stst", sdk.NewInt(1000000000))),
}
```