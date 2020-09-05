package merchant

import (
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain/merchant"
)

type MerchantController struct {
	usecase merchant.MerchantUsecase
}

func NewMerchantController(e *echo.Echo, m merchant.MerchantUsecase) merchant.MerchantController {
	c := &MerchantController{usecase: m}
	return c
}
