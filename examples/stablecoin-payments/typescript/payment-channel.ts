/**
 * Payment Channel Example
 *
 * Demonstrates off-chain payment channels for high-frequency micropayments:
 * - Open a payment channel with initial deposit
 * - Send multiple off-chain payments (no gas per payment)
 * - Close channel and settle on-chain
 *
 * Use cases:
 * - Streaming payments (pay-per-second)
 * - Micropayments for content
 * - Gaming microtransactions
 * - IoT device payments
 *
 * Usage: npx ts-node payment-channel.ts
 */

import { StatesetClient } from '@stateset/core';
import type { PaymentChannel } from '@stateset/core/types';
import { createHash } from 'crypto';

const RPC_ENDPOINT = process.env.STATESET_RPC || 'http://localhost:26657';
const SENDER_MNEMONIC = process.env.SENDER_MNEMONIC || 'sender mnemonic here';
const RECIPIENT_MNEMONIC = process.env.RECIPIENT_MNEMONIC || 'recipient mnemonic here';

const SSUSD_DENOM = 'ssusd';
const SSUSD_DECIMALS = 6;

function formatAmount(amount: string): string {
  return (parseInt(amount) / Math.pow(10, SSUSD_DECIMALS)).toFixed(6);
}

function toBaseUnits(amount: number): string {
  return Math.floor(amount * Math.pow(10, SSUSD_DECIMALS)).toString();
}

/**
 * Simple signing utility for payment channel updates
 * In production, use proper cryptographic signing
 */
function signChannelUpdate(channelId: string, amount: string, nonce: number, privateKey: string): string {
  const message = `${channelId}:${amount}:${nonce}`;
  return createHash('sha256').update(message + privateKey).digest('hex');
}

/**
 * Represents off-chain state for payment channel
 */
interface ChannelState {
  channelId: string;
  totalSent: number;
  nonce: number;
  signatures: string[];
}

class PaymentChannelManager {
  private channels: Map<string, ChannelState> = new Map();

  constructor(
    private senderClient: StatesetClient,
    private recipientAddress: string,
  ) {}

  /**
   * Open a new payment channel
   */
  async openChannel(depositAmount: number, expiresInHours: number): Promise<string> {
    const expiresAt = new Date(Date.now() + expiresInHours * 60 * 60 * 1000);

    console.log('Opening payment channel...');
    console.log(`  Deposit: ${depositAmount} ssUSD`);
    console.log(`  Expires: ${expiresAt.toISOString()}`);

    const result = await this.senderClient.settlement.openChannel(
      this.recipientAddress,
      { denom: SSUSD_DENOM, amount: toBaseUnits(depositAmount) },
      expiresAt,
    );

    const channelId = result.channelId;
    console.log(`  Channel ID: ${channelId}`);
    console.log(`  TX Hash: ${result.transactionHash}`);

    // Initialize local state
    this.channels.set(channelId, {
      channelId,
      totalSent: 0,
      nonce: 0,
      signatures: [],
    });

    return channelId;
  }

  /**
   * Send an off-chain payment through the channel
   * This doesn't hit the blockchain - just updates local state
   */
  async sendPayment(channelId: string, amount: number): Promise<ChannelState> {
    const state = this.channels.get(channelId);
    if (!state) {
      throw new Error(`Channel ${channelId} not found`);
    }

    // Update state
    state.totalSent += amount;
    state.nonce += 1;

    // Create signature for this state update
    const signature = signChannelUpdate(
      channelId,
      toBaseUnits(state.totalSent),
      state.nonce,
      'sender-private-key', // In production, use actual private key
    );
    state.signatures.push(signature);

    this.channels.set(channelId, state);

    return state;
  }

  /**
   * Get current channel state
   */
  getChannelState(channelId: string): ChannelState | undefined {
    return this.channels.get(channelId);
  }

  /**
   * Close channel and settle on-chain
   */
  async closeChannel(channelId: string): Promise<void> {
    const state = this.channels.get(channelId);
    if (!state) {
      throw new Error(`Channel ${channelId} not found`);
    }

    console.log('Closing payment channel...');
    console.log(`  Final Amount: ${state.totalSent} ssUSD`);
    console.log(`  Total Payments: ${state.nonce}`);

    const latestSignature = state.signatures[state.signatures.length - 1] || '';

    const result = await this.senderClient.settlement.closeChannel(
      channelId,
      { denom: SSUSD_DENOM, amount: toBaseUnits(state.totalSent) },
      latestSignature,
    );

    console.log(`  TX Hash: ${result.transactionHash}`);

    // Clear local state
    this.channels.delete(channelId);
  }
}

/**
 * Demo: Streaming Payment
 * Pay-per-second for content consumption
 */
async function streamingPaymentDemo() {
  console.log('═'.repeat(60));
  console.log('  Demo 1: Streaming Payment (Pay-per-second)');
  console.log('═'.repeat(60));
  console.log();

  const senderClient = await StatesetClient.connectWithMnemonic(SENDER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });

  const recipientAddress = 'stateset1content-creator...';
  const manager = new PaymentChannelManager(senderClient, recipientAddress);

  // Open channel with 10 ssUSD deposit
  const channelId = await manager.openChannel(10, 24);
  console.log();

  // Simulate streaming - pay 0.001 ssUSD per second
  const ratePerSecond = 0.001;
  const durationSeconds = 10;

  console.log('Simulating content streaming...');
  console.log(`  Rate: ${ratePerSecond} ssUSD/second`);
  console.log(`  Duration: ${durationSeconds} seconds`);
  console.log();

  console.log('Streaming progress:');
  for (let i = 1; i <= durationSeconds; i++) {
    const state = await manager.sendPayment(channelId, ratePerSecond);
    const bar = '█'.repeat(i) + '░'.repeat(durationSeconds - i);
    process.stdout.write(`\r  [${bar}] ${i}s - Total: ${state.totalSent.toFixed(4)} ssUSD`);

    // In real scenario, wait for actual time
    // await new Promise(r => setTimeout(r, 1000));
  }
  console.log('\n');

  // Close channel and settle
  await manager.closeChannel(channelId);
  console.log();

  senderClient.disconnect();
}

