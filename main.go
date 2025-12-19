package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	ImageURL string `json:"image_url"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{
			ID:       1,
			Title:    "Air Jordan I",
			Price:    "$180",
			ImageURL: "/images/pizdec.png",
		},
		{
			ID:       2,
			Title:    "Air Jordan II",
			Price:    "$190",
			ImageURL: "/images/air-jordan1.png",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.Handle(
		"/images/",
		http.StripPrefix(
			"/images/",
			http.FileServer(http.Dir("./images")),
		),
	)
	http.HandleFunc("/products", productsHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
