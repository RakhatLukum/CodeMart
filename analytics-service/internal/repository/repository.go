package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"
)

type viewRepo struct {
	db *sql.DB
}

func NewViewRepository(db *sql.DB) ViewRepository {
	return &viewRepo{db: db}
}

func (r *viewRepo) CreateView(ctx context.Context, view model.View) error {
	query := `INSERT INTO views (user_id, product_id, timestamp) VALUES (?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, view.UserID, view.ProductID, view.Timestamp)
	return err
}

func (r *viewRepo) GetViewsByUser(ctx context.Context, userID int) ([]model.View, error) {
	query := `SELECT id, user_id, product_id, timestamp FROM views WHERE user_id = ?`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []model.View
	for rows.Next() {
		var v model.View
		if err := rows.Scan(&v.ID, &v.UserID, &v.ProductID, &v.Timestamp); err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	return views, nil
}

func (r *viewRepo) GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error) {
	query := `SELECT id, user_id, product_id, timestamp FROM views WHERE product_id = ?`
	rows, err := r.db.QueryContext(ctx, query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []model.View
	for rows.Next() {
		var v model.View
		if err := rows.Scan(&v.ID, &v.UserID, &v.ProductID, &v.Timestamp); err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	return views, nil
}

func (r *viewRepo) GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error) {
	query := `SELECT id, user_id, product_id, timestamp FROM views WHERE user_id = ? AND product_id = ?`
	rows, err := r.db.QueryContext(ctx, query, userID, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []model.View
	for rows.Next() {
		var v model.View
		if err := rows.Scan(&v.ID, &v.UserID, &v.ProductID, &v.Timestamp); err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	return views, nil
}

func (r *viewRepo) GetRecentViews(ctx context.Context, limit int) ([]model.View, error) {
	query := `SELECT id, user_id, product_id, timestamp FROM views ORDER BY timestamp DESC LIMIT ?`
	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []model.View
	for rows.Next() {
		var v model.View
		if err := rows.Scan(&v.ID, &v.UserID, &v.ProductID, &v.Timestamp); err != nil {
			return nil, err
		}
		views = append(views, v)
	}
	return views, nil
}

func (r *viewRepo) GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
	query := `SELECT product_id, COUNT(*) as views FROM views GROUP BY product_id ORDER BY views DESC LIMIT ?`
	rows, err := r.db.QueryContext(ctx, query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.ProductViewCount
	for rows.Next() {
		var row dto.ProductViewCount
		if err := rows.Scan(&row.ProductID, &row.ViewCount); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r *viewRepo) GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error) {
	query := `SELECT product_id, COUNT(*) as views FROM views WHERE user_id = ? GROUP BY product_id ORDER BY views DESC LIMIT ?`
	rows, err := r.db.QueryContext(ctx, query, userID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.ProductViewCount
	for rows.Next() {
		var row dto.ProductViewCount
		if err := rows.Scan(&row.ProductID, &row.ViewCount); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

func (r *viewRepo) GetProductViewCount(ctx context.Context, productID int) (int, error) {
	query := `SELECT COUNT(*) FROM views WHERE product_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, productID).Scan(&count)
	return count, err
}

func (r *viewRepo) GetUserViewCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM views WHERE user_id = ?`
	var count int
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&count)
	return count, err
}

func (r *viewRepo) GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error) {
	query := `SELECT DATE(timestamp) as day, COUNT(*) FROM views GROUP BY day`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.DailyViewStat
	for rows.Next() {
		var stat dto.DailyViewStat
		if err := rows.Scan(&stat.Date, &stat.ViewCount); err != nil {
			return nil, err
		}
		result = append(result, stat)
	}
	return result, nil
}

func (r *viewRepo) GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error) {
	query := `SELECT HOUR(timestamp) as hour, COUNT(*) FROM views GROUP BY hour`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.HourlyViewStat
	for rows.Next() {
		var stat dto.HourlyViewStat
		if err := rows.Scan(&stat.Hour, &stat.ViewCount); err != nil {
			return nil, err
		}
		result = append(result, stat)
	}
	return result, nil
}

func (r *viewRepo) DeleteOldViews(ctx context.Context, olderThan time.Time) error {
	query := `DELETE FROM views WHERE timestamp < ?`
	_, err := r.db.ExecContext(ctx, query, olderThan)
	return err
}
