package product

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/product"
)

type Product struct {
	domain.Base
	Name        string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}

type ProductUsecase interface {
	GetProducts() ([]Product, error)
	GetProductById(string) (*Product, error)
	CreateProduct(product.ProductRequest) error
	UpdateProduct(string, product.ProductRequest) error
	DeleteProduct(string) error
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductById(string) (*Product, error)
	CreateProduct(Product) error
	UpdateProduct(string, Product) error
	DeleteProduct(string) error
}

type ProductController interface {
	GetProducts(echo.Context) error
	GetProductById(echo.Context) error
	CreateProduct(echo.Context) error
	UpdateProduct(echo.Context) error
	DeleteProduct(echo.Context) error
}
