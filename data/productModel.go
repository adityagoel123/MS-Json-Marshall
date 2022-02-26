package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	NAME        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Sku         string  `json:"sku"`
	CreatedOn   string  `json:"createdOn"`
	UpdatedOn   string  `json:"-"`  // Drop this field from ResponseJSON.
	DeletedOn   string  `json:"-"`  // Drop this field from ResponseJSON.
}

func GetProducts() []*Product {
	return productList
}

var productList = []*Product{

	&Product{
		ID:          1,
		NAME:        "LATTE",
		Description: "Frothy Milky Coffee",
		Price:       2.45,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		NAME:        "ESSPRESSO",
		Description: "Strong coffe without Milk",
		Price:       1.99,
		Sku:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
