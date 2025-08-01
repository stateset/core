#!/bin/bash

# Run go mod tidy in a container with Go 1.23
docker run --rm -v $(pwd):/app -w /app golang:1.23-alpine sh -c "
  apk add --no-cache git gcc musl-dev linux-headers && \
  go mod tidy -compat=1.21 && \
  go mod download && \
  mkdir -p build && \
  go build -o build/cored ./cmd/cored
"

# Check if build succeeded
if [ -f build/cored ]; then
  echo "Build successful! Binary is at build/cored"
  ls -lah build/cored
else
  echo "Build failed!"
  exit 1
fi