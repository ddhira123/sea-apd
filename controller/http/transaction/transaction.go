package transaction

import (
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/domain/transaction"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
	"github.com/williamchang80/sea-apd/dto/response/base"
	"net/http"
)

type TransactionController struct {
	usecase transaction.TransactionUsecase
}

func NewTransactionController(e *echo.Echo, t transaction.TransactionUsecase) transaction.TransactionController {
	c := &TransactionController{usecase: t}
	e.POST("/api/transaction", c.CreateTransaction)
	e.POST("/api/transaction/status", c.UpdateTransactionStatus)
	return c
}

func (t *TransactionController) CreateTransaction(c echo.Context) error {
	var request transaction2.TransactionRequest
	c.Bind(&request)
	err := t.usecase.CreateTransaction(request)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t *TransactionController) UpdateTransactionStatus(c echo.Context) error {
	var request transaction2.UpdateTransactionRequest
	c.Bind(&request)
	err := t.usecase.UpdateTransactionStatus(request)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, &base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}
