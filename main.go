package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sherpaurgen/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "productapi-log", log.LstdFlags)
	hh := handlers.NewLogger(l)
	servemux := http.NewServeMux()
	servemux.Handle("/", hh)
	http.ListenAndServe(":9090", servemux)
}
