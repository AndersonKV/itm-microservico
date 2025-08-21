// main.go
package main

import (
	"github.com/AndersonKV/instagram-microservice/internal/controller"
	"github.com/AndersonKV/instagram-microservice/internal/db"
	authHandler "github.com/AndersonKV/instagram-microservice/internal/handler/auth"
	userHandler "github.com/AndersonKV/instagram-microservice/internal/handler/user"
	authRepo "github.com/AndersonKV/instagram-microservice/internal/repository/auth"
	"github.com/AndersonKV/instagram-microservice/internal/repository/user"
	"github.com/AndersonKV/instagram-microservice/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB()

	// ✅ Cadeia completa de dependências
	userRepo := user.NewUserRepository(database)

	// User flow
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	userHandler := userHandler.NewUserHandler(userController) // ✅ Agora tipos compatíveis

	// Auth flow
	authService := authRepo.NewAuthService(userRepo)
	authHandler := authHandler.NewAuthHandler(authService)

	// Rotas
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/login", authHandler.Login)
		}

		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/create", userHandler.Create)
			userRoutes.GET("/:id", userHandler.FindById)
			userRoutes.DELETE("/:id", userHandler.Delete)
		}
	}

	r.Run(":8080")
}
