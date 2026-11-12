package main

import (
	"fmt"
	"net/http"
	"pack/internal/hello"
)

func main() {
	router := http.NewServeMux()

	hello.NewHandler(router)
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
