package main

import "fmt"

var currentID int

var products Products

func init() {
	RepoCreateProduct(
		Product{
			Name: "Monster",
			Category: "Drink",
			Inventory: 10,
			Price: 165,
		},
	)
	RepoCreateProduct(
		Product{
			Name: "RedBull",
			Category: "Drink",
			Inventory: 5,
			Price: 100,
		},
	)
}

// RepoFindProduct function to get product details
func RepoFindProduct(id int) Product {
	for _, t := range products {
		if t.ID == id {
			return t
		}
	}
	// return empty Product if not found
	return Product{}
}

// RepoCreateProduct function to add product
func RepoCreateProduct(t Product) Product {
	currentID ++
	t.ID = currentID
	products = append(products, t)
	return t
}

// RepoDestroyProduct to delete a product using id
func RepoDestroyProduct(id int) error {
    for i, t := range products {
        if t.ID == id {
            products = append(products[:i], products[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
