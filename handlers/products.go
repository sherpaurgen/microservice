package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sherpaurgen/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.l.Println(r)
		regx := regexp.MustCompile(`/([0-9]+)`)
		g := regx.FindAllStringSubmatch(r.URL.Path, -1)
		if g == nil {
			http.Error(rw, "Invalid URLs, empty url part", http.StatusBadRequest)
			return
		}

		if len(g) != 1 {
			http.Error(rw, "Invalid URLs, no match int", http.StatusBadRequest)
			return
		}
		p.l.Println(g)
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("failed converting Atoi - ", idString)
		}
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		m := make(map[string]string)
		m["id"] = idString

		reply, err := json.Marshal(m)
		rw.Write(reply)
		p.l.Println("Got id - ", id)
		return

	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts() //a list product
	rw.Header().Add("Content-Type", "application/json")
	//d, err := json.Marshal(lp) //it is standard but its slower than JSON Enconde
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal", http.StatusInternalServerError)
	}

	//rw.Write()
}
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle post product")
	prod := &data.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
	p.l.Printf("Prod: %v", prod)
}
func (p Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
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
