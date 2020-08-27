package product

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamchang80/sea-apd/usecase/product"
	"log"
	"net/http"
)

type ProductController struct {
	pc product.ProductUseCase
}

func NewProductController(r *mux.Router, p product.ProductUseCase) {
	c := &ProductController{
		pc: p,
	}
	r.HandleFunc("/products", c.GetProducts)
}

func (p *ProductController) GetProducts(r http.ResponseWriter, w *http.Request) {
	c := w.Context()
	products, err := p.pc.GetProducts(c)
	if err != nil {
		log.Panic("Error")
	}
	s, err := json.Marshal(products)
	if err != nil {
		log.Panic("Error")
	}
	fmt.Fprintf(r, string(s))
}
