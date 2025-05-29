package redis

import (
	"context"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type ClientInterface interface {
	Set(ctx context.Context, product model.Product) error
	SetMany(ctx context.Context, products []model.Product) error
	Get(ctx context.Context, productID int) (model.Product, error)
	Delete(ctx context.Context, productID int) error
	GetAll(ctx context.Context) ([]model.Product, error)
}
