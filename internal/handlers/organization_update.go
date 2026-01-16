package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateOrganization обновляет информацию об организации
func (bh *BaseHandler) UpdateOrganization(ctx context.Context, in *api_pb.OrganizationUpdateRequest) (out *api_pb.Organization, err error) {
	// Проверяем входные данные
	if in == nil || in.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Преобразуем запрос в модель
	org := models.Organization{
		ID:   int(in.Id),
		Name: in.Name,
	}

	// Используем репозиторий для обновления организации
	err = bh.orgRepo.UpdateOrganization(ctx, &org)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to update organization")
	}

	// Возвращаем ответ
	return &api_pb.Organization{
		Id:   int32(org.ID),
		Name: org.Name,
	}, nil
}
