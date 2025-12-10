/**
 * Stablecoin Module
 *
 * Interact with the Stateset stablecoin (ssUSD) system.
 * Create vaults, mint/burn stablecoins, manage collateral.
 */

import type { Coin } from '@cosmjs/stargate';
import type { StatesetClient } from '../client';
import type {
  Vault,
  VaultStats,
  CreateVaultParams,
  MintParams,
  RepayParams,
  DepositParams,
  WithdrawParams,
  LiquidationInfo,
  TxResult,
  PaginationParams,
} from '../types';
import { STABLECOIN_DENOM, MIN_COLLATERAL_RATIO } from '../constants';

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
   * Get all vaults for an owner
   */
  async getVaultsByOwner(
    owner: string,
    pagination?: PaginationParams,
  ): Promise<Vault[]> {
    const result = await this.client.queryModule('stablecoin', 'vaults_by_owner', {
      owner,
      ...pagination,
    }) as { vaults: unknown[] };
    return (result.vaults || []).map((v) => this.parseVault(v));
  }

  /**
   * Get stablecoin module statistics
   */
  async getStats(): Promise<VaultStats> {
    const result = await this.client.queryModule('stablecoin', 'stats', {}) as {
      total_vaults: number;
      total_collateral: { denom: string; amount: string };
      total_debt: { denom: string; amount: string };
      average_collateral_ratio: string;
    };
    return {
      totalVaults: result.total_vaults,
      totalCollateral: result.total_collateral,
      totalDebt: result.total_debt,
      averageCollateralRatio: parseFloat(result.average_collateral_ratio),
    };
  }

  /**
   * Calculate the collateral ratio for a vault
   */
  async calculateCollateralRatio(vaultId: string): Promise<number> {
    const vault = await this.getVault(vaultId);
    if (!vault) {
      throw new Error(`Vault ${vaultId} not found`);
    }

    const price = await this.client.oracle.getPrice(vault.collateral.denom);
    if (!price) {
      throw new Error(`Price not found for ${vault.collateral.denom}`);
    }

    const collateralValue =
      parseFloat(vault.collateral.amount) * parseFloat(price.amount);
    const debtValue = parseFloat(vault.debt.amount);

    if (debtValue === 0) {
      return Infinity;
    }

    return collateralValue / debtValue;
  }

  /**
   * Check if a vault is liquidatable
   */
  async getLiquidationInfo(vaultId: string): Promise<LiquidationInfo> {
    const ratio = await this.calculateCollateralRatio(vaultId);
    const vault = await this.getVault(vaultId);

    if (!vault) {
      throw new Error(`Vault ${vaultId} not found`);
    }

    const price = await this.client.oracle.getPrice(vault.collateral.denom);
    const currentPrice = price ? parseFloat(price.amount) : 0;

    // Calculate liquidation price
    const liquidationPrice =
      vault.debt.amount !== '0'
        ? (parseFloat(vault.debt.amount) * MIN_COLLATERAL_RATIO) /
          parseFloat(vault.collateral.amount)
        : 0;

    return {
      vaultId,
      collateralRatio: ratio,
      isLiquidatable: ratio < MIN_COLLATERAL_RATIO,
      liquidationPrice: liquidationPrice.toFixed(6),
    };
  }

  /**
   * Get all liquidatable vaults
   */
  async getLiquidatableVaults(): Promise<Vault[]> {
    const result = await this.client.queryModule(
      'stablecoin',
      'liquidatable_vaults',
      {},
    ) as { vaults: unknown[] };
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
      typeUrl: '/stateset.stablecoin.MsgRepayDebt',
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
   * Close a vault (repay all debt and withdraw all collateral)
   */
  async closeVault(vaultId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgCloseVault',
      value: {
        owner: this.client.getAddress(),
        vaultId,
      },
    };

    const result = await this.client.sendTx([msg], 'Close vault');
    return { ...result, events: [] };
  }

  /**
   * Liquidate an undercollateralized vault
   */
  async liquidate(vaultId: string): Promise<TxResult> {
    const msg = {
      typeUrl: '/stateset.stablecoin.MsgLiquidate',
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
      id: string;
      owner: string;
      collateral: { denom: string; amount: string };
      debt: { denom: string; amount: string };
      created_at: string;
      last_updated: string;
    };
    return {
      id: v.id,
      owner: v.owner,
      collateral: v.collateral,
      debt: v.debt,
      createdAt: new Date(v.created_at),
      lastUpdated: new Date(v.last_updated),
    };
  }

  private extractVaultIdFromEvents(result: { transactionHash: string }): string {
    // In a real implementation, parse transaction events
    // For now, return a placeholder
    return result.transactionHash.slice(0, 16);
  }
}
