package auth

type LoginResponse struct { //описываем возвращаемую структуру
	Token string `json:"token"`
}
