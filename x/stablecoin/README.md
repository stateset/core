# Stablecoin Module (`x/stablecoin`)

The Stablecoin module manages the native stablecoin `ssUSD` on Stateset. It supports two minting paths:

1. **Vault-based CDPs**: Users lock approved collateral (e.g., `stst`) in personal vaults to mint `ssusd`. Vault minting is **disabled by default** for a strictly reserve‑backed ssUSD; governance must enable `vault_minting_enabled=true` before CDPs can be used.
2. **Reserve-backed issuance**: Users deposit approved tokenized US Treasury assets (e.g., `usdy`, `stbt`, `ousg`) to mint `ssusd` 1:1 and redeem back into those assets.

## Features

### Vault-based collateral
- Individual vaults per user
- Over-collateralized debt positions
- Oracle-valued collateral with automatic liquidation

### Reserve-backed stablecoin
- 100%+ reserve ratio enforced via `ReserveParams`
- Minting from tokenized treasuries with haircuts and allocation limits
- Redemption queue with optional delay, KYC gating, and daily limits
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
| `MsgExecuteRedemption` | Execute a pending redemption (authority or auto) |
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
- Approved `tokenized_treasuries` list (haircuts, allocation caps, oracle denom)
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

## Events

**Vault events**
- `vault_created`, `collateral_deposited`, `collateral_withdrawn`
- `stablecoin_minted`, `stablecoin_repaid`, `vault_liquidated`

**Reserve events**
- `reserve_deposit`, `reserve_mint`
- `redemption_requested`, `redemption_executed`, `redemption_cancelled`
- `reserve_attestation`, `reserve_params_updated`
