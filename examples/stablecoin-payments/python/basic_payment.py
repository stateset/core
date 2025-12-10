#!/usr/bin/env python3
"""
Basic ssUSD Payment Example

Demonstrates fundamental stablecoin operations:
- Connect to Stateset network
- Check balances
- Send instant ssUSD transfers
- Query transaction results

Usage: python basic_payment.py
"""

import json
import subprocess
import os
from typing import Optional, Dict, Any
from dataclasses import dataclass
from decimal import Decimal


# Configuration
CHAIN_ID = os.getenv("STATESET_CHAIN_ID", "stateset-1")
NODE = os.getenv("STATESET_NODE", "tcp://localhost:26657")
KEYRING = os.getenv("STATESET_KEYRING", "test")
GAS_PRICES = "0.025ustake"
GAS_LIMIT = "500000"
BINARY = os.getenv("STATESET_BINARY", "statesetd")

# Stablecoin constants
SSUSD_DENOM = "ssusd"
SSUSD_DECIMALS = 6


@dataclass
class TransactionResult:
    """Result of a blockchain transaction"""
    success: bool
    tx_hash: Optional[str] = None
    height: Optional[int] = None
    error: Optional[str] = None


class StatesetClient:
    """Simple client for interacting with Stateset blockchain"""

    def __init__(self, key_name: str):
        self.key_name = key_name
        self.address = self._get_address()

    def _get_address(self) -> str:
        """Get the address for the configured key"""
        cmd = [
            BINARY, "keys", "show", self.key_name,
            "--keyring-backend", KEYRING,
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)
        if result.returncode != 0:
            raise Exception(f"Failed to get address: {result.stderr}")
        return json.loads(result.stdout)["address"]

    def _execute_tx(self, module: str, *args) -> TransactionResult:
        """Execute a transaction"""
        cmd = [
            BINARY, "tx", module, *args,
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
                height=data.get("height")
            )
        except json.JSONDecodeError:
            return TransactionResult(success=False, error="Failed to parse response")

    def _query(self, module: str, *args) -> Dict[str, Any]:
        """Execute a query"""
        cmd = [
            BINARY, "query", module, *args,
            "--chain-id", CHAIN_ID,
            "--node", NODE,
            "--output", "json"
        ]
        result = subprocess.run(cmd, capture_output=True, text=True)

        if result.returncode != 0:
            raise Exception(f"Query failed: {result.stderr}")

        return json.loads(result.stdout)

    def get_balance(self, denom: str = SSUSD_DENOM) -> int:
        """Get balance for a specific denomination"""
        result = self._query("bank", "balances", self.address)
        for balance in result.get("balances", []):
            if balance["denom"] == denom:
                return int(balance["amount"])
        return 0

    def get_all_balances(self) -> Dict[str, int]:
        """Get all balances"""
        result = self._query("bank", "balances", self.address)
        return {b["denom"]: int(b["amount"]) for b in result.get("balances", [])}

    def send(self, recipient: str, amount: int, denom: str = SSUSD_DENOM) -> TransactionResult:
        """Send tokens to a recipient"""
        amount_str = f"{amount}{denom}"
        return self._execute_tx("bank", "send", self.address, recipient, amount_str)


def format_amount(amount: int, decimals: int = SSUSD_DECIMALS) -> str:
    """Format amount from base units to human-readable"""
    value = Decimal(amount) / Decimal(10 ** decimals)
    return f"{value:.2f}"


def to_base_units(amount: float, decimals: int = SSUSD_DECIMALS) -> int:
    """Convert human-readable amount to base units"""
    return int(amount * (10 ** decimals))


def print_separator(char: str = "=", length: int = 60):
    print(char * length)


def print_header(title: str):
    print_separator()
    print(f"  {title}")
    print_separator()
    print()


def main():
    print_header("ssUSD Basic Payment Example")

    # Step 1: Connect to Stateset
    print("1. Connecting to Stateset network...")
    try:
        sender = StatesetClient("alice")  # Use your key name
        print(f"   Connected as: {sender.address}")
    except Exception as e:
        print(f"   Error: {e}")
        print("   Make sure you have created an account: statesetd keys add alice --keyring-backend test")
        return
    print()

    # Step 2: Check balance
    print("2. Checking sender balance...")
    balance = sender.get_balance()
    print(f"   ssUSD Balance: {format_amount(balance)} ssUSD")

    if balance == 0:
        print()
        print("   ⚠️  No ssUSD balance. Please fund your account first.")
        print("   See: scripts/test-ssusd-transfers.sh")
        return
    print()

    # Step 3: Define payment
    recipient_address = "stateset1recipient..."  # Replace with actual address
    payment_amount = 100.0  # 100 ssUSD

    print("3. Payment Details:")
    print(f"   From:   {sender.address}")
    print(f"   To:     {recipient_address}")
    print(f"   Amount: {payment_amount} ssUSD")
    print()

    # Step 4: Execute transfer
    print("4. Executing transfer...")
    result = sender.send(recipient_address, to_base_units(payment_amount))

    if result.success:
        print(f"   ✓ Transaction successful!")
        print(f"   TX Hash: {result.tx_hash}")
    else:
        print(f"   ✗ Transaction failed: {result.error}")
    print()

    # Step 5: Verify new balance
    print("5. Verifying balance after transfer...")
    new_balance = sender.get_balance()
    print(f"   New ssUSD Balance: {format_amount(new_balance)} ssUSD")
    print()

    print_header("Payment Completed Successfully!")


if __name__ == "__main__":
    main()
