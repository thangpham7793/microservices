package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product represents a drink product API
type Product struct {
	ID          int     `json:"id"` //no spaces!
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

//tags make working with MongoDb so nice (no need to manually skip the fields) + reduces the payload sent over the wire as well

//ProductModel defines methods of a valid prod model
type ProductModel interface {
	GetProducts() ([]*Product, error)
}

//Products define a slice of pointers to product
type Products []*Product

//ToJSON Methods on types helps with encapsulation
func (ps *Products) ToJSON(w io.Writer) error {
	//this allocates memory to hold the json before writing it, so using encoder skips this step
	//j, err := json.Marshal(pl)
	e := json.NewEncoder(w) //creating a new encoder is nowhere as expensive
	return e.Encode(ps)
}

//GetProducts abstracts this from handlers and presents itself as a service
func GetProducts() (Products, error) {
	return ProductList, nil
}

//ProductList dummy data
var ProductList Products = []*Product{
	{1, "Latte", "Frothy milk coffee", 2.45, "abc123", time.Now().UTC().String(), time.Now().UTC().String(), ""},
	{2, "Expresso", "Strong coffee", 4.45, "bcd123", time.Now().UTC().String(), time.Now().UTC().String(), ""},
}
