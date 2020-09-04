package transaction

import (
	"errors"
	"github.com/golang/mock/gomock"
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