// /internal/handler/auth/auth_handlers.go
package auth

import (
	"net/http"

	"github.com/AndersonKV/instagram-microservice/internal/dto"
	"github.com/AndersonKV/instagram-microservice/internal/repository/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	authService *auth.AuthService // ✅ auth.AuthService (pacote completo)
}

func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

var validate = validator.New()

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validação falhou: " + err.Error()})
		return
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
