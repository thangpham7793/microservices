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
	m *data.ProductService
}

//NewProducts inits a new product handler
func NewProducts(l *log.Logger, m *data.ProductService) *Products {
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

//HandleError handles errors
func HandleError(err error, w http.ResponseWriter, desc APIError) {
	if err != nil {
		switch desc {
		case ErrorRetrieveProducts, ErrorSaveProduct, ErrorEncodeJSON:
			http.Error(w, string(desc), http.StatusInternalServerError)
		case ErrorDecodeJSON:
			http.Error(w, string(desc), http.StatusBadRequest)
		case ErrorPageNotFound:
			http.Error(w, string(desc), http.StatusNotFound)
		}
		return
	}
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	//should this be an interface type ?
	//s := *(p.m)
	pl, err := (*p.m).GetProducts()
	HandleError(err, w, ErrorRetrieveProducts)

	err = pl.ToJSON(w)
	HandleError(err, w, ErrorEncodeJSON)
}

func (p *Products) saveProduct(w http.ResponseWriter, r *http.Request) {
	//should the actual work be defined on type Product?

	var product data.Product
	prod, err := product.FromJSON(r.Body)
	HandleError(err, w, ErrorDecodeJSON)

	fmt.Printf("%#v\n", *prod)

	pID, err := (*p.m).SaveProduct(prod)
	HandleError(err, w, ErrorSaveProduct)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "saved product with id %d\n", pID)
}
