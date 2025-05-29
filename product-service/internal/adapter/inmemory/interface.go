package inmemory

import "github.com/RakhatLukum/CodeMart/product-service/internal/model"

type ClientInterface interface {
	Set(product model.Product)
	SetMany(products []model.Product)
	Get(productID int) (model.Product, bool)
	Delete(productID int)
	GetAll() []model.Product
}
