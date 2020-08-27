package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/product"
)

type ProductRepository struct {
	db *gorm.DB
}

var Products []product.Product

func NewProductRepository(db *gorm.DB) ProductRepository {
	db.AutoMigrate(&product.Product{})
	return ProductRepository{db: db}
}

func (p *ProductRepository) GetProducts(ctx *context.Context) ([]product.Product, error) {
	p.db.Find(&Products)
	return Products, nil
}
