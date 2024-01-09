package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sherpaurgen/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "Productapi-Log", log.LstdFlags)
	ph := handlers.NewProducts(l) //a product handler

	//servemux := http.NewServeMux()
	servemux := chi.NewRouter()

	// servemux.Use(middleware.Logger)
	servemux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	apiRouter := chi.NewRouter()
	//apiRouter.Use(ph.MiddlewareProductValidation)
	apiRouter.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://127.0.0.1:8080", "http://127.0.0.1:8080"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	apiRouter.Get("/products", ph.GetProducts)

	apiRouter.With(ph.MiddlewareProductValidation).Put("/products/{productid}", ph.UpdateProduct)

	apiRouter.With(ph.MiddlewareProductValidation).Post("/products", ph.AddProduct)

	servemux.Mount("/api", apiRouter)
	//with above it will work for 127.0.0.1:8080/api/items , it will NOT work for 127.0.0.1:8080/api/items/ , notice trailing slash
	servemux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	// servemux.Handle("/about", gg)
	webserver := &http.Server{
		Addr:         ":8080",
		Handler:      servemux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		err := webserver.ListenAndServe()
		fmt.Println("Server listening on :8080")
		// err := http.ListenAndServe(":8080", servemux)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	//signal.Notify(sigChan, os.Kill)
	sig := <-sigChan //this is blocking operation --reading from channel
	l.Println("Received terminate graceful shutdown", sig)
	d := time.Now().Add(30 * time.Second)
	// Create a context with the calculated deadline
	//context.Background() returns the base context that is empty (its like blank canvas /template) and its passed to context.WithDeadline
	tc, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	webserver.Shutdown(tc)

}
