package handler

import (
	"context"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/usecase"
	pb "github.com/RakhatLukum/CodeMart/cart-service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CartHandler struct {
	uc usecase.CartUsecase
	pb.UnimplementedCartServiceServer
}

func NewCartHandler(uc usecase.CartUsecase) *CartHandler {
	return &CartHandler{uc: uc}
}

func (h *CartHandler) AddToCart(ctx context.Context, req *pb.CreateCartRequest) (*pb.CreateCartResponse, error) {
	cart := model.Cart{
		UserID:    int(req.GetUserId()),
		ProductID: int(req.GetProductId()),
	}
	id, err := h.uc.AddToCart(ctx, cart)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCartResponse{
		Id:        int32(id),
		UserId:    int32(cart.UserID),
		ProductId: int32(cart.ProductID),
	}, nil
}

func (h *CartHandler) RemoveFromCart(ctx context.Context, req *pb.DeleteCartItemRequest) (*pb.DeleteCartItemResponse, error) {
	err := h.uc.RemoveFromCart(ctx, int(req.GetUserId()), int(req.GetProductId()))
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCartItemResponse{Deleted: true}, nil
}

func (h *CartHandler) ClearCart(ctx context.Context, req *pb.UserIDRequest) (*emptypb.Empty, error) {
	if err := h.uc.ClearCart(ctx, int(req.GetUserId())); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *CartHandler) GetCart(ctx context.Context, req *pb.UserIDRequest) (*pb.CartListResponse, error) {
	carts, err := h.uc.GetCart(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.CartListResponse{Carts: toProtoCartList(carts)}, nil
}

func (h *CartHandler) GetCartItems(ctx context.Context, req *pb.UserIDRequest) (*pb.CartItemsResponse, error) {
	products, err := h.uc.GetCartItems(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.CartItemsResponse{Items: toProtoProductList(products)}, nil
}

func (h *CartHandler) UpdateCartItem(ctx context.Context, req *pb.UpdateCartItemRequest) (*emptypb.Empty, error) {
	cart := model.Cart{
		ID:        int(req.GetId()),
		UserID:    int(req.GetUserId()),
		ProductID: int(req.GetProductId()),
	}
	if err := h.uc.UpdateCartItem(ctx, cart); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (h *CartHandler) HasProductInCart(ctx context.Context, req *pb.HasProductInCartRequest) (*pb.HasProductInCartResponse, error) {
	hasProduct, err := h.uc.HasProductInCart(ctx, int(req.GetUserId()), int(req.GetProductId()))
	if err != nil {
		return nil, err
	}
	return &pb.HasProductInCartResponse{HasProduct: hasProduct}, nil
}

func (h *CartHandler) GetCartItemCount(ctx context.Context, req *pb.UserIDRequest) (*pb.CartItemCountResponse, error) {
	count, err := h.uc.GetCartItemCount(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.CartItemCountResponse{Count: int32(count)}, nil
}

func (h *CartHandler) GetCartTotalPrice(ctx context.Context, req *pb.UserIDRequest) (*pb.CartTotalPriceResponse, error) {
	totalPrice, err := h.uc.GetCartTotalPrice(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}
	return &pb.CartTotalPriceResponse{TotalPrice: totalPrice}, nil
}

func (h *CartHandler) SendCartSummaryEmail(ctx context.Context, req *pb.SendCartSummaryEmailRequest) (*pb.EmailStatusResponse, error) {
	err := h.uc.SendCartSummaryEmail(ctx, req.GetToEmail(), req.GetToName(), int(req.GetUserId()))
	if err != nil {
		return &pb.EmailStatusResponse{Sent: false, Message: err.Error()}, nil
	}
	return &pb.EmailStatusResponse{Sent: true}, nil
}

func (h *CartHandler) InvalidateCartCache(ctx context.Context, req *pb.UserIDRequest) (*pb.CacheResponse, error) {
	if err := h.uc.InvalidateCartCache(ctx, int(req.GetUserId())); err != nil {
		return &pb.CacheResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.CacheResponse{Success: true}, nil
}

func (h *CartHandler) GetAllFromRedis(ctx context.Context, _ *emptypb.Empty) (*pb.CartListResponse, error) {
	carts, err := h.uc.GetAllFromRedis(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CartListResponse{Carts: toProtoCartList(carts)}, nil
}

func (h *CartHandler) GetAllFromCache(ctx context.Context, _ *emptypb.Empty) (*pb.CartListResponse, error) {
	carts := h.uc.GetAllFromCache(ctx)
	return &pb.CartListResponse{Carts: toProtoCartList(carts)}, nil
}

func toProtoCartList(carts []model.Cart) []*pb.Cart {
	var protoCarts []*pb.Cart
	for _, c := range carts {
		protoCarts = append(protoCarts, &pb.Cart{
			Id:        int32(c.ID),
			UserId:    int32(c.UserID),
			ProductId: int32(c.ProductID),
		})
	}
	return protoCarts
}

func toProtoProductList(products []model.Product) []*pb.Product {
	var protoProducts []*pb.Product
	for _, p := range products {
		protoProducts = append(protoProducts, &pb.Product{
			Id:    int32(p.ID),
			Name:  p.Name,
			Price: p.Price,
			Tags:  p.Tags,
		})
	}
	return protoProducts
}
