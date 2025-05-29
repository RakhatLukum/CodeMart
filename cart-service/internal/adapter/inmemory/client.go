package inmemory

import (
	"sync"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
)

type Client struct {
	carts map[int]model.Cart
	mu    sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		carts: make(map[int]model.Cart),
	}
}

func (c *Client) Set(cart model.Cart) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.carts[cart.ID] = cart
}

func (c *Client) SetMany(carts []model.Cart) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, cart := range carts {
		c.carts[cart.ID] = cart
	}
}

func (c *Client) Get(cartID int) (model.Cart, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	cart, ok := c.carts[cartID]
	return cart, ok
}

func (c *Client) Delete(cartID int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.carts, cartID)
}

func (c *Client) GetAll() []model.Cart {
	c.mu.RLock()
	defer c.mu.RUnlock()

	carts := make([]model.Cart, 0, len(c.carts))
	for _, cart := range c.carts {
		carts = append(carts, cart)
	}
	return carts
}
