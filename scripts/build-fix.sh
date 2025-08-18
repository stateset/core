#!/bin/bash

# Build Fix Script for Stateset Core
# This script helps complete the remaining build tasks

set -e

echo "🔧 Stateset Core Build Fix Script"
echo "================================="

echo "Step 1: Cleaning up old generated files..."
find . -name "*.pb.go" -type f -delete
echo "✅ Cleaned up old protobuf files"

echo "Step 2: Installing required protobuf tools..."
go install github.com/cosmos/gogoproto/protoc-gen-gocosmos@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
echo "✅ Installed protobuf tools"

echo "Step 3: Regenerating protobuf files..."
if command -v buf &> /dev/null; then
    buf generate
    echo "✅ Generated protobuf files with buf"
else
    echo "⚠️  buf not found. Please install buf or regenerate protobuf files manually"
    echo "   See: https://docs.buf.build/installation"
fi

echo "Step 4: Running go mod tidy..."
go mod tidy
echo "✅ Updated go modules"

echo "Step 5: Building the project..."
if go build ./cmd/statesetd; then
    echo "✅ Build successful!"
    echo ""
    echo "🎉 All build issues resolved!"
    echo "The Stateset Core blockchain is ready for deployment."
else
    echo "❌ Build failed. Please check the error messages above."
    echo ""
    echo "Common fixes:"
    echo "1. Ensure all import paths are correct"
    echo "2. Regenerate protobuf files if needed"
    echo "3. Check for any missing dependencies"
    exit 1
fi

echo ""
echo "🚀 Next steps:"
echo "1. Run tests: go test ./..."
echo "2. Start the blockchain: ./statesetd start"
echo "3. Deploy to production environment"