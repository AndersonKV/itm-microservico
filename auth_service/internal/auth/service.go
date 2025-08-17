package auth

import (
	"errors"
	"time"

	"github.com/AndersonKV/auth_service/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("secret_key") // trocar para env

type AuthService struct {
    repo AuthRepositoryInterface
}

func NewAuthService(repo AuthRepositoryInterface) *AuthService {
    return &AuthService{repo: repo}
}


func (s *AuthService) Register(name, username, email, hashedPassword string) error {
   user := models.User{
    Username:    username,
    Email:       email,
    Password:    hashedPassword,
    ConfirmPassword:  hashedPassword,
    ProfilePic:  "",
    Name: name,
    Description: "",
}

    return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
    user, err := s.repo.GetUserByEmail(email)
    if err != nil {
        return "", err
    }
    // Aqui você deveria comparar hash da senha
    if user.Password != password {
        return "", errors.New("senha inválida")
    }

    // Cria JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })
    tokenString, _ := token.SignedString(jwtKey)
    return tokenString, nil
}
