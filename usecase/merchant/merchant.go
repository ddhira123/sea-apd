package merchant

import (
	"github.com/williamchang80/sea-apd/common/constants/merchant_status"
	"github.com/williamchang80/sea-apd/domain/merchant"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
)

var merchantStatus = merchant_status.GetMerchantStatus()

type MerchantUsecase struct {
	mc merchant.MerchantRepository
}

func NewMerchantUsecase(m merchant.MerchantRepository) merchant.MerchantUsecase {
	mc := MerchantUsecase{mc: m}
	return mc
}

func ConvertRequest(m request.MerchantRequest) merchant.Merchant {
	return merchant.Merchant{
		Name:     m.Name,
		Balance:  m.Balance,
		UserId:   m.UserId,
		Brand:    m.Brand,
		Address:  m.Address,
		Approval: merchantStatus["WAITING"],
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
	merch := ConvertRequest(request)
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

func (m MerchantUsecase) GetMerchantsByUser(userId string) ([]merchant.Merchant, error) {
	mhs, err := m.mc.GetMerchantsByUser(userId)
	if err != nil {
		return nil, err
	}
	return mhs, nil
}
