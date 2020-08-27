package product

import (
	"context"
	"github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/repository"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context) ([]product.Product, error)
}

type ProductUseCaseImpl struct {
	pr repository.ProductRepository
}

func NewProductUseCaseImpl(p repository.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		pr: p,
	}
}

func (s ProductUseCaseImpl) GetProducts(ctx context.Context) ([]product.Product, error) {
	return s.pr.GetProducts(&ctx)
}
