package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims расширенные claims для JWT токена
type Claims struct {
	UserID      int      `json:"user_id"`
	Email       string   `json:"email"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

// GenerateJWT генерирует JWT токен для пользователя
func GenerateJWT(secretKey string, userID int, email, name string, permissions []string) (string, error) {
	// Устанавливаем срок действия токена (например, 24 часа)
	expirationTime := time.Now().Add(24 * time.Hour)
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("Error marshaling private key: %w", err)
	}

	claims := &Claims{
		UserID:      userID,
		Email:       email,
		Name:        name,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(pkey)
}

// ValidateJWT проверяет валидность JWT токена
func ValidateJWT(tokenString, secretKey string) (*Claims, error) {
	claims := &Claims{}
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(secretKey))
	if err != nil {
		return nil, fmt.Errorf("Error marshaling private key: %w", err)
	}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return &pkey.PublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

// nolint unused
func showPublicPEM(pkey *rsa.PrivateKey) (string, error) {
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&pkey.PublicKey)
	if err != nil {
		return "", fmt.Errorf("Error marshaling public key: %w", err)
	}

	// 3. Encode the DER bytes to PEM format
	publicKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	})

	return string(publicKeyPem), nil
}
