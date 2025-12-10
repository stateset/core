/**
 * Settlement Module
 *
 * Handle instant transfers, escrow settlements, batch payments,
 * and payment channels on Stateset.
 */

import type { StatesetClient } from '../client';
import type {
  Settlement,
  EscrowSettlement,
  PaymentChannel,
  MerchantConfig,
  InstantTransferParams,
  CreateEscrowParams,
  BatchPaymentParams,
  TxResult,
  PaginationParams,
} from '../types';

export class SettlementModule {
  constructor(private readonly client: StatesetClient) {}

  // ============================================================================
  // Queries
  // ============================================================================

  /**
   * Get a settlement by ID
   */
  async getSettlement(settlementId: string): Promise<Settlement | null> {
    try {
      const result = await this.client.queryModule('settlement', 'settlement', {
        settlement_id: settlementId,
      });
      return this.parseSettlement(result);
    } catch {
      return null;
    }
  }

  /**
   * Get settlements for an address
   */
  async getSettlementsByAddress(
    address: string,
    pagination?: PaginationParams,
  ): Promise<Settlement[]> {
    const result = await this.client.queryModule('settlement', 'settlements', {
      address,
      ...pagination,
    }) as { settlements: unknown[] };
    return (result.settlements || []).map((s) => this.parseSettlement(s));
  }

  /**
   * Get an escrow by ID
   */
  async getEscrow(escrowId: string): Promise<EscrowSettlement | null> {
    try {
      const result = await this.client.queryModule('settlement', 'escrow', {
        escrow_id: escrowId,
      });
      return this.parseEscrow(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all active escrows for an address
   */
  async getActiveEscrows(address: string): Promise<EscrowSettlement[]> {
    const result = await this.client.queryModule('settlement', 'active_escrows', {
      address,
    }) as { escrows: unknown[] };
    return (result.escrows || []).map((e) => this.parseEscrow(e));
  }

  /**
   * Get a payment channel by ID
   */
  async getPaymentChannel(channelId: string): Promise<PaymentChannel | null> {
    try {
      const result = await this.client.queryModule('settlement', 'channel', {
        channel_id: channelId,
      });
      return this.parseChannel(result);
    } catch {
      return null;
    }
  }

  /**
   * Get merchant configuration
   */
  async getMerchant(address: string): Promise<MerchantConfig | null> {
    try {
      const result = await this.client.queryModule('settlement', 'merchant', {
        address,
      });
      return this.parseMerchant(result);
    } catch {
      return null;
    }
  }

  // ============================================================================
  // Transactions
  // ============================================================================

  /**
   * Send an instant transfer
   */
  async instantTransfer(params: InstantTransferParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgInstantTransfer',
      value: {
        sender: this.client.getAddress(),
        recipient: params.recipient,
        amount: params.amount,
        memo: params.memo || '',
      },
    };

    const result = await this.client.sendTx([msg], params.memo);
    return { ...result, events: [] };
  }

  /**
   * Create an escrow settlement
   */
  async createEscrow(params: CreateEscrowParams): Promise<TxResult & { escrowId: string }> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgCreateEscrow',
      value: {
        depositor: this.client.getAddress(),
        recipient: params.recipient,
        amount: params.amount,
        releaseTime: params.releaseTime.toISOString(),
        conditions: params.conditions || '',
      },
    };

    const result = await this.client.sendTx([msg], 'Create escrow');
    return {
      ...result,
      escrowId: result.transactionHash.slice(0, 16),
      events: [],
    };
  }

  /**
   * Release funds from escrow
   */
  async releaseEscrow(escrowId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgReleaseEscrow',
      value: {
        sender: this.client.getAddress(),
        escrowId,
      },
    };

