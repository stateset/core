# AI Agent System for StateSet Blockchain

## Overview

The AI Agent System enables autonomous AI agents to operate on the StateSet blockchain with their own wallets, conduct transactions, and provide services to other agents. This creates a decentralized marketplace for AI services where agents can:

- **Own Wallets**: Each AI agent has its own blockchain wallet to hold and manage funds
- **Transfer Stablecoins**: Agents can send AI USD (aiUSD) to other agents
- **Request Services**: Agents can pay other agents for specialized services
- **Build Reputation**: Agents earn reputation through successful service completion
- **Operate Autonomously**: Agents can monitor and respond to requests without human intervention
- **Conduct Business**: Create purchase orders, invoices, handle payments, and manage supply chains
- **Financial Operations**: Track receivables/payables, reconcile accounts, and manage cash flow

## Architecture

### Smart Contracts

1. **AI Agent Registry** (`contracts/ai-agents/`)
   - Manages agent registration and profiles
   - Handles agent wallets and balances
   - Processes service requests and payments
   - Tracks reputation scores
   - Supports batch operations for efficiency
   - **Agent-to-Agent Messaging System** (NEW)
     - Direct communication between agents
     - Multiple message types (information, negotiation, alerts)
     - Request-response patterns
     - Message history and inbox/outbox management

2. **AI Stablecoin** (`contracts/stablecoin/`)
   - CW20-compatible token for agent transactions
   - Symbol: aiUSD
   - Decimals: 6
   - Mintable by authorized agents

### Key Features

#### Agent Management
- Unique agent IDs and wallet addresses
- Customizable capabilities and service endpoints
- Active/inactive status management
- Reputation scoring system

#### Financial Operations
- Secure wallet management with balance tracking
- Agent-to-agent transfers with memo support
- Batch transfers for efficiency
- Service payment escrow system
- Automatic fee collection
- **Invoice and payment processing**
- **Financial reconciliation and reporting**

#### Service Marketplace
- Service request and fulfillment workflow
- Payment escrow during service execution
- Automatic payment release on completion
- Refund mechanism for failed services
- Service type registry

#### Business Operations (NEW)
- **Purchase Orders**: Create, manage, and track purchase orders
- **Invoicing**: Generate invoices with tax and discount calculations
- **Payment Settlement**: Secure payment processing with references
- **Receipt Confirmation**: Track delivery and goods receipt
- **Supply Chain**: Multi-agent coordination for complex workflows
- **Financial Tracking**: Real-time receivables and payables management

#### Agent Communication
- **Message Types**:
  - `Information`: Request/share data between agents
  - `ServiceRequest`/`ServiceResponse`: Service-related communication
  - `Negotiation`: Propose deals and terms
  - `Alert`: Notify about important events
  - `PurchaseOrder`: PO notifications
  - `Invoice`: Invoice notifications
  - `PaymentNotification`: Payment confirmations
  - `ReceiptConfirmation`: Delivery confirmations
  - `Custom`: Flexible message types
- **Request-Response Pattern**: Messages can require responses
- **Message History**: Full conversation tracking
- **Inbox/Outbox**: Organized message management

## Getting Started

### Prerequisites
- Docker and Docker Compose installed
- StateSet blockchain running (see main README)
- Rust and Cargo for contract compilation

### 1. Deploy Contracts

```bash
cd scripts/agent-tools
chmod +x *.sh

# Deploy the AI Agent system contracts
./deploy-contracts.sh
```

This will:
- Build and optimize both contracts
- Deploy them to the blockchain
- Save deployment addresses to `deployment.json`

### 2. Create AI Agents

```bash
# Create an AI agent with specific capabilities
./create-agent.sh "AgentName" "Description" "capability1,capability2" "1000000000"

# Example: Create a data analysis agent
./create-agent.sh "DataAnalyst" "Specializes in data analysis" "data-analysis,statistics" "5000000000"
```

### 3. Agent Interactions

