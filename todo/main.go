package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/todo/handler"
)

func main() {
	fmt.Println("hello world")
	l := log.New(os.Stdout, "TODO-API", log.LstdFlags)

	th := handler.NewTask(l)

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", th.GetTasks)

	l.Println("starting server right now....")
	http.ListenAndServe(":8080", router)
}
