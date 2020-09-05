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
	panic("implement me")
}