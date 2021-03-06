package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

var (
	emptyCreateCartRequest              = transaction.CreateCartRequest{}
	emptyUpdateTransactionStatusRequest = transaction.UpdateTransactionRequest{}
	mockTransactionSlice                = []domain.Transaction{}
	mockPaymentRequest                  = transaction.PaymentRequest{}
	mockCartRequest                     = transaction.CartRequest{}
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) CreateCart(request transaction.CreateCartRequest) error {
	if request == emptyCreateCartRequest {
		return errors.New("Request cannot be empty")
	}
	return nil
}

func (m MockUsecase) UpdateTransactionStatus(request transaction.UpdateTransactionRequest) error {
	if request == emptyUpdateTransactionStatusRequest {
		return errors.New("Request cannot be empty")
	}
	return nil
}

func (m MockUsecase) GetTransactionById(id string) (*domain.Transaction, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be empty")
	}
	return &domain.Transaction{}, nil
}

func (m MockUsecase) GetTransactionHistory(userId string) ([]domain.Transaction, error) {
	if len(userId) != 0 {
		return mockTransactionSlice, nil
	}
	return nil, errors.New("User Id cannot be empty")
}

func (m MockUsecase) GetMerchantRequestItem(merchantId string) ([]domain.Transaction, error) {
	if len(merchantId) == 0 {
		return nil, errors.New("cannot get merchant request item")
	}
	return mockTransactionSlice, nil
}

func (m MockUsecase) PayTransaction(request transaction.PaymentRequest) error {
	if request == mockPaymentRequest {
		return errors.New("cannot pay transaction")
	}
	return nil
}

func (m MockUsecase) AddCartItem(request transaction.CartRequest) error {
	if request == mockCartRequest {
		return errors.New("cannot add item to cart")
	}
	return nil
}

func (m MockUsecase) RemoveCartItem(request transaction.CartRequest) error {
	if request == mockCartRequest {
		return errors.New("cannot remove item from cart")
	}
	return nil
}

func (m MockUsecase) UpdateCartItem(request transaction.CartRequest) error {
	if request == mockCartRequest {
		return errors.New("cannot update item from cart")
	}
	return nil
}

func (m MockUsecase) GetCartItems(id string) ([]domain.ProductTransaction, error) {
	if len(id) == 0 {
		return nil, errors.New("cannot get items from cart")
	}
	return []domain.ProductTransaction{}, nil
}
