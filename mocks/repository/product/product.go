package product

import (
	"errors"
	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/product"
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
		return &domain.Product{
			Name:        "Mock Name",
			Description: "Mock Desc",
			Price:       20,
			Image:       "Mock image",
			Stock:       30,
		}, nil
	}
	return nil, errors.New("Cannot Get Product By Id")
}

func (m MockRepository) CreateProduct(product domain.Product) error {
	var p = domain.Product{}
	if product == p {
		return errors.New("Cannot Create Product")
	}
	return nil
}

func (m MockRepository) UpdateProduct(s string, product domain.Product) error {
	var p = domain.Product{}
	if s != "" || product == p {
		return nil
	}
	return errors.New("Cannot Update Product")
}

func (m MockRepository) DeleteProduct(s string) error {
	if s != "" {
		return nil
	}
	return errors.New("Cannot Delete Product")
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) GetProductsByMerchant(merchantId string) ([]domain.Product, error) {
	if merchantId != "" {
		return []domain.Product{}, nil
	}
	return nil, errors.New("Cannot Delete Product")
}

func (m MockRepository) GetProductPriceSumByTransactionId(transactionId string) int {
	if len(transactionId) == 0 {
		return 0
	}
	return 10
}
