package inmemory

import (
	"sync"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type Client struct {
	products map[int]model.Product
	mu       sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		products: make(map[int]model.Product),
	}
}

func (c *Client) Set(product model.Product) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.products[product.ID] = product
}

func (c *Client) SetMany(products []model.Product) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, product := range products {
		c.products[product.ID] = product
	}
}

func (c *Client) Get(productID int) (model.Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	product, ok := c.products[productID]
	return product, ok
}

func (c *Client) Delete(productID int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.products, productID)
}

func (c *Client) GetAll() []model.Product {
	c.mu.RLock()
	defer c.mu.RUnlock()

	products := make([]model.Product, 0, len(c.products))
	for _, product := range c.products {
		products = append(products, product)
	}
	return products
}
