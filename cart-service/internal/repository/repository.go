package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
)

type cartRepository struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(ctx context.Context, cart model.Cart) (int, error) {
	res, err := r.db.ExecContext(ctx, `INSERT INTO carts (user_id, product_id) VALUES (?, ?)`, cart.UserID, cart.ProductID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert cart item: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}
	return int(id), nil
}

func (r *cartRepository) RemoveFromCart(ctx context.Context, userID, productID int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM carts WHERE user_id = ? AND product_id = ?`, userID, productID)
	return err
}

func (r *cartRepository) ClearCart(ctx context.Context, userID int) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM carts WHERE user_id = ?`, userID)
	return err
}

func (r *cartRepository) GetCart(ctx context.Context, userID int) ([]model.Cart, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, user_id, product_id FROM carts WHERE user_id = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carts []model.Cart
	for rows.Next() {
		var c model.Cart
		if err := rows.Scan(&c.ID, &c.UserID, &c.ProductID); err != nil {
			return nil, err
		}
		carts = append(carts, c)
	}
	return carts, nil
}

func (r *cartRepository) GetCartItems(ctx context.Context, userID int) ([]model.Product, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT p.id, p.name, p.price, p.tags
		FROM carts c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = ?
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Tags); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *cartRepository) UpdateCartItem(ctx context.Context, cart model.Cart) error {
	_, err := r.db.ExecContext(ctx, `UPDATE carts SET product_id = ? WHERE id = ? AND user_id = ?`, cart.ProductID, cart.ID, cart.UserID)
	return err
}

func (r *cartRepository) HasProductInCart(ctx context.Context, userID, productID int) (bool, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM carts WHERE user_id = ? AND product_id = ?`, userID, productID).Scan(&count)
	return count > 0, err
}

func (r *cartRepository) GetCartItemCount(ctx context.Context, userID int) (int, error) {
	var count int
	err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM carts WHERE user_id = ?`, userID).Scan(&count)
	return count, err
}

func (r *cartRepository) GetCartTotalPrice(ctx context.Context, userID int) (float64, error) {
	var total float64
	err := r.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(p.price), 0)
		FROM carts c
		JOIN products p ON c.product_id = p.id
		WHERE c.user_id = ?
	`, userID).Scan(&total)
	return total, err
}
