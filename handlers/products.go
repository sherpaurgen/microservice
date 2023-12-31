package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sherpaurgen/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
	//its constructor for Products structor for example
	// if there is struct type person struct { name string age  int }
	//then a new person will be s := person{name: "Sean", age: 50} or person{"Sean",50}
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
	// prod := &data.Product{}
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Failed to unmarshal JSON to product data", http.StatusBadRequest)
	}
	data.AddProduct(&prod)
	p.l.Printf("Prod: %v", prod)
}

func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "productid")) //extract product id from url
	if err != nil {
		http.Error(rw, "Invalid url in put request", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	status, err := data.UpdateProduct(id, &prod)
	if status {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		jsonData := map[string]string{"response": "update successful"}
		jsonResponse, _ := json.Marshal(jsonData)
		rw.Write(jsonResponse)
	}
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{} //create the product obj

		err := prod.FromJson(r.Body) //gets httpbody + json-encode and save to prod
		p.l.Printf("problem body %v", prod)
		if err != nil {
			p.l.Printf("problem body %v", r.Body)
			http.Error(rw, "Failed to unmarshal JSON to product data: UpdateProduct", http.StatusBadRequest)

		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx) //adding the updated context to request
		next.ServeHTTP(rw, r)
	})
}
