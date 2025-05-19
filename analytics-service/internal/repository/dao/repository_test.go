package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestView_SelectLatestByProductID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	viewRepo := NewView(db)

	expected := model.View{
		ID:        1,
		UserID:    2,
		ProductID: 3,
		Timestamp: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "product_id", "timestamp"}).
		AddRow(expected.ID, expected.UserID, expected.ProductID, expected.Timestamp)

	mock.ExpectQuery(`SELECT id, user_id, product_id, timestamp FROM views WHERE product_id = \? ORDER BY timestamp DESC LIMIT 1`).
		WithArgs(expected.ProductID).
		WillReturnRows(rows)

	view, err := viewRepo.SelectLatestByProductID(context.Background(), expected.ProductID)
	assert.NoError(t, err)
	assert.Equal(t, expected, view)
}

func TestView_DeleteByProductID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	viewRepo := NewView(db)
	productID := 3

	mock.ExpectExec(`DELETE FROM views WHERE product_id = \?`).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := viewRepo.DeleteByProductID(context.Background(), productID)
	assert.NoError(t, err)
}
