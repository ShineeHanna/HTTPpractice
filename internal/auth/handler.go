package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pack/configs"
	"pack/pkg"
	"regexp"

	"github.com/go-playground/validator/v10"
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
	return func(w http.ResponseWriter, req *http.Request) {
		//прочитать body
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			pkg.Json(w, err.Error(), 402)
		}
		reg, _ := regexp.Compile(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z{2,}]`)
		if !reg.MatchString(payload.Email) { //проверка на корректно введенный email
			pkg.Json(w, "Wrong email", 402)
			return
		}
		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {
			pkg.Json(w, "Wrong email", 402)
			return
		}
		fmt.Println(payload)
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
