# Product Service

The `product-service` is a core microservice in the **Microshop** project — a gRPC-based demo application designed with Clean Architecture principles. This service manages product-related operations in the online store, including CRUD operations, search, and caching.

## Responsibilities

The `product-service` handles the following functionalities:

* Creating, retrieving, updating, and deleting products.
* Listing all available products.
* Searching products by name or tags.
* Filtering products by a specific tag.
* Caching products in Redis and in-memory cache.
* Invalidating product cache entries.
* Sending product notification emails (e.g., on creation).
* Bulk creation of multiple products.

## Service Details

* **Port**: 50052
* **Transport**: gRPC
* **Database**: MySQL 8 (shared across Microshop services)

## Directory Structure

The service adheres to Clean Architecture and uses the following layout:

```
├── cmd/
│   └── main.go                    # Entry point
├── config/
│   ├── config.go
│   └── config_test.go
├── Dockerfile                     # Container definition
├── go.mod / go.sum                # Dependencies
├── internal/
│   ├── adapter/
│   │   ├── grpc/
│   │   │   ├── handler/
│   │   │   │   ├── handler.go
│   │   │   │   └── handler_test.go
│   │   │   └── server.go
│   │   ├── inmemory/
│   │   │   └── client.go
│   │   ├── mailer/
│   │   │   └── client.go
│   │   ├── redis/
│   │   │   └── client.go
│   │   └── nats/                  # NATS client (optional future use)
│   │       └── client.go
│   ├── app/
│   │   └── app.go                 # Dependency injection
│   ├── model/
│   │   ├── dto/
│   │   │   └── product.go
│   │   └── product.go
│   ├── repository/
│   │   ├── dao/
│   │   │   └── repository.go
│   │   └── interface.go
│   └── usecase/
│       ├── usecase.go
│       └── interface.go
├── migrations/
│   ├── 001_create_products_table.sql
│   └── 002_create_users_and_carts.sql
├── pkg/
│   ├── mysql/
│   │   └── mysql.go
│   └── redis/
│       └── redis.go
├── proto/
│   ├── product.proto              # gRPC definitions
│   └── product.pb.go             # Generated code
├── script/
│   └── trigger.sh                # gRPC testing script
```

## Setup and Quick Start

### Prerequisites

* **Go**: ≥ 1.22
* **protoc**: With `protoc-gen-go` and `protoc-gen-go-grpc`
* **MySQL 8** with schema from `/migrations`
* **grpcurl**: Install via `go install github.com/fullstorydev/grpcurl/...@latest`

### Start with Docker

```bash
docker-compose up
```

Ensure that the `product-service` image is properly built and the MySQL database is initialized with schema.

## Testing

### Run All Unit Tests

```bash
go test -v ./...
```

### Test gRPC Endpoints

Trigger all endpoints using the helper script:

```bash
./script/trigger.sh CreateProduct
```

You can call other methods by name, e.g.:

```bash
./script/trigger.sh ListProducts
```

Ensure the service is running and `grpcurl` is installed.

## gRPC Endpoints

The following gRPC endpoints are exposed in `proto/product.proto`:

* `CreateProduct`
* `GetProduct`
* `UpdateProduct`
* `DeleteProduct`
* `ListProducts`
* `SearchProducts`
* `GetProductsByTag`
* `SetProductCache`
* `InvalidateProductCache`
* `SendProductEmail`
* `GetAllFromRedis`
* `GetAllFromCache`
* `BulkCreateProducts`

## Dependencies

* **Database**: MySQL 8 (tables: products, users, carts)
* **Mailer**: Mailjet API
* **Cache**: Redis + in-memory Go map
* **gRPC**: For external communication

## Notes

* Ensure that foreign key constraints (`ON DELETE CASCADE`) are set to avoid deletion errors.
* The trigger script assumes gRPC reflection is enabled. If not, update the script to use `.proto` definitions directly.
* The service is built to scale horizontally and integrates smoothly with analytics and user services in the Microshop project.

## About Microshop

The `product-service` is part of **Microshop**, a microservices-based demo project for an online store. The ecosystem includes:

* **User Service** (port 50051)
* **Product Service** (port 50052)
* **Cart Service** (port 50053)
* **Analytics Service** (port 50054)

Explore the entire project to see how gRPC, Clean Architecture, and microservices work in harmony! 🌟
