package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qascade/qascadeservice/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
	} else if r.Method == http.MethodPut {

	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		err = fmt.Errorf("not able to convert product list to json, %s", err)
		p.l.Fatal(err)
	}
}
