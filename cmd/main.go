package main

import (
	"fmt"
	"net/http"
	"pack/configs"
	"pack/internal/auth"
	"pack/internal/link"
	"pack/pkg/db"
	"pack/pkg/middleware"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	//Repositories
	linkRepository := link.NewLinkRepositiry(db)
	//Handler
	auth.NewHandler(router, auth.HandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	// Запускаем сервер
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}
	fmt.Println("Starting server at port 8080")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Error starting the server:", nil)
	}
}
