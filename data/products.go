package data

import "time"

type Product struct {
	ID          int     `json:"id"`   //this is called struct tag or field tags
	Name        string  `json:"name"` //
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` // omit this completely on output
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       10.2,
		SKU:         "122-GR-3k",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Black roasted coffee",
		Price:       12.2,
		SKU:         "213-GS-2k",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func GetProducts() []*Product {
	return productList
}
