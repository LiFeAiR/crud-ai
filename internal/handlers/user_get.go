package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/utils"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetUser общий метод для получения пользователя по ID
func (bh *BaseHandler) GetUser(ctx context.Context, in *api_pb.Id) (out *api_pb.User, err error) {
	// Проверяем входные данные
	if in == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Используем репозиторий для получения пользователя
	user, err := bh.userRepo.GetUserByID(ctx, int(in.Id))
	if err != nil {
		return nil, status.Error(codes.NotFound, "User not found")
	}

	// Возвращаем ответ
	return &api_pb.User{
		Id:           int32(user.ID),
		Name:         user.Name,
		Email:        user.Email,
		Organization: utils.FromPtr(user.Organization),
	}, nil
}
