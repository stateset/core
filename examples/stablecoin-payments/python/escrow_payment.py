#!/usr/bin/env python3
"""
Escrow-Based Payment Example

Demonstrates secure escrow payments for e-commerce:
- Create escrow for purchases
- Hold funds until delivery
- Release to seller on confirmation
- Handle refunds

Usage: python escrow_payment.py
"""

import json
import subprocess
import os
import time
from typing import Optional, Dict, Any, List
from dataclasses import dataclass
from decimal import Decimal
from datetime import datetime, timedelta
from enum import Enum


# Configuration
CHAIN_ID = os.getenv("STATESET_CHAIN_ID", "stateset-1")
NODE = os.getenv("STATESET_NODE", "tcp://localhost:26657")
KEYRING = os.getenv("STATESET_KEYRING", "test")
GAS_PRICES = "0.025ustake"
GAS_LIMIT = "500000"
BINARY = os.getenv("STATESET_BINARY", "statesetd")

SSUSD_DENOM = "ssusd"
SSUSD_DECIMALS = 6


class EscrowStatus(Enum):
    ACTIVE = "active"
    RELEASED = "released"
    REFUNDED = "refunded"
    DISPUTED = "disputed"


@dataclass
class Escrow:
    """Escrow details"""
    escrow_id: str
    depositor: str
    recipient: str
    amount: int
    release_time: datetime
    conditions: str
    status: EscrowStatus


@dataclass
class TransactionResult:
    success: bool
    tx_hash: Optional[str] = None
    escrow_id: Optional[str] = None
    error: Optional[str] = None


