# Stateset Smart Contracts

This directory contains the CosmWasm smart contracts for the Stateset blockchain platform. These contracts provide secure, gas-optimized implementations for various business processes including escrow, loans, proofs, and workgroup management.

## 🚀 Recent Improvements

### Security Enhancements
- **Reentrancy Protection**: All contracts now implement proper state management to prevent reentrancy attacks
- **Access Control**: Enhanced authorization checks with proper role-based permissions
- **Input Validation**: Comprehensive validation of all user inputs with appropriate error handling
- **Gas Optimization**: Constants and limits to prevent gas limit attacks

### Code Quality
- **Updated Dependencies**: All CosmWasm dependencies updated to latest stable versions (1.5.0)
- **Error Handling**: Improved error types with descriptive messages
- **Event Emission**: Proper events for all state changes to improve observability
- **Migration Support**: Contract upgrade mechanisms with version validation
- **Comprehensive Tests**: Full test coverage including edge cases and security scenarios

## 📋 Available Contracts

### 1. Stateset Escrow (`stateset-escrow`)

A secure escrow contract for holding and conditionally releasing funds between parties.

**Key Features:**
- Multi-token support (native coins and CW20 tokens)
- Time or block height-based expiration
- Arbiter-controlled release mechanism
- Token whitelist support for CW20 tokens
- Reentrancy protection

**Security Improvements:**
- Input validation with size limits
- Proper authorization checks
- State cleanup on completion
- Gas-optimized constants

**Usage:**
```bash
# Create an escrow
statesetd tx wasm execute <contract_addr> '{
  "create": {
    "id": "deal123",
    "arbiter": "stateset1...",
    "recipient": "stateset1...",
    "title": "Service Agreement",
    "description": "Payment for development services",
    "end_height": 1000000
  }
}' --amount 1000usdc --from sender

# Approve release (arbiter only)
statesetd tx wasm execute <contract_addr> '{
  "approve": {"id": "deal123"}
}' --from arbiter

# Refund (after expiration or by arbiter)
statesetd tx wasm execute <contract_addr> '{
  "refund": {"id": "deal123"}
}' --from anyone
```

### 2. Stateset Loan (`stateset-loan`)

Advanced lending protocol with collateralization, interest rates, and liquidation mechanisms.

**Key Features:**
- Multi-asset lending and borrowing
- Dynamic interest rate calculations
- Health factor monitoring
- Liquidation protection
- Uncollateralized loan limits

**Security Features:**
- Proper liquidation thresholds
- Interest rate safeguards
- Asset validation
- Owner controls with emergency mechanisms

### 3. Stateset Proof (`stateset-proof`)

Document and data verification system using cryptographic proofs.

**Key Features:**
- DID-based identity management
- Payload verification
- Status tracking
- Provider authentication

**Recent Fixes:**
- Fixed compilation errors
- Improved error handling
- Added proper validation
- Enhanced event emission

### 4. Stateset Workgroup (`stateset-workgroup`)

Group membership and governance contract compatible with CW3 multisigs.

**Key Features:**
- Weighted membership system
- Hook system for external integrations
- Admin controls
- Snapshot support for historical queries

## 🛡️ Security Best Practices

All contracts implement the following security measures:

### 1. Input Validation
- Maximum length limits for strings
- Required field validation
- Address format verification
- Numeric range checks

### 2. Access Control
- Role-based permissions
- Owner/admin pattern implementation
- Emergency controls where appropriate
- Unauthorized access prevention

### 3. Reentrancy Protection
- State updates before external calls
- Proper cleanup on completion
- Lock mechanisms where needed

### 4. Gas Optimization
- Efficient storage patterns
- Minimal computation in loops
- Optimized data structures
- Gas limit protections

## 🧪 Testing

Each contract includes comprehensive test suites covering:

### Unit Tests
- Individual function testing
- Edge case validation
- Error condition verification
- Gas usage optimization verification

### Integration Tests
- Cross-contract interactions
- End-to-end workflows
- Security scenario testing
- Performance benchmarking

### Security Tests
- Reentrancy attack prevention
- Authorization bypass attempts
- Input fuzzing
- Gas limit testing

Run tests with:
```bash
cargo test --workspace
```

## 📦 Building and Deployment

### Prerequisites
- Rust 1.70+
- `wasm32-unknown-unknown` target
- CosmWasm tools

### Build
```bash
# Build all contracts
cargo build --release --target wasm32-unknown-unknown

# Optimize for deployment
cosmwasm-check artifacts/*.wasm
```

### Deploy
```bash
# Store contract code
statesetd tx wasm store artifacts/stateset_escrow.wasm --from deployer

# Instantiate contract
statesetd tx wasm instantiate <code_id> '{}' --from deployer --label "escrow-v1"
```

## 🔄 Migration

Contracts support migration for upgrades:

```bash
# Migrate to new version
statesetd tx wasm migrate <contract_addr> <new_code_id> '{}' --from admin
```

## 📊 Gas Usage

Optimized gas usage across all operations:

| Operation | Gas Usage | Notes |
|-----------|-----------|-------|
| Escrow Create | ~150k | Including validation |
| Escrow Approve | ~100k | With token transfers |
| Loan Borrow | ~200k | With health checks |
| Proof Verify | ~80k | Simple verification |

## 🤝 Contributing

When contributing to smart contracts:

1. **Security First**: All changes must pass security review
2. **Test Coverage**: Maintain 90%+ test coverage
3. **Gas Efficiency**: Profile gas usage for optimizations
4. **Documentation**: Update documentation for all changes
5. **Migration Path**: Consider upgrade implications

## 📞 Support

For questions about smart contract implementation:

- **Documentation**: See individual contract README files
- **Issues**: Report security issues privately
- **Community**: Join our developer Discord

## ⚠️ Security Considerations

**Important**: These contracts handle valuable assets. Always:

- Audit code before mainnet deployment
- Test thoroughly on testnets
- Use proper key management
- Monitor contract activity
- Have incident response plans

## 📄 License

Licensed under Apache 2.0. See [LICENSE](../../LICENSE) for details.