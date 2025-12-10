/**
 * Treasury Module
 *
 * Manage treasury spend proposals, budgets, and fund allocations.
 * All spend proposals are subject to governance timelock.
 */

import type { Coin } from '@cosmjs/stargate';
import type { StatesetClient } from '../client';
import type {
  SpendProposal,
  Budget,
  Allocation,
  TreasuryStats,
  TreasuryParams,
  ProposeSpendParams,
  SetBudgetParams,
  TxResult,
  PaginationParams,
} from '../types';
import { BUDGET_CATEGORIES } from '../constants';

export class TreasuryModule {
  constructor(private readonly client: StatesetClient) {}

  // ============================================================================
  // Queries
  // ============================================================================

  /**
   * Get treasury balance
   */
  async getBalance(): Promise<Coin[]> {
    const result = await this.client.queryModule('treasury', 'balance', {}) as {
      balance: Coin[];
    };
    return result.balance || [];
  }

  /**
   * Get treasury parameters
   */
  async getParams(): Promise<TreasuryParams> {
    const result = await this.client.queryModule('treasury', 'params', {}) as {
      params: TreasuryParams;
    };
    return result.params;
  }

  /**
   * Get treasury statistics
   */
  async getStats(): Promise<TreasuryStats> {
    const result = await this.client.queryModule('treasury', 'stats', {}) as {
      balance: Coin[];
      total_allocated: Coin[];
      total_disbursed: Coin[];
      pending_proposals: number;
      executed_proposals: number;
    };
    return {
      balance: result.balance,
      totalAllocated: result.total_allocated,
      totalDisbursed: result.total_disbursed,
      pendingProposals: result.pending_proposals,
      executedProposals: result.executed_proposals,
    };
  }

