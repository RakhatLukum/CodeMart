package repository

import (
	"context"
	"database/sql"
	"strings"

	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(ctx context.Context, product model.Product) (int, error) {
	query := `INSERT INTO products (name, price, tags) VALUES (?, ?, ?)`
	res, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Tags)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (r *productRepo) UpdateProduct(ctx context.Context, product model.Product) error {
	query := `UPDATE products SET name = ?, price = ?, tags = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, product.Name, product.Price, product.Tags, product.ID)
	return err
}

func (r *productRepo) DeleteProduct(ctx context.Context, id int) error {
	query := `DELETE FROM products WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *productRepo) GetProductByID(ctx context.Context, id int) (model.Product, error) {
	var p model.Product
	query := `SELECT id, name, price, tags FROM products WHERE id = ?`
	err := r.db.QueryRowContext(ctx, query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Tags)
	if err != nil {
		return p, err
	}
	return p, nil
}

func (r *productRepo) ListProducts(ctx context.Context) ([]model.Product, error) {
	query := `SELECT id, name, price, tags FROM products`
	rows, err := r.db.QueryContext(ctx, query)
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

func (r *productRepo) SearchProducts(ctx context.Context, query, tags string) ([]model.Product, error) {
	q := `SELECT id, name, price, tags FROM products WHERE name LIKE ?`
	args := []interface{}{"%" + query + "%"}

	if tags != "" {
		tagList := strings.Split(tags, ",")
		for _, tag := range tagList {
			q += ` AND tags LIKE ?`
			args = append(args, "%"+strings.TrimSpace(tag)+"%")
		}
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
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

func (r *productRepo) GetProductsByTag(ctx context.Context, tag string) ([]model.Product, error) {
	query := `SELECT id, name, price, tags FROM products WHERE tags LIKE ?`
	rows, err := r.db.QueryContext(ctx, query, "%"+tag+"%")
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

func (r *productRepo) BulkInsertProducts(ctx context.Context, products []model.Product) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO products (name, price, tags) VALUES (?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()

	for _, p := range products {
		if _, err := stmt.ExecContext(ctx, p.Name, p.Price, p.Tags); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return len(products), nil
}
