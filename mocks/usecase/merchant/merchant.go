package merchant

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

var (
	emptyUpdateMerchantBalanceRequest = merchant.UpdateMerchantBalanceRequest{}
)

type MockUsecase struct {
	ctrl *gomock.Controller
}

func NewMockUsecase(repo *gomock.Controller) *MockUsecase {
	return &MockUsecase{
		ctrl: repo,
	}
}

func (m MockUsecase) UpdateMerchantBalance(request merchant.UpdateMerchantBalanceRequest) error {
	if request == emptyUpdateMerchantBalanceRequest {
		return errors.New("Request cannot be empty")
	}
	return nil
}

func (m MockUsecase) GetMerchantBalance(merchantId string) (int, error) {
	if len(merchantId) == 0 {
		return 0, errors.New("Merchant id cannot be empty")
	}
	return 1000, nil
}
