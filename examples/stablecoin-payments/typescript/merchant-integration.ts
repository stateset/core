/**
 * Merchant Integration Example
 *
 * Complete e-commerce payment integration demonstrating:
 * - Merchant registration
 * - Order creation with ssUSD payment
 * - Payment processing with escrow
 * - Order fulfillment and settlement
 * - Refund handling
 *
 * This is a reference implementation for integrating
 * ssUSD payments into your e-commerce platform.
 *
 * Usage: npx ts-node merchant-integration.ts
 */

import { StatesetClient } from '@stateset/core';
import type { MerchantConfig, EscrowSettlement } from '@stateset/core/types';

const RPC_ENDPOINT = process.env.STATESET_RPC || 'http://localhost:26657';

const SSUSD_DENOM = 'ssusd';
const SSUSD_DECIMALS = 6;

function formatAmount(amount: string | number): string {
  const num = typeof amount === 'string' ? parseInt(amount) : amount;
  return (num / Math.pow(10, SSUSD_DECIMALS)).toFixed(2);
}

function toBaseUnits(amount: number): string {
  return Math.floor(amount * Math.pow(10, SSUSD_DECIMALS)).toString();
}

// ============================================================================
// Data Types
// ============================================================================

interface OrderItem {
  productId: string;
  name: string;
  quantity: number;
  unitPrice: number;
}

interface Order {
  id: string;
  customerId: string;
  merchantId: string;
  items: OrderItem[];
  subtotal: number;
  tax: number;
  shipping: number;
  total: number;
  status: 'pending' | 'paid' | 'processing' | 'shipped' | 'delivered' | 'refunded';
  escrowId?: string;
  createdAt: Date;
}

interface PaymentResult {
  success: boolean;
  transactionHash?: string;
  escrowId?: string;
  error?: string;
}

// ============================================================================
// Merchant Service
// ============================================================================

class MerchantPaymentService {
  private client: StatesetClient;
  private merchantConfig?: MerchantConfig;

  constructor(client: StatesetClient) {
    this.client = client;
  }

  /**
   * Register as a merchant
   */
  async registerMerchant(name: string, feeRate: number, webhookUrl?: string): Promise<void> {
    console.log('Registering merchant...');
    console.log(`  Name: ${name}`);
    console.log(`  Fee Rate: ${(feeRate * 100).toFixed(2)}%`);
    if (webhookUrl) console.log(`  Webhook: ${webhookUrl}`);

    await this.client.settlement.registerMerchant(name, feeRate, webhookUrl);

    this.merchantConfig = await this.client.settlement.getMerchant(this.client.getAddress());
    console.log('  ✓ Merchant registered!');
  }

  /**
   * Get merchant configuration
   */
  async getMerchantInfo(): Promise<MerchantConfig | null> {
    if (!this.merchantConfig) {
      this.merchantConfig = await this.client.settlement.getMerchant(this.client.getAddress());
    }
    return this.merchantConfig;
  }

  /**
   * Process a payment for an order
   */
  async processPayment(
    order: Order,
    customerClient: StatesetClient,
    useEscrow: boolean = true,
  ): Promise<PaymentResult> {
    console.log(`Processing payment for order ${order.id}...`);
    console.log(`  Amount: ${order.total} ssUSD`);
    console.log(`  Escrow: ${useEscrow ? 'Yes' : 'No'}`);

    try {
      if (useEscrow) {
        // Create escrow - funds held until delivery
        const releaseTime = new Date(Date.now() + 14 * 24 * 60 * 60 * 1000); // 14 days

        const result = await customerClient.settlement.createEscrow({
          recipient: this.client.getAddress(),
          amount: {
            denom: SSUSD_DENOM,
            amount: toBaseUnits(order.total),
          },
          releaseTime,
          conditions: JSON.stringify({
            orderId: order.id,
            requiresDeliveryConfirmation: true,
          }),
        });

        console.log(`  ✓ Payment held in escrow`);
        console.log(`  Escrow ID: ${result.escrowId}`);

        return {
          success: true,
          transactionHash: result.transactionHash,
          escrowId: result.escrowId,
        };
      } else {
        // Instant transfer
        const result = await customerClient.settlement.instantTransfer({
          recipient: this.client.getAddress(),
          amount: {
            denom: SSUSD_DENOM,
            amount: toBaseUnits(order.total),
          },
          memo: `Payment for order ${order.id}`,
        });

        console.log(`  ✓ Payment received`);

        return {
          success: true,
          transactionHash: result.transactionHash,
        };
      }
    } catch (error) {
      console.log(`  ✗ Payment failed: ${error}`);
      return {
        success: false,
        error: String(error),
      };
    }
  }

