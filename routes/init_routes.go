package routes

import (
	"github.com/labstack/echo"
)

type Routes struct {
	controller interface{}
	usecase    interface{}
	repository interface{}
}

func InitMainRoutes(echo *echo.Echo) {
	NewMerchantRoute(echo)
	NewProductRoutes(echo)
	NewTransactionRoute(echo)
	NewUserRoute(echo)
	NewTransferRoute(echo)
}
