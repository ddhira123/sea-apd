package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	if db != nil {
		db.AutoMigrate(&domain.Product{})
	}
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) GetProductById(productId string) (*domain.Product, error) {
	var product []domain.Product
	err := p.db.Find(&product, productId).Error
	if err != nil || len(product) == 0 {
		return nil, err
	}
	return &product[0], nil
}

func (p *ProductRepository) CreateProduct(product domain.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) UpdateProduct(productId string, product domain.Product) error {
	panic("implement me")
}

func (p *ProductRepository) DeleteProduct(productId string) error {
	if err := p.db.Delete(&domain.Product{}, productId).Error; err != nil {
		return err
	}
	return nil
}
