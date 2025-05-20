package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/blog/data"
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

type PostKey struct{}

func (p *Post) AddPost(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post posts method....")

	pst := r.Context().Value(PostKey{})
	post, ok := pst.(data.Post)
	if !ok {
		http.Error(w, "Could not add Post", http.StatusMethodNotAllowed)
		return
	}
	data.AddPost(&post)
	json.NewEncoder(w).Encode(post)
}
