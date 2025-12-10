#!/usr/bin/env python3
"""
Batch Payments Example

Demonstrates efficient multi-recipient payments:
- Payroll disbursement
- Revenue splitting
- Affiliate payouts
- Mass refunds

Usage: python batch_payments.py
"""

import json
import subprocess
import os
from typing import List, Dict, Any, Optional
from dataclasses import dataclass
from decimal import Decimal


# Configuration
CHAIN_ID = os.getenv("STATESET_CHAIN_ID", "stateset-1")
NODE = os.getenv("STATESET_NODE", "tcp://localhost:26657")
KEYRING = os.getenv("STATESET_KEYRING", "test")
GAS_PRICES = "0.025ustake"
GAS_LIMIT = "800000"  # Higher for batch operations
BINARY = os.getenv("STATESET_BINARY", "statesetd")

SSUSD_DENOM = "ssusd"
SSUSD_DECIMALS = 6


@dataclass
class Payment:
    """Single payment in a batch"""
    recipient: str
    amount: int
    name: str = ""
    category: str = ""


@dataclass
class BatchResult:
    """Result of batch payment"""
    success: bool
    tx_hash: Optional[str] = None
    payments_count: int = 0
    total_amount: int = 0
    error: Optional[str] = None


class BatchPaymentClient:
    """Client for batch payment operations"""

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

    def get_balance(self, denom: str = SSUSD_DENOM) -> int:
        cmd = [
            BINARY, "query", "bank", "balances", self.address,
            "--chain-id", CHAIN_ID,
            "--node", NODE,
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            return 0
        data = json.loads(result.stdout)
        for balance in data.get("balances", []):
            if balance["denom"] == denom:
                return int(balance["amount"])
        return 0

    def batch_send(self, payments: List[Payment], memo: str = "") -> BatchResult:
        """Send batch payments"""
        total_amount = sum(p.amount for p in payments)

        # Build multi-send transaction
        batch_data = json.dumps({
            "sender": self.address,
            "payments": [
                {
                    "recipient": p.recipient,
                    "amount": {"denom": SSUSD_DENOM, "amount": str(p.amount)}
                }
                for p in payments
            ],
            "memo": memo
        })

        cmd = [
            BINARY, "tx", "settlement", "batch-payment", batch_data,
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
            return BatchResult(
                success=False,
                payments_count=len(payments),
                total_amount=total_amount,
                error=result.stderr
            )

        try:
            data = json.loads(result.stdout)
            return BatchResult(
                success=True,
                tx_hash=data.get("txhash"),
                payments_count=len(payments),
                total_amount=total_amount
            )
        except json.JSONDecodeError:
            return BatchResult(
                success=False,
                payments_count=len(payments),
                total_amount=total_amount,
                error="Failed to parse response"
            )


def format_amount(amount: int) -> str:
    return f"{Decimal(amount) / Decimal(10 ** SSUSD_DECIMALS):.2f}"


def to_base_units(amount: float) -> int:
    return int(amount * (10 ** SSUSD_DECIMALS))


def print_table(headers: List[str], rows: List[List[str]], footer: Optional[List[str]] = None):
    """Print a formatted table"""
    widths = [max(len(str(h)), max(len(str(row[i])) for row in rows)) + 2
              for i, h in enumerate(headers)]

    # Header
    print("┌" + "┬".join("─" * w for w in widths) + "┐")
    print("│" + "│".join(h.center(widths[i]) for i, h in enumerate(headers)) + "│")
    print("├" + "┼".join("─" * w for w in widths) + "┤")

    # Rows
    for row in rows:
        print("│" + "│".join(str(row[i]).center(widths[i]) for i in range(len(headers))) + "│")

    # Footer
    if footer:
        print("├" + "┼".join("─" * w for w in widths) + "┤")
        print("│" + "│".join(str(footer[i]).center(widths[i]) for i in range(len(headers))) + "│")

    print("└" + "┴".join("─" * w for w in widths) + "┘")


def payroll_example(client: BatchPaymentClient):
    """Example 1: Payroll Disbursement"""
    print("━" * 60)
    print("  Example 1: Payroll Disbursement")
    print("━" * 60)
    print()

    employees = [
        Payment("stateset1alice...", to_base_units(5000), "Alice", "Engineering"),
        Payment("stateset1bob...", to_base_units(4500), "Bob", "Marketing"),
        Payment("stateset1carol...", to_base_units(4000), "Carol", "Sales"),
        Payment("stateset1dave...", to_base_units(3500), "Dave", "Operations"),
        Payment("stateset1eve...", to_base_units(3000), "Eve", "Support"),
    ]

    print("Payroll Recipients:")
    headers = ["Employee", "Address", "Amount"]
    rows = [[f"{p.name} ({p.category})", p.recipient[:14], f"${format_amount(p.amount)}"]
            for p in employees]
    total = sum(p.amount for p in employees)
    print_table(headers, rows, ["TOTAL", "", f"${format_amount(total)}"])
    print()

    # Execute batch payment
    print("Executing batch payroll transaction...")
    result = client.batch_send(employees, f"Payroll - {__import__('datetime').date.today()}")

    if result.success:
        print(f"✓ Payroll disbursed!")
        print(f"  TX Hash: {result.tx_hash}")
        print(f"  Total Paid: ${format_amount(result.total_amount)}")
        print(f"  Recipients: {result.payments_count}")
    else:
        print(f"✗ Payroll failed: {result.error}")
    print()


def revenue_split_example(client: BatchPaymentClient):
    """Example 2: Revenue Split"""
    print("━" * 60)
    print("  Example 2: Revenue Split (Marketplace Sale)")
    print("━" * 60)
    print()

    sale_amount = 1000  # Total sale in USD

    distribution = {
        "Seller": {"address": "stateset1seller...", "percentage": 85},
        "Platform": {"address": "stateset1platform...", "percentage": 10},
        "Affiliate": {"address": "stateset1affiliate...", "percentage": 5},
    }

    print(f"Sale Amount: ${sale_amount} ssUSD")
    print()
    print("Revenue Distribution:")

    payments = []
    headers = ["Recipient", "Percentage", "Amount"]
    rows = []

    for role, config in distribution.items():
        amount = int(sale_amount * config["percentage"] / 100)
        payments.append(Payment(config["address"], to_base_units(amount), role))
        rows.append([role, f"{config['percentage']}%", f"${amount:.2f}"])

    print_table(headers, rows)
    print()

    print("Executing revenue split...")
    result = client.batch_send(payments, "Revenue split - Sale #12345")

    if result.success:
        print(f"✓ Revenue distributed!")
        print(f"  TX Hash: {result.tx_hash}")
    else:
        print(f"✗ Revenue split failed: {result.error}")
    print()


def affiliate_payouts_example(client: BatchPaymentClient):
    """Example 3: Affiliate Payouts"""
    print("━" * 60)
    print("  Example 3: Monthly Affiliate Payouts")
    print("━" * 60)
    print()

    affiliates = [
        {"id": "AFF001", "address": "stateset1aff1...", "sales": 50, "commission": 500},
        {"id": "AFF002", "address": "stateset1aff2...", "sales": 35, "commission": 350},
        {"id": "AFF003", "address": "stateset1aff3...", "sales": 28, "commission": 280},
        {"id": "AFF004", "address": "stateset1aff4...", "sales": 22, "commission": 220},
        {"id": "AFF005", "address": "stateset1aff5...", "sales": 15, "commission": 150},
    ]

    min_payout = 100  # Minimum payout threshold
    eligible = [a for a in affiliates if a["commission"] >= min_payout]

    print(f"Minimum Payout Threshold: ${min_payout}")
    print()
    print("Affiliate Earnings:")

    headers = ["Affiliate", "Sales", "Commission", "Status"]
    rows = []
    for aff in affiliates:
        status = "✓ Eligible" if aff["commission"] >= min_payout else "Below min"
        rows.append([aff["id"], str(aff["sales"]), f"${aff['commission']:.2f}", status])

    total_payout = sum(a["commission"] for a in eligible)
    print_table(headers, rows, ["TOTAL", "", f"${total_payout:.2f}", ""])
    print()

    payments = [
        Payment(a["address"], to_base_units(a["commission"]), a["id"])
        for a in eligible
    ]

    print(f"Processing {len(eligible)} payouts...")
    result = client.batch_send(payments, "Affiliate payouts - monthly")

    if result.success:
        print(f"✓ Affiliate payouts complete!")
        print(f"  TX Hash: {result.tx_hash}")
        print(f"  Affiliates Paid: {len(eligible)}")
        print(f"  Below Threshold: {len(affiliates) - len(eligible)}")
    else:
        print(f"✗ Payouts failed: {result.error}")
    print()


def mass_refund_example(client: BatchPaymentClient):
    """Example 4: Mass Refunds"""
    print("━" * 60)
    print("  Example 4: Mass Refund Processing")
    print("━" * 60)
    print()

    refunds = [
        {"order_id": "ORD-1001", "customer": "stateset1cust1...", "amount": 150, "reason": "Product defect"},
        {"order_id": "ORD-1045", "customer": "stateset1cust2...", "amount": 89, "reason": "Wrong item"},
        {"order_id": "ORD-1089", "customer": "stateset1cust3...", "amount": 225, "reason": "Cancelled"},
        {"order_id": "ORD-1102", "customer": "stateset1cust4...", "amount": 75, "reason": "Never shipped"},
    ]

    print("Approved Refunds:")
    headers = ["Order ID", "Customer", "Amount", "Reason"]
    rows = [[r["order_id"], r["customer"][:14], f"${r['amount']:.2f}", r["reason"]]
            for r in refunds]
    total_refunds = sum(r["amount"] for r in refunds)
    print_table(headers, rows, ["TOTAL", "", f"${total_refunds:.2f}", ""])
    print()

    payments = [
        Payment(r["customer"], to_base_units(r["amount"]), r["order_id"])
        for r in refunds
    ]

    print("Processing refunds...")
    result = client.batch_send(payments, "Batch refund processing")

    if result.success:
        print(f"✓ Refunds processed!")
        print(f"  TX Hash: {result.tx_hash}")
        print(f"  Refunds Issued: {len(refunds)}")
        print(f"  Total Refunded: ${total_refunds:.2f}")
    else:
        print(f"✗ Refunds failed: {result.error}")
    print()


def main():
    print("═" * 60)
    print("  ssUSD Batch Payments Examples")
    print("  Efficient Multi-Recipient Transfers")
    print("═" * 60)
    print()

    # Connect
    print("Connecting to Stateset network...")
    client = BatchPaymentClient("treasury")  # Use account with funds
    print(f"Connected as: {client.address}")
    print()

    # Check balance
    balance = client.get_balance()
    print(f"Available Balance: ${format_amount(balance)} ssUSD")
    print()

    # Run examples
    payroll_example(client)
    revenue_split_example(client)
    affiliate_payouts_example(client)
    mass_refund_example(client)

    print("═" * 60)
    print("  All Batch Payment Examples Completed!")
    print("═" * 60)


if __name__ == "__main__":
    main()
