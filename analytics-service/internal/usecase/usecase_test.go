package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model"
	"github.com/RakhatLukum/CodeMart/analytics-service/internal/model/dto"

	"github.com/stretchr/testify/assert"
)

type mockRepo struct {
	createViewCalled    bool
	getDailyViewsCalled bool
	views               []model.View
	dailyStats          []dto.DailyViewStat
	createViewErr       error
	getDailyViewsErr    error
}

func (m *mockRepo) CreateView(ctx context.Context, view model.View) error {
	m.createViewCalled = true
	return m.createViewErr
}
func (m *mockRepo) GetViewsByUser(ctx context.Context, userID int) ([]model.View, error) {
	return m.views, nil
}
func (m *mockRepo) GetViewsByProduct(ctx context.Context, productID int) ([]model.View, error) {
	return m.views, nil
}
func (m *mockRepo) GetViewsByUserAndProduct(ctx context.Context, userID, productID int) ([]model.View, error) {
	return m.views, nil
}
func (m *mockRepo) GetRecentViews(ctx context.Context, limit int) ([]model.View, error) {
	return m.views, nil
}
func (m *mockRepo) GetMostViewedProducts(ctx context.Context, limit int) ([]dto.ProductViewCount, error) {
	return nil, nil
}
func (m *mockRepo) GetUserTopProducts(ctx context.Context, userID, limit int) ([]dto.ProductViewCount, error) {
	return nil, nil
}
func (m *mockRepo) GetProductViewCount(ctx context.Context, productID int) (int, error) {
	return 0, nil
}
func (m *mockRepo) GetUserViewCount(ctx context.Context, userID int) (int, error) {
	return 0, nil
}
func (m *mockRepo) GetDailyViews(ctx context.Context) ([]dto.DailyViewStat, error) {
	m.getDailyViewsCalled = true
	return m.dailyStats, m.getDailyViewsErr
}
func (m *mockRepo) GetHourlyViews(ctx context.Context) ([]dto.HourlyViewStat, error) {
	return nil, nil
}
func (m *mockRepo) DeleteOldViews(ctx context.Context, olderThan time.Time) error {
	return nil
}

type mockCache struct {
	setCalled bool
	setErr    error
}

func (m *mockCache) Set(ctx context.Context, view model.View) error {
	m.setCalled = true
	return m.setErr
}
func (m *mockCache) SetMany(ctx context.Context, views []model.View) error {
	return nil
}
func (m *mockCache) Get(ctx context.Context, productID int) (model.View, error) {
	return model.View{}, nil
}
func (m *mockCache) Delete(ctx context.Context, productID int) error {
	return nil
}

type mockMemory struct {
	setCalled bool
}

func (m *mockMemory) Set(view model.View) {
	m.setCalled = true
}
func (m *mockMemory) SetMany(views []model.View) {}
func (m *mockMemory) Get(productID int) (model.View, bool) {
	return model.View{}, false
}
func (m *mockMemory) Delete(productID int) {}

type mockMailer struct {
	sendCalled bool
	sendErr    error
}

func (m *mockMailer) SendDailyReportEmail(email, name, report string) error {
	m.sendCalled = true
	return m.sendErr
}

func TestCreateView_Success(t *testing.T) {
	repo := &mockRepo{}
	cache := &mockCache{}
	memory := &mockMemory{}
	mailer := &mockMailer{}

	uc := NewViewUsecase(repo, cache, memory, mailer)

	view := model.View{UserID: 1, ProductID: 2}
	err := uc.CreateView(context.Background(), view)

	assert.NoError(t, err)
	assert.True(t, repo.createViewCalled)
	assert.True(t, cache.setCalled)
	assert.True(t, memory.setCalled)
}

func TestCreateView_FailRepo(t *testing.T) {
	repo := &mockRepo{createViewErr: errors.New("repo error")}
	cache := &mockCache{}
	memory := &mockMemory{}
	mailer := &mockMailer{}

	uc := NewViewUsecase(repo, cache, memory, mailer)

	err := uc.CreateView(context.Background(), model.View{UserID: 1})
	assert.Error(t, err)
	assert.Equal(t, "repo error", err.Error())
}

func TestGenerateDailyViewReportEmail_Success(t *testing.T) {
	repo := &mockRepo{
		dailyStats: []dto.DailyViewStat{
			{Date: "2024-01-01", ViewCount: 100},
			{Date: "2024-01-02", ViewCount: 200},
		},
	}
	cache := &mockCache{}
	memory := &mockMemory{}
	mailer := &mockMailer{}

	uc := NewViewUsecase(repo, cache, memory, mailer)

	err := uc.GenerateDailyViewReportEmail(context.Background(), "test@example.com", "Test User")

	assert.NoError(t, err)
	assert.True(t, mailer.sendCalled)
	assert.True(t, repo.getDailyViewsCalled)
}

func TestGenerateDailyViewReportEmail_EmptyStats(t *testing.T) {
	repo := &mockRepo{}
	cache := &mockCache{}
	memory := &mockMemory{}
	mailer := &mockMailer{}

	uc := NewViewUsecase(repo, cache, memory, mailer)

	err := uc.GenerateDailyViewReportEmail(context.Background(), "test@example.com", "Test User")
	assert.Error(t, err)
	assert.Equal(t, "no views recorded yet", err.Error())
}

func TestGetViewsByUser(t *testing.T) {
	repo := &mockRepo{
		views: []model.View{{UserID: 1, ProductID: 10}},
	}
	uc := NewViewUsecase(repo, &mockCache{}, &mockMemory{}, &mockMailer{})

	views, err := uc.GetViewsByUser(context.Background(), 1)
	assert.NoError(t, err)
	assert.Len(t, views, 1)
	assert.Equal(t, 10, views[0].ProductID)
}
