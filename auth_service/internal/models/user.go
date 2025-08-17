package models

type User  struct {
	ID              int64  `db:"id"`
	Name            string `db:"name"`
	Username        string `db:"username"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	ConfirmPassword string `db:"confirmpassword"`
	ProfilePic      string `db:"profile_pic"`
	Description     string `db:"description"`
}
 