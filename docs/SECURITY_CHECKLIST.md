# Security Checklist for Developers

Quick reference checklist for implementing and reviewing financial module features.

## Before Writing Code

- [ ] Read the [Security Architecture](./security_architecture.md) document
- [ ] Understand the threat model for your module
- [ ] Review similar existing code for patterns
- [ ] Identify all user inputs and state changes
- [ ] Plan authorization checks
- [ ] Consider edge cases and error conditions

## Input Validation

### Addresses
- [ ] Validate all addresses using `sdk.AccAddressFromBech32()`
- [ ] Check for non-empty addresses
- [ ] Handle validation errors appropriately
- [ ] Prevent self-transfer (payer == payee) where appropriate

### Amounts
- [ ] Verify amounts are positive using `IsPositive()`
- [ ] Check against minimum amounts (prevent dust spam)
- [ ] Check against maximum amounts (prevent overflow, enforce limits)
- [ ] Use `sdkmath.Int` or `sdkmath.LegacyDec` for calculations
- [ ] Never use raw integer types for financial amounts

### Denoms
- [ ] Validate denom is supported/whitelisted
- [ ] Check collateral type is active (stablecoin module)
- [ ] Use constants for expected denoms

### Other Inputs
- [ ] Validate expiration times are within bounds
- [ ] Validate nonces strictly increase
- [ ] Validate signatures cryptographically
- [ ] Sanitize or limit metadata fields
- [ ] Validate webhook URLs (HTTPS only, not blacklisted)

## Authorization Checks

- [ ] Verify caller has permission for the operation
- [ ] For vault operations: check `vault.Owner == msg.Sender`
- [ ] For escrow: only sender can release, only recipient can refund
- [ ] For payments: only payee can settle, only payer can cancel
- [ ] For system operations: check `msg.Authority == keeper.GetAuthority()`
- [ ] Check authorization BEFORE any state modifications

## State Modifications

### Ordering
- [ ] Check status/preconditions first
- [ ] Update status before fund transfers
- [ ] Perform fund transfers
- [ ] Update other state
- [ ] Emit events last

### Status Machines
- [ ] Define valid state transitions
- [ ] Prevent invalid transitions (e.g., settling cancelled payment)
- [ ] Check current status before operations
- [ ] Update status atomically

### Example Pattern
```go
// 1. Validate inputs
if !amount.IsPositive() {
    return err
}

// 2. Get and validate state
item, found := k.GetItem(ctx, id)
if !found {
    return ErrNotFound
}

// 3. Check authorization
if item.Owner != sender.String() {
    return ErrUnauthorized
}

// 4. Check status
if item.Status != StatusPending {
    return ErrInvalidStatus
}

// 5. Update status first
item.Status = StatusCompleted

// 6. Perform fund transfers
if err := k.bankKeeper.SendCoins(...); err != nil {
    return err
}

// 7. Update other state
k.SetItem(ctx, item)

// 8. Emit events
ctx.EventManager().EmitEvent(...)
```

## Compliance Checks

- [ ] Check compliance for all parties before fund movements
- [ ] Use `complianceKeeper.AssertCompliant()` on wrapped context
- [ ] Check payer compliance
- [ ] Check payee/recipient compliance
- [ ] Handle compliance failures gracefully
- [ ] Document compliance requirements

```go
wrappedCtx := sdk.WrapSDKContext(ctx)
if err := k.compKeeper.AssertCompliant(wrappedCtx, userAddr); err != nil {
    return types.ErrComplianceCheckFailed
}
```

## Oracle Integration

### For Operations Using Prices

- [ ] Get price using staleness-checked method
- [ ] Handle missing price error (fail operation)
- [ ] Handle stale price error (fail operation)
- [ ] Log oracle errors for monitoring
- [ ] Document oracle dependency

