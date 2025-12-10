/**
 * Escrow-Based Payment Example
 *
 * Demonstrates secure escrow payments for e-commerce transactions:
 * - Create an escrow for a purchase
 * - Hold funds until conditions are met
 * - Release to seller upon delivery confirmation
 * - Handle refunds if needed
 *
 * This pattern is ideal for:
 * - Marketplace transactions
 * - Freelance payments
 * - Any buyer-seller scenario requiring trust
 *
 * Usage: npx ts-node escrow-payment.ts
 */

import { StatesetClient } from '@stateset/core';
import type { EscrowSettlement } from '@stateset/core/types';

// Configuration
const RPC_ENDPOINT = process.env.STATESET_RPC || 'http://localhost:26657';
const BUYER_MNEMONIC = process.env.BUYER_MNEMONIC || 'buyer mnemonic here';
const SELLER_MNEMONIC = process.env.SELLER_MNEMONIC || 'seller mnemonic here';

const SSUSD_DENOM = 'ssusd';
const SSUSD_DECIMALS = 6;

function formatAmount(amount: string): string {
  return (parseInt(amount) / Math.pow(10, SSUSD_DECIMALS)).toFixed(2);
}

function toBaseUnits(amount: number): string {
  return Math.floor(amount * Math.pow(10, SSUSD_DECIMALS)).toString();
}

function printEscrowStatus(escrow: EscrowSettlement) {
  console.log('┌─────────────────────────────────────────────────────────┐');
  console.log('│                    ESCROW STATUS                        │');
  console.log('├─────────────────────────────────────────────────────────┤');
  console.log(`│ ID:         ${escrow.id.padEnd(44)}│`);
  console.log(`│ Status:     ${escrow.status.padEnd(44)}│`);
  console.log(`│ Amount:     ${formatAmount(escrow.amount.amount).padEnd(44)}│`);
  console.log(`│ Depositor:  ${escrow.depositor.slice(0, 20).padEnd(44)}│`);
  console.log(`│ Recipient:  ${escrow.recipient.slice(0, 20).padEnd(44)}│`);
  console.log(`│ Release:    ${escrow.releaseTime.toISOString().slice(0, 19).padEnd(44)}│`);
  console.log('└─────────────────────────────────────────────────────────┘');
}

