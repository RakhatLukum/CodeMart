package repository

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockRepo(t *testing.T) (*sql.DB, sqlmock.Sqlmock, ViewRepository) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	repo := NewViewRepository(db)
	return db, mock, repo
}

func TestCreateView(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	view := model.View{UserID: 1, ProductID: 2, Timestamp: time.Now()}
	mock.ExpectExec(`INSERT INTO views`).WithArgs(view.UserID, view.ProductID, view.Timestamp).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.CreateView(context.Background(), view)
	assert.NoError(t, err)
}

func TestGetViewsByUser(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "product_id", "timestamp"}).
		AddRow(1, 1, 2, time.Now())

	mock.ExpectQuery(`SELECT id, user_id, product_id, timestamp FROM views WHERE user_id = \?`).
		WithArgs(1).WillReturnRows(rows)

	views, err := repo.GetViewsByUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, views, 1)
}

func TestGetViewsByProduct(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "product_id", "timestamp"}).
		AddRow(1, 2, 3, time.Now())

	mock.ExpectQuery(`SELECT id, user_id, product_id, timestamp FROM views WHERE product_id = \?`).
		WithArgs(3).WillReturnRows(rows)

	views, err := repo.GetViewsByProduct(context.Background(), 3)
	assert.NoError(t, err)
	assert.Len(t, views, 1)
}

func TestGetViewsByUserAndProduct(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "product_id", "timestamp"}).
		AddRow(1, 2, 3, time.Now())

	mock.ExpectQuery(`SELECT id, user_id, product_id, timestamp FROM views WHERE user_id = \? AND product_id = \?`).
		WithArgs(2, 3).WillReturnRows(rows)

	views, err := repo.GetViewsByUserAndProduct(context.Background(), 2, 3)
	assert.NoError(t, err)
	assert.Len(t, views, 1)
}

func TestGetRecentViews(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "user_id", "product_id", "timestamp"}).
		AddRow(1, 2, 3, time.Now())

	mock.ExpectQuery(`SELECT id, user_id, product_id, timestamp FROM views ORDER BY timestamp DESC LIMIT \?`).
		WithArgs(1).WillReturnRows(rows)

	views, err := repo.GetRecentViews(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, views, 1)
}

func TestGetMostViewedProducts(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"product_id", "views"}).
		AddRow(3, 10)

	mock.ExpectQuery(`SELECT product_id, COUNT\(\*\) as views FROM views GROUP BY product_id ORDER BY views DESC LIMIT \?`).
		WithArgs(1).WillReturnRows(rows)

	result, err := repo.GetMostViewedProducts(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, []dto.ProductViewCount{{ProductID: 3, ViewCount: 10}}, result)
}

func TestGetUserTopProducts(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"product_id", "views"}).
		AddRow(3, 5)

	mock.ExpectQuery(`SELECT product_id, COUNT\(\*\) as views FROM views WHERE user_id = \? GROUP BY product_id ORDER BY views DESC LIMIT \?`).
		WithArgs(2, 1).WillReturnRows(rows)

	result, err := repo.GetUserTopProducts(context.Background(), 2, 1)
	assert.NoError(t, err)
	assert.Equal(t, []dto.ProductViewCount{{ProductID: 3, ViewCount: 5}}, result)
}

func TestGetProductViewCount(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM views WHERE product_id = \?`).
		WithArgs(3).WillReturnRows(sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(12))

	count, err := repo.GetProductViewCount(context.Background(), 3)
	assert.NoError(t, err)
	assert.Equal(t, 12, count)
}

func TestGetUserViewCount(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM views WHERE user_id = \?`).
		WithArgs(2).WillReturnRows(sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(8))

	count, err := repo.GetUserViewCount(context.Background(), 2)
	assert.NoError(t, err)
	assert.Equal(t, 8, count)
}

func TestGetDailyViews(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"day", "COUNT(*)"}).
		AddRow("2024-05-01", 100)

	mock.ExpectQuery(`SELECT DATE\(timestamp\) as day, COUNT\(\*\) FROM views GROUP BY day`).
		WillReturnRows(rows)

	stats, err := repo.GetDailyViews(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "2024-05-01", stats[0].Date)
	assert.Equal(t, 100, stats[0].ViewCount)
}

func TestGetHourlyViews(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"hour", "COUNT(*)"}).
		AddRow(14, 50)

	mock.ExpectQuery(`SELECT HOUR\(timestamp\) as hour, COUNT\(\*\) FROM views GROUP BY hour`).
		WillReturnRows(rows)

	stats, err := repo.GetHourlyViews(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 14, stats[0].Hour)
	assert.Equal(t, 50, stats[0].ViewCount)
}

func TestDeleteOldViews(t *testing.T) {
	db, mock, repo := newMockRepo(t)
	defer db.Close()

	threshold := time.Now().Add(-24 * time.Hour)

	mock.ExpectExec(`DELETE FROM views WHERE timestamp < \?`).
		WithArgs(threshold).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteOldViews(context.Background(), threshold)
	assert.NoError(t, err)
}