```go
price, err := k.oracleKeeper.GetPriceDecSafe(ctx, denom)
if err != nil {
    ctx.Logger().Error("oracle price unavailable", "denom", denom, "error", err)
    return types.ErrPriceNotFound
}
```

### Price Validation
- [ ] Use deviation limits for price updates
- [ ] Check update frequency
- [ ] Verify provider authorization
- [ ] Record price history

## Arithmetic Operations

- [ ] Always use `sdkmath.Int` for integer amounts
- [ ] Use `sdkmath.LegacyDec` for decimal calculations
- [ ] Never use Go's built-in int/float types for financial values
- [ ] Check for zero before division
- [ ] Handle fee calculations carefully (rounding)

```go
// Good
feeAmount := amount.Amount.Mul(sdkmath.NewInt(int64(feeRateBps))).Quo(sdkmath.NewInt(10000))
fee := sdk.NewCoin(amount.Denom, feeAmount)

// Bad
feeAmount := amount.Amount.Int64() * feeRateBps / 10000 // NEVER DO THIS
```

## Error Handling

- [ ] Return specific error types (use errorsmod.Wrap)
- [ ] Include context in error messages
- [ ] Don't leak sensitive information in errors
- [ ] Log errors with appropriate level
- [ ] Return errors, don't panic (unless truly exceptional)

```go
// Good
if err != nil {
    return errorsmod.Wrapf(types.ErrInsufficientFunds,
        "required %s, available %s", required, available)
}

// Bad
if err != nil {
    panic(err) // Don't panic in normal error conditions
}
```

## Event Emission

- [ ] Emit events for all significant operations
- [ ] Include relevant attributes (IDs, addresses, amounts)
- [ ] Emit events AFTER state changes complete
- [ ] Use consistent event naming
- [ ] Document event schema

```go
ctx.EventManager().EmitEvent(
    sdk.NewEvent(
        types.EventTypeSettlementCompleted,
        sdk.NewAttribute(types.AttributeKeySettlementID, fmt.Sprintf("%d", id)),
        sdk.NewAttribute(types.AttributeKeySender, sender),
        sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
    ),
)
```

## Module Accounts

- [ ] Use correct module account name constant
- [ ] Verify module account exists in init
- [ ] Only mint/burn if module has permission
- [ ] Never directly access another module's account
- [ ] Track module balance in state where needed

## Rate Limiting

- [ ] Enforce minimum amounts (prevent dust)
- [ ] Enforce maximum amounts (prevent overflow, DoS)
- [ ] Implement batch size limits
- [ ] Consider time-based rate limits (daily limits)
- [ ] Document rate limit rationale

## Invariants

- [ ] Define financial invariants for your module
- [ ] Register invariants in module
- [ ] Ensure invariants always hold after operations
- [ ] Test that invariants catch errors
- [ ] Document what each invariant guarantees

```go
func ModuleBalanceInvariant(k Keeper) sdk.Invariant {
    return func(ctx sdk.Context) (string, bool) {
        // Sum all pending amounts
        total := sdk.NewCoins()
        k.IterateItems(ctx, func(item Item) bool {
            if item.Status == StatusPending {
                total = total.Add(item.Amount)
            }
            return false
        })

        // Compare to module balance
        moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
        balance := k.bankKeeper.GetAllBalances(ctx, moduleAddr)

        if !balance.IsAllGTE(total) {
            return sdk.FormatInvariant(
                types.ModuleName, "module-balance",
                fmt.Sprintf("module balance %s < pending total %s", balance, total),
            ), true
        }
        return "", false
    }
}
```

## Testing

### Unit Tests
- [ ] Test happy path
- [ ] Test each error condition
- [ ] Test authorization checks (unauthorized attempts)
- [ ] Test input validation (invalid inputs)
- [ ] Test edge cases (zero, max, negative)
- [ ] Test state consistency after errors

### Security Tests
- [ ] Test unauthorized access attempts
- [ ] Test double-spend prevention
- [ ] Test overflow/underflow scenarios
- [ ] Test invariant violations
- [ ] Test compliance bypass attempts
- [ ] Document threat model in test comments

