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
	hh := handlers.NewLogger(l)

	servemux := http.NewServeMux()
	servemux.Handle("/", hh)
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

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan //this is blocking operation --reading from channel
	l.Println("Received terminate graceful shutdown", sig)
	d := time.Now().Add(30 * time.Second)
	tc, _ := context.WithDeadline(context.Background(), d)
	webserver.Shutdown(tc)

}
