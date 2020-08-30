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

func NewProductController(e *echo.Echo, p domain.ProductUsecase) domain.ProductController {
	c := &ProductController{
		usecase: p,
	}
	e.GET("/products", c.GetProducts)
	e.POST("/products", c.CreateProduct)
	e.GET("/product", c.GetProductById)
	e.DELETE("/product", c.DeleteProduct)
	return c
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
	var productRequest request.Product
	c.Bind(&productRequest)
	if err := p.usecase.CreateProduct(productRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, "err")
}

func (p *ProductController) GetProductById(context echo.Context) error {
	id := context.QueryParam("productId")
	res, err := p.usecase.GetProductById(id)
	if err != nil {
		return context.JSON(http.StatusNotFound, "Not Found")
	}
	s, err := json.Marshal(res)
	if err != nil {
		log.Panic("Error")
	}
	return context.JSON(http.StatusOK, string(s))
}

func (p *ProductController) UpdateProduct(context echo.Context) error {
	context.FormValue("productRequest")
	//TODO implement later
	return nil
}

func (p *ProductController) DeleteProduct(context echo.Context) error {
	id := context.QueryParam("productId")
	err := p.usecase.DeleteProduct(id)
	if err != nil {
		return context.JSON(http.StatusNotFound, "Not Found")
	}
	return context.JSON(http.StatusOK, "Success")
}
