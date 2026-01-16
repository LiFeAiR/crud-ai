package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
	"github.com/LiFeAiR/crud-ai/internal/utils"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateUser общий метод для обновления пользователя
func (bh *BaseHandler) UpdateUser(ctx context.Context, in *api_pb.UserUpdateRequest) (out *api_pb.User, err error) {
	// Проверяем входные данные
	if in == nil || in.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Преобразуем запрос в модель
	user := models.User{
		ID:           int(in.Id),
		Name:         in.Name,
		Email:        in.Email,
		Password:     in.Password,
		Organization: utils.Ptr(in.Organization),
	}

	// Используем репозиторий для обновления пользователя
	if err := bh.userRepo.UpdateUser(ctx, &user); err != nil {
		return nil, status.Error(codes.Internal, "Failed to update user")
	}

	// Возвращаем ответ
	return &api_pb.User{
		Id:           int32(user.ID),
		Name:         user.Name,
		Email:        user.Email,
		Organization: utils.FromPtr(user.Organization),
	}, nil
}
