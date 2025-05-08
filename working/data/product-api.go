package data

import (
	"encoding/json"
	"io"
	"math/rand"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatedON   string  `json:"-"`
	UpdatedON   string  `json:"-"`
	DeletedON   string  `json:"-"`
}

var ProductList = []*Product{
	{
		ID:          rand.Intn(100),
		Name:        "Coffee Espresso",
		Description: "Mango Flavour",
		Price:       rand.Float64(),
		CreatedON:   time.Now().UTC().String(),
		UpdatedON:   time.Now().UTC().String(),
		DeletedON:   time.Now().UTC().String(),
	},
	{
		ID:          rand.Intn(100),
		Name:        "Coffee Capucino",
		Description: "yellow Flavour",
		Price:       rand.Float64(),
		CreatedON:   time.Now().UTC().String(),
		UpdatedON:   time.Now().UTC().String(),
		DeletedON:   time.Now().UTC().String(),
	},
}

type Products []*Product

func (p *Product) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() []*Product {
	return ProductList
}

func AddProduct(p *Product) {
	return
}
