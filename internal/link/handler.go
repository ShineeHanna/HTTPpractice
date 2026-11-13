package link

import (
	"net/http"
)

type LinkHandler struct {
}
type LinkHandlerDeps struct {
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) { //уже обработчик(типо в нём всё уже регистрируем)
	handler := &LinkHandler{} //создаем структруу для использования её методов

	router.HandleFunc("POST /link/login", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
	router.HandleFunc("GET /{alias}", handler.GoTo())
}

// создание, обновление,удаление, чтение
func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
