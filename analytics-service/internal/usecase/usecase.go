package usecase

import (
	"CodeMart/analytics-service/internal/model"
	"CodeMart/analytics-service/internal/model/dto"
	"CodeMart/analytics-service/internal/repository"
	"context"
	"time"
)

type viewUsecase struct {
	repo         repository.ViewRepository
	redisClient  ViewCacheUsecase
	memoryClient ViewMemoryUsecase
}

func NewViewUsecase(
	repo repository.ViewRepository,
	redisClient ViewCacheUsecase,
	memoryClient ViewMemoryUsecase,
) ViewUsecase {
	return &viewUsecase{
		repo:         repo,
		redisClient:  redisClient,
		memoryClient: memoryClient,
	}
}

func (uc *viewUsecase) CreateView(ctx context.Context, view model.View) error {
	if err := uc.repo.CreateView(ctx, view); err != nil {
		return err
	}

	if err := uc.redisClient.Set(ctx, view); err != nil {
		return err
	}

	uc.memoryClient.Set(view)

	return nil
}

func (uc *viewUsecase) GetViewsByUser(ctx context.Context, userID int) ([]model.View, error) {
	return uc.repo.GetViewsByUser(ctx, userID)
}

func (uc *viewUsecase) GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error) {
	return uc.repo.GetViewsByProduct(ctx, productID)
}

func (uc *viewUsecase) GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error) {
	return uc.repo.GetViewsByUserAndProduct(ctx, userID, productID)
}

func (uc *viewUsecase) GetRecentViews(ctx context.Context, limit int) ([]model.View, error) {
	return uc.repo.GetRecentViews(ctx, limit)
}

func (uc *viewUsecase) GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
	return uc.repo.GetMostViewedProducts(ctx, limit)
}

func (uc *viewUsecase) GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error) {
	return uc.repo.GetUserTopProducts(ctx, userID, limit)
}

func (uc *viewUsecase) GetProductViewCount(ctx context.Context, productID int) (int, error) {
	return uc.repo.GetProductViewCount(ctx, productID)
}

func (uc *viewUsecase) GetUserViewCount(ctx context.Context, userID int) (int, error) {
	return uc.repo.GetUserViewCount(ctx, userID)
}

func (uc *viewUsecase) GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error) {
	return uc.repo.GetDailyViews(ctx)
}

func (uc *viewUsecase) GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error) {
	return uc.repo.GetHourlyViews(ctx)
}

func (uc *viewUsecase) DeleteOldViews(ctx context.Context, olderThan time.Time) error {
	return uc.repo.DeleteOldViews(ctx, olderThan)
}
