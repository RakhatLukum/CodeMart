package service

import (
	"context"
	"errors"
	"net"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/config"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

type mockViewUsecase struct {
	createViewFunc                   func(context.Context, model.View) error
	getViewsByUserFunc               func(context.Context, int) ([]model.View, error)
	getViewsByProductFunc            func(context.Context, int) ([]model.View, error)
	getViewsByUserAndProductFunc     func(context.Context, int, int) ([]model.View, error)
	getRecentViewsFunc               func(context.Context, int) ([]model.View, error)
	getMostViewedProductsFunc        func(context.Context, int) ([]dto.ProductViewCount, error)
	getUserTopProductsFunc           func(context.Context, int, int) ([]dto.ProductViewCount, error)
	getProductViewCountFunc          func(context.Context, int) (int, error)
	getUserViewCountFunc             func(context.Context, int) (int, error)
	getDailyViewsFunc                func(context.Context) ([]dto.DailyViewStat, error)
	getHourlyViewsFunc               func(context.Context) ([]dto.HourlyViewStat, error)
	deleteOldViewsFunc               func(context.Context, time.Time) error
	generateDailyViewReportEmailFunc func(context.Context, string, string) error
}

func (m *mockViewUsecase) CreateView(ctx context.Context, view model.View) error {
	if m.createViewFunc != nil {
		return m.createViewFunc(ctx, view)
	}
	return errors.New("not implemented")
}

func (m *mockViewUsecase) GetViewsByUser(ctx context.Context, userID int) ([]model.View, error) {
	if m.getViewsByUserFunc != nil {
		return m.getViewsByUserFunc(ctx, userID)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error) {
	if m.getViewsByProductFunc != nil {
		return m.getViewsByProductFunc(ctx, productID)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error) {
	if m.getViewsByUserAndProductFunc != nil {
		return m.getViewsByUserAndProductFunc(ctx, userID, productID)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetRecentViews(ctx context.Context, limit int) ([]model.View, error) {
	if m.getRecentViewsFunc != nil {
		return m.getRecentViewsFunc(ctx, limit)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
	if m.getMostViewedProductsFunc != nil {
		return m.getMostViewedProductsFunc(ctx, limit)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error) {
	if m.getUserTopProductsFunc != nil {
		return m.getUserTopProductsFunc(ctx, userID, limit)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetProductViewCount(ctx context.Context, productID int) (int, error) {
	if m.getProductViewCountFunc != nil {
		return m.getProductViewCountFunc(ctx, productID)
	}
	return 0, errors.New("not implemented")
}

func (m *mockViewUsecase) GetUserViewCount(ctx context.Context, userID int) (int, error) {
	if m.getUserViewCountFunc != nil {
		return m.getUserViewCountFunc(ctx, userID)
	}
	return 0, errors.New("not implemented")
}

func (m *mockViewUsecase) GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error) {
	if m.getDailyViewsFunc != nil {
		return m.getDailyViewsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error) {
	if m.getHourlyViewsFunc != nil {
		return m.getHourlyViewsFunc(ctx)
	}
	return nil, errors.New("not implemented")
}

func (m *mockViewUsecase) DeleteOldViews(ctx context.Context, olderThan time.Time) error {
	if m.deleteOldViewsFunc != nil {
		return m.deleteOldViewsFunc(ctx, olderThan)
	}
	return errors.New("not implemented")
}

func (m *mockViewUsecase) GenerateDailyViewReportEmail(ctx context.Context, email, name string) error {
	if m.generateDailyViewReportEmailFunc != nil {
		return m.generateDailyViewReportEmailFunc(ctx, email, name)
	}
	return errors.New("not implemented")
}

type mockViewCacheUsecase struct {
	getFunc func(context.Context, int) (model.View, error)
}

func (m *mockViewCacheUsecase) Set(ctx context.Context, view model.View) error {
	return errors.New("not implemented")
}

func (m *mockViewCacheUsecase) SetMany(ctx context.Context, views []model.View) error {
	return errors.New("not implemented")
}

func (m *mockViewCacheUsecase) Get(ctx context.Context, productID int) (model.View, error) {
	if m.getFunc != nil {
		return m.getFunc(ctx, productID)
	}
	return model.View{}, errors.New("not implemented")
}

func (m *mockViewCacheUsecase) Delete(ctx context.Context, productID int) error {
	return errors.New("not implemented")
}

type mockViewMemoryUsecase struct {
	getFunc func(int) (model.View, bool)
}

func (m *mockViewMemoryUsecase) Set(view model.View) {
}

func (m *mockViewMemoryUsecase) SetMany(views []model.View) {
}

func (m *mockViewMemoryUsecase) Get(productID int) (model.View, bool) {
	if m.getFunc != nil {
		return m.getFunc(productID)
	}
	return model.View{}, false
}

func (m *mockViewMemoryUsecase) Delete(productID int) {
}

func TestNewGRPCServer_Success(t *testing.T) {
	cfg := config.Config{
		GRPC: config.GRPCConfig{
			Port: 50051,
		},
	}

	server, err := NewGRPCServer(
		cfg,
		&mockViewUsecase{},
		&mockViewCacheUsecase{},
		&mockViewMemoryUsecase{},
	)

	assert.NoError(t, err)
	assert.NotNil(t, server)
	assert.Equal(t, "0.0.0.0:50051", server.addr)
	assert.NotNil(t, server.server)
	assert.NotNil(t, server.listener)
	server.Stop()
}

func TestNewGRPCServer_ListenError(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	require.NoError(t, err)
	defer lis.Close()

	cfg := config.Config{
		GRPC: config.GRPCConfig{
			Port: 50051,
		},
	}

	server, err := NewGRPCServer(
		cfg,
		&mockViewUsecase{},
		&mockViewCacheUsecase{},
		&mockViewMemoryUsecase{},
	)

	assert.Error(t, err)
	assert.Nil(t, server)
	assert.Contains(t, err.Error(), "failed to listen on")
}

func TestGRPCServer_Run_Error(t *testing.T) {
	lis, err := net.Listen("tcp", "0.0.0.0:0")
	require.NoError(t, err)
	lis.Close()

	server := &GRPCServer{
		server:   grpc.NewServer(),
		listener: lis,
	}

	err = server.Run()
	assert.Error(t, err)
}

func TestGRPCServer_Stop(t *testing.T) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()

	server := &GRPCServer{
		server:   s,
		listener: lis,
	}

	go func() {
		_ = server.Run()
	}()

	time.Sleep(100 * time.Millisecond)

	server.Stop()

	err := server.Run()
	assert.Error(t, err)
}
