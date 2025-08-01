#!/bin/bash

# AI Agent Business Operations Tool

set -e

# Default values
CHAIN_ID="stateset-1"
NODE="http://localhost:26657"
KEYRING="--keyring-backend test"
GAS_FLAGS="--gas auto --gas-adjustment 1.5 --gas-prices 0.025stake"

# Load deployment info
if [ ! -f deployment.json ]; then
    echo "Error: deployment.json not found. Run deploy-contracts.sh first."
    exit 1
fi

AGENT_CONTRACT=$(jq -r '.agent_registry.address' deployment.json)

# Function to show usage
usage() {
    echo "AI Agent Business Operations Tool"
    echo ""
    echo "Usage: $0 <command> [arguments]"
    echo ""
    echo "Commands:"
    echo "  create-po <buyer_id> <seller_id> <items_json> <delivery_terms> <payment_terms_json>"
    echo "      Create a purchase order"
    echo ""
    echo "  update-po <po_id> <status> <agent_id> [notes]"
    echo "      Update purchase order status (submitted, accepted, rejected, etc.)"
    echo ""
    echo "  create-invoice <po_id> <seller_id> <line_items_json> <due_days> [tax_rate] [discount_rate]"
    echo "      Create an invoice for a PO"
    echo ""
    echo "  pay-invoice <invoice_id> <buyer_id> [payment_ref]"
    echo "      Pay an invoice"
    echo ""
    echo "  confirm-receipt <po_id> <buyer_id> <items_received_json> [notes]"
    echo "      Confirm receipt of goods"
    echo ""
    echo "  query-po <po_id>"
    echo "      Query purchase order details"
    echo ""
    echo "  query-invoice <invoice_id>"
    echo "      Query invoice details"
    echo ""
    echo "  account-summary <agent_id>"
    echo "      Get financial summary for an agent"
    echo ""
    echo "Examples:"
    echo "  $0 create-po agent_123 agent_456 '[{\"item_id\":\"ITEM-001\",\"description\":\"Product\",\"quantity\":\"10\",\"unit_price\":\"1000000\",\"unit\":\"piece\"}]' \"FOB Warehouse\" '{\"payment_type\":\"net\",\"net_days\":\"30\"}'"
    echo ""
    exit 1
}

