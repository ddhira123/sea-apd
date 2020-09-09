package product

import (
	"github.com/williamchang80/sea-apd/domain/product"
	request "github.com/williamchang80/sea-apd/dto/request/product"
)

type ProductUsecase struct {
	pr product.ProductRepository
}

func ConvertToDomain(p request.ProductRequest) product.Product {
	return product.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Image:       "",
		Stock:       p.Stock,
		MerchantId: p.MerchantId,
	}
}
func NewProductUseCase(p product.ProductRepository) product.ProductUsecase {
	return &ProductUsecase{
		pr: p,
	}
}
func (s *ProductUsecase) GetProducts() ([]product.Product, error) {
	p, err := s.pr.GetProducts()
	if err != nil {
		return nil, err
	}
	return p, nil
}
func (s *ProductUsecase) GetProductById(productId string) (*product.Product, error) {
	p, err := s.pr.GetProductById(productId)
	if err != nil || p == nil {
		return nil, err
	}
	return p, nil
}
func (s *ProductUsecase) CreateProduct(product request.ProductRequest) error {
	p := ConvertToDomain(product)
	err := s.pr.CreateProduct(p)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductUsecase) UpdateProduct(productId string, request request.ProductRequest) error {
	p := ConvertToDomain(request)
	err := s.pr.UpdateProduct(productId, p)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductUsecase) DeleteProduct(productId string) error {
	err := s.pr.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductUsecase) GetProductsByMerchant(merchantId string) ([]product.Product, error) {
	products, err := s.pr.GetProductsByMerchant(merchantId)
	if err != nil {
		return nil, err
	}
	return products, nil
}
