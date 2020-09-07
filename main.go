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