package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func Connection() (*gorm.DB, sqlmock.Sqlmock) {
	const driver = "postgres"
	db, mock, _ := sqlmock.New()
	g,_ := gorm.Open(driver, db)
	return g, mock
}
