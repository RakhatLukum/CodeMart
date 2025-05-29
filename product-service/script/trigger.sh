#!/bin/bash

# Configuration
HOST="localhost:50052"
PROTO_PATH="proto/product.proto" # Update with actual path to proto file
SERVICE="product.ProductService"

# Check if grpcurl is installed
if ! command -v grpcurl &> /dev/null; then
    echo "grpcurl is not installed. Please install it using 'go install github.com/fullstorydev/grpcurl/...@latest'"
    exit 1
fi

# Function to execute grpcurl command
call_grpc() {
    local method=$1
    local data=$2
    echo "Calling $method..."
    if [ -z "$data" ]; then
        grpcurl -proto "$PROTO_PATH" -plaintext -d "{}" $HOST $SERVICE/$method
    else
        grpcurl -proto "$PROTO_PATH" -plaintext -d "$data" $HOST $SERVICE/$method
    fi
    if [ $? -eq 0 ]; then
        echo "Success: $method"
    else
        echo "Error: Failed to call $method"
    fi
    echo "----------------------------------------"
}

# 1. CreateProduct
call_grpc "CreateProduct" '{"name": "Test Product", "price": 99.99, "tags": "test"}'

# 2. GetProduct
call_grpc "GetProduct" '{"id": 1}'

# 3. UpdateProduct
call_grpc "UpdateProduct" '{"id": 1, "name": "Updated Product", "price": 149.99, "tags": "updated"}'

# 4. DeleteProduct
call_grpc "DeleteProduct" '{"id": 1}'

# 5. ListProducts
call_grpc "ListProducts" ''

# 6. SearchProducts
call_grpc "SearchProducts" '{"query": "test", "tags": "test"}'

# 7. GetProductsByTag
call_grpc "GetProductsByTag" '{"tag": "test"}'

# 8. SetProductCache
call_grpc "SetProductCache" '{"id": 1, "name": "Cached Product", "price": 99.99, "tags": "cache"}'

# 9. InvalidateProductCache
call_grpc "InvalidateProductCache" '{"id": 1}'

# 10. SendProductEmail
call_grpc "SendProductEmail" '{"product_id": 1, "email": "test@example.com"}'

# 11. GetAllFromRedis
call_grpc "GetAllFromRedis" ''

# 12. GetAllFromCache
call_grpc "GetAllFromCache" ''

# 13. BulkCreateProducts
call_grpc "BulkCreateProducts" '{"products": [{"name": "Bulk Product 1", "price": 49.99, "tags": "bulk"}, {"name": "Bulk Product 2", "price": 59.99, "tags": "bulk"}]}'

echo "All endpoints triggered successfully!"