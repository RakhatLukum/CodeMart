package grpc

import (
	"CodeMart/cart-service/internal/cart/usecase"
	cartpb "CodeMart/proto/cart"
	productpb "CodeMart/proto/product"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	cartpb.UnimplementedCartServiceServer
	uc            *usecase.CartUsecase
	productClient productpb.ProductServiceClient
}

func Register(s *grpc.Server, uc *usecase.CartUsecase) {
	cartpb.RegisterCartServiceServer(s, &server{uc: uc})
}

func RegisterWithProductClient(s *grpc.Server, uc *usecase.CartUsecase, productClient productpb.ProductServiceClient) {
	cartpb.RegisterCartServiceServer(s, &server{uc: uc, productClient: productClient})
}

func (s *server) AddToCart(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if req.ProductId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product_id")
	}

	if err := s.uc.Add(int64(userID), int64(req.ProductId)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "added"}, nil
}

func (s *server) RemoveFromCart(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if req.ProductId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product_id")
	}

	if err := s.uc.Remove(int64(userID), int64(req.ProductId)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "removed"}, nil
}

func (s *server) GetCart(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartList, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	ids, err := s.uc.List(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Use a map to deduplicate product IDs
	seen := make(map[int64]bool)
	var uniqueIds []int64
	for _, id := range ids {
		if !seen[id] {
			seen[id] = true
			uniqueIds = append(uniqueIds, id)
		}
	}

	var res cartpb.CartList
	for _, id := range uniqueIds {
		res.Items = append(res.Items, &cartpb.CartItem{ProductId: int32(id)})
	}
	return &res, nil
}

func (s *server) ClearCart(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	err = s.uc.Clear(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "cart cleared"}, nil
}

func (s *server) CartItemCount(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartCountResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	count, err := s.uc.Count(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartCountResponse{Count: int32(count)}, nil
}

func (s *server) HasProduct(ctx context.Context, req *cartpb.CartItemRequest) (*cartpb.CartHasResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if req.ProductId <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid product_id")
	}

	has, err := s.uc.Has(int64(userID), int64(req.ProductId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartHasResponse{Has: has}, nil
}

func (s *server) ReplaceCart(ctx context.Context, req *cartpb.ReplaceCartRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if len(req.ProductIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "product_ids cannot be empty")
	}

	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		if id <= 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid product_id")
		}
		pids[i] = int64(id)
	}

	err = s.uc.Replace(int64(userID), pids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "cart replaced"}, nil
}

func (s *server) GetCartTotal(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartTotalResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if s.productClient == nil {
		return nil, status.Error(codes.Unimplemented, "product service client not configured")
	}

	ids, err := s.uc.List(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var total float32
	seen := make(map[int64]bool)
	for _, id := range ids {
		if seen[id] {
			continue // Skip duplicates
		}
		seen[id] = true

		prod, err := s.productClient.GetProductById(ctx, &productpb.ProductIdRequest{ProductId: int32(id)})
		if err != nil {
			continue // Skip if product not found
		}
		total += prod.Price
	}
	return &cartpb.CartTotalResponse{Total: total}, nil
}

func (s *server) GetCartProducts(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartProductList, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if s.productClient == nil {
		return nil, status.Error(codes.Unimplemented, "product service client not configured")
	}

	ids, err := s.uc.List(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var res cartpb.CartProductList
	seen := make(map[int64]bool)
	for _, id := range ids {
		if seen[id] {
			continue // Skip duplicates
		}
		seen[id] = true

		prod, err := s.productClient.GetProductById(ctx, &productpb.ProductIdRequest{ProductId: int32(id)})
		if err != nil {
			continue // Skip if product not found
		}
		res.Products = append(res.Products, &cartpb.Product{
			Id:    prod.Id,
			Name:  prod.Name,
			Price: prod.Price,
			Tags:  prod.Tags,
		})
	}
	return &res, nil
}

func (s *server) AddMultipleToCart(ctx context.Context, req *cartpb.AddMultipleRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if len(req.ProductIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "product_ids cannot be empty")
	}

	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		if id <= 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid product_id")
		}
		pids[i] = int64(id)
	}

	err = s.uc.AddMultiple(int64(userID), pids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "multiple products added"}, nil
}

func (s *server) RemoveMultipleFromCart(ctx context.Context, req *cartpb.RemoveMultipleRequest) (*cartpb.CartResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if len(req.ProductIds) == 0 {
		return nil, status.Error(codes.InvalidArgument, "product_ids cannot be empty")
	}

	pids := make([]int64, len(req.ProductIds))
	for i, id := range req.ProductIds {
		if id <= 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid product_id")
		}
		pids[i] = int64(id)
	}

	err = s.uc.RemoveMultiple(int64(userID), pids)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &cartpb.CartResponse{Message: "multiple products removed"}, nil
}

func (s *server) GetCartSummary(ctx context.Context, req *cartpb.UserIdRequest) (*cartpb.CartSummaryResponse, error) {
	userID, err := GetUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if s.productClient == nil {
		return nil, status.Error(codes.Unimplemented, "product service client not configured")
	}

	ids, err := s.uc.List(int64(userID))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var total float32
	seen := make(map[int64]bool)
	for _, id := range ids {
		if seen[id] {
			continue // Skip duplicates
		}
		seen[id] = true

		prod, err := s.productClient.GetProductById(ctx, &productpb.ProductIdRequest{ProductId: int32(id)})
		if err != nil {
			continue // Skip if product not found
		}
		total += prod.Price
	}

	return &cartpb.CartSummaryResponse{
		Count: int32(len(seen)),
		Total: total,
	}, nil
}
