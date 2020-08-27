package product

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamchang80/sea-apd/domain"
	"log"
	"net/http"
)

type ProductController struct {
	usecase domain.ProductUsecase
}

func NewProductController(r *mux.Router, p domain.ProductUsecase) {
	c := &ProductController{
		usecase: p,
	}
	r.HandleFunc("/products", c.GetProducts)
}

func (p *ProductController) GetProducts(r http.ResponseWriter, w *http.Request) {
	products, err := p.usecase.GetProducts()
	if err != nil {
		log.Panic("Error")
	}
	s, err := json.Marshal(products)
	if err != nil {
		log.Panic("Error")
	}
	fmt.Fprintf(r, string(s))
}
