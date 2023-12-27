package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

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

type Products []*Product

// notice that its for array "[]*Product"
func (p *Products) ToJSON(rw io.Writer) error {
	e := json.NewEncoder(rw) //returns new eoncoder
	return e.Encode(p)       //this will encode , if not it will return errror
}

// method for "Product" notice that its singluar product
func (p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}
func getNextID() int {
	lastproduct := productList[len(productList)-1]
	return lastproduct.ID + 1
}
func UpdateProduct(id int, p *Product) error {
	fp, position, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[position] = fp
	return nil
}

var ErrProductNotFound = fmt.Errorf("product not found")

func findProduct(id int) (*Product, int, error) {
	for position, p := range productList {
		if p.ID == id {
			return p, position, nil
		}
	}
	return nil, -1, ErrProductNotFound
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
