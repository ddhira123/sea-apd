package repository

import (
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

func (m MockRepository) GetProductById(s string) (domain.Product, error) {
	panic("implement me")
}

func (m MockRepository) CreateProduct(product domain.Product) error {
	panic("implement me")
}

func (m MockRepository) UpdateProduct(s string, product domain.Product) error {
	panic("implement me")
}

func (m MockRepository) DeleteProduct(s string) error {
	panic("implement me")
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}


