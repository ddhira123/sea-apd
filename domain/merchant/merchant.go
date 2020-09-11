package merchant

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request/merchant"
)

type Merchant struct {
	domain.Base
	Name     string `json:"name"`
	Balance  int    `json:"balance"`
	UserId   string `json:"user_id"`
	Brand    string `json:"brand"`
	Address  string `json:"address"`
	Approval string `json:"approval"`
}

type MerchantRepository interface {
	UpdateMerchantBalance(amount int, merchantId string) error
	GetMerchantBalance(merchantId string) (int, error)
	RegisterMerchant(merchant Merchant) error
	GetMerchants() ([]Merchant, error)
	GetMerchantById(merchantId string) (*Merchant, error)
	UpdateMerchantApprovalStatus(merchantId string, status string) error
}

type MerchantUsecase interface {
	UpdateMerchantBalance(request merchant.UpdateMerchantBalanceRequest) error
	GetMerchantBalance(merchantId string) (int, error)
	RegisterMerchant(request merchant.MerchantRequest) error
	GetMerchants() ([]Merchant, error)
	GetMerchantById(merchantId string) (*Merchant, error)
	UpdateMerchantApprovalStatus(request merchant.UpdateMerchantApprovalStatusRequest) error
}
type MerchantController interface {
	GetMerchantBalance(echo echo.Context) error
	GetMerchants(echo echo.Context) error
	GetMerchantById(echo echo.Context) error
	RegisterMerchant(echo echo.Context) error
	UpdateMerchantApprovalStatus(echo echo.Context) error
}
