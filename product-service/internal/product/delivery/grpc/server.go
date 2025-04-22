package grpc

import (
	"CodeMart/product-service/internal/product/usecase"
	"context"
	productpb "proto/product"

	"google.golang.org/grpc"
)

type server struct {
	productpb.UnimplementedProductServiceServer
	uc *usecase.ProductUsecase
}

func Register(s *grpc.Server, uc *usecase.ProductUsecase) {
	productpb.RegisterProductServiceServer(s, &server{uc: uc})
}

func (s *server) GetAllProducts(ctx context.Context, _ *productpb.Empty) (*productpb.ProductList, error) {
	list, err := s.uc.All()
	if err != nil {
		return nil, err
	}
	var res productpb.ProductList
	for _, p := range list {
		res.Products = append(res.Products, &productpb.Product{
			Id: int32(p.ID), Name: p.Name, Price: float32(p.Price), Tags: p.Tags,
		})
	}
	return &res, nil
}
func (s *server) GetProductById(ctx context.Context, req *productpb.ProductIdRequest) (*productpb.Product, error) {
	p, err := s.uc.ByID(int64(req.ProductId))
	if err != nil {
		return nil, err
	}
	return &productpb.Product{Id: int32(p.ID), Name: p.Name, Price: float32(p.Price), Tags: p.Tags}, nil
}
func (s *server) GetProductsByTag(ctx context.Context, req *productpb.TagRequest) (*productpb.ProductList, error) {
	list, err := s.uc.ByTag(req.Tag)
	if err != nil {
		return nil, err
	}
	var res productpb.ProductList
	for _, p := range list {
		res.Products = append(res.Products, &productpb.Product{
			Id: int32(p.ID), Name: p.Name, Price: float32(p.Price), Tags: p.Tags,
		})
	}
	return &res, nil
}
