package handler

import (
	"context"
	"fmt"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/usecase"
	proto "github.com/RakhatLukum/CodeMart/analytics-service/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ViewHandler struct {
	viewUC   usecase.ViewUsecase
	cacheUC  usecase.ViewCacheUsecase
	memoryUC usecase.ViewMemoryUsecase
	proto.UnimplementedViewServiceServer
}

func NewViewHandler(viewUC usecase.ViewUsecase, cacheUC usecase.ViewCacheUsecase, memoryUC usecase.ViewMemoryUsecase) *ViewHandler {
	return &ViewHandler{
		viewUC:   viewUC,
		cacheUC:  cacheUC,
		memoryUC: memoryUC,
	}
}

func (h *ViewHandler) CreateView(ctx context.Context, req *proto.CreateViewRequest) (*proto.CreateViewResponse, error) {
	view := model.View{
		UserID:    int(req.GetUserId()),
		ProductID: int(req.GetProductId()),
	}

	if err := h.viewUC.CreateView(ctx, view); err != nil {
		return nil, err
	}

	return &proto.CreateViewResponse{
		UserId:    req.GetUserId(),
		ProductId: req.GetProductId(),
		Timestamp: timestamppb.Now(),
	}, nil
}

func (h *ViewHandler) GetViewsByUser(ctx context.Context, req *proto.UserRequest) (*proto.UserViewsResponse, error) {
	views, err := h.viewUC.GetViewsByUser(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, fmt.Errorf("failed to get views by user: %w", err)
	}

	return &proto.UserViewsResponse{
		Views: mapViewsToProto(views),
	}, nil
}

func (h *ViewHandler) GetViewsByProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductViewsResponse, error) {
	views, err := h.viewUC.GetViewsByProduct(ctx, int(req.GetProductId()))
	if err != nil {
		return nil, err
	}

	return &proto.ProductViewsResponse{
		Views: mapViewsToProto(views),
	}, nil
}

func (h *ViewHandler) GetViewsByUserAndProduct(ctx context.Context, req *proto.UserProductRequest) (*proto.UserProductViewsResponse, error) {
	views, err := h.viewUC.GetViewsByUserAndProduct(ctx, int(req.GetUserId()), int(req.GetProductId()))
	if err != nil {
		return nil, err
	}

	return &proto.UserProductViewsResponse{
		Views: mapViewsToProto(views),
	}, nil
}

func (h *ViewHandler) GetRecentViews(ctx context.Context, req *proto.RecentViewsRequest) (*proto.RecentViewsResponse, error) {
	views, err := h.viewUC.GetRecentViews(ctx, int(req.GetLimit()))
	if err != nil {
		return nil, err
	}

	return &proto.RecentViewsResponse{
		Views: mapViewsToProto(views),
	}, nil
}

func (h *ViewHandler) GetMostViewedProducts(ctx context.Context, req *proto.Empty) (*proto.MostViewedProductsResponse, error) {
	products, err := h.viewUC.GetMostViewedProducts(ctx, 10) // Default limit
	if err != nil {
		return nil, fmt.Errorf("failed to get most viewed products: %w", err)
	}

	return &proto.MostViewedProductsResponse{
		Data: mapProductViewCountsToProto(products),
	}, nil
}

func (h *ViewHandler) GetUserTopProducts(ctx context.Context, req *proto.UserTopProductsRequest) (*proto.UserTopProductsResponse, error) {
	products, err := h.viewUC.GetUserTopProducts(ctx, int(req.GetUserId()), int(req.GetLimit()))
	if err != nil {
		return nil, err
	}

	return &proto.UserTopProductsResponse{
		UserId: req.GetUserId(),
		Data:   mapProductViewCountsToProto(products),
	}, nil
}

func (h *ViewHandler) GetProductViewCount(ctx context.Context, req *proto.ProductRequest) (*proto.ProductViewCountResponse, error) {
	count, err := h.viewUC.GetProductViewCount(ctx, int(req.GetProductId()))
	if err != nil {
		return nil, err
	}

	return &proto.ProductViewCountResponse{
		ProductId: req.GetProductId(),
		ViewCount: int32(count),
	}, nil
}

