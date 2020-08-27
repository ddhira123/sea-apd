package usecase

import (
	"github.com/williamchang80/sea-apd/domain"
)

type ProductUseCaseImpl struct {
	pr domain.ProductRepository
}

func NewProductUseCaseImpl(p domain.ProductRepository) domain.ProductUsecase {
	return &ProductUseCaseImpl{
		pr: p,
	}
}

func (s *ProductUseCaseImpl) GetProducts() ([]domain.Product, error) {
	p, err := s.pr.GetProducts()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductUseCaseImpl) GetProductById(product domain.Product) domain.Product {
	panic("implement me")
}

func (s *ProductUseCaseImpl) CreateProduct(product domain.Product) error {
	panic("implement me")
}

func (s *ProductUseCaseImpl) UpdateProduct(productId string, product domain.Product) error {
	panic("implement me")
}

func (s *ProductUseCaseImpl) DeleteProduct(productId domain.Product) error {
	panic("implement me")
}
