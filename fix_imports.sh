#!/bin/bash

# Fix tendermint imports to cometbft
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i \
  -e 's|github.com/tendermint/tendermint/abci/types|github.com/cometbft/cometbft/abci/types|g' \
  -e 's|github.com/tendermint/tendermint/libs/log|github.com/cometbft/cometbft/libs/log|g' \
  -e 's|github.com/tendermint/tendermint/libs/cli|github.com/cometbft/cometbft/libs/cli|g' \
  -e 's|github.com/tendermint/tendermint/libs/json|github.com/cometbft/cometbft/libs/json|g' \
  -e 's|github.com/tendermint/tendermint/libs/os|github.com/cometbft/cometbft/libs/os|g' \
  -e 's|github.com/tendermint/tendermint/crypto|github.com/cometbft/cometbft/crypto|g' \
  -e 's|github.com/tendermint/tendermint/config|github.com/cometbft/cometbft/config|g' \
  -e 's|github.com/tendermint/tendermint/types|github.com/cometbft/cometbft/types|g' \
  -e 's|github.com/tendermint/tendermint/proto/tendermint|github.com/cometbft/cometbft/proto/tendermint|g' \
  -e 's|github.com/tendermint/tendermint/libs/rand|github.com/cometbft/cometbft/libs/rand|g' \
  -e 's|github.com/tendermint/tm-db|github.com/cometbft/cometbft-db|g' \
  {} \;

# Fix cosmos-sdk store imports
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i \
  -e 's|github.com/cosmos/cosmos-sdk/store/types|cosmossdk.io/store/types|g' \
  {} \;

echo "Import fixes completed"