async function demonstrateEscrowPayment() {
  console.log('═'.repeat(60));
  console.log('  ssUSD Escrow Payment Example');
  console.log('  Secure Marketplace Transaction');
  console.log('═'.repeat(60));
  console.log();

  // Connect both parties
  console.log('1. Connecting buyer and seller...');
  const buyerClient = await StatesetClient.connectWithMnemonic(BUYER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });
  const sellerClient = await StatesetClient.connectWithMnemonic(SELLER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });

  const buyerAddress = buyerClient.getAddress();
  const sellerAddress = sellerClient.getAddress();

  console.log(`   Buyer:  ${buyerAddress}`);
  console.log(`   Seller: ${sellerAddress}`);
  console.log();

  // Check initial balances
  console.log('2. Initial Balances:');
  const buyerBalance = await buyerClient.getBalance(SSUSD_DENOM);
  const sellerBalance = await sellerClient.getBalance(SSUSD_DENOM);
  console.log(`   Buyer:  ${formatAmount(buyerBalance)} ssUSD`);
  console.log(`   Seller: ${formatAmount(sellerBalance)} ssUSD`);
  console.log();

  // Create escrow
  const purchaseAmount = 500; // 500 ssUSD for the product
  const escrowDuration = 7 * 24 * 60 * 60 * 1000; // 7 days in ms
  const releaseTime = new Date(Date.now() + escrowDuration);

  console.log('3. Creating Escrow:');
  console.log(`   Purchase Amount: ${purchaseAmount} ssUSD`);
  console.log(`   Release Time:    ${releaseTime.toISOString()}`);
  console.log(`   Conditions:      Delivery confirmation required`);
  console.log();

  console.log('   Submitting escrow transaction...');
  const escrowResult = await buyerClient.settlement.createEscrow({
    recipient: sellerAddress,
    amount: {
      denom: SSUSD_DENOM,
      amount: toBaseUnits(purchaseAmount),
    },
    releaseTime,
    conditions: JSON.stringify({
      type: 'delivery_confirmation',
      orderId: 'ORD-12345',
      requiresTracking: true,
    }),
  });

  console.log(`   ✓ Escrow created!`);
  console.log(`   TX Hash:   ${escrowResult.transactionHash}`);
  console.log(`   Escrow ID: ${escrowResult.escrowId}`);
  console.log();

  // Check escrow status
  console.log('4. Escrow Status:');
  const escrow = await buyerClient.settlement.getEscrow(escrowResult.escrowId);
  if (escrow) {
    printEscrowStatus(escrow);
  }
  console.log();

  // Check balances after escrow
  console.log('5. Balances After Escrow (funds locked):');
  const buyerBalanceAfterEscrow = await buyerClient.getBalance(SSUSD_DENOM);
  const sellerBalanceAfterEscrow = await sellerClient.getBalance(SSUSD_DENOM);
  console.log(`   Buyer:  ${formatAmount(buyerBalanceAfterEscrow)} ssUSD`);
  console.log(`   Seller: ${formatAmount(sellerBalanceAfterEscrow)} ssUSD`);
  console.log(`   Locked: ${purchaseAmount} ssUSD (in escrow)`);
  console.log();

  // Simulate delivery confirmation and release
  console.log('6. Simulating Order Fulfillment...');
  console.log('   → Seller ships product');
  console.log('   → Tracking: TRACK-12345-ABC');
  console.log('   → Buyer confirms delivery');
  console.log();

  // Release escrow
  console.log('7. Releasing Escrow to Seller:');
  const releaseResult = await buyerClient.settlement.releaseEscrow(escrowResult.escrowId);
  console.log(`   ✓ Escrow released!`);
  console.log(`   TX Hash: ${releaseResult.transactionHash}`);
  console.log();

  // Final balances
  console.log('8. Final Balances:');
  const buyerFinalBalance = await buyerClient.getBalance(SSUSD_DENOM);
  const sellerFinalBalance = await sellerClient.getBalance(SSUSD_DENOM);
  console.log(`   Buyer:  ${formatAmount(buyerFinalBalance)} ssUSD`);
  console.log(`   Seller: ${formatAmount(sellerFinalBalance)} ssUSD`);
  console.log();

  // Calculate changes
  const buyerChange = parseInt(buyerFinalBalance) - parseInt(buyerBalance);
  const sellerChange = parseInt(sellerFinalBalance) - parseInt(sellerBalance);
  console.log('   Balance Changes:');
  console.log(`   Buyer:  ${buyerChange >= 0 ? '+' : ''}${formatAmount(buyerChange.toString())} ssUSD`);
  console.log(`   Seller: ${sellerChange >= 0 ? '+' : ''}${formatAmount(sellerChange.toString())} ssUSD`);

  // Cleanup
  buyerClient.disconnect();
  sellerClient.disconnect();

  console.log();
  console.log('═'.repeat(60));
  console.log('  Escrow Payment Completed Successfully!');
  console.log('═'.repeat(60));
}

/**
 * Demonstrate refund flow
 */
async function demonstrateRefundFlow() {
  console.log();
  console.log('═'.repeat(60));
  console.log('  BONUS: Escrow Refund Example');
  console.log('═'.repeat(60));
  console.log();

  const buyerClient = await StatesetClient.connectWithMnemonic(BUYER_MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });
  const sellerAddress = 'stateset1seller...';

  // Create escrow
  console.log('Creating escrow for potential refund...');
  const escrowResult = await buyerClient.settlement.createEscrow({
    recipient: sellerAddress,
    amount: {
      denom: SSUSD_DENOM,
      amount: toBaseUnits(200),
    },
    releaseTime: new Date(Date.now() + 3600000),
    conditions: 'Order fulfillment',
  });

  console.log(`Escrow ID: ${escrowResult.escrowId}`);
  console.log();

  // Simulate dispute scenario
  console.log('Scenario: Order cancelled, initiating refund...');

  // Refund escrow (seller must call this, or buyer after release time)
  const refundResult = await buyerClient.settlement.refundEscrow(escrowResult.escrowId);
  console.log(`✓ Escrow refunded!`);
  console.log(`TX Hash: ${refundResult.transactionHash}`);

  buyerClient.disconnect();
}

// Run the demonstration
async function main() {
  await demonstrateEscrowPayment();
  // Uncomment to see refund flow:
  // await demonstrateRefundFlow();
}

main().catch(console.error);
