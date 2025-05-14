package usecase

import (
	"CodeMart/analytics-service/internal/adapter/redis"
	"CodeMart/analytics-service/internal/model"
	"context"
)

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
