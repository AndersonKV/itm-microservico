package service

import (
	"github.com/AndersonKV/instagram-microservice/internal/models"
	repository "github.com/AndersonKV/instagram-microservice/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
}

// APENAS UMA DECLARAÇÃO DO NewAuthService!
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     name,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(user)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.FindById(id)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
