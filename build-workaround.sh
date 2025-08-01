#!/bin/bash

# Create a temporary directory for the build
BUILD_DIR=$(mktemp -d)
echo "Building in temporary directory: $BUILD_DIR"

# Copy necessary files
cp -r cmd $BUILD_DIR/
cp -r app $BUILD_DIR/
cp -r x $BUILD_DIR/
cp -r docs $BUILD_DIR/ 2>/dev/null || true
cp -r config $BUILD_DIR/ 2>/dev/null || true
cp go.mod $BUILD_DIR/
cp go.sum $BUILD_DIR/

# Run build in Docker with minimal dependencies
docker run --rm -v $BUILD_DIR:/app -w /app golang:1.23-alpine sh -c '
  set -e
  
  # Install dependencies
  apk add --no-cache git gcc musl-dev linux-headers
  
  # Add build tags to avoid test dependencies
  export CGO_ENABLED=1
  export GOFLAGS="-buildvcs=false"
  
  # Try to build without updating dependencies
  mkdir -p build
  go build -mod=readonly -tags "netgo" -ldflags "-w -s" -o build/cored ./cmd/cored || {
    echo "Direct build failed, trying with vendor..."
    
    # If that fails, try vendoring
    go mod vendor
    go build -mod=vendor -tags "netgo" -ldflags "-w -s" -o build/cored ./cmd/cored
  }
'

# Copy the binary back if successful
if [ -f "$BUILD_DIR/build/cored" ]; then
  cp -f "$BUILD_DIR/build/cored" ./build/cored
  echo "Build successful! Binary copied to ./build/cored"
  chmod +x ./build/cored
  ls -lah ./build/cored
else
  echo "Build failed!"
  exit 1
fi

# Clean up
rm -rf $BUILD_DIR