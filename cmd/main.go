package main

import (
	"fmt"
	"net/http"
	"pack/configs"
	"pack/internal/auth"
	"pack/internal/link"
	"pack/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)
	router := http.NewServeMux()
	//Handler

	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{})
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
