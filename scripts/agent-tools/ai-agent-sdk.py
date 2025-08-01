#!/usr/bin/env python3
"""
AI Agent SDK for StateSet Blockchain
Enables AI agents to interact with the blockchain autonomously
"""

import json
import subprocess
import time
from typing import Dict, List, Optional, Any
from dataclasses import dataclass
from enum import Enum

class ServiceStatus(Enum):
    PENDING = "pending"
    IN_PROGRESS = "in_progress"
    COMPLETED = "completed"
    FAILED = "failed"
    REFUNDED = "refunded"

class MessageType(Enum):
    SERVICE_REQUEST = "service_request"
    SERVICE_RESPONSE = "service_response"
    NEGOTIATION = "negotiation"
    INFORMATION = "information"
    ALERT = "alert"
    PURCHASE_ORDER = "purchase_order"
    INVOICE = "invoice"
    PAYMENT_NOTIFICATION = "payment_notification"
    RECEIPT_CONFIRMATION = "receipt_confirmation"
    CUSTOM = "custom"

class PaymentType(Enum):
    IMMEDIATE = "immediate"
    NET = "net"
    DEPOSIT = "deposit"
    MILESTONE = "milestone"

class PurchaseOrderStatus(Enum):
    DRAFT = "draft"
    SUBMITTED = "submitted"
    ACCEPTED = "accepted"
    REJECTED = "rejected"
    IN_PROGRESS = "in_progress"
    DELIVERED = "delivered"
    COMPLETED = "completed"
    CANCELLED = "cancelled"

class ItemCondition(Enum):
    GOOD = "good"
    DAMAGED = "damaged"
    MISSING = "missing"
    WRONG = "wrong"

@dataclass
class AgentInfo:
    agent_id: str
    name: str
    address: str
    balance: int
    capabilities: List[str]
    is_active: bool
    reputation_score: int

@dataclass
class ServiceRequest:
    service_id: str
    requester: str
    provider: str
    service_type: str
    payment: int
    status: ServiceStatus
    parameters: Dict[str, Any]
    result: Optional[str] = None

@dataclass
class Message:
    message_id: str
    from_agent_id: str
    to_agent_id: str
    message_type: MessageType
    content: str
    requires_response: bool
    timestamp: int
    response: Optional[Dict[str, Any]] = None

@dataclass
class PurchaseOrderItem:
    item_id: str
    description: str
    quantity: int
    unit_price: int
    unit: str

@dataclass
class PaymentTerms:
    payment_type: PaymentType
    deposit_percentage: Optional[int]
    net_days: int

@dataclass
class PurchaseOrder:
    po_id: str
    buyer_agent_id: str
    seller_agent_id: str
    items: List[PurchaseOrderItem]
    total_amount: int
    status: PurchaseOrderStatus
    created_at: int
    updated_at: int
    delivery_terms: str
    payment_terms: PaymentTerms
    invoice_id: Optional[str] = None

@dataclass
class InvoiceLineItem:
    description: str
    quantity: int
    unit_price: int
    po_item_id: Optional[str] = None

@dataclass
class Invoice:
    invoice_id: str
    po_id: str
    seller_agent_id: str
    buyer_agent_id: str
    line_items: List[InvoiceLineItem]
    subtotal: int
    tax_amount: int
    discount_amount: int
    total_amount: int
    paid: bool
    paid_at: Optional[int]
    created_at: int
    due_date: int
    payment_reference: Optional[str] = None

@dataclass
class ItemReceipt:
    po_item_id: str
    quantity_received: int
    condition: ItemCondition
    notes: Optional[str] = None

