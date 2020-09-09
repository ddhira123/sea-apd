package merchant

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/dto/domain"
	"github.com/williamchang80/sea-apd/dto/response/base"
	response "github.com/williamchang80/sea-apd/dto/response/merchant"
	"net/http"
)

type MerchantController struct {
	usecase merchant.MerchantUsecase
}

func NewMerchantController(e *echo.Echo, m merchant.MerchantUsecase) merchant.MerchantController {
	c := &MerchantController{usecase: m}
	e.GET("/api/merchant/balance", c.GetMerchantBalance)
	return c
}

func (m *MerchantController) GetMerchantBalance(e echo.Context) error {
	merchantId := e.QueryParam("merchantId")
	balance, err := m.usecase.GetMerchantBalance(merchantId)
	if err != nil {
		return e.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return e.JSON(http.StatusOK, &response.GetMerchantBalanceResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		}, Data: domain.MerchantBalanceDto{
			Balance: balance,
		},
	})
}
