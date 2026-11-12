package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Запускаем сервер
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting the server:", nil)
	}
}
