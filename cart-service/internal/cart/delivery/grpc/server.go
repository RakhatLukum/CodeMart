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
func (s *server) ClearCart(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartResponse, error) {
	err := s.uc.Clear(int64(req.UserId))
	if err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "cart cleared"}, nil
}
func (s *server) CartItemCount(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartCountResponse, error) {
	count, err := s.uc.Count(int64(req.UserId))
	if err != nil {
		return nil, err
	}
	return &cartpb.CartCountResponse{Count: int32(count)}, nil
}
func (s *server) HasProduct(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartHasResponse, error) {
	has, err := s.uc.Has(int64(req.UserId), int64(req.ProductId))
	if err != nil {
		return nil, err
	}
	return &cartpb.CartHasResponse{Has: has}, nil
}
func (s *server) ReplaceCart(ctx context.Context, req *cartpb.ReplaceCartRequest) (*cartpb.CartResponse, error) {
	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		pids[i] = int64(id)
	}
	err := s.uc.Replace(int64(req.UserId), pids)
	if err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "cart replaced"}, nil
}
func (s *server) GetCartTotal(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartTotalResponse, error) {
	// TODO: Call product-service for real prices
	return &cartpb.CartTotalResponse{Total: 0}, nil
}
func (s *server) GetCartProducts(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartProductList, error) {
	// TODO: Call product-service for real product details
	return &cartpb.CartProductList{}, nil
}
func (s *server) AddMultipleToCart(ctx context.Context, req *cartpb.AddMultipleRequest) (*cartpb.CartResponse, error) {
	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		pids[i] = int64(id)
	}
	err := s.uc.AddMultiple(int64(req.UserId), pids)
	if err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "multiple products added"}, nil
}
func (s *server) RemoveMultipleFromCart(ctx context.Context, req *cartpb.RemoveMultipleRequest) (*cartpb.CartResponse, error) {
	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		pids[i] = int64(id)
	}
	err := s.uc.RemoveMultiple(int64(req.UserId), pids)
	if err != nil {
		return nil, err
	}
	return &cartpb.CartResponse{Message: "multiple products removed"}, nil
}
func (s *server) GetCartSummary(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartSummaryResponse, error) {
	// TODO: Call product-service for real count/total
	return &cartpb.CartSummaryResponse{Count: 0, Total: 0}, nil
}
