package merchant

import (
	"net/http"

	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/dto/domain"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
	"github.com/williamchang80/sea-apd/dto/response/base"
	response "github.com/williamchang80/sea-apd/dto/response/merchant"
)

type MerchantController struct {
	usecase merchant.MerchantUsecase
}

func NewMerchantController(e *echo.Echo, m merchant.MerchantUsecase) merchant.MerchantController {
	c := &MerchantController{usecase: m}
	e.GET("/api/merchant/balance", c.GetMerchantBalance)
	e.POST("/api/merchant", c.RegisterMerchant)
	e.GET("/api/merchant", c.GetMerchantById)
	e.GET("/api/merchants", c.GetMerchants)
	e.PUT("/api/merchant/status", c.UpdateMerchantApprovalStatus)
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

func (m *MerchantController) RegisterMerchant(c echo.Context) error {
	var merchantRequest request.MerchantRequest
	c.Bind(&merchantRequest)

	if err := m.usecase.RegisterMerchant(merchantRequest); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: message.BAD_REQUEST,
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: message.SUCCESS,
	})
}

func (m *MerchantController) GetMerchantById(context echo.Context) error {
	id := context.QueryParam("merchantId")
	merch, err := m.usecase.GetMerchantById(id)
	if err != nil {
		return context.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return context.JSON(http.StatusOK, &response.GetMerchantByIdResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.MerchantDto{
			Merchant: merch,
		},
	})
}

func (m *MerchantController) GetMerchants(c echo.Context) error {
	merchants, err := m.usecase.GetMerchants()
	if err != nil {
		c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}

	return c.JSON(http.StatusOK, &response.GetMerchantsResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.MerchantListDto{Merchants: merchants},
	})
}

func (m *MerchantController) UpdateMerchantApprovalStatus(c echo.Context) error {
	var request request.UpdateMerchantApprovalStatusRequest
	c.Bind(&request)

	if err := m.usecase.UpdateMerchantApprovalStatus(request); err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: message.SUCCESS,
	})
}
