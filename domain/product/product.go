package product

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Stock       int    `json:"stock"`
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
