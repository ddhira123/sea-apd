package repository

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func (m MockRepository) GetProducts() ([]domain.Product, error) {
	m.ctrl.T.Helper()
	return []domain.Product{}, nil
}

func (m MockRepository) GetProductById(id string) (*domain.Product, error) {
	if id != "" {
		return &domain.Product{}, nil
	}
	return nil, errors.New("Mock error")
}

func (m MockRepository) CreateProduct(product domain.Product) error {
	var p = domain.Product{}
	if product == p {
		return errors.New("Mock error")
	}
	return nil
}

func (m MockRepository) UpdateProduct(s string, product domain.Product) error {
	var p = domain.Product{}
	if s != "" || product == p {
		return nil
	}
	return errors.New("Mock Error")
}

func (m MockRepository) DeleteProduct(s string) error {
	if s != "" {
		return nil
	}
	return errors.New("Mock error")
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}
