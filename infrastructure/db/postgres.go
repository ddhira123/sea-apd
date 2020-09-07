package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const driver = "postgres"

func Postgres() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(driver, psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}
