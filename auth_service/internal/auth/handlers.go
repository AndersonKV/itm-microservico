package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name            string `db:"name"`
	Username        string `db:"username"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	ConfirmPassword string `db:"confirmpassword"`
	ProfilePic      string `db:"profile_pic"`
	Description     string `db:"description"`
}

func RegisterHandler(s *AuthService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req RegisterRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // chama o service para registrar o usuário
        if err := s.Register(req.Name, req.Username, req.Email, req.Password ); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusCreated, gin.H{"message": "usuário criado"})
    }
}
