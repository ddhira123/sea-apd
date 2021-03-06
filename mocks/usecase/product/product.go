package product

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/product"
	"github.com/williamchang80/sea-apd/domain/transaction"
	product2 "github.com/williamchang80/sea-apd/dto/request/product"
	"reflect"
)

var (
	emptyProduct        = product.Product{}
	emptyProductRequest = product2.ProductRequest{}
	emptyProductSlice   = []product.Product{}
	emptyTransaction    = transaction.Transaction{}
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

func (m MockUsecase) GetProducts() ([]product.Product, error) {
	return []product.Product{}, nil
}

func (m MockUsecase) GetProductById(id string) (*product.Product, error) {
	if id != "" {
		return &emptyProduct, nil
	}
	return nil, errors.New("Cannot Get Product By Id")
}

func (m MockUsecase) CreateProduct(request product2.ProductRequest) error {
	if request == emptyProductRequest {
		return errors.New("Cannot Create Product")
	}
	return nil
}

func (m MockUsecase) UpdateProduct(id string, request product2.ProductRequest) error {
	if request == emptyProductRequest {
		return nil
	}
	return errors.New("Cannot Update Product")
}

func (m MockUsecase) DeleteProduct(id string) error {
	if len(id) != 0 {
		return nil
	}
	return errors.New("Cannot Delete Product")
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}
func (m MockUsecase) GetProductsByMerchant(merchantId string) ([]product.Product, error) {
	if len(merchantId) == 0 {
		return nil, errors.New("Cannot Get Products by Merchant")
	}
	return emptyProductSlice, nil
}

func (m MockUsecase) GetProductPriceTotal(transaction transaction.Transaction) (int, error) {
	if reflect.DeepEqual(transaction, emptyTransaction) {
		return 0, errors.New("Cannot get product price total")
	}
	return 1000, nil
}
