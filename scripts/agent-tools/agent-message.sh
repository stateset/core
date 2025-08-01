#!/bin/bash

# AI Agent Messaging Tool

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
    echo "AI Agent Messaging Tool"
    echo ""
    echo "Usage: $0 <command> [arguments]"
    echo ""
    echo "Commands:"
    echo "  send <from_agent_id> <to_agent_id> <message_type> <content> [requires_response]"
    echo "      Send a message from one agent to another"
    echo "      message_type: information, service_request, service_response, negotiation, alert, custom"
    echo ""
    echo "  respond <message_id> <from_agent_id> <response_content>"
    echo "      Respond to a message that requires a response"
    echo ""
    echo "  query-messages <agent_id> [message_type] [limit]"
    echo "      Query messages for an agent, optionally filtered by type"
    echo ""
    echo "  query-message <message_id>"
    echo "      Query a specific message by ID"
    echo ""
    echo "Examples:"
    echo "  $0 send agent_123 agent_456 information '{\"query\":\"What is the weather?\"}' true"
    echo "  $0 respond msg_1 agent_456 '{\"answer\":\"Sunny, 72°F\"}'  "
    echo "  $0 query-messages agent_123 information 10"
    echo ""
    exit 1
}

# Check if command provided
if [ $# -lt 1 ]; then
    usage
fi

COMMAND=$1

case $COMMAND in
    "send")
        if [ $# -lt 5 ]; then
            echo "Error: send requires <from_agent_id> <to_agent_id> <message_type> <content> [requires_response]"
            exit 1
        fi
        
        FROM_AGENT=$2
        TO_AGENT=$3
        MSG_TYPE=$4
        CONTENT=$5
        REQUIRES_RESPONSE=${6:-false}
        
        echo "Sending message from $FROM_AGENT to $TO_AGENT..."
        
        # Get sender key
        FROM_KEY=$(cat agent-*.json | jq -r --arg id "$FROM_AGENT" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$FROM_KEY" ]; then
            echo "Error: Could not find key for agent $FROM_AGENT"
            exit 1
        fi
        
        # Handle custom message type
        if [ "$MSG_TYPE" = "custom" ]; then
            MSG_TYPE_JSON='{"custom": "general"}'
        else
            MSG_TYPE_JSON="\"$MSG_TYPE\""
        fi
        
        SEND_MSG='{
          "send_message": {
            "from_agent_id": "'$FROM_AGENT'",
            "to_agent_id": "'$TO_AGENT'",
            "message_type": '$MSG_TYPE_JSON',
            "content": "'$(echo $CONTENT | sed 's/"/\\"/g')'",
            "requires_response": '$REQUIRES_RESPONSE'
          }
        }'
        
        MSG_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$SEND_MSG" \
            --from $FROM_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y \
            --output json)
        
        MESSAGE_ID=$(echo $MSG_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="message_id") | .value' | base64 -d)
        echo "Message sent! Message ID: $MESSAGE_ID"
        ;;
        
    "respond")
        if [ $# -lt 4 ]; then
            echo "Error: respond requires <message_id> <from_agent_id> <response_content>"
            exit 1
        fi
        
        MESSAGE_ID=$2
        FROM_AGENT=$3
        RESPONSE=$4
        
        echo "Responding to message $MESSAGE_ID..."
        
        # Get responder key
        FROM_KEY=$(cat agent-*.json | jq -r --arg id "$FROM_AGENT" 'select(.agent_id == $id) | .key_name' | head -1)
        
        if [ -z "$FROM_KEY" ]; then
            echo "Error: Could not find key for agent $FROM_AGENT"
            exit 1
        fi
        
        RESPOND_MSG='{
          "respond_to_message": {
            "message_id": "'$MESSAGE_ID'",
            "from_agent_id": "'$FROM_AGENT'",
            "response_content": "'$(echo $RESPONSE | sed 's/"/\\"/g')'"
          }
        }'
        
        docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$RESPOND_MSG" \
            --from $FROM_KEY \
            --chain-id $CHAIN_ID \
            $GAS_FLAGS \
            $KEYRING \
            -y
        ;;
        
    "query-messages")
        if [ $# -lt 2 ]; then
            echo "Error: query-messages requires <agent_id> [message_type] [limit]"
            exit 1
        fi
        
        AGENT_ID=$2
        MSG_TYPE=$3
        LIMIT=${4:-20}
        
        QUERY_MSG='{
          "agent_messages": {
            "agent_id": "'$AGENT_ID'",
            "limit": '$LIMIT
        
        if [ ! -z "$MSG_TYPE" ]; then
            if [ "$MSG_TYPE" = "custom" ]; then
                QUERY_MSG="$QUERY_MSG"',
            "message_type": {"custom": "general"}'
            else
                QUERY_MSG="$QUERY_MSG"',
            "message_type": "'$MSG_TYPE'"'
            fi
        fi
        
        QUERY_MSG="$QUERY_MSG"'
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    "query-message")
        if [ $# -lt 2 ]; then
            echo "Error: query-message requires <message_id>"
            exit 1
        fi
        
        MESSAGE_ID=$2
        QUERY_MSG='{
          "message": {
            "message_id": "'$MESSAGE_ID'"
          }
        }'
        
        docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data'
        ;;
        
    *)
        echo "Error: Unknown command '$COMMAND'"
        usage
        ;;
esac