class EscrowClient:
    """Client for escrow operations"""

    def __init__(self, key_name: str):
        self.key_name = key_name
        self.address = self._get_address()

    def _get_address(self) -> str:
        cmd = [
            BINARY, "keys", "show", self.key_name,
            "--keyring-backend", KEYRING,
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            raise Exception(f"Failed to get address: {result.stderr}")
        return json.loads(result.stdout)["address"]

    def _execute_tx(self, *args) -> TransactionResult:
        cmd = [
            BINARY, "tx", *args,
            "--from", self.key_name,
            "--chain-id", CHAIN_ID,
            "--node", NODE,
            "--keyring-backend", KEYRING,
            "--gas", GAS_LIMIT,
            "--gas-prices", GAS_PRICES,
            "--broadcast-mode", "sync",
            "--yes",
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)

        if result.returncode != 0:
            return TransactionResult(success=False, error=result.stderr)

        try:
            data = json.loads(result.stdout)
            return TransactionResult(
                success=True,
                tx_hash=data.get("txhash"),
                escrow_id=data.get("txhash", "")[:16]  # Simplified escrow ID
            )
        except json.JSONDecodeError:
            return TransactionResult(success=False, error="Failed to parse response")

    def _query(self, *args) -> Dict[str, Any]:
        cmd = [
            BINARY, "query", *args,
            "--chain-id", CHAIN_ID,
            "--node", NODE,
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            return {}
        return json.loads(result.stdout)

    def get_balance(self, denom: str = SSUSD_DENOM) -> int:
        result = self._query("bank", "balances", self.address)
        for balance in result.get("balances", []):
            if balance["denom"] == denom:
                return int(balance["amount"])
        return 0

    def create_escrow(
        self,
        recipient: str,
        amount: int,
        release_days: int = 7,
        conditions: Optional[Dict] = None
    ) -> TransactionResult:
        """Create an escrow payment"""
        release_time = datetime.now() + timedelta(days=release_days)
        conditions_json = json.dumps(conditions or {})

        escrow_data = json.dumps({
            "recipient": recipient,
            "amount": {"denom": SSUSD_DENOM, "amount": str(amount)},
            "release_time": release_time.isoformat(),
            "conditions": conditions_json
        })

        return self._execute_tx("settlement", "create-escrow", escrow_data)

    def release_escrow(self, escrow_id: str) -> TransactionResult:
        """Release escrow funds to recipient"""
        return self._execute_tx("settlement", "release-escrow", escrow_id)

    def refund_escrow(self, escrow_id: str) -> TransactionResult:
        """Refund escrow funds to depositor"""
        return self._execute_tx("settlement", "refund-escrow", escrow_id)

    def get_escrow(self, escrow_id: str) -> Optional[Escrow]:
        """Get escrow details"""
        result = self._query("settlement", "escrow", escrow_id)
        if not result:
            return None

        escrow_data = result.get("escrow", {})
        return Escrow(
            escrow_id=escrow_data.get("id", escrow_id),
            depositor=escrow_data.get("depositor", ""),
            recipient=escrow_data.get("recipient", ""),
            amount=int(escrow_data.get("amount", {}).get("amount", 0)),
            release_time=datetime.fromisoformat(escrow_data.get("release_time", datetime.now().isoformat())),
            conditions=escrow_data.get("conditions", ""),
            status=EscrowStatus(escrow_data.get("status", "active").lower())
        )


def format_amount(amount: int) -> str:
    return f"{Decimal(amount) / Decimal(10 ** SSUSD_DECIMALS):.2f}"


def to_base_units(amount: float) -> int:
    return int(amount * (10 ** SSUSD_DECIMALS))


def print_box(lines: List[str], title: str = "", width: int = 60):
    """Print a formatted box"""
    print("┌" + "─" * (width - 2) + "┐")
    if title:
        padding = (width - 2 - len(title)) // 2
        print(f"│{' ' * padding}{title}{' ' * (width - 2 - padding - len(title))}│")
        print("├" + "─" * (width - 2) + "┤")
    for line in lines:
        print(f"│ {line:<{width - 4}} │")
    print("└" + "─" * (width - 2) + "┘")


def print_escrow_status(escrow: Escrow):
    """Print formatted escrow status"""
    lines = [
        f"ID:        {escrow.escrow_id}",
        f"Status:    {escrow.status.value}",
        f"Amount:    {format_amount(escrow.amount)} ssUSD",
        f"Depositor: {escrow.depositor[:20]}...",
        f"Recipient: {escrow.recipient[:20]}...",
        f"Release:   {escrow.release_time.strftime('%Y-%m-%d %H:%M')}",
    ]
    print_box(lines, "ESCROW STATUS")


def demo_escrow_payment():
    """Demonstrate escrow payment flow"""
    print("═" * 60)
    print("  ssUSD Escrow Payment Example")
    print("  Secure Marketplace Transaction")
    print("═" * 60)
    print()

    # Setup participants
    print("1. Setting up participants...")
    buyer = EscrowClient("buyer")
    seller = EscrowClient("seller")
    print(f"   Buyer:  {buyer.address}")
    print(f"   Seller: {seller.address}")
    print()

    # Check initial balances
    print("2. Initial Balances:")
    buyer_balance = buyer.get_balance()
    seller_balance = seller.get_balance()
    print(f"   Buyer:  {format_amount(buyer_balance)} ssUSD")
    print(f"   Seller: {format_amount(seller_balance)} ssUSD")
    print()

    # Create escrow
    purchase_amount = 500.0
    print("3. Creating Escrow:")
    print(f"   Purchase Amount: {purchase_amount} ssUSD")
    print(f"   Escrow Duration: 7 days")
    print(f"   Conditions: Delivery confirmation required")
    print()

    print("   Submitting escrow transaction...")
    result = buyer.create_escrow(
        recipient=seller.address,
        amount=to_base_units(purchase_amount),
        release_days=7,
        conditions={
            "type": "delivery_confirmation",
            "order_id": "ORD-12345",
            "requires_tracking": True
        }
    )

    if result.success:
        print(f"   ✓ Escrow created!")
        print(f"   TX Hash:   {result.tx_hash}")
        print(f"   Escrow ID: {result.escrow_id}")
    else:
        print(f"   ✗ Failed: {result.error}")
        return
    print()

    # Check balances after escrow
    print("4. Balances After Escrow (funds locked):")
    buyer_after = buyer.get_balance()
    seller_after = seller.get_balance()
    print(f"   Buyer:  {format_amount(buyer_after)} ssUSD")
    print(f"   Seller: {format_amount(seller_after)} ssUSD")
    print(f"   Locked: {purchase_amount} ssUSD (in escrow)")
    print()

    # Simulate order fulfillment
    print("5. Order Fulfillment:")
    print("   → Seller ships product")
    print("   → Tracking: TRACK-12345-ABC")
    print("   → Buyer confirms delivery")
    print()

    # Release escrow
    print("6. Releasing Escrow to Seller:")
    release_result = buyer.release_escrow(result.escrow_id)

    if release_result.success:
        print(f"   ✓ Escrow released!")
        print(f"   TX Hash: {release_result.tx_hash}")
    else:
        print(f"   ✗ Release failed: {release_result.error}")
    print()

    # Final balances
    print("7. Final Balances:")
    buyer_final = buyer.get_balance()
    seller_final = seller.get_balance()
    print(f"   Buyer:  {format_amount(buyer_final)} ssUSD")
    print(f"   Seller: {format_amount(seller_final)} ssUSD")
    print()

    # Balance changes
    buyer_change = buyer_final - buyer_balance
    seller_change = seller_final - seller_balance
    print("   Balance Changes:")
    print(f"   Buyer:  {'+' if buyer_change >= 0 else ''}{format_amount(buyer_change)} ssUSD")
    print(f"   Seller: {'+' if seller_change >= 0 else ''}{format_amount(seller_change)} ssUSD")
    print()

    print("═" * 60)
    print("  Escrow Payment Completed Successfully!")
    print("═" * 60)


def demo_refund_flow():
    """Demonstrate escrow refund flow"""
    print()
    print("═" * 60)
    print("  BONUS: Escrow Refund Example")
    print("═" * 60)
    print()

    buyer = EscrowClient("buyer")
    seller_address = "stateset1seller..."

    # Create escrow
    print("Creating escrow for potential refund...")
    result = buyer.create_escrow(
        recipient=seller_address,
        amount=to_base_units(200.0),
        release_days=1,
        conditions={"type": "order_fulfillment"}
    )

    if result.success:
        print(f"Escrow ID: {result.escrow_id}")
        print()

        # Scenario: Order cancelled
        print("Scenario: Order cancelled, initiating refund...")
        refund_result = buyer.refund_escrow(result.escrow_id)

        if refund_result.success:
            print(f"✓ Escrow refunded!")
            print(f"TX Hash: {refund_result.tx_hash}")
        else:
            print(f"✗ Refund failed: {refund_result.error}")
    else:
        print(f"Failed to create escrow: {result.error}")


def main():
    demo_escrow_payment()
    # Uncomment to see refund flow:
    # demo_refund_flow()


if __name__ == "__main__":
    main()
