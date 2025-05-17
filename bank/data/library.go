package data

import (
	"math/rand"
)

type Book struct {
	ID     int
	Name   string
	Author string
	Review string
	views  int
}

var BookList = []*Book{
	{
		ID:     rand.Intn(10),
		Name:   "Room",
		Author: "Emmma Donoughue",
		Review: "A story about Kidnappping",
		views:  rand.Intn(500),
	},
	{
		ID:     rand.Intn(10),
		Name:   "Room",
		Author: "Emmma Donoughue",
		Review: "A story about Kidnappping",
		views:  rand.Intn(500),
	},
}

type Books []*Book

func (b Book) GetBooks() []*Book {
	return BookList
}
func (b *Book) AddBooks() {
	b.ID = getNextBookID()
	BookList = append(BookList, b)
}

func getNextBookID() int {
	lastBook := BookList[len(BookList)-1]
	return lastBook.ID + 1
}
