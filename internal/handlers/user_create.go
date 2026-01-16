package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
	"github.com/LiFeAiR/crud-ai/internal/utils"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateUser общий метод для создания пользователя
func (bh *BaseHandler) CreateUser(ctx context.Context, in *api_pb.UserCreateRequest) (out *api_pb.User, err error) {
	user := models.User{
		Name:         in.Name,
		Email:        in.Email,
		Password:     in.Password,
		Organization: utils.Ptr(in.Organization),
	}

	// Validate request
	if false {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Используем репозиторий для создания пользователя
	dbUser, err := bh.userRepo.CreateUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create user")
	}

	// Send response
	return &api_pb.User{
		Id:           int32(dbUser.ID),
		Name:         user.Name,
		Email:        user.Email,
		Organization: utils.FromPtr(user.Organization),
	}, nil
}
