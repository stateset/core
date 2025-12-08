# ZKP Verify Module

The ZKP Verify module provides zero-knowledge proof verification capabilities for the Stateset blockchain.

## Overview

The zkpverify module handles:
- **Proof Verification**: Verify zero-knowledge proofs on-chain
- **Verifier Management**: Register and manage verification keys
- **Proof Records**: Store verification results

## Features

### Proof Types Supported
- Groth16 proofs
- PLONK proofs
- Custom verification schemes

### Verifier Management
- Register verification keys
- Update verifier configurations
- Deactivate verifiers

### Verification Records
- Store verification results
- Audit trail for proofs
- Verification timestamps

## Messages

| Message | Description |
|---------|-------------|
| `MsgVerifyProof` | Submit proof for verification |
| `MsgRegisterVerifier` | Register new verifier |
| `MsgUpdateVerifier` | Update verifier configuration |

## State

| Key | Value |
|-----|-------|
| `0x01{verifier_id}` | Verifier |
| `0x02{proof_id}` | VerificationRecord |
| `0x03` | Params |

## Events

| Event | Attributes |
|-------|------------|
| `proof_verified` | verifier_id, result, timestamp |
| `proof_rejected` | verifier_id, reason |
| `verifier_registered` | id, scheme |
