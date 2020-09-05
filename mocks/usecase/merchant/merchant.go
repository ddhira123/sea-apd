package merchant

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
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
	panic("implement me")
}