  /**
   * Release escrow after delivery confirmation
   */
  async releaseEscrow(escrowId: string, customerClient: StatesetClient): Promise<boolean> {
    console.log(`Releasing escrow ${escrowId}...`);

    try {
      await customerClient.settlement.releaseEscrow(escrowId);
      console.log('  ✓ Escrow released to merchant');
      return true;
    } catch (error) {
      console.log(`  ✗ Release failed: ${error}`);
      return false;
    }
  }

  /**
   * Process a refund
   */
  async processRefund(
    order: Order,
    amount: number,
    reason: string,
    customerAddress: string,
  ): Promise<PaymentResult> {
    console.log(`Processing refund for order ${order.id}...`);
    console.log(`  Amount: ${amount} ssUSD`);
    console.log(`  Reason: ${reason}`);

    try {
      // If escrow is still active, refund from escrow
      if (order.escrowId) {
        const escrow = await this.client.settlement.getEscrow(order.escrowId);
        if (escrow && escrow.status === 'active') {
          await this.client.settlement.refundEscrow(order.escrowId);
          console.log('  ✓ Escrow refunded to customer');
          return { success: true };
        }
      }

      // Otherwise, send direct refund
      const result = await this.client.settlement.instantTransfer({
        recipient: customerAddress,
        amount: {
          denom: SSUSD_DENOM,
          amount: toBaseUnits(amount),
        },
        memo: `Refund for order ${order.id}: ${reason}`,
      });

      console.log('  ✓ Refund sent');
      return {
        success: true,
        transactionHash: result.transactionHash,
      };
    } catch (error) {
      console.log(`  ✗ Refund failed: ${error}`);
      return {
        success: false,
        error: String(error),
      };
    }
  }
}

// ============================================================================
// Demo: Complete E-Commerce Flow
// ============================================================================

async function ecommerceDemo() {
  console.log('═'.repeat(70));
  console.log('  Complete E-Commerce Payment Flow');
  console.log('═'.repeat(70));
  console.log();

  // Setup: Connect merchant and customer
  console.log('Setup: Connecting participants...');
  const merchantClient = await StatesetClient.connectWithMnemonic(
    process.env.MERCHANT_MNEMONIC || 'merchant mnemonic',
    { rpcEndpoint: RPC_ENDPOINT },
  );
  const customerClient = await StatesetClient.connectWithMnemonic(
    process.env.CUSTOMER_MNEMONIC || 'customer mnemonic',
    { rpcEndpoint: RPC_ENDPOINT },
  );

  const merchantAddress = merchantClient.getAddress();
  const customerAddress = customerClient.getAddress();

  console.log(`  Merchant: ${merchantAddress.slice(0, 20)}...`);
  console.log(`  Customer: ${customerAddress.slice(0, 20)}...`);
  console.log();

  // Initialize merchant service
  const merchantService = new MerchantPaymentService(merchantClient);

  // Step 1: Register Merchant
  console.log('━'.repeat(70));
  console.log('  Step 1: Merchant Registration');
  console.log('━'.repeat(70));
  await merchantService.registerMerchant(
    'Acme Commerce Store',
    0.025, // 2.5% fee rate
    'https://acme.store/webhooks/stateset',
  );
  console.log();

  // Step 2: Customer creates order
  console.log('━'.repeat(70));
  console.log('  Step 2: Customer Creates Order');
  console.log('━'.repeat(70));

  const order: Order = {
    id: `ORD-${Date.now()}`,
    customerId: customerAddress,
    merchantId: merchantAddress,
    items: [
      { productId: 'SKU-001', name: 'Wireless Headphones', quantity: 1, unitPrice: 79.99 },
      { productId: 'SKU-002', name: 'Phone Case', quantity: 2, unitPrice: 19.99 },
      { productId: 'SKU-003', name: 'USB-C Cable', quantity: 3, unitPrice: 9.99 },
    ],
    subtotal: 149.94,
    tax: 12.00,
    shipping: 5.99,
    total: 167.93,
    status: 'pending',
    createdAt: new Date(),
  };

  console.log('Order Created:');
  console.log('┌────────────────────────────────────────────────────────────────┐');
  console.log(`│ Order ID: ${order.id.padEnd(53)}│`);
  console.log('├────────────────────────────────────────────────────────────────┤');
  console.log('│ Items:                                                         │');
  for (const item of order.items) {
    const line = `  ${item.quantity}x ${item.name}`;
    const price = `$${(item.unitPrice * item.quantity).toFixed(2)}`;
    console.log(`│ ${line.padEnd(50)} ${price.padStart(12)} │`);
  }
  console.log('├────────────────────────────────────────────────────────────────┤');
  console.log(`│ ${'Subtotal:'.padEnd(50)} $${order.subtotal.toFixed(2).padStart(10)} │`);
  console.log(`│ ${'Tax:'.padEnd(50)} $${order.tax.toFixed(2).padStart(10)} │`);
  console.log(`│ ${'Shipping:'.padEnd(50)} $${order.shipping.toFixed(2).padStart(10)} │`);
  console.log('├────────────────────────────────────────────────────────────────┤');
  console.log(`│ ${'TOTAL:'.padEnd(50)} $${order.total.toFixed(2).padStart(10)} │`);
  console.log('└────────────────────────────────────────────────────────────────┘');
  console.log();

  // Step 3: Process Payment
  console.log('━'.repeat(70));
  console.log('  Step 3: Process Payment with Escrow');
  console.log('━'.repeat(70));

  const paymentResult = await merchantService.processPayment(order, customerClient, true);

  if (paymentResult.success) {
    order.status = 'paid';
    order.escrowId = paymentResult.escrowId;
    console.log(`  Order status: ${order.status}`);
  } else {
    console.log('  Payment failed, order cancelled');
    return;
  }
  console.log();

  // Step 4: Merchant processes order
  console.log('━'.repeat(70));
  console.log('  Step 4: Order Processing & Fulfillment');
  console.log('━'.repeat(70));

  console.log('  → Merchant receives order notification');
  order.status = 'processing';
  console.log(`  → Order status: ${order.status}`);

  console.log('  → Merchant packs and ships order');
  order.status = 'shipped';
  console.log(`  → Order status: ${order.status}`);
  console.log('  → Tracking: SHIP-12345-ACME');
  console.log();

  // Step 5: Customer confirms delivery
  console.log('━'.repeat(70));
  console.log('  Step 5: Delivery Confirmation & Settlement');
  console.log('━'.repeat(70));

  console.log('  → Customer receives package');
  console.log('  → Customer confirms delivery');
  order.status = 'delivered';

  if (order.escrowId) {
    const released = await merchantService.releaseEscrow(order.escrowId, customerClient);
    if (released) {
      console.log(`  → Final order status: ${order.status}`);
    }
  }
  console.log();

  // Step 6: Check final balances
  console.log('━'.repeat(70));
  console.log('  Step 6: Final Balance Check');
  console.log('━'.repeat(70));

  const merchantBalance = await merchantClient.getBalance(SSUSD_DENOM);
  const customerBalance = await customerClient.getBalance(SSUSD_DENOM);

  console.log(`  Merchant balance: ${formatAmount(merchantBalance)} ssUSD`);
  console.log(`  Customer balance: ${formatAmount(customerBalance)} ssUSD`);
  console.log();

  // Cleanup
  merchantClient.disconnect();
  customerClient.disconnect();

  console.log('═'.repeat(70));
  console.log('  E-Commerce Flow Completed Successfully!');
  console.log('═'.repeat(70));
}

