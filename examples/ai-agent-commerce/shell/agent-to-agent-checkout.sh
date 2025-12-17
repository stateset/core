#!/usr/bin/env bash

set -euo pipefail

# -----------------------------------------------------------------------------
# Demo: AI agent → AI agent commerce using native ssUSD (ssusd) on-chain.
#
# This script boots a local single-node chain in a temp directory, mints `ssusd`
# via a CDP vault, then runs on-chain checkout flows between two agent wallets:
# - Scenario A: instant checkout (immediate settlement) via x/settlement
# - Scenario B: escrow checkout + buyer release via x/settlement
# - Scenario C: service purchase via x/payments (payment intent + seller settle)
#
# Usage:
#   bash examples/ai-agent-commerce/shell/agent-to-agent-checkout.sh
#
# Keep state for inspection:
#   KEEP_HOME=1 bash examples/ai-agent-commerce/shell/agent-to-agent-checkout.sh
# -----------------------------------------------------------------------------

BINARY="${BINARY:-}"
if [[ -z "$BINARY" ]]; then
  if [[ -x "./build/statesetd" ]]; then
    BINARY="./build/statesetd"
  else
    BINARY="statesetd"
  fi
fi

CHAIN_ID="${CHAIN_ID:-stateset-agent-commerce-demo}"
KEYRING="${KEYRING:-test}"
RPC_LADDR="${RPC_LADDR:-tcp://127.0.0.1:26657}"
NODE="${NODE:-$RPC_LADDR}"
GRPC_ADDR="${GRPC_ADDR:-127.0.0.1:9090}"
HOME_DIR="${HOME_DIR:-$(mktemp -d -t stateset-agent-commerce-XXXX)}"
LOG_FILE="${LOG_FILE:-$HOME_DIR/statesetd.log}"
KEEP_HOME="${KEEP_HOME:-}"

require_cmd() {
  local cmd="$1"
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "Missing dependency: $cmd" >&2
    exit 1
  fi
}

require_cmd jq

if ! command -v "$BINARY" >/dev/null 2>&1 && [[ ! -x "$BINARY" ]]; then
  echo "statesetd binary not found (BINARY=$BINARY)" >&2
  echo "Build it with: make build" >&2
  exit 1
fi

echo "Using state directory: $HOME_DIR"

"$BINARY" init agent-commerce --chain-id "$CHAIN_ID" --home "$HOME_DIR" >/dev/null

add_key() {
  local name=$1
  "$BINARY" keys add "$name" --keyring-backend "$KEYRING" --home "$HOME_DIR" --output json >/dev/null
}

add_key validator
add_key treasury
add_key buyer-agent
add_key seller-agent

VAL_ADDR=$("$BINARY" keys show validator --keyring-backend "$KEYRING" --home "$HOME_DIR" -a)
TREASURY_ADDR=$("$BINARY" keys show treasury --keyring-backend "$KEYRING" --home "$HOME_DIR" -a)
BUYER_ADDR=$("$BINARY" keys show buyer-agent --keyring-backend "$KEYRING" --home "$HOME_DIR" -a)
SELLER_ADDR=$("$BINARY" keys show seller-agent --keyring-backend "$KEYRING" --home "$HOME_DIR" -a)

# `stst` is used as demo collateral for CDP minting; `stake` pays gas.
GENESIS_FUNDS_VALIDATOR="100000000stake,10000000000stst"
GENESIS_FUNDS_TREASURY="200000000stake,10000000000stst"
GENESIS_FUNDS_BUYER="200000000stake,10000000000stst"
GENESIS_FUNDS_SELLER="200000000stake,10000000000stst"

"$BINARY" genesis add-genesis-account "$VAL_ADDR" "$GENESIS_FUNDS_VALIDATOR" --home "$HOME_DIR" >/dev/null
"$BINARY" genesis add-genesis-account "$TREASURY_ADDR" "$GENESIS_FUNDS_TREASURY" --home "$HOME_DIR" >/dev/null
"$BINARY" genesis add-genesis-account "$BUYER_ADDR" "$GENESIS_FUNDS_BUYER" --home "$HOME_DIR" >/dev/null
"$BINARY" genesis add-genesis-account "$SELLER_ADDR" "$GENESIS_FUNDS_SELLER" --home "$HOME_DIR" >/dev/null

"$BINARY" genesis gentx validator 50000000stake --chain-id "$CHAIN_ID" --keyring-backend "$KEYRING" --home "$HOME_DIR" >/dev/null
"$BINARY" genesis collect-gentxs --home "$HOME_DIR" >/dev/null

