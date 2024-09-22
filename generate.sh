#!/bin/bash

rm -rf ./api
rm -rf ./core

# Check and install dependencies
if ! command -v protoc-gen-go &>/dev/null; then
  echo "Installing protoc-gen-go..."
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
else
  echo "protoc-gen-go is already installed"
fi

if ! command -v protoc-gen-go-grpc &>/dev/null; then
  echo "Installing protoc-gen-go-grpc..."
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
else
  echo "protoc-gen-go-grpc is already installed"
fi

# Generate code
protoc -I=./protocol -I./googleapis --go_out=. --go-grpc_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  ./protocol/api/*.proto ./protocol/core/*.proto ./protocol/core/contract/*.proto

# Replace import paths
find . -name "*.pb.go" -type f -exec sed -i 's|github.com/tronprotocol/grpc-gateway|github.com/gofika/fikatron|g' {} +

# Move files and remove empty directories
mv ./core/contract/* ./core/
rm -rf ./core/contract

echo "Installation and generation steps completed"
