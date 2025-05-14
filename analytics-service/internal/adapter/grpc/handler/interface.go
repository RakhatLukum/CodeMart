package handler

import (
	proto "CodeMart/analytics-service/proto"
	"context"
)

type ViewServiceHandler interface {
	CreateView(ctx context.Context, req *proto.CreateViewRequest) (*proto.CreateViewResponse, error)
	GetViewsByUser(ctx context.Context, req *proto.UserRequest) (*proto.UserViewsResponse, error)
	GetViewsByProduct(ctx context.Context, req *proto.ProductRequest) (*proto.ProductViewsResponse, error)
	GetViewsByUserAndProduct(ctx context.Context, req *proto.UserProductRequest) (*proto.UserProductViewsResponse, error)
	GetRecentViews(ctx context.Context, req *proto.RecentViewsRequest) (*proto.RecentViewsResponse, error)
	GetMostViewedProducts(ctx context.Context, req *proto.Empty) (*proto.MostViewedProductsResponse, error)
	GetUserTopProducts(ctx context.Context, req *proto.UserTopProductsRequest) (*proto.UserTopProductsResponse, error)
	GetProductViewCount(ctx context.Context, req *proto.ProductRequest) (*proto.ProductViewCountResponse, error)
	GetUserViewCount(ctx context.Context, req *proto.UserRequest) (*proto.UserViewCountResponse, error)
	GetDailyViews(ctx context.Context, req *proto.Empty) (*proto.DailyViewsResponse, error)
	GenerateDailyViewReportEmail(ctx context.Context, req *proto.ReportEmailRequest) (*proto.ReportEmailResponse, error)
	GetHourlyViews(ctx context.Context, req *proto.ProductRequest) (*proto.HourlyViewsResponse, error)
	DeleteOldViews(ctx context.Context, req *proto.DeleteOldViewsRequest) (*proto.DeleteOldViewsResponse, error)
	GetCachedView(ctx context.Context, req *proto.ProductRequest) (*proto.ViewResponse, error)
	GetMemoryCachedView(ctx context.Context, req *proto.ProductRequest) (*proto.ViewResponse, error)
}
