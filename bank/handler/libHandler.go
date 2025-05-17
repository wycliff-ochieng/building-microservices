package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/working2/data"
)

type Book struct {
	l *log.Logger
}

func NewBook(l *log.Logger) *Book {
	return &Book{l}
}

func (b *Book) GetBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("handling GET method requests")

	bl := data.GetBooks()

	d, err := json.Marshal(bl)
	if err != nil {
		http.Error(w, "Unable to list books", http.StatusInternalServerError)
	}
	w.Write(d)
}

func (b *Book) AddBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("handling POST method requests")

	bk := r.Context().Value(bookkey{})
	book, ok := bk.(data.Book)
	if !ok {
		http.Error(w, "Could validate task", http.StatusBadRequest)
	}
	data.AddBook(&book)
}
