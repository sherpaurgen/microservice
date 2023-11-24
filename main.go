package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		logger.Info("Get received at /")
		dataBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops", http.StatusBadRequest)
			return
		}
		logger.Info("Data received", zap.ByteString("body", dataBody))
		fmt.Fprintf(rw, "HelloWorld %s", dataBody)
	})

	if err != nil {
		log.Fatalf("Cannot initialize zap logger %v", err)
	}

	defer logger.Sync()
	http.ListenAndServe(":9090", r)
}
