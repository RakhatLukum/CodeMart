package usecase

import (
	"context"
	"fmt"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/inmemory"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/mailer"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/nats"
	cache "github.com/RakhatLukum/CodeMart/cart-service/internal/adapter/redis"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/repository"
)

type cartUsecase struct {
	repo         repository.CartRepository
	redisClient  *cache.Client
	memoryClient *inmemory.Client
	mailer       mailer.Mailer
	publisher    *nats.Publisher
}

func NewCartUsecase(
	repo repository.CartRepository,
	redisClient *cache.Client,
	memoryClient *inmemory.Client,
	mailer mailer.Mailer,
	publisher *nats.Publisher,
) CartUsecase {
	return &cartUsecase{
		repo:         repo,
		redisClient:  redisClient,
		memoryClient: memoryClient,
		mailer:       mailer,
		publisher:    publisher,
	}
}

func (uc *cartUsecase) AddToCart(ctx context.Context, cart model.Cart) (int, error) {
	id, err := uc.repo.AddToCart(ctx, cart)
	if err != nil {
		return 0, err
	}
	uc.memoryClient.Set(cart)
	uc.redisClient.Set(ctx, cart)
	return id, nil
}

func (uc *cartUsecase) RemoveFromCart(ctx context.Context, userID, productID int) error {
	carts, err := uc.repo.GetCart(ctx, userID)
	if err != nil {
		return err
	}

	var cartIDToDelete int
	for _, c := range carts {
		if c.ProductID == productID {
			cartIDToDelete = c.ID
			break
		}
	}

	if cartIDToDelete == 0 {
		return fmt.Errorf("cart item not found for userID %d and productID %d", userID, productID)
	}

	err = uc.repo.RemoveFromCart(ctx, userID, productID)
	if err == nil {
		uc.memoryClient.Delete(cartIDToDelete)
		uc.redisClient.Delete(ctx, cartIDToDelete)
	}

	return err
}

func (uc *cartUsecase) ClearCart(ctx context.Context, userID int) error {
	return uc.repo.ClearCart(ctx, userID)
}

func (uc *cartUsecase) GetCart(ctx context.Context, userID int) ([]model.Cart, error) {
	return uc.repo.GetCart(ctx, userID)
}

func (uc *cartUsecase) GetCartItems(ctx context.Context, userID int) ([]model.Product, error) {
	items, err := uc.repo.GetCartItems(ctx, userID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *cartUsecase) UpdateCartItem(ctx context.Context, cart model.Cart) error {
	return uc.repo.UpdateCartItem(ctx, cart)
}

func (uc *cartUsecase) HasProductInCart(ctx context.Context, userID, productID int) (bool, error) {
	return uc.repo.HasProductInCart(ctx, userID, productID)
}

func (uc *cartUsecase) GetCartItemCount(ctx context.Context, userID int) (int, error) {
	return uc.repo.GetCartItemCount(ctx, userID)
}

func (uc *cartUsecase) GetCartTotalPrice(ctx context.Context, userID int) (float64, error) {
	return uc.repo.GetCartTotalPrice(ctx, userID)
}

func (uc *cartUsecase) SendCartSummaryEmail(ctx context.Context, toEmail, toName string, userID int) error {
	carts, err := uc.repo.GetCart(ctx, userID)
	if err != nil {
		return err
	}
	products, err := uc.repo.GetCartItems(ctx, userID)
	if err != nil {
		return err
	}
	return uc.mailer.SendCartSummaryEmail(toEmail, toName, carts, products)
}

func (uc *cartUsecase) InvalidateCartCache(ctx context.Context, userID int) error {
	return uc.redisClient.Delete(ctx, userID)
}

func (uc *cartUsecase) GetAllFromRedis(ctx context.Context) ([]model.Cart, error) {
	return uc.redisClient.GetAll(ctx)
}

func (uc *cartUsecase) GetAllFromCache(ctx context.Context) []model.Cart {
	return uc.memoryClient.GetAll()
}
