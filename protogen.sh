#!/bin/sh

# remove pkg/relayer directory if it exists
if [ -d pkg/proto ]; then
    rm -rf pkg/proto
fi

mkdir -p pkg/proto

protoc --proto_path=. --go_out=pkg/proto --go_opt=paths=source_relative \
    --go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative \
    relayer.proto