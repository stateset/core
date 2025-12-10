/**
 * Stateset Core TypeScript SDK
 *
 * Official client library for interacting with Stateset Core blockchain.
 * Supports stablecoin operations, settlements, treasury, and compliance.
 */

export { StatesetClient, StatesetClientOptions } from './client';
export { StablecoinModule } from './modules/stablecoin';
export { SettlementModule } from './modules/settlement';
export { TreasuryModule } from './modules/treasury';
export { OracleModule } from './modules/oracle';
export { ComplianceModule } from './modules/compliance';

// Types
export * from './types';

// Constants
export {
  MAINNET_RPC,
  MAINNET_REST,
  TESTNET_RPC,
  TESTNET_REST,
  CHAIN_ID_MAINNET,
  CHAIN_ID_TESTNET,
  STABLECOIN_DENOM,
  NATIVE_DENOM,
} from './constants';

// Utils
export { formatCoins, parseCoins, calculateCollateralRatio } from './utils';
