/**
 * Basic ssUSD Payment Example
 *
 * Demonstrates the simplest form of stablecoin payment:
 * - Connect to Stateset network
 * - Check balances
 * - Send an instant ssUSD transfer
 * - Verify the transaction
 *
 * Usage: npx ts-node basic-payment.ts
 */

import { StatesetClient } from '@stateset/core';

// Configuration
const RPC_ENDPOINT = process.env.STATESET_RPC || 'http://localhost:26657';
const MNEMONIC = process.env.STATESET_MNEMONIC || 'your mnemonic here';

// Denomination constants
const SSUSD_DENOM = 'ssusd';
const SSUSD_DECIMALS = 6;

/**
 * Format amount from base units to human-readable
 */
function formatAmount(amount: string, decimals: number = SSUSD_DECIMALS): string {
  const value = parseInt(amount) / Math.pow(10, decimals);
  return value.toFixed(2);
}

/**
 * Convert human-readable amount to base units
 */
function toBaseUnits(amount: number, decimals: number = SSUSD_DECIMALS): string {
  return Math.floor(amount * Math.pow(10, decimals)).toString();
}

async function main() {
  console.log('='.repeat(60));
  console.log('  ssUSD Basic Payment Example');
  console.log('  Stateset Commerce Network');
  console.log('='.repeat(60));
  console.log();

  // Step 1: Connect to Stateset
  console.log('1. Connecting to Stateset network...');
  const client = await StatesetClient.connectWithMnemonic(MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });
  const senderAddress = client.getAddress();
  console.log(`   Connected as: ${senderAddress}`);
  console.log();

  // Step 2: Check sender's balance
  console.log('2. Checking sender balance...');
  const senderBalance = await client.getBalance(SSUSD_DENOM);
  console.log(`   ssUSD Balance: ${formatAmount(senderBalance)} ssUSD`);

  if (parseInt(senderBalance) === 0) {
    console.log('\n   ⚠️  No ssUSD balance. Please fund your account first.');
    console.log('   See: scripts/test-ssusd-transfers.sh');
    client.disconnect();
    return;
  }
  console.log();

  // Step 3: Define the payment
  const recipientAddress = 'stateset1recipient...'; // Replace with actual address
  const paymentAmount = 100; // 100 ssUSD

  console.log('3. Payment Details:');
  console.log(`   From:   ${senderAddress}`);
  console.log(`   To:     ${recipientAddress}`);
  console.log(`   Amount: ${paymentAmount} ssUSD`);
  console.log();

  // Step 4: Execute the transfer
  console.log('4. Executing instant transfer...');
  try {
    const result = await client.settlement.instantTransfer({
      recipient: recipientAddress,
      amount: {
        denom: SSUSD_DENOM,
        amount: toBaseUnits(paymentAmount),
      },
      memo: 'Payment for services',
    });

    console.log(`   ✓ Transaction successful!`);
    console.log(`   TX Hash: ${result.transactionHash}`);
    console.log(`   Block:   ${result.height}`);
    console.log(`   Gas:     ${result.gasUsed}/${result.gasWanted}`);
  } catch (error) {
    console.log(`   ✗ Transaction failed: ${error}`);
  }
  console.log();

  // Step 5: Verify new balances
  console.log('5. Verifying balances after transfer...');
  const newSenderBalance = await client.getBalance(SSUSD_DENOM);
  console.log(`   Sender new balance: ${formatAmount(newSenderBalance)} ssUSD`);

  // Cleanup
  client.disconnect();
  console.log();
  console.log('='.repeat(60));
  console.log('  Payment completed successfully!');
  console.log('='.repeat(60));
}

main().catch(console.error);
