package transfer

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/transfer"
	request "github.com/williamchang80/sea-apd/dto/request/transfer"
	"github.com/williamchang80/sea-apd/dto/response/base"
	transfer2 "github.com/williamchang80/sea-apd/dto/response/transfer"
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
	userId := ctx.QueryParam("merchantId")
	transfers, err := t.usecase.GetTransferHistory(userId)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, &base.BaseResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: message.UNPROCESSABLE_ENTITY,
		})
	}
	return ctx.JSON(http.StatusOK, &transfer2.GetTransferResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: transfers,
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
