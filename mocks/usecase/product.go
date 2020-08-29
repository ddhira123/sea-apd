package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) GetProductById(s string) (domain.Product, error) {
	panic("implement me")
}

func (m MockUsecase) CreateProduct(product domain.Product) error {
	panic("implement me")
}

func (m MockUsecase) UpdateProduct(s string, product domain.Product) error {
	panic("implement me")
}

func (m MockUsecase) DeleteProduct(s string) error {
	panic("implement me")
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{
		ctrl: repo,
	}
	return mock
}

func (m MockUsecase) GetProducts() ([]domain.Product, error) {
	m.ctrl.T.Helper()
	return []domain.Product{}, nil
}
