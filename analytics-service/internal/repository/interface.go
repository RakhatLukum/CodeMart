package repository

import (
	"context"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"
)

type ViewRepository interface {
	CreateView(ctx context.Context, view model.View) error
	GetViewsByUser(ctx context.Context, userID int) ([]model.View, error)
	GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error)
	GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error)
	GetRecentViews(ctx context.Context, limit int) ([]model.View, error)
	GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error)
	GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error)
	GetProductViewCount(ctx context.Context, productID int) (int, error)
	GetUserViewCount(ctx context.Context, userID int) (int, error)
	GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error)
	GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error)
	DeleteOldViews(ctx context.Context, olderThan time.Time) error
}
