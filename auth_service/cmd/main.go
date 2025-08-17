package main

import (
	"log"

	"github.com/AndersonKV/auth_service/internal/auth"
	"github.com/AndersonKV/auth_service/internal/db"
	"github.com/gin-gonic/gin"
)


func main() {
    // Conectar ao Oracle
    oracleDB, err := db.ConnectOracle("usuario", "senha", "localhost", "1521", "XE")
    if err != nil {
        log.Fatal(err)
    }

    // Criar repositório e serviço
    repo := auth.NewAuthRepository(oracleDB)
    service := auth.NewAuthService(repo)

    // Inicializar Gin
    r := gin.Default()

    // Rotas do Auth
    r.POST("/register", auth.RegisterHandler(service))
    r.POST("/login", auth.LoginHandler(service))

    // Rodar servidor
    r.Run(":8081") // porta do Auth Service
}
