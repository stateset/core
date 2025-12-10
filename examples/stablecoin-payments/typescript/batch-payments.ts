/**
 * Batch Payments Example
 *
 * Demonstrates efficient multi-recipient payments:
 * - Payroll disbursement
 * - Revenue splitting
 * - Affiliate payouts
 * - Mass refunds
 *
 * Benefits of batch payments:
 * - Single transaction for multiple transfers
 * - Reduced gas costs
 * - Atomic execution (all succeed or all fail)
 *
 * Usage: npx ts-node batch-payments.ts
 */

import { StatesetClient } from '@stateset/core';

const RPC_ENDPOINT = process.env.STATESET_RPC || 'http://localhost:26657';
const MNEMONIC = process.env.STATESET_MNEMONIC || 'your mnemonic here';

const SSUSD_DENOM = 'ssusd';
const SSUSD_DECIMALS = 6;

function formatAmount(amount: string): string {
  return (parseInt(amount) / Math.pow(10, SSUSD_DECIMALS)).toFixed(2);
}

function toBaseUnits(amount: number): string {
  return Math.floor(amount * Math.pow(10, SSUSD_DECIMALS)).toString();
}

interface PaymentRecipient {
  name: string;
  address: string;
  amount: number;
  category: string;
}

/**
 * Example 1: Payroll Disbursement
 */
async function payrollExample(client: StatesetClient) {
  console.log('━'.repeat(60));
  console.log('  Example 1: Payroll Disbursement');
  console.log('━'.repeat(60));
  console.log();

  const employees: PaymentRecipient[] = [
    { name: 'Alice (Engineering)', address: 'stateset1alice...', amount: 5000, category: 'salary' },
    { name: 'Bob (Marketing)', address: 'stateset1bob...', amount: 4500, category: 'salary' },
    { name: 'Carol (Sales)', address: 'stateset1carol...', amount: 4000, category: 'salary' },
    { name: 'Dave (Operations)', address: 'stateset1dave...', amount: 3500, category: 'salary' },
    { name: 'Eve (Support)', address: 'stateset1eve...', amount: 3000, category: 'salary' },
  ];

  console.log('Payroll Recipients:');
  console.log('┌────────────────────────┬──────────────────┬─────────────┐');
  console.log('│ Employee               │ Address          │ Amount      │');
  console.log('├────────────────────────┼──────────────────┼─────────────┤');

  let totalPayroll = 0;
  for (const emp of employees) {
    console.log(`│ ${emp.name.padEnd(22)} │ ${emp.address.slice(0, 14).padEnd(16)} │ $${emp.amount.toFixed(2).padStart(9)} │`);
    totalPayroll += emp.amount;
  }

  console.log('├────────────────────────┼──────────────────┼─────────────┤');
  console.log(`│ ${'TOTAL'.padEnd(22)} │ ${' '.padEnd(16)} │ $${totalPayroll.toFixed(2).padStart(9)} │`);
  console.log('└────────────────────────┴──────────────────┴─────────────┘');
  console.log();

  // Execute batch payment
  console.log('Executing batch payroll transaction...');
  const result = await client.settlement.batchPayment({
    payments: employees.map(emp => ({
      recipient: emp.address,
      amount: { denom: SSUSD_DENOM, amount: toBaseUnits(emp.amount) },
    })),
    memo: `Payroll - ${new Date().toISOString().slice(0, 10)}`,
  });

  console.log(`✓ Payroll disbursed!`);
  console.log(`  TX Hash: ${result.transactionHash}`);
  console.log(`  Gas Used: ${result.gasUsed}`);
  console.log(`  Total Paid: ${totalPayroll} ssUSD`);
  console.log(`  Recipients: ${employees.length}`);
  console.log();
}

/**
 * Example 2: Revenue Split (Marketplace)
 */