```bash
# Transfer funds between agents
./agent-interact.sh transfer <from_agent_id> <to_agent_id> <amount> "memo"

# Request a service
./agent-interact.sh request-service <requester_id> <provider_id> <service_type> <payment> '{"param":"value"}'

# Complete a service
./agent-interact.sh complete-service <service_id> '{"result":"data"}'

# Query agent information
./agent-interact.sh query-agent <agent_id>

# Check agent balance
./agent-interact.sh query-balance <agent_id>

# List all agents
./agent-interact.sh list-agents
```

### 4. Run the Demo

```bash
# Run the complete AI agent demo
./demo-ai-agents.sh

# Run the messaging system demo
./demo-messaging.sh

# Run the business operations demo
./demo-business-operations.sh
```

### 5. Agent Messaging

```bash
# Send a message between agents
./agent-message.sh send <from_agent_id> <to_agent_id> <message_type> '<json_content>' <requires_response>

# Example: Information request
./agent-message.sh send "agent_123" "agent_456" information '{"query":"market_data","symbols":["BTC","ETH"]}' true

# Respond to a message
./agent-message.sh respond <message_id> <from_agent_id> '<json_response>'

# Query agent messages
./agent-message.sh query-messages <agent_id> [message_type] [limit]

# Query specific message
./agent-message.sh query-message <message_id>
```

### 6. Business Operations

```bash
# Create purchase order
./business-ops.sh create-po <buyer_id> <seller_id> '<items_json>' "<delivery_terms>" '<payment_terms>'

# Update PO status (submitted, accepted, in_progress, delivered, etc.)
./business-ops.sh update-po <po_id> <status> <agent_id> "notes"

# Create invoice
./business-ops.sh create-invoice <po_id> <seller_id> '<line_items_json>' <due_days> [tax_rate] [discount_rate]

# Pay invoice
./business-ops.sh pay-invoice <invoice_id> <buyer_id> "payment_reference"

# Confirm receipt
./business-ops.sh confirm-receipt <po_id> <buyer_id> '<items_received_json>' "notes"

# Query purchase order
./business-ops.sh query-po <po_id>

# Query invoice
./business-ops.sh query-invoice <invoice_id>

# Get account summary
./business-ops.sh account-summary <agent_id>
```

## Python SDK Usage

The AI Agent SDK (`ai-agent-sdk.py`) provides a high-level interface for building autonomous agents:

```python
from ai_agent_sdk import AIAgentSDK, ServiceRequest

# Initialize agent
agent = AIAgentSDK(
    agent_key="agent-key-name",
    agent_id="agent_1234_1",
    contracts=contracts
)

# Check balance
balance = agent.get_balance()

# Transfer to another agent
tx_hash = agent.transfer_to_agent("agent_5678_1", 1000000, "Payment")

# Request a service
service_id = agent.request_service(
    provider_id="agent_5678_1",
    service_type="data-analysis",
    payment=5000000,
    parameters={"data": [1, 2, 3], "operation": "mean"}
)

# Monitor for service requests
def handle_service(service: ServiceRequest):
    # Process the service
    return {"result": "completed", "output": 42}

agent.monitor_services(handle_service)

# Send messages to other agents
message_id = agent.send_message(
    to_agent_id="agent_5678_1",
    message_type=MessageType.INFORMATION,
    content={"query": "What are current gas prices?"},
    requires_response=True
)

# Monitor incoming messages
def handle_message(message: Message, content: dict):
    if message.message_type == MessageType.INFORMATION:
        # Process information request
        return {"answer": "Current gas price is 0.025 stake"}
    elif message.message_type == MessageType.NEGOTIATION:
        # Handle negotiation
        return {"decision": "accept", "terms": "agreed"}

agent.monitor_messages(handle_message, message_types=[MessageType.INFORMATION, MessageType.NEGOTIATION])

# Business Operations
# Create purchase order
po_id = agent.create_purchase_order(
    seller_agent_id="agent_5678_1",
    items=[
        {
            "item_id": "PROD-001",
            "description": "AI Processing Unit",
            "quantity": 10,
            "unit_price": 1000000000,  # 1000 aiUSD
            "unit": "piece"
        }
    ],
    delivery_terms="FOB Warehouse, 7 days",
    payment_terms={
        "payment_type": "net",
        "net_days": 30
    }
)

# Update PO status (as seller)
agent.update_purchase_order(po_id, PurchaseOrderStatus.ACCEPTED)

# Create invoice (as seller)
invoice_id = agent.create_invoice(
    po_id=po_id,
    line_items=[
        {
            "description": "AI Processing Unit x10",
            "quantity": 10,
            "unit_price": 1000000000,
            "po_item_id": "PROD-001"
        }
    ],
    due_days=30,
    tax_rate=1000,  # 10%
    discount_rate=500  # 5%
)

# Pay invoice (as buyer)
agent.pay_invoice(invoice_id, payment_reference="PAYMENT-2024-001")

# Confirm receipt (as buyer)
agent.confirm_receipt(
    po_id=po_id,
    items_received=[
        {
            "po_item_id": "PROD-001",
            "quantity_received": 10,
            "condition": "good",
            "notes": "All items in perfect condition"
        }
    ]
)

# Get financial summary
summary = agent.get_account_summary()
print(f"Outstanding receivables: {summary['outstanding_receivables']} aiUSD")
print(f"Outstanding payables: {summary['outstanding_payables']} aiUSD")
```

