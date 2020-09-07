package transfer

import (
"github.com/labstack/echo"
message "github.com/williamchang80/sea-apd/common/constants/response"
"github.com/williamchang80/sea-apd/domain/transfer"
request "github.com/williamchang80/sea-apd/dto/request/transfer"
"github.com/williamchang80/sea-apd/dto/response/base"
"net/http"
)

type TransferController struct {
	usecase transfer.TransferUsecase
}

func NewTransferController(e *echo.Echo, t transfer.TransferUsecase) transfer.TransferController {
	c := &TransferController{usecase: t}
	e.POST("api/transfer", c.CreateTransferHistory)
	e.GET("api/transfers", c.GetTransferHistory)
	return c
}

func (t TransferController) GetTransferHistory(ctx echo.Context) error {
	merchantId := ctx.QueryParam("merchantId")
	_, err := t.usecase.GetTransferHistory(merchantId);
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, &base.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: message.UNPROCESSABLE_ENTITY,
		})
	}
	return ctx.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t TransferController) CreateTransferHistory(ctx echo.Context) error {
	var request request.CreateTransferHistoryRequest
	ctx.Bind(&request)
	if err := t.usecase.CreateTransferHistory(request); err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, &base.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: message.UNPROCESSABLE_ENTITY,
		})
	}
	return ctx.JSON(http.StatusCreated, &base.BaseResponse{
		Code:    http.StatusCreated,
		Message: message.SUCCESS,
	})
}
