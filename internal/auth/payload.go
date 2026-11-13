package auth

type LoginResponse struct { //описываем возвращаемую структуру НАМИ
	Token string `json:"token"`
}

type LoginRequest struct { //то, что приходит и как его интерпретировать
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}
