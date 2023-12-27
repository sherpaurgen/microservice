package handlers

import (
	"log"
	"net/http"

	"github.com/sherpaurgen/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts() //a list product
	rw.Header().Add("Content-Type", "application/json")
	//d, err := json.Marshal(lp) //it is standard but its slower than JSON Enconde
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
	}

	//rw.Write()
}
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle post product")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	p.l.Printf("Prod: %v", prod)
}

func (p Products) UpdateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshall", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
}
