package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/wycliff-ochieng/blog/data"
	"github.com/wycliff-ochieng/blog/middleware"
)

type Post struct {
	l *log.Logger
}

func NewPost(l *log.Logger) *Post {
	return &Post{l}
}

func (p *Post) GetPost(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle GET posts method....")

	lp := data.GetPost()

	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Data could not be marshalled", http.StatusBadRequest)
	}

	w.Write(d)
}

//type PostKey struct{}

func (p *Post) AddPost(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post posts method....")

	pst := r.Context().Value(middleware.PostKey{})
	post, ok := pst.(data.Post)
	if !ok {
		http.Error(w, "Could not add Post", http.StatusMethodNotAllowed)
		return
	}
	data.AddPost(&post)
	json.NewEncoder(w).Encode(post)
}

func (p *Post) GetPostByID(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle GET posts by ID....")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Cannot convert ID to integer", http.StatusInternalServerError)
		return
	}

	for _, p := range data.PostList {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")

			if err := json.NewEncoder(w).Encode(p); err != nil {
				http.Error(w, "Cant Encode response", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}
