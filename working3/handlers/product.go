package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/wycliff-ochieng/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusBadRequest)
	}
	w.Write(d)
}

func (p *Products) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(vars[len(vars)-1])
	if err != nil {
		http.Error(w, "Invalid URI", http.StatusMethodNotAllowed)
	}
	for _, p := range data.ProductList {
		if p.ID == id {
			w.Header().Set("Context-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	return
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "cannot Unamrshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) UpdateProduct(id int, w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}

}
