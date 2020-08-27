package domain

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name        string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
}

type ProductUsecase interface {
	GetProducts() ([]Product, error)
	GetProductById(product Product) Product
	CreateProduct(product Product) error
	UpdateProduct(productId string, product Product) error
	DeleteProduct(productId Product) error
}

type ProductRepository interface {
	GetProducts() ([]Product, error)
	GetProductById(product Product) Product
	CreateProduct(product Product) error
	UpdateProduct(productId string, product Product) error
	DeleteProduct(productId Product) error
}

func NewProduct(name string, desc string, price int, image string, stock int) *Product {
	return &Product{
		Name:        name,
		Description: desc,
		Price:       price,
		Image:       image,
		Stock:       stock,
	}
}
