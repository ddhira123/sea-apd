package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository"
	product2 "github.com/williamchang80/sea-apd/usecase/product"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	db := db.Postgres()
	k := repository.NewProductRepository(db)
	t := product2.NewProductUseCaseImpl(k)
	product.NewProductController(r, t)
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, r)
}
