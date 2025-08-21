package user

import (
	"database/sql"
	"fmt"

	"github.com/AndersonKV/instagram-microservice/internal/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE username = $1`
	err := r.db.Get(&user, query, username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByUsernameOrEmail(login string) (*models.User, error) {
	var user models.User
	query := `
        SELECT id, name, username, email, password, profile_pic, description, created_at, updated_at 
        FROM users 
        WHERE username = $1 OR email = $2
    `

	err := r.db.Get(&user, query, login, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user models.User) error {
	query := `
		INSERT INTO users (name, username, email, password, profile_pic, description, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
	`
	_, err := r.db.Exec(query, user.Name, user.Username, user.Email, user.Password, user.ProfilePic, user.Description)
	return err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, name, username, email, password, profile_pic, description, created_at, updated_at 
		FROM users WHERE email = $1
	`
	var user models.User
	err := r.db.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // usuário não encontrado
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindById(id int) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE id = $1`
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %d not found", id)
	}

	return nil
}
