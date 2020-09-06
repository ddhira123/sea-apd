package merchant

import (
	"errors"
	"github.com/golang/mock/gomock"
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

func (m MockRepository) UpdateMerchantBalance(amount int, merchantId string) error {
	if len(merchantId) == 0 || amount == 0 {
		return errors.New("Id and amount cannot be empty")
	}
	return nil
}

func (m MockRepository) GetMerchantBalance(merchantId string) (int, error) {
	if len(merchantId) == 0 {
		return 0, errors.New("Id cannot be empty")
	}
	return 100, nil
}
