package utils

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func GenerateToken(auth *jwtauth.JWTAuth, userUUID string, expiresIn time.Duration) (string, error) {
	_, tokenString, err := auth.Encode(map[string]any{
		"user_id": userUUID,
		"exp":     jwtauth.ExpireIn(expiresIn),
		"iat":     time.Now().Unix(),
	})
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
