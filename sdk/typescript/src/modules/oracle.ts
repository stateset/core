/**
 * Oracle Module
 *
 * Query price feeds and oracle provider information.
 */

import type { StatesetClient } from '../client';
import type {
  Price,
  PriceHistory,
  OracleProvider,
  OracleConfig,
} from '../types';

export class OracleModule {
  constructor(private readonly client: StatesetClient) {}

  // ============================================================================
  // Queries
  // ============================================================================

  /**
   * Get current price for a denom
   */
  async getPrice(denom: string): Promise<Price | null> {
    try {
      const result = await this.client.queryModule('oracle', 'price', {
        denom,
      });
      return this.parsePrice(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all current prices
   */
  async getAllPrices(): Promise<Price[]> {
    const result = await this.client.queryModule('oracle', 'prices', {}) as {
      prices: unknown[];
    };
    return (result.prices || []).map((p) => this.parsePrice(p));
  }

  /**
   * Get price history for a denom
   */
  async getPriceHistory(denom: string, limit?: number): Promise<PriceHistory> {
    const result = await this.client.queryModule('oracle', 'price_history', {
      denom,
      limit: limit || 100,
    }) as {
      denom: string;
      prices: Array<{
        amount: string;
        timestamp: string;
        height: string;
      }>;
    };

    return {
      denom: result.denom,
      prices: result.prices.map((p) => ({
        amount: p.amount,
        timestamp: new Date(p.timestamp),
        height: p.height,
      })),
    };
  }

  /**
   * Get oracle configuration for a denom
   */
  async getConfig(denom: string): Promise<OracleConfig | null> {
    try {
      const result = await this.client.queryModule('oracle', 'config', {
        denom,
      });
      return this.parseConfig(result);
    } catch {
      return null;
    }
  }

  /**
   * Get oracle provider information
   */
  async getProvider(address: string): Promise<OracleProvider | null> {
    try {
      const result = await this.client.queryModule('oracle', 'provider', {
        address,
      });
      return this.parseProvider(result);
    } catch {
      return null;
    }
  }

  /**
   * Get all registered providers
   */
  async getAllProviders(): Promise<OracleProvider[]> {
    const result = await this.client.queryModule('oracle', 'providers', {}) as {
      providers: unknown[];
    };
    return (result.providers || []).map((p) => this.parseProvider(p));
  }

  /**
   * Check if a price is stale
   */
  async isPriceStale(denom: string): Promise<boolean> {
    const price = await this.getPrice(denom);
    if (!price) {
      return true;
    }

    const config = await this.getConfig(denom);
    if (!config) {
      return false;
    }

    const staleness = Date.now() - price.updatedAt.getTime();
    return staleness > config.stalenessThresholdSeconds * 1000;
  }

  /**
   * Calculate price change percentage from history
   */
  async getPriceChange(denom: string, periodMs: number): Promise<number | null> {
    const history = await this.getPriceHistory(denom, 100);
    if (history.prices.length < 2) {
      return null;
    }

    const now = Date.now();
    const targetTime = now - periodMs;

    const currentPrice = parseFloat(history.prices[0].amount);
    const pastPrice = history.prices.find(
      (p) => p.timestamp.getTime() <= targetTime,
    );

    if (!pastPrice) {
      return null;
    }

    const pastPriceValue = parseFloat(pastPrice.amount);
    return ((currentPrice - pastPriceValue) / pastPriceValue) * 100;
  }

  // ============================================================================
  // Helpers
  // ============================================================================

  private parsePrice(data: unknown): Price {
    const p = data as {
      denom: string;
      amount: string;
      last_updater: string;
      last_height: string;
      updated_at: string;
    };
    return {
      denom: p.denom,
      amount: p.amount,
      lastUpdater: p.last_updater,
      lastHeight: p.last_height,
      updatedAt: new Date(p.updated_at),
    };
  }

  private parseConfig(data: unknown): OracleConfig {
    const c = data as {
      denom: string;
      enabled: boolean;
      min_update_interval_seconds: number;
      max_deviation_bps: number;
      staleness_threshold_seconds: number;
    };
    return {
      denom: c.denom,
      enabled: c.enabled,
      minUpdateIntervalSeconds: c.min_update_interval_seconds,
      maxDeviationBps: c.max_deviation_bps,
      stalenessThresholdSeconds: c.staleness_threshold_seconds,
    };
  }

  private parseProvider(data: unknown): OracleProvider {
    const p = data as {
      address: string;
      is_active: boolean;
      slashed: boolean;
      slash_count: number;
      total_submissions: number;
      successful_submissions: number;
    };
    return {
      address: p.address,
      isActive: p.is_active,
      slashed: p.slashed,
      slashCount: p.slash_count,
      totalSubmissions: p.total_submissions,
      successfulSubmissions: p.successful_submissions,
    };
  }
}
