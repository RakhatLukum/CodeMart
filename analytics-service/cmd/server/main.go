package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	delivery "CodeMart/analytics-service/internal/analytics/delivery/grpc"
	repoPkg "CodeMart/analytics-service/internal/analytics/repository/mysqlrepo"
	usePkg "CodeMart/analytics-service/internal/analytics/usecase"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:MyStrongPassword123!@tcp(127.0.0.1:3306)/shop?parseTime=true"
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("db err: %v", err)
	}

	repo := repoPkg.New(db)
	uc := usePkg.New(repo)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	s := grpc.NewServer()
	delivery.Register(s, uc)

	log.Printf("analytics service on :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("serve: %v", err)
	}
}
