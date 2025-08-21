package user

import (
	"net/http"
	"strconv"

	"github.com/AndersonKV/instagram-microservice/internal/controller"
	"github.com/AndersonKV/instagram-microservice/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userController *controller.UserController
}

func NewUserHandler(userController *controller.UserController) *UserHandler {
	return &UserHandler{userController: userController}
}

var validate = validator.New()

func (h *UserHandler) Create(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validação falhou"})
		return
	}

	// Chama o CONTROLLER em vez do service diretamente
	if err := h.userController.Create(req.Name, req.Username, req.Email, req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o CONTROLLER
	if err := h.userController.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado"})
}

func (h *UserHandler) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Chama o CONTROLLER
	user, err := h.userController.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
