package handlers

import (
	"fmt"
	"log"
	"net/http"
)

//Goodbye struct
type Goodbye struct {
	l *log.Logger
}

//NewGoodbye constructor
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

//Goodbye logs goodbye world
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye world!")
	fmt.Fprint(rw, "Goodbye world!\n")
}