// ============================================================================
// Demo: Refund Scenario
// ============================================================================

async function refundDemo() {
  console.log();
  console.log('═'.repeat(70));
  console.log('  Refund Scenario Demo');
  console.log('═'.repeat(70));
  console.log();

  const merchantClient = await StatesetClient.connectWithMnemonic(
    process.env.MERCHANT_MNEMONIC || 'merchant mnemonic',
    { rpcEndpoint: RPC_ENDPOINT },
  );
  const customerClient = await StatesetClient.connectWithMnemonic(
    process.env.CUSTOMER_MNEMONIC || 'customer mnemonic',
    { rpcEndpoint: RPC_ENDPOINT },
  );

  const merchantService = new MerchantPaymentService(merchantClient);

  // Simulate order that needs refund
  const order: Order = {
    id: 'ORD-REFUND-001',
    customerId: customerClient.getAddress(),
    merchantId: merchantClient.getAddress(),
    items: [{ productId: 'SKU-DEFECT', name: 'Defective Product', quantity: 1, unitPrice: 99.99 }],
    subtotal: 99.99,
    tax: 8.00,
    shipping: 0,
    total: 107.99,
    status: 'delivered',
    createdAt: new Date(),
  };

  console.log('Scenario: Customer reports defective product');
  console.log(`  Order: ${order.id}`);
  console.log(`  Original Amount: ${order.total} ssUSD`);
  console.log();

  // Full refund
  console.log('Processing full refund...');
  const refundResult = await merchantService.processRefund(
    order,
    order.total,
    'Defective product - customer return',
    order.customerId,
  );

  if (refundResult.success) {
    order.status = 'refunded';
    console.log(`  Order status: ${order.status}`);
  }

  merchantClient.disconnect();
  customerClient.disconnect();
}

// ============================================================================
// Main
// ============================================================================

async function main() {
  console.log('═'.repeat(70));
  console.log('  ssUSD Merchant Integration Examples');
  console.log('  Complete E-Commerce Payment Solutions');
  console.log('═'.repeat(70));
  console.log();

  await ecommerceDemo();
  // Uncomment to see refund flow:
  // await refundDemo();
}

main().catch(console.error);
