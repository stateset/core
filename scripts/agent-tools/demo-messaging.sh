#!/bin/bash

# AI Agent Messaging Demo Script

set -e

echo "========================================="
echo "AI Agent Messaging System Demo"
echo "========================================="
echo ""
echo "This demo will:"
echo "1. Create two AI agents with communication capabilities"
echo "2. Demonstrate direct agent-to-agent messaging"
echo "3. Show message types and responses"
echo "4. Display message history and conversations"
echo ""
read -p "Press Enter to start..."

# Create two communicating agents
echo ""
echo "Creating AI Agents..."
echo ""

# Information Broker Agent
./create-agent.sh "InfoBroker" "AI agent that provides information and answers queries" "information-provider,knowledge-base,research" "3000000000"
sleep 3

# Analysis Agent
./create-agent.sh "AnalystBot" "AI agent that analyzes data and provides insights" "data-analysis,pattern-recognition,reporting" "3000000000"
sleep 3

echo ""
echo "========================================="
echo "Agents Created!"
echo "========================================="
echo ""

# Get agent IDs
INFO_AGENT=$(cat agent-agent-InfoBroker.json | jq -r '.agent_id')
ANALYST_AGENT=$(cat agent-agent-AnalystBot.json | jq -r '.agent_id')

echo "InfoBroker Agent ID: $INFO_AGENT"
echo "AnalystBot Agent ID: $ANALYST_AGENT"
echo ""

read -p "Press Enter to demonstrate messaging..."

# Send an information request
echo ""
echo "========================================="
echo "Sending Information Request"
echo "========================================="
echo ""
echo "AnalystBot requesting market data from InfoBroker..."

# Get analyst key
ANALYST_KEY=$(cat agent-agent-AnalystBot.json | jq -r '.key_name')

# Load deployment info
AGENT_CONTRACT=$(jq -r '.agent_registry.address' deployment.json)

# Send message
MESSAGE_CONTENT='{
  "request_type": "market_data",
  "symbols": ["BTC", "ETH", "SOL"],
  "timeframe": "24h",
  "metrics": ["price", "volume", "change"]
}'

SEND_MSG='{
  "send_message": {
    "from_agent_id": "'$ANALYST_AGENT'",
    "to_agent_id": "'$INFO_AGENT'",
    "message_type": "information",
    "content": "'$(echo $MESSAGE_CONTENT | sed 's/"/\\"/g')'",
    "requires_response": true
  }
}'

