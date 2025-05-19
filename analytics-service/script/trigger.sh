#!/bin/bash

# Configuration
HOST="localhost:50054"
SERVICE="view.ViewService"

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
        grpcurl -plaintext -d "{}" $HOST $SERVICE/$method
    else
        grpcurl -plaintext -d "$data" $HOST $SERVICE/$method
    fi
    echo "----------------------------------------"
}

# 1. CreateView
call_grpc "CreateView" '{"user_id": 1, "product_id": 1}'

# 2. GetViewsByUser
call_grpc "GetViewsByUser" '{"user_id": 1}'

# 3. GetViewsByProduct
call_grpc "GetViewsByProduct" '{"product_id": 1}'

# 4. GetViewsByUserAndProduct
call_grpc "GetViewsByUserAndProduct" '{"user_id": 1, "product_id": 1}'

# 5. GetRecentViews
call_grpc "GetRecentViews" '{"limit": 10}'

# 6. GetMostViewedProducts
call_grpc "GetMostViewedProducts" ''

# 7. GetUserTopProducts
call_grpc "GetUserTopProducts" '{"user_id": 1, "limit": 5}'

# 8. GetProductViewCount
call_grpc "GetProductViewCount" '{"product_id": 1}'

# 9. GetUserViewCount
call_grpc "GetUserViewCount" '{"user_id": 1}'

# 10. GetDailyViews
call_grpc "GetDailyViews" ''

# 11. GenerateDailyViewReportEmail
call_grpc "GenerateDailyViewReportEmail" '{"email": "test@example.com", "name": "Test User"}'

# 12. GetHourlyViews
call_grpc "GetHourlyViews" '{"product_id": 1}'

# 13. DeleteOldViews
call_grpc "DeleteOldViews" '{"before": "2025-01-01T00:00:00Z"}'

# 14. GetCachedView
call_grpc "GetCachedView" '{"product_id": 1}'

# 15. GetMemoryCachedView
call_grpc "GetMemoryCachedView" '{"product_id": 1}'

echo "All endpoints triggered successfully!"