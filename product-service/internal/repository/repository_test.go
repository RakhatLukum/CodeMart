package repository_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/RakhatLukum/CodeMart/product-service/internal/model"
	"github.com/RakhatLukum/CodeMart/product-service/internal/repository"
	"github.com/stretchr/testify/require"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewProductRepository(db)

	ctx := context.Background()
	product := model.Product{
		Name:  "TestProduct",
		Price: 10.99,
		Tags:  "fruit",
	}

	mock.ExpectExec("INSERT INTO products").
		WithArgs(product.Name, product.Price, product.Tags).
		WillReturnResult(sqlmock.NewResult(1, 1))

	id, err := repo.CreateProduct(ctx, product)
	require.NoError(t, err)
	require.Equal(t, 1, id)
}

func TestGetProductByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewProductRepository(db)

	expected := model.Product{ID: 1, Name: "Apple", Price: 2.5, Tags: "fruit"}

	rows := sqlmock.NewRows([]string{"id", "name", "price", "tags"}).
		AddRow(expected.ID, expected.Name, expected.Price, expected.Tags)

	mock.ExpectQuery("SELECT id, name, price, tags FROM products WHERE id = ?").
		WithArgs(1).
		WillReturnRows(rows)

	result, err := repo.GetProductByID(context.Background(), 1)
	require.NoError(t, err)
	require.Equal(t, expected, result)
}

func TestDeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repository.NewProductRepository(db)

	mock.ExpectExec("DELETE FROM products WHERE id = ?").
		WithArgs(42).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteProduct(context.Background(), 42)
	require.NoError(t, err)
}
