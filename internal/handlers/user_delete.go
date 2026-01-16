package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteUser общий метод для удаления пользователя
func (bh *BaseHandler) DeleteUser(ctx context.Context, req *grpc.Id) (*grpc.Empty, error) {
	// Проверяем входные данные
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Используем репозиторий для удаления пользователя
	if err := bh.userRepo.DeleteUser(ctx, int(req.GetId())); err != nil {
		return nil, status.Error(codes.Internal, "Failed to delete user")
	}

	// Возвращаем пустой ответ
	return &grpc.Empty{}, nil
}
