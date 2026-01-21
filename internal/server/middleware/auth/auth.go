package auth

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/LiFeAiR/crud-ai/internal/utils"
)

type ctxKey string

var (
	ErrInvalidToken = errors.New("invalid token")

	UserIDKey  ctxKey = "UserID"
	ErrorKey   ctxKey = "Error"
	IsAdminKey ctxKey = "IsAdmin"

	IsAdmin = "admin"
)

// New creates new auth middleware.
func New(
	appSecret string,
) func(next http.Handler) http.Handler {
	const op = "middleware.auth"

	// Возвращаем функцию-обработчик
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Получаем JWT-токен из запроса
			tokenStr := extractBearerToken(r)
			if tokenStr == "" {
				// It's ok, if user is not authorized
				next.ServeHTTP(w, r)
				return
			}

			// Парсим и валидируем токен, использая appSecret
			claims, err := utils.ValidateJWT(tokenStr, appSecret)
			if err != nil {
				log.Printf("%s, failed to parse token: %v", op, err)

				// But if token is invalid, we shouldn't handle request
				ctx := context.WithValue(r.Context(), ErrorKey, ErrInvalidToken)
				next.ServeHTTP(w, r.WithContext(ctx))

				return
			}

			log.Printf("%s, user authorized: %v", op, claims.UserID)

			// Полученны данные сохраняем в контекст,
			// откуда его смогут получить следующие хэндлеры.
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, IsAdminKey, isAdmin(claims.Permissions))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isAdmin(permissions []string) bool {
	for _, permission := range permissions {
		if permission == IsAdmin {
			return true
		}
	}

	return false
}
