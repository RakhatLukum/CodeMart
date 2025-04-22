# Microshop – gRPC Clean Architecture Demo

This repository contains four independent microservices communicating over **gRPC**,
organised according to **Clean Architecture** principles.

## Services

| Service | Port | Responsibilities |
|---------|------|------------------|
| user     | 50051 | Register, login, profile |
| product  | 50052 | Product catalogue & filtering |
| cart     | 50053 | Shopping cart per user |
| analytics| 50054 | Product‑view statistics |

## Quick start

```bash
# 1. Generate protobuf stubs
make proto

# 2. Start a local MySQL (or use your own)
docker run -d --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=shop mysql:8
mysql -h127.0.0.1 -uroot -ppassword shop < database/schema.sql
mysql -h127.0.0.1 -uroot -ppassword shop < database/seed.sql

# 3. Run each service in its own terminal
go run user-service/cmd/server/main.go
go run product-service/cmd/server/main.go
go run cart-service/cmd/server/main.go
go run analytics-service/cmd/server/main.go
```

Environment variable `DB_DSN` overrides the default:

```
root:password@tcp(localhost:3306)/shop?parseTime=true
```

## Directory layout (per service)

```
<service>/
    cmd/server/main.go
    internal/
        entity/        # Enterprise objects
        repository/    # Interface + MySQL impl
        usecase/       # Application logic
        delivery/grpc/ # Transport layer
```

## Tooling requirements

* Go ≥ 1.22
* protoc + plugins (`protoc-gen-go`, `protoc-gen-go-grpc`)
* MySQL 8
