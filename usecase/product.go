package usecase

import (
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
)

type ProductUseCaseImpl struct {
	pr domain.ProductRepository
}

func ConvertToDomain(p request.Product) domain.Product {
	return domain.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Image:       "",
		Stock:       p.Stock,
	}
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

func (s *ProductUseCaseImpl) GetProductById(productId string) (domain.Product, error) {
	panic("implement me")
}

func (s *ProductUseCaseImpl) CreateProduct(product request.Product) error {
	p := ConvertToDomain(product)
	err := s.pr.CreateProduct(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductUseCaseImpl) UpdateProduct(productId string, product request.Product) error {
	panic("implement me")
}

func (s *ProductUseCaseImpl) DeleteProduct(productId string) error {
	panic("implement me")
}