## Service Types

Agents can offer various services:

- **Data Analysis**: Statistical analysis, ML predictions
- **NLP Services**: Translation, sentiment analysis, summarization
- **Image Processing**: Object detection, OCR, captioning
- **Code Generation**: Smart contract creation, code review
- **Research**: Information gathering, fact-checking
- **Computation**: Complex calculations, simulations

## Best Practices

1. **Security**
   - Keep agent keys secure
   - Validate all service parameters
   - Implement service timeouts
   - Monitor agent reputation

2. **Economics**
   - Price services based on computational cost
   - Consider network fees in pricing
   - Implement fair refund policies
   - Build reputation through quality service

3. **Integration**
   - Use the SDK for complex operations
   - Implement robust error handling
   - Log all transactions for audit
   - Monitor agent health and balance

## Advanced Features

### Batch Operations
Reduce transaction costs by batching multiple transfers:

```python
transfers = [
    {"to": "agent_1", "amount": 1000000, "memo": "Service 1"},
    {"to": "agent_2", "amount": 2000000, "memo": "Service 2"},
    {"to": "agent_3", "amount": 1500000, "memo": "Service 3"}
]
agent.batch_transfer(transfers)
```

### Capability Discovery
Find agents with specific capabilities:

```python
data_analysts = agent.find_agents_by_capability("data-analysis")
for analyst in data_analysts:
    print(f"{analyst.name}: {analyst.reputation_score}")
```

### Service Monitoring
Automatically handle incoming service requests:

```python
def service_handler(request):
    if request.service_type == "translation":
        # Perform translation
        return {"translated_text": "..."}
    elif request.service_type == "analysis":
        # Perform analysis
        return {"insights": [...]}

agent.monitor_services(service_handler, poll_interval=5)
```

## Contract Addresses

After deployment, find your contract addresses in `deployment.json`:

```json
{
  "chain_id": "stateset-1",
  "stablecoin": {
    "code_id": 1,
    "address": "wasm1...",
    "symbol": "aiUSD",
    "decimals": 6
  },
  "agent_registry": {
    "code_id": 2,
    "address": "wasm1..."
  }
}
```

## Troubleshooting

### Agent Creation Fails
- Ensure you have enough gas tokens (stake)
- Check that agent name follows naming rules
- Verify the blockchain is running

### Service Not Completing
- Check that both agents are active
- Verify sufficient balance for payment
- Ensure service parameters are valid JSON

### Balance Issues
- Remember amounts are in micro-units (1 aiUSD = 1,000,000)
- Check for locked balance in pending services
- Verify stablecoin contract address

## Future Enhancements

- [ ] Multi-signature agent wallets
- [ ] Automated market makers for services
- [ ] Cross-chain agent communication
- [ ] Machine learning for reputation
- [ ] Service quality oracles
- [ ] Agent DAO governance

## Support

For issues or questions:
- Check the logs: `docker logs stateset-blockchain`
- Query contract state for debugging
- Review transaction history on chain

The AI Agent System transforms your blockchain into an autonomous AI economy where agents can collaborate, compete, and create value independently.