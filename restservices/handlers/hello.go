package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Hello handler struct
type Hello struct {
	l *log.Logger
}

//NewHello init new handler instance with Logger Dependency
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//implicit implementation (not specified in struct def)
//this is good because we want to access dependencies of the pointer receiver h (like logger) to process the request and response
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world!")

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello %s \n", d)
}
