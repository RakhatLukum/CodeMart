package mysqlrepo

import (
	"CodeMart/product-service/internal/product/entity"
	"database/sql"
	"strings"
)

type repo struct{ db *sql.DB }

func New(db *sql.DB) *repo { return &repo{db: db} }

func (r *repo) GetAll() ([]*entity.Product, error) {
	rows, err := r.db.Query(`SELECT id,name,price,tags FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*entity.Product
	for rows.Next() {
		var p entity.Product
		var tags string
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &tags); err != nil {
			return nil, err
		}
		if tags != "" {
			p.Tags = strings.Split(tags, ",")
		}
		out = append(out, &p)
	}
	return out, nil
}
func (r *repo) GetByID(id int64) (*entity.Product, error) {
	var p entity.Product
	var tags string
	err := r.db.QueryRow(`SELECT id,name,price,tags FROM products WHERE id=?`, id).Scan(&p.ID, &p.Name, &p.Price, &tags)
	if err != nil {
		return nil, err
	}
	if tags != "" {
		p.Tags = strings.Split(tags, ",")
	}
	return &p, nil
}
func (r *repo) GetByTag(tag string) ([]*entity.Product, error) {
	rows, err := r.db.Query(`SELECT id,name,price,tags FROM products WHERE tags LIKE ?`, "%"+tag+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*entity.Product
	for rows.Next() {
		var p entity.Product
		var tags string
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &tags); err != nil {
			return nil, err
		}
		if tags != "" {
			p.Tags = strings.Split(tags, ",")
		}
		out = append(out, &p)
	}
	return out, nil
}
