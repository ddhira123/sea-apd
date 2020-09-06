package merchant

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/dto/domain"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
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
	e.PUT("/api/merchant/balance", c.UpdateMerchantBalance)
	return c
}

func (m *MerchantController) UpdateMerchantBalance(e echo.Context) error {
	var updateMerchantBalanceRequest request.UpdateMerchantBalanceRequest
	e.Bind(&updateMerchantBalanceRequest)
	if err := m.usecase.UpdateMerchantBalance(updateMerchantBalanceRequest); err != nil {
		return e.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return e.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
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
