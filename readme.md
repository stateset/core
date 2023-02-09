# Stateset Core

**Stateset Core** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

```

Transactions subcommands

Usage:
  statesetd tx [flags]
  statesetd tx [command]

Available Commands:
                      
  agreement           agreement transactions subcommands
  bank                Bank transaction subcommands
  broadcast           Broadcast transactions generated offline
  crisis              Crisis transactions subcommands
  decode              Decode a binary encoded transaction string
  did                 did transactions subcommands
  distribution        Distribution transactions subcommands
  encode              Encode transactions generated offline
  evidence            Evidence transaction subcommands
  feegrant            Feegrant transactions subcommands
  gov                 Governance transactions subcommands
  ibc                 IBC transaction subcommands
  ibc-transfer        IBC fungible token transfer transaction subcommands
  invoice             invoice transactions subcommands
  loan                loan transactions subcommands
  multisign           Generate multisig signatures for transactions generated offline
  proof               proof transactions subcommands
  purchaseorder       purchaseorder transactions subcommands
  sign                Sign a transaction generated offline
  sign-batch          Sign transaction batch files
  slashing            Slashing transaction subcommands
  staking             Staking transaction subcommands
  validate-signatures validate transactions signatures
  vesting             Vesting transaction subcommands
  wasm                Wasm transaction subcommands

  ```

### CosmWasm

CosmWasm is used for deploying smart contracts on the Stateset Network.

```
Wasm transaction subcommands

Usage:
  statesetd tx wasm [flags]
  statesetd tx wasm [command]

Available Commands:
  clear-contract-admin Clears admin for a contract to prevent further migrations
  execute              Execute a command on a wasm contract
  instantiate          Instantiate a wasm contract
  migrate              Migrate a wasm contract to a new code version
  set-contract-admin   Set new admin for a contract
  store                Upload a wasm binary

  ```