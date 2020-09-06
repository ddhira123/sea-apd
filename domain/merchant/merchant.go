package merchant

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

type Merchant struct {
	domain.Base
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	UserId  string `json:"user_id"`
}

type MerchantRepository interface {
	UpdateMerchantBalance(amount int, merchantId string) error
	GetMerchantBalance(merchantId string) (int, error)
}

type MerchantUsecase interface {
	UpdateMerchantBalance(request merchant.UpdateMerchantBalanceRequest) error
	GetMerchantBalance(merchantId string) (int, error)
}
type MerchantController interface {
	UpdateMerchantBalance(echo echo.Context) error
	GetMerchantBalance(echo echo.Context) error
}
