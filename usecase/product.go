package usecase

import (
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
)

type ProductUsecase struct {
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

func NewProductUseCase(p domain.ProductRepository) domain.ProductUsecase {
	return &ProductUsecase{
		pr: p,
	}
}

func (s *ProductUsecase) GetProducts() ([]domain.Product, error) {
	p, err := s.pr.GetProducts()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductUsecase) GetProductById(productId string) (domain.Product, error) {
	panic("implement me")
}

func (s *ProductUsecase) CreateProduct(product request.Product) error {
	p := ConvertToDomain(product)
	err := s.pr.CreateProduct(p)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductUsecase) UpdateProduct(productId string, product request.Product) error {
	panic("implement me")
}

func (s *ProductUsecase) DeleteProduct(productId string) error {
	panic("implement me")
}
