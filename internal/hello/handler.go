package hello

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{} //структура, чтобы мы могли решистрировать определенные методы

func NewHandler(router *http.ServeMux) { //уже обработчик(типо в нём всё уже регистрируем)
	handler := HelloHandler{} //создаем структруу для использования её методов

	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc { //методы для структуры handler
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello")
	}
}
