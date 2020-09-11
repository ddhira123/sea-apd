package routes

import (
	"github.com/labstack/echo"
	middleware2 "github.com/labstack/echo/middleware"
	"github.com/williamchang80/sea-apd/common/auth"
	message "github.com/williamchang80/sea-apd/common/constants/response"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/dto/response/base"
	"net/http"
	"os"
	"strings"
)

type Routes struct {
	Controller interface{}
	Usecase    interface{}
	Repository interface{}
}

func InitMainRoutes(echo *echo.Echo) {
	NewUserRoute(echo)
	NewMerchantRoute(echo)
	NewProductRoutes(echo)
	NewTransactionRoute(echo)
	NewTransferRoute(echo)
	NewAuthRoute(echo)

	mailer.InitMail()
	//InitMiddleware(echo)
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
