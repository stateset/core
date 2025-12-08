# Treasury Module

The Treasury module manages protocol treasury and fund allocation for the Stateset blockchain.

## Overview

The treasury module handles:
- **Treasury Management**: Manage protocol treasury funds
- **Fund Allocation**: Allocate funds to various purposes
- **Revenue Distribution**: Distribute protocol revenues

## Features

### Treasury Operations
- Receive protocol fees and revenues
- Manage treasury balance
- Execute approved disbursements

### Governance Integration
- Governance-controlled spending
- Proposal-based fund allocation

## Messages

| Message | Description |
|---------|-------------|
| `MsgWithdraw` | Withdraw funds from treasury |
| `MsgAllocate` | Allocate funds to address |
| `MsgUpdateParams` | Update treasury parameters |

## State

| Key | Value |
|-----|-------|
| `0x01` | TreasuryBalance |
| `0x02` | AllocationRecords |
| `0x03` | Params |

## Events

| Event | Attributes |
|-------|------------|
| `funds_received` | source, amount |
| `funds_withdrawn` | recipient, amount, purpose |
| `allocation_executed` | recipient, amount |
