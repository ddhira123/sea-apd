package merchant

import (
	"github.com/williamchang80/sea-apd/domain/merchant"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
)

type MerchantUsecase struct {
	mc merchant.MerchantRepository
}

func NewMerchantUsecase(m merchant.MerchantRepository) merchant.MerchantUsecase {
	mc := MerchantUsecase{mc: m}
	return mc
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
