package merchant

import (
	"errors"

	"github.com/golang/mock/gomock"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

var (
	emptyUpdateMerchantBalanceRequest = merchant.UpdateMerchantBalanceRequest{}
	emptyMerchantRequest              = merchant.MerchantRequest{}
	emptyMerchant                     = domain.Merchant{}
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

func (m MockUsecase) RegisterMerchant(request merchant.MerchantRequest) error {
	if request == emptyMerchantRequest {
		return errors.New("Cannot Create Merchant")
	}
	return nil
}

func (m MockUsecase) GetMerchants() ([]domain.Merchant, error) {
	return []domain.Merchant{}, nil
}

func (m MockUsecase) GetMerchantById(merchantId string) (*domain.Merchant, error) {
	if merchantId != "" {
		return &emptyMerchant, nil
	}
	return nil, errors.New("Cannot Get Merchant By Id")
}

func (m MockUsecase) GetMerchantsByUser(userId string) ([]domain.Merchant, error) {
	if len(userId) == 0 {
		return nil, errors.New("Cannot Get Merchants by User")
	}
	return []domain.Merchant{}, nil
}

func (m MockUsecase) UpdateMerchantApprovalStatus(request merchant.UpdateMerchantApprovalStatusRequest) error {
	panic("implement me")
}
