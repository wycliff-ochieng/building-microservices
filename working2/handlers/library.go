package handlers

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

func (b *Book) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		b.getBook(w, r)
		return
	}
	if r.Method == http.MethodPost {
		b.addBook(w, r)
		return
	}
}

func (b *Book) getBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handling GET requests")
	lb := data.GetBooks()

	d, err := json.Marshal(lb)
	if err != nil {
		http.Error(w, "Unable to retrieve books", http.StatusMethodNotAllowed)
	}
	w.Write(d)

}

func (b Book) addBook(w http.ResponseWriter, r *http.Request) {
	b.l.Println("Handling POST requests")
	bk := &data.Book{}
	err := bk.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "invalid URI", http.StatusMethodNotAllowed)
	}
	data.AddBook(bk)
}

func (b *Book) updateBook(w http.ResponseWriter, r *http.Request) {
	return
}

func (b *Book) deleteBook(w http.ResponseWriter, r *http.Request) {
	return
}
