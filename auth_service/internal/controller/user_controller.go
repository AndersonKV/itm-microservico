package controller

import (
	"github.com/AndersonKV/instagram-microservice/internal/models"
	"github.com/AndersonKV/instagram-microservice/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) Create(name, username, email, password string) error {
	return c.userService.Create(name, username, email, password)
}

func (c *UserController) GetUserByID(id int) (*models.User, error) {
	return c.userService.GetUserByID(id)
}

func (c *UserController) DeleteUser(id int) error {
	return c.userService.DeleteUser(id)
}
