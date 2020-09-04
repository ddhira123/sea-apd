package transaction

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/transaction"
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
	panic("implement me")
}

func (m MockRepository) GetTransactionById(s string) (*transaction.Transaction, error) {
	panic("implement me")
}

func (m MockRepository) UpdateTransactionStatus(s string, s2 string) error {
	panic("implement me")
}
