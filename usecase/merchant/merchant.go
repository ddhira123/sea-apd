package merchant

import (
	"github.com/williamchang80/sea-apd/common/constants/merchant_status"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/domain/merchant"
	user "github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
)

type MerchantUsecase struct {
	mc      merchant.MerchantRepository
	usecase user.UserUsecase
}

func NewMerchantUsecase(m merchant.MerchantRepository, usecase user.
UserUsecase) merchant.MerchantUsecase {
	mc := MerchantUsecase{mc: m, usecase: usecase}
	return mc
}

func ConvertMerchantRequestToEntity(m request.MerchantRequest) merchant.Merchant {
	return merchant.Merchant{
		Name:     m.Name,
		Balance:  m.Balance,
		UserId:   m.UserId,
		Brand:    m.Brand,
		Address:  m.Address,
		Approval: merchant_status.ToString(merchant_status.WAITING),
	}
}

func (m MerchantUsecase) UpdateMerchantBalance(request request.UpdateMerchantBalanceRequest) error {
	if err := m.mc.UpdateMerchantBalance(request.Amount, request.MerchantId); err != nil {
		return err
	}
	return nil
}

func (m MerchantUsecase) GetMerchantBalance(merchantId string) (int, error) {
	balance, err := m.mc.GetMerchantBalance(merchantId)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (m MerchantUsecase) RegisterMerchant(request request.MerchantRequest) error {
	merch := ConvertMerchantRequestToEntity(request)
	err := m.mc.RegisterMerchant(merch)
	return err
}

func (m MerchantUsecase) GetMerchants() ([]merchant.Merchant, error) {
	mh, err := m.mc.GetMerchants()
	if err != nil {
		return nil, err
	}
	return mh, nil
}

func (m MerchantUsecase) GetMerchantById(merchantId string) (*merchant.Merchant, error) {
	mh, err := m.mc.GetMerchantById(merchantId)
	if err != nil || mh == nil {
		return nil, err
	}
	return mh, nil
}

func (m MerchantUsecase) UpdateMerchantApprovalStatus(request request.
UpdateMerchantApprovalStatusRequest) error {
	if err := m.mc.UpdateMerchantApprovalStatus(request.MerchantId,
		merchant_status.ToString(request.Status)); err != nil {
		return err
	}
	if request.Status == merchant_status.ACCEPTED {
		merch, _ := m.mc.GetMerchantById(request.MerchantId)
		updateRequest := user2.UpdateUserRoleRequest{Role: user_role.MERCHANT,
			UserId: merch.UserId}
		if err := m.usecase.UpdateUserRole(updateRequest); err != nil {
			return err
		}
	}
	return nil
}
