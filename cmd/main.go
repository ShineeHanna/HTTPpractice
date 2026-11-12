package main

import (
	"fmt"
	"net/http"
	"pack/configs"
	"pack/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()

	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})
	// Запускаем сервер
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	fmt.Println("Starting server at port 8080")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server:", nil)
	}
}
