package repository_test

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/model"
	"github.com/RakhatLukum/CodeMart/cart-service/internal/repository"
	"github.com/stretchr/testify/assert"
)

func setupRepo(t *testing.T) (*sql.DB, sqlmock.Sqlmock, repository.CartRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock db: %s", err)
	}
	repo := repository.NewCartRepository(db)
	return db, mock, repo
}

func TestAddToCart(t *testing.T) {
	db, mock, repo := setupRepo(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO carts (user_id, product_id) VALUES (?, ?)")).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(10, 1))

	id, err := repo.AddToCart(context.Background(), model.Cart{UserID: 1, ProductID: 2})

	assert.NoError(t, err)
	assert.Equal(t, 10, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveFromCart(t *testing.T) {
	db, mock, repo := setupRepo(t)
	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM carts WHERE user_id = ? AND product_id = ?")).WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.RemoveFromCart(context.Background(), 1, 2)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestHasProductInCart(t *testing.T) {
	db, mock, repo := setupRepo(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM carts WHERE user_id = ? AND product_id = ?")).WithArgs(1, 2).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	has, err := repo.HasProductInCart(context.Background(), 1, 2)
	assert.NoError(t, err)
	assert.True(t, has)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetCartItemCount(t *testing.T) {
	db, mock, repo := setupRepo(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT COUNT(*) FROM carts WHERE user_id = ?")).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))

	count, err := repo.GetCartItemCount(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
	assert.NoError(t, mock.ExpectationsWereMet())
}
