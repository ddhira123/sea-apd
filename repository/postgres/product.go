package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	db.AutoMigrate(&domain.Product{})
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetProducts() ([]domain.Product, error) {
	var Products []domain.Product
	err := p.db.Find(&Products).Error
	if err != nil {
		return nil, err
	}
	return Products, nil
}

func (p *ProductRepository) GetProductById(product domain.Product) domain.Product {
	panic("implement me")
}

func (p *ProductRepository) CreateProduct(product domain.Product) error {
	panic("implement me")
}

func (p *ProductRepository) UpdateProduct(productId string, product domain.Product) error {
	panic("implement me")
}

func (p *ProductRepository) DeleteProduct(productId domain.Product) error {
	panic("implement me")
}
