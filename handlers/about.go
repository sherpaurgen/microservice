package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"go.uber.org/zap"
)

type GG struct {
	l *log.Logger
}

func NewGG(l *log.Logger) *GG {
	return &GG{l}
}

func (gg *GG) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	gg.l.Println("aboutpage log:hello world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops from about", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hi from About handler: %s", d)
}
