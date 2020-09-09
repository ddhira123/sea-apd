package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/request/transaction"
)

var emptyTransactionRequest = transaction.TransactionRequest{}
var emptyUpdateTransactionStatusRequest = transaction.UpdateTransactionRequest{}

type MockUsecase struct {
	ctrl *gomock.Controller
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) CreateTransaction(request transaction.TransactionRequest) error {
	if request == emptyTransactionRequest {
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
		return []domain.Transaction{}, nil
	}
	return nil, errors.New("User Id cannot be empty")
}
