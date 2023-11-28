package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"go.uber.org/zap"
)

type Logging struct {
	applogger *log.Logger
}

func NewLogger(l *log.Logger) *Logging {
	return &Logging{}
}

func (lg *Logging) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lg.applogger.Println("hello world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
