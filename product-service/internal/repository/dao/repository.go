package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type Product struct {
	db *sql.DB
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (d *Product) Insert(ctx context.Context, tx *sql.Tx, product model.Product) error {
	query := `INSERT INTO products (name, price, tags) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, product.Name, product.Price, product.Tags)
	return err
}

func (d *Product) InsertMany(ctx context.Context, tx *sql.Tx, products []model.Product) error {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO products (name, price, tags) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range products {
		if _, err := stmt.ExecContext(ctx, p.Name, p.Price, p.Tags); err != nil {
			return err
		}
	}
	return nil
}

func (d *Product) SelectLatest(ctx context.Context) (model.Product, error) {
	var p model.Product
	query := `SELECT id, name, price, tags FROM products ORDER BY id DESC LIMIT 1`
	err := d.db.QueryRowContext(ctx, query).Scan(&p.ID, &p.Name, &p.Price, &p.Tags)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		return model.Product{}, fmt.Errorf("select error: %w", err)
	}
	return p, nil
}

func (d *Product) DeleteByID(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := d.db.ExecContext(ctx, query, id)
	return err
}
