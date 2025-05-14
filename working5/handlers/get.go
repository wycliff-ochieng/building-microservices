package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/wycliff-ochieng/working5/data"
)

// swagger: route GET /products products listProducts
// Returns a list of products
// responses:
//       200: productResponse

// list All handles GET requests and returns all products
func (p *Products) GetProduct(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(w, "Unable to marshal", http.StatusBadRequest)
	}
	w.Write(d)
}
