package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/bank/handler"
)

func main() {

	l := log.New(os.Stdout, "LIBRARY MANAGEMENT SERVICE", log.LstdFlags)

	bh := handler.NewBook(l)

	router := mux.NewRouter()
	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/books", bh.GetBook)

	postRouter := router.Methods("GET").Subrouter()
	postRouter.HandleFunc("/add", bh.AddBook)
	postRouter.Use(handler.BookValidateMiddleware)

	http.ListenAndServe(":9000", router)
}
