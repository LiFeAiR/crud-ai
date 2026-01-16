package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateOrganization создает новую организацию
func (bh *BaseHandler) CreateOrganization(ctx context.Context, in *api_pb.OrganizationCreateRequest) (out *api_pb.Organization, err error) {
	// Проверяем входные данные
	if in == nil || in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid argument")
	}

	// Преобразуем запрос в модель
	org := models.Organization{
		Name: in.Name,
	}

	// Используем репозиторий для создания организации
	dbOrg, err := bh.orgRepo.CreateOrganization(ctx, &org)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to create organization")
	}

	// Возвращаем ответ
	return &api_pb.Organization{
		Id:   int32(dbOrg.ID),
		Name: dbOrg.Name,
	}, nil
}
