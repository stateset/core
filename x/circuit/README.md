# Circuit Breaker Module

The Circuit module provides system-wide resilience and protection mechanisms for the Stateset blockchain, implementing circuit breakers, rate limiting, and emergency controls.

## Overview

The circuit breaker module provides:
- **Global Pause**: Emergency system-wide transaction halt
- **Module Circuit Breakers**: Per-module transaction control
- **Rate Limiting**: Configurable transaction rate limits
- **Oracle Deviation Protection**: Price manipulation prevention
- **Liquidation Surge Protection**: Cascading failure prevention

## Circuit Breaker States

| State | Description |
|-------|-------------|
| `CLOSED` | Normal operation, all transactions allowed |
| `OPEN` | Circuit tripped, transactions blocked |
| `HALF_OPEN` | Recovery mode, limited transactions allowed |

## Features

### Global Pause
- Immediate system-wide transaction halt
- Optional auto-resume duration
- Authority-controlled activation/deactivation

### Module-Level Circuit Breakers
- Independent circuit state per module
- Selective message type disabling
- Automatic recovery after cooldown

### Rate Limiting
- Per-address rate limits
- Global rate limits
- Sliding window implementation
- Configurable limits per message type

### Oracle Deviation Protection
- Price deviation thresholds
- Staleness detection
- Daily deviation caps

### Liquidation Surge Protection
- Maximum liquidations per block
- Maximum liquidation value per block
- Automatic cooldown periods

## Messages

| Message | Description |
|---------|-------------|
| `MsgPauseSystem` | Activate global pause |
| `MsgResumeSystem` | Deactivate global pause |
| `MsgTripCircuit` | Trip a module circuit breaker |
| `MsgResetCircuit` | Reset a module circuit breaker |
| `MsgUpdateParams` | Update module parameters |

## Parameters

| Parameter | Default | Description |
|-----------|---------|-------------|
| `global_rate_limit` | 1000/min | Global transaction rate |
| `per_address_rate_limit` | 100/min | Per-address transaction rate |
| `max_liquidations_per_block` | 10 | Liquidation surge limit |
| `max_liquidation_value_per_block` | 1M | Value-based surge limit |
| `price_deviation_threshold` | 5% | Oracle deviation trigger |
| `staleness_threshold` | 1 hour | Price staleness limit |

## Ante Handler Decorators

The module provides security decorators for the ante handler:

1. **CombinedSecurityDecorator**: Global pause, module circuits, rate limits
2. **CircuitBreakerDecorator**: Per-message circuit checking
3. **LiquidationSurgeDecorator**: Per-block liquidation limits

## Events

| Event | Attributes |
|-------|------------|
| `global_pause` | authority, auto_resume_height |
| `global_resume` | authority |
| `circuit_tripped` | module, reason |
| `circuit_reset` | module |
| `rate_limit_exceeded` | address, message_type |

## State

| Key | Value |
|-----|-------|
| `0x01` | GlobalPauseState |
| `0x02{module}` | ModuleCircuitState |
| `0x03{address}` | RateLimitState |
| `0x04` | Params |
