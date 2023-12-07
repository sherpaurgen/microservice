package handlers

import (
	"encoding/json"
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

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	rw.Header().Add("Content-Type", "application/json")
	//d, err := json.Marshal(lp) //it is standard but its slower than JSON Enconde
	e := json.NewEncoder(rw) //returns new encoder
	err := e.Encode(lp)
	if err != nil {
		http.Error(rw, "unable to marshal", http.StatusInternalServerError)
	}

	//rw.Write()
}
