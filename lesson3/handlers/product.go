package handlers

import (
	"lesson3/data"
	"log"
	"net/http"
)

//Products defines API of a product handler
type Products struct {
	l *log.Logger
}

//NewProducts inits a new product handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
	case http.MethodPost:
		p.saveProduct(w, r)
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	//should this be an interface type ?
	pl, err := data.GetProducts()
	if err != nil {
		http.Error(w, "could not retrieve products", http.StatusInternalServerError)
	}

	err = pl.ToJSON(w)
	if err != nil {
		http.Error(w, "could not encode json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) saveProduct(w http.ResponseWriter, r *http.Request) {
	//should the actual work be defined on type Product?
}
