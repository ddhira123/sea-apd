package product

import (
	"context"
	"github.com/williamchang80/sea-apd/domain/product"
)

type ProductUseCase interface {
	GetProducts(ctx context.Context) (*product.Product, error)
}

type ProductUseCaseImpl struct {
}

func NewProductUseCaseImpl() ProductUseCase {
	return &ProductUseCaseImpl{}
}

func (s ProductUseCaseImpl) GetProducts(ctx context.Context) (*product.Product, error) {
	return product.NewProduct("Test", "desc", 23, "iamge", 23), nil
}
