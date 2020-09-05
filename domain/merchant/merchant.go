package merchant

import (
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

type Merchant struct {
	domain.Base
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type MerchantRepository interface {
	UpdateMerchantBalance(amount int, merchantId string) error
}

type MerchantUsecase interface {
	UpdateMerchantBalance(request merchant.UpdateMerchantBalanceRequest) error
}

type MerchantController interface {
}
