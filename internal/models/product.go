package models

type Product struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    string `json:"price"`
	ImageURL string `json:"image_url"`
}
