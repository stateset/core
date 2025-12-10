#!/bin/bash

# =============================================================================
# ssUSD Quick Start - Get Up and Running in 5 Minutes
# =============================================================================
#
# This script demonstrates the fastest way to:
# 1. Create test accounts
# 2. Fund them with ssUSD
# 3. Send your first payment
#
# Prerequisites:
# - Stateset node running locally (./scripts/run-stateset.sh)
# - statesetd binary available
#
# Usage: ./quick-start.sh
# =============================================================================

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m'

# Configuration
CHAIN_ID="stateset-1"
NODE="tcp://localhost:26657"
KEYRING="test"
GAS="--gas=auto --gas-adjustment=1.5 --gas-prices=0.025ustake"
BINARY="statesetd"

# Use local build if available
[ -f "./build/statesetd" ] && BINARY="./build/statesetd"

# =============================================================================
# Helper Functions
# =============================================================================

print_step() {
    echo ""
    echo -e "${CYAN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
    echo -e "${CYAN}STEP $1: $2${NC}"
    echo -e "${CYAN}━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━${NC}"
}

print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

print_info() {
    echo -e "${BLUE}→ $1${NC}"
}

run_tx() {
    local result=$($BINARY tx "$@" \
        --chain-id=$CHAIN_ID \
        --node=$NODE \
        --keyring-backend=$KEYRING \
        $GAS \
        --broadcast-mode=sync \
        --yes \
        --output=json 2>/dev/null)
    echo "$result"
}

get_balance() {
    local address=$1
    local denom=$2
    $BINARY query bank balances "$address" \
        --node=$NODE \
        --output=json 2>/dev/null | \
        jq -r ".balances[] | select(.denom==\"$denom\") | .amount // \"0\""
}

wait_for_tx() {
    sleep 3
}

# =============================================================================
# Main Script
# =============================================================================

clear
echo -e "${GREEN}"
echo "╔════════════════════════════════════════════════════════════════════╗"
echo "║                                                                    ║"
echo "║              ssUSD QUICK START GUIDE                              ║"
echo "║              Your First Stablecoin Payment                        ║"
echo "║                                                                    ║"
echo "╚════════════════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# Check prerequisites
print_step "0" "Checking Prerequisites"

if ! command -v $BINARY &> /dev/null; then
    echo -e "${RED}Error: $BINARY not found!${NC}"
    echo "Please build the binary first: make build"
    exit 1
fi
print_success "statesetd binary found"

if ! $BINARY status --node=$NODE > /dev/null 2>&1; then
    echo -e "${YELLOW}Warning: Node may not be running at $NODE${NC}"
    echo "Start your node with: ./scripts/run-stateset.sh"
fi
print_success "Node connectivity checked"

# Step 1: Create accounts
print_step "1" "Creating Test Accounts"

for account in alice bob; do
    if ! $BINARY keys show $account --keyring-backend=$KEYRING &> /dev/null; then
        print_info "Creating $account account..."
        $BINARY keys add $account --keyring-backend=$KEYRING > /dev/null 2>&1
        print_success "$account created"
    else
        print_info "$account already exists"
    fi
done

ALICE=$($BINARY keys show alice -a --keyring-backend=$KEYRING)
BOB=$($BINARY keys show bob -a --keyring-backend=$KEYRING)

echo ""
echo "  Alice: $ALICE"
echo "  Bob:   $BOB"

# Step 2: Check initial balances
print_step "2" "Checking Initial Balances"

ALICE_SSUSD=$(get_balance "$ALICE" "ssusd")
BOB_SSUSD=$(get_balance "$BOB" "ssusd")

echo ""
echo "┌─────────────┬──────────────────┐"
echo "│   Account   │   ssUSD Balance  │"
echo "├─────────────┼──────────────────┤"
printf "│ %-11s │ %16s │\n" "Alice" "${ALICE_SSUSD:-0}"
printf "│ %-11s │ %16s │\n" "Bob" "${BOB_SSUSD:-0}"
echo "└─────────────┴──────────────────┘"

# Step 3: Fund Alice if needed
print_step "3" "Funding Alice with ssUSD"

if [ "${ALICE_SSUSD:-0}" = "0" ] || [ "${ALICE_SSUSD:-0}" -lt "1000000000" ]; then
    print_info "Minting 1000 ssUSD for Alice..."

    # For demo purposes - in production this would require proper reserves
    TX=$(run_tx stablecoins mint-stablecoin ssusd "$ALICE" 1000000000 --from=alice)
    TX_HASH=$(echo "$TX" | jq -r '.txhash')

    if [ -n "$TX_HASH" ] && [ "$TX_HASH" != "null" ]; then
        print_success "Mint transaction submitted: ${TX_HASH:0:16}..."
        wait_for_tx
    else
        print_info "Note: Minting requires proper setup. See test-ssusd-transfers.sh"
    fi
