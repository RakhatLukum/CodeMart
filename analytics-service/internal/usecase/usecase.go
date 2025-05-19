package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/inmemory"
	mailjet "github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/mailer"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/adapter/redis"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/repository"
)

type viewUsecase struct {
	repo         repository.ViewRepository
	redisClient  ViewCacheUsecase
	memoryClient ViewMemoryUsecase
	mailer       mailjet.Mailer
}

func NewViewUsecase(
	repo repository.ViewRepository,
	redisClient ViewCacheUsecase,
	memoryClient ViewMemoryUsecase,
	mailer mailjet.Mailer,
) ViewUsecase {
	return &viewUsecase{
		repo:         repo,
		redisClient:  redisClient,
		memoryClient: memoryClient,
		mailer:       mailer,
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

type viewCacheUsecase struct {
	redisClient *redis.Client
}

func NewViewCacheUsecase(redisClient *redis.Client) ViewCacheUsecase {
	return &viewCacheUsecase{redisClient: redisClient}
}

func (uc *viewCacheUsecase) Set(ctx context.Context, view model.View) error {
	return uc.redisClient.Set(ctx, view)
}

func (uc *viewCacheUsecase) SetMany(ctx context.Context, views []model.View) error {
	return uc.redisClient.SetMany(ctx, views)
}

func (uc *viewCacheUsecase) Get(ctx context.Context, productID int) (model.View, error) {
	return uc.redisClient.Get(ctx, productID)
}

func (uc *viewCacheUsecase) Delete(ctx context.Context, productID int) error {
	return uc.redisClient.Delete(ctx, productID)
}

type viewMemoryUsecase struct {
	memoryClient *inmemory.Client
}

func NewViewMemoryUsecase(memoryClient *inmemory.Client) ViewMemoryUsecase {
	return &viewMemoryUsecase{memoryClient: memoryClient}
}

func (uc *viewMemoryUsecase) Set(view model.View) {
	uc.memoryClient.Set(view)
}

func (uc *viewMemoryUsecase) SetMany(views []model.View) {
	uc.memoryClient.SetMany(views)
}

func (uc *viewMemoryUsecase) Get(productID int) (model.View, bool) {
	return uc.memoryClient.Get(productID)
}

func (uc *viewMemoryUsecase) Delete(productID int) {
	uc.memoryClient.Delete(productID)
}

func (uc *viewUsecase) GenerateDailyViewReportEmail(ctx context.Context, email string, name string) error {
	stats, err := uc.GetDailyViews(ctx)
	if err != nil {
		return err
	}

	if len(stats) == 0 {
		return fmt.Errorf("no views recorded yet")
	}

	report := "Daily View Stats:\n\n"
	for _, s := range stats {
		report += fmt.Sprintf("%s — %d views\n", s.Date, s.ViewCount)
	}

	err = uc.mailer.SendDailyReportEmail(email, name, report)
	if err != nil {
		return err
	}

	return nil
}
