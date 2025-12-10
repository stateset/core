/**
 * Compliance Module
 *
 * Query compliance status and KYC information.
 */

import type { StatesetClient } from '../client';
import type {
  ComplianceProfile,
  TransactionLimits,
  ComplianceStatus,
} from '../types';

export class ComplianceModule {
  constructor(private readonly client: StatesetClient) {}

  // ============================================================================
  // Queries
  // ============================================================================

  /**
   * Get compliance profile for an address
   */
  async getProfile(address: string): Promise<ComplianceProfile | null> {
    try {
      const result = await this.client.queryModule('compliance', 'profile', {
        address,
      });
      return this.parseProfile(result);
    } catch {
      return null;
    }
  }

  /**
   * Check if an address is compliant
   */
  async isCompliant(address: string): Promise<boolean> {
    const profile = await this.getProfile(address);
    if (!profile) {
      return false;
    }
    return profile.status === 'compliant';
  }

  /**
   * Get transaction limits for an address
   */
  async getLimits(address: string): Promise<TransactionLimits | null> {
    try {
      const result = await this.client.queryModule('compliance', 'limits', {
        address,
      });
      return this.parseLimits(result);
    } catch {
      return null;
    }
  }

  /**
   * Check if a transaction amount would exceed limits
   */
  async checkTransactionLimit(
    address: string,
    amount: { denom: string; amount: string },
  ): Promise<{
    allowed: boolean;
    reason?: string;
  }> {
    const limits = await this.getLimits(address);
    if (!limits) {
      return { allowed: true };
    }

    const amountValue = parseInt(amount.amount, 10);
    const singleLimit = parseInt(limits.singleTransactionLimit.amount, 10);

    if (amountValue > singleLimit) {
      return {
        allowed: false,
        reason: `Amount ${amount.amount} exceeds single transaction limit ${singleLimit}`,
      };
    }

    const dailyUsed = parseInt(limits.dailyUsed.amount, 10);
    const dailyLimit = parseInt(limits.dailyLimit.amount, 10);

    if (dailyUsed + amountValue > dailyLimit) {
      return {
        allowed: false,
        reason: `Transaction would exceed daily limit. Used: ${dailyUsed}, Limit: ${dailyLimit}`,
      };
    }

    const monthlyUsed = parseInt(limits.monthlyUsed.amount, 10);
    const monthlyLimit = parseInt(limits.monthlyLimit.amount, 10);

    if (monthlyUsed + amountValue > monthlyLimit) {
      return {
        allowed: false,
        reason: `Transaction would exceed monthly limit. Used: ${monthlyUsed}, Limit: ${monthlyLimit}`,
      };
    }

    return { allowed: true };
  }

  /**
   * Check if an address is on the sanctions list
   */
  async isSanctioned(address: string): Promise<boolean> {
    const profile = await this.getProfile(address);
    if (!profile) {
      return false;
    }
    return profile.status === 'sanctioned';
  }

  /**
   * Get compliance status description
   */
  getStatusDescription(status: ComplianceStatus): string {
    switch (status) {
      case 'compliant':
        return 'Address has passed all compliance checks and is fully verified.';
      case 'pending':
        return 'Compliance verification is in progress.';
      case 'blocked':
        return 'Address has been blocked due to compliance issues.';
      case 'sanctioned':
        return 'Address is on a sanctions list and cannot transact.';
      case 'expired':
        return 'Compliance verification has expired and needs renewal.';
      default:
        return 'Unknown compliance status.';
    }
  }

  /**
   * Get list of blocked jurisdictions
   */
  async getBlockedJurisdictions(): Promise<string[]> {
    const result = await this.client.queryModule(
      'compliance',
      'blocked_jurisdictions',
      {},
    ) as { jurisdictions: string[] };
    return result.jurisdictions || [];
  }

  // ============================================================================
  // Helpers
  // ============================================================================

  private parseProfile(data: unknown): ComplianceProfile {
    const p = data as {
      address: string;
      status: string;
      risk_level: string;
      kyc_verified: boolean;
      jurisdiction: string;
      daily_limit: { denom: string; amount: string };
      monthly_limit: { denom: string; amount: string };
      last_verified: string;
      expires_at?: string;
    };
    return {
      address: p.address,
      status: p.status as ComplianceStatus,
      riskLevel: p.risk_level as ComplianceProfile['riskLevel'],
      kycVerified: p.kyc_verified,
      jurisdiction: p.jurisdiction,
      dailyLimit: p.daily_limit,
      monthlyLimit: p.monthly_limit,
      lastVerified: new Date(p.last_verified),
      expiresAt: p.expires_at ? new Date(p.expires_at) : undefined,
    };
  }

  private parseLimits(data: unknown): TransactionLimits {
    const l = data as {
      daily_limit: { denom: string; amount: string };
      monthly_limit: { denom: string; amount: string };
      single_transaction_limit: { denom: string; amount: string };
      daily_used: { denom: string; amount: string };
      monthly_used: { denom: string; amount: string };
    };
    return {
      dailyLimit: l.daily_limit,
      monthlyLimit: l.monthly_limit,
      singleTransactionLimit: l.single_transaction_limit,
      dailyUsed: l.daily_used,
      monthlyUsed: l.monthly_used,
    };
  }
}
