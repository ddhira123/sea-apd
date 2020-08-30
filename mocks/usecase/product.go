package usecase

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
)

var emptyProduct = domain.Product{}

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) GetProductById(s string) (*domain.Product, error) {
	panic("implement me")
}

func (m MockUsecase) CreateProduct(product domain.Product) error {
	if product == emptyProduct {
		return errors.New("Mock Error")
	}
	return nil
}

func (m MockUsecase) UpdateProduct(s string, product domain.Product) error {
	panic("implement me")
}

func (m MockUsecase) DeleteProduct(id string) error {
	if id != "" {
		return nil
	}
	return errors.New("Mock Error")
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) GetProducts() ([]domain.Product, error) {
	m.ctrl.T.Helper()
	return []domain.Product{}, nil
}
