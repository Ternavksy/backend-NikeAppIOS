package handlers

import (
	"encoding/json"
	"net/http"
	"nike-backend/internal/models"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []models.Product{
		{
			ID:       1,
			Title:    "Air Jordan I",
			Price:    "$180",
			ImageURL: "http://localhost:8080/images/pizdec.png",
		},
		{
			ID:       2,
			Title:    "Air Jordan II",
			Price:    "$190",
			ImageURL: "http://localhost:8080/images/air-jordan1.png",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(products)

}
