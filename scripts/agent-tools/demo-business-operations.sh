#!/bin/bash

# AI Agent Business Operations Demo Script

set -e

echo "========================================="
echo "AI Agent Business Operations Demo"
echo "========================================="
echo ""
echo "This demo will showcase a complete business cycle:"
echo "1. Create manufacturer and retailer agents"
echo "2. Retailer creates purchase order for goods"
echo "3. Manufacturer accepts PO and fulfills order"
echo "4. Manufacturer creates invoice"
echo "5. Retailer confirms receipt of goods"
echo "6. Retailer pays invoice"
echo "7. Financial reconciliation"
echo ""
read -p "Press Enter to start..."

# Create business agents
echo ""
echo "Creating Business Agents..."
echo ""

# Manufacturer Agent
./create-agent.sh "TechManufacturer" "AI agent representing a technology equipment manufacturer" "manufacturing,wholesale,b2b-sales" "10000000000"
sleep 3

# Retailer Agent  
./create-agent.sh "SmartRetailer" "AI agent representing a retail electronics store" "retail,purchasing,inventory-management" "15000000000"
sleep 3

# Logistics Agent (for future use)
./create-agent.sh "SwiftLogistics" "AI agent providing shipping and logistics services" "shipping,tracking,delivery" "5000000000"
sleep 3

echo ""
echo "========================================="
echo "Business Agents Created!"
echo "========================================="
echo ""

# Get agent IDs
MANUFACTURER=$(cat agent-agent-TechManufacturer.json | jq -r '.agent_id')
RETAILER=$(cat agent-agent-SmartRetailer.json | jq -r '.agent_id')
LOGISTICS=$(cat agent-agent-SwiftLogistics.json | jq -r '.agent_id')

echo "TechManufacturer Agent ID: $MANUFACTURER"
echo "SmartRetailer Agent ID: $RETAILER"
echo "SwiftLogistics Agent ID: $LOGISTICS"
echo ""

# Show initial balances
echo "Initial Agent Balances:"
echo "----------------------"
./agent-interact.sh query-balance $MANUFACTURER | jq -r '"TechManufacturer: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $RETAILER | jq -r '"SmartRetailer: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $LOGISTICS | jq -r '"SwiftLogistics: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
echo ""

read -p "Press Enter to create purchase order..."

# Create Purchase Order
echo ""
echo "========================================="
echo "Step 1: Create Purchase Order"
echo "========================================="
echo ""
echo "SmartRetailer creating PO for electronic goods from TechManufacturer..."

# Get retailer key
RETAILER_KEY=$(cat agent-agent-SmartRetailer.json | jq -r '.key_name')
AGENT_CONTRACT=$(jq -r '.agent_registry.address' deployment.json)

# Create PO with items
PO_MSG='{
  "create_purchase_order": {
    "buyer_agent_id": "'$RETAILER'",
    "seller_agent_id": "'$MANUFACTURER'",
    "items": [
      {
        "item_id": "LAPTOP-001",
        "description": "High-performance laptop with AI chip",
        "quantity": "50",
        "unit_price": "1500000000",
        "unit": "piece"
      },
      {
        "item_id": "TABLET-002",
        "description": "Smart tablet with neural processor",
        "quantity": "100",
        "unit_price": "800000000",
        "unit": "piece"
      },
      {
        "item_id": "PHONE-003",
        "description": "AI-powered smartphone",
        "quantity": "200",
        "unit_price": "1000000000",
        "unit": "piece"
      }
    ],
    "delivery_terms": "FOB Warehouse - Delivery within 14 days",
    "payment_terms": {
      "payment_type": "net",
      "net_days": "30"
    },
    "metadata": "Q1 2024 Electronics Order"
  }
}'

PO_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$PO_MSG" \
    --from $RETAILER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y \
    --output json)

PO_ID=$(echo $PO_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="po_id") | .value' | base64 -d)
echo "Purchase Order created! PO ID: $PO_ID"
echo ""

# Query PO details
echo "Purchase Order Details:"
QUERY_PO='{
  "purchase_order": {
    "po_id": "'$PO_ID'"
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_PO" --output json | jq '.data.purchase_order | {
  po_id: .po_id,
  total_amount: (.total_amount | tonumber / 1000000),
  items_count: (.items | length),
  status: .status,
  payment_terms: .payment_terms
}'

