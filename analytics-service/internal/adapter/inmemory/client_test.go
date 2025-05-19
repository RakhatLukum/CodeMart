package inmemory

import (
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
	assert.NotNil(t, client.views)
	assert.Empty(t, client.views)
}

func TestSetAndGet(t *testing.T) {
	client := NewClient()
	now := time.Now()

	view := model.View{
		ID:        1,
		UserID:    100,
		ProductID: 200,
		Timestamp: now,
	}

	client.Set(view)

	retrievedView, found := client.Get(view.ProductID)
	assert.True(t, found)
	assert.Equal(t, view, retrievedView)

	_, found = client.Get(999)
	assert.False(t, found)
}

func TestSetMany(t *testing.T) {
	client := NewClient()
	now := time.Now()

	views := []model.View{
		{ID: 1, UserID: 100, ProductID: 200, Timestamp: now},
		{ID: 2, UserID: 101, ProductID: 201, Timestamp: now.Add(time.Hour)},
	}

	client.SetMany(views)

	for _, view := range views {
		retrievedView, found := client.Get(view.ProductID)
		assert.True(t, found)
		assert.Equal(t, view, retrievedView)
	}
}

func TestDelete(t *testing.T) {
	client := NewClient()
	view := model.View{ProductID: 200}

	client.Set(view)

	_, found := client.Get(view.ProductID)
	assert.True(t, found)

	client.Delete(view.ProductID)

	_, found = client.Get(view.ProductID)
	assert.False(t, found)
}

func TestConcurrentAccess(t *testing.T) {
	client := NewClient()
	numRoutines := 100
	done := make(chan bool)

	for i := 0; i < numRoutines; i++ {
		go func(id int) {
			view := model.View{
				ID:        id,
				UserID:    id * 10,
				ProductID: id * 100,
				Timestamp: time.Now(),
			}

			client.Set(view)
			_, _ = client.Get(view.ProductID)
			done <- true
		}(i)
	}

	for i := 0; i < numRoutines; i++ {
		<-done
	}

	for i := 0; i < numRoutines; i++ {
		_, found := client.Get(i * 100)
		assert.True(t, found)
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("ZeroValueView", func(t *testing.T) {
		client := NewClient()
		zeroView := model.View{}

		client.Set(zeroView)
		retrieved, found := client.Get(0)
		assert.True(t, found)
		assert.Equal(t, zeroView, retrieved)
	})

	t.Run("DuplicateProductID", func(t *testing.T) {
		client := NewClient()
		productID := 123
		view1 := model.View{ProductID: productID, UserID: 1}
		view2 := model.View{ProductID: productID, UserID: 2}

		client.Set(view1)
		client.Set(view2)

		retrieved, found := client.Get(productID)
		assert.True(t, found)
		assert.Equal(t, view2, retrieved)
	})

	t.Run("DeleteNonExistent", func(t *testing.T) {
		client := NewClient()
		client.Delete(999)
	})
}
