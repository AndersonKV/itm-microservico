package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/AndersonKV/instagram-microservice/internal/models"
	repository "github.com/AndersonKV/instagram-microservice/internal/repository/user"
	"github.com/AndersonKV/instagram-microservice/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepository
}

// // Variável global (menos recomendado, mas funciona)
var jwtKey = []byte("secret_key") // TODO: Trocar por variável de ambiente

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Login - usuário pode enviar username OU email
func (s *AuthService) Login(login string, password string) (string, error) {
	// Busca usuário por username OU email
	user, err := s.userRepo.FindByUsernameOrEmail(login)
	if err != nil {
		return "", errors.New("erro ao buscar usuário")
	}
	if user == nil {
		return "", errors.New("credenciais inválidas")
	}

	// Verifica senha (com bcrypt)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

	// Gera token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("erro ao gerar token")
	}

	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenString string) (*models.User, error) {
	// Usa a utility function para parsear
	token, err := utils.ParseToken(tokenString)
	if err != nil {
		return nil, fmt.Errorf("token inválido: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims inválidas")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("user_id inválido no token")
	}
	userID := int(userIDFloat)

	// Busca usuário no banco
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}
