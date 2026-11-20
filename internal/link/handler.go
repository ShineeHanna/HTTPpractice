package link

import (
	"fmt"
	"net/http"
	"pack/pkg"
	"pack/pkg/req"
)

type LinkHandler struct { //структура для зависимостей
	LinkRepository *LinkRepository
}
type LinkHandlerDeps struct {
	LinkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) { //уже обработчик(типо в нём всё уже регистрируем)
	handler := &LinkHandler{ //создаем структруу для использования её методов
		LinkRepository: deps.LinkRepository,
	}

	router.HandleFunc("POST /link/login", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete()) //динамические параметры {id}
	router.HandleFunc("GET /{hash}", handler.GoTo())
}

// создание, обновление,удаление, чтение
func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r) //записать пришедший ответ в переменную body типа LinkCreateRequest
		if err != nil {
			return
		}
		link := NewLink(body.Url)                               //достать ссылку из объекта
		createdLink, err := handler.LinkRepository.Create(link) //созданная новая ссылка
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		pkg.Json(w, createdLink, 201)
	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") //получить название по переменной
		fmt.Println(id)
	}
}
func (handler *LinkHandler) GoTo() http.HandlerFunc { //перейти по ссылке
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect) //редирект
	}
}
