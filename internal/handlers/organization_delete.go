package handlers

import (
	"context"

	"github.com/LiFeAiR/crud-ai/pkg/server/grpc"
)

// DeleteOrganization удаляет организацию
func (bh *BaseHandler) DeleteOrganization(ctx context.Context, req *grpc.Id) (*grpc.Empty, error) {
	// Используем репозиторий для удаления организации
	err := bh.orgRepo.DeleteOrganization(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}

	// Возвращаем пустой ответ
	return &grpc.Empty{}, nil
}
