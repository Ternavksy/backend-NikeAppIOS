package server

import (
	"net/http"
	"nike-backend/internal/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/products", handlers.ProductsHandler)

	http.Handle(
		"/images",
		http.StripPrefix(
			"images/",
			http.FileServer(http.Dir("./images")),
		),
	)
}
