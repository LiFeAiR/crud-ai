package repository

import (
	"context"

	"github.com/LiFeAiR/crud-ai/internal/models"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id int) error
	GetUsers(ctx context.Context, limit, offset int) ([]*models.User, error)
	InitDB() error
}

// OrganizationRepository интерфейс для работы с организациями
type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, org *models.Organization) (*models.Organization, error)
	GetOrganizationByID(ctx context.Context, id int) (*models.Organization, error)
	UpdateOrganization(ctx context.Context, org *models.Organization) error
	DeleteOrganization(ctx context.Context, id int) error
	GetOrganizations(ctx context.Context, limit, offset int) ([]*models.Organization, error)
	InitDB() error
}