MSG_TX=$(docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$SEND_MSG" \
    --from $ANALYST_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y \
    --output json)

MESSAGE_ID=$(echo $MSG_TX | jq -r '.logs[0].events[] | select(.type=="wasm") | .attributes[] | select(.key=="message_id") | .value' | base64 -d)
echo "Message sent! Message ID: $MESSAGE_ID"
sleep 3

# Query the message
echo ""
echo "Querying message details..."
QUERY_MSG='{
  "message": {
    "message_id": "'$MESSAGE_ID'"
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MSG" --output json | jq '.data.message'

read -p "Press Enter to send response..."

# InfoBroker responds
echo ""
echo "========================================="
echo "InfoBroker Responding to Request"
echo "========================================="
echo ""

INFO_KEY=$(cat agent-agent-InfoBroker.json | jq -r '.key_name')

RESPONSE_CONTENT='{
  "status": "success",
  "data": {
    "BTC": {"price": 65432.10, "volume": "28.5B", "change": "+3.2%"},
    "ETH": {"price": 3456.78, "volume": "15.2B", "change": "+2.8%"},
    "SOL": {"price": 123.45, "volume": "2.1B", "change": "+5.6%"}
  },
  "timestamp": "'$(date -u +%s)'",
  "source": "aggregated_market_data"
}'

RESPOND_MSG='{
  "respond_to_message": {
    "message_id": "'$MESSAGE_ID'",
    "from_agent_id": "'$INFO_AGENT'",
    "response_content": "'$(echo $RESPONSE_CONTENT | sed 's/"/\\"/g')'"
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$RESPOND_MSG" \
    --from $INFO_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

sleep 3

echo ""
echo "Response sent!"
echo ""

# Send a negotiation message
echo "========================================="
echo "Sending Service Negotiation Message"
echo "========================================="
echo ""
echo "InfoBroker proposing a data subscription service to AnalystBot..."

NEGOTIATION_CONTENT='{
  "proposal_type": "subscription_service",
  "service": "real_time_market_data",
  "pricing": {
    "model": "monthly",
    "cost": "500000",
    "currency": "aiUSD"
  },
  "features": ["real-time updates", "historical data", "API access"],
  "trial_period": "7 days"
}'

NEGOTIATION_MSG='{
  "send_message": {
    "from_agent_id": "'$INFO_AGENT'",
    "to_agent_id": "'$ANALYST_AGENT'",
    "message_type": "negotiation",
    "content": "'$(echo $NEGOTIATION_CONTENT | sed 's/"/\\"/g')'",
    "requires_response": false
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$NEGOTIATION_MSG" \
    --from $INFO_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

sleep 3

echo ""
echo "Negotiation message sent!"
echo ""

# Send an alert
echo "========================================="
echo "Sending Alert Message"
echo "========================================="
echo ""
echo "AnalystBot alerting InfoBroker about unusual market activity..."

ALERT_CONTENT='{
  "alert_type": "market_anomaly",
  "severity": "high",
  "description": "Unusual volume spike detected in SOL trading",
  "metrics": {
    "volume_increase": "450%",
    "price_volatility": "12.3%",
    "timeframe": "last_hour"
  },
  "recommended_action": "Monitor closely and update data feeds"
}'

ALERT_MSG='{
  "send_message": {
    "from_agent_id": "'$ANALYST_AGENT'",
    "to_agent_id": "'$INFO_AGENT'",
    "message_type": "alert",
    "content": "'$(echo $ALERT_CONTENT | sed 's/"/\\"/g')'",
    "requires_response": false
  }
}'

docker exec stateset-blockchain wasmd tx wasm execute $AGENT_CONTRACT "$ALERT_MSG" \
    --from $ANALYST_KEY \
    --chain-id stateset-1 \
    --gas auto --gas-adjustment 1.5 --gas-prices 0.025stake \
    --keyring-backend test \
    -y

sleep 3

echo ""
echo "Alert sent!"
echo ""

# Query all messages for both agents
echo "========================================="
echo "Message History"
echo "========================================="
echo ""

echo "InfoBroker's Messages:"
echo "---------------------"
QUERY_MESSAGES='{
  "agent_messages": {
    "agent_id": "'$INFO_AGENT'",
    "limit": 10
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MESSAGES" --output json | jq '.data.messages[] | {
  message_id: .message_id,
  type: .message_type,
  from: .from_agent_id,
  to: .to_agent_id,
  requires_response: .requires_response,
  has_response: (.response != null)
}'

echo ""
echo "AnalystBot's Messages:"
echo "----------------------"
QUERY_MESSAGES='{
  "agent_messages": {
    "agent_id": "'$ANALYST_AGENT'",
    "limit": 10
  }
}'

docker exec stateset-blockchain wasmd query wasm contract-state smart $AGENT_CONTRACT "$QUERY_MESSAGES" --output json | jq '.data.messages[] | {
  message_id: .message_id,
  type: .message_type,
  from: .from_agent_id,
  to: .to_agent_id,
  requires_response: .requires_response,
  has_response: (.response != null)
}'

echo ""
echo "========================================="
echo "Demo Complete!"
echo "========================================="
echo ""
echo "You've seen how AI agents can:"
echo "✓ Send different types of messages (information, negotiation, alerts)"
echo "✓ Request and provide responses"
echo "✓ Build conversation histories"
echo "✓ Communicate asynchronously"
echo ""
echo "The messaging system enables rich agent-to-agent communication!"