class AIAgentSDK:
    """SDK for AI Agents to interact with StateSet blockchain"""
    
    def __init__(self, agent_key: str, agent_id: str, contracts: Dict[str, str]):
        self.agent_key = agent_key
        self.agent_id = agent_id
        self.contracts = contracts
        self.chain_id = "stateset-1"
        self.keyring = "--keyring-backend test"
        self.gas_flags = "--gas auto --gas-adjustment 1.5 --gas-prices 0.025stake"
        
    def _execute_contract(self, contract: str, msg: Dict[str, Any], funds: Optional[str] = None) -> Dict:
        """Execute a contract transaction"""
        msg_json = json.dumps(msg)
        
        cmd = [
            "docker", "exec", "stateset-blockchain",
            "wasmd", "tx", "wasm", "execute",
            contract, msg_json,
            "--from", self.agent_key,
            "--chain-id", self.chain_id,
            "--output", "json",
            "-y"
        ]
        
        # Add keyring and gas flags
        cmd.extend(self.keyring.split())
        cmd.extend(self.gas_flags.split())
        
        # Add funds if specified
        if funds:
            cmd.extend(["--amount", funds])
        
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            raise Exception(f"Contract execution failed: {result.stderr}")
            
        return json.loads(result.stdout)
    
    def _query_contract(self, contract: str, msg: Dict[str, Any]) -> Dict:
        """Query a contract"""
        msg_json = json.dumps(msg)
        
        cmd = [
            "docker", "exec", "stateset-blockchain",
            "wasmd", "query", "wasm", "contract-state", "smart",
            contract, msg_json,
            "--output", "json"
        ]
        
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            raise Exception(f"Contract query failed: {result.stderr}")
            
        return json.loads(result.stdout)["data"]
    
    def get_balance(self) -> int:
        """Get agent's current balance in aiUSD"""
        msg = {
            "agent_balance": {
                "agent_id": self.agent_id
            }
        }
        result = self._query_contract(self.contracts["agent_registry"], msg)
        return int(result["balance"]["amount"])
    
    def get_info(self) -> AgentInfo:
        """Get agent's full information"""
        msg = {
            "agent": {
                "agent_id": self.agent_id
            }
        }
        result = self._query_contract(self.contracts["agent_registry"], msg)
        
        return AgentInfo(
            agent_id=result["agent_id"],
            name=result["name"],
            address=result["wallet_address"],
            balance=int(result["balance"]["amount"]),
            capabilities=result["capabilities"],
            is_active=result["is_active"],
            reputation_score=result["reputation_score"]
        )
    
    def transfer_to_agent(self, to_agent_id: str, amount: int, memo: str = "") -> str:
        """Transfer aiUSD to another agent"""
        msg = {
            "agent_transfer": {
                "from_agent_id": self.agent_id,
                "to_agent_id": to_agent_id,
                "amount": {
                    "denom": "ibc/aiUSD",
                    "amount": str(amount)
                },
                "memo": memo
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def batch_transfer(self, transfers: List[Dict[str, Any]]) -> str:
        """Execute multiple transfers in one transaction"""
        transfer_list = []
        for t in transfers:
            transfer_list.append({
                "to_agent_id": t["to"],
                "amount": {
                    "denom": "ibc/aiUSD",
                    "amount": str(t["amount"])
                },
                "memo": t.get("memo", "")
            })
        
        msg = {
            "batch_agent_transfer": {
                "from_agent_id": self.agent_id,
                "transfers": transfer_list
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def request_service(self, provider_id: str, service_type: str, 
                       payment: int, parameters: Dict[str, Any]) -> str:
        """Request a service from another agent"""
        msg = {
            "request_service": {
                "requester_agent_id": self.agent_id,
                "provider_agent_id": provider_id,
                "service_type": service_type,
                "payment": {
                    "denom": "ibc/aiUSD",
                    "amount": str(payment)
                },
                "parameters": json.dumps(parameters)
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        
        # Extract service ID from events
        for event in result.get("logs", [{}])[0].get("events", []):
            if event["type"] == "wasm":
                for attr in event["attributes"]:
                    if attr["key"] == "service_id":
                        import base64
                        return base64.b64decode(attr["value"]).decode()
        
        raise Exception("Service ID not found in transaction result")
    
    def complete_service(self, service_id: str, result: Dict[str, Any]) -> str:
        """Complete a service request"""
        msg = {
            "complete_service": {
                "service_id": service_id,
                "result": json.dumps(result)
            }
        }
        
        tx_result = self._execute_contract(self.contracts["agent_registry"], msg)
        return tx_result["txhash"]
    
    def get_pending_services(self) -> List[ServiceRequest]:
        """Get all pending service requests for this agent"""
        msg = {
            "list_services": {
                "agent_id": self.agent_id,
                "status": "pending",
                "limit": 50
            }
        }
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        services = []
        
        for service in result.get("services", []):
            services.append(ServiceRequest(
                service_id=service["service_id"],
                requester=service["requester_agent_id"],
                provider=service["provider_agent_id"],
                service_type=service["service_type"],
                payment=int(service["payment"]["amount"]),
                status=ServiceStatus(service["status"].lower()),
                parameters=json.loads(service["parameters"]),
                result=service.get("result")
            ))
        
        return services
    
    def find_agents_by_capability(self, capability: str, limit: int = 10) -> List[AgentInfo]:
        """Find agents with a specific capability"""
        msg = {
            "agents_by_capability": {
                "capability": capability,
                "limit": limit
            }
        }
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        agents = []
        
        for agent in result.get("agents", []):
            # Get full agent info
            agent_info = self._query_agent(agent["agent_id"])
            agents.append(agent_info)
        
        return agents
    
    def _query_agent(self, agent_id: str) -> AgentInfo:
        """Query information about any agent"""
        msg = {
            "agent": {
                "agent_id": agent_id
            }
        }
        result = self._query_contract(self.contracts["agent_registry"], msg)
        
        return AgentInfo(
            agent_id=result["agent_id"],
            name=result["name"],
            address=result["wallet_address"],
            balance=int(result["balance"]["amount"]),
            capabilities=result["capabilities"],
            is_active=result["is_active"],
            reputation_score=result["reputation_score"]
        )
    
    def monitor_services(self, callback, poll_interval: int = 5):
        """Monitor for new service requests and handle them with callback"""
        processed_services = set()
        
        while True:
            try:
                pending_services = self.get_pending_services()
                
                for service in pending_services:
                    if service.service_id not in processed_services:
                        # Only process services where we are the provider
                        if service.provider == self.agent_id:
                            processed_services.add(service.service_id)
                            
                            # Call the callback to handle the service
                            try:
                                result = callback(service)
                                
                                # Complete the service
                                self.complete_service(service.service_id, result)
                                print(f"Completed service {service.service_id}")
                                
                            except Exception as e:
                                print(f"Error processing service {service.service_id}: {e}")
                                # Could implement service refund here
                
                time.sleep(poll_interval)
                
            except KeyboardInterrupt:
                print("Service monitoring stopped")
                break
            except Exception as e:
                print(f"Error in service monitoring: {e}")
                time.sleep(poll_interval)
    
    def send_message(self, to_agent_id: str, message_type: MessageType, 
                    content: Dict[str, Any], requires_response: bool = False) -> str:
        """Send a message to another agent"""
        # Convert message type enum to string
        if isinstance(message_type, MessageType):
            if message_type == MessageType.CUSTOM:
                msg_type = {"custom": content.get("custom_type", "general")}
            else:
                msg_type = message_type.value
        else:
            msg_type = str(message_type)
        
        msg = {
            "send_message": {
                "from_agent_id": self.agent_id,
                "to_agent_id": to_agent_id,
                "message_type": msg_type,
                "content": json.dumps(content),
                "requires_response": requires_response
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        
        # Extract message ID from events
        for event in result.get("logs", [{}])[0].get("events", []):
            if event["type"] == "wasm":
                for attr in event["attributes"]:
                    if attr["key"] == "message_id":
                        import base64
                        return base64.b64decode(attr["value"]).decode()
        
        raise Exception("Message ID not found in transaction result")
    
    def respond_to_message(self, message_id: str, response_content: Dict[str, Any]) -> str:
        """Respond to a message that requires a response"""
        msg = {
            "respond_to_message": {
                "message_id": message_id,
                "from_agent_id": self.agent_id,
                "response_content": json.dumps(response_content)
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def get_messages(self, message_type: Optional[MessageType] = None, 
                    limit: int = 20) -> List[Message]:
        """Get messages for this agent"""
        query_msg = {
            "agent_messages": {
                "agent_id": self.agent_id,
                "limit": limit
            }
        }
        
        if message_type:
            if isinstance(message_type, MessageType):
                if message_type == MessageType.CUSTOM:
                    query_msg["agent_messages"]["message_type"] = {"custom": "general"}
                else:
                    query_msg["agent_messages"]["message_type"] = message_type.value
            else:
                query_msg["agent_messages"]["message_type"] = str(message_type)
        
        result = self._query_contract(self.contracts["agent_registry"], query_msg)
        messages = []
        
        for msg in result.get("messages", []):
            # Parse message type
            msg_type_str = msg["message_type"]
            if isinstance(msg_type_str, dict) and "custom" in msg_type_str:
                msg_type = MessageType.CUSTOM
            else:
                msg_type = MessageType(msg_type_str) if msg_type_str in [e.value for e in MessageType] else MessageType.CUSTOM
            
            messages.append(Message(
                message_id=msg["message_id"],
                from_agent_id=msg["from_agent_id"],
                to_agent_id=msg["to_agent_id"],
                message_type=msg_type,
                content=msg["content"],
                requires_response=msg["requires_response"],
                timestamp=msg["timestamp"],
                response=msg.get("response")
            ))
        
        return messages
    
    def get_message(self, message_id: str) -> Message:
        """Get a specific message by ID"""
        msg = {
            "message": {
                "message_id": message_id
            }
        }
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        msg_info = result["message"]
        
        # Parse message type
        msg_type_str = msg_info["message_type"]
        if isinstance(msg_type_str, dict) and "custom" in msg_type_str:
            msg_type = MessageType.CUSTOM
        else:
            msg_type = MessageType(msg_type_str) if msg_type_str in [e.value for e in MessageType] else MessageType.CUSTOM
        
        return Message(
            message_id=msg_info["message_id"],
            from_agent_id=msg_info["from_agent_id"],
            to_agent_id=msg_info["to_agent_id"],
            message_type=msg_type,
            content=msg_info["content"],
            requires_response=msg_info["requires_response"],
            timestamp=msg_info["timestamp"],
            response=msg_info.get("response")
        )
    
    def monitor_messages(self, callback, message_types: Optional[List[MessageType]] = None,
                        poll_interval: int = 5):
        """Monitor for new messages and handle them with callback"""
        processed_messages = set()
        
        while True:
            try:
                messages = self.get_messages(limit=50)
                
                for message in messages:
                    if message.message_id not in processed_messages:
                        # Filter by message types if specified
                        if message_types and message.message_type not in message_types:
                            continue
                        
                        # Only process messages where we are the recipient
                        if message.to_agent_id == self.agent_id:
                            processed_messages.add(message.message_id)
                            
                            try:
                                # Parse content as JSON
                                content = json.loads(message.content)
                                
                                # Call the callback
                                response = callback(message, content)
                                
                                # If message requires response and callback returned something
                                if message.requires_response and response is not None:
                                    self.respond_to_message(message.message_id, response)
                                    print(f"Responded to message {message.message_id}")
                                
                            except Exception as e:
                                print(f"Error processing message {message.message_id}: {e}")
                
                time.sleep(poll_interval)
                
            except KeyboardInterrupt:
                print("Message monitoring stopped")
                break
            except Exception as e:
                print(f"Error in message monitoring: {e}")
                time.sleep(poll_interval)
    
    # Business Operations Functions
    
    def create_purchase_order(self, seller_agent_id: str, items: List[Dict[str, Any]], 
                            delivery_terms: str, payment_terms: Dict[str, Any],
                            metadata: Optional[str] = None) -> str:
        """Create a purchase order"""
        # Convert items to proper format
        po_items = []
        for item in items:
            po_items.append({
                "item_id": item["item_id"],
                "description": item["description"],
                "quantity": str(item["quantity"]),
                "unit_price": str(item["unit_price"]),
                "unit": item["unit"]
            })
        
        # Convert payment terms
        payment_terms_msg = {
            "payment_type": payment_terms["payment_type"],
            "deposit_percentage": payment_terms.get("deposit_percentage"),
            "net_days": str(payment_terms["net_days"])
        }
        
        msg = {
            "create_purchase_order": {
                "buyer_agent_id": self.agent_id,
                "seller_agent_id": seller_agent_id,
                "items": po_items,
                "delivery_terms": delivery_terms,
                "payment_terms": payment_terms_msg,
                "metadata": metadata
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        
        # Extract PO ID from events
        for event in result.get("logs", [{}])[0].get("events", []):
            if event["type"] == "wasm":
                for attr in event["attributes"]:
                    if attr["key"] == "po_id":
                        import base64
                        return base64.b64decode(attr["value"]).decode()
        
        raise Exception("PO ID not found in transaction result")
    
    def update_purchase_order(self, po_id: str, status: PurchaseOrderStatus, 
                            notes: Optional[str] = None) -> str:
        """Update purchase order status"""
        msg = {
            "update_purchase_order": {
                "po_id": po_id,
                "status": status.value,
                "updater_agent_id": self.agent_id,
                "notes": notes
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def create_invoice(self, po_id: str, line_items: List[Dict[str, Any]],
                      due_days: int = 30, tax_rate: Optional[int] = None,
                      discount_rate: Optional[int] = None, metadata: Optional[str] = None) -> str:
        """Create an invoice for a purchase order"""
        # Convert line items
        invoice_items = []
        for item in line_items:
            invoice_items.append({
                "description": item["description"],
                "quantity": str(item["quantity"]),
                "unit_price": str(item["unit_price"]),
                "po_item_id": item.get("po_item_id")
            })
        
        # Calculate due date (current time + due_days)
        import time
        due_date = int(time.time()) + (due_days * 86400)
        
        msg = {
            "create_invoice": {
                "po_id": po_id,
                "seller_agent_id": self.agent_id,
                "line_items": invoice_items,
                "tax_rate": tax_rate,
                "discount_rate": discount_rate,
                "due_date": str(due_date),
                "metadata": metadata
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        
        # Extract invoice ID from events
        for event in result.get("logs", [{}])[0].get("events", []):
            if event["type"] == "wasm":
                for attr in event["attributes"]:
                    if attr["key"] == "invoice_id":
                        import base64
                        return base64.b64decode(attr["value"]).decode()
        
        raise Exception("Invoice ID not found in transaction result")
    
    def pay_invoice(self, invoice_id: str, payment_reference: Optional[str] = None) -> str:
        """Pay an invoice"""
        msg = {
            "pay_invoice": {
                "invoice_id": invoice_id,
                "buyer_agent_id": self.agent_id,
                "payment_reference": payment_reference
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def confirm_receipt(self, po_id: str, items_received: List[Dict[str, Any]], 
                       notes: Optional[str] = None) -> str:
        """Confirm receipt of items from a purchase order"""
        # Convert items received
        receipt_items = []
        for item in items_received:
            receipt_items.append({
                "po_item_id": item["po_item_id"],
                "quantity_received": str(item["quantity_received"]),
                "condition": item["condition"],
                "notes": item.get("notes")
            })
        
        msg = {
            "confirm_receipt": {
                "po_id": po_id,
                "buyer_agent_id": self.agent_id,
                "items_received": receipt_items,
                "notes": notes
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]
    
    def get_purchase_order(self, po_id: str) -> PurchaseOrder:
        """Get purchase order details"""
        msg = {
            "purchase_order": {
                "po_id": po_id
            }
        }
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        po = result["purchase_order"]
        
        # Convert items
        items = []
        for item in po["items"]:
            items.append(PurchaseOrderItem(
                item_id=item["item_id"],
                description=item["description"],
                quantity=int(item["quantity"]),
                unit_price=int(item["unit_price"]),
                unit=item["unit"]
            ))
        
        # Convert payment terms
        pt = po["payment_terms"]
        payment_terms = PaymentTerms(
            payment_type=PaymentType(pt["payment_type"]),
            deposit_percentage=pt.get("deposit_percentage"),
            net_days=int(pt["net_days"])
        )
        
        return PurchaseOrder(
            po_id=po["po_id"],
            buyer_agent_id=po["buyer_agent_id"],
            seller_agent_id=po["seller_agent_id"],
            items=items,
            total_amount=int(po["total_amount"]),
            status=PurchaseOrderStatus(po["status"].lower()),
            created_at=po["created_at"],
            updated_at=po["updated_at"],
            delivery_terms=po["delivery_terms"],
            payment_terms=payment_terms,
            invoice_id=po.get("invoice_id")
        )
    
    def get_invoice(self, invoice_id: str) -> Invoice:
        """Get invoice details"""
        msg = {
            "invoice": {
                "invoice_id": invoice_id
            }
        }
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        inv = result["invoice"]
        
        # Convert line items
        line_items = []
        for item in inv["line_items"]:
            line_items.append(InvoiceLineItem(
                description=item["description"],
                quantity=int(item["quantity"]),
                unit_price=int(item["unit_price"]),
                po_item_id=item.get("po_item_id")
            ))
        
        return Invoice(
            invoice_id=inv["invoice_id"],
            po_id=inv["po_id"],
            seller_agent_id=inv["seller_agent_id"],
            buyer_agent_id=inv["buyer_agent_id"],
            line_items=line_items,
            subtotal=int(inv["subtotal"]),
            tax_amount=int(inv["tax_amount"]),
            discount_amount=int(inv["discount_amount"]),
            total_amount=int(inv["total_amount"]),
            paid=inv["paid"],
            paid_at=inv.get("paid_at"),
            created_at=inv["created_at"],
            due_date=inv["due_date"],
            payment_reference=inv.get("payment_reference")
        )
    
    def get_account_summary(self, period_start: Optional[int] = None, 
                          period_end: Optional[int] = None) -> Dict[str, Any]:
        """Get financial summary for the agent"""
        msg = {
            "account_summary": {
                "agent_id": self.agent_id
            }
        }
        
        if period_start:
            msg["account_summary"]["period_start"] = str(period_start)
        if period_end:
            msg["account_summary"]["period_end"] = str(period_end)
        
        result = self._query_contract(self.contracts["agent_registry"], msg)
        
        return {
            "total_sales": int(result["total_sales"]),
            "total_purchases": int(result["total_purchases"]),
            "outstanding_receivables": int(result["outstanding_receivables"]),
            "outstanding_payables": int(result["outstanding_payables"]),
            "completed_orders": result["completed_orders"],
            "pending_orders": result["pending_orders"]
        }
    
    def reconcile_accounts(self, period_start: int, period_end: int) -> str:
        """Reconcile accounts for a specific period"""
        msg = {
            "reconcile_accounts": {
                "agent_id": self.agent_id,
                "period_start": str(period_start),
                "period_end": str(period_end)
            }
        }
        
        result = self._execute_contract(self.contracts["agent_registry"], msg)
        return result["txhash"]


# Example usage
if __name__ == "__main__":
    # Load deployment info
    with open("deployment.json", "r") as f:
        deployment = json.load(f)
    
    contracts = {
        "agent_registry": deployment["agent_registry"]["address"],
        "stablecoin": deployment["stablecoin"]["address"]
    }
    
    # Example: Create an AI agent that provides data analysis services
    agent = AIAgentSDK(
        agent_key="agent-data-analyst",
        agent_id="agent_1234_1",
        contracts=contracts
    )
    
    # Check balance
    balance = agent.get_balance()
    print(f"Current balance: {balance} aiUSD")
    
    # Find agents that need data analysis
    analysts = agent.find_agents_by_capability("data-analysis")
    print(f"Found {len(analysts)} data analysis agents")
    
    # Define service handler
    def handle_data_analysis(service: ServiceRequest) -> Dict[str, Any]:
        """Handle incoming data analysis requests"""
        params = service.parameters
        
        # Simulate data analysis
        if params.get("type") == "statistical":
            return {
                "mean": 42.5,
                "median": 40.0,
                "std_dev": 5.2,
                "analysis": "Data shows normal distribution"
            }
        else:
            return {
                "result": "Analysis completed",
                "insights": ["Pattern detected", "Anomaly at index 5"]
            }
    
    # Start monitoring for service requests
    print("Starting service monitor...")
    agent.monitor_services(handle_data_analysis)