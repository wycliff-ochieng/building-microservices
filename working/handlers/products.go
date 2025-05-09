package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"github.com/wycliff-ochieng/working/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	if r.Method == http.Method{
		reg:=regexp.MustCompile(`/([0-9]+)`)
		g:=reg.FindAllStringSubmatch(r.URI.Path-1)
		if len(g)!=1{
			http.Error(w,"invalid URI", http.StatusBadRequest)
		}
		if len(g[0]) != 2{
			p.l.Println("invalid uri more than cpature group")
			http.Error(w,"invalid uri",http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id,err := strconv.Atoi(idString)
		if err != nil{
			p.l.Println("invalid URI cannot convert to number")
			http.Error(w,"invlaid uri", http.StatusMethodNotAllowed)
			return
		}
		p.l.Println("got id", id)
		p.updateProducts(id,w,r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
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
	p.l.Println("Handle POST product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "cannot Unamrshal json", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *Product) updateProduct(id,w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle UPDATE product")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "cannot Unamrshal json", http.StatusBadRequest)
	}
	data.UpdateProduct(prod)

}
