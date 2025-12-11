/**
 * Stablecoin Module
 *
 * Interact with the Stateset stablecoin (ssUSD) system.
 * Supports vault-based CDPs and reserve-backed issuance/redemption.
 */

import type { Coin } from '@cosmjs/stargate';
import type { StatesetClient } from '../client';
import type {
  Vault,
  CreateVaultParams,
  MintParams,
  RepayParams,
  DepositParams,
  WithdrawParams,
  TxResult,
} from '../types';
import { STABLECOIN_DENOM } from '../constants';

export class StablecoinModule {
  constructor(private readonly client: StatesetClient) {}

  // ============================================================================
  // Queries
  // ============================================================================

  /**
   * Get a vault by ID
   */
  async getVault(vaultId: string): Promise<Vault | null> {
    try {
      const result = await this.client.queryModule('stablecoin', 'vault', {
        vault_id: vaultId,
      });
      return this.parseVault(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all vaults, optionally filtered by owner.
   */
  async getVaults(owner?: string): Promise<Vault[]> {
    const result = await this.client.queryModule('stablecoin', 'vaults', {
      owner: owner || '',
    }) as { vaults: unknown[] };
    return (result.vaults || []).map((v) => this.parseVault(v));
  }

  // ============================================================================
  // Transactions
  // ============================================================================

  /**
   * Create a new vault
   */
  async createVault(params: CreateVaultParams): Promise<TxResult & { vaultId: string }> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgCreateVault',
      value: {
        owner: this.client.getAddress(),
        collateral: params.collateral,
        debt: params.initialDebt || { denom: STABLECOIN_DENOM, amount: '0' },
      },
    };

    const result = await this.client.sendTx([msg], 'Create stablecoin vault');

    // Extract vault ID from events
    const vaultId = this.extractVaultIdFromEvents(result);

    return {
      ...result,
      vaultId,
      events: [],
    };
  }

  /**
   * Deposit additional collateral to a vault
   */
  async depositCollateral(params: DepositParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgDepositCollateral',
      value: {
        owner: this.client.getAddress(),
        vaultId: params.vaultId,
        amount: params.amount,
      },
    };

    const result = await this.client.sendTx([msg], 'Deposit collateral');
    return { ...result, events: [] };
  }

  /**
   * Withdraw collateral from a vault
   */
  async withdrawCollateral(params: WithdrawParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgWithdrawCollateral',
      value: {
        owner: this.client.getAddress(),
        vaultId: params.vaultId,
        amount: params.amount,
      },
    };

    const result = await this.client.sendTx([msg], 'Withdraw collateral');
    return { ...result, events: [] };
  }

  /**
   * Mint stablecoins against vault collateral
   */
  async mint(params: MintParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgMintStablecoin',
      value: {
        owner: this.client.getAddress(),
        vaultId: params.vaultId,
        amount: params.amount,
      },
    };

    const result = await this.client.sendTx([msg], 'Mint stablecoin');
    return { ...result, events: [] };
  }

  /**
   * Repay stablecoin debt
   */
  async repay(params: RepayParams): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgRepayStablecoin',
      value: {
        owner: this.client.getAddress(),
        vaultId: params.vaultId,
        amount: params.amount,
      },
    };

    const result = await this.client.sendTx([msg], 'Repay debt');
    return { ...result, events: [] };
  }

  /**
   * Liquidate an undercollateralized vault
   */
  async liquidate(vaultId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgLiquidateVault',
      value: {
        liquidator: this.client.getAddress(),
        vaultId,
      },
    };

    const result = await this.client.sendTx([msg], 'Liquidate vault');
    return { ...result, events: [] };
  }

  // ============================================================================
  // Helpers
  // ============================================================================

  private parseVault(data: unknown): Vault {
    const v = data as {
      id: string | number;
      owner: string;
      collateral: { denom: string; amount: string };
      collateral_denom?: string;
      debt: string;
      last_accrued?: string | number;
    };
    const debtAmount = typeof v.debt === 'string' ? v.debt : String(v.debt);
    return {
      id: String(v.id),
      owner: v.owner,
      collateral: v.collateral,
      debt: { denom: STABLECOIN_DENOM, amount: debtAmount },
      collateralDenom: v.collateral_denom || v.collateral.denom,
      lastAccrued: v.last_accrued ? Number(v.last_accrued) : 0,
    };
  }

  private extractVaultIdFromEvents(result: { transactionHash: string }): string {
    // In a real implementation, parse transaction events
    // For now, return a placeholder
    return result.transactionHash.slice(0, 16);
  }
}
