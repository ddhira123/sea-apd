package routes

import (
	"github.com/labstack/echo"
	v4 "github.com/labstack/echo/v4"
	"github.com/williamchang80/sea-apd/docs"
	"log"
	"net/http"
	"os"
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

	serveSwaggerUI()
}

func serveSwaggerUI() {
	e := v4.New()
	e.GET("/docs/*", docs.WrapHandler)
	go func() {
		swaggerPort := ":" + os.Getenv("SWAGGER_PORT")
		log.Panic(http.ListenAndServe(swaggerPort, e))
	}()
}