# Check if command provided
if [ $# -lt 1 ]; then
    usage
fi

COMMAND=$1

case $COMMAND in
    "create-po")
        if [ $# -lt 6 ]; then
            echo "Error: create-po requires <buyer_id> <seller_id> <items_json> <delivery_terms> <payment_terms_json>"
            exit 1
        fi
        
        BUYER_ID=$2
        SELLER_ID=$3
        ITEMS=$4
        DELIVERY_TERMS=$5
        PAYMENT_TERMS=$6
        
        echo "Creating purchase order..."
        
        # Get buyer key
        BUYER_KEY=$(cat agent-*.json | jq -r --arg id "$BUYER_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$BUYER_KEY" ]; then
            echo "Error: Could not find key for buyer agent $BUYER_ID"
            exit 1
        fi
        
        CREATE_PO_MSG='{
          "create_purchase_order": {
            "buyer_agent_id": "'$BUYER_ID'",
            "seller_agent_id": "'$SELLER_ID'",
            "items": '$ITEMS',
            "delivery_terms": "'$DELIVERY_TERMS'",
            "payment_terms": '$PAYMENT_TERMS'
          }
        }'
        
        PO_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$CREATE_PO_MSG" \
            --from $BUYER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y \
            --output json)
        
        PO_ID=$(echo $PO_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="po_id") | .value' | base64 -d)
        echo "Purchase order created! PO ID: $PO_ID"
        ;;
        
    "update-po")
        if [ $# -lt 4 ]; then
            echo "Error: update-po requires <po_id> <status> <agent_id> [notes]"
            exit 1
        fi
        
        PO_ID=$2
        STATUS=$3
        AGENT_ID=$4
        NOTES=$5
        
        echo "Updating purchase order $PO_ID..."
        
        # Get agent key
        AGENT_KEY=$(cat agent-*.json | jq -r --arg id "$AGENT_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$AGENT_KEY" ]; then
            echo "Error: Could not find key for agent $AGENT_ID"
            exit 1
        fi
        
        UPDATE_MSG='{
          "update_purchase_order": {
            "po_id": "'$PO_ID'",
            "status": "'$STATUS'",
            "updater_agent_id": "'$AGENT_ID'",
            "notes": "'$NOTES'"
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$UPDATE_MSG" \
            --from $AGENT_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "create-invoice")
        if [ $# -lt 5 ]; then
            echo "Error: create-invoice requires <po_id> <seller_id> <line_items_json> <due_days> [tax_rate] [discount_rate]"
            exit 1
        fi
        
        PO_ID=$2
        SELLER_ID=$3
        LINE_ITEMS=$4
        DUE_DAYS=$5
        TAX_RATE=$6
        DISCOUNT_RATE=$7
        
        echo "Creating invoice for PO $PO_ID..."
        
        # Get seller key
        SELLER_KEY=$(cat agent-*.json | jq -r --arg id "$SELLER_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$SELLER_KEY" ]; then
            echo "Error: Could not find key for seller agent $SELLER_ID"
            exit 1
        fi
        
        # Calculate due date
        DUE_DATE=$(($(date +%s) + (DUE_DAYS * 86400)))
        
        INVOICE_MSG='{
          "create_invoice": {
            "po_id": "'$PO_ID'",
            "seller_agent_id": "'$SELLER_ID'",
            "line_items": '$LINE_ITEMS',
            "due_date": "'$DUE_DATE'"'
        
        if [ ! -z "$TAX_RATE" ]; then
            INVOICE_MSG="$INVOICE_MSG"',
            "tax_rate": "'$TAX_RATE'"'
        fi
        
        if [ ! -z "$DISCOUNT_RATE" ]; then
            INVOICE_MSG="$INVOICE_MSG"',
            "discount_rate": "'$DISCOUNT_RATE'"'
        fi
        
        INVOICE_MSG="$INVOICE_MSG"'
          }
        }'
        
        INVOICE_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$INVOICE_MSG" \
            --from $SELLER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y \
            --output json)
        
        INVOICE_ID=$(echo $INVOICE_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="invoice_id") | .value' | base64 -d)
        echo "Invoice created! Invoice ID: $INVOICE_ID"
        ;;
        
    "pay-invoice")
        if [ $# -lt 3 ]; then
            echo "Error: pay-invoice requires <invoice_id> <buyer_id> [payment_ref]"
            exit 1
        fi
        
        INVOICE_ID=$2
        BUYER_ID=$3
        PAYMENT_REF=$4
        
        echo "Paying invoice $INVOICE_ID..."
        
        # Get buyer key
        BUYER_KEY=$(cat agent-*.json | jq -r --arg id "$BUYER_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$BUYER_KEY" ]; then
            echo "Error: Could not find key for buyer agent $BUYER_ID"
            exit 1
        fi
        
        PAY_MSG='{
          "pay_invoice": {
            "invoice_id": "'$INVOICE_ID'",
            "buyer_agent_id": "'$BUYER_ID'"'
        
        if [ ! -z "$PAYMENT_REF" ]; then
            PAY_MSG="$PAY_MSG"',
            "payment_reference": "'$PAYMENT_REF'"'
        fi
        
        PAY_MSG="$PAY_MSG"'
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$PAY_MSG" \
            --from $BUYER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "confirm-receipt")
        if [ $# -lt 4 ]; then
            echo "Error: confirm-receipt requires <po_id> <buyer_id> <items_received_json> [notes]"
            exit 1
        fi
        
        PO_ID=$2
        BUYER_ID=$3
        ITEMS_RECEIVED=$4
        NOTES=$5
        
        echo "Confirming receipt for PO $PO_ID..."
        
        # Get buyer key
        BUYER_KEY=$(cat agent-*.json | jq -r --arg id "$BUYER_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$BUYER_KEY" ]; then
            echo "Error: Could not find key for buyer agent $BUYER_ID"
            exit 1
        fi
        
        CONFIRM_MSG='{
          "confirm_receipt": {
            "po_id": "'$PO_ID'",
            "buyer_agent_id": "'$BUYER_ID'",
            "items_received": '$ITEMS_RECEIVED',
            "notes": "'$NOTES'"
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$CONFIRM_MSG" \
            --from $BUYER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "query-po")
        if [ $# -lt 2 ]; then
            echo "Error: query-po requires <po_id>"
            exit 1
        fi
        
        PO_ID=$2
        QUERY_MSG='{
          "purchase_order": {
            "po_id": "'$PO_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "query-invoice")
        if [ $# -lt 2 ]; then
            echo "Error: query-invoice requires <invoice_id>"
            exit 1
        fi
        
        INVOICE_ID=$2
        QUERY_MSG='{
          "invoice": {
            "invoice_id": "'$INVOICE_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "account-summary")
        if [ $# -lt 2 ]; then
            echo "Error: account-summary requires <agent_id>"
            exit 1
        fi
        
        AGENT_ID=$2
        QUERY_MSG='{
          "account_summary": {
            "agent_id": "'$AGENT_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    *)
        echo "Error: Unknown command '$COMMAND'"
        usage
        ;;
esac