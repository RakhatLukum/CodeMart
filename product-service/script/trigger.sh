#!/bin/bash

# Configuration
HOST="localhost:50052"
SERVICE="product.ProductService"

# Check if grpcurl is installed
if ! command -v grpcurl &> /dev/null; then
    echo "grpcurl is not installed. Please install it using 'go install github.com/fullstorydev/grpcurl/...@latest'"
    exit 1
fi

# Map of methods to JSON payloads
declare -A PAYLOADS

PAYLOADS["CreateProduct"]='{"name": "Test Product", "price": 99.99, "tags": "test"}'
PAYLOADS["GetProduct"]='{"id": 1}'
PAYLOADS["UpdateProduct"]='{"id": 1, "name": "Updated Product", "price": 149.99, "tags": "updated"}'
PAYLOADS["DeleteProduct"]='{"id": 1}'
PAYLOADS["ListProducts"]=""
PAYLOADS["SearchProducts"]='{"query": "test", "tags": "test"}'
PAYLOADS["GetProductsByTag"]='{"tag": "test"}'
PAYLOADS["SetProductCache"]='{"id": 1, "name": "Cached Product", "price": 99.99, "tags": "cache"}'
PAYLOADS["InvalidateProductCache"]='{"id": 1}'
PAYLOADS["SendProductEmail"]='{"product_id": 2, "email": "test@example.com"}'
PAYLOADS["GetAllFromRedis"]=""
PAYLOADS["GetAllFromCache"]=""
PAYLOADS["BulkCreateProducts"]='{"products": [{"name": "Bulk Product 1", "price": 49.99, "tags": "bulk"}, {"name": "Bulk Product 2", "price": 59.99, "tags": "bulk"}]}'

# Handle input
if [ -z "$1" ]; then
    echo "Usage: $0 <MethodName>"
    echo "Available methods:"
    for m in "${!PAYLOADS[@]}"; do echo "  - $m"; done
    exit 1
fi

METHOD="$1"
DATA="${PAYLOADS[$METHOD]}"

# Check for unknown method
if [ -z "${PAYLOADS[$METHOD]+_}" ]; then
    echo "Unknown method: $METHOD"
    exit 1
fi

echo "Calling $METHOD..."

if [ -z "$DATA" ]; then
    grpcurl -plaintext "$HOST" "$SERVICE/$METHOD"
else
    grpcurl -plaintext -d "$DATA" "$HOST" "$SERVICE/$METHOD"
fi

echo "Done ✅"
