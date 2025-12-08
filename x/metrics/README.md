# Metrics Module

The Metrics module provides on-chain metrics and monitoring capabilities for the Stateset blockchain.

## Overview

The metrics module handles:
- **Performance Metrics**: Track blockchain performance
- **Usage Statistics**: Record module usage patterns
- **Health Monitoring**: Monitor system health

## Features

### Metric Collection
- Block production metrics
- Transaction throughput
- Module-specific statistics

### Aggregation
- Time-based aggregation
- Rolling averages
- Peak detection

## State

| Key | Value |
|-----|-------|
| `0x01{metric_id}` | MetricData |
| `0x02` | AggregatedMetrics |

## Events

| Event | Attributes |
|-------|------------|
| `metric_recorded` | metric_type, value |
| `threshold_exceeded` | metric_type, threshold, value |
