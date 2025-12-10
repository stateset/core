/**
 * Utility functions for Stateset SDK
 */

import type { Coin } from '@cosmjs/stargate';
import { STABLECOIN_DENOM, NATIVE_DENOM, MIN_COLLATERAL_RATIO } from './constants';

/**
 * Format coins for display
 */
export function formatCoins(coins: Coin[]): string {
  return coins
    .map((c) => formatCoin(c))
    .join(', ');
}

/**
 * Format a single coin for display
 */
export function formatCoin(coin: Coin): string {
  const amount = formatAmount(coin.amount, getDenomDecimals(coin.denom));
  const symbol = getDenomSymbol(coin.denom);
  return `${amount} ${symbol}`;
}

/**
 * Get decimals for a denom
 */
export function getDenomDecimals(denom: string): number {
  switch (denom) {
    case NATIVE_DENOM:
    case STABLECOIN_DENOM:
      return 6;
    default:
      return 6;
  }
}

/**
 * Get display symbol for a denom
 */
export function getDenomSymbol(denom: string): string {
  switch (denom) {
    case NATIVE_DENOM:
      return 'STATE';
    case STABLECOIN_DENOM:
      return 'ssUSD';
    default:
      return denom.toUpperCase();
  }
}

/**
 * Format amount with proper decimals
 */
export function formatAmount(amount: string, decimals: number): string {
  const value = parseInt(amount, 10);
  const divisor = Math.pow(10, decimals);
  const formatted = (value / divisor).toFixed(decimals);

  // Remove trailing zeros
  return parseFloat(formatted).toString();
}

/**
 * Parse human-readable amount to chain format
 */
export function parseAmount(amount: string | number, decimals: number): string {
  const value = typeof amount === 'string' ? parseFloat(amount) : amount;
  const multiplier = Math.pow(10, decimals);
  return Math.floor(value * multiplier).toString();
}

/**
 * Parse human-readable coins to chain format
 */
export function parseCoins(amount: string | number, denom: string): Coin {
  const decimals = getDenomDecimals(denom);
  return {
    denom,
    amount: parseAmount(amount, decimals),
  };
}

/**
 * Calculate collateral ratio
 */
export function calculateCollateralRatio(
  collateralValue: number,
  debtValue: number,
): number {
  if (debtValue === 0) {
    return Infinity;
  }
  return collateralValue / debtValue;
}

/**
 * Calculate maximum mintable stablecoin
 */
export function calculateMaxMintable(
  collateralAmount: string,
  collateralPrice: string,
  existingDebt: string = '0',
  minRatio: number = MIN_COLLATERAL_RATIO,
): string {
  const collateralValue = parseFloat(collateralAmount) * parseFloat(collateralPrice);
  const maxDebt = collateralValue / minRatio;
  const existingDebtValue = parseFloat(existingDebt);
  const mintable = Math.max(0, maxDebt - existingDebtValue);
  return Math.floor(mintable).toString();
}

/**
 * Calculate liquidation price
 */
export function calculateLiquidationPrice(
  collateralAmount: string,
  debtAmount: string,
  liquidationRatio: number = MIN_COLLATERAL_RATIO,
): string {
  const collateral = parseFloat(collateralAmount);
  const debt = parseFloat(debtAmount);

  if (collateral === 0 || debt === 0) {
    return '0';
  }

  const liquidationPrice = (debt * liquidationRatio) / collateral;
  return liquidationPrice.toFixed(6);
}

/**
 * Check if a vault is safe
 */
export function isVaultSafe(
  collateralRatio: number,
  minRatio: number = MIN_COLLATERAL_RATIO,
): boolean {
  return collateralRatio >= minRatio;
}

/**
 * Calculate health factor (1.0 = at liquidation threshold)
 */
export function calculateHealthFactor(
  collateralRatio: number,
  liquidationRatio: number = 1.3, // 130%
): number {
  return collateralRatio / liquidationRatio;
}

/**
 * Format percentage
 */
export function formatPercentage(value: number, decimals: number = 2): string {
  return `${(value * 100).toFixed(decimals)}%`;
}

/**
 * Format collateral ratio for display
 */
export function formatCollateralRatio(ratio: number): string {
  if (!isFinite(ratio)) {
    return '∞';
  }
  return formatPercentage(ratio);
}

/**
 * Calculate annual stability fee
 */
export function calculateStabilityFee(
  debtAmount: string,
  annualRate: number,
  daysElapsed: number,
): string {
  const debt = parseFloat(debtAmount);
  const dailyRate = annualRate / 365;
  const fee = debt * dailyRate * daysElapsed;
  return Math.floor(fee).toString();
}

/**
 * Validate bech32 address
 */
export function isValidAddress(address: string, prefix: string = 'stateset'): boolean {
  try {
    if (!address.startsWith(prefix)) {
      return false;
    }
    // Basic length check (bech32 addresses are typically 39-59 chars)
    return address.length >= 39 && address.length <= 59;
  } catch {
    return false;
  }
}

/**
 * Truncate address for display
 */
export function truncateAddress(address: string, chars: number = 8): string {
  if (address.length <= chars * 2 + 3) {
    return address;
  }
  return `${address.slice(0, chars)}...${address.slice(-chars)}`;
}

/**
 * Convert nanoseconds to human-readable duration
 */
export function formatDuration(nanoseconds: number): string {
  const seconds = nanoseconds / 1e9;
  const minutes = seconds / 60;
  const hours = minutes / 60;
  const days = hours / 24;

  if (days >= 1) {
    return `${Math.floor(days)} day${days >= 2 ? 's' : ''}`;
  }
  if (hours >= 1) {
    return `${Math.floor(hours)} hour${hours >= 2 ? 's' : ''}`;
  }
  if (minutes >= 1) {
    return `${Math.floor(minutes)} minute${minutes >= 2 ? 's' : ''}`;
  }
  return `${Math.floor(seconds)} second${seconds >= 2 ? 's' : ''}`;
}

/**
 * Sleep for a given number of milliseconds
 */
export function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