    const result = await this.client.sendTx([msg], 'Release escrow');
    return { ...result, events: [] };
  }

  /**
   * Refund escrow to depositor
   */
  async refundEscrow(escrowId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgRefundEscrow',
      value: {
        sender: this.client.getAddress(),
        escrowId,
      },
    };

    const result = await this.client.sendTx([msg], 'Refund escrow');
    return { ...result, events: [] };
  }

  /**
   * Send a batch payment
   */
  async batchPayment(params: BatchPaymentParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgBatchPayment',
      value: {
        sender: this.client.getAddress(),
        payments: params.payments,
        memo: params.memo || '',
      },
    };

    const result = await this.client.sendTx([msg], params.memo || 'Batch payment');
    return { ...result, events: [] };
  }

  /**
   * Open a payment channel
   */
  async openChannel(
    recipient: string,
    deposit: { denom: string; amount: string },
    expiresAt: Date,
  ): Promise<TxResult & { channelId: string }> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgOpenChannel',
      value: {
        sender: this.client.getAddress(),
        recipient,
        deposit,
        expiresAt: expiresAt.toISOString(),
      },
    };

    const result = await this.client.sendTx([msg], 'Open payment channel');
    return {
      ...result,
      channelId: result.transactionHash.slice(0, 16),
      events: [],
    };
  }

  /**
   * Close a payment channel
   */
  async closeChannel(
    channelId: string,
    finalAmount: { denom: string; amount: string },
    signature: string,
  ): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgCloseChannel',
      value: {
        sender: this.client.getAddress(),
        channelId,
        finalAmount,
        signature,
      },
    };

    const result = await this.client.sendTx([msg], 'Close payment channel');
    return { ...result, events: [] };
  }

  /**
   * Register as a merchant
   */
  async registerMerchant(
    name: string,
    feeRate: number,
    webhookUrl?: string,
  ): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.settlement.MsgRegisterMerchant',
      value: {
        address: this.client.getAddress(),
        name,
        feeRate: Math.floor(feeRate * 10000), // Convert to basis points
        webhookUrl: webhookUrl || '',
      },
    };

    const result = await this.client.sendTx([msg], 'Register merchant');
    return { ...result, events: [] };
  }

  // ============================================================================
  // Helpers
  // ============================================================================

  private parseSettlement(data: unknown): Settlement {
    const s = data as {
      id: string;
      sender: string;
      recipient: string;
      amount: { denom: string; amount: string };
      status: string;
      created_at: string;
      completed_at?: string;
    };
    return {
      id: s.id,
      sender: s.sender,
      recipient: s.recipient,
      amount: s.amount,
      status: s.status as Settlement['status'],
      createdAt: new Date(s.created_at),
      completedAt: s.completed_at ? new Date(s.completed_at) : undefined,
    };
  }

  private parseEscrow(data: unknown): EscrowSettlement {
    const e = data as {
      id: string;
      depositor: string;
      recipient: string;
      amount: { denom: string; amount: string };
      release_time: string;
      conditions: string;
      status: string;
    };
    return {
      id: e.id,
      depositor: e.depositor,
      recipient: e.recipient,
      amount: e.amount,
      releaseTime: new Date(e.release_time),
      conditions: e.conditions,
      status: e.status as EscrowSettlement['status'],
    };
  }

  private parseChannel(data: unknown): PaymentChannel {
    const c = data as {
      id: string;
      sender: string;
      recipient: string;
      total_deposit: { denom: string; amount: string };
      withdrawn: { denom: string; amount: string };
      nonce: number;
      expires_at: string;
    };
    return {
      id: c.id,
      sender: c.sender,
      recipient: c.recipient,
      totalDeposit: c.total_deposit,
      withdrawn: c.withdrawn,
      nonce: c.nonce,
      expiresAt: new Date(c.expires_at),
    };
  }

  private parseMerchant(data: unknown): MerchantConfig {
    const m = data as {
      address: string;
      name: string;
      fee_rate: number;
      webhook_url?: string;
      is_active: boolean;
    };
    return {
      address: m.address,
      name: m.name,
      feeRate: m.fee_rate / 10000, // Convert from basis points
      webhookUrl: m.webhook_url,
      isActive: m.is_active,
    };
  }
}
