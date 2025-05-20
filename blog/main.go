package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/blog/handler"
	"github.com/wycliff-ochieng/blog/middleware"
)

func main() {
	fmt.Println("Blogging")

	l := log.New(os.Stdout, "BLOG PLATFORM API", log.LstdFlags)

	ph := handler.NewPost(l)

	router := mux.NewRouter()

	getRouter := router.Methods("GET").Subrouter()
	getRouter.HandleFunc("/posts", ph.GetPost)

	postRouter := router.Methods("POST").Subrouter()
	postRouter.HandleFunc("/posts", ph.AddPost)
	postRouter.Use(middleware.ValidatePostMiddleware)

	http.ListenAndServe(":8000", router)

}