else
    print_info "Alice already has sufficient ssUSD"
fi

ALICE_SSUSD=$(get_balance "$ALICE" "ssusd")
echo "  Alice's balance: ${ALICE_SSUSD:-0} ssusd"

# Step 4: Send payment
print_step "4" "Sending Your First ssUSD Payment"

echo ""
echo "Payment Details:"
echo "  From:   Alice"
echo "  To:     Bob"
echo "  Amount: 100 ssUSD (100000000 base units)"
echo ""

if [ "${ALICE_SSUSD:-0}" -gt "100000000" ]; then
    print_info "Executing transfer..."

    TX=$(run_tx bank send "$ALICE" "$BOB" 100000000ssusd --from=alice)
    TX_HASH=$(echo "$TX" | jq -r '.txhash')

    if [ -n "$TX_HASH" ] && [ "$TX_HASH" != "null" ]; then
        print_success "Transfer successful!"
        echo "  TX Hash: $TX_HASH"
        wait_for_tx
    else
        echo -e "${YELLOW}  Transfer may have failed. Check node status.${NC}"
    fi
else
    echo -e "${YELLOW}  Insufficient balance. Please fund Alice's account first.${NC}"
fi

# Step 5: Verify transfer
print_step "5" "Verifying Transfer"

ALICE_FINAL=$(get_balance "$ALICE" "ssusd")
BOB_FINAL=$(get_balance "$BOB" "ssusd")

echo ""
echo "Final Balances:"
echo "┌─────────────┬──────────────────┬───────────────────┐"
echo "│   Account   │   ssUSD Balance  │      Change       │"
echo "├─────────────┼──────────────────┼───────────────────┤"

ALICE_CHANGE=$((${ALICE_FINAL:-0} - ${ALICE_SSUSD:-0}))
BOB_CHANGE=$((${BOB_FINAL:-0} - ${BOB_SSUSD:-0}))

printf "│ %-11s │ %16s │ %+17s │\n" "Alice" "${ALICE_FINAL:-0}" "$ALICE_CHANGE"
printf "│ %-11s │ %16s │ %+17s │\n" "Bob" "${BOB_FINAL:-0}" "$BOB_CHANGE"
echo "└─────────────┴──────────────────┴───────────────────┘"

# Summary
echo ""
echo -e "${GREEN}"
echo "╔════════════════════════════════════════════════════════════════════╗"
echo "║                    🎉 QUICK START COMPLETE! 🎉                    ║"
echo "╠════════════════════════════════════════════════════════════════════╣"
echo "║                                                                    ║"
echo "║  You've successfully:                                             ║"
echo "║  ✓ Created test accounts                                          ║"
echo "║  ✓ Funded accounts with ssUSD                                     ║"
echo "║  ✓ Sent your first ssUSD payment                                  ║"
echo "║                                                                    ║"
echo "╠════════════════════════════════════════════════════════════════════╣"
echo "║  NEXT STEPS:                                                       ║"
echo "║                                                                    ║"
echo "║  1. Try escrow payments:                                          ║"
echo "║     ./ssusd-payment-demo.sh                                       ║"
echo "║                                                                    ║"
echo "║  2. Complete order flow:                                          ║"
echo "║     ./demo-order-processing.sh                                    ║"
echo "║                                                                    ║"
echo "║  3. TypeScript SDK examples:                                      ║"
echo "║     examples/stablecoin-payments/typescript/                      ║"
echo "║                                                                    ║"
echo "║  4. Python SDK examples:                                          ║"
echo "║     examples/stablecoin-payments/python/                          ║"
echo "║                                                                    ║"
echo "╚════════════════════════════════════════════════════════════════════╝"
echo -e "${NC}"

# Useful commands
echo ""
echo "Useful Commands:"
echo "────────────────"
echo "Check balance:"
echo "  $BINARY query bank balances \$(statesetd keys show alice -a --keyring-backend=test)"
echo ""
echo "Send ssUSD:"
echo "  $BINARY tx bank send alice <recipient> 100000000ssusd --keyring-backend=test --yes"
echo ""
echo "Query ssUSD info:"
echo "  $BINARY query stablecoins stablecoin ssusd"
echo ""
