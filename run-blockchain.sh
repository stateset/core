#!/bin/bash

echo "========================================="
echo "Starting StateSet Blockchain"
echo "========================================="

# Create data directory
mkdir -p data

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo "Error: Docker is not running. Please start Docker first."
    exit 1
fi

# Stop any existing containers
echo "Stopping any existing containers..."
docker-compose down 2>/dev/null || true

# Start the blockchain
echo "Starting blockchain node..."
docker-compose up -d stateset-node

# Wait for node to be ready
echo "Waiting for blockchain to initialize..."
sleep 10

# Check node status
echo "Checking node status..."
docker exec stateset-blockchain wasmd status 2>/dev/null || {
    echo "Node is still initializing, waiting..."
    sleep 10
    docker exec stateset-blockchain wasmd status 2>/dev/null || {
        echo "Node initialization might be taking longer. Check logs with:"
        echo "  docker logs stateset-blockchain"
    }
}

echo ""
echo "========================================="
echo "Blockchain is starting!"
echo "========================================="
echo ""
echo "Services available at:"
echo "  - RPC:      http://localhost:26657"
echo "  - REST API: http://localhost:1317"
echo "  - gRPC:     localhost:9090"
echo "  - P2P:      localhost:26656"
echo ""
echo "Useful commands:"
echo "  - View logs:        docker logs -f stateset-blockchain"
echo "  - Check status:     docker exec stateset-blockchain wasmd status"
echo "  - Stop blockchain:  docker-compose down"
echo "  - Get validator address: docker exec stateset-blockchain wasmd keys show validator -a --keyring-backend test"
echo ""
echo "To interact with the blockchain (using STST tokens):"
echo "  docker exec -it stateset-blockchain wasmd query bank balances \$(wasmd keys show validator -a --keyring-backend test)"
echo ""