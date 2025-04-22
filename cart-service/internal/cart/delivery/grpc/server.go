package grpc

import (
	"CodeMart/cart-service/internal/cart/usecase"
	"context"
	cartpb "proto/cart"

	"google.golang.org/grpc"
)

type server struct {
	cartpb.UnimplementedCartServiceServer
	uc *usecase.CartUsecase
}

func Register(s *grpc.Server, uc *usecase.CartUsecase) {
	cartpb.RegisterCartServiceServer(s, &server{uc: uc})
}

func (s *server) AddToCart(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartResponse, error) {
	if err := s.uc.Add(int64(req.UserId), int64(req.ProductId)); err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "added"}, nil
}
func (s *server) RemoveFromCart(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartResponse, error) {
	if err := s.uc.Remove(int64(req.UserId), int64(req.ProductId)); err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "removed"}, nil
}
func (s *server) GetCart(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartList, error) {
	ids, err := s.uc.List(int64(req.UserId))
	if err != nil {
		return nil, err
	}
	var res cartpb.CartList
	for _, id := range ids {
		res.Items = append(res.Items, &cartpb.CartItem{ProductId: int32(id)})
	}
	return &res, nil
}
