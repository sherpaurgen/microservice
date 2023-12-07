package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sherpaurgen/microservice/handlers"
)

func main() {
	l := log.New(os.Stdout, "productapi-log", log.LstdFlags)
	// hh := handlers.NewLogger(l)
	// gg := handlers.NewGG(l)
	ph := handlers.NewProducts(l) //a product handler
	servemux := http.NewServeMux()
	servemux.Handle("/", ph)
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
