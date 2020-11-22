package handlers

import (
	"fmt"
	"lesson3/data"
	"log"
	"net/http"
)

//Products defines API of a product handler
type Products struct {
	l *log.Logger
	m data.ProductService
}

//NewProducts inits a new product handler
func NewProducts(l *log.Logger, m data.ProductService) *Products {
	return &Products{l, m}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(w, r)
	case http.MethodPost:
		p.saveProduct(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	//should this be an interface type ?
	pl, err := p.m.GetProducts()
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

	var product data.Product
	prod, err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "could not decode product", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%#v\n", *prod)

	pID, err := p.m.SaveProduct(prod)
	if err != nil {
		http.Error(w, "could not save product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "saved product with id %d\n", pID)
}
