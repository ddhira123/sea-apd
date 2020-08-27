package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	"github.com/williamchang80/sea-apd/usecase"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	db := db.Postgres()
	k := postgres.NewProductRepository(db)
	t := usecase.NewProductUseCaseImpl(k)
	product.NewProductController(r, t)
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, r)
}
