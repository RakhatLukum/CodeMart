package repository

import (
	"context"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
)

type CartRepository interface {
	AddToCart(ctx context.Context, cart model.Cart) (int, error)
	RemoveFromCart(ctx context.Context, userID, productID int) error
	ClearCart(ctx context.Context, userID int) error
	GetCart(ctx context.Context, userID int) ([]model.Cart, error)
	GetCartItems(ctx context.Context, userID int) ([]model.Product, error)
	UpdateCartItem(ctx context.Context, cart model.Cart) error
	HasProductInCart(ctx context.Context, userID, productID int) (bool, error)
	GetCartItemCount(ctx context.Context, userID int) (int, error)
	GetCartTotalPrice(ctx context.Context, userID int) (float64, error)
}
