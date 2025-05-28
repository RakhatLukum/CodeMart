package usecase

import (
	"context"

	"github.com/RakhatLukum/CodeMart/product-service/internal/adapter/mailer"
	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/RakhatLukum/CodeMart/product-service/internal/repository"
)

type productUsecase struct {
	repo         repository.ProductRepository
	redisClient  ProductCacheUsecase
	memoryClient ProductMemoryUsecase
	mailer       mailer.Mailer
}

func NewProductUsecase(
	repo repository.ProductRepository,
	redisClient ProductCacheUsecase,
	memoryClient ProductMemoryUsecase,
	mailer mailer.Mailer,
) ProductUsecase {
	return &productUsecase{
		repo:         repo,
		redisClient:  redisClient,
		memoryClient: memoryClient,
		mailer:       mailer,
	}
}

func (uc *productUsecase) CreateProduct(ctx context.Context, product model.Product) (int, error) {
	id, err := uc.repo.CreateProduct(ctx, product)
	if err != nil {
		return 0, err
	}
	product.ID = id
	_ = uc.redisClient.Set(ctx, product)
	uc.memoryClient.Set(product)
	return id, nil
}

func (uc *productUsecase) GetProduct(ctx context.Context, id int) (model.Product, error) {
	return uc.repo.GetProductByID(ctx, id)
}

func (uc *productUsecase) UpdateProduct(ctx context.Context, product model.Product) error {
	err := uc.repo.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}
	_ = uc.redisClient.Set(ctx, product)
	uc.memoryClient.Set(product)
	return nil
}

func (uc *productUsecase) DeleteProduct(ctx context.Context, id int) error {
	err := uc.repo.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	_ = uc.redisClient.Delete(ctx, id)
	uc.memoryClient.Delete(id)
	return nil
}

func (uc *productUsecase) ListProducts(ctx context.Context) ([]model.Product, error) {
	return uc.repo.ListProducts(ctx)
}

func (uc *productUsecase) SearchProducts(ctx context.Context, query, tags string) ([]model.Product, error) {
	return uc.repo.SearchProducts(ctx, query, tags)
}

func (uc *productUsecase) GetProductsByTag(ctx context.Context, tag string) ([]model.Product, error) {
	return uc.repo.GetProductsByTag(ctx, tag)
}

func (uc *productUsecase) SetProductCache(ctx context.Context, product model.Product) error {
	return uc.redisClient.Set(ctx, product)
}

func (uc *productUsecase) InvalidateProductCache(ctx context.Context, id int) error {
	return uc.redisClient.Delete(ctx, id)
}

func (uc *productUsecase) SendProductEmail(ctx context.Context, toEmail, toName string, product model.Product) error {
	return uc.mailer.SendProductCreatedEmail(toEmail, toName, product)
}

func (uc *productUsecase) GetAllFromRedis(ctx context.Context) ([]model.Product, error) {
	return uc.redisClient.GetAll(ctx)
}

func (uc *productUsecase) GetAllFromCache(ctx context.Context) []model.Product {
	return uc.memoryClient.GetAll()
}