async function revenueSplitExample(client: StatesetClient) {
  console.log('━'.repeat(60));
  console.log('  Example 2: Revenue Split (Marketplace Sale)');
  console.log('━'.repeat(60));
  console.log();

  const saleAmount = 1000; // Total sale amount

  // Revenue distribution
  const distribution = {
    seller: { address: 'stateset1seller...', percentage: 85 },
    platform: { address: 'stateset1platform...', percentage: 10 },
    affiliate: { address: 'stateset1affiliate...', percentage: 5 },
  };

  console.log(`Sale Amount: ${saleAmount} ssUSD`);
  console.log();
  console.log('Revenue Distribution:');
  console.log('┌─────────────────┬────────────┬─────────────┐');
  console.log('│ Recipient       │ Percentage │ Amount      │');
  console.log('├─────────────────┼────────────┼─────────────┤');

  const payments: Array<{ recipient: string; amount: { denom: string; amount: string } }> = [];

  for (const [role, config] of Object.entries(distribution)) {
    const amount = (saleAmount * config.percentage) / 100;
    console.log(`│ ${role.padEnd(15)} │ ${(config.percentage + '%').padEnd(10)} │ $${amount.toFixed(2).padStart(9)} │`);
    payments.push({
      recipient: config.address,
      amount: { denom: SSUSD_DENOM, amount: toBaseUnits(amount) },
    });
  }

  console.log('└─────────────────┴────────────┴─────────────┘');
  console.log();

  console.log('Executing revenue split...');
  const result = await client.settlement.batchPayment({
    payments,
    memo: 'Revenue split - Sale #12345',
  });

  console.log(`✓ Revenue distributed!`);
  console.log(`  TX Hash: ${result.transactionHash}`);
  console.log();
}

/**
 * Example 3: Affiliate Payouts
 */
async function affiliatePayoutsExample(client: StatesetClient) {
  console.log('━'.repeat(60));
  console.log('  Example 3: Monthly Affiliate Payouts');
  console.log('━'.repeat(60));
  console.log();

  // Affiliate earnings for the month
  const affiliates = [
    { id: 'AFF001', address: 'stateset1aff1...', sales: 50, commission: 500 },
    { id: 'AFF002', address: 'stateset1aff2...', sales: 35, commission: 350 },
    { id: 'AFF003', address: 'stateset1aff3...', sales: 28, commission: 280 },
    { id: 'AFF004', address: 'stateset1aff4...', sales: 22, commission: 220 },
    { id: 'AFF005', address: 'stateset1aff5...', sales: 15, commission: 150 },
  ];

  // Filter affiliates meeting minimum threshold
  const minPayout = 100;
  const eligibleAffiliates = affiliates.filter(a => a.commission >= minPayout);

  console.log(`Minimum Payout Threshold: ${minPayout} ssUSD`);
  console.log();
  console.log('Affiliate Earnings:');
  console.log('┌────────────┬───────────┬─────────────┬───────────┐');
  console.log('│ Affiliate  │ Sales     │ Commission  │ Status    │');
  console.log('├────────────┼───────────┼─────────────┼───────────┤');

  let totalPayout = 0;
  for (const aff of affiliates) {
    const eligible = aff.commission >= minPayout;
    const status = eligible ? '✓ Eligible' : 'Below min';
    console.log(`│ ${aff.id.padEnd(10)} │ ${aff.sales.toString().padEnd(9)} │ $${aff.commission.toFixed(2).padStart(9)} │ ${status.padEnd(9)} │`);
    if (eligible) totalPayout += aff.commission;
  }

  console.log('├────────────┼───────────┼─────────────┼───────────┤');
  console.log(`│ ${'TOTAL'.padEnd(10)} │ ${' '.padEnd(9)} │ $${totalPayout.toFixed(2).padStart(9)} │ ${' '.padEnd(9)} │`);
  console.log('└────────────┴───────────┴─────────────┴───────────┘');
  console.log();

  console.log(`Processing ${eligibleAffiliates.length} payouts...`);
  const result = await client.settlement.batchPayment({
    payments: eligibleAffiliates.map(aff => ({
      recipient: aff.address,
      amount: { denom: SSUSD_DENOM, amount: toBaseUnits(aff.commission) },
    })),
    memo: `Affiliate payouts - ${new Date().toISOString().slice(0, 7)}`,
  });

  console.log(`✓ Affiliate payouts complete!`);
  console.log(`  TX Hash: ${result.transactionHash}`);
  console.log(`  Affiliates Paid: ${eligibleAffiliates.length}`);
  console.log(`  Below Threshold: ${affiliates.length - eligibleAffiliates.length}`);
  console.log();
}

