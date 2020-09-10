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
		UserId:     "",
	}
)

type MockRepository struct {
	ctrl *gomock.Controller
}

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) CreateTransaction(transaction transaction.Transaction) error {
	if reflect.DeepEqual(transaction, emptyTransaction) {
		return errors.New("Transaction cannot be empty")
	}
	return nil
}

func (m MockRepository) GetTransactionById(id string) (*transaction.Transaction, error) {
	if len(id) == 0 {
		return nil, errors.New("Id cannot be empty")
	}
	return &emptyTransaction, nil
}

func (m MockRepository) UpdateTransactionStatus(status string, id string) (*transaction.Transaction, error) {
	if len(status) == 0 || len(id) == 0 {
		return nil, errors.New("Cannot Update with empty object")
	}
	return &emptyTransaction, nil
}

func (m MockRepository) GetTransactionByRequiredStatus(requiredStatus []string, userId string) ([]transaction.Transaction, error) {
	if len(userId) == 0 || len(requiredStatus) == 0 {
		return nil, errors.New("Cannot Get Required status with empty user id")
	}
	return []transaction.Transaction{}, nil
}
