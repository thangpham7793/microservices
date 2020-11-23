package handlers

import (
	"log"
	"net/http"
	"regexp"
	"restservices/data"
	"strconv"
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
		p.addProduct(w, r)
	case http.MethodPut:
		rgx := regexp.MustCompile(`/([0-9]+)`)
		g := rgx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(w, "invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(w, "invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		p.l.Println("Received id", id)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
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

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post!")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "could not decode json", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}
