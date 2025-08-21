package user

import (
	"github.com/AndersonKV/instagram-microservice/internal/models"
)

type UserRepository interface {
	Create(user models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByUsernameOrEmail(login string) (*models.User, error)
	FindById(id int) (*models.User, error)
	Delete(id int) error
}
