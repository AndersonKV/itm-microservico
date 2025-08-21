package models

import "time"

type User struct {
	ID              int64     `db:"id"`
	Name            string    `db:"name"`
	Username        string    `db:"username"`
	Email           string    `db:"email"`
	Password        string    `json:"-" db:"password"` // "-" omite do JSON por segurança
	ConfirmPassword string    `db:"-"`
	ProfilePic      *string   `db:"profile_pic"` // default 'default_user.png'
	Description     *string   `db:"description"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
