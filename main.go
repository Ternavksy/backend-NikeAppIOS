package main

import (
	"log"
	"net/http"
	"nike-backend/internal/server"
)

func main() {
	server.SetupRoutes()
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
