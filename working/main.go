package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wycliff-ochieng/working/handlers"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeeeeee"))
}

func goodbyeHandler(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		errors.New("Try again")
	}
	fmt.Fprintf(w, "Melisa %s", d)
}

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := NewGoodBye(l)
	sm := http.NewServeMux()
	sm.Handle("/bye", gh)
	http.HandleFunc("/what", goodbyeHandler)
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         "8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()
}

//go func(){
//	l.Println("Startin server at..")
//}
