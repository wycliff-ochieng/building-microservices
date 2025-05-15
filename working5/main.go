package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/working3/handlers"
)

func main() {
	fmt.Println("Gorilla handler")

	//create logger
	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	//create a new gorilla mux
	r := mux.NewRouter()

	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/product", ph.GetProduct)
	getRouter.HandleFunc("/product/{id:[0-100]+}", ph.GetProductByID)

	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/product", ph.AddProduct)
	postRouter.Use(handlers.MiddlewareValidateProduct)

	//create server

	s := http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	s.ListenAndServe()

}
