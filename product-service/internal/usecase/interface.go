package usecase

import (
	"context"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, product model.Product) (int, error)
	GetProduct(ctx context.Context, id int) (model.Product, error)
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, id int) error
	ListProducts(ctx context.Context) ([]model.Product, error)
	SearchProducts(ctx context.Context, query, tags string) ([]model.Product, error)
	GetProductsByTag(ctx context.Context, tag string) ([]model.Product, error)
	SetProductCache(ctx context.Context, product model.Product) error
	InvalidateProductCache(ctx context.Context, id int) error
	SendProductEmail(ctx context.Context, toEmail, toName string, product model.Product) error
	GetAllFromRedis(ctx context.Context) ([]model.Product, error)
	GetAllFromCache(ctx context.Context) []model.Product
}

type ProductCacheUsecase interface {
	Set(ctx context.Context, product model.Product) error
	SetMany(ctx context.Context, products []model.Product) error
	Get(ctx context.Context, id int) (model.Product, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]model.Product, error)
}

type ProductMemoryUsecase interface {
	Set(product model.Product)
	SetMany(products []model.Product)
	Get(id int) (model.Product, bool)
	Delete(id int)
	GetAll() []model.Product
}
