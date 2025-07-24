# Stateset Core

**Stateset Core** is a blockchain built using Cosmos SDK and Tendermint. It provides a decentralized platform for business processes including agreements, invoices, purchase orders, loans, and more.

## Prerequisites

Before building and running the Stateset blockchain node, ensure you have the following dependencies installed:

- **Go 1.21+**: [Download and install Go](https://golang.org/dl/)
- **Git**: For cloning the repository
- **Make**: For using build scripts (optional)

### Verify Go Installation

```bash
go version
```

You should see output similar to `go version go1.21.0 linux/amd64`.

## Building the Blockchain Node

### 1. Clone the Repository

```bash
git clone https://github.com/stateset/core.git
cd core
```

### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

### 3. Build the Node Binary

#### Using Make (Recommended)

```bash
# Build the binary (outputs to ./build/statesetd)
make build

# Or install directly to GOPATH/bin
make install
```

#### Using Go directly

```bash
# Build the binary
go build -o statesetd ./cmd/cored

# Or install it to your GOPATH/bin
go install ./cmd/cored
```

After successful compilation, you should have a `statesetd` binary in your current directory (or in your `$GOPATH/bin` if you used `go install`).

### 4. Verify the Build

```bash
# If built with make
./build/statesetd version

# If built with go build
./statesetd version

# If installed with make install or go install
statesetd version
```

### Available Make Targets

The Makefile provides several convenient targets:

```bash
make build          # Build the binary
make install        # Install to GOPATH/bin
make dev            # Quick dev setup (build + init + start)
make init           # Initialize development blockchain
make start-dev      # Start development node
make reset          # Reset blockchain data
make test           # Run tests
make deps           # Download dependencies
make clean          # Clean build directory
```

## Running the Blockchain Node

### Quick Start for Development

The fastest way to get a development blockchain running:

```bash
# Build and start a development node (automated setup)
make dev
```

This command will:
1. Build the binary
2. Initialize the blockchain with default settings
3. Create a validator key
4. Start the node

### Development Mode (Manual Setup)

For development and testing purposes, you can run a local single-node blockchain:

#### 1. Initialize the Node

```bash
# Initialize the blockchain with a moniker (name for your node)
./statesetd init my-stateset-node --chain-id stateset-1

# Create a validator key
./statesetd keys add validator

# Add genesis account
./statesetd add-genesis-account $(./statesetd keys show validator -a) 1000000000stake

# Create genesis transaction
./statesetd gentx validator 1000000stake --chain-id stateset-1

# Collect genesis transactions
./statesetd collect-gentxs
```

#### 2. Start the Node

```bash
# Start the blockchain node
./statesetd start
```

The node will start and begin producing blocks. You should see log output indicating the blockchain is running.

### Production Mode

For production deployments:

#### 1. Configuration

Edit the configuration files in `~/.statesetd/config/`:

- `config.toml`: Node configuration (P2P, RPC, consensus settings)
- `app.toml`: Application configuration (API, gRPC, state sync)
- `genesis.json`: Genesis state of the blockchain

#### 2. Key Management

```bash
# Import or create validator keys
./statesetd keys add validator --recover  # To import existing key
# OR
./statesetd keys add validator            # To create new key
```

#### 3. Join Existing Network

```bash
# Get the genesis file for the network you want to join
curl -s https://raw.githubusercontent.com/stateset/networks/main/stateset-1/genesis.json > ~/.statesetd/config/genesis.json

# Add persistent peers to config.toml
# Edit ~/.statesetd/config/config.toml and add peer nodes
```

#### 4. Start the Production Node

```bash
# Start the node
./statesetd start --home ~/.statesetd
```

## Using the Node

### Query Commands

```bash
# Check node status
./statesetd status

# Query account balance
./statesetd query bank balances <address>

# Query all validators
./statesetd query staking validators
```

### Transaction Commands

```bash
# Send tokens
./statesetd tx bank send <from_address> <to_address> <amount> --chain-id stateset-1

# Delegate to a validator
./statesetd tx staking delegate <validator_address> <amount> --from <your_key> --chain-id stateset-1

# Vote on governance proposals
./statesetd tx gov vote <proposal_id> yes --from <your_key> --chain-id stateset-1
```

## API and Services

Once your node is running, you can access various services:

- **RPC Endpoint**: `http://localhost:26657` (Tendermint RPC)
- **REST API**: `http://localhost:1317` (Cosmos SDK REST)
- **gRPC**: `localhost:9090` (Cosmos SDK gRPC)

### Enable API Services

To enable REST and gRPC APIs, edit `~/.statesetd/config/app.toml`:

```toml
[api]
enable = true
swagger = true
address = "tcp://0.0.0.0:1317"

[grpc]
enable = true
address = "0.0.0.0:9090"
```

## Docker Support (Optional)

You can also run the node using Docker:

```bash
# Build Docker image
docker build -t stateset/core .

# Run the container
docker run -it --rm \
  -p 26656:26656 \
  -p 26657:26657 \
  -p 1317:1317 \
  -p 9090:9090 \
  -v ~/.statesetd:/root/.statesetd \
  stateset/core:latest
```

## Troubleshooting

### Common Issues

1. **Build Errors**: Ensure you have Go 1.21+ installed and all dependencies are downloaded
   ```bash
   go version
   go mod download
   go mod tidy
   ```

2. **Port Conflicts**: If ports are already in use, modify the configuration in `~/.statesetd/config/config.toml` and `app.toml`

3. **Permission Issues**: Ensure the `~/.statesetd` directory has proper permissions

### Logs

View logs for debugging:
```bash
# View logs in real-time
./statesetd start --log_level info

# Or redirect to a file
./statesetd start > stateset.log 2>&1 &
```

## Contributing

Please read our [contributing guidelines](CONTRIBUTING.md) before submitting pull requests.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

---

## Available Transactions subcommands

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