package handler

import (
	"context"

	"user-service/internal/service"
	userpb "user-service/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	userpb.UnimplementedUserServiceServer
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.UserResponse, error) {
	id, err := h.svc.Register(ctx, req.Email, req.Password)
	if err != nil {
		if err == service.ErrEmailTaken {
			return nil, status.Errorf(codes.AlreadyExists, "email already taken")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &userpb.UserResponse{
		UserId: id, // id — строка (hex)
		Email:  req.Email,
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.UserResponse, error) {
	id, err := h.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		if err == service.ErrInvalidCreds {
			return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &userpb.UserResponse{
		UserId: id, // строковый ID
		Email:  req.Email,
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *userpb.UserIdRequest) (*userpb.UserResponse, error) {
	user, err := h.svc.GetUser(ctx, req.UserId)
	if err != nil {
		if err == service.ErrUserNotFound {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &userpb.UserResponse{
		UserId: req.UserId,
		Email:  user.Email,
	}, nil
}
