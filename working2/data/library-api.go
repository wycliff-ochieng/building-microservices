package data

import (
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

type Book struct {
	ID          int
	Name        string
	Author      string
	Published   string
	Description string
	Reads       int
}

var BookList []*Book = []*Book{
	{
		ID:          rand.Intn(100),
		Name:        "If it Bleeds",
		Author:      "Stepehen King",
		Published:   time.Now().UTC().String(),
		Description: "Series of Kings fanstsy",
		Reads:       rand.Intn(10000),
	},
	{
		ID:          rand.Intn(100),
		Name:        "Room",
		Author:      "Emma Donougue",
		Published:   time.Now().UTC().String(),
		Description: "Kidnapping at its worst",
		Reads:       rand.Intn(10000),
	},
}

type Books []*Book

func GetBooks() []*Book {
	return BookList
}

func AddBook(b *Book) []*Book {
	b.ID = getNextID()
	BookList = append(BookList, b)
	return BookList
}

func getNextID() int {
	lb := BookList[len(BookList)-1]
	return lb.ID + 1
}

func (b *Book) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(b)
}
