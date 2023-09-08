package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"my-auth-service/internal/middleware/dependency"
)

const secretKey = "some-secret-key"

type Claims struct {
	jwt.RegisteredClaims
	UserID string
}

func SignToken(repository dependency.Repository, claims *Claims) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(2 * time.Hour)) // TODO hard-coded 2-hour expiration time
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	if err = repository.CreateToken(tokenString); err != nil {
		return "", err
	}
	return tokenString, err
}

func ParseToken(repository dependency.Repository, tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		if repoErr := repository.DeleteToken(tokenString); repoErr != nil {
			return nil, repoErr
		}
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if err = repository.CheckToken(tokenString); err != nil {
			return nil, err
		}
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