GENESIS="$HOME_DIR/config/genesis.json"
TMP_GENESIS="$GENESIS.tmp"

# - Set module authorities to treasury (demo convenience)
# - Enable stablecoin vault minting for local CDP issuance
jq --arg addr "$TREASURY_ADDR" \
  '.app_state.compliance.authority = $addr
   | .app_state.oracle.authority = $addr
   | .app_state.treasury.authority = $addr
   | .app_state.stablecoin.params.vault_minting_enabled = true' \
  "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

echo "Starting node..."
"$BINARY" start \
  --home "$HOME_DIR" \
  --rpc.laddr "$RPC_LADDR" \
  --grpc.address "$GRPC_ADDR" \
  --minimum-gas-prices "0stake" >"$LOG_FILE" 2>&1 &
NODE_PID=$!

cleanup() {
  kill "$NODE_PID" >/dev/null 2>&1 || true
  wait "$NODE_PID" >/dev/null 2>&1 || true
  if [[ -z "$KEEP_HOME" ]]; then
    rm -rf "$HOME_DIR"
  else
    echo "KEEP_HOME set; leaving state dir at $HOME_DIR"
    echo "Node log: $LOG_FILE"
  fi
}
trap cleanup EXIT

for _ in {1..40}; do
  if "$BINARY" status --node "$NODE" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

COMMON_FLAGS=(--home "$HOME_DIR" --keyring-backend "$KEYRING" --chain-id "$CHAIN_ID" --node "$NODE" --yes --broadcast-mode block --gas auto --gas-adjustment 1.3 --fees 2000stake)

tx() {
  "$BINARY" tx "$@" "${COMMON_FLAGS[@]}" --output json
}

query() {
  "$BINARY" query "$@" --node "$NODE" --output json
}

balance() {
  local address=$1
  local denom=$2
  query bank balances "$address" | jq -r --arg denom "$denom" '[.balances[]? | select(.denom==$denom) | .amount][0] // "0"'
}

