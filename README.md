# 🛍️ CodeMart – gRPC Microservices with Clean Architecture

**CodeMart** is a demo e-commerce platform composed of microservices communicating over **gRPC**, organized using **Clean Architecture**. It includes a full product catalog, cart management, and analytics tracking — all running with Docker Compose and easily extendable.

---

## Services Overview

| Service             | gRPC Port      | Description                                      |
| ------------------- | -------------- | ------------------------------------------------ |
| `product-service`   | `50052`        | Handles products: create, list, cache, filter    |
| `cart-service`      | `50053`        | Manages user carts and email summaries           |
| `analytics-service` | `50054`        | Tracks and analyzes product view interactions    |
| `api-gateway`       | `50050` (HTTP) | Forwards HTTP requests to internal gRPC services |

---

## Quick Start

### 1. Clone Repository

```bash
git clone https://github.com/RakhatLukum/CodeMart.git
cd CodeMart
```

### 2. Start All Services

All dependencies (MySQL, Redis, NATS) and services are included:

```bash
docker-compose up --build
```

> This runs everything: `product-service`, `cart-service`, `analytics-service`, `api-gateway`, `mysql`, `redis`, `nats`, and `adminer`.

### 3. Test Endpoints

Use Postman, `curl`, or your browser to interact via:

```
http://localhost:50050/api/v1/
```

Example requests:

```bash
# Add item to cart
curl -X POST http://localhost:50050/api/v1/cart \
  -H "Content-Type: application/json" \
  -d '{"user_id": 1, "product_id": 2}'

# Get user's cart
curl http://localhost:50050/api/v1/cart?user_id=1
```

---

## Directory Structure

```
<service>/
│
├── cmd/                # Entry point
│   └── main.go
├── config/             # Configuration loader
├── internal/
│   ├── adapter/        # Redis, gRPC handlers, NATS, Mailjet, memory
│   ├── app/            # Service setup (wiring)
│   ├── model/          # Domain models & DTOs
│   ├── repository/     # Interfaces + SQL implementation
│   └── usecase/        # Business logic
├── migrations/         # SQL migration files
├── proto/              # gRPC .proto and generated .pb.go files
├── Dockerfile          # Docker service build
```

---

## Protobuf Compilation

Each service has its own `.proto` definition.

Example for `cart-service`:

```bash
cd cart-service
protoc --go_out=./proto --go-grpc_out=./proto \
  -I./proto proto/cart.proto
```
---

## Environment Variables (Used in Docker Compose)

```env
MYSQL_DSN=root:MyStrongPassword123!@tcp(mysql:3306)/shop?parseTime=true

REDIS_ADDR=redis:6379
REDIS_DB=0
REDIS_TTL_SECONDS=86400

NATS_URL=nats://nats:4222

MAILJET_API_KEY=<your-key>
MAILJET_SECRET_KEY=<your-secret>
MAILJET_SENDER_EMAIL=your@email.com
MAILJET_SENDER_NAME=CodeMart

GRPC_PORT=50053
```

> Defined per service under `environment:` in `docker-compose.yml`


---

## gRPC Endpoints

### cart-service

* `AddToCart(CreateCartRequest) returns (CreateCartResponse)`
* `RemoveFromCart(DeleteCartItemRequest) returns (DeleteCartItemResponse)`
* `ClearCart(UserIDRequest) returns (google.protobuf.Empty)`
* `GetCart(UserIDRequest) returns (CartListResponse)`
* `GetCartItems(UserIDRequest) returns (CartItemsResponse)`
* `UpdateCartItem(UpdateCartItemRequest) returns (google.protobuf.Empty)`
* `HasProductInCart(HasProductInCartRequest) returns (HasProductInCartResponse)`
* `GetCartItemCount(UserIDRequest) returns (CartItemCountResponse)`
* `GetCartTotalPrice(UserIDRequest) returns (CartTotalPriceResponse)`
* `SendCartSummaryEmail(SendCartSummaryEmailRequest) returns (EmailStatusResponse)`
* `InvalidateCartCache(UserIDRequest) returns (CacheResponse)`
* `GetAllFromRedis(google.protobuf.Empty) returns (CartListResponse)`
* `GetAllFromCache(google.protobuf.Empty) returns (CartListResponse)`

---

### product-service

* `CreateProduct(CreateProductRequest) returns (CreateProductResponse)`
* `GetProduct(ProductIDRequest) returns (ProductResponse)`
* `UpdateProduct(UpdateProductRequest) returns (Empty)`
* `DeleteProduct(ProductIDRequest) returns (DeleteProductResponse)`
* `ListProducts(Empty) returns (ProductListResponse)`
* `SearchProducts(SearchProductsRequest) returns (ProductListResponse)`
* `GetProductsByTag(TagRequest) returns (ProductListResponse)`
* `SetProductCache(Product) returns (CacheResponse)`
* `InvalidateProductCache(ProductIDRequest) returns (CacheResponse)`
* `SendProductEmail(SendProductEmailRequest) returns (EmailStatusResponse)`
* `GetAllFromRedis(Empty) returns (ProductListResponse)`
* `GetAllFromCache(Empty) returns (ProductListResponse)`
* `BulkCreateProducts(BulkCreateProductsRequest) returns (BulkCreateProductsResponse)`

---

### analytics-service

* `CreateView(CreateViewRequest) returns (CreateViewResponse)`
* `GetViewsByUser(UserRequest) returns (UserViewsResponse)`
* `GetViewsByProduct(ProductRequest) returns (ProductViewsResponse)`
* `GetViewsByUserAndProduct(UserProductRequest) returns (UserProductViewsResponse)`
* `GetRecentViews(RecentViewsRequest) returns (RecentViewsResponse)`
* `GetMostViewedProducts(Empty) returns (MostViewedProductsResponse)`
* `GetUserTopProducts(UserTopProductsRequest) returns (UserTopProductsResponse)`
* `GetProductViewCount(ProductRequest) returns (ProductViewCountResponse)`
* `GetUserViewCount(UserRequest) returns (UserViewCountResponse)`
* `GetDailyViews(Empty) returns (DailyViewsResponse)`
* `GenerateDailyViewReportEmail(ReportEmailRequest) returns (ReportEmailResponse)`
* `GetHourlyViews(ProductRequest) returns (HourlyViewsResponse)`
* `DeleteOldViews(DeleteOldViewsRequest) returns (DeleteOldViewsResponse)`
* `GetCachedView(ProductRequest) returns (ViewResponse)`
* `GetMemoryCachedView(ProductRequest) returns (ViewResponse)`

---

## Running Tests

Run all tests:

```bash
go test ./... -v
```

---

## Developer Tools

* Adminer UI (DB viewer): [http://localhost:8080](http://localhost:8080)
* grpcurl (optional gRPC testing):

  ```bash
  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
  grpcurl -plaintext localhost:50053 list
  ```

---

## About

CodeMart is a learning-oriented microservices app for exploring gRPC, Redis caching, event messaging (NATS), and clean service design using Go.

Services are designed to scale horizontally and plug easily into observability, tracing, and CI/CD setups.
