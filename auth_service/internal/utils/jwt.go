package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// JWTKey retorna a chave secreta para JWT
func JWTKey() []byte {
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		// Fallback para desenvolvimento (NÃO usar em produção)
		return []byte("secret_key_development_change_in_production")
	}
	return []byte(key)
}

// CreateToken cria um novo token JWT
func CreateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expira em 24h
	})

	return token.SignedString(JWTKey())
}

// ParseToken valida e parseia um token JWT
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWTKey(), nil
	})
}
