#!/bin/bash

# AI Agent System Deployment Script

set -e

CHAIN_ID="stateset-1"
NODE="http://localhost:26657"
KEYRING="--keyring-backend test"
GAS_FLAGS="--gas auto --gas-adjustment 1.5 --gas-prices 0.025stake"

echo "========================================="
echo "Deploying AI Agent System Contracts"
echo "========================================="

# Get validator address
VALIDATOR=$(docker exec stateset-blockchain wasmd keys show validator -a $KEYRING)
echo "Deploying from: $VALIDATOR"

# Build contracts
echo "Building contracts..."
cd contracts

# Build AI Stablecoin
echo "Building AI Stablecoin contract..."
cd stablecoin
cargo wasm
cargo run-script optimize
cd ..

# Build AI Agent Registry
echo "Building AI Agent Registry contract..."
cd ai-agents
cargo wasm
cargo run-script optimize
cd ..

cd ..

# Copy optimized contracts
echo "Copying optimized contracts..."
mkdir -p artifacts
cp contracts/stablecoin/artifacts/*.wasm artifacts/
cp contracts/ai-agents/artifacts/*.wasm artifacts/

# Deploy Stablecoin
echo ""
echo "Deploying AI Stablecoin..."
STABLECOIN_TX=$(docker exec stateset-blockchain wasmd tx wasm store /artifacts/ai_stablecoin.wasm \
    --from validator \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y \
    --output json \
    --broadcast-mode sync)

STABLECOIN_CODE_ID=$(echo $STABLECOIN_TX | jq -r '.logs[0].events[] | select(.type=="store_code") | .attributes[] | select(.key=="code_id") | .value')
echo "Stablecoin Code ID: $STABLECOIN_CODE_ID"

# Instantiate Stablecoin
echo "Instantiating AI Stablecoin..."
STABLECOIN_INIT='{
  "name": "AI USD",
  "symbol": "aiUSD",
  "decimals": 6,
  "initial_balances": [{
    "address": "'$VALIDATOR'",
    "amount": "1000000000000000"
  }],
  "mint": {
    "minter": "'$VALIDATOR'",
    "cap": null
  },
  "marketing": null
}'

STABLECOIN_INSTANTIATE_TX=$(docker exec stateset-blockchain wasmd tx wasm instantiate $STABLECOIN_CODE_ID "$STABLECOIN_INIT" \
    --from validator \
    --label "AI Stablecoin" \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y \
    --output json \
    --broadcast-mode sync)

# Get stablecoin contract address
STABLECOIN_CONTRACT=$(docker exec stateset-blockchain wasmd query wasm list-contract-by-code $STABLECOIN_CODE_ID --output json | jq -r '.contracts[0]')
echo "Stablecoin Contract: $STABLECOIN_CONTRACT"

# Deploy AI Agent Registry
echo ""
echo "Deploying AI Agent Registry..."
AGENT_TX=$(docker exec stateset-blockchain wasmd tx wasm store /artifacts/ai_agent_registry.wasm \
    --from validator \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y \
    --output json \
    --broadcast-mode sync)

AGENT_CODE_ID=$(echo $AGENT_TX | jq -r '.logs[0].events[] | select(.type=="store_code") | .attributes[] | select(.key=="code_id") | .value')
echo "Agent Registry Code ID: $AGENT_CODE_ID"

# Instantiate Agent Registry
echo "Instantiating AI Agent Registry..."
AGENT_INIT='{
  "stablecoin_denom": "ibc/aiUSD",
  "min_agent_balance": "1000000",
  "service_fee_percentage": 250
}'

AGENT_INSTANTIATE_TX=$(docker exec stateset-blockchain wasmd tx wasm instantiate $AGENT_CODE_ID "$AGENT_INIT" \
    --from validator \
    --label "AI Agent Registry" \
    --chain-id $CHAIN_ID \
    $GAS_FLAGS \
    $KEYRING \
    -y \
    --output json \
    --broadcast-mode sync)

# Get agent registry contract address
AGENT_CONTRACT=$(docker exec stateset-blockchain wasmd query wasm list-contract-by-code $AGENT_CODE_ID --output json | jq -r '.contracts[0]')
echo "Agent Registry Contract: $AGENT_CONTRACT"

# Save contract addresses
echo ""
echo "Saving deployment info..."
cat > deployment.json <<EOF
{
  "chain_id": "$CHAIN_ID",
  "stablecoin": {
    "code_id": $STABLECOIN_CODE_ID,
    "address": "$STABLECOIN_CONTRACT",
    "symbol": "aiUSD",
    "decimals": 6
  },
  "agent_registry": {
    "code_id": $AGENT_CODE_ID,
    "address": "$AGENT_CONTRACT"
  },
  "deployer": "$VALIDATOR"
}
EOF

echo ""
echo "========================================="
echo "Deployment Complete!"
echo "========================================="
echo "Stablecoin: $STABLECOIN_CONTRACT"
echo "Agent Registry: $AGENT_CONTRACT"
echo ""
echo "Deployment info saved to deployment.json"