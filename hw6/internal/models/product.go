package models

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	ImageUrl    string  `json:"image"`
}