### Integration Tests
- [ ] Test cross-module interactions
- [ ] Test with real keeper instances
- [ ] Test end-to-end user flows
- [ ] Test genesis export/import

## Code Review Checklist

### For Reviewers

- [ ] All user inputs validated?
- [ ] Authorization checks present and correct?
- [ ] Safe arithmetic (no raw int/float)?
- [ ] Status checked before operations?
- [ ] Compliance enforced?
- [ ] Errors handled properly?
- [ ] Events emitted?
- [ ] Tests cover security aspects?
- [ ] Documentation updated?
- [ ] Invariants still hold?

## Pre-Commit Checklist

- [ ] All tests pass (unit + integration + security)
- [ ] No compiler warnings
- [ ] Code formatted with `gofmt`
- [ ] Linter passes
- [ ] Documentation updated if needed
- [ ] Changelog updated if user-facing change

## Pre-PR Checklist

- [ ] Branch up to date with main/master
- [ ] All tests pass on clean build
- [ ] Self-review completed using code review checklist
- [ ] Security implications documented
- [ ] Breaking changes documented
- [ ] Migration plan for state changes (if any)

## Common Pitfalls to Avoid

### ❌ DON'T
```go
// Don't use raw integer types
fee := int64(amount) * rate / 10000

// Don't skip authorization
k.TransferFunds(ctx, from, to, amount) // Missing auth check!

// Don't panic in normal cases
if err != nil {
    panic(err)
}

// Don't modify state before validation
k.SetItem(ctx, item)
if !item.Amount.IsPositive() {
    return err // Already modified state!
}

// Don't skip compliance
k.bankKeeper.SendCoins(...) // Missing compliance check!

// Don't use stale prices
price, _ := k.oracleKeeper.GetPrice(ctx, denom) // No staleness check!
```

### ✅ DO
```go
// Use SDK math types
feeAmount := amount.Amount.Mul(sdkmath.NewInt(rate)).Quo(sdkmath.NewInt(10000))

// Check authorization first
if item.Owner != sender.String() {
    return types.ErrUnauthorized
}
k.TransferFunds(ctx, from, to, amount)

// Return errors
if err != nil {
    return errorsmod.Wrap(err, "operation failed")
}

// Validate before modifying state
if !item.Amount.IsPositive() {
    return types.ErrInvalidAmount
}
k.SetItem(ctx, item)

// Check compliance
wrappedCtx := sdk.WrapSDKContext(ctx)
if err := k.compKeeper.AssertCompliant(wrappedCtx, userAddr); err != nil {
    return err
}
k.bankKeeper.SendCoins(...)

// Use staleness-checked prices
price, err := k.oracleKeeper.GetPriceDecSafe(ctx, denom)
if err != nil {
    return types.ErrPriceNotFound
}
```

## Emergency Response

If you discover a security issue:

1. **DO NOT** commit the fix to a public branch
2. **DO NOT** discuss publicly
3. **Contact security team immediately**: security@stateset.io
4. **Document** the issue, impact, and reproduction steps
5. **Wait** for security team guidance before proceeding

## Security Resources

- **Security Architecture**: `/docs/security_architecture.md`
- **Security Tests**: `/tests/security/`
- **Test README**: `/tests/security/README.md`
- **This Checklist**: `/docs/SECURITY_CHECKLIST.md`
- **Review Summary**: `/docs/SECURITY_REVIEW_SUMMARY.md`

## Questions?

When in doubt:
1. Review similar existing code
2. Check the security architecture document
3. Ask in team chat/code review
4. Prefer failing safely over risky behavior
5. Document your security assumptions

Remember: **Security is everyone's responsibility**. When you write code that handles user funds, you hold a position of trust. Take it seriously.

---

**Last Updated**: 2025-12-10
**Maintained By**: Stateset Security Team
