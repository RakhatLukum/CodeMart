package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserContextKey is the key used to store user ID in context
type UserContextKey struct{}

// UserContextMiddleware validates user context in incoming requests
func UserContextMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Extract user ID from request
		var userID int32
		switch r := req.(type) {
		case interface{ GetUserId() int32 }:
			userID = r.GetUserId()
		default:
			return nil, status.Error(codes.InvalidArgument, "request must contain user_id")
		}

		// Validate user ID
		if userID <= 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid user_id")
		}

		// Add user ID to context
		ctx = context.WithValue(ctx, UserContextKey{}, userID)

		return handler(ctx, req)
	}
}

// GetUserIDFromContext extracts user ID from context
func GetUserIDFromContext(ctx context.Context) (int32, error) {
	userID, ok := ctx.Value(UserContextKey{}).(int32)
	if !ok {
		return 0, errors.New("user_id not found in context")
	}
	return userID, nil
}
