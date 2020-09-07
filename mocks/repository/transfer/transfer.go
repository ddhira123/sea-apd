package transfer

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/transfer"
)

type MockRepository struct {
	ctrl *gomock.Controller
}
var (
	emptyCreateTransferDomain= transfer.Transfer{}
)

func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{
		ctrl: ctrl,
	}
	return mock
}

func (m MockRepository) GetTransferHistory(merchantId string) ([]transfer.Transfer, error) {
	if len(merchantId) == 0 {
		return nil, errors.New("Merchant id cannot be empty")
	}
	return []transfer.Transfer{}, nil
}

func (m MockRepository) CreateTransferHistory(transfer transfer.Transfer) error {
	if transfer == emptyCreateTransferDomain {
		return errors.New("Transfer request cannot be empty")
	}
	return nil
}
