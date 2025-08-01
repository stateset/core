#!/bin/bash

echo "Fixing all import issues..."

# Fix IBC v2 to v8
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i 's#github\.com/cosmos/ibc-go/v2#github.com/cosmos/ibc-go/v8#g' {} \;

# Fix simapp imports
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i 's#cosmossdk\.io/simapp#github.com/cosmos/cosmos-sdk/simapp#g' {} \;

# Remove unused imports - cctp/client/cli
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i '/github\.com\/stateset\/core\/x\/cctp\/client\/cli/d' {} \;

# Remove gogo/protobuf/grpc import if not used
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i '/github\.com\/gogo\/protobuf\/grpc/d' {} \;

# Remove core/appconfig and core/coins imports
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i '/cosmossdk\.io\/core\/appconfig/d' {} \;
find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.git/*" -exec sed -i '/cosmossdk\.io\/core\/coins/d' {} \;

echo "Import fixes completed"