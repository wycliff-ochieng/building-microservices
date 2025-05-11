package data

import (
	"encoding/json"
	"fmt"
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
		ID:          2,
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
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	ProductList[pos] = p

	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}