func (h *ViewHandler) GetUserViewCount(ctx context.Context, req *proto.UserRequest) (*proto.UserViewCountResponse, error) {
	count, err := h.viewUC.GetUserViewCount(ctx, int(req.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &proto.UserViewCountResponse{
		UserId:    req.GetUserId(),
		ViewCount: int32(count),
	}, nil
}

func (h *ViewHandler) GetDailyViews(ctx context.Context, req *proto.Empty) (*proto.DailyViewsResponse, error) {
	stats, err := h.viewUC.GetDailyViews(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily views: %w", err)
	}

	return &proto.DailyViewsResponse{
		Data: mapDailyViewStatsToProto(stats),
	}, nil
}

func (h *ViewHandler) GenerateDailyViewReportEmail(ctx context.Context, req *proto.ReportEmailRequest) (*proto.ReportEmailResponse, error) {
	err := h.viewUC.GenerateDailyViewReportEmail(ctx, req.GetEmail(), req.GetName())
	if err != nil {
		return nil, err
	}

	return &proto.ReportEmailResponse{
		Message: "Report sent successfully",
	}, nil
}

func (h *ViewHandler) GetHourlyViews(ctx context.Context, req *proto.ProductRequest) (*proto.HourlyViewsResponse, error) {
	stats, err := h.viewUC.GetHourlyViews(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.HourlyViewsResponse{
		ProductId: req.GetProductId(),
		Data:      mapHourlyViewStatsToProto(stats),
	}, nil
}

func (h *ViewHandler) DeleteOldViews(ctx context.Context, req *proto.DeleteOldViewsRequest) (*proto.DeleteOldViewsResponse, error) {
	err := h.viewUC.DeleteOldViews(ctx, req.GetBefore().AsTime())
	if err != nil {
		return nil, err
	}

	return &proto.DeleteOldViewsResponse{
		DeletedCount: 0,
	}, nil
}

func (h *ViewHandler) GetCachedView(ctx context.Context, req *proto.ProductRequest) (*proto.ViewResponse, error) {
	view, err := h.cacheUC.Get(ctx, int(req.GetProductId()))
	if err != nil {
		return nil, err
	}

	return mapViewToProto(view), nil
}

func (h *ViewHandler) GetMemoryCachedView(ctx context.Context, req *proto.ProductRequest) (*proto.ViewResponse, error) {
	view, exists := h.memoryUC.Get(int(req.GetProductId()))
	if !exists {
		return nil, fmt.Errorf("view not found in memory cache")
	}

	return mapViewToProto(view), nil
}

func mapViewToProto(v model.View) *proto.ViewResponse {
	return &proto.ViewResponse{
		Id:        int32(v.ID),
		UserId:    int32(v.UserID),
		ProductId: int32(v.ProductID),
		Timestamp: timestamppb.New(v.Timestamp),
	}
}

func mapViewsToProto(views []model.View) []*proto.ViewResponse {
	var protoViews []*proto.ViewResponse
	for _, v := range views {
		protoViews = append(protoViews, mapViewToProto(v))
	}
	return protoViews
}

func mapProductViewCountsToProto(counts []dto.ProductViewCount) []*proto.ProductViewCount {
	var protoCounts []*proto.ProductViewCount
	for _, c := range counts {
		protoCounts = append(protoCounts, &proto.ProductViewCount{
			ProductId: int32(c.ProductID),
			ViewCount: int32(c.ViewCount),
		})
	}
	return protoCounts
}

func mapDailyViewStatsToProto(stats []dto.DailyViewStat) []*proto.DailyViewStat {
	var protoStats []*proto.DailyViewStat
	for _, s := range stats {
		protoStats = append(protoStats, &proto.DailyViewStat{
			Date:      s.Date,
			ViewCount: int32(s.ViewCount),
		})
	}
	return protoStats
}

func mapHourlyViewStatsToProto(stats []dto.HourlyViewStat) []*proto.HourlyViewStat {
	var protoStats []*proto.HourlyViewStat
	for _, s := range stats {
		protoStats = append(protoStats, &proto.HourlyViewStat{
			Hour:      int32(s.Hour),
			ViewCount: int32(s.ViewCount),
		})
	}
	return protoStats
}
