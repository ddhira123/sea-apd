package postgres

import (
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain/product"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepository {
	if db != nil {
		db.AutoMigrate(&product.Product{})
	}
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetProducts() ([]product.Product, error) {
	var products []product.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *ProductRepository) GetProductById(productId string) (*product.Product, error) {
	var product product.Product
	err := p.db.Find(&product, productId).Limit(1).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductRepository) CreateProduct(product product.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) UpdateProduct(productId string, product product.Product) error {
	if err := p.db.Model(&product).Where("id = " + productId).Update(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) DeleteProduct(productId string) error {
	if err := p.db.Delete(&product.Product{}, productId).Error; err != nil {
		return err
	}
	return nil
}
