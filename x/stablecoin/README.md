# Stablecoin Module (`x/stablecoin`)

The Stablecoin module manages the native USD stablecoin `ssUSD` (`ssusd`) on Stateset. It supports two minting paths:

1. **Vault-based CDPs**: Users lock approved collateral (e.g., `stst`) in personal vaults to mint `ssusd`. Vault minting is **disabled by default** for a strictly reserve‑backed ssUSD; governance must enable `vault_minting_enabled=true` before CDPs can be used.
2. **Reserve-backed issuance (Path B, default)**: Users deposit approved **tokenized US Treasury Notes** (T-Notes) to mint `ssusd` and redeem back into those notes. The default and only supported on-chain reserve asset is `ustn` with `underlying_type="t_note"`.

## Features

### Vault-based collateral
- Individual vaults per user
- Over-collateralized debt positions
- Oracle-valued collateral with automatic liquidation

### Reserve-backed stablecoin
- 100%+ reserve ratio enforced via `ReserveParams`
- Minting from tokenized US Treasury Notes with haircuts and allocation limits
- Redemption requests with optional delay, KYC gating, daily limits, and reserve locking
- Off-chain attestations folded into total backing

## Messages

| Message | Purpose |
|---------|---------|
| `MsgCreateVault` | Create a new collateral vault |
| `MsgDepositCollateral` | Add collateral to a vault |
| `MsgWithdrawCollateral` | Withdraw collateral from a vault |
| `MsgMintStablecoin` | Mint `ssusd` against vault collateral |
| `MsgRepayStablecoin` | Repay vault debt with `ssusd` |
| `MsgLiquidateVault` | Liquidate an unhealthy vault |
| `MsgDepositReserve` | Deposit approved tokenized treasuries to mint `ssusd` |
| `MsgRequestRedemption` | Request redemption of `ssusd` into an approved reserve asset |
| `MsgExecuteRedemption` | Execute a pending redemption (anyone after delay) |
| `MsgCancelRedemption` | Cancel a pending redemption (authority only) |
| `MsgUpdateReserveParams` | Update reserve policy (governance) |
| `MsgRecordAttestation` | Record off-chain reserve attestation (approved attester) |
| `MsgSetApprovedAttester` | Add/remove approved attesters (authority only) |

## Parameters

### Vault params (`Params`)
- `vault_minting_enabled`: global gate for CDP minting (default `false`)
- `collateral_params`: per-denom risk config (`liquidation_ratio`, `stability_fee`, `debt_limit`, `active`)

### Reserve params (`ReserveParams`)
- Reserve ratio targets and daily mint/redeem limits
- Mint/redeem fees, minimum amounts, optional redemption delay
- Approved `tokenized_treasuries` list (haircuts, allocation caps, oracle denom). On-chain reserve assets are restricted to `underlying_type="t_note"`.
- `require_kyc`, `mint_paused`, `redeem_paused`

## State

| Key Prefix | Value |
|------------|-------|
| `0x01{id}` | Vault |
| `0x02` | Vault `Params` |
| `0x03` | Next vault ID |
| `0x10` | `ReserveParams` |
| `0x11` | `Reserve` |
| `0x12{id}` | `ReserveDeposit` |
| `0x13{id}` | `RedemptionRequest` |
| `0x14{date}` | `DailyMintStats` |
| `0x15{id}` | `OffChainReserveAttestation` |
| `0x16` | Next deposit ID |
| `0x17` | Next redemption ID |
| `0x18` | Next attestation ID |
| `0x19{addr}` | Approved attester flag |
| `0x1A` | Locked reserves for pending redemptions (JSON-encoded `sdk.Coins`) |

## Reserve-backed Mint/Redeem Semantics (Path B)

- **Mint (`MsgDepositReserve`)**: transfers approved reserve assets (default: `ustn`) into the stablecoin module account, applies haircut + mint fee using the oracle price, and mints `ssusd` to the depositor.
- **Redeem request (`MsgRequestRedemption`)**: transfers `ssusd` into the module account and **burns immediately**, computes and stores `output_amount` at the current oracle price (after redemption fee), and **locks** the corresponding reserve assets so later redemptions cannot overbook reserves.
- **Redeem execute (`MsgExecuteRedemption`)**: after `executable_after`, transfers the stored `output_amount` to the requester and unlocks it (then updates the on-chain reserve totals).
- **Cancel (`MsgCancelRedemption`)**: authority-only; unlocks the stored `output_amount` and mints back the burned `ssusd` to the requester.

## Events

**Vault events**
- `vault_created`, `collateral_deposited`, `collateral_withdrawn`
- `stablecoin_minted`, `stablecoin_repaid`, `vault_liquidated`

**Reserve events**
- `reserve_deposit`
- `redemption_requested`, `redemption_executed`, `redemption_cancelled`
- `reserve_attestation`, `reserve_params_updated`
