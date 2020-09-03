package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

func Connection() (*gorm.DB, sqlmock.Sqlmock) {
	const driver = "postgres"
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	g,_ := gorm.Open(driver, db)
	return g, mock
}
