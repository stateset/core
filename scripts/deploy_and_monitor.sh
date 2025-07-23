#!/bin/bash

# Stateset Core Blockchain Deployment and Monitoring Script
# This script handles deployment, health checks, and monitoring setup

set -e

# Configuration
PROJECT_NAME="stateset-core"
CHAIN_ID="stateset-1"
MONIKER="stateset-validator"
VERSION="v1.0.0"
HOME_DIR="$HOME/.stateset"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Logging function
log() {
    echo -e "${GREEN}[$(date +'%Y-%m-%d %H:%M:%S')] $1${NC}"
}

error() {
    echo -e "${RED}[ERROR] $1${NC}" >&2
}

warn() {
    echo -e "${YELLOW}[WARNING] $1${NC}"
}

info() {
    echo -e "${BLUE}[INFO] $1${NC}"
}

# Check prerequisites
check_prerequisites() {
    log "Checking prerequisites..."
    
    # Check if Go is installed
    if ! command -v go &> /dev/null; then
        error "Go is not installed. Please install Go 1.21 or later."
        exit 1
    fi
    
    # Check Go version
    GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
    if [[ $(printf '%s\n' "1.21" "$GO_VERSION" | sort -V | head -n1) != "1.21" ]]; then
        error "Go version 1.21 or later is required. Current version: $GO_VERSION"
        exit 1
    fi
    
    # Check if Docker is installed (for monitoring stack)
    if ! command -v docker &> /dev/null; then
        warn "Docker is not installed. Monitoring stack will be limited."
    fi
    
    log "Prerequisites check completed successfully"
}

# Build the blockchain binary
build_binary() {
    log "Building Stateset Core binary..."
    
    # Clean previous builds
    go clean -cache
    
    # Build the binary
    make build
    
    if [ $? -eq 0 ]; then
        log "Binary built successfully"
    else
        error "Failed to build binary"
        exit 1
    fi
}

# Initialize the blockchain
initialize_chain() {
    log "Initializing blockchain..."
    
    # Remove existing data if any
    if [ -d "$HOME_DIR" ]; then
        warn "Existing blockchain data found. Removing..."
        rm -rf "$HOME_DIR"
    fi
    
    # Initialize the chain
    ./build/statesetd init $MONIKER --chain-id $CHAIN_ID --home $HOME_DIR
    
    # Add genesis account
    ./build/statesetd keys add validator --keyring-backend test --home $HOME_DIR
    ./build/statesetd add-genesis-account validator 100000000000ustate --keyring-backend test --home $HOME_DIR
    
    # Create genesis transaction
    ./build/statesetd gentx validator 1000000ustate --keyring-backend test --chain-id $CHAIN_ID --home $HOME_DIR
    
    # Collect genesis transactions
    ./build/statesetd collect-gentxs --home $HOME_DIR
    
    log "Blockchain initialized successfully"
}

# Start the blockchain
start_blockchain() {
    log "Starting Stateset Core blockchain..."
    
    # Start the node in background
    nohup ./build/statesetd start --home $HOME_DIR > blockchain.log 2>&1 &
    BLOCKCHAIN_PID=$!
    
    echo $BLOCKCHAIN_PID > blockchain.pid
    
    # Wait for the node to start
    sleep 10
    
    # Check if the process is running
    if ps -p $BLOCKCHAIN_PID > /dev/null; then
        log "Blockchain started successfully (PID: $BLOCKCHAIN_PID)"
    else
        error "Failed to start blockchain"
        exit 1
    fi
}

