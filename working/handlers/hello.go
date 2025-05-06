package handlers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		errors.New("ooops")
	}
	fmt.Fprintf(rw, " hello %s", d)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Helloooo")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		errors.New("Oooops")
	}
	fmt.Fprintf(w, "hello %s", d)
}
