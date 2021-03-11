package main



// Product model
type Product struct {
	ID        int       `json:"id"`
	Name string `json:"name"`
	Category string `json:"category"`
	Inventory int16 `json:"inventory"`
	Price int16 `json:"price"`
}

//Products array/slice composed of Product's
type Products []Product

