package main

import (
	"context"
	"fmt"

	pb "github.com/RakhatLukum/CodeMart/proto"
)

// Cardserver implements the Cart Service Server interface
type CartServer struct {
	pb.UnimplementedCartServiceServer
	carts map[int32][]int32 // user_id: list of product_ids
}

func (s *CartServer) AddToCart(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {
	s.carts[req.UserId] = append(s.carts[req.UserId], req.ProductId)
	return &pb.CartResponse{Message: fmt.Sprintf("Added product %d to user %d's cart", req.ProductId, req.UserId)}, nil
}

func (s *CartServer) RemoveFromCart(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {
	products := s.carts[req.UserId]
	newProducts := []int32{}
	for _, pid := range products {
		if pid != req.ProductId {
			newProducts = append(newProducts, pid)
		}
	}
	s.carts[req.UserId] = newProducts
	return &pb.CartResponse{Message: fmt.Sprintf("Removed product %d from user %d's cart", req.ProductId, req.UserId)}, nil
}

func (s *CartServer) GetCart(ctx context.Context, req *pb.UserIdRequest) (*pb.CartItems, error) {
	products := s.carts[req.UserId]
	var items []*pb.CartItem
	for _, pid := range products {
		items = append(items, &pb.CartItem{ProductId: pid})
	}
	return &pb.CartItems{Items: items}, nil
}
