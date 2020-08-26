package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	product2 "github.com/williamchang80/sea-apd/usecase/product"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot read .env file")
	}
	r := mux.NewRouter()
	product.NewProductController(r, product2.NewProductUseCaseImpl())
	db.Postgres()
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, r)
}
