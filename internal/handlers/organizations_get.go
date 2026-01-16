package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
	api_pb "github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrganizationsResponse struct {
	Data []*models.Organization `json:"data"`
}

// GetOrganizations получает список организаций с пагинацией
func (bh *BaseHandler) GetOrganizations(
	ctx context.Context,
	in *api_pb.ListRequest,
) (out *api_pb.OrganizationsResponse, err error) {
	// Устанавливаем значения по умолчанию
	limit := 10
	offset := 0

	// Парсим limit
	if in.Limit > 0 && in.Limit < 100 {
		limit = int(in.GetLimit())
	}

	// Парсим offset
	if in.Offset > 0 {
		offset = int(in.GetOffset())
	}

	// Используем репозиторий для получения организаций
	organizations, err := bh.orgRepo.GetOrganizations(ctx, limit, offset)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get organizations")
	}

	// Отправляем ответ клиенту
	data := make([]*api_pb.Organization, len(organizations))
	for i, u := range organizations {
		data[i] = &api_pb.Organization{
			Id:   int32(u.ID),
			Name: u.Name,
		}
	}

	return &api_pb.OrganizationsResponse{
		Data: data,
	}, nil
}
