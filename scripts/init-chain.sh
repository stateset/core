#!/bin/sh

# Initialize the chain with custom stateset prefix
echo "Initializing StateSet blockchain with stateset prefix..."

# Set custom bech32 prefixes
export WASMD_BECH32_PREFIX="stateset"

# Initialize the chain
wasmd init stateset-node --chain-id stateset-1

# Update genesis configuration
echo "Configuring genesis file..."

# Update staking denomination to stst
sed -i 's/"bond_denom": "stake"/"bond_denom": "stst"/' /root/.wasmd/config/genesis.json

# Update mint denomination to stst  
sed -i 's/"mint_denom": "stake"/"mint_denom": "stst"/' /root/.wasmd/config/genesis.json

# Update crisis denomination to stst
sed -i 's/"denom": "stake"/"denom": "stst"/g' /root/.wasmd/config/genesis.json

# Update bech32 prefixes in genesis
jq '.app_state.auth.params.bech32_prefix_account = "stateset"' /root/.wasmd/config/genesis.json > /tmp/genesis.json && mv /tmp/genesis.json /root/.wasmd/config/genesis.json
jq '.app_state.auth.params.bech32_prefix_validator = "statesetvaloper"' /root/.wasmd/config/genesis.json > /tmp/genesis.json && mv /tmp/genesis.json /root/.wasmd/config/genesis.json
jq '.app_state.auth.params.bech32_prefix_consensus = "statesetvalcons"' /root/.wasmd/config/genesis.json > /tmp/genesis.json && mv /tmp/genesis.json /root/.wasmd/config/genesis.json

# Create validator key with custom prefix
WASMD_BECH32_PREFIX=stateset wasmd keys add validator --keyring-backend test

# Get validator address
VALIDATOR_ADDR=$(WASMD_BECH32_PREFIX=stateset wasmd keys show validator -a --keyring-backend test)
echo "Validator address: $VALIDATOR_ADDR"

# Add genesis account with tokens
WASMD_BECH32_PREFIX=stateset wasmd genesis add-genesis-account $VALIDATOR_ADDR 100000000000stst --keyring-backend test

# Create genesis transaction
WASMD_BECH32_PREFIX=stateset wasmd genesis gentx validator 100000000stst --chain-id stateset-1 --keyring-backend test

# Collect genesis transactions
wasmd genesis collect-gentxs

# Configure for development
sed -i 's/minimum-gas-prices = ""/minimum-gas-prices = "0.025stst"/' /root/.wasmd/config/app.toml
sed -i 's/enable = false/enable = true/' /root/.wasmd/config/app.toml
sed -i 's/swagger = false/swagger = true/' /root/.wasmd/config/app.toml

echo "Chain initialization complete!"