  /**
   * Get a spend proposal by ID
   */
  async getProposal(proposalId: string): Promise<SpendProposal | null> {
    try {
      const result = await this.client.queryModule('treasury', 'proposal', {
        proposal_id: proposalId,
      });
      return this.parseProposal(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all spend proposals
   */
  async getProposals(
    status?: string,
    pagination?: PaginationParams,
  ): Promise<SpendProposal[]> {
    const result = await this.client.queryModule('treasury', 'proposals', {
      status,
      ...pagination,
    }) as { proposals: unknown[] };
    return (result.proposals || []).map((p) => this.parseProposal(p));
  }

  /**
   * Get pending proposals that can be executed
   */
  async getExecutableProposals(): Promise<SpendProposal[]> {
    const result = await this.client.queryModule(
      'treasury',
      'executable_proposals',
      {},
    ) as { proposals: unknown[] };
    return (result.proposals || []).map((p) => this.parseProposal(p));
  }

  /**
   * Get budget for a category
   */
  async getBudget(category: string): Promise<Budget | null> {
    try {
      const result = await this.client.queryModule('treasury', 'budget', {
        category,
      });
      return this.parseBudget(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all budgets
   */
  async getAllBudgets(): Promise<Budget[]> {
    const result = await this.client.queryModule('treasury', 'budgets', {}) as {
      budgets: unknown[];
    };
    return (result.budgets || []).map((b) => this.parseBudget(b));
  }

  /**
   * Get allocation for a recipient
   */
  async getAllocation(recipient: string): Promise<Allocation | null> {
    try {
      const result = await this.client.queryModule('treasury', 'allocation', {
        recipient,
      });
      return this.parseAllocation(result);
    } catch {
      return null;
    }
  }

  /**
   * Check if a category is valid
   */
  isValidCategory(category: string): boolean {
    return BUDGET_CATEGORIES.includes(category as typeof BUDGET_CATEGORIES[number]);
  }

  // ============================================================================
  // Transactions
  // ============================================================================

  /**
   * Create a spend proposal (requires authority)
   */
  async proposeSpend(params: ProposeSpendParams): Promise<TxResult & { proposalId: string }> {
    if (!this.isValidCategory(params.category)) {
      throw new Error(`Invalid category: ${params.category}`);
    }

    const msg = {
      typeUrl: '/stateset.treasury.MsgProposeSpend',
      value: {
        authority: this.client.getAddress(),
        recipient: params.recipient,
        amount: params.amount,
        category: params.category,
        description: params.description,
        timelockSeconds: params.timelockSeconds.toString(),
      },
    };

    const result = await this.client.sendTx([msg], `Treasury spend: ${params.description}`);
    return {
      ...result,
      proposalId: result.transactionHash.slice(0, 16),
      events: [],
    };
  }

  /**
   * Execute a spend proposal after timelock expires
   */
  async executeSpend(proposalId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.treasury.MsgExecuteSpend',
      value: {
        authority: this.client.getAddress(),
        proposalId,
      },
    };

    const result = await this.client.sendTx([msg], 'Execute treasury spend');
    return { ...result, events: [] };
  }

  /**
   * Cancel a pending spend proposal
   */
  async cancelSpend(proposalId: string, reason: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.treasury.MsgCancelSpend',
      value: {
        authority: this.client.getAddress(),
        proposalId,
        reason,
      },
    };

    const result = await this.client.sendTx([msg], 'Cancel treasury spend');
    return { ...result, events: [] };
  }

  /**
   * Set budget for a category (requires authority)
   */
  async setBudget(params: SetBudgetParams): Promise<TxResult> {
    if (!this.isValidCategory(params.category)) {
      throw new Error(`Invalid category: ${params.category}`);
    }

    const msg = {
      typeUrl: '/stateset.treasury.MsgSetBudget',
      value: {
        authority: this.client.getAddress(),
        category: params.category,
        totalLimit: params.totalLimit,
        periodLimit: params.periodLimit,
        periodDuration: (params.periodDurationDays * 24 * 60 * 60 * 1e9).toString(), // nanoseconds
        enabled: params.enabled,
      },
    };

    const result = await this.client.sendTx([msg], `Set budget: ${params.category}`);
    return { ...result, events: [] };
  }

  /**
   * Update treasury parameters (requires authority)
   */
  async updateParams(params: Partial<TreasuryParams>): Promise<TxResult> {
    const currentParams = await this.getParams();
    const newParams = { ...currentParams, ...params };

    const msg = {
      typeUrl: '/stateset.treasury.MsgUpdateParams',
      value: {
        authority: this.client.getAddress(),
        params: newParams,
      },
    };

    const result = await this.client.sendTx([msg], 'Update treasury params');
    return { ...result, events: [] };
  }

  // ============================================================================
  // Helpers
  // ============================================================================

  /**
   * Calculate time until a proposal can be executed
   */
  async getTimeUntilExecutable(proposalId: string): Promise<number> {
    const proposal = await this.getProposal(proposalId);
    if (!proposal) {
      throw new Error(`Proposal ${proposalId} not found`);
    }

    const now = Date.now();
    const executeAfter = proposal.executeAfter.getTime();
    return Math.max(0, executeAfter - now);
  }

  /**
   * Check if a proposal can be executed now
   */
  async canExecute(proposalId: string): Promise<boolean> {
    const proposal = await this.getProposal(proposalId);
    if (!proposal || proposal.status !== 'pending') {
      return false;
    }

    const now = Date.now();
    return (
      now >= proposal.executeAfter.getTime() &&
      now < proposal.expiresAt.getTime()
    );
  }

  private parseProposal(data: unknown): SpendProposal {
    const p = data as {
      id: string;
      proposer: string;
      recipient: string;
      amount: Coin[];
      category: string;
      description: string;
      status: string;
      created_at: string;
      execute_after: string;
      expires_at: string;
      executed_at?: string;
    };
    return {
      id: p.id,
      proposer: p.proposer,
      recipient: p.recipient,
      amount: p.amount,
      category: p.category,
      description: p.description,
      status: p.status as SpendProposal['status'],
      createdAt: new Date(p.created_at),
      executeAfter: new Date(p.execute_after),
      expiresAt: new Date(p.expires_at),
      executedAt: p.executed_at ? new Date(p.executed_at) : undefined,
    };
  }

  private parseBudget(data: unknown): Budget {
    const b = data as {
      category: string;
      total_limit: Coin[];
      period_limit: Coin[];
      period_spent: Coin[];
      total_spent: Coin[];
      period_start: string;
      period_duration: string;
      enabled: boolean;
    };
    return {
      category: b.category,
      totalLimit: b.total_limit,
      periodLimit: b.period_limit,
      periodSpent: b.period_spent,
      totalSpent: b.total_spent,
      periodStart: new Date(b.period_start),
      periodDuration: parseInt(b.period_duration, 10),
      enabled: b.enabled,
    };
  }

  private parseAllocation(data: unknown): Allocation {
    const a = data as {
      recipient: string;
      total_allocated: Coin[];
      total_disbursed: Coin[];
      pending: Coin[];
      last_updated: string;
    };
    return {
      recipient: a.recipient,
      totalAllocated: a.total_allocated,
      totalDisbursed: a.total_disbursed,
      pending: a.pending,
      lastUpdated: new Date(a.last_updated),
    };
  }
}
