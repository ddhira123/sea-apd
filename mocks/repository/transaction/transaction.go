package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/domain/transaction"
	"reflect"
)

var (
	emptyTransaction = transaction.Transaction{
		Base:       domain.Base{},
		Status:     "",
		BankNumber: "",
		BankName:   "",
		Amount:     0,
		CustomerId: "",
		MerchantId: "",
	}
	mockTransaction = transaction.Transaction{
		Base:       domain.Base{},
		Status:     "123",
		BankNumber: "123",
		BankName:   "123",
		Amount:     10,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockTransactionSlice   = []transaction.Transaction{}
	mockProductTransaction = transaction.ProductTransaction{}
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func (m MockRepository) CreateCart(transaction transaction.Transaction) error {
	if reflect.DeepEqual(transaction, emptyTransaction) {
		return errors.New("Transaction cannot be empty")
	}
	return nil
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) GetTransactionById(id string) (*transaction.Transaction, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be empty")
	}
	return &mockTransaction, nil
}

func (m MockRepository) UpdateTransactionStatus(status string, id string) (*transaction.Transaction, error) {
	if len(status) == 0 || len(id) == 0 {
		return nil, errors.New("Cannot Update with empty object")
	}
	return &mockTransaction, nil
}

func (m MockRepository) GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]transaction.Transaction, error) {
	if len(userId) == 0 || len(requiredStatus) == 0 {
		return nil, errors.New("Cannot Get Required status with empty user id")
	}
	return []transaction.Transaction{}, nil
}

func (m MockRepository) GetMerchantRequestItem(merchantId string) ([]transaction.Transaction, error) {
	if len(merchantId) == 0 {
		return nil, errors.New("cannot get merchant request item")
	}
	return mockTransactionSlice, nil
}

func (m MockRepository) UpdateTransaction(transaction transaction.Transaction) error {
	if reflect.DeepEqual(transaction, emptyTransaction) {
		return errors.New("Cannot update transaction")
	}
	return nil
}

func (m MockRepository) AddCartItem(cart transaction.ProductTransaction) error {
	if cart == mockProductTransaction {
		return errors.New("cannot add item to cart")
	}
	return nil
}

func (m MockRepository) RemoveCartItem(cart transaction.ProductTransaction) error {
	if cart == mockProductTransaction {
		return errors.New("cannot remove item from cart")
	}
	return nil
}

func (m MockRepository) UpdateCartItem(cart transaction.ProductTransaction) error {
	if cart == mockProductTransaction {
		return errors.New("cannot update item from cart")
	}
	return nil
}

func (m MockRepository) GetCartItems(id string) ([]transaction.ProductTransaction, error) {
	if len(id) == 0 {
		return nil, errors.New("cannot get cart items")
	}
	return []transaction.ProductTransaction{}, nil
}
