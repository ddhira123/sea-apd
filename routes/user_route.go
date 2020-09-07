package routes

import (
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
)

func NewUserRoute(e *echo.Echo) {
	db := db.Postgres()
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}
}
