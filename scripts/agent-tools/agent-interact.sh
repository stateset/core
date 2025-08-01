#!/bin/bash

# AI Agent Interaction Script

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
    echo "AI Agent Interaction Tool"
    echo ""
    echo "Usage: $0 <command> [arguments]"
    echo ""
    echo "Commands:"
    echo "  transfer <from_agent_id> <to_agent_id> <amount> [memo]"
    echo "      Transfer stablecoins between agents"
    echo ""
    echo "  request-service <requester_id> <provider_id> <service_type> <payment> <parameters>"
    echo "      Request a service from another agent"
    echo ""
    echo "  complete-service <service_id> <result>"
    echo "      Complete a service request"
    echo ""
    echo "  query-agent <agent_id>"
    echo "      Query agent information"
    echo ""
    echo "  query-balance <agent_id>"
    echo "      Query agent balance"
    echo ""
    echo "  list-agents [limit]"
    echo "      List all registered agents"
    echo ""
    echo "  list-services [agent_id]"
    echo "      List services (optionally filtered by agent)"
    echo ""
    exit 1
}

# Check if command provided
if [ $# -lt 1 ]; then
    usage
fi

COMMAND=$1

case $COMMAND in
    "transfer")
        if [ $# -lt 4 ]; then
            echo "Error: transfer requires <from_agent_id> <to_agent_id> <amount> [memo]"
            exit 1
        fi
        
        FROM_AGENT=$2
        TO_AGENT=$3
        AMOUNT=$4
        MEMO=${5:-"Agent transfer"}
        
        echo "Transferring $AMOUNT aiUSD from $FROM_AGENT to $TO_AGENT..."
        
        # Get agent info to find the key
        FROM_KEY=$(cat agent-*.json | jq -r --arg id "$FROM_AGENT" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$FROM_KEY" ]; then
            echo "Error: Could not find key for agent $FROM_AGENT"
            exit 1
        fi
        
        TRANSFER_MSG='{
          "agent_transfer": {
            "from_agent_id": "'$FROM_AGENT'",
            "to_agent_id": "'$TO_AGENT'",
            "amount": {
              "denom": "ibc/aiUSD",
              "amount": "'$AMOUNT'"
            },
            "memo": "'$MEMO'"
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$TRANSFER_MSG" \
            --from $FROM_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "request-service")
        if [ $# -lt 5 ]; then
            echo "Error: request-service requires <requester_id> <provider_id> <service_type> <payment> <parameters>"
            exit 1
        fi
        
        REQUESTER=$2
        PROVIDER=$3
        SERVICE_TYPE=$4
        PAYMENT=$5
        PARAMS=${6:-"{}"}
        
        echo "Requesting service $SERVICE_TYPE from $PROVIDER..."
        
        # Get requester key
        REQUESTER_KEY=$(cat agent-*.json | jq -r --arg id "$REQUESTER" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$REQUESTER_KEY" ]; then
            echo "Error: Could not find key for agent $REQUESTER"
            exit 1
        fi
        
        SERVICE_MSG='{
          "request_service": {
            "requester_agent_id": "'$REQUESTER'",
            "provider_agent_id": "'$PROVIDER'",
            "service_type": "'$SERVICE_TYPE'",
            "payment": {
              "denom": "ibc/aiUSD",
              "amount": "'$PAYMENT'"
            },
            "parameters": "'$(echo $PARAMS | sed 's/"/\\"/g')'"
          }
        }'
        
        SERVICE_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$SERVICE_MSG" \
            --from $REQUESTER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y \
            --output json)
        
        SERVICE_ID=$(echo $SERVICE_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="service_id") | .value' | base64 -d)
        echo "Service requested! Service ID: $SERVICE_ID"
        ;;
        
    "complete-service")
        if [ $# -lt 3 ]; then
            echo "Error: complete-service requires <service_id> <result>"
            exit 1
        fi
        
        SERVICE_ID=$2
        RESULT=$3
        
        echo "Completing service $SERVICE_ID..."
        
        # Query service to get provider
        SERVICE_QUERY='{
          "service": {
            "service_id": "'$SERVICE_ID'"
          }
        }'
        
        SERVICE_INFO=$(docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$SERVICE_QUERY" --output json)
        PROVIDER_ID=$(echo $SERVICE_INFO | jq -r '.data.provider_agent_id')
        
        # Get provider key
        PROVIDER_KEY=$(cat agent-*.json | jq -r --arg id "$PROVIDER_ID" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$PROVIDER_KEY" ]; then
            echo "Error: Could not find key for provider agent $PROVIDER_ID"
            exit 1
        fi
        
        COMPLETE_MSG='{
          "complete_service": {
            "service_id": "'$SERVICE_ID'",
            "result": "'$(echo $RESULT | sed 's/"/\\"/g')'"
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$COMPLETE_MSG" \
            --from $PROVIDER_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "query-agent")
        if [ $# -lt 2 ]; then
            echo "Error: query-agent requires <agent_id>"
            exit 1
        fi
        
        AGENT_ID=$2
        QUERY_MSG='{
          "agent": {
            "agent_id": "'$AGENT_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "query-balance")
        if [ $# -lt 2 ]; then
            echo "Error: query-balance requires <agent_id>"
            exit 1
        fi
        
        AGENT_ID=$2
        QUERY_MSG='{
          "agent_balance": {
            "agent_id": "'$AGENT_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "list-agents")
        LIMIT=${2:-20}
        QUERY_MSG='{
          "list_agents": {
            "limit": '$LIMIT'
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "list-services")
        AGENT_ID=$2
        if [ -z "$AGENT_ID" ]; then
            QUERY_MSG='{
              "list_services": {
                "limit": 20
              }
            }'
        else
            QUERY_MSG='{
              "list_services": {
                "agent_id": "'$AGENT_ID'",
                "limit": 20
              }
            }'
        fi
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    *)
        echo "Error: Unknown command '$COMMAND'"
        usage
        ;;
esac