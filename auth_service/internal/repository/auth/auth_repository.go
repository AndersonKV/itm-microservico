package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/AndersonKV/instagram-microservice/internal/models"
	repository "github.com/AndersonKV/instagram-microservice/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// ✅ ValidateToken DEVE estar no Service!
func (s *AuthService) ValidateToken(tokenString string) (*models.User, error) {
	// 1. Parse do token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método inesperado: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("token inválido: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	// 2. Extrai claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("claims inválidas")
	}

	// 3. Obtém ID do usuário
	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("user_id inválido")
	}
	userID := int(userIDFloat)

	// 4. ✅ Busca usuário usando o UserRepository
	user, err := s.userRepo.FindById(userID)
	if err != nil {
		return nil, errors.New("erro ao buscar usuário")
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}

// ✅ Login também no Service
func (s *AuthService) Login(login string, password string) (string, error) {
	var user *models.User
	var err error

	if strings.Contains(login, "@") {
		user, err = s.userRepo.FindByEmail(login)
	} else {
		user, err = s.userRepo.FindByUsername(login)
	}

	if err != nil || user == nil {
		return "", errors.New("credenciais inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("credenciais inválidas")
	}

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
