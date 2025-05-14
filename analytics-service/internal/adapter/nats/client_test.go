package nats

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"CodeMart/analytics-service/internal/model"
	"CodeMart/analytics-service/internal/model/dto"

	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockNATSConn struct {
	mock.Mock
	nats.Conn
}

func (m *mockNATSConn) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	args := m.Called(subject, cb)
	return args.Get(0).(*nats.Subscription), args.Error(1)
}

type mockSubscription struct {
	mock.Mock
	nats.Subscription
}

type mockViewUsecase struct {
	mock.Mock
}

func (m *mockViewUsecase) CreateView(ctx context.Context, view model.View) error {
	return m.Called(ctx, view).Error(0)
}

func (m *mockViewUsecase) GetViewsByUser(ctx context.Context, userID int) ([]model.View, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]model.View), args.Error(1)
}

func (m *mockViewUsecase) GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error) {
	args := m.Called(ctx, productID)
	return args.Get(0).([]model.View), args.Error(1)
}

func (m *mockViewUsecase) GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error) {
	args := m.Called(ctx, userID, productID)
	return args.Get(0).([]model.View), args.Error(1)
}

func (m *mockViewUsecase) GetRecentViews(ctx context.Context, limit int) ([]model.View, error) {
	args := m.Called(ctx, limit)
	return args.Get(0).([]model.View), args.Error(1)
}

func (m *mockViewUsecase) GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
	args := m.Called(ctx, limit)
	return args.Get(0).([]dto.ProductViewCount), args.Error(1)
}

func (m *mockViewUsecase) GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error) {
	args := m.Called(ctx, userID, limit)
	return args.Get(0).([]dto.ProductViewCount), args.Error(1)
}

func (m *mockViewUsecase) GetProductViewCount(ctx context.Context, productID int) (int, error) {
	args := m.Called(ctx, productID)
	return args.Int(0), args.Error(1)
}

func (m *mockViewUsecase) GetUserViewCount(ctx context.Context, userID int) (int, error) {
	args := m.Called(ctx, userID)
	return args.Int(0), args.Error(1)
}

func (m *mockViewUsecase) GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error) {
	args := m.Called(ctx)
	return args.Get(0).([]dto.DailyViewStat), args.Error(1)
}

func (m *mockViewUsecase) GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error) {
	args := m.Called(ctx)
	return args.Get(0).([]dto.HourlyViewStat), args.Error(1)
}

func (m *mockViewUsecase) DeleteOldViews(ctx context.Context, olderThan time.Time) error {
	return m.Called(ctx, olderThan).Error(0)
}

func (m *mockViewUsecase) GenerateDailyViewReportEmail(ctx context.Context, email, name string) error {
	return m.Called(ctx, email, name).Error(0)
}

func TestNewSubscriber(t *testing.T) {
	mockConn := &mockNATSConn{}
	mockUC := &mockViewUsecase{}

	sub := NewSubscriber(&nats.Conn{}, mockUC, "products", "users", "carts")
	assert.Equal(t, mockConn, sub.conn)
	assert.Equal(t, mockUC, sub.viewUsecase)
	assert.Equal(t, "products", sub.subjectProds)
	assert.Equal(t, "users", sub.subjectUsers)
	assert.Equal(t, "carts", sub.subjectCarts)
}

func TestSubscribe_Success(t *testing.T) {
	mockConn := &mockNATSConn{}
	mockConn.On("Subscribe", "products", mock.Anything).Return(&mockSubscription{}, nil)
	mockConn.On("Subscribe", "users", mock.Anything).Return(&mockSubscription{}, nil)
	mockConn.On("Subscribe", "carts", mock.Anything).Return(&mockSubscription{}, nil)

	mockUC := &mockViewUsecase{}

	sub := NewSubscriber(&nats.Conn{}, mockUC, "products", "users", "carts")
	err := sub.Subscribe()

	assert.NoError(t, err)
	mockConn.AssertExpectations(t)
}

func TestSubscribe_ProductError(t *testing.T) {
	mockConn := &mockNATSConn{}
	mockConn.On("Subscribe", "products", mock.Anything).Return(&mockSubscription{}, errors.New("subscribe error"))

	mockUC := &mockViewUsecase{}

	sub := NewSubscriber(&nats.Conn{}, mockUC, "products", "users", "carts")
	err := sub.Subscribe()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to subscribe to products")
}

func TestMessageHandlers(t *testing.T) {
	mockConn := &mockNATSConn{}
	mockConn.On("Subscribe", "products", mock.Anything).Run(func(args mock.Arguments) {
		cb := args.Get(1).(nats.MsgHandler)
		product := model.Product{ID: 1, Name: "Test", Price: 9.99}
		data, _ := json.Marshal(product)
		cb(&nats.Msg{Data: data})
	}).Return(&mockSubscription{}, nil)

	mockConn.On("Subscribe", "users", mock.Anything).Return(&mockSubscription{}, nil)
	mockConn.On("Subscribe", "carts", mock.Anything).Return(&mockSubscription{}, nil)

	mockUC := &mockViewUsecase{}

	sub := NewSubscriber(&nats.Conn{}, mockUC, "products", "users", "carts")
	err := sub.Subscribe()

	assert.NoError(t, err)
	mockConn.AssertExpectations(t)
}