/**
 * Example 4: Mass Refunds
 */
async function massRefundExample(client: StatesetClient) {
  console.log('━'.repeat(60));
  console.log('  Example 4: Mass Refund Processing');
  console.log('━'.repeat(60));
  console.log();

  // Refund requests
  const refunds = [
    { orderId: 'ORD-1001', customer: 'stateset1cust1...', amount: 150, reason: 'Product defect' },
    { orderId: 'ORD-1045', customer: 'stateset1cust2...', amount: 89, reason: 'Wrong item' },
    { orderId: 'ORD-1089', customer: 'stateset1cust3...', amount: 225, reason: 'Cancelled' },
    { orderId: 'ORD-1102', customer: 'stateset1cust4...', amount: 75, reason: 'Never shipped' },
  ];

  console.log('Approved Refunds:');
  console.log('┌───────────────┬──────────────────┬─────────────┬────────────────┐');
  console.log('│ Order ID      │ Customer         │ Amount      │ Reason         │');
  console.log('├───────────────┼──────────────────┼─────────────┼────────────────┤');

  let totalRefunds = 0;
  for (const refund of refunds) {
    console.log(`│ ${refund.orderId.padEnd(13)} │ ${refund.customer.slice(0, 14).padEnd(16)} │ $${refund.amount.toFixed(2).padStart(9)} │ ${refund.reason.padEnd(14)} │`);
    totalRefunds += refund.amount;
  }

  console.log('├───────────────┼──────────────────┼─────────────┼────────────────┤');
  console.log(`│ ${'TOTAL'.padEnd(13)} │ ${' '.padEnd(16)} │ $${totalRefunds.toFixed(2).padStart(9)} │ ${' '.padEnd(14)} │`);
  console.log('└───────────────┴──────────────────┴─────────────┴────────────────┘');
  console.log();

  console.log('Processing refunds...');
  const result = await client.settlement.batchPayment({
    payments: refunds.map(r => ({
      recipient: r.customer,
      amount: { denom: SSUSD_DENOM, amount: toBaseUnits(r.amount) },
    })),
    memo: 'Batch refund processing',
  });

  console.log(`✓ Refunds processed!`);
  console.log(`  TX Hash: ${result.transactionHash}`);
  console.log(`  Refunds Issued: ${refunds.length}`);
  console.log(`  Total Refunded: ${totalRefunds} ssUSD`);
  console.log();
}

async function main() {
  console.log('═'.repeat(60));
  console.log('  ssUSD Batch Payments Examples');
  console.log('  Efficient Multi-Recipient Transfers');
  console.log('═'.repeat(60));
  console.log();

  // Connect to Stateset
  console.log('Connecting to Stateset network...');
  const client = await StatesetClient.connectWithMnemonic(MNEMONIC, {
    rpcEndpoint: RPC_ENDPOINT,
  });
  console.log(`Connected as: ${client.getAddress()}`);
  console.log();

  // Check balance
  const balance = await client.getBalance(SSUSD_DENOM);
  console.log(`Available Balance: ${formatAmount(balance)} ssUSD`);
  console.log();

  // Run examples
  await payrollExample(client);
  await revenueSplitExample(client);
  await affiliatePayoutsExample(client);
  await massRefundExample(client);

  // Cleanup
  client.disconnect();

  console.log('═'.repeat(60));
  console.log('  All Batch Payment Examples Completed!');
  console.log('═'.repeat(60));
}

main().catch(console.error);