/**
 * Demo: Micropayments for API usage
 */
async function apiMicropaymentDemo() {
  console.log('═'.repeat(60));
  console.log('  Demo 2: API Micropayments');
  console.log('═'.repeat(60));
  console.log();

  const senderClient = await StatesetClient.connectWithMnemonic(SENDER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });

  const apiProviderAddress = 'stateset1api-provider...';
  const manager = new PaymentChannelManager(senderClient, apiProviderAddress);

  // Open channel
  const channelId = await manager.openChannel(50, 168); // 1 week
  console.log();

  // API pricing
  const pricing = {
    basic: 0.001,    // $0.001 per basic request
    premium: 0.01,   // $0.01 per premium request
    compute: 0.05,   // $0.05 per compute unit
  };

  // Simulate API usage
  const usage = [
    { type: 'basic', count: 100 },
    { type: 'premium', count: 25 },
    { type: 'compute', count: 5 },
    { type: 'basic', count: 50 },
  ];

  console.log('Simulating API usage...');
  console.log('┌──────────────┬───────────┬───────────────┬───────────────┐');
  console.log('│ Request Type │ Count     │ Cost Each     │ Total         │');
  console.log('├──────────────┼───────────┼───────────────┼───────────────┤');

  for (const call of usage) {
    const cost = pricing[call.type as keyof typeof pricing] * call.count;
    await manager.sendPayment(channelId, cost);
    const state = manager.getChannelState(channelId)!;
    console.log(`│ ${call.type.padEnd(12)} │ ${call.count.toString().padEnd(9)} │ $${pricing[call.type as keyof typeof pricing].toFixed(3).padStart(11)} │ $${cost.toFixed(4).padStart(11)} │`);
  }

  console.log('└──────────────┴───────────┴───────────────┴───────────────┘');
  console.log();

  const finalState = manager.getChannelState(channelId)!;
  console.log(`Total API Cost: ${finalState.totalSent.toFixed(4)} ssUSD`);
  console.log(`Total Requests: ${finalState.nonce}`);
  console.log();

  // Close channel
  await manager.closeChannel(channelId);
  console.log();

  senderClient.disconnect();
}

/**
 * Demo: Gaming Micropayments
 */
async function gamingMicropaymentDemo() {
  console.log('═'.repeat(60));
  console.log('  Demo 3: Gaming Micropayments');
  console.log('═'.repeat(60));
  console.log();

  const senderClient = await StatesetClient.connectWithMnemonic(SENDER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });

  const gameServerAddress = 'stateset1game-server...';
  const manager = new PaymentChannelManager(senderClient, gameServerAddress);

  // Open channel with gaming budget
  const channelId = await manager.openChannel(100, 720); // 30 days
  console.log();

  // Game events that cost money
  const gameEvents = [
    { event: 'Extra Life', cost: 0.50 },
    { event: 'Power-up Pack', cost: 1.99 },
    { event: 'Skip Level', cost: 0.25 },
    { event: 'Cosmetic Item', cost: 2.50 },
    { event: 'Extra Life', cost: 0.50 },
    { event: 'Tournament Entry', cost: 5.00 },
    { event: 'Power-up Pack', cost: 1.99 },
  ];

  console.log('Gaming Session:');
  console.log('┌──────────────────────┬─────────────┬─────────────────┐');
  console.log('│ Event                │ Cost        │ Running Total   │');
  console.log('├──────────────────────┼─────────────┼─────────────────┤');

  for (const ge of gameEvents) {
    const state = await manager.sendPayment(channelId, ge.cost);
    console.log(`│ ${ge.event.padEnd(20)} │ $${ge.cost.toFixed(2).padStart(9)} │ $${state.totalSent.toFixed(2).padStart(13)} │`);
  }

  console.log('└──────────────────────┴─────────────┴─────────────────┘');
  console.log();

  const finalState = manager.getChannelState(channelId)!;
  console.log(`Session Summary:`);
  console.log(`  Total Spent: ${finalState.totalSent.toFixed(2)} ssUSD`);
  console.log(`  Transactions: ${finalState.nonce}`);
  console.log(`  Avg per Transaction: ${(finalState.totalSent / finalState.nonce).toFixed(2)} ssUSD`);
  console.log();

  // End session - close channel
  console.log('Ending gaming session...');
  await manager.closeChannel(channelId);
  console.log();

  senderClient.disconnect();
}

async function main() {
  console.log('═'.repeat(60));
  console.log('  ssUSD Payment Channels');
  console.log('  High-Frequency Micropayments');
  console.log('═'.repeat(60));
  console.log();

  console.log('Payment channels enable instant, gas-free micropayments');
  console.log('by settling only the final balance on-chain.');
  console.log();

  // Run all demos
  await streamingPaymentDemo();
  await apiMicropaymentDemo();
  await gamingMicropaymentDemo();

  console.log('═'.repeat(60));
  console.log('  Payment Channel Examples Completed!');
  console.log('═'.repeat(60));
}

main().catch(console.error);
