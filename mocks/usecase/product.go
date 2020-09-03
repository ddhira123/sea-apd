package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/product"
)

var emptyProduct = product.Product{}

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) GetProductById(s string) (*product.Product, error) {
	if s != "" {
		return &emptyProduct, nil
	}
	return nil, errors.New("Cannot Get Product By Id")
}

func (m MockUsecase) CreateProduct(product product.Product) error {
	if product == emptyProduct {
		return errors.New("Cannot Create Product")
	}
	return nil
}

func (m MockUsecase) UpdateProduct(productId string, product product.Product) error {
	if productId != "" && product != emptyProduct {
		return nil
	}
	return errors.New("Cannot Update Product")
}

func (m MockUsecase) DeleteProduct(productId string) error {
	if productId != "" {
		return nil
	}
	return errors.New("Cannot Delete Product")
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) GetProducts() ([]product.Product, error) {
	m.ctrl.T.Helper()
	return []product.Product{}, nil
}
