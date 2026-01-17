package handlers

import (
	"context"
	"errors"
	"testing"

	"github.com/LiFeAiR/crud-ai/internal/models"
	"github.com/LiFeAiR/crud-ai/pkg/server/grpc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock UserRepository для тестирования
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CheckPassword(ctx context.Context, userID int, password string) (bool, error) {
	args := m.Called(ctx, userID, password)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) GetUsers(ctx context.Context, limit, offset int) ([]*models.User, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]*models.User), args.Error(1)
}

func (m *MockUserRepository) InitDB() error {
	panic("implement me")
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	args := m.Called(ctx, user)

	if u, ok := args.Get(0).(*models.User); ok {
		return u, args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// TestBaseHandler_GetUser тестирует метод GetUser базового обработчика
func TestBaseHandler_GetUser(t *testing.T) {
	ctx := context.Background()
	// Test 1: Успешное получение пользователя
	t.Run("GetUserSuccess", func(t *testing.T) {
		// Создаем мок репозиторий
		mockRepo := new(MockUserRepository)

		// Подготавливаем тестовую организацию
		testUser := &models.User{
			ID:   1,
			Name: "Test Org",
		}

		// Определяем ожидаемое поведение мока
		mockRepo.On("GetUserByID", ctx, 1).Return(testUser, nil)

		// Создаем базовый обработчик с моком
		baseHandler := &BaseHandler{
			userRepo: mockRepo,
		}

		// Вызываем метод GetUser
		u, err := baseHandler.GetUser(ctx, &grpc.Id{Id: 1})

		// Проверяем результат
		assert.NoError(t, err)
		assert.NotNil(t, u)
		assert.Equal(t, int32(1), u.Id)
		assert.Equal(t, "Test Org", u.Name)

		// Проверяем, что мок был вызван правильно
		mockRepo.AssertExpectations(t)
	})

	// Test 2: Некорректный аргумент (nil)
	t.Run("GetUserNilArgument", func(t *testing.T) {
		// Создаем базовый обработчик
		baseHandler := &BaseHandler{
			orgRepo: nil,
		}

		// Вызываем метод GetUser с nil аргументом
		org, err := baseHandler.GetUser(ctx, nil)

		// Проверяем результат
		assert.Error(t, err)
		assert.Nil(t, org)
	})

	// Test 3: Пользователь не найден
	t.Run("GetUserNotFound", func(t *testing.T) {
		// Создаем мок репозиторий
		mockRepo := new(MockUserRepository)

		// Определяем ожидаемое поведение мока - возвращаем ошибку
		mockRepo.On("GetUserByID", ctx, 1).Return((*models.User)(nil), errors.New("user not found"))

		// Создаем базовый обработчик с моком
		baseHandler := &BaseHandler{
			userRepo: mockRepo,
		}

		// Вызываем метод GetUser
		org, err := baseHandler.GetUser(ctx, &grpc.Id{Id: 1})

		// Проверяем результат
		assert.Error(t, err)
		assert.Nil(t, org)

		// Проверяем, что мок был вызван правильно
		mockRepo.AssertExpectations(t)
	})
}
