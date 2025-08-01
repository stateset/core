#!/bin/bash

# Script to create an AI Agent

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
STABLECOIN_CONTRACT=$(jq -r '.stablecoin.address' deployment.json)

# Parse arguments
AGENT_NAME=${1:-"AI-Agent-$(date +%s)"}
AGENT_DESC=${2:-"Autonomous AI Agent"}
CAPABILITIES=${3:-"data-analysis,nlp,prediction"}
INITIAL_BALANCE=${4:-"1000000000"} # 1000 aiUSD

echo "========================================="
echo "Creating AI Agent"
echo "========================================="
echo "Name: $AGENT_NAME"
echo "Description: $AGENT_DESC"
echo "Capabilities: $CAPABILITIES"
echo "Initial Balance: $INITIAL_BALANCE aiUSD"
echo ""

# Create agent account
AGENT_KEY="agent-${AGENT_NAME// /-}"
echo "Creating agent account: $AGENT_KEY"
docker exec stateset-blockchain wasmd keys add $AGENT_KEY $KEYRING

# Get agent address
AGENT_ADDRESS=$(docker exec stateset-blockchain wasmd keys show $AGENT_KEY -a $KEYRING)
echo "Agent Address: $AGENT_ADDRESS"

# Fund agent account with gas tokens
echo "Funding agent with gas tokens..."
docker exec stateset-blockchain wasmd tx bank send validator $AGENT_ADDRESS 10000000stake \
    --from validator \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y

sleep 5

# Mint stablecoins for the agent
echo "Minting stablecoins for agent..."
MINT_MSG='{
  "mint": {
    "recipient": "'$AGENT_ADDRESS'",
    "amount": "'$INITIAL_BALANCE'"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $STABLECOIN_CONTRACT "$MINT_MSG" \
    --from validator \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y

sleep 5

# Approve agent registry to handle agent's funds
echo "Approving agent registry..."
APPROVE_MSG='{
  "increase_allowance": {
    "spender": "'$AGENT_CONTRACT'",
    "amount": "'$INITIAL_BALANCE'"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $STABLECOIN_CONTRACT "$APPROVE_MSG" \
    --from $AGENT_KEY \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y

sleep 5

# Register the agent
echo "Registering agent..."
# Convert capabilities to JSON array
CAPS_JSON=$(echo $CAPABILITIES | awk -F',' '{printf "["; for(i=1;i<=NF;i++) printf "\"%s\"%s", $i, (i<NF?",":""); printf "]"}')

REGISTER_MSG='{
  "register_agent": {
    "name": "'$AGENT_NAME'",
    "description": "'$AGENT_DESC'",
    "capabilities": '$CAPS_JSON',
    "service_endpoints": ["http://localhost:8080/agent/'$AGENT_KEY'"],
    "initial_balance": {
      "denom": "ibc/aiUSD",
      "amount": "'$INITIAL_BALANCE'"
    }
  }
}'

REGISTER_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$REGISTER_MSG" \
    --from $AGENT_KEY \
    --amount "${INITIAL_BALANCE}ibc/aiUSD" \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y \
    --output json)

# Extract agent ID from events
AGENT_ID=$(echo $REGISTER_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="agent_id") | .value' | base64 -d)

echo ""
echo "========================================="
echo "AI Agent Created Successfully!"
echo "========================================="
echo "Agent ID: $AGENT_ID"
echo "Agent Key: $AGENT_KEY"
echo "Agent Address: $AGENT_ADDRESS"
echo "Initial Balance: $INITIAL_BALANCE aiUSD"
echo ""

# Save agent info
cat > "agent-${AGENT_KEY}.json" <<EOF
{
  "agent_id": "$AGENT_ID",
  "key_name": "$AGENT_KEY",
  "address": "$AGENT_ADDRESS",
  "name": "$AGENT_NAME",
  "description": "$AGENT_DESC",
  "capabilities": $CAPS_JSON,
  "created_at": "$(date -u +%Y-%m-%dT%H:%M:%SZ)"
}
EOF

echo "Agent info saved to agent-${AGENT_KEY}.json"