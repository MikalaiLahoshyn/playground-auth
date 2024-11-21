package models

type LoginRequest struct {
	Login    string `json:"login" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type LoginResponse struct {
	Token string
}
