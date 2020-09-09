package routes

import (
	"github.com/labstack/echo"
	middleware2 "github.com/labstack/echo/middleware"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/common/auth"
	"github.com/williamchang80/sea-apd/dto/response/base"
	"net/http"
	"os"
	"strings"
)

type Routes struct {
	controller interface{}
	usecase    interface{}
	repository interface{}
}

func InitMainRoutes(echo *echo.Echo) {
	NewUserRoute(echo)
	NewMerchantRoute(echo)
	NewProductRoutes(echo)
	NewTransactionRoute(echo)
	NewTransferRoute(echo)
	NewAuthRoute(echo)

	InitMiddleware(echo)
}

func InitMiddleware(e *echo.Echo) {
	key := os.Getenv("SECRET_AUTH_KEY")
	e.Use(middleware2.JWTWithConfig(middleware2.JWTConfig{
		SigningKey:  []byte(key),
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
		Skipper: func(context echo.Context) bool {
			if strings.HasPrefix(context.Request().URL.Path, "/api/auth") {
				return true
			}
			return false
		},
		BeforeFunc: func(e echo.Context) {
			if !auth.IsValidTokenLifetime(e.Request().Header.Get("Authorization")) {
				e.JSON(http.StatusUnauthorized, base.BaseResponse{
					Code:    http.StatusUnauthorized,
					Message: message.UNAUTHORIZED,
				})
			}
		},
	}))
}
