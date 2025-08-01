#!/bin/bash

# StateSet Blockchain Client Script

CONTAINER="stateset-blockchain"
KEYRING="--keyring-backend test"

case "$1" in
    "status")
        echo "Blockchain Status:"
        docker exec $CONTAINER wasmd status 2>/dev/null | jq '.' || echo "Node not running"
        ;;
    
    "balance")
        ADDRESS=${2:-$(docker exec $CONTAINER wasmd keys show validator -a $KEYRING 2>/dev/null)}
        echo "Checking balance for: $ADDRESS"
        docker exec $CONTAINER wasmd query bank balances $ADDRESS
        ;;
    
    "send")
        if [ -z "$2" ] || [ -z "$3" ]; then
            echo "Usage: $0 send <to_address> <amount>"
            echo "Example: $0 send wasm1... 1000stake"
            exit 1
        fi
        echo "Sending $3 to $2..."
        docker exec $CONTAINER wasmd tx bank send validator $2 $3 \
            --chain-id stateset-1 \
            --gas auto \
            --gas-adjustment 1.5 \
            --gas-prices 0.025stake \
            $KEYRING \
            -y
        ;;
    
    "accounts")
        echo "Available accounts:"
        docker exec $CONTAINER wasmd keys list $KEYRING
        ;;
    
    "create-account")
        NAME=${2:-"user$(date +%s)"}
        echo "Creating new account: $NAME"
        docker exec $CONTAINER wasmd keys add $NAME $KEYRING
        ;;
    
    "height")
        echo "Current block height:"
        docker exec $CONTAINER wasmd status 2>/dev/null | jq -r '.sync_info.latest_block_height'
        ;;
    
    "logs")
        echo "Following blockchain logs (Ctrl+C to stop)..."
        docker logs -f $CONTAINER
        ;;
    
    *)
        echo "StateSet Blockchain Client"
        echo ""
        echo "Usage: $0 <command> [arguments]"
        echo ""
        echo "Commands:"
        echo "  status              - Show blockchain status"
        echo "  balance [address]   - Check account balance"
        echo "  send <to> <amount>  - Send tokens"
        echo "  accounts            - List all accounts"
        echo "  create-account [name] - Create new account"
        echo "  height              - Show current block height"
        echo "  logs                - Follow blockchain logs"
        echo ""
        echo "Examples:"
        echo "  $0 status"
        echo "  $0 balance"
        echo "  $0 send wasm1abc... 1000stake"
        echo "  $0 create-account alice"
        ;;
esac