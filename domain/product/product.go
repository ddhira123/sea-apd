package product

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/product"
)

type Product struct {
	domain.Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
	MerchantId  string `json:"merchant_id"`
}

type ProductUsecase interface {
	GetProducts() ([]Product, error)
	GetProductById(productId string) (*Product, error)
	CreateProduct(productRequest product.ProductRequest) error
	UpdateProduct(productId string, productRequest product.ProductRequest) error
	DeleteProduct(productId string) error
	GetProductsByMerchant(merchantId string) ([]Product, error)
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductById(productId string) (*Product, error)
	CreateProduct(Product) error
	UpdateProduct(productId string, product Product) error
	DeleteProduct(productId string) error
	GetProductsByMerchant(merchantId string) ([]Product, error)
}

type ProductController interface {
	GetProducts(echo.Context) error
	GetProductById(echo.Context) error
	CreateProduct(echo.Context) error
	UpdateProduct(echo.Context) error
	DeleteProduct(echo.Context) error
	GetProductsByMerchant(echo.Context) error
}