get_event_attr() {
  local tx_json=$1
  local event_type=$2
  local key=$3
  echo "$tx_json" | jq -r --arg type "$event_type" --arg key "$key" '
    (.tx_response.logs[0].events // [])
    | map(select(.type==$type))
    | (.[0].attributes // [])
    | map(select(.key==$key))
    | (.[0].value // "")
  ' | head -n 1
}

echo "Bootstrapping demo prerequisites..."
echo "  Buyer agent:  $BUYER_ADDR"
echo "  Seller agent: $SELLER_ADDR"

echo "Registering compliance profiles..."
tx compliance upsert-profile "$BUYER_ADDR" basic low --metadata "demo buyer AI agent" --from treasury >/dev/null
tx compliance upsert-profile "$SELLER_ADDR" basic low --metadata "demo seller AI agent" --from treasury >/dev/null

echo "Setting oracle price for stst..."
tx oracle update-price stst 1.00 --from treasury >/dev/null

echo "Minting ssUSD for buyer agent (CDP vault)..."
VAULT_TX=$(tx stablecoin create-vault 2000000000stst --debt 1000000000ssusd --from buyer-agent)
VAULT_ID=$(get_event_attr "$VAULT_TX" "vault_created" "vault_id")
VAULT_ID="${VAULT_ID:-1}"
echo "  Vault ID: $VAULT_ID"

echo "Balances (after mint):"
echo "  Buyer ssusd:  $(balance "$BUYER_ADDR" ssusd)"
echo "  Seller ssusd: $(balance "$SELLER_ADDR" ssusd)"

echo ""
echo "=== Scenario A: Agent checkout (instant settlement) ==="
ORDER_A="AGENT-ORDER-A-$(date +%s)"
META_A=$(jq -nc \
  --arg flow "agent_checkout_v1" \
  --arg order_id "$ORDER_A" \
  --arg buyer_agent_id "agent:buyer:001" \
  --arg seller_agent_id "agent:seller:042" \
  '{
    flow: $flow,
    order_id: $order_id,
    buyer_agent: { id: $buyer_agent_id, role: "checkout" },
    seller_agent: { id: $seller_agent_id, role: "merchant" },
    cart: { items: [{ sku: "SKU-DIGI-001", qty: 1, unit_price_ssusd: "25000000" }], total_ssusd: "25000000" },
    settlement: { method: "instant_checkout", escrow: false }
  }')

CHECKOUT_A_TX=$(tx settlement instant-checkout "$SELLER_ADDR" 25000000ssusd "$ORDER_A" --metadata "$META_A" --from buyer-agent)
SETTLEMENT_A_ID=$(get_event_attr "$CHECKOUT_A_TX" "instant_checkout" "settlement_id")
SETTLEMENT_A_ID="${SETTLEMENT_A_ID:-1}"
echo "  Settlement ID: $SETTLEMENT_A_ID"
echo "  Settlement state:"
query settlement settlement "$SETTLEMENT_A_ID" | jq '.'
echo "  Balances:"
echo "    Buyer ssusd:  $(balance "$BUYER_ADDR" ssusd)"
echo "    Seller ssusd: $(balance "$SELLER_ADDR" ssusd)"

echo ""
echo "=== Scenario B: Agent checkout (escrow + buyer release) ==="
ORDER_B="AGENT-ORDER-B-$(date +%s)"
META_B=$(jq -nc \
  --arg flow "agent_checkout_v1" \
  --arg order_id "$ORDER_B" \
  --arg buyer_agent_id "agent:buyer:001" \
  --arg seller_agent_id "agent:seller:042" \
  '{
    flow: $flow,
    order_id: $order_id,
    buyer_agent: { id: $buyer_agent_id, role: "checkout" },
    seller_agent: { id: $seller_agent_id, role: "merchant" },
    cart: { items: [{ sku: "SKU-PHYS-101", qty: 2, unit_price_ssusd: "25000000" }], total_ssusd: "50000000" },
    settlement: { method: "instant_checkout", escrow: true, release_condition: "delivery_confirmed" }
  }')

CHECKOUT_B_TX=$(tx settlement instant-checkout "$SELLER_ADDR" 50000000ssusd "$ORDER_B" --use-escrow --metadata "$META_B" --from buyer-agent)
SETTLEMENT_B_ID=$(get_event_attr "$CHECKOUT_B_TX" "instant_checkout" "settlement_id")
SETTLEMENT_B_ID="${SETTLEMENT_B_ID:-2}"
echo "  Settlement ID: $SETTLEMENT_B_ID"
echo "  Settlement state (pending escrow):"
query settlement settlement "$SETTLEMENT_B_ID" | jq '.'

echo "  Simulating delivery... buyer agent releases escrow"
tx settlement release-escrow "$SETTLEMENT_B_ID" --from buyer-agent >/dev/null
echo "  Settlement state (completed):"
query settlement settlement "$SETTLEMENT_B_ID" | jq '.'
echo "  Balances:"
echo "    Buyer ssusd:  $(balance "$BUYER_ADDR" ssusd)"
echo "    Seller ssusd: $(balance "$SELLER_ADDR" ssusd)"

echo ""
echo "=== Scenario C: AI service purchase via payment intent (seller claims) ==="
SERVICE_REQ="SRV-$(date +%s)"
PAYMENT_META=$(jq -nc \
  --arg flow "agent_service_purchase_v1" \
  --arg request_id "$SERVICE_REQ" \
  --arg buyer_agent_id "agent:buyer:001" \
  --arg seller_agent_id "agent:seller:042" \
  '{
    flow: $flow,
    request_id: $request_id,
    buyer_agent: { id: $buyer_agent_id, role: "requester" },
    seller_agent: { id: $seller_agent_id, role: "provider" },
    service: { type: "product_recommendation", input: { category: "office", constraints: ["budget<=75ssUSD"] } },
    payment: { denom: "ssusd", amount: "75000000", settle_by: "seller_agent" }
  }')

PAYMENT_TX=$(tx payments create "$SELLER_ADDR" 75000000ssusd --metadata "$PAYMENT_META" --from buyer-agent)
PAYMENT_ID=$(get_event_attr "$PAYMENT_TX" "payment_created" "payment_id")
PAYMENT_ID="${PAYMENT_ID:-1}"
echo "  Payment ID: $PAYMENT_ID"
echo "  Payment state (pending):"
query payments payment "$PAYMENT_ID" | jq '.'

echo "  Seller agent settles after delivering service output"
tx payments settle "$PAYMENT_ID" --from seller-agent >/dev/null
echo "  Payment state (settled):"
query payments payment "$PAYMENT_ID" | jq '.'
echo "  Balances:"
echo "    Buyer ssusd:  $(balance "$BUYER_ADDR" ssusd)"
echo "    Seller ssusd: $(balance "$SELLER_ADDR" ssusd)"

echo ""
echo "Demo complete."

