package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestView_Insert(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	tx, _ := db.Begin()
	viewRepo := NewView(db)

	view := model.View{UserID: 1, ProductID: 2, Timestamp: time.Now()}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO views`).
		WithArgs(view.UserID, view.ProductID, view.Timestamp).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := tx.Commit()
	assert.NoError(t, err)

	err = viewRepo.Insert(context.Background(), tx, view)
	assert.NoError(t, err)
}

func TestView_InsertMany(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	tx, _ := db.Begin()
	viewRepo := NewView(db)

	views := []model.View{
		{UserID: 1, ProductID: 2, Timestamp: time.Now()},
		{UserID: 3, ProductID: 4, Timestamp: time.Now()},
	}

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO views`).
		ExpectExec().WithArgs(views[0].UserID, views[0].ProductID, views[0].Timestamp).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`INSERT INTO views`).
		WithArgs(views[1].UserID, views[1].ProductID, views[1].Timestamp).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectClose()
	mock.ExpectCommit()

	err := tx.Commit()
	assert.NoError(t, err)

	err = viewRepo.InsertMany(context.Background(), tx, views)
	assert.NoError(t, err)
}

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
