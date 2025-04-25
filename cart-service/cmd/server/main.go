package main

import (
	cartgrpc "CodeMart/cart-service/internal/cart/delivery/grpc"
	"CodeMart/cart-service/internal/cart/repository/mysqlrepo"
	"CodeMart/cart-service/internal/cart/usecase"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	productpb "CodeMart/proto/product"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 50053, "The server port")
)

func main() {
	flag.Parse()

	// Get database connection string from environment or use default
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		dbDSN = "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true"
	}

	// Connect to MySQL
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// Initialize repository
	repo := mysqlrepo.New(db)

	// Initialize usecase
	uc := usecase.New(repo)

	// Connect to product service
	productConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	productClient := productpb.NewProductServiceClient(productConn)

	// Create gRPC server with middleware
	s := grpc.NewServer(
		grpc.UnaryInterceptor(cartgrpc.UserContextMiddleware()),
	)

	// Register cart service with product client
	cartgrpc.RegisterWithProductClient(s, uc, productClient)

	// Start server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
