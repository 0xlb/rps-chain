#!/usr/bin/env bash

cd ..

# Get the RPS chain protos
mkdir -p proto/lb/rps
cp -r ../proto/lb/rps/v1 ./proto/lb/rps

# Get Cosmos-SDK protos
mkdir -p proto/cosmos/query/v1
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.50.4/proto/cosmos/query/v1/query.proto -o proto/cosmos/query/v1/query.proto

mkdir -p proto/cosmos/msg/v1
curl https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.50.4/proto/cosmos/msg/v1/msg.proto -o proto/cosmos/msg/v1/msg.proto

mkdir -p proto/cosmos_proto
curl https://raw.githubusercontent.com/cosmos/cosmos-proto/v1.0.0-beta.4/proto/cosmos_proto/cosmos.proto -o proto/cosmos_proto/cosmos.proto

# Get dependencies protos
mkdir -p proto/google/api
curl https://raw.githubusercontent.com/googleapis/googleapis/common-protos-1_3_1/google/api/annotations.proto -o proto/google/api/annotations.proto
curl https://raw.githubusercontent.com/googleapis/googleapis/common-protos-1_3_1/google/api/http.proto -o proto/google/api/http.proto

mkdir -p proto/google/protobuf
curl https://raw.githubusercontent.com/protocolbuffers/protobuf/v22.2/src/google/protobuf/descriptor.proto -o proto/google/protobuf/descriptor.proto

mkdir -p proto/gogoproto
curl https://raw.githubusercontent.com/cosmos/gogoproto/v1.4.11/gogoproto/gogo.proto -o proto/gogoproto/gogo.proto

mkdir -p src/types/generated
ls ../proto/lb/rps/v1 | xargs -I {} protoc \
    --plugin="./node_modules/.bin/protoc-gen-ts_proto" \
    --ts_proto_out="./src/types/generated" \
    --proto_path="./proto" \
    --ts_proto_opt="esModuleInterop=true,forceLong=long,useOptionals=messages" \
    lb/rps/v1/{}