read -p "Press Enter for manufacturer to accept PO..."

# Manufacturer accepts PO
echo ""
echo "========================================="
echo "Step 2: Manufacturer Accepts PO"
echo "========================================="
echo ""

MANUFACTURER_KEY=$(cat agent-agent-TechManufacturer.json | jq -r '.key_name')

UPDATE_PO_MSG='{
  "update_purchase_order": {
    "po_id": "'$PO_ID'",
    "status": "accepted",
    "updater_agent_id": "'$MANUFACTURER'",
    "notes": "Order confirmed. Production will begin immediately."
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$UPDATE_PO_MSG" \
    --from $MANUFACTURER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo "PO accepted by manufacturer!"
sleep 3

# Update to in progress
UPDATE_PO_MSG='{
  "update_purchase_order": {
    "po_id": "'$PO_ID'",
    "status": "in_progress",
    "updater_agent_id": "'$MANUFACTURER'",
    "notes": "Manufacturing in progress. 50% complete."
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$UPDATE_PO_MSG" \
    --from $MANUFACTURER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo "Manufacturing in progress..."
sleep 3

read -p "Press Enter to mark items as delivered and create invoice..."

# Mark as delivered
UPDATE_PO_MSG='{
  "update_purchase_order": {
    "po_id": "'$PO_ID'",
    "status": "delivered",
    "updater_agent_id": "'$MANUFACTURER'",
    "notes": "All items shipped via SwiftLogistics. Tracking: SL-2024-001"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$UPDATE_PO_MSG" \
    --from $MANUFACTURER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo "Items marked as delivered!"
sleep 3

# Create Invoice
echo ""
echo "========================================="
echo "Step 3: Create Invoice"
echo "========================================="
echo ""
echo "Manufacturer creating invoice for delivered goods..."

# Calculate due date (30 days from now)
DUE_DATE=$(($(date +%s) + 2592000))

INVOICE_MSG='{
  "create_invoice": {
    "po_id": "'$PO_ID'",
    "seller_agent_id": "'$MANUFACTURER'",
    "line_items": [
      {
        "description": "High-performance laptop with AI chip (50 units @ 1500 aiUSD)",
        "quantity": "50",
        "unit_price": "1500000000",
        "po_item_id": "LAPTOP-001"
      },
      {
        "description": "Smart tablet with neural processor (100 units @ 800 aiUSD)",
        "quantity": "100",
        "unit_price": "800000000",
        "po_item_id": "TABLET-002"
      },
      {
        "description": "AI-powered smartphone (200 units @ 1000 aiUSD)",
        "quantity": "200",
        "unit_price": "1000000000",
        "po_item_id": "PHONE-003"
      }
    ],
    "tax_rate": "1000",
    "discount_rate": "200",
    "due_date": "'$DUE_DATE'",
    "metadata": "Invoice for PO '$PO_ID'"
  }
}'

INVOICE_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$INVOICE_MSG" \
    --from $MANUFACTURER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y \
    --output json)

INVOICE_ID=$(echo $INVOICE_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="invoice_id") | .value' | base64 -d)
echo "Invoice created! Invoice ID: $INVOICE_ID"
echo ""

# Query invoice details
echo "Invoice Details:"
QUERY_INVOICE='{
  "invoice": {
    "invoice_id": "'$INVOICE_ID'"
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_INVOICE" --output json | jq '.data.invoice | {
  invoice_id: .invoice_id,
  subtotal: (.subtotal | tonumber / 1000000),
  tax_amount: (.tax_amount | tonumber / 1000000),
  discount_amount: (.discount_amount | tonumber / 1000000),
  total_amount: (.total_amount | tonumber / 1000000),
  due_date: .due_date,
  paid: .paid
}'

read -p "Press Enter for retailer to confirm receipt..."

# Confirm Receipt
echo ""
echo "========================================="
echo "Step 4: Confirm Receipt"
echo "========================================="
echo ""
echo "Retailer confirming receipt of goods..."

CONFIRM_MSG='{
  "confirm_receipt": {
    "po_id": "'$PO_ID'",
    "buyer_agent_id": "'$RETAILER'",
    "items_received": [
      {
        "po_item_id": "LAPTOP-001",
        "quantity_received": "50",
        "condition": "good",
        "notes": "All units received in perfect condition"
      },
      {
        "po_item_id": "TABLET-002",
        "quantity_received": "100",
        "condition": "good",
        "notes": "All units received as ordered"
      },
      {
        "po_item_id": "PHONE-003",
        "quantity_received": "198",
        "condition": "good",
        "notes": "198 units in good condition, 2 units damaged in transit"
      }
    ],
    "notes": "Overall delivery satisfactory. Will coordinate with logistics for damaged units."
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$CONFIRM_MSG" \
    --from $RETAILER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo "Receipt confirmed!"
sleep 3

read -p "Press Enter for retailer to pay invoice..."

# Pay Invoice
echo ""
echo "========================================="
echo "Step 5: Pay Invoice"
echo "========================================="
echo ""
echo "Retailer paying invoice..."

PAY_MSG='{
  "pay_invoice": {
    "invoice_id": "'$INVOICE_ID'",
    "buyer_agent_id": "'$RETAILER'",
    "payment_reference": "PAYMENT-2024-Q1-001"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$PAY_MSG" \
    --from $RETAILER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo "Invoice paid!"
sleep 5

# Show updated balances
echo ""
echo "Updated Agent Balances:"
echo "----------------------"
./agent-interact.sh query-balance $MANUFACTURER | jq -r '"TechManufacturer: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $RETAILER | jq -r '"SmartRetailer: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
echo ""

read -p "Press Enter to view financial summary..."

# Financial Summary
echo ""
echo "========================================="
echo "Step 6: Financial Summary"
echo "========================================="
echo ""

# Get account summaries
echo "TechManufacturer Account Summary:"
QUERY_SUMMARY='{
  "account_summary": {
    "agent_id": "'$MANUFACTURER'"
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_SUMMARY" --output json | jq '.data | {
  total_sales: (.total_sales | tonumber / 1000000),
  total_purchases: (.total_purchases | tonumber / 1000000),
  outstanding_receivables: (.outstanding_receivables | tonumber / 1000000),
  outstanding_payables: (.outstanding_payables | tonumber / 1000000),
  completed_orders: .completed_orders,
  pending_orders: .pending_orders
}'

echo ""
echo "SmartRetailer Account Summary:"
QUERY_SUMMARY='{
  "account_summary": {
    "agent_id": "'$RETAILER'"
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_SUMMARY" --output json | jq '.data | {
  total_sales: (.total_sales | tonumber / 1000000),
  total_purchases: (.total_purchases | tonumber / 1000000),
  outstanding_receivables: (.outstanding_receivables | tonumber / 1000000),
  outstanding_payables: (.outstanding_payables | tonumber / 1000000),
  completed_orders: .completed_orders,
  pending_orders: .pending_orders
}'

# Reconcile accounts
echo ""
echo "Reconciling accounts for both agents..."

CURRENT_TIME=$(date +%s)
PERIOD_START=$((CURRENT_TIME - 86400)) # 24 hours ago

RECONCILE_MSG='{
  "reconcile_accounts": {
    "agent_id": "'$MANUFACTURER'",
    "period_start": "'$PERIOD_START'",
    "period_end": "'$CURRENT_TIME'"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$RECONCILE_MSG" \
    --from $MANUFACTURER_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

echo ""
echo "========================================="
echo "Business Cycle Complete!"
echo "========================================="
echo ""
echo "Summary of completed business operations:"
echo "✓ Purchase Order created and accepted"
echo "✓ Goods manufactured and delivered"
echo "✓ Invoice generated with tax and discount"
echo "✓ Receipt confirmed by buyer"
echo "✓ Payment processed successfully"
echo "✓ Financial records updated"
echo ""
echo "The AI agents have successfully completed a full B2B transaction cycle!"
echo ""
echo "Additional capabilities available:"
echo "- Refund processing"
echo "- Multi-agent supply chains"
echo "- Automated reordering"
echo "- Credit terms and financing"
echo "- Integration with logistics agents"