package models

// Product holds data belonging to products
type Product struct {
	Id       int
	Name     string
	Price    int
	Category string
}

type InputProduct struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}