# Health check function
health_check() {
    log "Performing health check..."
    
    # Check if the process is still running
    if [ -f "blockchain.pid" ]; then
        PID=$(cat blockchain.pid)
        if ps -p $PID > /dev/null; then
            info "✓ Blockchain process is running (PID: $PID)"
        else
            error "✗ Blockchain process is not running"
            return 1
        fi
    else
        error "✗ PID file not found"
        return 1
    fi
    
    # Check RPC endpoint
    if curl -s http://localhost:26657/status > /dev/null; then
        info "✓ RPC endpoint is responsive"
        
        # Get blockchain status
        STATUS=$(curl -s http://localhost:26657/status)
        BLOCK_HEIGHT=$(echo $STATUS | jq -r '.result.sync_info.latest_block_height')
        CATCHING_UP=$(echo $STATUS | jq -r '.result.sync_info.catching_up')
        
        info "✓ Current block height: $BLOCK_HEIGHT"
        info "✓ Catching up: $CATCHING_UP"
    else
        error "✗ RPC endpoint is not responsive"
        return 1
    fi
    
    # Check API endpoint
    if curl -s http://localhost:1317/cosmos/base/tendermint/v1beta1/node_info > /dev/null; then
        info "✓ API endpoint is responsive"
    else
        warn "! API endpoint is not responsive"
    fi
    
    log "Health check completed"
}

# Setup monitoring
setup_monitoring() {
    log "Setting up monitoring..."
    
    # Create monitoring directory
    mkdir -p monitoring
    
    # Create Prometheus configuration
    cat > monitoring/prometheus.yml << EOF
global:
  scrape_interval: 15s
  evaluation_interval: 15s

rule_files:
  - "alert_rules.yml"

scrape_configs:
  - job_name: 'stateset-core'
    static_configs:
      - targets: ['localhost:26660']  # Tendermint metrics
    scrape_interval: 5s
    metrics_path: /metrics

  - job_name: 'node-exporter'
    static_configs:
      - targets: ['localhost:9100']
    scrape_interval: 15s

  - job_name: 'performance-monitor'
    static_configs:
      - targets: ['localhost:8080']  # Our custom performance monitor
    scrape_interval: 10s

alerting:
  alertmanagers:
    - static_configs:
        - targets:
          - localhost:9093
EOF

    # Create alert rules
    cat > monitoring/alert_rules.yml << EOF
groups:
  - name: stateset_core_alerts
    rules:
      - alert: HighBlockTime
        expr: tendermint_consensus_block_interval_seconds > 10
        for: 2m
        labels:
          severity: warning
        annotations:
          summary: "High block time detected"
          description: "Block time is {{ \$value }} seconds, which is above normal."

      - alert: NodeDown
        expr: up{job="stateset-core"} == 0
        for: 1m
        labels:
          severity: critical
        annotations:
          summary: "Stateset Core node is down"
          description: "The Stateset Core node has been down for more than 1 minute."

      - alert: HighMemoryUsage
        expr: (node_memory_MemTotal_bytes - node_memory_MemAvailable_bytes) / node_memory_MemTotal_bytes * 100 > 90
        for: 5m
        labels:
          severity: warning
        annotations:
          summary: "High memory usage"
          description: "Memory usage is above 90% for more than 5 minutes."

      - alert: SecurityAlert
        expr: stateset_security_alerts_total > 0
        for: 0s
        labels:
          severity: critical
        annotations:
          summary: "Security alert triggered"
          description: "A security alert has been triggered in the Stateset Core blockchain."
EOF

    # Start performance monitor
    if [ -f "performance_monitor.go" ]; then
        log "Starting performance monitor..."
        nohup go run performance_monitor.go > performance_monitor.log 2>&1 &
        echo $! > performance_monitor.pid
    fi
    
    log "Monitoring setup completed"
}

# Performance benchmarking
benchmark_performance() {
    log "Running performance benchmarks..."
    
    # Wait for blockchain to be ready
    sleep 30
    
    # Run transaction throughput test
    info "Testing transaction throughput..."
    
    # Create test accounts
    for i in {1..10}; do
        ./build/statesetd keys add testuser$i --keyring-backend test --home $HOME_DIR > /dev/null 2>&1
    done
    
    # Send test transactions
    START_TIME=$(date +%s)
    for i in {1..100}; do
        ./build/statesetd tx bank send validator $(./build/statesetd keys show testuser1 -a --keyring-backend test --home $HOME_DIR) 1000ustate \
            --keyring-backend test --chain-id $CHAIN_ID --home $HOME_DIR --yes --broadcast-mode async > /dev/null 2>&1 &
        
        if [ $((i % 10)) -eq 0 ]; then
            wait  # Wait for batch to complete
            sleep 1
        fi
    done
    wait
    
    END_TIME=$(date +%s)
    DURATION=$((END_TIME - START_TIME))
    TPS=$(echo "scale=2; 100 / $DURATION" | bc)
    
    info "Transaction throughput: $TPS TPS"
    
    # Test query performance
    info "Testing query performance..."
    QUERY_START=$(date +%s%3N)
    ./build/statesetd query bank balances $(./build/statesetd keys show validator -a --keyring-backend test --home $HOME_DIR) --home $HOME_DIR > /dev/null 2>&1
    QUERY_END=$(date +%s%3N)
    QUERY_TIME=$((QUERY_END - QUERY_START))
    
    info "Query response time: ${QUERY_TIME}ms"
    
    log "Performance benchmarking completed"
}

# Backup function
backup_chain_data() {
    log "Creating backup of chain data..."
    
    BACKUP_DIR="backups/$(date +%Y%m%d_%H%M%S)"
    mkdir -p "$BACKUP_DIR"
    
    # Stop the blockchain temporarily
    if [ -f "blockchain.pid" ]; then
        PID=$(cat blockchain.pid)
        kill $PID
        sleep 5
    fi
    
    # Create backup
    tar -czf "$BACKUP_DIR/chain_data.tar.gz" -C $HOME_DIR .
    
    # Restart blockchain
    start_blockchain
    
    log "Backup created at $BACKUP_DIR/chain_data.tar.gz"
}

# Upgrade function
upgrade_blockchain() {
    log "Upgrading blockchain..."
    
    # Stop current instance
    if [ -f "blockchain.pid" ]; then
        PID=$(cat blockchain.pid)
        kill $PID
        wait $PID 2>/dev/null || true
    fi
    
    # Create backup before upgrade
    backup_chain_data
    
    # Build new version
    build_binary
    
    # Start upgraded version
    start_blockchain
    
    log "Blockchain upgrade completed"
}

# Main deployment function
deploy() {
    log "Starting Stateset Core blockchain deployment..."
    
    check_prerequisites
    build_binary
    initialize_chain
    start_blockchain
    setup_monitoring
    
    # Wait for blockchain to stabilize
    sleep 30
    
    health_check
    benchmark_performance
    
    log "Deployment completed successfully!"
    log "Blockchain is running on:"
    log "  RPC: http://localhost:26657"
    log "  API: http://localhost:1317"
    log "  Monitoring: http://localhost:8080"
}

# Stop function
stop() {
    log "Stopping Stateset Core blockchain..."
    
    # Stop blockchain
    if [ -f "blockchain.pid" ]; then
        PID=$(cat blockchain.pid)
        kill $PID
        rm blockchain.pid
        log "Blockchain stopped"
    fi
    
    # Stop performance monitor
    if [ -f "performance_monitor.pid" ]; then
        PID=$(cat performance_monitor.pid)
        kill $PID
        rm performance_monitor.pid
        log "Performance monitor stopped"
    fi
}

# Status function
status() {
    log "Checking Stateset Core blockchain status..."
    health_check
}

# Help function
show_help() {
    echo "Stateset Core Blockchain Deployment Script"
    echo ""
    echo "Usage: $0 {deploy|stop|status|health|benchmark|backup|upgrade|help}"
    echo ""
    echo "Commands:"
    echo "  deploy     - Full deployment of the blockchain"
    echo "  stop       - Stop the running blockchain"
    echo "  status     - Check blockchain status"
    echo "  health     - Perform health check"
    echo "  benchmark  - Run performance benchmarks"
    echo "  backup     - Create backup of chain data"
    echo "  upgrade    - Upgrade the blockchain"
    echo "  help       - Show this help message"
}

# Main script logic
case "$1" in
    deploy)
        deploy
        ;;
    stop)
        stop
        ;;
    status)
        status
        ;;
    health)
        health_check
        ;;
    benchmark)
        benchmark_performance
        ;;
    backup)
        backup_chain_data
        ;;
    upgrade)
        upgrade_blockchain
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        show_help
        exit 1
        ;;
esac