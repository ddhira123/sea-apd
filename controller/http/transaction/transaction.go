package transaction

import (
	"fmt"
	"github.com/labstack/echo"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/transaction"
	"github.com/williamchang80/sea-apd/dto/domain"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
	"github.com/williamchang80/sea-apd/dto/response/base"
	response "github.com/williamchang80/sea-apd/dto/response/transaction"
	"net/http"
)

type TransactionController struct {
	usecase transaction.TransactionUsecase
}



func NewTransactionController(e *echo.Echo, t transaction.TransactionUsecase) transaction.TransactionController {
	c := &TransactionController{usecase: t}
	e.POST("/api/transaction/status", c.UpdateTransactionStatus)
	e.GET("/api/transaction", c.GetTransactionById)
	e.GET("/api/transactions/history", c.GetTransactionHistory)
	e.GET("/api/transactions/request", c.GetMerchantRequestItem)
	e.POST("/api/transaction/payment", c.PayTransaction)
	e.POST("/api/cart", c.CreateCart)
	e.POST("/api/cart/item", c.AddCartItem)
	e.DELETE("/api/cart/item", c.RemoveCartItem)
	e.PUT("/api/cart/item", c.UpdateCartItem)
	e.GET("/api/cart/item", c.GetCartItems)
	return c
}

func (t *TransactionController) UpdateTransactionStatus(c echo.Context) error {
	var request transaction2.UpdateTransactionRequest
	c.Bind(&request)
	fmt.Println(transaction_status.ToString(request.Status))
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

func (t *TransactionController) GetTransactionById(c echo.Context) error {
	id := c.QueryParam("transactionId")
	tr, err := t.usecase.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, response.GetTransactionByIdResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.TransactionDto{
			Transaction: *tr,
		},
	})
}

func (t *TransactionController) GetTransactionHistory(c echo.Context) error {
	id := c.QueryParam("userId")
	tr, err := t.usecase.GetTransactionHistory(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, response.GetTransactionHistoryResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.TransactionListDto{
			Transactions: tr,
		},
	})
}

func (t *TransactionController) GetMerchantRequestItem(c echo.Context) error {
	id := c.QueryParam("merchantId")
	tr, err := t.usecase.GetMerchantRequestItem(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, response.GetTransactionHistoryResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.TransactionListDto{
			Transactions: tr,
		},
	})
}

func (t *TransactionController) PayTransaction(c echo.Context) error {
	var request transaction2.PaymentRequest
	c.Bind(&request)
	err := t.usecase.PayTransaction(request)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t *TransactionController) AddCartItem(c echo.Context) error {
	var request transaction2.CartRequest
	c.Bind(&request)
	err := t.usecase.AddCartItem(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t *TransactionController) RemoveCartItem(c echo.Context) error {
	var request transaction2.CartRequest
	c.Bind(&request)
	err := t.usecase.RemoveCartItem(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t *TransactionController) UpdateCartItem(c echo.Context) error {
	var request transaction2.CartRequest
	c.Bind(&request)
	err := t.usecase.UpdateCartItem(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &base.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, base.BaseResponse{
		Code:    http.StatusOK,
		Message: message.SUCCESS,
	})
}

func (t *TransactionController) GetCartItems(c echo.Context) error {
	id := c.QueryParam("transactionId")
	carts, err := t.usecase.GetCartItems(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, &base.BaseResponse{
			Code:    http.StatusNotFound,
			Message: message.NOT_FOUND,
		})
	}
	return c.JSON(http.StatusOK, &response.GetCartItemsResponse{
		BaseResponse: base.BaseResponse{
			Code:    http.StatusOK,
			Message: message.SUCCESS,
		},
		Data: domain.CartDto{
			CartItems: carts,
		},
	})
}

func (t *TransactionController) CreateCart(c echo.Context) error {
	var request transaction2.CreateCartRequest
	c.Bind(&request)
	err := t.usecase.CreateCart(request)
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