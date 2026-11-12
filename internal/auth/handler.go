package auth

import (
	"fmt"
	"net/http"
)

type Handler struct{} //структура, чтобы мы могли решистрировать определенные методы

func NewHandler(router *http.ServeMux) { //уже обработчик(типо в нём всё уже регистрируем)
	handler := Handler{} //создаем структруу для использования её методов

	router.HandleFunc("/auth/login", handler.Login())
	router.HandleFunc("/auth/register", handler.Register())
}

func (handler *Handler) Login() http.HandlerFunc { //методы для структуры handler
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Login")
	}
}
func (handler Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Register")
	}
}
