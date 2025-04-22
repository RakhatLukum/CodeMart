package grpc

import (
	"CodeMart/user-service/internal/user/usecase"
	"context"
	userpb "proto/user"

	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
	uc *usecase.UserUsecase
}

func Register(s *grpc.Server, uc *usecase.UserUsecase) {
	userpb.RegisterUserServiceServer(s, &server{uc: uc})
}

func (s *server) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.UserResponse, error) {
	u, err := s.uc.Register(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &userpb.UserResponse{UserId: int32(u.ID), Email: u.Email}, nil
}
func (s *server) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.UserResponse, error) {
	u, err := s.uc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &userpb.UserResponse{UserId: int32(u.ID), Email: u.Email}, nil
}
func (s *server) GetUser(ctx context.Context, req *userpb.UserIdRequest) (*userpb.UserResponse, error) {
	u, err := s.uc.Get(int64(req.UserId))
	if err != nil {
		return nil, err
	}
	return &userpb.UserResponse{UserId: int32(u.ID), Email: u.Email}, nil
}
