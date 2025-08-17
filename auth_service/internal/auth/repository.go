package auth

import (
	"database/sql"

	"github.com/AndersonKV/auth_service/internal/models"
)

 

type AuthRepository struct {
    db *sql.DB
}
 


type AuthRepositoryInterface interface {
    CreateUser(user models.User) error
    GetUserByEmail(email string) (models.User, error)
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
    return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) error {
    _, err := r.db.Exec(
        "INSERT INTO users (username, email, password, profile_pic, description) VALUES (:1, :2, :3)",
        user.Username, user.Email, user.Password,
    )
    return err
}


func (r *AuthRepository) GetUserByEmail(email string) (models.User, error) {
    var user models.User
    row := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE email=:1", email)
    err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
    return user, err
}
