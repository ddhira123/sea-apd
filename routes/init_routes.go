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
	NewProductRoutes(echo)
	NewTransactionRoute(echo)
	NewMerchantRoute(echo)
}
