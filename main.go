package main

import (
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
		dataBody, _ := io.ReadAll(r.Body)
		logger.Info("Data received", zap.ByteString("body", dataBody))
	})

	if err != nil {
		log.Fatalf("Cannot initialize zap logger %v", err)
	}

	defer logger.Sync()
	http.ListenAndServe(":9090", r)
}
