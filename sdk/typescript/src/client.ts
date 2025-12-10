/**
 * Main Stateset Client
 */

import {
  SigningStargateClient,
  StargateClient,
  GasPrice,
} from '@cosmjs/stargate';
import { DirectSecp256k1HdWallet } from '@cosmjs/proto-signing';
import { Tendermint37Client } from '@cosmjs/tendermint-rpc';

import { StablecoinModule } from './modules/stablecoin';
import { SettlementModule } from './modules/settlement';
import { TreasuryModule } from './modules/treasury';
import { OracleModule } from './modules/oracle';
import { ComplianceModule } from './modules/compliance';
import {
  MAINNET_RPC,
  CHAIN_ID_MAINNET,
  DEFAULT_GAS_PRICE,
} from './constants';

export interface StatesetClientOptions {
  rpcEndpoint?: string;
  restEndpoint?: string;
  chainId?: string;
  gasPrice?: string;
}

export class StatesetClient {
  private signingClient?: SigningStargateClient;
  private queryClient?: StargateClient;
  private tmClient?: Tendermint37Client;
  private address?: string;

  public readonly stablecoin: StablecoinModule;
  public readonly settlement: SettlementModule;
  public readonly treasury: TreasuryModule;
  public readonly oracle: OracleModule;
  public readonly compliance: ComplianceModule;

  private constructor(
    private readonly options: StatesetClientOptions,
  ) {
    this.stablecoin = new StablecoinModule(this);
    this.settlement = new SettlementModule(this);
    this.treasury = new TreasuryModule(this);
    this.oracle = new OracleModule(this);
    this.compliance = new ComplianceModule(this);
  }

  /**
   * Create a read-only client (no signing capabilities)
   */
  static async connect(options: StatesetClientOptions = {}): Promise<StatesetClient> {
    const client = new StatesetClient(options);
    const rpcEndpoint = options.rpcEndpoint || MAINNET_RPC;

    client.tmClient = await Tendermint37Client.connect(rpcEndpoint);
    client.queryClient = await StargateClient.create(client.tmClient);

    return client;
  }

  /**
   * Create a signing client from mnemonic
   */
  static async connectWithMnemonic(
    mnemonic: string,
    options: StatesetClientOptions = {},
  ): Promise<StatesetClient> {
    const client = new StatesetClient(options);
    const rpcEndpoint = options.rpcEndpoint || MAINNET_RPC;
    const gasPrice = GasPrice.fromString(options.gasPrice || DEFAULT_GAS_PRICE);

    const wallet = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
      prefix: 'stateset',
    });
    const [account] = await wallet.getAccounts();
    client.address = account.address;

    client.tmClient = await Tendermint37Client.connect(rpcEndpoint);
    client.queryClient = await StargateClient.create(client.tmClient);
    client.signingClient = await SigningStargateClient.connectWithSigner(
      rpcEndpoint,
      wallet,
      { gasPrice },
    );

    return client;
  }

  /**
   * Create a signing client from an existing signer
   */
  static async connectWithSigner(
    signer: DirectSecp256k1HdWallet,
    options: StatesetClientOptions = {},
  ): Promise<StatesetClient> {
    const client = new StatesetClient(options);
    const rpcEndpoint = options.rpcEndpoint || MAINNET_RPC;
    const gasPrice = GasPrice.fromString(options.gasPrice || DEFAULT_GAS_PRICE);

    const [account] = await signer.getAccounts();
    client.address = account.address;

    client.tmClient = await Tendermint37Client.connect(rpcEndpoint);
    client.queryClient = await StargateClient.create(client.tmClient);
    client.signingClient = await SigningStargateClient.connectWithSigner(
      rpcEndpoint,
      signer,
      { gasPrice },
    );

    return client;
  }

  /**
   * Get the connected wallet address
   */
  getAddress(): string {
    if (!this.address) {
      throw new Error('No wallet connected. Use connectWithMnemonic or connectWithSigner.');
    }
    return this.address;
  }

  /**
   * Get the signing client
   */
  getSigningClient(): SigningStargateClient {
    if (!this.signingClient) {
      throw new Error('No signing client available. Connect with a wallet first.');
    }
    return this.signingClient;
  }

  /**
   * Get the query client
   */
  getQueryClient(): StargateClient {
    if (!this.queryClient) {
      throw new Error('Client not connected.');
    }
    return this.queryClient;
  }

  /**
   * Get the Tendermint client for raw queries
   */
  getTmClient(): Tendermint37Client {
    if (!this.tmClient) {
      throw new Error('Client not connected.');
    }
    return this.tmClient;
  }

  /**
   * Get account balance
   */
  async getBalance(denom: string, address?: string): Promise<string> {
    const addr = address || this.address;
    if (!addr) {
      throw new Error('No address specified.');
    }
    const balance = await this.getQueryClient().getBalance(addr, denom);
    return balance.amount;
  }

  /**
   * Get all balances for an account
   */
  async getAllBalances(address?: string): Promise<readonly { denom: string; amount: string }[]> {
    const addr = address || this.address;
    if (!addr) {
      throw new Error('No address specified.');
    }
    return this.getQueryClient().getAllBalances(addr);
  }

  /**
   * Get current block height
   */
  async getHeight(): Promise<number> {
    return this.getQueryClient().getHeight();
  }

  /**
   * Get chain ID
   */
  async getChainId(): Promise<string> {
    return this.getQueryClient().getChainId();
  }

  /**
   * Disconnect the client
   */
  disconnect(): void {
    if (this.tmClient) {
      this.tmClient.disconnect();
    }
  }

  /**
   * Query a module endpoint via ABCI
   */
  async queryModule(
    module: string,
    method: string,
    params: Record<string, unknown> = {},
  ): Promise<unknown> {
    const path = `/stateset/${module}/${method}`;
    const data = new TextEncoder().encode(JSON.stringify(params));
    const response = await this.getTmClient().abciQuery({ path, data });

    if (response.code !== 0) {
      throw new Error(`Query failed: ${response.log}`);
    }

    return JSON.parse(new TextDecoder().decode(response.value));
  }

  /**
   * Send a transaction
   */
  async sendTx(
    messages: readonly { typeUrl: string; value: unknown }[],
    memo?: string,
  ): Promise<{
    transactionHash: string;
    height: number;
    gasUsed: number;
    gasWanted: number;
  }> {
    const client = this.getSigningClient();
    const address = this.getAddress();

    const result = await client.signAndBroadcast(
      address,
      messages,
      'auto',
      memo,
    );

    if (result.code !== 0) {
      throw new Error(`Transaction failed: ${result.rawLog}`);
    }

    return {
      transactionHash: result.transactionHash,
      height: result.height,
      gasUsed: result.gasUsed,
      gasWanted: result.gasWanted,
    };
  }
}
