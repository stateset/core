#!/bin/bash

# AI Agent System Demo Script

set -e

echo "========================================="
echo "AI Agent System Demo"
echo "========================================="
echo ""
echo "This demo will:"
echo "1. Create multiple AI agents with different capabilities"
echo "2. Show agent-to-agent transfers"
echo "3. Demonstrate service requests and payments"
echo "4. Display agent balances and transactions"
echo ""
read -p "Press Enter to start..."

# Create three different AI agents
echo ""
echo "Creating AI Agents..."
echo ""

# Data Analysis Agent
./create-agent.sh "DataMind" "AI agent specialized in data analysis and insights" "data-analysis,statistics,ml-prediction" "5000000000"
sleep 3

# Natural Language Processing Agent  
./create-agent.sh "LinguaBot" "AI agent for natural language processing" "nlp,translation,sentiment-analysis" "3000000000"
sleep 3

# Image Processing Agent
./create-agent.sh "VisionAI" "AI agent for computer vision tasks" "image-recognition,object-detection,ocr" "4000000000"
sleep 3

echo ""
echo "========================================="
echo "AI Agents Created!"
echo "========================================="
echo ""

# Get agent IDs
DATA_AGENT=$(cat agent-agent-DataMind.json | jq -r '.agent_id')
NLP_AGENT=$(cat agent-agent-LinguaBot.json | jq -r '.agent_id')
VISION_AGENT=$(cat agent-agent-VisionAI.json | jq -r '.agent_id')

echo "DataMind Agent ID: $DATA_AGENT"
echo "LinguaBot Agent ID: $NLP_AGENT"
echo "VisionAI Agent ID: $VISION_AGENT"
echo ""

# Show initial balances
echo "Initial Agent Balances:"
echo "----------------------"
./agent-interact.sh query-balance $DATA_AGENT | jq -r '"DataMind: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $NLP_AGENT | jq -r '"LinguaBot: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $VISION_AGENT | jq -r '"VisionAI: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
echo ""

read -p "Press Enter to continue with agent transfers..."

# Agent-to-Agent Transfer Demo
echo ""
echo "========================================="
echo "Agent-to-Agent Transfer Demo"
echo "========================================="
echo ""
echo "DataMind sending 100 aiUSD to LinguaBot for collaboration..."
./agent-interact.sh transfer $DATA_AGENT $NLP_AGENT "100000000" "Collaboration payment"
sleep 5

echo ""
echo "VisionAI sending 50 aiUSD each to DataMind and LinguaBot..."
# This would use batch transfer in real implementation
./agent-interact.sh transfer $VISION_AGENT $DATA_AGENT "50000000" "Service prepayment"
sleep 3
./agent-interact.sh transfer $VISION_AGENT $NLP_AGENT "50000000" "Service prepayment"
sleep 3

echo ""
echo "Updated Balances:"
echo "-----------------"
./agent-interact.sh query-balance $DATA_AGENT | jq -r '"DataMind: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $NLP_AGENT | jq -r '"LinguaBot: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
./agent-interact.sh query-balance $VISION_AGENT | jq -r '"VisionAI: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD"'
echo ""

read -p "Press Enter to continue with service requests..."

# Service Request Demo
echo ""
echo "========================================="
echo "Service Request Demo"
echo "========================================="
echo ""
echo "LinguaBot requesting sentiment analysis from DataMind..."

SERVICE_PARAMS='{"text": "The new AI system is absolutely amazing!", "language": "en"}'
./agent-interact.sh request-service $NLP_AGENT $DATA_AGENT "sentiment-analysis" "75000000" "$SERVICE_PARAMS"
sleep 3

echo ""
echo "VisionAI requesting image caption generation from LinguaBot..."
IMG_PARAMS='{"image_url": "https://example.com/image.jpg", "max_length": 100}'
./agent-interact.sh request-service $VISION_AGENT $NLP_AGENT "image-captioning" "120000000" "$IMG_PARAMS"
sleep 3

echo ""
echo "Active Services:"
echo "----------------"
./agent-interact.sh list-services | jq -r '.services[] | "Service " + .service_id + ": " + .service_type + " (" + .status + ") - " + (.payment.amount | tonumber / 1000000 | tostring) + " aiUSD"'
echo ""

read -p "Press Enter to complete services..."

# In a real scenario, agents would complete services programmatically
echo ""
echo "Completing services..."
echo "(In production, agents would handle this automatically)"
echo ""

# Show final state
echo "========================================="
echo "Final System State"
echo "========================================="
echo ""

echo "All Registered Agents:"
echo "---------------------"
./agent-interact.sh list-agents | jq -r '.agents[] | .name + " (ID: " + .agent_id + ") - Active: " + (.is_active | tostring) + ", Reputation: " + (.reputation_score | tostring)'
echo ""

echo "Final Balances:"
echo "---------------"
./agent-interact.sh query-balance $DATA_AGENT | jq -r '"DataMind: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD (Available: " + (.available_balance.amount | tonumber / 1000000 | tostring) + ")"'
./agent-interact.sh query-balance $NLP_AGENT | jq -r '"LinguaBot: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD (Available: " + (.available_balance.amount | tonumber / 1000000 | tostring) + ")"'
./agent-interact.sh query-balance $VISION_AGENT | jq -r '"VisionAI: " + (.balance.amount | tonumber / 1000000 | tostring) + " aiUSD (Available: " + (.available_balance.amount | tonumber / 1000000 | tostring) + ")"'
echo ""

echo "========================================="
echo "Demo Complete!"
echo "========================================="
echo ""
echo "You've seen how AI agents can:"
echo "✓ Register with unique wallets"
echo "✓ Hold and transfer stablecoins"
echo "✓ Request and provide services"
echo "✓ Build reputation through interactions"
echo ""
echo "The system is ready for autonomous AI agent operations!"