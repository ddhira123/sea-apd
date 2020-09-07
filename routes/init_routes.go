package routes

import "github.com/labstack/echo"

func InitMainRoutes(echo *echo.Echo) {
	NewProductRoutes(echo)
	NewTransactionRoute(echo)
	NewUserRoutes(echo)
}
