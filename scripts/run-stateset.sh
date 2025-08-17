#!/bin/bash

echo "========================================="
echo "Starting StateSet Blockchain with STST"
echo "========================================="

# Clean up any existing data
docker-compose down 2>/dev/null || true
docker run --rm -v $(pwd):/workspace alpine sh -c "rm -rf /workspace/stateset-data" 2>/dev/null || true
mkdir -p stateset-data

# Start the blockchain with custom configuration
docker run -d \
  --name stateset-blockchain \
  -p 26656:26656 \
  -p 26657:26657 \
  -p 1317:1317 \
  -p 9090:9090 \
  -v $(pwd)/stateset-data:/root/.wasmd \
  cosmwasm/wasmd:v0.50.0 \
  sh -c '
    if [ ! -f /root/.wasmd/config/genesis.json ]; then
      echo "Initializing blockchain..."
      wasmd init stateset-node --chain-id stateset-1
      
      # Update denominations to stst
      sed -i "s/\"stake\"/\"stst\"/g" /root/.wasmd/config/genesis.json
      
      # Create validator key
      wasmd keys add validator --keyring-backend test
      
      # Add genesis account
      wasmd genesis add-genesis-account validator 100000000000stst --keyring-backend test
      
      # Create genesis transaction
      wasmd genesis gentx validator 100000000stst --chain-id stateset-1 --keyring-backend test
      
      # Collect genesis transactions
      wasmd genesis collect-gentxs
      
      # Configure for development
      sed -i "s/minimum-gas-prices = \"\"/minimum-gas-prices = \"0.025stst\"/" /root/.wasmd/config/app.toml
      sed -i "s/enable = false/enable = true/" /root/.wasmd/config/app.toml
    fi
    
    echo "Starting node with STST tokens..."
    echo "Note: Addresses will show as wasm1... but represent StateSet addresses"
    wasmd start --api.enable --api.swagger
  '

echo ""
echo "Waiting for chain to start..."
sleep 10

# Check if chain is running
docker exec stateset-blockchain wasmd status 2>/dev/null | grep -q "catching_up" && {
  echo "✅ Chain is running!"
  echo ""
  echo "Services available at:"
  echo "  - RPC:      http://localhost:26657"
  echo "  - REST API: http://localhost:1317"
  echo "  - gRPC:     localhost:9090"
  echo ""
  echo "To create test accounts and send STST:"
  echo "  docker exec stateset-blockchain wasmd keys add alice --keyring-backend test"
  echo "  docker exec stateset-blockchain wasmd keys add bob --keyring-backend test"
  echo ""
  echo "To check balances:"
  echo "  docker exec stateset-blockchain wasmd query bank total"
  echo ""
  echo "To send STST tokens:"
  echo "  docker exec stateset-blockchain wasmd tx bank send <from> <to> <amount>stst --keyring-backend test --chain-id stateset-1 --yes"
} || {
  echo "⚠️  Chain is still starting. Check logs with:"
  echo "  docker logs stateset-blockchain"
}