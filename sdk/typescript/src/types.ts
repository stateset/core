/**
 * Type definitions for Stateset Core SDK
 */

import type { Coin } from '@cosmjs/stargate';

// ============================================================================
// Common Types
// ============================================================================

export interface PaginationParams {
  offset?: number;
  limit?: number;
  countTotal?: boolean;
  reverse?: boolean;
}

export interface PaginationResponse {
  nextKey?: Uint8Array;
  total?: string;
}

// ============================================================================
// Stablecoin Module Types
// ============================================================================

export interface Vault {
  id: string;
  owner: string;
  collateral: Coin;
  debt: Coin;
  collateralDenom: string;
  lastAccrued: number;
}

export interface CreateVaultParams {
  collateral: Coin;
  initialDebt?: Coin;
}

export interface MintParams {
  vaultId: string;
  amount: Coin;
}

export interface RepayParams {
  vaultId: string;
  amount: Coin;
}

export interface WithdrawParams {
  vaultId: string;
  amount: Coin;
}

export interface DepositParams {
  vaultId: string;
  amount: Coin;
}

export interface LiquidationInfo {
  vaultId: string;
  collateralRatio: number;
  isLiquidatable: boolean;
  liquidationPrice: string;
}

// ============================================================================
// Settlement Module Types
// ============================================================================

export interface Settlement {
  id: string;
  sender: string;
  recipient: string;
  amount: Coin;
  status: SettlementStatus;
  createdAt: Date;
  completedAt?: Date;
}

export type SettlementStatus =
  | 'pending'
  | 'completed'
  | 'failed'
  | 'cancelled';

export interface EscrowSettlement {
  id: string;
  depositor: string;
  recipient: string;
  amount: Coin;
  releaseTime: Date;
  conditions: string;
  status: EscrowStatus;
}

export type EscrowStatus =
  | 'active'
  | 'released'
  | 'refunded'
  | 'disputed';

export interface PaymentChannel {
  id: string;
  sender: string;
  recipient: string;
  totalDeposit: Coin;
  withdrawn: Coin;
  nonce: number;
  expiresAt: Date;
}

export interface MerchantConfig {
  address: string;
  name: string;
  feeRate: number;
  webhookUrl?: string;
  isActive: boolean;
}

export interface InstantTransferParams {
  recipient: string;
  amount: Coin;
  memo?: string;
}

export interface CreateEscrowParams {
  recipient: string;
  amount: Coin;
  releaseTime: Date;
  conditions?: string;
}

export interface BatchPaymentParams {
  payments: Array<{
    recipient: string;
    amount: Coin;
  }>;
  memo?: string;
}

// ============================================================================
// Treasury Module Types
// ============================================================================

export interface SpendProposal {
  id: string;
  proposer: string;
  recipient: string;
  amount: Coin[];
  category: string;
  description: string;
  status: ProposalStatus;
  createdAt: Date;
  executeAfter: Date;
  expiresAt: Date;
  executedAt?: Date;
}

export type ProposalStatus =
  | 'pending'
  | 'executed'
  | 'cancelled'
  | 'expired';

export interface Budget {
  category: string;
  totalLimit: Coin[];
  periodLimit: Coin[];
  periodSpent: Coin[];
  totalSpent: Coin[];
  periodStart: Date;
  periodDuration: number; // nanoseconds
  enabled: boolean;
}

export interface Allocation {
  recipient: string;
  totalAllocated: Coin[];
  totalDisbursed: Coin[];
  pending: Coin[];
  lastUpdated: Date;
}

export interface TreasuryStats {
  balance: Coin[];
  totalAllocated: Coin[];
  totalDisbursed: Coin[];
  pendingProposals: number;
  executedProposals: number;
}

export interface ProposeSpendParams {
  recipient: string;
  amount: Coin[];
  category: string;
  description: string;
  timelockSeconds: number;
}

export interface SetBudgetParams {
  category: string;
  totalLimit: Coin[];
  periodLimit: Coin[];
  periodDurationDays: number;
  enabled: boolean;
}

export interface TreasuryParams {
  minTimelockDuration: number;
  maxTimelockDuration: number;
  proposalExpiryDuration: number;
  maxPendingProposals: number;
  baseBurnRate: number;
  validatorRewardRate: number;
  communityPoolRate: number;
}

// ============================================================================
// Oracle Module Types
// ============================================================================

export interface Price {
  denom: string;
  amount: string;
  lastUpdater: string;
  lastHeight: string;
  updatedAt: Date;
}

export interface PriceHistory {
  denom: string;
  prices: PricePoint[];
}

export interface PricePoint {
  amount: string;
  timestamp: Date;
  height: string;
}

export interface OracleProvider {
  address: string;
  isActive: boolean;
  slashed: boolean;
  slashCount: number;
  totalSubmissions: number;
  successfulSubmissions: number;
}

export interface OracleConfig {
  denom: string;
  enabled: boolean;
  minUpdateIntervalSeconds: number;
  maxDeviationBps: number;
  stalenessThresholdSeconds: number;
}

// ============================================================================
// Compliance Module Types
// ============================================================================

export type ComplianceStatus =
  | 'compliant'
  | 'pending'
  | 'blocked'
  | 'sanctioned'
  | 'expired';

export interface ComplianceProfile {
  address: string;
  status: ComplianceStatus;
  riskLevel: RiskLevel;
  kycVerified: boolean;
  jurisdiction: string;
  dailyLimit: Coin;
  monthlyLimit: Coin;
  lastVerified: Date;
  expiresAt?: Date;
}

export type RiskLevel = 'low' | 'medium' | 'high' | 'critical';

export interface TransactionLimits {
  dailyLimit: Coin;
  monthlyLimit: Coin;
  singleTransactionLimit: Coin;
  dailyUsed: Coin;
  monthlyUsed: Coin;
}

// ============================================================================
// Circuit Breaker Types
// ============================================================================

export type CircuitState = 'closed' | 'open' | 'half_open';

export interface CircuitStatus {
  moduleName: string;
  state: CircuitState;
  tripCount: number;
  lastTripped?: Date;
  lastReset?: Date;
}

export interface RateLimitStatus {
  globalRequestsPerMinute: number;
  globalLimit: number;
  addressRequestsPerMinute: Map<string, number>;
  addressLimit: number;
}

export interface SystemStatus {
  globalPaused: boolean;
  pausedAt?: Date;
  pausedBy?: string;
  circuits: CircuitStatus[];
  rateLimits: RateLimitStatus;
}

// ============================================================================
// Transaction Types
// ============================================================================

export interface TxResult {
  transactionHash: string;
  height: number;
  gasUsed: number;
  gasWanted: number;
  events: TxEvent[];
}

export interface TxEvent {
  type: string;
  attributes: Array<{
    key: string;
    value: string;
  }>;
}

export interface SigningOptions {
  gas?: number;
  gasPrice?: string;
  memo?: string;
}
