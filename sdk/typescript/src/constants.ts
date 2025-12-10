/**
 * Network endpoints and constants
 */

// Mainnet
export const MAINNET_RPC = 'https://rpc.stateset.network';
export const MAINNET_REST = 'https://api.stateset.network';
export const CHAIN_ID_MAINNET = 'stateset-1';

// Testnet
export const TESTNET_RPC = 'https://rpc.testnet.stateset.network';
export const TESTNET_REST = 'https://api.testnet.stateset.network';
export const CHAIN_ID_TESTNET = 'stateset-testnet-1';

// Token denominations
export const NATIVE_DENOM = 'ustate';
export const STABLECOIN_DENOM = 'ssusd';

// Module names
export const MODULE_STABLECOIN = 'stablecoin';
export const MODULE_SETTLEMENT = 'settlement';
export const MODULE_TREASURY = 'treasury';
export const MODULE_ORACLE = 'oracle';
export const MODULE_COMPLIANCE = 'compliance';
export const MODULE_CIRCUIT = 'circuit';

// Default gas
export const DEFAULT_GAS_PRICE = '0.025ustate';
export const DEFAULT_GAS_LIMIT = 200000;

// Stablecoin parameters
export const MIN_COLLATERAL_RATIO = 1.5; // 150%
export const LIQUIDATION_THRESHOLD = 1.3; // 130%
export const LIQUIDATION_PENALTY = 0.1; // 10%
export const STABILITY_FEE_RATE = 0.02; // 2% annual

// Treasury parameters
export const MIN_TIMELOCK_HOURS = 24;
export const MAX_TIMELOCK_DAYS = 30;
export const PROPOSAL_EXPIRY_DAYS = 7;

// Budget categories
export const BUDGET_CATEGORIES = [
  'development',
  'marketing',
  'operations',
  'grants',
  'security',
  'infrastructure',
  'reserve',
] as const;

export type BudgetCategory = (typeof BUDGET_CATEGORIES)[number];
