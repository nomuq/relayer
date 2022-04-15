#!/bin/sh

# check if protoc is installed
if ! which protoc > /dev/null; then
  echo "Protobuf compiler not found in PATH."
fi

# Update PATH so that the protoc compiler can find the plugins
export PATH="$PATH:$(go env GOPATH)/bin"

# check if protoc-gen-go is installed
if ! which protoc-gen-go > /dev/null; then
  echo "Protobuf Go compiler not found in PATH."
  echo "Installing it..."
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
fi

# remove pkg/relayer directory if it exists
if [ -d pkg/proto ]; then
    rm -rf pkg/proto
fi

mkdir -p pkg/proto

protoc --proto_path=protos --go_out=pkg/proto --go_opt=paths=source_relative \
    --go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative \
    protos/*.proto