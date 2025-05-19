# Analytics Service

The `analytics-service` is a microservice within the **Microshop** project, a gRPC-based demo application built with Clean Architecture principles. This service is responsible for tracking and analyzing product view statistics, providing insights into user interactions with products in the online grocery store.

## Responsibilities

The `analytics-service` handles the following functionalities:
- Recording product views by users.
- Retrieving views by user, product, or user-product combinations.
- Fetching recent views with a configurable limit.
- Identifying the most viewed products overall and per user.
- Counting views for specific products or users.
- Generating daily and hourly view statistics.
- Sending daily view reports via email.
- Deleting old views based on a timestamp.
- Caching view data in both persistent and in-memory stores.

## Service Details

- **Port**: 50054
- **Transport**: gRPC
- **Database**: MySQL 8 (shared with other services)

## Directory Structure

The service follows Clean Architecture principles, with the following layout:

```
├── cmd/                            # Application entry point
│   └── main.go                     # Main function to start the gRPC server
├── config/                         # Configuration management
│   ├── config.go                   # Loads and parses environment or config files
│   └── config_test.go              # Tests for configuration logic
├── Dockerfile                      # Dockerfile for containerizing the application
├── go.mod                          # Go module definition
├── go.sum                          # Module dependency checksums
├── internal/                       # Internal application logic (not exposed publicly)
│   ├── adapter/                    # External service adapters
│   │   ├── grpc/                   # gRPC transport layer
│   │   │   ├── handler/            # gRPC handler implementations
│   │   │   │   ├── handler.go
│   │   │   │   ├── handler_test.go
│   │   │   │   └── interface.go    # Interface for handler abstraction
│   │   │   ├── server.go           # gRPC server setup
│   │   │   └── server_test.go
│   │   ├── inmemory/               # In-memory implementations for testing/mocking
│   │   │   ├── client.go
│   │   │   └── client_test.go
│   │   ├── mailer/                 # Mail sending client
│   │   │   └── client.go
│   │   ├── nats/                   # NATS messaging client
│   │   │   └── client.go
│   │   └── redis/                  # Redis client for caching
│   │       └── client.go
│   ├── app/                        # Application orchestration
│   │   └── app.go
│   ├── model/                      # Domain models and DTOs
│   │   ├── cart.go
│   │   ├── dto/                    # Data transfer objects
│   │   │   └── view.go
│   │   ├── product.go
│   │   ├── user.go
│   │   └── view.go
│   ├── repository/                # Database repositories and interfaces
│   │   ├── dao/                    # Data access object implementations
│   │   │   ├── repository.go
│   │   │   └── repository_test.go
│   │   ├── interface.go
│   │   ├── repository.go
│   │   └── repository_test.go
│   └── usecase/                   # Business logic
│       ├── interface.go
│       ├── usecase.go
│       └── usecase_test.go
├── migrations/                    # SQL migration scripts
│   ├── 001_create_views_table.sql
│   ├── 002_add_indexes_to_views.sql
│   ├── 003_create_aggregated_product_views.sql
│   ├── 004_create_user_view_counts.sql
│   ├── 005_create_daily_view_stats.sql
│   ├── 006_create_hourly_view_stats.sql
│   └── 007_create_recent_views.sql
├── pkg/                           # Reusable helper packages
│   ├── mysql/
│   │   ├── mysql.go
│   │   └── mysql_test.go
│   ├── nats/
│   │   ├── nats.go
│   │   └── nats_test.go
│   └── redis/
│       ├── redis.go
│       └── redis_test.go
├── proto/                         # Protobuf definitions and generated files
│   ├── analytics_grpc.pb.go
│   ├── analytics.pb.go
│   └── analytics.proto
├── script/                        # Utility scripts
│   └── trigger.sh                 # Script to invoke all gRPC endpoints
├── test.txt                       # Possibly temp/test file
└── TODO                           # Task or feature tracking
```

## Setup and Quick Start

### Prerequisites

- **Go**: ≥ 1.22
- **protoc**: With plugins `protoc-gen-go` and `protoc-gen-go-grpc`
- **MySQL**: Version 8
- **grpcurl**: For testing endpoints (install via `go install github.com/fullstorydev/grpcurl/...@latest`)

### Steps

1. **Clone the Repository**

   Ensure you have the Microshop repository cloned:
   ```bash
   git clone https://github.com/RakhatLukum/CodeMart.git
   cd analytics-service
   ```

2. **Run the Service**

   Start the analytics service:
   ```bash
   docker-compose up
   ```

## Testing

### Run Unit Tests

Execute all tests for the service:
```bash
go test -v ./...
```

### Trigger All Endpoints

Use the provided Bash script to call all gRPC endpoints on port 50054:
```bash
./script/trigger.sh
```

Ensure `grpcurl` is installed and the server is running. The script sends sample requests to endpoints like `CreateView`, `GetViewsByUser`, `GetMostViewedProducts`, and others. Modify the script's request data in `script/trigger.sh` for specific test cases.

## gRPC Endpoints

The service exposes the following gRPC endpoints (defined in `proto/view.proto`):

- `CreateView`: Records a new product view.
- `GetViewsByUser`: Retrieves views for a specific user.
- `GetViewsByProduct`: Retrieves views for a specific product.
- `GetViewsByUserAndProduct`: Retrieves views for a user-product pair.
- `GetRecentViews`: Fetches the most recent views (with a limit).
- `GetMostViewedProducts`: Lists products with the highest view counts.
- `GetUserTopProducts`: Lists top products viewed by a specific user.
- `GetProductViewCount`: Returns the total views for a product.
- `GetUserViewCount`: Returns the total views by a user.
- `GetDailyViews`: Provides daily view statistics.
- `GenerateDailyViewReportEmail`: Sends a daily view report to an email.
- `GetHourlyViews`: Provides hourly view statistics for a product.
- `DeleteOldViews`: Deletes views before a specified timestamp.
- `GetCachedView`: Retrieves a cached view for a product.
- `GetMemoryCachedView`: Retrieves a view from in-memory cache.

Refer to `proto/view.proto` for request/response schemas.

## Dependencies

- **Go Modules**: Managed via `go.mod` (e.g., `google.golang.org/grpc`, `github.com/go-sql-driver/mysql`).
- **Database**: MySQL 8 for persistent storage.
- **Caching**: Supports both persistent and in-memory caching (implementation details in `usecase`).

## Notes

- The service integrates with other Microshop services (user, product, cart) via the shared MySQL database but operates independently over gRPC.
- Ensure the MySQL schema (`database/schema.sql`) and seed data (`database/seed.sql`) are applied before running.
- The `trigger.sh` script assumes the server supports gRPC reflection. If not, update the script to include the proto file path.

## About Microshop

The `analytics-service` is part of **Microshop**, an online grocery store demo showcasing microservices with gRPC and Clean Architecture. The project includes four services: user (port 50051), product (port 50052), cart (port 50053), and analytics (port 50054).
