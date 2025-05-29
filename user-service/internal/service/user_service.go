package service

import (
	"context"

	"user-service/internal/model"
	"user-service/internal/repository"
)

// Переэкспортируем ошибки из репозитория
var (
	ErrEmailTaken   = repository.ErrEmailTaken
	ErrInvalidCreds = repository.ErrInvalidCreds
	ErrUserNotFound = repository.ErrUserNotFound
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, email, password string) (string, error) {
	// создаём пользователя, возвращаем hex-строку ID
	id, err := s.repo.Create(ctx, &model.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (string, error) {
	u, err := s.repo.FindByEmailAndPassword(ctx, email, password)
	if err != nil {
		return "", err
	}
	return u.ID.Hex(), nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*model.User, error) {
	return s.repo.FindByID(ctx, id)
}
