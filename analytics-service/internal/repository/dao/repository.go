package mysql

import (
	"CodeMart/analytics-service/internal/model"
	"context"
	"database/sql"
	"fmt"
)

type View struct {
	db *sql.DB
}

func NewView(db *sql.DB) *View {
	return &View{db: db}
}

func (d *View) Insert(ctx context.Context, tx *sql.Tx, view model.View) error {
	query := `INSERT INTO views (user_id, product_id, timestamp) VALUES (?, ?, ?)`
	_, err := tx.ExecContext(ctx, query, view.UserID, view.ProductID, view.Timestamp)
	return err
}

func (d *View) InsertMany(ctx context.Context, tx *sql.Tx, views []model.View) error {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO views (user_id, product_id, timestamp) VALUES (?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range views {
		if _, err := stmt.ExecContext(ctx, v.UserID, v.ProductID, v.Timestamp); err != nil {
			return err
		}
	}
	return nil
}

func (d *View) SelectLatestByProductID(ctx context.Context, productID int) (model.View, error) {
	var view model.View
	query := `SELECT id, user_id, product_id, timestamp FROM views WHERE product_id = ? ORDER BY timestamp DESC LIMIT 1`
	err := d.db.QueryRowContext(ctx, query, productID).Scan(
		&view.ID, &view.UserID, &view.ProductID, &view.Timestamp,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.View{}, nil
		}
		return model.View{}, fmt.Errorf("select error: %w", err)
	}
	return view, nil
}

func (d *View) DeleteByProductID(ctx context.Context, productID int) error {
	query := `DELETE FROM views WHERE product_id = ?`
	_, err := d.db.ExecContext(ctx, query, productID)
	return err
}
