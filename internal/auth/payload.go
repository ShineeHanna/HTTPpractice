package auth

type LoginResponse struct { //описываем возвращаемую структуру НАМИ
	Token string `json:"token"`
}

type LoginRequest struct { //то, что приходит и как его интерпретировать
	Email    string `json:"email"`
	Password string `json:"password"`
}
