package auth

import "github.com/AndersonKV/instagram-microservice/internal/models"

type AuthRepository interface {
	Login(login string) (string, error)
	ValidateToken(tokenString string) (*models.User, error)
	RefreshToken(tokenString string) (string, error)
}
