package product

import (
	"encoding/json"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
	"log"
	"net/http"
)

type ProductController struct {
	usecase domain.ProductUsecase
}

func NewProductController(e *echo.Echo, p domain.ProductUsecase) {
	c := &ProductController{
		usecase: p,
	}
	e.GET("/products", c.GetProducts)
	e.POST("/products", c.CreateProduct)
}

func (p *ProductController) GetProducts(c echo.Context) error {
	products, err := p.usecase.GetProducts()
	if err != nil {
		log.Panic("Error")
	}
	s, err := json.Marshal(products)
	if err != nil {
		log.Panic("Error")
	}
	return c.JSON(http.StatusOK, string(s))
}

func (p *ProductController) CreateProduct(c echo.Context) error {
	//uploadedFile, _, err := r.FormFile("file")
	//if err != nil {
	//	panic("Error processing image")
	//}
	//defer uploadedFile.Close()
	var productRequest request.Product
	c.Bind(&productRequest)
	if err := p.usecase.CreateProduct(productRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, "err")
}
