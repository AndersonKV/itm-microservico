package dto

type RegisterRequest struct {
	Name            string `json:"name" validate:"required,min=2,max=50"`
	Username        string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	ProfilePic  *string `json:"profile_pic"`
	Description *string `json:"description"`
}
