package service

import (
	"fmt"
	"log"
	"net"

	"github.com/RakhatLukum/CodeMart/product-service/config"
	"github.com/RakhatLukum/CodeMart/product-service/internal/adapter/grpc/handler"
	"github.com/RakhatLukum/CodeMart/product-service/internal/usecase"
	pb "github.com/RakhatLukum/CodeMart/product-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	cfg      config.GRPCConfig
	server   *grpc.Server
	addr     string
	listener net.Listener
}

func NewGRPCServer(
	cfg config.Config,
	productUC usecase.ProductUsecase,
) (*GRPCServer, error) {
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.GRPC.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to listen on %s: %w", addr, err)
	}

	s := grpc.NewServer()
	productHandler := handler.NewProductHandler(productUC)

	pb.RegisterProductServiceServer(s, productHandler)

	reflection.Register(s)

	return &GRPCServer{
		cfg:      cfg.GRPC,
		server:   s,
		addr:     addr,
		listener: lis,
	}, nil
}

func (s *GRPCServer) Run() error {
	log.Printf("gRPC server running on %s", s.addr)
	return s.server.Serve(s.listener)
}

func (s *GRPCServer) Stop() {
	s.server.GracefulStop()
	log.Println("gRPC server stopped gracefully")
}
