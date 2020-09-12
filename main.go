// @title SEA-APD Store API
// @version 1.0
// @description This is an API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	routes.InitMainRoutes(e)
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	log.Panic(http.ListenAndServe(appPort, e))
}
