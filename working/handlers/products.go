package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wycliff-ochieng/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	/*lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Unable to chnage to json", http.StatusBadRequest)
	}*/
	if r.Method == http.MethodGet {
		lp := data.GetProducts()
		d, err := json.Marshal(lp)
		if err != nil {
			http.Error(w, "cannot marshal", http.StatusBadRequest)
		}
		w.Write(d)
	}
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return

	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	//w.Write(d)
}

/*
	func (p *Products) getProduct(w http.ResponseWriter, r *http.Request) {
		lp := data.GetProducts()
		err := data.Products(lp).ToJSON(w)
		if err != nil {
			http.Error(w, "Unable to marshal", http.StatusBadRequest)
		}
	}
*/
func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "cannot Unamrshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}
