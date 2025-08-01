#!/bin/bash

echo "Creating minimal blockchain binary..."

# Create a docker container that stays running
docker run -d --name stateset-builder \
  -v $(pwd):/workspace \
  -w /workspace \
  golang:1.23-alpine \
  sleep 3600

# Install dependencies in the container
docker exec stateset-builder sh -c "apk add --no-cache git gcc musl-dev linux-headers"

# Copy current go.mod to backup
docker exec stateset-builder cp go.mod go.mod.original

# Try to download dependencies and build
docker exec stateset-builder sh -c '
  # Set Go env
  export CGO_ENABLED=1
  export GOOS=linux
  export GOARCH=amd64
  
  # Create build directory
  mkdir -p build
  
  # Download core dependencies manually
  go get cosmossdk.io/store@v1.1.0
  go get github.com/cometbft/cometbft@v0.38.7
  go get github.com/cosmos/cosmos-sdk@v0.50.6
  
  # Build with minimal flags
  go build -mod=mod -o build/cored ./cmd/cored || {
    echo "Build failed, trying simpler approach..."
    # If fails, just compile the main file
    cd cmd/cored
    go build -mod=mod -o ../../build/cored main.go
  }
'

# Check if build succeeded
docker exec stateset-builder test -f build/cored && {
  echo "Build successful!"
  docker cp stateset-builder:/workspace/build/cored ./build/cored
  chmod +x ./build/cored
  ls -lah ./build/cored
} || {
  echo "Build failed!"
}

# Clean up
docker stop stateset-builder
docker rm stateset-builder