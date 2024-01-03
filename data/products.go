package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`                       //this is called struct tag or field tags
	Name        string  `json:"name" validate:"required"` //
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // omit this completely on output
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("sku", validateSku)

	return validate.Struct(p)
}

// custom validator func for sku that is registered in Validate() func above
func validateSku(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
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
func UpdateProduct(id int, p *Product) (bool, error) {
	_, position, err := findProduct(id)
	if err != nil {
		return false, err
	}
	p.ID = id
	productList[position] = p
	return true, nil
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
		SKU:         "wxy-woi-poi",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Black roasted coffee",
		Price:       12.2,
		SKU:         "abc-def-ghi",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
