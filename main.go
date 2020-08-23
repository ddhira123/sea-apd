package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot read .env file")
	}
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, nil)
}
