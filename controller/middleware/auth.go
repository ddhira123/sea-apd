package middleware

import (
	"crypto/subtle"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

var (
	u = os.Getenv("BASIC_AUTH_USERNAME")
	p = os.Getenv("BASIC_AUTH_PASSWORD")
)

// Init ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// BasicAuthUser ...
func BasicAuthUser(username, password string, c echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte(u)) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte(p)) == 1 {
		return true, nil
	}
	return false, nil
}