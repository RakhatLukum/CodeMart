package repository

import (
	"context"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product model.Product) (int, error)
	UpdateProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetProductByID(ctx context.Context, id int) (model.Product, error)
	ListProducts(ctx context.Context) ([]model.Product, error)
	SearchProducts(ctx context.Context, query, tags string) ([]model.Product, error)
	GetProductsByTag(ctx context.Context, tag string) ([]model.Product, error)
	BulkInsertProducts(ctx context.Context, products []model.Product) (int, error)
}
