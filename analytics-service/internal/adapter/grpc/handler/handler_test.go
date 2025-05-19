package handler

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"
	proto "github.com/RakhatLukum/CodeMart/analytics-service/proto"

	"github.com/stretchr/testify/assert"
)

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

func TestViewHandler_CreateView_Success(t *testing.T) {
	mockViewUC := &mockViewUsecase{
		createViewFunc: func(ctx context.Context, view model.View) error {
			return nil
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	req := &proto.CreateViewRequest{
		UserId:    1,
		ProductId: 2,
	}
	resp, err := handler.CreateView(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, int32(2), resp.ProductId)
	assert.NotNil(t, resp.Timestamp)
}

func TestViewHandler_CreateView_Error(t *testing.T) {
	mockViewUC := &mockViewUsecase{
		createViewFunc: func(ctx context.Context, view model.View) error {
			return errors.New("create view error")
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	req := &proto.CreateViewRequest{
		UserId:    1,
		ProductId: 2,
	}
	resp, err := handler.CreateView(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestViewHandler_GetViewsByUser_Success(t *testing.T) {
	expectedViews := []model.View{
		{ID: 1, UserID: 1, ProductID: 10, Timestamp: time.Now()},
	}
	mockViewUC := &mockViewUsecase{
		getViewsByUserFunc: func(ctx context.Context, userID int) ([]model.View, error) {
			assert.Equal(t, 1, userID)
			return expectedViews, nil
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	req := &proto.UserRequest{UserId: 1}
	resp, err := handler.GetViewsByUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Len(t, resp.Views, 1)
	assert.Equal(t, int32(10), resp.Views[0].ProductId)
}

func TestViewHandler_GetViewsByUser_Error(t *testing.T) {
	mockViewUC := &mockViewUsecase{
		getViewsByUserFunc: func(ctx context.Context, userID int) ([]model.View, error) {
			return nil, errors.New("database error")
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	req := &proto.UserRequest{UserId: 1}
	resp, err := handler.GetViewsByUser(context.Background(), req)

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestViewHandler_GetMostViewedProducts_Success(t *testing.T) {
	expected := []dto.ProductViewCount{
		{ProductID: 1, ViewCount: 100},
	}
	mockViewUC := &mockViewUsecase{
		getMostViewedProductsFunc: func(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
			assert.Equal(t, 10, limit)
			return expected, nil
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	resp, err := handler.GetMostViewedProducts(context.Background(), &proto.Empty{})

	assert.NoError(t, err)
	assert.Len(t, resp.Data, 1)
	assert.Equal(t, int32(1), resp.Data[0].ProductId)
	assert.Equal(t, int32(100), resp.Data[0].ViewCount)
}

func TestViewHandler_GetMostViewedProducts_Error(t *testing.T) {
	mockViewUC := &mockViewUsecase{
		getMostViewedProductsFunc: func(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
			return nil, errors.New("database error")
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	resp, err := handler.GetMostViewedProducts(context.Background(), &proto.Empty{})

	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestViewHandler_GetDailyViews_Success(t *testing.T) {
	expected := []dto.DailyViewStat{
		{Date: "2024-01-01", ViewCount: 100},
	}
	mockViewUC := &mockViewUsecase{
		getDailyViewsFunc: func(ctx context.Context) ([]dto.DailyViewStat, error) {
			return expected, nil
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	resp, err := handler.GetDailyViews(context.Background(), &proto.Empty{})

	assert.NoError(t, err)
	assert.Len(t, resp.Data, 1)
	assert.Equal(t, "2024-01-01", resp.Data[0].Date)
	assert.Equal(t, int32(100), resp.Data[0].ViewCount)
}

func TestViewHandler_GenerateDailyViewReportEmail_Success(t *testing.T) {
	mockViewUC := &mockViewUsecase{
		generateDailyViewReportEmailFunc: func(ctx context.Context, email, name string) error {
			assert.Equal(t, "test@example.com", email)
			assert.Equal(t, "Test User", name)
			return nil
		},
	}
	handler := NewViewHandler(mockViewUC, nil, nil)

	req := &proto.ReportEmailRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	resp, err := handler.GenerateDailyViewReportEmail(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "Report sent successfully", resp.Message)
}

func TestViewHandler_GetCachedView_Success(t *testing.T) {
	expected := model.View{ID: 1, ProductID: 5}
	mockCacheUC := &mockViewCacheUsecase{
		getFunc: func(ctx context.Context, productID int) (model.View, error) {
			assert.Equal(t, 5, productID)
			return expected, nil
		},
	}
	handler := NewViewHandler(nil, mockCacheUC, nil)

	req := &proto.ProductRequest{ProductId: 5}
	resp, err := handler.GetCachedView(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, int32(5), resp.ProductId)
}

func TestViewHandler_GetMemoryCachedView_Success(t *testing.T) {
	expected := model.View{ID: 1, ProductID: 5}
	mockMemoryUC := &mockViewMemoryUsecase{
		getFunc: func(productID int) (model.View, bool) {
			assert.Equal(t, 5, productID)
			return expected, true
		},
	}
	handler := NewViewHandler(nil, nil, mockMemoryUC)

	req := &proto.ProductRequest{ProductId: 5}
	resp, err := handler.GetMemoryCachedView(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, int32(5), resp.ProductId)
}

func TestMapViewToProto(t *testing.T) {
	now := time.Now()
	view := model.View{
		ID:        1,
		UserID:    2,
		ProductID: 3,
		Timestamp: now,
	}
	protoView := mapViewToProto(view)

	assert.Equal(t, int32(1), protoView.Id)
	assert.Equal(t, int32(2), protoView.UserId)
	assert.Equal(t, int32(3), protoView.ProductId)
	assert.Equal(t, now.Unix(), protoView.Timestamp.AsTime().Unix())
}

func TestMapViewsToProto(t *testing.T) {
	views := []model.View{
		{ID: 1, UserID: 2, ProductID: 3},
		{ID: 4, UserID: 5, ProductID: 6},
	}
	protoViews := mapViewsToProto(views)

	assert.Len(t, protoViews, 2)
	assert.Equal(t, int32(1), protoViews[0].Id)
	assert.Equal(t, int32(5), protoViews[1].UserId)
}

func TestMapDailyViewStatsToProto(t *testing.T) {
	stats := []dto.DailyViewStat{
		{Date: "2024-01-01", ViewCount: 100},
		{Date: "2024-01-02", ViewCount: 200},
	}
	protoStats := mapDailyViewStatsToProto(stats)

	assert.Len(t, protoStats, 2)
	assert.Equal(t, "2024-01-02", protoStats[1].Date)
	assert.Equal(t, int32(200), protoStats[1].ViewCount)
}

func TestMapHourlyViewStatsToProto(t *testing.T) {
	stats := []dto.HourlyViewStat{
		{Hour: 10, ViewCount: 50},
		{Hour: 12, ViewCount: 60},
	}
	protoStats := mapHourlyViewStatsToProto(stats)

	assert.Len(t, protoStats, 2)
	assert.Equal(t, int32(12), protoStats[1].Hour)
	assert.Equal(t, int32(60), protoStats[1].ViewCount)
}
