#!/bin/bash

# Папки
AUTH_PROTO_DIR="/path/to/your/project/backend/proto/auth"
MARKETPLACE_PROTO_DIR="/path/to/your/project/backend/proto/marketplace"
INCLUDE_DIR="/path/to/your/project/backend/proto/include"
OUT_DIR="/path/to/your/project/backend/proto/protos_gen"
AUTH_OUT_DIR="$OUT_DIR/auck.authservice1.v1"
MARKETPLACE_OUT_DIR="$OUT_DIR/auck.marketplaceservice1.v1"

# Генерация для Auth
protoc \
  -I="$AUTH_PROTO_DIR" \
  -I="$INCLUDE_DIR" \
  --go_out="$AUTH_OUT_DIR" --go_opt=paths=source_relative \
  --go-grpc_out="$AUTH_OUT_DIR" --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out="$AUTH_OUT_DIR" --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out="$OUT_DIR" \
  "$AUTH_PROTO_DIR"/*.proto

# Генерация для Marketplace
protoc \
  -I="$MARKETPLACE_PROTO_DIR" \
  -I="$INCLUDE_DIR" \
  --go_out="$MARKETPLACE_OUT_DIR" --go_opt=paths=source_relative \
  --go-grpc_out="$MARKETPLACE_OUT_DIR" --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out="$MARKETPLACE_OUT_DIR" --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out="$OUT_DIR" \
  "$MARKETPLACE_PROTO_DIR"/*.proto

echo
echo "Done"
