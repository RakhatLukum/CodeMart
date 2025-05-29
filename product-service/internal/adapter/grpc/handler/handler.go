package handler

import (
	"context"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/RakhatLukum/CodeMart/product-service/internal/usecase"
	pb "github.com/RakhatLukum/CodeMart/product-service/proto"
)

type ProductHandler struct {
	uc usecase.ProductUsecase
	pb.UnimplementedProductServiceServer
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (h *ProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := model.Product{
		Name:  req.GetName(),
		Price: req.GetPrice(),
		Tags:  req.GetTags(),
	}
	id, err := h.uc.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{
		Id:    int32(id),
		Name:  product.Name,
		Price: product.Price,
		Tags:  product.Tags,
	}, nil
}

func (h *ProductHandler) GetProduct(ctx context.Context, req *pb.ProductIDRequest) (*pb.ProductResponse, error) {
	product, err := h.uc.GetProduct(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return toProtoResponse(product), nil
}

func (h *ProductHandler) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.Empty, error) {
	product := model.Product{
		ID:    int(req.GetId()),
		Name:  req.GetName(),
		Price: req.GetPrice(),
		Tags:  req.GetTags(),
	}
	if err := h.uc.UpdateProduct(ctx, product); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (h *ProductHandler) DeleteProduct(ctx context.Context, req *pb.ProductIDRequest) (*pb.DeleteProductResponse, error) {
	if err := h.uc.DeleteProduct(ctx, int(req.GetId())); err != nil {
		return nil, err
	}
	return &pb.DeleteProductResponse{Deleted: true}, nil
}

func (h *ProductHandler) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductListResponse, error) {
	products, err := h.uc.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ProductListResponse{Products: toProtoResponseList(products)}, nil
}

func (h *ProductHandler) SearchProducts(ctx context.Context, req *pb.SearchProductsRequest) (*pb.ProductListResponse, error) {
	products, err := h.uc.SearchProducts(ctx, req.GetQuery(), req.GetTags())
	if err != nil {
		return nil, err
	}
	return &pb.ProductListResponse{Products: toProtoResponseList(products)}, nil
}

func (h *ProductHandler) GetProductsByTag(ctx context.Context, req *pb.TagRequest) (*pb.ProductListResponse, error) {
	products, err := h.uc.GetProductsByTag(ctx, req.GetTag())
	if err != nil {
		return nil, err
	}
	return &pb.ProductListResponse{Products: toProtoResponseList(products)}, nil
}

func (h *ProductHandler) SetProductCache(ctx context.Context, req *pb.Product) (*pb.CacheResponse, error) {
	product := model.Product{
		ID:    int(req.GetId()),
		Name:  req.GetName(),
		Price: req.GetPrice(),
		Tags:  req.GetTags(),
	}
	if err := h.uc.SetProductCache(ctx, product); err != nil {
		return &pb.CacheResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.CacheResponse{Success: true}, nil
}

func (h *ProductHandler) InvalidateProductCache(ctx context.Context, req *pb.ProductIDRequest) (*pb.CacheResponse, error) {
	if err := h.uc.InvalidateProductCache(ctx, int(req.GetId())); err != nil {
		return &pb.CacheResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.CacheResponse{Success: true}, nil
}

func (h *ProductHandler) SendProductEmail(ctx context.Context, req *pb.SendProductEmailRequest) (*pb.EmailStatusResponse, error) {
	product, err := h.uc.GetProduct(ctx, int(req.GetProductId()))
	if err != nil {
		return &pb.EmailStatusResponse{Sent: false, Message: err.Error()}, nil
	}
	err = h.uc.SendProductEmail(ctx, req.GetEmail(), "", product)
	if err != nil {
		return &pb.EmailStatusResponse{Sent: false, Message: err.Error()}, nil
	}
	return &pb.EmailStatusResponse{Sent: true}, nil
}

func (h *ProductHandler) GetAllFromRedis(ctx context.Context, _ *pb.Empty) (*pb.ProductListResponse, error) {
	products, err := h.uc.GetAllFromRedis(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.ProductListResponse{Products: toProtoResponseList(products)}, nil
}

func (h *ProductHandler) GetAllFromCache(ctx context.Context, _ *pb.Empty) (*pb.ProductListResponse, error) {
	products := h.uc.GetAllFromCache(ctx)
	return &pb.ProductListResponse{Products: toProtoResponseList(products)}, nil
}

func (h *ProductHandler) BulkCreateProducts(ctx context.Context, req *pb.BulkCreateProductsRequest) (*pb.BulkCreateProductsResponse, error) {
	var createdCount int
	for _, p := range req.GetProducts() {
		product := model.Product{
			Name:  p.GetName(),
			Price: p.GetPrice(),
			Tags:  p.GetTags(),
		}
		id, err := h.uc.CreateProduct(ctx, product)
		if err != nil {
			continue
		}
		createdCount++
		product.ID = id
		_ = h.uc.SetProductCache(ctx, product)
	}
	return &pb.BulkCreateProductsResponse{CreatedCount: int32(createdCount)}, nil
}

func toProtoResponse(p model.Product) *pb.ProductResponse {
	return &pb.ProductResponse{
		Id:    int32(p.ID),
		Name:  p.Name,
		Price: p.Price,
		Tags:  p.Tags,
	}
}

func toProtoResponseList(products []model.Product) []*pb.ProductResponse {
	var protoProducts []*pb.ProductResponse
	for _, p := range products {
		protoProducts = append(protoProducts, toProtoResponse(p))
	}
	return protoProducts
}
