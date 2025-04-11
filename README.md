# Online Grocery Store with Microservices

This project is an **Online Grocery Store** built using **Microservices Architecture**. Each service is responsible for handling a specific domain such as user management, product catalog, shopping cart, and analytics. These services communicate with each other using **gRPC** for efficient and high-performance communication.

## Architecture Overview

The application is divided into the following four main microservices:

- **User Service**: Handles user registration, login, and profile management.
- **Product Service**: Manages the product catalog, allowing retrieval of products and filtering by tags (e.g., halal, vegan).
- **Cart Service**: Manages the shopping cart for each user, including adding and removing products.
- **Analytics Service**: Tracks product views and provides insights like the top 5 most viewed products.

### Communication Between Services

All microservices communicate with each other using **gRPC**. The service interfaces are defined in **.proto** files, which are used to generate client and server code. Each service provides gRPC endpoints to be consumed by other services.

## Features

- **User Registration and Authentication**: Secure user management with basic registration and login functionality.
- **Product Catalog Management**: Browse and filter products, fetch product details by ID, and filter products by tags (e.g., halal, vegan).
- **Shopping Cart**: Add, remove, and view products in the shopping cart.
- **Analytics**: Track and view the most popular products based on user activity.

## Technology Stack

- **Programming Language**: Go (Golang)
- **Communication**: gRPC
- **Database**: SQLite or JSON (for storing data locally)
- **Containerization**: Docker (Optional)
- **Service Discovery**: Can be extended to support service discovery for dynamic scaling.

## Microservices Breakdown

### 1. **User Service**
- **Functions**: Registration, login, and fetching user profile by ID.
- **Endpoints**:
  - `Register (email + password)`
  - `Login (email + password)`
  - `GetUser (user_id)`
- **Proto File**: `user.proto`
  
### 2. **Product Service**
- **Functions**: Manage the product catalog, including retrieval by ID and filtering by tags (e.g., halal, vegan).
- **Endpoints**:
  - `GetAllProducts()`
  - `GetProductById(product_id)`
  - `GetProductsByTag(tag)`
- **Proto File**: `product.proto`

### 3. **Cart Service**
- **Functions**: Manage the shopping cart (add/remove products).
- **Endpoints**:
  - `AddToCart(user_id, product_id)`
  - `RemoveFromCart(user_id, product_id)`
  - `GetCart(user_id)`
- **Proto File**: `cart.proto`

### 4. **Analytics Service**
- **Functions**: Track product views and provide insights on the most popular products.
- **Endpoints**:
  - `LogProductView(user_id, product_id)`
  - `GetTopProducts()`
- **Proto File**: `analytics.proto`

## Getting Started

Follow the instructions below to get the project up and running on your local machine.

### Prerequisites

- **Go 1.18+**
- **gRPC**: Install the required gRPC tools for Go.
- **Protobuf**: Install Protobuf compiler (`protoc`) and the Go plugin for Protobuf.

To install `protoc` and the necessary plugin, run:

```bash
# Install Protobuf Compiler
sudo apt install protobuf-compiler

# Install the Go plugin for Protobuf
go get -u github.com/golang/protobuf/protoc-gen-go

