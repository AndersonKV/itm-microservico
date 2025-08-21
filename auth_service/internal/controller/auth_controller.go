package controller

import (
	"github.com/AndersonKV/instagram-microservice/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Login(login, password string) (string, error) {
	return c.authService.Login(login, password)
}
