package auth

import (
	"fmt"
	"net/http"
	"pack/configs"
	"pack/pkg"
)

type Handler struct { //структура, чтобы мы могли решистрировать определенные методы
	*configs.Config
}

type HandlerDeps struct { //структура в структуре, берем все наши зависимости
	*configs.Config
}

func NewHandler(router *http.ServeMux, deps HandlerDeps) { //уже обработчик(типо в нём всё уже регистрируем)
	handler := Handler{ //создаем структyру для использования её методов
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc { //методы для структуры handler
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
		data := LoginResponse{
			Token: "123",
		}
		pkg.Json(w, data, 200)
	}
}
func (handler Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Register")
	}
}
