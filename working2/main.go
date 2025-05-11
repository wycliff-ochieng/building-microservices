package main

import (
	"log"
	"net/http"
	"os"

	"github.com/wycliff-ochieng/working2/handlers"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	//create handlers
	bh := handlers.NewBook(l)

	//create servermux
	sm := http.NewServeMux()
	sm.Handle("/book", bh)

	s := http.Server{
		Addr:     ":9000",
		Handler:  sm,
		ErrorLog: l,
	}
	s.ListenAndServe()
}
