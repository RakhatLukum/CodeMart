# Cart Service

The `cart-service` is a core microservice in the **Microshop** project — a gRPC-based application following Clean Architecture. This service manages users’ shopping carts, offering flexible and efficient operations backed by Redis, MySQL, and in-memory caching.

## Responsibilities

The `cart-service` is responsible for:

* Adding, removing, and clearing cart items.
* Updating items in a user’s cart.
* Fetching cart contents and item count.
* Calculating total cart value.
* Sending cart summary emails.
* Providing cart data from Redis and in-memory cache.
* Supporting cache invalidation.

## Service Details

* **Port**: `50053`
* **Transport**: `gRPC`
* **Database**: `MySQL 8`

## Directory Structure

```
├── cmd/
│   └── main.go                    # Entry point
├── config/
│   ├── config.go
│   └── config_test.go
├── Dockerfile
├── go.mod / go.sum
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
│   │   └── nats/
│   │       └── client.go
│   ├── app/
│   │   └── app.go
│   ├── model/
│   │   ├── dto/
│   │   │   └── cart.go
│   │   └── cart.go
│   ├── repository/
│   │   ├── dao/
│   │   │   └── repository.go
│   │   └── interface.go
│   └── usecase/
│       ├── usecase.go
│       └── interface.go
├── migrations/
│   ├── 001_create_products_table.sql
│   ├── 002_create_users_table.sql
│   ├── 003_create_carts_table.sql
│   └── 004_add_cart_foreign_keys.sql
├── pkg/
│   ├── mysql/
│   │   └── mysql.go
│   └── redis/
│       └── redis.go
├── proto/
│   ├── cart.proto
│   └── cart.pb.go
├── script/
│   └── trigger.sh
```

## Setup and Quick Start

### Requirements

* **Go**: ≥ 1.22
* **protoc** with `protoc-gen-go` and `protoc-gen-go-grpc`
* **MySQL 8**
* **grpcurl**

### Start via Docker

```bash
docker-compose up
```

Make sure the database schema is initialized from `/migrations`.

## Testing

### Run Unit Tests

```bash
go test -v ./...
```

### Call gRPC Methods

```bash
./script/trigger.sh AddToCart
```

Use other methods like:

```bash
./script/trigger.sh GetCart
```

## gRPC Endpoints

The following endpoints are defined in `proto/cart.proto`:

* `AddToCart`
* `RemoveFromCart`
* `ClearCart`
* `GetCart`
* `UpdateCartItem`
* `HasProductInCart`
* `GetCartItemCount`
* `GetCartTotalPrice`
* `SendCartSummaryEmail`
* `InvalidateCartCache`
* `GetAllFromRedis`
* `GetAllFromCache`

## Dependencies

* **Database**: MySQL 8 (`products`, `users`, `carts`)
* **Cache**: Redis + Go in-memory map
* **Email**: Mailjet API
* **NATS**: For pub/sub events (optional)
* **gRPC**: External service communication

## Notes

* Cache fallback is supported between Redis and memory.
* `SendCartSummaryEmail` sends a digest of current cart contents.
* Ensure `product-service` is running for accurate pricing info.
* The service supports horizontal scaling via stateless logic.

## About Microshop

This service is part of **Microshop**, a gRPC-based microservices e-commerce system. It works alongside:

* **User Service** – Port `50051`
* **Product Service** – Port `50052`
* **Cart Service** – Port `50053`
* **Analytics Service** – Port `50054`

Let your services talk like pros — via gRPC! 
