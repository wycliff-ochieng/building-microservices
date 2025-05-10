package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/working2/data"
)

type book struct {
	l *log.Logger
}

func NewBook(l *log.Logger) *book {
	return &book{l}
}

func (b *book) getBook(w http.ResponseWriter, r *http.Request) {
	lb := data.GetBooks()

	d, err := json.Marshal(lb)
	if err != nil {
		http.Error(w, "Unable to retrieve books", http.StatusMethodNotAllowed)
	}
	w.Write(d)

}

func (b *book) addBook(w http.ResponseWriter, r *http.Request) {
	b := AddBook()
	if err != nil {
		http.Error(w, "Unable to add Book to list", http.StatusMethodNotAllowed)
	}
}
